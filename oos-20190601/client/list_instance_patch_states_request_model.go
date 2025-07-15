// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancePatchStatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *ListInstancePatchStatesRequest
	GetInstanceIds() *string
	SetMaxResults(v int32) *ListInstancePatchStatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstancePatchStatesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListInstancePatchStatesRequest
	GetRegionId() *string
}

type ListInstancePatchStatesRequest struct {
	// The ID of the Elastic Compute Service (ECS) instance. The value can be a JSON array that consists of up to 100 instance IDs. Separate the instance IDs with commas (,).
	//
	// example:
	//
	// ["i-bp1jaxa2bs4bps7*****", "i-bp67acfmxazb4p****", … "i-bp67acfmxazb4p****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region in which the instance whose patches you want to query resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancePatchStatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancePatchStatesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancePatchStatesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ListInstancePatchStatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstancePatchStatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstancePatchStatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancePatchStatesRequest) SetInstanceIds(v string) *ListInstancePatchStatesRequest {
	s.InstanceIds = &v
	return s
}

func (s *ListInstancePatchStatesRequest) SetMaxResults(v int32) *ListInstancePatchStatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancePatchStatesRequest) SetNextToken(v string) *ListInstancePatchStatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancePatchStatesRequest) SetRegionId(v string) *ListInstancePatchStatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancePatchStatesRequest) Validate() error {
	return dara.Validate(s)
}
