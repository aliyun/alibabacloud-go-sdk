// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptAgreementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *AcceptAgreementRequest
	GetJwtToken() *string
}

type AcceptAgreementRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s AcceptAgreementRequest) String() string {
	return dara.Prettify(s)
}

func (s AcceptAgreementRequest) GoString() string {
	return s.String()
}

func (s *AcceptAgreementRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *AcceptAgreementRequest) SetJwtToken(v string) *AcceptAgreementRequest {
	s.JwtToken = &v
	return s
}

func (s *AcceptAgreementRequest) Validate() error {
	return dara.Validate(s)
}
