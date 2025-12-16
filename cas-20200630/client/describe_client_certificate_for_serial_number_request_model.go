// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientCertificateForSerialNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSerialNumber(v string) *DescribeClientCertificateForSerialNumberRequest
	GetSerialNumber() *string
}

type DescribeClientCertificateForSerialNumberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 084bde9cd233f0ddae33adc438cfbbbd****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s DescribeClientCertificateForSerialNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateForSerialNumberRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateForSerialNumberRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeClientCertificateForSerialNumberRequest) SetSerialNumber(v string) *DescribeClientCertificateForSerialNumberRequest {
	s.SerialNumber = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberRequest) Validate() error {
	return dara.Validate(s)
}
