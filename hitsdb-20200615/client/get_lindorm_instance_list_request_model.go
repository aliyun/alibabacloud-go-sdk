// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormInstanceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetLindormInstanceListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetLindormInstanceListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *GetLindormInstanceListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetLindormInstanceListRequest
	GetPageSize() *int32
	SetQueryStr(v string) *GetLindormInstanceListRequest
	GetQueryStr() *string
	SetRegionId(v string) *GetLindormInstanceListRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *GetLindormInstanceListRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *GetLindormInstanceListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetLindormInstanceListRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetLindormInstanceListRequest
	GetSecurityToken() *string
	SetServiceType(v string) *GetLindormInstanceListRequest
	GetServiceType() *string
	SetSupportEngine(v int32) *GetLindormInstanceListRequest
	GetSupportEngine() *int32
	SetTag(v []*GetLindormInstanceListRequestTag) *GetLindormInstanceListRequest
	GetTag() []*GetLindormInstanceListRequestTag
}

type GetLindormInstanceListRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of instances to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword contained in the names of Lindorm instances you want to query. Fuzzy queries based on the keyword is supported.
	//
	// example:
	//
	// test
	QueryStr *string `json:"QueryStr,omitempty" xml:"QueryStr,omitempty"`
	// The ID of the region in which the instances that you want to query is located. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/426062.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aek3b63arvg27vi
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The series of instances that you want to query. Valid values:
	//
	// 	- **lindorm**: The instance is a single-zone Lindorm instance.
	//
	// 	- **lindorm_multizone**: The instance is a multi-zone Lindorm instance.
	//
	// 	- **serverless_lindorm**: The instance is a Lindorm Serverless instance.
	//
	// 	- **lindorm_standalone**: The instance is a single-node Lindorm instance.
	//
	// 	- **lts**: The instance is an LTS instance.
	//
	// example:
	//
	// lindorm
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The engine supported by the instances that you want to query. The engines are indicated by different numbers:
	//
	// 	- **1**: LindormSearch.
	//
	// 	- **2**: LindormTSDB
	//
	// 	- **4**: LindormTable
	//
	// 	- **8**: LindormDFS
	//
	// >  The value of this parameter is the sum of all numbers that indicate the engines supported by the instance. For example, if you set the value of this parameter to 15, which is the sum of 1, 2, 4, and 8, this operation queries instances that support all four engines. If you set the value of this parameter to 6, which is the sum of 2 and 4, this operation queries instances that support LindormTSDB and LindormTable.
	//
	// example:
	//
	// 15
	SupportEngine *int32 `json:"SupportEngine,omitempty" xml:"SupportEngine,omitempty"`
	// The list of tags associated with the specified instances.
	Tag []*GetLindormInstanceListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetLindormInstanceListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLindormInstanceListRequest) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetLindormInstanceListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetLindormInstanceListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetLindormInstanceListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetLindormInstanceListRequest) GetQueryStr() *string {
	return s.QueryStr
}

func (s *GetLindormInstanceListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLindormInstanceListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetLindormInstanceListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetLindormInstanceListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetLindormInstanceListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetLindormInstanceListRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetLindormInstanceListRequest) GetSupportEngine() *int32 {
	return s.SupportEngine
}

func (s *GetLindormInstanceListRequest) GetTag() []*GetLindormInstanceListRequestTag {
	return s.Tag
}

func (s *GetLindormInstanceListRequest) SetOwnerAccount(v string) *GetLindormInstanceListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetOwnerId(v int64) *GetLindormInstanceListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetPageNumber(v int32) *GetLindormInstanceListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetPageSize(v int32) *GetLindormInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetQueryStr(v string) *GetLindormInstanceListRequest {
	s.QueryStr = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetRegionId(v string) *GetLindormInstanceListRequest {
	s.RegionId = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetResourceGroupId(v string) *GetLindormInstanceListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetResourceOwnerAccount(v string) *GetLindormInstanceListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetResourceOwnerId(v int64) *GetLindormInstanceListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetSecurityToken(v string) *GetLindormInstanceListRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetServiceType(v string) *GetLindormInstanceListRequest {
	s.ServiceType = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetSupportEngine(v int32) *GetLindormInstanceListRequest {
	s.SupportEngine = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetTag(v []*GetLindormInstanceListRequestTag) *GetLindormInstanceListRequest {
	s.Tag = v
	return s
}

func (s *GetLindormInstanceListRequest) Validate() error {
	return dara.Validate(s)
}

type GetLindormInstanceListRequestTag struct {
	// The key of tag N of the instances you want to query. You can specify 1 to 20 tag keys.
	//
	// > You can specify the keys of multiple tags. For example, you can specify the key of the first tag in the first key-value pair contained in the value of this parameter and specify the key of the second tag in the second key-value pair.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the instances you want to query. You can specify 1 to 20 tag values.
	//
	// > You can specify the values of multiple tags. For example, you can specify the value of the first tag in the first key-value pair contained in the value of this parameter and specify the value of the second tag in the second key-value pair.
	//
	// example:
	//
	// 2.2.18
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetLindormInstanceListRequestTag) String() string {
	return dara.Prettify(s)
}

func (s GetLindormInstanceListRequestTag) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListRequestTag) GetKey() *string {
	return s.Key
}

func (s *GetLindormInstanceListRequestTag) GetValue() *string {
	return s.Value
}

func (s *GetLindormInstanceListRequestTag) SetKey(v string) *GetLindormInstanceListRequestTag {
	s.Key = &v
	return s
}

func (s *GetLindormInstanceListRequestTag) SetValue(v string) *GetLindormInstanceListRequestTag {
	s.Value = &v
	return s
}

func (s *GetLindormInstanceListRequestTag) Validate() error {
	return dara.Validate(s)
}
