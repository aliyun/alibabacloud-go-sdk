// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckGraylist(v bool) *ValidateEmailRequest
	GetCheckGraylist() *bool
	SetEmail(v string) *ValidateEmailRequest
	GetEmail() *string
	SetTimeout(v int64) *ValidateEmailRequest
	GetTimeout() *int64
}

type ValidateEmailRequest struct {
	CheckGraylist *bool `json:"CheckGraylist,omitempty" xml:"CheckGraylist,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx@yyy.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 20
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ValidateEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateEmailRequest) GoString() string {
	return s.String()
}

func (s *ValidateEmailRequest) GetCheckGraylist() *bool {
	return s.CheckGraylist
}

func (s *ValidateEmailRequest) GetEmail() *string {
	return s.Email
}

func (s *ValidateEmailRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *ValidateEmailRequest) SetCheckGraylist(v bool) *ValidateEmailRequest {
	s.CheckGraylist = &v
	return s
}

func (s *ValidateEmailRequest) SetEmail(v string) *ValidateEmailRequest {
	s.Email = &v
	return s
}

func (s *ValidateEmailRequest) SetTimeout(v int64) *ValidateEmailRequest {
	s.Timeout = &v
	return s
}

func (s *ValidateEmailRequest) Validate() error {
	return dara.Validate(s)
}
