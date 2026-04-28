// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAgentRuntimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DisableAgentRuntimeRequest
	GetClientToken() *string
	SetInstanceName(v string) *DisableAgentRuntimeRequest
	GetInstanceName() *string
	SetRegionId(v string) *DisableAgentRuntimeRequest
	GetRegionId() *string
}

type DisableAgentRuntimeRequest struct {
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88**********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableAgentRuntimeRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableAgentRuntimeRequest) GoString() string {
	return s.String()
}

func (s *DisableAgentRuntimeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableAgentRuntimeRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DisableAgentRuntimeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableAgentRuntimeRequest) SetClientToken(v string) *DisableAgentRuntimeRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableAgentRuntimeRequest) SetInstanceName(v string) *DisableAgentRuntimeRequest {
	s.InstanceName = &v
	return s
}

func (s *DisableAgentRuntimeRequest) SetRegionId(v string) *DisableAgentRuntimeRequest {
	s.RegionId = &v
	return s
}

func (s *DisableAgentRuntimeRequest) Validate() error {
	return dara.Validate(s)
}
