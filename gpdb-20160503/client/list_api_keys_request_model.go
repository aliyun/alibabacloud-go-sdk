// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyName(v string) *ListApiKeysRequest
	GetKeyName() *string
	SetMaxResults(v int32) *ListApiKeysRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListApiKeysRequest
	GetNextToken() *string
	SetRegionId(v string) *ListApiKeysRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *ListApiKeysRequest
	GetWorkspaceId() *string
}

type ListApiKeysRequest struct {
	// example:
	//
	// my api
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListApiKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysRequest) GoString() string {
	return s.String()
}

func (s *ListApiKeysRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *ListApiKeysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApiKeysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApiKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApiKeysRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListApiKeysRequest) SetKeyName(v string) *ListApiKeysRequest {
	s.KeyName = &v
	return s
}

func (s *ListApiKeysRequest) SetMaxResults(v int32) *ListApiKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApiKeysRequest) SetNextToken(v string) *ListApiKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListApiKeysRequest) SetRegionId(v string) *ListApiKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListApiKeysRequest) SetWorkspaceId(v string) *ListApiKeysRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListApiKeysRequest) Validate() error {
	return dara.Validate(s)
}
