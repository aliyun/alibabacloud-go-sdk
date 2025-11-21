// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListInstancesV2Request
	GetRegionId() *string
	SetInstanceId(v string) *ListInstancesV2Request
	GetInstanceId() *string
	SetInstanceName(v string) *ListInstancesV2Request
	GetInstanceName() *string
	SetMaxResults(v int32) *ListInstancesV2Request
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstancesV2Request
	GetNextToken() *string
	SetPageNumber(v int32) *ListInstancesV2Request
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesV2Request
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListInstancesV2Request
	GetResourceGroupId() *string
	SetTag(v []*ListInstancesV2RequestTag) *ListInstancesV2Request
	GetTag() []*ListInstancesV2RequestTag
}

type ListInstancesV2Request struct {
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// milvus-test
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// None
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// rg-123xxx
	ResourceGroupId *string                      `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Tag             []*ListInstancesV2RequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s ListInstancesV2Request) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesV2Request) GoString() string {
	return s.String()
}

func (s *ListInstancesV2Request) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesV2Request) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesV2Request) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesV2Request) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstancesV2Request) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstancesV2Request) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesV2Request) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesV2Request) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesV2Request) GetTag() []*ListInstancesV2RequestTag {
	return s.Tag
}

func (s *ListInstancesV2Request) SetRegionId(v string) *ListInstancesV2Request {
	s.RegionId = &v
	return s
}

func (s *ListInstancesV2Request) SetInstanceId(v string) *ListInstancesV2Request {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesV2Request) SetInstanceName(v string) *ListInstancesV2Request {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesV2Request) SetMaxResults(v int32) *ListInstancesV2Request {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesV2Request) SetNextToken(v string) *ListInstancesV2Request {
	s.NextToken = &v
	return s
}

func (s *ListInstancesV2Request) SetPageNumber(v int32) *ListInstancesV2Request {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesV2Request) SetPageSize(v int32) *ListInstancesV2Request {
	s.PageSize = &v
	return s
}

func (s *ListInstancesV2Request) SetResourceGroupId(v string) *ListInstancesV2Request {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesV2Request) SetTag(v []*ListInstancesV2RequestTag) *ListInstancesV2Request {
	s.Tag = v
	return s
}

func (s *ListInstancesV2Request) Validate() error {
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

type ListInstancesV2RequestTag struct {
	// example:
	//
	// k1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// v1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListInstancesV2RequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesV2RequestTag) GoString() string {
	return s.String()
}

func (s *ListInstancesV2RequestTag) GetKey() *string {
	return s.Key
}

func (s *ListInstancesV2RequestTag) GetValue() *string {
	return s.Value
}

func (s *ListInstancesV2RequestTag) SetKey(v string) *ListInstancesV2RequestTag {
	s.Key = &v
	return s
}

func (s *ListInstancesV2RequestTag) SetValue(v string) *ListInstancesV2RequestTag {
	s.Value = &v
	return s
}

func (s *ListInstancesV2RequestTag) Validate() error {
	return dara.Validate(s)
}
