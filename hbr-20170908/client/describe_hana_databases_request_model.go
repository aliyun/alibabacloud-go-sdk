// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaDatabasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeHanaDatabasesRequest
	GetClusterId() *string
	SetPageNumber(v int32) *DescribeHanaDatabasesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHanaDatabasesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeHanaDatabasesRequest
	GetResourceGroupId() *string
	SetVaultId(v string) *DescribeHanaDatabasesRequest
	GetVaultId() *string
}

type DescribeHanaDatabasesRequest struct {
	// The ID of the SAP HANA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cl-000hpc******uv14x
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
	// rg-acfmw2azsegupmi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-00063fq******8xjr
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaDatabasesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaDatabasesRequest) GoString() string {
	return s.String()
}

func (s *DescribeHanaDatabasesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeHanaDatabasesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHanaDatabasesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHanaDatabasesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeHanaDatabasesRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeHanaDatabasesRequest) SetClusterId(v string) *DescribeHanaDatabasesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaDatabasesRequest) SetPageNumber(v int32) *DescribeHanaDatabasesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaDatabasesRequest) SetPageSize(v int32) *DescribeHanaDatabasesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaDatabasesRequest) SetResourceGroupId(v string) *DescribeHanaDatabasesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHanaDatabasesRequest) SetVaultId(v string) *DescribeHanaDatabasesRequest {
	s.VaultId = &v
	return s
}

func (s *DescribeHanaDatabasesRequest) Validate() error {
	return dara.Validate(s)
}
