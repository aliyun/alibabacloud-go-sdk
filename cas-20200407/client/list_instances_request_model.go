// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrand(v string) *ListInstancesRequest
	GetBrand() *string
	SetCertificateStatus(v string) *ListInstancesRequest
	GetCertificateStatus() *string
	SetCertificateType(v string) *ListInstancesRequest
	GetCertificateType() *string
	SetCurrentPage(v int32) *ListInstancesRequest
	GetCurrentPage() *int32
	SetInstanceType(v string) *ListInstancesRequest
	GetInstanceType() *string
	SetKeyword(v string) *ListInstancesRequest
	GetKeyword() *string
	SetResourceGroupId(v string) *ListInstancesRequest
	GetResourceGroupId() *string
	SetShowSize(v int32) *ListInstancesRequest
	GetShowSize() *int32
	SetStatus(v string) *ListInstancesRequest
	GetStatus() *string
}

type ListInstancesRequest struct {
	// example:
	//
	// Digicert
	Brand *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	// example:
	//
	// issued
	CertificateStatus *string `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty"`
	// example:
	//
	// DV
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// BUY
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// rg-ae******4wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 10
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// example:
	//
	// inactive
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetBrand() *string {
	return s.Brand
}

func (s *ListInstancesRequest) GetCertificateStatus() *string {
	return s.CertificateStatus
}

func (s *ListInstancesRequest) GetCertificateType() *string {
	return s.CertificateType
}

func (s *ListInstancesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListInstancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListInstancesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesRequest) SetBrand(v string) *ListInstancesRequest {
	s.Brand = &v
	return s
}

func (s *ListInstancesRequest) SetCertificateStatus(v string) *ListInstancesRequest {
	s.CertificateStatus = &v
	return s
}

func (s *ListInstancesRequest) SetCertificateType(v string) *ListInstancesRequest {
	s.CertificateType = &v
	return s
}

func (s *ListInstancesRequest) SetCurrentPage(v int32) *ListInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceType(v string) *ListInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesRequest) SetKeyword(v string) *ListInstancesRequest {
	s.Keyword = &v
	return s
}

func (s *ListInstancesRequest) SetResourceGroupId(v string) *ListInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesRequest) SetShowSize(v int32) *ListInstancesRequest {
	s.ShowSize = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
