// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCACertificateLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifier(v string) *ListCACertificateLogRequest
	GetIdentifier() *string
}

type ListCACertificateLogRequest struct {
	// The unique identifier of the CA certificate to query.
	//
	// > Call [DescribeCACertificateList](https://help.aliyun.com/document_detail/328095.html) to query the unique identifiers of all CA certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s ListCACertificateLogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCACertificateLogRequest) GoString() string {
	return s.String()
}

func (s *ListCACertificateLogRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListCACertificateLogRequest) SetIdentifier(v string) *ListCACertificateLogRequest {
	s.Identifier = &v
	return s
}

func (s *ListCACertificateLogRequest) Validate() error {
	return dara.Validate(s)
}
