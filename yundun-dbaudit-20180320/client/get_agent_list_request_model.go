// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIp(v string) *GetAgentListRequest
	GetAgentIp() *string
	SetAgentOs(v string) *GetAgentListRequest
	GetAgentOs() *string
	SetAgentStatus(v int32) *GetAgentListRequest
	GetAgentStatus() *int32
	SetInstanceId(v string) *GetAgentListRequest
	GetInstanceId() *string
	SetLang(v string) *GetAgentListRequest
	GetLang() *string
	SetRegionId(v string) *GetAgentListRequest
	GetRegionId() *string
}

type GetAgentListRequest struct {
	// example:
	//
	// 192.168.XX.XX
	AgentIp *string `json:"AgentIp,omitempty" xml:"AgentIp,omitempty"`
	// example:
	//
	// Windows
	AgentOs *string `json:"AgentOs,omitempty" xml:"AgentOs,omitempty"`
	// example:
	//
	// 1
	AgentStatus *int32 `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dbaudit-cn-78v1gc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAgentListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentListRequest) GoString() string {
	return s.String()
}

func (s *GetAgentListRequest) GetAgentIp() *string {
	return s.AgentIp
}

func (s *GetAgentListRequest) GetAgentOs() *string {
	return s.AgentOs
}

func (s *GetAgentListRequest) GetAgentStatus() *int32 {
	return s.AgentStatus
}

func (s *GetAgentListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAgentListRequest) GetLang() *string {
	return s.Lang
}

func (s *GetAgentListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAgentListRequest) SetAgentIp(v string) *GetAgentListRequest {
	s.AgentIp = &v
	return s
}

func (s *GetAgentListRequest) SetAgentOs(v string) *GetAgentListRequest {
	s.AgentOs = &v
	return s
}

func (s *GetAgentListRequest) SetAgentStatus(v int32) *GetAgentListRequest {
	s.AgentStatus = &v
	return s
}

func (s *GetAgentListRequest) SetInstanceId(v string) *GetAgentListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentListRequest) SetLang(v string) *GetAgentListRequest {
	s.Lang = &v
	return s
}

func (s *GetAgentListRequest) SetRegionId(v string) *GetAgentListRequest {
	s.RegionId = &v
	return s
}

func (s *GetAgentListRequest) Validate() error {
	return dara.Validate(s)
}
