// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoReissueFlag(v int32) *ListInstancesRequest
	GetAutoReissueFlag() *int32
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
	SetServerDeployFlag(v int32) *ListInstancesRequest
	GetServerDeployFlag() *int32
	SetShowSize(v int32) *ListInstancesRequest
	GetShowSize() *int32
	SetStatus(v string) *ListInstancesRequest
	GetStatus() *string
}

type ListInstancesRequest struct {
	// Specifies whether the instance is managed. Valid values:
	//
	// - 1: Managed.
	//
	// - 0: Not managed.
	//
	// example:
	//
	// 1
	AutoReissueFlag *int32 `json:"AutoReissueFlag,omitempty" xml:"AutoReissueFlag,omitempty"`
	// The CA brand. Valid values: WoSign, CFCA, DigiCert, GeoTrust, GlobalSign, vTrus, and Alibaba.
	//
	// example:
	//
	// Digicert
	Brand *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	// The status of the certificate. Valid values:
	//
	// - **issued**: Issued.
	//
	// - **revoked**: Revoked.
	//
	// - **willExpire**: About to expire.
	//
	// - **expired**: Expired.
	//
	// example:
	//
	// issued
	CertificateStatus *string `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty"`
	// The type of the certificate. Valid values: DV, OV, and EV.
	//
	// example:
	//
	// DV
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The page number of the current page in a paged query. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The instance type. Valid values:
	//
	// - BUY: formal certificate.
	//
	// - TEST: test certificate.
	//
	// example:
	//
	// BUY
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The keyword for fuzzy search. Matches domain names, instance names, or corresponding resource IDs.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-ae******4wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to return only instances that meet server deployment conditions. Valid values:
	//
	// - 1: Yes.
	//
	// - 0: No.
	//
	// example:
	//
	// 1
	ServerDeployFlag *int32 `json:"ServerDeployFlag,omitempty" xml:"ServerDeployFlag,omitempty"`
	// The number of instances to display per page in a paged query. Default value: **10**. Maximum value: **100**.
	//
	// example:
	//
	// 10
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The instance status. Valid values:
	//
	// - **inactive**: Pending use.
	//
	// - **pending**: Under review. The latest certificate is being reviewed.
	//
	// - **willExpire**: The instance is about to expire.
	//
	// - **expired**: The instance has expired.
	//
	// - **refund**: Refunded.
	//
	// - **normal**: Normal.
	//
	// - **closed**: Closed. The instance is unavailable.
	//
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

func (s *ListInstancesRequest) GetAutoReissueFlag() *int32 {
	return s.AutoReissueFlag
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

func (s *ListInstancesRequest) GetServerDeployFlag() *int32 {
	return s.ServerDeployFlag
}

func (s *ListInstancesRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesRequest) SetAutoReissueFlag(v int32) *ListInstancesRequest {
	s.AutoReissueFlag = &v
	return s
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

func (s *ListInstancesRequest) SetServerDeployFlag(v int32) *ListInstancesRequest {
	s.ServerDeployFlag = &v
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
