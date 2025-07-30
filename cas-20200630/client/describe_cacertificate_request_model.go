// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCACertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifier(v string) *DescribeCACertificateRequest
	GetIdentifier() *string
}

type DescribeCACertificateRequest struct {
	// The unique identifier of the CA certificate that you want to query.
	//
	// >  You can call the [DescribeCACertificateList](https://help.aliyun.com/document_detail/328095.html) operation to query the unique identifiers of all CA certificates.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s DescribeCACertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificateRequest) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *DescribeCACertificateRequest) SetIdentifier(v string) *DescribeCACertificateRequest {
	s.Identifier = &v
	return s
}

func (s *DescribeCACertificateRequest) Validate() error {
	return dara.Validate(s)
}
