// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterShareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *AlterShareRequest
	GetComment() *string
	SetEnableWrite(v bool) *AlterShareRequest
	GetEnableWrite() *bool
	SetShareName(v string) *AlterShareRequest
	GetShareName() *string
}

type AlterShareRequest struct {
	// example:
	//
	// description
	Comment     *string `json:"comment,omitempty" xml:"comment,omitempty"`
	EnableWrite *bool   `json:"enableWrite,omitempty" xml:"enableWrite,omitempty"`
	// example:
	//
	// share_name
	ShareName *string `json:"shareName,omitempty" xml:"shareName,omitempty"`
}

func (s AlterShareRequest) String() string {
	return dara.Prettify(s)
}

func (s AlterShareRequest) GoString() string {
	return s.String()
}

func (s *AlterShareRequest) GetComment() *string {
	return s.Comment
}

func (s *AlterShareRequest) GetEnableWrite() *bool {
	return s.EnableWrite
}

func (s *AlterShareRequest) GetShareName() *string {
	return s.ShareName
}

func (s *AlterShareRequest) SetComment(v string) *AlterShareRequest {
	s.Comment = &v
	return s
}

func (s *AlterShareRequest) SetEnableWrite(v bool) *AlterShareRequest {
	s.EnableWrite = &v
	return s
}

func (s *AlterShareRequest) SetShareName(v string) *AlterShareRequest {
	s.ShareName = &v
	return s
}

func (s *AlterShareRequest) Validate() error {
	return dara.Validate(s)
}
