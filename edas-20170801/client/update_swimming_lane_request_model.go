// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSwimmingLaneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInfos(v string) *UpdateSwimmingLaneRequest
	GetAppInfos() *string
	SetEnableRules(v bool) *UpdateSwimmingLaneRequest
	GetEnableRules() *bool
	SetEntryRules(v string) *UpdateSwimmingLaneRequest
	GetEntryRules() *string
	SetLaneId(v int64) *UpdateSwimmingLaneRequest
	GetLaneId() *int64
	SetName(v string) *UpdateSwimmingLaneRequest
	GetName() *string
}

type UpdateSwimmingLaneRequest struct {
	// The list of applications that are related to the lane.
	//
	// example:
	//
	// [{"appId":"8e7689af-6ddd-4676-8ee6-5fbecdf2****"},{"appId":"f72deaac-26ba-429a-948d-5fa47c4a****"},{"appId":"99a2d4b5-99a5-4e25-a964-1bd03a17****"}]
	AppInfos *string `json:"AppInfos,omitempty" xml:"AppInfos,omitempty"`
	// Specifies whether to enable the throttling rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableRules *bool `json:"EnableRules,omitempty" xml:"EnableRules,omitempty"`
	// The configuration of the throttling rule.
	//
	// example:
	//
	// [{"priority":1,"path":"/traffictest","condition":"AND","restItems":[{"type":"header","name":"testheader","value":"testheadervalue","cond":"==","operator":"rawvalue"}]}]
	EntryRules *string `json:"EntryRules,omitempty" xml:"EntryRules,omitempty"`
	// The ID of the lane.
	//
	// This parameter is required.
	//
	// example:
	//
	// 224
	LaneId *int64 `json:"LaneId,omitempty" xml:"LaneId,omitempty"`
	// The name of the lane.
	//
	// example:
	//
	// test-swimlane
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateSwimmingLaneRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimmingLaneRequest) GoString() string {
	return s.String()
}

func (s *UpdateSwimmingLaneRequest) GetAppInfos() *string {
	return s.AppInfos
}

func (s *UpdateSwimmingLaneRequest) GetEnableRules() *bool {
	return s.EnableRules
}

func (s *UpdateSwimmingLaneRequest) GetEntryRules() *string {
	return s.EntryRules
}

func (s *UpdateSwimmingLaneRequest) GetLaneId() *int64 {
	return s.LaneId
}

func (s *UpdateSwimmingLaneRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSwimmingLaneRequest) SetAppInfos(v string) *UpdateSwimmingLaneRequest {
	s.AppInfos = &v
	return s
}

func (s *UpdateSwimmingLaneRequest) SetEnableRules(v bool) *UpdateSwimmingLaneRequest {
	s.EnableRules = &v
	return s
}

func (s *UpdateSwimmingLaneRequest) SetEntryRules(v string) *UpdateSwimmingLaneRequest {
	s.EntryRules = &v
	return s
}

func (s *UpdateSwimmingLaneRequest) SetLaneId(v int64) *UpdateSwimmingLaneRequest {
	s.LaneId = &v
	return s
}

func (s *UpdateSwimmingLaneRequest) SetName(v string) *UpdateSwimmingLaneRequest {
	s.Name = &v
	return s
}

func (s *UpdateSwimmingLaneRequest) Validate() error {
	return dara.Validate(s)
}
