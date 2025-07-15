// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBuildConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListBuildConfigsRequest
	GetAgentKey() *string
	SetRegionId(v string) *ListBuildConfigsRequest
	GetRegionId() *string
	SetType(v string) *ListBuildConfigsRequest
	GetType() *string
}

type ListBuildConfigsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cd327c3d5d5e44159cc716e23bfa530e_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// media
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListBuildConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBuildConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListBuildConfigsRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListBuildConfigsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBuildConfigsRequest) GetType() *string {
	return s.Type
}

func (s *ListBuildConfigsRequest) SetAgentKey(v string) *ListBuildConfigsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListBuildConfigsRequest) SetRegionId(v string) *ListBuildConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *ListBuildConfigsRequest) SetType(v string) *ListBuildConfigsRequest {
	s.Type = &v
	return s
}

func (s *ListBuildConfigsRequest) Validate() error {
	return dara.Validate(s)
}
