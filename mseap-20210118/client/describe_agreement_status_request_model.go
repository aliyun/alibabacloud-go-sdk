// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgreementStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgreementCode(v string) *DescribeAgreementStatusRequest
	GetAgreementCode() *string
}

type DescribeAgreementStatusRequest struct {
	// example:
	//
	// 10aa40008e081ad7b1fb50bffc3a70b1
	AgreementCode *string `json:"AgreementCode,omitempty" xml:"AgreementCode,omitempty"`
}

func (s DescribeAgreementStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgreementStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgreementStatusRequest) GetAgreementCode() *string {
	return s.AgreementCode
}

func (s *DescribeAgreementStatusRequest) SetAgreementCode(v string) *DescribeAgreementStatusRequest {
	s.AgreementCode = &v
	return s
}

func (s *DescribeAgreementStatusRequest) Validate() error {
	return dara.Validate(s)
}
