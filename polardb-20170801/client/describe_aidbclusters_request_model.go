// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAiNodeType(v string) *DescribeAIDBClustersRequest
	GetAiNodeType() *string
	SetDBClusterDescription(v string) *DescribeAIDBClustersRequest
	GetDBClusterDescription() *string
	SetDBClusterIds(v string) *DescribeAIDBClustersRequest
	GetDBClusterIds() *string
	SetDBClusterStatus(v string) *DescribeAIDBClustersRequest
	GetDBClusterStatus() *string
	SetOwnerAccount(v string) *DescribeAIDBClustersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAIDBClustersRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAIDBClustersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAIDBClustersRequest
	GetPageSize() *int32
	SetPayType(v string) *DescribeAIDBClustersRequest
	GetPayType() *string
	SetRegionId(v string) *DescribeAIDBClustersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAIDBClustersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAIDBClustersRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeAIDBClustersRequestTag) *DescribeAIDBClustersRequest
	GetTag() []*DescribeAIDBClustersRequestTag
}

type DescribeAIDBClustersRequest struct {
	// example:
	//
	// vnode,container
	AiNodeType *string `json:"AiNodeType,omitempty" xml:"AiNodeType,omitempty"`
	// example:
	//
	// pc-****************
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// example:
	//
	// pc-***************
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                           `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                            `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tag                  []*DescribeAIDBClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAIDBClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClustersRequest) GetAiNodeType() *string {
	return s.AiNodeType
}

func (s *DescribeAIDBClustersRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeAIDBClustersRequest) GetDBClusterIds() *string {
	return s.DBClusterIds
}

func (s *DescribeAIDBClustersRequest) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeAIDBClustersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAIDBClustersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAIDBClustersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAIDBClustersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAIDBClustersRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribeAIDBClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAIDBClustersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAIDBClustersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAIDBClustersRequest) GetTag() []*DescribeAIDBClustersRequestTag {
	return s.Tag
}

func (s *DescribeAIDBClustersRequest) SetAiNodeType(v string) *DescribeAIDBClustersRequest {
	s.AiNodeType = &v
	return s
}

func (s *DescribeAIDBClustersRequest) SetDBClusterDescription(v string) *DescribeAIDBClustersRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeAIDBClustersRequest) SetDBClusterIds(v string) *DescribeAIDBClustersRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeAIDBClustersRequest) SetDBClusterStatus(v string) *DescribeAIDBClustersRequest {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeAIDBClustersRequest) SetOwnerAccount(v string) *DescribeAIDBClustersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAIDBClustersRequest) SetOwnerId(v int64) *DescribeAIDBClustersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAIDBClustersRequest) SetPageNumber(v int32) *DescribeAIDBClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAIDBClustersRequest) SetPageSize(v int32) *DescribeAIDBClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAIDBClustersRequest) SetPayType(v string) *DescribeAIDBClustersRequest {
	s.PayType = &v
	return s
}

func (s *DescribeAIDBClustersRequest) SetRegionId(v string) *DescribeAIDBClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAIDBClustersRequest) SetResourceOwnerAccount(v string) *DescribeAIDBClustersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAIDBClustersRequest) SetResourceOwnerId(v int64) *DescribeAIDBClustersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAIDBClustersRequest) SetTag(v []*DescribeAIDBClustersRequestTag) *DescribeAIDBClustersRequest {
	s.Tag = v
	return s
}

func (s *DescribeAIDBClustersRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeAIDBClustersRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAIDBClustersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClustersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClustersRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeAIDBClustersRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeAIDBClustersRequestTag) SetKey(v string) *DescribeAIDBClustersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeAIDBClustersRequestTag) SetValue(v string) *DescribeAIDBClustersRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeAIDBClustersRequestTag) Validate() error {
	return dara.Validate(s)
}
