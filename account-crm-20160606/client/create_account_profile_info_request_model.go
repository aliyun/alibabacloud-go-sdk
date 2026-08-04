// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountProfileInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountJson(v string) *CreateAccountProfileInfoRequest
	GetAccountJson() *string
}

type CreateAccountProfileInfoRequest struct {
	AccountJson *string `json:"AccountJson,omitempty" xml:"AccountJson,omitempty"`
}

func (s CreateAccountProfileInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountProfileInfoRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountProfileInfoRequest) GetAccountJson() *string {
	return s.AccountJson
}

func (s *CreateAccountProfileInfoRequest) SetAccountJson(v string) *CreateAccountProfileInfoRequest {
	s.AccountJson = &v
	return s
}

func (s *CreateAccountProfileInfoRequest) Validate() error {
	return dara.Validate(s)
}
