// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptOption(v string) *GetConnectionRequest
	GetEncryptOption() *string
}

type GetConnectionRequest struct {
	// The encryption settings. Valid values:
	//
	// 	- PlainText
	//
	// 	- Secret
	//
	// example:
	//
	// PlainText
	EncryptOption *string `json:"EncryptOption,omitempty" xml:"EncryptOption,omitempty"`
}

func (s GetConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionRequest) GetEncryptOption() *string {
	return s.EncryptOption
}

func (s *GetConnectionRequest) SetEncryptOption(v string) *GetConnectionRequest {
	s.EncryptOption = &v
	return s
}

func (s *GetConnectionRequest) Validate() error {
	return dara.Validate(s)
}
