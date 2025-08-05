// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *DescribeClientsRequest
	GetClientId() *string
	SetClientType(v string) *DescribeClientsRequest
	GetClientType() *string
	SetClusterId(v string) *DescribeClientsRequest
	GetClusterId() *string
	SetPageNumber(v int32) *DescribeClientsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeClientsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeClientsRequest
	GetResourceGroupId() *string
	SetSourceType(v string) *DescribeClientsRequest
	GetSourceType() *string
	SetVaultId(v string) *DescribeClientsRequest
	GetVaultId() *string
}

type DescribeClientsRequest struct {
	// The ID of the Cloud Backup client.
	//
	// example:
	//
	// c-000ed600******6b0
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The type of the Cloud Backup client. Valid value: **ECS_AGENT**, which indicates an SAP HANA backup client.
	//
	// example:
	//
	// ECS_AGENT
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The ID of the SAP HANA instance.
	//
	// example:
	//
	// cl-0005ni1******2l87
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmw6bxl7o5qyq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the data source. Valid value:**HANA**, which indicates SAP HANA backup.
	//
	// example:
	//
	// HANA
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-0001vk0z******xdyr
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientsRequest) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeClientsRequest) GetClientType() *string {
	return s.ClientType
}

func (s *DescribeClientsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClientsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClientsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClientsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeClientsRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeClientsRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeClientsRequest) SetClientId(v string) *DescribeClientsRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeClientsRequest) SetClientType(v string) *DescribeClientsRequest {
	s.ClientType = &v
	return s
}

func (s *DescribeClientsRequest) SetClusterId(v string) *DescribeClientsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClientsRequest) SetPageNumber(v int32) *DescribeClientsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClientsRequest) SetPageSize(v int32) *DescribeClientsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClientsRequest) SetResourceGroupId(v string) *DescribeClientsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClientsRequest) SetSourceType(v string) *DescribeClientsRequest {
	s.SourceType = &v
	return s
}

func (s *DescribeClientsRequest) SetVaultId(v string) *DescribeClientsRequest {
	s.VaultId = &v
	return s
}

func (s *DescribeClientsRequest) Validate() error {
	return dara.Validate(s)
}
