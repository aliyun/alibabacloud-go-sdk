// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCACertificateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaStatus(v string) *DescribeCACertificateListRequest
	GetCaStatus() *string
	SetCertType(v string) *DescribeCACertificateListRequest
	GetCertType() *string
	SetCurrentPage(v int32) *DescribeCACertificateListRequest
	GetCurrentPage() *int32
	SetIdentifier(v string) *DescribeCACertificateListRequest
	GetIdentifier() *string
	SetIssuerType(v string) *DescribeCACertificateListRequest
	GetIssuerType() *string
	SetResourceGroupId(v string) *DescribeCACertificateListRequest
	GetResourceGroupId() *string
	SetShowSize(v int32) *DescribeCACertificateListRequest
	GetShowSize() *int32
	SetValidStatus(v string) *DescribeCACertificateListRequest
	GetValidStatus() *string
}

type DescribeCACertificateListRequest struct {
	// Ca status.
	//
	// - issue: inUse.
	//
	// - forbidden: forbidden.
	//
	// - revoke: revoked.
	//
	// example:
	//
	// issue
	CaStatus *string `json:"CaStatus,omitempty" xml:"CaStatus,omitempty"`
	// The type of the certificate. Valid values:
	//
	// - root: rootCA.
	//
	// - subRoot: subCA.
	//
	// - externalCa: import.
	//
	// example:
	//
	// subRoot
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The unique identifier of the CA certificate.
	//
	// >  You can call the [DescribeCACertificateList](https://help.aliyun.com/document_detail/328095.html) operation to query the unique identifiers of all CA certificates.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The CA Issuer Type.
	//
	// - local: Private certificate.
	//
	// - iTrusChina: Compliance CA.
	//
	// - external: External Import.
	//
	// example:
	//
	// local
	IssuerType      *string `json:"IssuerType,omitempty" xml:"IssuerType,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of CA certificates per page. Default value: **20**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// valid time.
	//
	// - valid: means in the valid period.
	//
	// - notValid: means expired.
	//
	// example:
	//
	// valid
	ValidStatus *string `json:"ValidStatus,omitempty" xml:"ValidStatus,omitempty"`
}

func (s DescribeCACertificateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateListRequest) GetCaStatus() *string {
	return s.CaStatus
}

func (s *DescribeCACertificateListRequest) GetCertType() *string {
	return s.CertType
}

func (s *DescribeCACertificateListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCACertificateListRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *DescribeCACertificateListRequest) GetIssuerType() *string {
	return s.IssuerType
}

func (s *DescribeCACertificateListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCACertificateListRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *DescribeCACertificateListRequest) GetValidStatus() *string {
	return s.ValidStatus
}

func (s *DescribeCACertificateListRequest) SetCaStatus(v string) *DescribeCACertificateListRequest {
	s.CaStatus = &v
	return s
}

func (s *DescribeCACertificateListRequest) SetCertType(v string) *DescribeCACertificateListRequest {
	s.CertType = &v
	return s
}

func (s *DescribeCACertificateListRequest) SetCurrentPage(v int32) *DescribeCACertificateListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCACertificateListRequest) SetIdentifier(v string) *DescribeCACertificateListRequest {
	s.Identifier = &v
	return s
}

func (s *DescribeCACertificateListRequest) SetIssuerType(v string) *DescribeCACertificateListRequest {
	s.IssuerType = &v
	return s
}

func (s *DescribeCACertificateListRequest) SetResourceGroupId(v string) *DescribeCACertificateListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCACertificateListRequest) SetShowSize(v int32) *DescribeCACertificateListRequest {
	s.ShowSize = &v
	return s
}

func (s *DescribeCACertificateListRequest) SetValidStatus(v string) *DescribeCACertificateListRequest {
	s.ValidStatus = &v
	return s
}

func (s *DescribeCACertificateListRequest) Validate() error {
	return dara.Validate(s)
}
