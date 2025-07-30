// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifier(v string) *DescribeClientCertificateRequest
	GetIdentifier() *string
}

type DescribeClientCertificateRequest struct {
	// The unique identifier of the client certificate or the server certificate that you want to query.
	//
	// >  You can call the [ListClientCertificate](https://help.aliyun.com/document_detail/330884.html) operation to query the unique identifiers of all client certificates and server certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// d3b95700998e47afc4d95f886579****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s DescribeClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *DescribeClientCertificateRequest) SetIdentifier(v string) *DescribeClientCertificateRequest {
	s.Identifier = &v
	return s
}

func (s *DescribeClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
