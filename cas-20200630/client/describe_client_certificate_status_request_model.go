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
	// The unique identifier of the client or server-side certificate that you want to query. Separate multiple identifiers with commas (,).
	//
	// > Call [ListClientCertificate](https://help.aliyun.com/document_detail/465990.html) to query the unique identifiers of all client and server-side certificates.
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
