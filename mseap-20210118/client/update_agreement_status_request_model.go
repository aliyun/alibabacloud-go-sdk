// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgreementStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgreementCode(v string) *UpdateAgreementStatusRequest
	GetAgreementCode() *string
}

type UpdateAgreementStatusRequest struct {
	// example:
	//
	// 10aa40008e081ad7b1fb50bffc3a70b1
	AgreementCode *string `json:"AgreementCode,omitempty" xml:"AgreementCode,omitempty"`
}

func (s UpdateAgreementStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgreementStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgreementStatusRequest) GetAgreementCode() *string {
	return s.AgreementCode
}

func (s *UpdateAgreementStatusRequest) SetAgreementCode(v string) *UpdateAgreementStatusRequest {
	s.AgreementCode = &v
	return s
}

func (s *UpdateAgreementStatusRequest) Validate() error {
	return dara.Validate(s)
}
