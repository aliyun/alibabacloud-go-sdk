// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncModifyAgLoginEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMpk(v string) *AsyncModifyAgLoginEmailRequest
	GetMpk() *string
	SetNewLoginEmail(v string) *AsyncModifyAgLoginEmailRequest
	GetNewLoginEmail() *string
	SetPk(v string) *AsyncModifyAgLoginEmailRequest
	GetPk() *string
}

type AsyncModifyAgLoginEmailRequest struct {
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	NewLoginEmail *string `json:"NewLoginEmail,omitempty" xml:"NewLoginEmail,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s AsyncModifyAgLoginEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncModifyAgLoginEmailRequest) GoString() string {
	return s.String()
}

func (s *AsyncModifyAgLoginEmailRequest) GetMpk() *string {
	return s.Mpk
}

func (s *AsyncModifyAgLoginEmailRequest) GetNewLoginEmail() *string {
	return s.NewLoginEmail
}

func (s *AsyncModifyAgLoginEmailRequest) GetPk() *string {
	return s.Pk
}

func (s *AsyncModifyAgLoginEmailRequest) SetMpk(v string) *AsyncModifyAgLoginEmailRequest {
	s.Mpk = &v
	return s
}

func (s *AsyncModifyAgLoginEmailRequest) SetNewLoginEmail(v string) *AsyncModifyAgLoginEmailRequest {
	s.NewLoginEmail = &v
	return s
}

func (s *AsyncModifyAgLoginEmailRequest) SetPk(v string) *AsyncModifyAgLoginEmailRequest {
	s.Pk = &v
	return s
}

func (s *AsyncModifyAgLoginEmailRequest) Validate() error {
	return dara.Validate(s)
}
