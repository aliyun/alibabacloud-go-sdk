// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEaisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientInstanceId(v string) *DescribeEaisRequest
	GetClientInstanceId() *string
	SetElasticAcceleratedInstanceIds(v string) *DescribeEaisRequest
	GetElasticAcceleratedInstanceIds() *string
	SetInstanceName(v string) *DescribeEaisRequest
	GetInstanceName() *string
	SetInstanceType(v string) *DescribeEaisRequest
	GetInstanceType() *string
	SetPageNumber(v int32) *DescribeEaisRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEaisRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeEaisRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeEaisRequest
	GetResourceGroupId() *string
	SetStatus(v string) *DescribeEaisRequest
	GetStatus() *string
	SetTag(v []*DescribeEaisRequestTag) *DescribeEaisRequest
	GetTag() []*DescribeEaisRequestTag
}

type DescribeEaisRequest struct {
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	// example:
	//
	// ["eais-id1", "eais-id2"]
	ElasticAcceleratedInstanceIds *string `json:"ElasticAcceleratedInstanceIds,omitempty" xml:"ElasticAcceleratedInstanceIds,omitempty"`
	// example:
	//
	// eais*
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// eais.ei-a6.2xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 200
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// InUse
	Status *string                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag    []*DescribeEaisRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeEaisRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEaisRequest) GoString() string {
	return s.String()
}

func (s *DescribeEaisRequest) GetClientInstanceId() *string {
	return s.ClientInstanceId
}

func (s *DescribeEaisRequest) GetElasticAcceleratedInstanceIds() *string {
	return s.ElasticAcceleratedInstanceIds
}

func (s *DescribeEaisRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeEaisRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeEaisRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEaisRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEaisRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEaisRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeEaisRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeEaisRequest) GetTag() []*DescribeEaisRequestTag {
	return s.Tag
}

func (s *DescribeEaisRequest) SetClientInstanceId(v string) *DescribeEaisRequest {
	s.ClientInstanceId = &v
	return s
}

func (s *DescribeEaisRequest) SetElasticAcceleratedInstanceIds(v string) *DescribeEaisRequest {
	s.ElasticAcceleratedInstanceIds = &v
	return s
}

func (s *DescribeEaisRequest) SetInstanceName(v string) *DescribeEaisRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeEaisRequest) SetInstanceType(v string) *DescribeEaisRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeEaisRequest) SetPageNumber(v int32) *DescribeEaisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEaisRequest) SetPageSize(v int32) *DescribeEaisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEaisRequest) SetRegionId(v string) *DescribeEaisRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEaisRequest) SetResourceGroupId(v string) *DescribeEaisRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeEaisRequest) SetStatus(v string) *DescribeEaisRequest {
	s.Status = &v
	return s
}

func (s *DescribeEaisRequest) SetTag(v []*DescribeEaisRequestTag) *DescribeEaisRequest {
	s.Tag = v
	return s
}

func (s *DescribeEaisRequest) Validate() error {
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

type DescribeEaisRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEaisRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeEaisRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeEaisRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeEaisRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeEaisRequestTag) SetKey(v string) *DescribeEaisRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeEaisRequestTag) SetValue(v string) *DescribeEaisRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeEaisRequestTag) Validate() error {
	return dara.Validate(s)
}
