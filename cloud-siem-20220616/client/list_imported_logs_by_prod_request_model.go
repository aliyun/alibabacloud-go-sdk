// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImportedLogsByProdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudCode(v string) *ListImportedLogsByProdRequest
	GetCloudCode() *string
	SetProdCode(v string) *ListImportedLogsByProdRequest
	GetProdCode() *string
	SetRegionId(v string) *ListImportedLogsByProdRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListImportedLogsByProdRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListImportedLogsByProdRequest
	GetRoleType() *int32
}

type ListImportedLogsByProdRequest struct {
	// The multicloud code. Valid values:
	//
	// - qcloud: Tencent Cloud.
	//
	// - aliyun: Alibaba Cloud.
	//
	// - hcloud: Huawei Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The code of the product.
	//
	// This parameter is required.
	//
	// example:
	//
	// qcloud_waf
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The region where the Data Management hub of threat analysis is deployed. Select the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID that the administrator uses to switch to the perspective of a member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts within the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ListImportedLogsByProdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImportedLogsByProdRequest) GoString() string {
	return s.String()
}

func (s *ListImportedLogsByProdRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ListImportedLogsByProdRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *ListImportedLogsByProdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListImportedLogsByProdRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListImportedLogsByProdRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListImportedLogsByProdRequest) SetCloudCode(v string) *ListImportedLogsByProdRequest {
	s.CloudCode = &v
	return s
}

func (s *ListImportedLogsByProdRequest) SetProdCode(v string) *ListImportedLogsByProdRequest {
	s.ProdCode = &v
	return s
}

func (s *ListImportedLogsByProdRequest) SetRegionId(v string) *ListImportedLogsByProdRequest {
	s.RegionId = &v
	return s
}

func (s *ListImportedLogsByProdRequest) SetRoleFor(v int64) *ListImportedLogsByProdRequest {
	s.RoleFor = &v
	return s
}

func (s *ListImportedLogsByProdRequest) SetRoleType(v int32) *ListImportedLogsByProdRequest {
	s.RoleType = &v
	return s
}

func (s *ListImportedLogsByProdRequest) Validate() error {
	return dara.Validate(s)
}
