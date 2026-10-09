// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignCertificateCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaIdentifier(v string) *AssignCertificateCountRequest
	GetCaIdentifier() *string
	SetCertTotalCount(v int32) *AssignCertificateCountRequest
	GetCertTotalCount() *int32
	SetId(v int64) *AssignCertificateCountRequest
	GetId() *int64
}

type AssignCertificateCountRequest struct {
	// The identifier of the CA certificate.
	//
	// example:
	//
	// 1f0167b4-ee84-XXX-49bc4d39fa68
	CaIdentifier *string `json:"CaIdentifier,omitempty" xml:"CaIdentifier,omitempty"`
	// The total number of certificate records.
	//
	// example:
	//
	// 5
	CertTotalCount *int32 `json:"CertTotalCount,omitempty" xml:"CertTotalCount,omitempty"`
	// The ID of the data source to which the certificate belongs.
	//
	// example:
	//
	// 33285
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AssignCertificateCountRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignCertificateCountRequest) GoString() string {
	return s.String()
}

func (s *AssignCertificateCountRequest) GetCaIdentifier() *string {
	return s.CaIdentifier
}

func (s *AssignCertificateCountRequest) GetCertTotalCount() *int32 {
	return s.CertTotalCount
}

func (s *AssignCertificateCountRequest) GetId() *int64 {
	return s.Id
}

func (s *AssignCertificateCountRequest) SetCaIdentifier(v string) *AssignCertificateCountRequest {
	s.CaIdentifier = &v
	return s
}

func (s *AssignCertificateCountRequest) SetCertTotalCount(v int32) *AssignCertificateCountRequest {
	s.CertTotalCount = &v
	return s
}

func (s *AssignCertificateCountRequest) SetId(v int64) *AssignCertificateCountRequest {
	s.Id = &v
	return s
}

func (s *AssignCertificateCountRequest) Validate() error {
	return dara.Validate(s)
}
