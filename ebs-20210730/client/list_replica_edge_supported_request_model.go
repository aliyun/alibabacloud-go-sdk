// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReplicaEdgeSupportedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAzone(v string) *ListReplicaEdgeSupportedRequest
	GetAzone() *string
	SetMaxResults(v int32) *ListReplicaEdgeSupportedRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListReplicaEdgeSupportedRequest
	GetNextToken() *string
	SetRegionId(v string) *ListReplicaEdgeSupportedRequest
	GetRegionId() *string
}

type ListReplicaEdgeSupportedRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-h
	Azone *string `json:"Azone,omitempty" xml:"Azone,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// e71d8a535bd9c****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListReplicaEdgeSupportedRequest) String() string {
	return dara.Prettify(s)
}

func (s ListReplicaEdgeSupportedRequest) GoString() string {
	return s.String()
}

func (s *ListReplicaEdgeSupportedRequest) GetAzone() *string {
	return s.Azone
}

func (s *ListReplicaEdgeSupportedRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListReplicaEdgeSupportedRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListReplicaEdgeSupportedRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListReplicaEdgeSupportedRequest) SetAzone(v string) *ListReplicaEdgeSupportedRequest {
	s.Azone = &v
	return s
}

func (s *ListReplicaEdgeSupportedRequest) SetMaxResults(v int32) *ListReplicaEdgeSupportedRequest {
	s.MaxResults = &v
	return s
}

func (s *ListReplicaEdgeSupportedRequest) SetNextToken(v string) *ListReplicaEdgeSupportedRequest {
	s.NextToken = &v
	return s
}

func (s *ListReplicaEdgeSupportedRequest) SetRegionId(v string) *ListReplicaEdgeSupportedRequest {
	s.RegionId = &v
	return s
}

func (s *ListReplicaEdgeSupportedRequest) Validate() error {
	return dara.Validate(s)
}
