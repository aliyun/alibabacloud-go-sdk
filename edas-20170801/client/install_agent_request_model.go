// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *InstallAgentRequest
	GetClusterId() *string
	SetDoAsync(v bool) *InstallAgentRequest
	GetDoAsync() *bool
	SetInstanceIds(v string) *InstallAgentRequest
	GetInstanceIds() *string
}

type InstallAgentRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// b3e3f77b-462e-****-****-bec8727a4dc8
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is discontinued.
	//
	// example:
	//
	// true
	DoAsync *bool `json:"DoAsync,omitempty" xml:"DoAsync,omitempty"`
	// The ID of the ECS instance. Separate multiple IDs with commas (,). Example: instanceId1,instanceId2.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ze7s2v0b789k60pk1af
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s InstallAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentRequest) GoString() string {
	return s.String()
}

func (s *InstallAgentRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *InstallAgentRequest) GetDoAsync() *bool {
	return s.DoAsync
}

func (s *InstallAgentRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *InstallAgentRequest) SetClusterId(v string) *InstallAgentRequest {
	s.ClusterId = &v
	return s
}

func (s *InstallAgentRequest) SetDoAsync(v bool) *InstallAgentRequest {
	s.DoAsync = &v
	return s
}

func (s *InstallAgentRequest) SetInstanceIds(v string) *InstallAgentRequest {
	s.InstanceIds = &v
	return s
}

func (s *InstallAgentRequest) Validate() error {
	return dara.Validate(s)
}
