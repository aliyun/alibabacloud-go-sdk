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
	// 证书的类型 。取值：
	//
	// - **CA**：表示CA证书。
	//
	// - **CERT**：表示签发的证书。
	//
	// example:
	//
	// CERT
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The keyword for the query. You can enter a name, domain name, or Subject Alternative Name (SAN) extension. Fuzzy match is supported.
	//
	// example:
	//
	// test_name
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The number of entries to return on each page. Default value: 50.
	//
	// example:
	//
	// 50
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The source of the certificate. Valid values:
	//
	// 	- **upload**: uploaded certificate
	//
	// 	- **aliyun**: Alibaba Cloud certificate
	//
	// example:
	//
	// aliyun
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The status of the certificate. Valid values:
	//
	// 	- **ISSUE**: issued
	//
	// 	- **REVOKE**: revoked
	//
	// example:
	//
	// ISSUE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the certificate repository. You can call the ListCertWarehouse API operation to query the IDs of certificate repositories.
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
