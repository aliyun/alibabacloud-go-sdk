// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWarehouseCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *DescribeWarehouseCertRequest
	GetCertIdentifier() *string
}

type DescribeWarehouseCertRequest struct {
	// The unique identifier of the certificate.
	//
	// example:
	//
	// 1ef1da5f-38ed-69b3-****-037781890265
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
}

func (s DescribeWarehouseCertRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarehouseCertRequest) GoString() string {
	return s.String()
}

func (s *DescribeWarehouseCertRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DescribeWarehouseCertRequest) SetCertIdentifier(v string) *DescribeWarehouseCertRequest {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeWarehouseCertRequest) Validate() error {
	return dara.Validate(s)
}
