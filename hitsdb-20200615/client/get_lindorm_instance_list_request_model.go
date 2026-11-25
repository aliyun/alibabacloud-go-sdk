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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page for paging in a paged query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword used to perform a fuzzy search by instance name.
	//
	// example:
	//
	// test
	QueryStr *string `json:"QueryStr,omitempty" xml:"QueryStr,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/426062.html) operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek3b63arvg27vi
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The instance type. Valid values:
	//
	// - **lindorm**: Lindorm single-zone instance
	//
	// - **lindorm_multizone**: Lindorm multi-zone instance
	//
	// - **serverless_lindorm**: Lindorm Serverless instance
	//
	// - **lindorm_standalone**: Lindorm single-node instance
	//
	// - **lts**: Lindorm Tunnel Service
	//
	// example:
	//
	// lindorm
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The data engine types supported by the instance. Valid values:
	//
	// - **1**: Search engine
	//
	// - **2**: LindormTSDB
	//
	// - **4**: LindormTable
	//
	// - **8**: File engine
	//
	// > For example, if the value is 15 (15 = 8 + 4 + 2 + 1), the instance supports the search engine, LindormTSDB, LindormTable, and file engine. If the value is 6 (6 = 4 + 2), the instance supports LindormTSDB and LindormTable.
	//
	// example:
	//
	// 6
	SupportEngine *int32 `json:"SupportEngine,omitempty" xml:"SupportEngine,omitempty"`
	// The list of tags.
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
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLindormInstanceListRequestTag struct {
	// The tag key. Valid values of N: 1 to 20.
	//
	// > You can specify multiple tag keys. For example, the first Key specifies the first tag key, and the second Key specifies the second tag key.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: 1 to 20.
	//
	// > You can specify multiple tag values. For example, the first Value specifies the first tag value, and the second Value specifies the second tag value.
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
