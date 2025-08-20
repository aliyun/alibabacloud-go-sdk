// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackResourceDriftsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListStackResourceDriftsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListStackResourceDriftsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListStackResourceDriftsRequest
	GetRegionId() *string
	SetResourceDriftStatus(v []*string) *ListStackResourceDriftsRequest
	GetResourceDriftStatus() []*string
	SetStackId(v string) *ListStackResourceDriftsRequest
	GetStackId() *string
}

type ListStackResourceDriftsRequest struct {
	// The time when the resource drift detection operation was initiated.
	//
	// example:
	//
	// 50
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// AAAAAdDWBF2****w==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The physical ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource properties as defined in the template, in JSON format.
	//
	// example:
	//
	// MODIFIED
	ResourceDriftStatus []*string `json:"ResourceDriftStatus,omitempty" xml:"ResourceDriftStatus,omitempty" type:"Repeated"`
	// The ID of the stack.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s ListStackResourceDriftsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStackResourceDriftsRequest) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListStackResourceDriftsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListStackResourceDriftsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStackResourceDriftsRequest) GetResourceDriftStatus() []*string {
	return s.ResourceDriftStatus
}

func (s *ListStackResourceDriftsRequest) GetStackId() *string {
	return s.StackId
}

func (s *ListStackResourceDriftsRequest) SetMaxResults(v int64) *ListStackResourceDriftsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListStackResourceDriftsRequest) SetNextToken(v string) *ListStackResourceDriftsRequest {
	s.NextToken = &v
	return s
}

func (s *ListStackResourceDriftsRequest) SetRegionId(v string) *ListStackResourceDriftsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackResourceDriftsRequest) SetResourceDriftStatus(v []*string) *ListStackResourceDriftsRequest {
	s.ResourceDriftStatus = v
	return s
}

func (s *ListStackResourceDriftsRequest) SetStackId(v string) *ListStackResourceDriftsRequest {
	s.StackId = &v
	return s
}

func (s *ListStackResourceDriftsRequest) Validate() error {
	return dara.Validate(s)
}
