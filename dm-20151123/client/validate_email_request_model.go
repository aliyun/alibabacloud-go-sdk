// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *ValidateEmailRequest
	GetEmail() *string
	SetTimeout(v int64) *ValidateEmailRequest
	GetTimeout() *int64
}

type ValidateEmailRequest struct {
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

func (s *ValidateEmailRequest) GetEmail() *string {
	return s.Email
}

func (s *ValidateEmailRequest) GetTimeout() *int64 {
	return s.Timeout
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
