// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterListRequest
	GetClusterId() *string
	SetClusterName(v string) *DescribeClusterListRequest
	GetClusterName() *string
	SetDbType(v string) *DescribeClusterListRequest
	GetDbType() *string
	SetOwnerId(v int64) *DescribeClusterListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeClusterListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeClusterListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeClusterListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeClusterListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeClusterListRequest
	GetResourceOwnerId() *int64
	SetStatusList(v []*string) *DescribeClusterListRequest
	GetStatusList() []*string
	SetTag(v []*DescribeClusterListRequestTag) *DescribeClusterListRequest
	GetTag() []*DescribeClusterListRequestTag
	SetZoneId(v string) *DescribeClusterListRequest
	GetZoneId() *string
}

type DescribeClusterListRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	DbType      *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                          `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                           `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StatusList           []*string                        `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	Tag                  []*DescribeClusterListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ZoneId               *string                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterListRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterListRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterListRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeClusterListRequest) GetDbType() *string {
	return s.DbType
}

func (s *DescribeClusterListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeClusterListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClusterListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClusterListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeClusterListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeClusterListRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *DescribeClusterListRequest) GetTag() []*DescribeClusterListRequestTag {
	return s.Tag
}

func (s *DescribeClusterListRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeClusterListRequest) SetClusterId(v string) *DescribeClusterListRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterListRequest) SetClusterName(v string) *DescribeClusterListRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterListRequest) SetDbType(v string) *DescribeClusterListRequest {
	s.DbType = &v
	return s
}

func (s *DescribeClusterListRequest) SetOwnerId(v int64) *DescribeClusterListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClusterListRequest) SetPageNumber(v int32) *DescribeClusterListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterListRequest) SetPageSize(v int32) *DescribeClusterListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterListRequest) SetRegionId(v string) *DescribeClusterListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterListRequest) SetResourceOwnerAccount(v string) *DescribeClusterListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClusterListRequest) SetResourceOwnerId(v int64) *DescribeClusterListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterListRequest) SetStatusList(v []*string) *DescribeClusterListRequest {
	s.StatusList = v
	return s
}

func (s *DescribeClusterListRequest) SetTag(v []*DescribeClusterListRequestTag) *DescribeClusterListRequest {
	s.Tag = v
	return s
}

func (s *DescribeClusterListRequest) SetZoneId(v string) *DescribeClusterListRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeClusterListRequest) Validate() error {
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

type DescribeClusterListRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeClusterListRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterListRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeClusterListRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeClusterListRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeClusterListRequestTag) SetKey(v string) *DescribeClusterListRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeClusterListRequestTag) SetValue(v string) *DescribeClusterListRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeClusterListRequestTag) Validate() error {
	return dara.Validate(s)
}
