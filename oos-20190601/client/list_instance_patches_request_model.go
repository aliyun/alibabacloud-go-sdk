// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancePatchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListInstancePatchesRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListInstancePatchesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstancePatchesRequest
	GetNextToken() *string
	SetPatchStatuses(v string) *ListInstancePatchesRequest
	GetPatchStatuses() *string
	SetRegionId(v string) *ListInstancePatchesRequest
	GetRegionId() *string
}

type ListInstancePatchesRequest struct {
	// The ID of the instance.
	//
	// example:
	//
	// i-bp1jaxa2bs4bps7*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// gAAAAABfTgv5ewUWmNdJ3g7JVLvX70sPH90GZOVGC6KPDUL0FIIb
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The status of the patches that you want to query. If you do not set this parameter, patches are not filtered.
	//
	// example:
	//
	// Installed
	PatchStatuses *string `json:"PatchStatuses,omitempty" xml:"PatchStatuses,omitempty"`
	// The ID of the region in which the instance whose patches you want to query resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancePatchesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancePatchesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancePatchesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancePatchesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstancePatchesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstancePatchesRequest) GetPatchStatuses() *string {
	return s.PatchStatuses
}

func (s *ListInstancePatchesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancePatchesRequest) SetInstanceId(v string) *ListInstancePatchesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancePatchesRequest) SetMaxResults(v int32) *ListInstancePatchesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancePatchesRequest) SetNextToken(v string) *ListInstancePatchesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancePatchesRequest) SetPatchStatuses(v string) *ListInstancePatchesRequest {
	s.PatchStatuses = &v
	return s
}

func (s *ListInstancePatchesRequest) SetRegionId(v string) *ListInstancePatchesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancePatchesRequest) Validate() error {
	return dara.Validate(s)
}
