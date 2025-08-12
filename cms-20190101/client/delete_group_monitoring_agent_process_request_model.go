// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupMonitoringAgentProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteGroupMonitoringAgentProcessRequest
	GetGroupId() *string
	SetId(v string) *DeleteGroupMonitoringAgentProcessRequest
	GetId() *string
	SetRegionId(v string) *DeleteGroupMonitoringAgentProcessRequest
	GetRegionId() *string
}

type DeleteGroupMonitoringAgentProcessRequest struct {
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the process monitoring task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 48F83746-C817-478C-9B06-7158F56B****
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteGroupMonitoringAgentProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupMonitoringAgentProcessRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupMonitoringAgentProcessRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteGroupMonitoringAgentProcessRequest) GetId() *string {
	return s.Id
}

func (s *DeleteGroupMonitoringAgentProcessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteGroupMonitoringAgentProcessRequest) SetGroupId(v string) *DeleteGroupMonitoringAgentProcessRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteGroupMonitoringAgentProcessRequest) SetId(v string) *DeleteGroupMonitoringAgentProcessRequest {
	s.Id = &v
	return s
}

func (s *DeleteGroupMonitoringAgentProcessRequest) SetRegionId(v string) *DeleteGroupMonitoringAgentProcessRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGroupMonitoringAgentProcessRequest) Validate() error {
	return dara.Validate(s)
}
