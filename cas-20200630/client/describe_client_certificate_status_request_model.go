// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientCertificateStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifier(v string) *DescribeClientCertificateStatusRequest
	GetIdentifier() *string
}

type DescribeClientCertificateStatusRequest struct {
	// The unique identifiers of the client certificates or server certificates that you want to query. Separate multiple unique identifiers with commas (,).
	//
	// >  You can call the [ListClientCertificate](https://help.aliyun.com/document_detail/330884.html) operation to query the unique identifiers of all client certificates and server certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s DescribeClientCertificateStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateStatusRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *DescribeClientCertificateStatusRequest) SetIdentifier(v string) *DescribeClientCertificateStatusRequest {
	s.Identifier = &v
	return s
}

func (s *DescribeClientCertificateStatusRequest) Validate() error {
	return dara.Validate(s)
}
