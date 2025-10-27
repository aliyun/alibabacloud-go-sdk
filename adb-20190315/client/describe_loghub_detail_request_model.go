// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoghubDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportName(v string) *DescribeLoghubDetailRequest
	GetExportName() *string
	SetOwnerAccount(v string) *DescribeLoghubDetailRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLoghubDetailRequest
	GetOwnerId() *int64
	SetProjectName(v string) *DescribeLoghubDetailRequest
	GetProjectName() *string
	SetRegionId(v string) *DescribeLoghubDetailRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeLoghubDetailRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeLoghubDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLoghubDetailRequest
	GetResourceOwnerId() *int64
}

type DescribeLoghubDetailRequest struct {
	// The name of the log shipping job.
	//
	// This parameter is required.
	//
	// example:
	//
	// wap_log_full_to_adb
	ExportName   *string `json:"ExportName,omitempty" xml:"ExportName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the Simple Log Service project.
	//
	// This parameter is required.
	//
	// example:
	//
	// sls-hz-php-ad
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLoghubDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoghubDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoghubDetailRequest) GetExportName() *string {
	return s.ExportName
}

func (s *DescribeLoghubDetailRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLoghubDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLoghubDetailRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeLoghubDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLoghubDetailRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeLoghubDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLoghubDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLoghubDetailRequest) SetExportName(v string) *DescribeLoghubDetailRequest {
	s.ExportName = &v
	return s
}

func (s *DescribeLoghubDetailRequest) SetOwnerAccount(v string) *DescribeLoghubDetailRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoghubDetailRequest) SetOwnerId(v int64) *DescribeLoghubDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoghubDetailRequest) SetProjectName(v string) *DescribeLoghubDetailRequest {
	s.ProjectName = &v
	return s
}

func (s *DescribeLoghubDetailRequest) SetRegionId(v string) *DescribeLoghubDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoghubDetailRequest) SetResourceGroupId(v string) *DescribeLoghubDetailRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLoghubDetailRequest) SetResourceOwnerAccount(v string) *DescribeLoghubDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoghubDetailRequest) SetResourceOwnerId(v int64) *DescribeLoghubDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLoghubDetailRequest) Validate() error {
	return dara.Validate(s)
}
