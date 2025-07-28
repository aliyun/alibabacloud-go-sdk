// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *ListInstancesShrinkRequest
	GetInstanceName() *string
	SetInstanceNameListShrink(v string) *ListInstancesShrinkRequest
	GetInstanceNameListShrink() *string
	SetMaxResults(v int32) *ListInstancesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstancesShrinkRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListInstancesShrinkRequest
	GetResourceGroupId() *string
	SetStatus(v string) *ListInstancesShrinkRequest
	GetStatus() *string
	SetTagShrink(v string) *ListInstancesShrinkRequest
	GetTagShrink() *string
}

type ListInstancesShrinkRequest struct {
	// The name of the instance. Fuzzy search is supported.
	//
	// example:
	//
	// instance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The names of the instances. This parameter is used to specify multiple instances that you want to query at the same time.
	InstanceNameListShrink *string `json:"InstanceNameList,omitempty" xml:"InstanceNameList,omitempty"`
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
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesShrinkRequest) GetInstanceNameListShrink() *string {
	return s.InstanceNameListShrink
}

func (s *ListInstancesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstancesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstancesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListInstancesShrinkRequest) SetInstanceName(v string) *ListInstancesShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetInstanceNameListShrink(v string) *ListInstancesShrinkRequest {
	s.InstanceNameListShrink = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetMaxResults(v int32) *ListInstancesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetNextToken(v string) *ListInstancesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetResourceGroupId(v string) *ListInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetStatus(v string) *ListInstancesShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetTagShrink(v string) *ListInstancesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
