// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsMappingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilesystemId(v string) *ListUserGroupsMappingsRequest
	GetFilesystemId() *string
	SetInputRegionId(v string) *ListUserGroupsMappingsRequest
	GetInputRegionId() *string
	SetLimit(v int32) *ListUserGroupsMappingsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListUserGroupsMappingsRequest
	GetNextToken() *string
}

type ListUserGroupsMappingsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	FilesystemId *string `json:"FilesystemId,omitempty" xml:"FilesystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// user1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListUserGroupsMappingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsMappingsRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsMappingsRequest) GetFilesystemId() *string {
	return s.FilesystemId
}

func (s *ListUserGroupsMappingsRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *ListUserGroupsMappingsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListUserGroupsMappingsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserGroupsMappingsRequest) SetFilesystemId(v string) *ListUserGroupsMappingsRequest {
	s.FilesystemId = &v
	return s
}

func (s *ListUserGroupsMappingsRequest) SetInputRegionId(v string) *ListUserGroupsMappingsRequest {
	s.InputRegionId = &v
	return s
}

func (s *ListUserGroupsMappingsRequest) SetLimit(v int32) *ListUserGroupsMappingsRequest {
	s.Limit = &v
	return s
}

func (s *ListUserGroupsMappingsRequest) SetNextToken(v string) *ListUserGroupsMappingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserGroupsMappingsRequest) Validate() error {
	return dara.Validate(s)
}
