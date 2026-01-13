// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *ListInstancesRequest
	GetInstanceName() *string
	SetInstanceNameList(v []*string) *ListInstancesRequest
	GetInstanceNameList() []*string
	SetMaxResults(v int32) *ListInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstancesRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListInstancesRequest
	GetResourceGroupId() *string
	SetStatus(v string) *ListInstancesRequest
	GetStatus() *string
	SetTag(v []*ListInstancesRequestTag) *ListInstancesRequest
	GetTag() []*ListInstancesRequestTag
}

type ListInstancesRequest struct {
	// The name of the instance. Fuzzy search is supported.
	//
	// example:
	//
	// instance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The names of the instances. This parameter is used to specify multiple instances that you want to query at the same time.
	InstanceNameList []*string `json:"InstanceNameList,omitempty" xml:"InstanceNameList,omitempty" type:"Repeated"`
	// The maximum number of instances that you want to return. Valid values: 0 to 200. If you do not configure this parameter or set this parameter to 0, the default value of 100 is used.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start position of the query. Set this parameter to the value of the NextToken parameter that is returned from the last call. Instances are returned in lexicographical order starting from the position that is specified by this parameter. The first time you call the operation, leave this parameter empty.
	//
	// example:
	//
	// CAESCG15aC1xxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource group ID. You can query the ID on the Resource Group page in the Resource Management console.
	//
	// example:
	//
	// rg-aek24upgom6p5ri
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the instance.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the instance.
	Tag []*ListInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesRequest) GetInstanceNameList() []*string {
	return s.InstanceNameList
}

func (s *ListInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesRequest) GetTag() []*ListInstancesRequestTag {
	return s.Tag
}

func (s *ListInstancesRequest) SetInstanceName(v string) *ListInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceNameList(v []*string) *ListInstancesRequest {
	s.InstanceNameList = v
	return s
}

func (s *ListInstancesRequest) SetMaxResults(v int32) *ListInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesRequest) SetNextToken(v string) *ListInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancesRequest) SetResourceGroupId(v string) *ListInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesRequest) SetTag(v []*ListInstancesRequestTag) *ListInstancesRequest {
	s.Tag = v
	return s
}

func (s *ListInstancesRequest) Validate() error {
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

type ListInstancesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// Owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Tester
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListInstancesRequestTag) SetKey(v string) *ListInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListInstancesRequestTag) SetValue(v string) *ListInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *ListInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
