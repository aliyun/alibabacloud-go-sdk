// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRdMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListInstanceRdMembersRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListInstanceRdMembersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstanceRdMembersRequest
	GetNextToken() *string
	SetRegionId(v string) *ListInstanceRdMembersRequest
	GetRegionId() *string
}

type ListInstanceRdMembersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 4ieSWJCwxvW3dk3wF.BqkrZmP72nWu5zJ5NWydMqyEs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstanceRdMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRdMembersRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRdMembersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceRdMembersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstanceRdMembersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstanceRdMembersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceRdMembersRequest) SetInstanceId(v string) *ListInstanceRdMembersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceRdMembersRequest) SetMaxResults(v int32) *ListInstanceRdMembersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstanceRdMembersRequest) SetNextToken(v string) *ListInstanceRdMembersRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstanceRdMembersRequest) SetRegionId(v string) *ListInstanceRdMembersRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstanceRdMembersRequest) Validate() error {
	return dara.Validate(s)
}
