// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateShareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateShareRequest
	GetComment() *string
	SetEnableWrite(v bool) *CreateShareRequest
	GetEnableWrite() *bool
	SetShareName(v string) *CreateShareRequest
	GetShareName() *string
}

type CreateShareRequest struct {
	// example:
	//
	// demo
	Comment     *string `json:"comment,omitempty" xml:"comment,omitempty"`
	EnableWrite *bool   `json:"enableWrite,omitempty" xml:"enableWrite,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// share_name
	ShareName *string `json:"shareName,omitempty" xml:"shareName,omitempty"`
}

func (s CreateShareRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateShareRequest) GoString() string {
	return s.String()
}

func (s *CreateShareRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateShareRequest) GetEnableWrite() *bool {
	return s.EnableWrite
}

func (s *CreateShareRequest) GetShareName() *string {
	return s.ShareName
}

func (s *CreateShareRequest) SetComment(v string) *CreateShareRequest {
	s.Comment = &v
	return s
}

func (s *CreateShareRequest) SetEnableWrite(v bool) *CreateShareRequest {
	s.EnableWrite = &v
	return s
}

func (s *CreateShareRequest) SetShareName(v string) *CreateShareRequest {
	s.ShareName = &v
	return s
}

func (s *CreateShareRequest) Validate() error {
	return dara.Validate(s)
}
