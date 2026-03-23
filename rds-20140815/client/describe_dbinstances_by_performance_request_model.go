// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesByPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTag(v []*DescribeDBInstancesByPerformanceRequestTag) *DescribeDBInstancesByPerformanceRequest
	GetTag() []*DescribeDBInstancesByPerformanceRequestTag
	SetClientToken(v string) *DescribeDBInstancesByPerformanceRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeDBInstancesByPerformanceRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeDBInstancesByPerformanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstancesByPerformanceRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDBInstancesByPerformanceRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstancesByPerformanceRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDBInstancesByPerformanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBInstancesByPerformanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstancesByPerformanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstancesByPerformanceRequest
	GetResourceOwnerId() *int64
	SetSortKey(v string) *DescribeDBInstancesByPerformanceRequest
	GetSortKey() *string
	SetSortMethod(v string) *DescribeDBInstancesByPerformanceRequest
	GetSortMethod() *string
	SetTags(v string) *DescribeDBInstancesByPerformanceRequest
	GetTags() *string
	SetProxyId(v string) *DescribeDBInstancesByPerformanceRequest
	GetProxyId() *string
}

type DescribeDBInstancesByPerformanceRequest struct {
	Tag                  []*DescribeDBInstancesByPerformanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ClientToken          *string                                       `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceId         *string                                       `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string                                       `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string                                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string                                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                        `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SortKey              *string                                       `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
	SortMethod           *string                                       `json:"SortMethod,omitempty" xml:"SortMethod,omitempty"`
	Tags                 *string                                       `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ProxyId              *string                                       `json:"proxyId,omitempty" xml:"proxyId,omitempty"`
}

func (s DescribeDBInstancesByPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesByPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesByPerformanceRequest) GetTag() []*DescribeDBInstancesByPerformanceRequestTag {
	return s.Tag
}

func (s *DescribeDBInstancesByPerformanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDBInstancesByPerformanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesByPerformanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstancesByPerformanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstancesByPerformanceRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancesByPerformanceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstancesByPerformanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesByPerformanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesByPerformanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstancesByPerformanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstancesByPerformanceRequest) GetSortKey() *string {
	return s.SortKey
}

func (s *DescribeDBInstancesByPerformanceRequest) GetSortMethod() *string {
	return s.SortMethod
}

func (s *DescribeDBInstancesByPerformanceRequest) GetTags() *string {
	return s.Tags
}

func (s *DescribeDBInstancesByPerformanceRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *DescribeDBInstancesByPerformanceRequest) SetTag(v []*DescribeDBInstancesByPerformanceRequestTag) *DescribeDBInstancesByPerformanceRequest {
	s.Tag = v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequest) SetClientToken(v string) *DescribeDBInstancesByPerformanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequest) SetDBInstanceId(v string) *DescribeDBInstancesByPerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequest) SetOwnerAccount(v string) *DescribeDBInstancesByPerformanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequest) SetOwnerId(v int64) *DescribeDBInstancesByPerformanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequest) SetPageNumber(v int32) *DescribeDBInstancesByPerformanceRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequest) SetPageSize(v int32) *DescribeDBInstancesByPerformanceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequest) SetRegionId(v string) *DescribeDBInstancesByPerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequest) SetResourceGroupId(v string) *DescribeDBInstancesByPerformanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequest) SetResourceOwnerAccount(v string) *DescribeDBInstancesByPerformanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequest) SetResourceOwnerId(v int64) *DescribeDBInstancesByPerformanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequest) SetSortKey(v string) *DescribeDBInstancesByPerformanceRequest {
	s.SortKey = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequest) SetSortMethod(v string) *DescribeDBInstancesByPerformanceRequest {
	s.SortMethod = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequest) SetTags(v string) *DescribeDBInstancesByPerformanceRequest {
	s.Tags = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequest) SetProxyId(v string) *DescribeDBInstancesByPerformanceRequest {
	s.ProxyId = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequest) Validate() error {
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

type DescribeDBInstancesByPerformanceRequestTag struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeDBInstancesByPerformanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesByPerformanceRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesByPerformanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstancesByPerformanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstancesByPerformanceRequestTag) SetKey(v string) *DescribeDBInstancesByPerformanceRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequestTag) SetValue(v string) *DescribeDBInstancesByPerformanceRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceRequestTag) Validate() error {
	return dara.Validate(s)
}
