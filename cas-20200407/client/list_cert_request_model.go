// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertType(v string) *ListCertRequest
	GetCertType() *string
	SetCurrentPage(v int64) *ListCertRequest
	GetCurrentPage() *int64
	SetIdentifiers(v string) *ListCertRequest
	GetIdentifiers() *string
	SetKeyWord(v string) *ListCertRequest
	GetKeyWord() *string
	SetShowSize(v int64) *ListCertRequest
	GetShowSize() *int64
	SetSourceType(v string) *ListCertRequest
	GetSourceType() *string
	SetStatus(v string) *ListCertRequest
	GetStatus() *string
	SetWarehouseId(v int64) *ListCertRequest
	GetWarehouseId() *int64
}

type ListCertRequest struct {
	// The certificate type. Valid values:
	//
	// - **CA**: CA certificate
	//
	// - **CERT**: issued certificate
	//
	// example:
	//
	// CERT
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The page number to return. The default value is 1.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// A comma-separated list of certificate identifiers.
	//
	// example:
	//
	// aaa,bbb
	Identifiers *string `json:"Identifiers,omitempty" xml:"Identifiers,omitempty"`
	// The keyword for a fuzzy search by name, domain name, or subject alternative name.
	//
	// example:
	//
	// test_name
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The number of entries per page. The default value is 50.
	//
	// example:
	//
	// 50
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The source of the certificate. Valid values:
	//
	// - **upload**: uploaded certificate
	//
	// - **aliyun**: Alibaba Cloud certificate
	//
	// example:
	//
	// aliyun
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The certificate status. Valid values:
	//
	// - **ISSUE**: issued
	//
	// - **REVOKE**: revoked
	//
	// example:
	//
	// ISSUE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The warehouse ID. You can obtain this ID by calling the [ListCertWarehouse](https://help.aliyun.com/document_detail/453246.html) API.
	//
	// example:
	//
	// 12
	WarehouseId *int64 `json:"WarehouseId,omitempty" xml:"WarehouseId,omitempty"`
}

func (s ListCertRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCertRequest) GoString() string {
	return s.String()
}

func (s *ListCertRequest) GetCertType() *string {
	return s.CertType
}

func (s *ListCertRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListCertRequest) GetIdentifiers() *string {
	return s.Identifiers
}

func (s *ListCertRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *ListCertRequest) GetShowSize() *int64 {
	return s.ShowSize
}

func (s *ListCertRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListCertRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCertRequest) GetWarehouseId() *int64 {
	return s.WarehouseId
}

func (s *ListCertRequest) SetCertType(v string) *ListCertRequest {
	s.CertType = &v
	return s
}

func (s *ListCertRequest) SetCurrentPage(v int64) *ListCertRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCertRequest) SetIdentifiers(v string) *ListCertRequest {
	s.Identifiers = &v
	return s
}

func (s *ListCertRequest) SetKeyWord(v string) *ListCertRequest {
	s.KeyWord = &v
	return s
}

func (s *ListCertRequest) SetShowSize(v int64) *ListCertRequest {
	s.ShowSize = &v
	return s
}

func (s *ListCertRequest) SetSourceType(v string) *ListCertRequest {
	s.SourceType = &v
	return s
}

func (s *ListCertRequest) SetStatus(v string) *ListCertRequest {
	s.Status = &v
	return s
}

func (s *ListCertRequest) SetWarehouseId(v int64) *ListCertRequest {
	s.WarehouseId = &v
	return s
}

func (s *ListCertRequest) Validate() error {
	return dara.Validate(s)
}
