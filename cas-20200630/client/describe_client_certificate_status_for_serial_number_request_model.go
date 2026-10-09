// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientCertificateStatusForSerialNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSerialNumber(v string) *DescribeClientCertificateStatusForSerialNumberRequest
	GetSerialNumber() *string
}

type DescribeClientCertificateStatusForSerialNumberRequest struct {
	// Certificate serial number of the client certificate or server certificate to query. Separate multiple serial numbers with commas (,).
	//
	//
	// > You can call [ListClientCertificate](https://help.aliyun.com/document_detail/330884.html) to query certificate serial numbers of all client certificates and server certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// b67e53ebcea9b77d65b0c3236646d715****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s DescribeClientCertificateStatusForSerialNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateStatusForSerialNumberRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateStatusForSerialNumberRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeClientCertificateStatusForSerialNumberRequest) SetSerialNumber(v string) *DescribeClientCertificateStatusForSerialNumberRequest {
	s.SerialNumber = &v
	return s
}

func (s *DescribeClientCertificateStatusForSerialNumberRequest) Validate() error {
	return dara.Validate(s)
}
