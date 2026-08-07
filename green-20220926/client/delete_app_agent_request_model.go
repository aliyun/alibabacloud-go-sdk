// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *DeleteAppAgentRequest
	GetAgentId() *string
	SetAppId(v string) *DeleteAppAgentRequest
	GetAppId() *string
	SetRegionId(v string) *DeleteAppAgentRequest
	GetRegionId() *string
	SetResourceType(v string) *DeleteAppAgentRequest
	GetResourceType() *string
}

type DeleteAppAgentRequest struct {
	// Agent ID。
	//
	// example:
	//
	// ag.abcxxx
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// App ID。
	//
	// example:
	//
	// txt_check_agent_01
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// agent_text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DeleteAppAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppAgentRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppAgentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *DeleteAppAgentRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAppAgentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAppAgentRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteAppAgentRequest) SetAgentId(v string) *DeleteAppAgentRequest {
	s.AgentId = &v
	return s
}

func (s *DeleteAppAgentRequest) SetAppId(v string) *DeleteAppAgentRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppAgentRequest) SetRegionId(v string) *DeleteAppAgentRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAppAgentRequest) SetResourceType(v string) *DeleteAppAgentRequest {
	s.ResourceType = &v
	return s
}

func (s *DeleteAppAgentRequest) Validate() error {
	return dara.Validate(s)
}
