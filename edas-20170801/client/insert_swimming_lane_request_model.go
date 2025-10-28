// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertSwimmingLaneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInfos(v string) *InsertSwimmingLaneRequest
	GetAppInfos() *string
	SetEnableRules(v bool) *InsertSwimmingLaneRequest
	GetEnableRules() *bool
	SetEntryRules(v string) *InsertSwimmingLaneRequest
	GetEntryRules() *string
	SetGroupId(v int64) *InsertSwimmingLaneRequest
	GetGroupId() *int64
	SetLogicalRegionId(v string) *InsertSwimmingLaneRequest
	GetLogicalRegionId() *string
	SetName(v string) *InsertSwimmingLaneRequest
	GetName() *string
	SetTag(v string) *InsertSwimmingLaneRequest
	GetTag() *string
}

type InsertSwimmingLaneRequest struct {
	// The information about applications related to the lane.
	//
	// example:
	//
	// [{"appId":"f72deaac-26ba-429a-948d-5fa47c4a****"},{"appId":"5049d2c8-f997-4fc9-92a2-153506a6****"}]
	AppInfos *string `json:"AppInfos,omitempty" xml:"AppInfos,omitempty"`
	// Specifies whether to enable the throttling rule.
	//
	// example:
	//
	// true
	EnableRules *bool `json:"EnableRules,omitempty" xml:"EnableRules,omitempty"`
	// The throttling conditions.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"priority":1,"path":"/traffic","condition":"AND","restItems":[{"type":"header","name":"testheader","value":"testvalue","cond":"==","operator":"rawvalue"}]}]
	EntryRules *string `json:"EntryRules,omitempty" xml:"EntryRules,omitempty"`
	// The ID of the lane group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the custom namespace. The ID is in the `physical region ID:custom namespace identifier` format. Example: `cn-hangzhou:test`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	// The name of the lane.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s InsertSwimmingLaneRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertSwimmingLaneRequest) GoString() string {
	return s.String()
}

func (s *InsertSwimmingLaneRequest) GetAppInfos() *string {
	return s.AppInfos
}

func (s *InsertSwimmingLaneRequest) GetEnableRules() *bool {
	return s.EnableRules
}

func (s *InsertSwimmingLaneRequest) GetEntryRules() *string {
	return s.EntryRules
}

func (s *InsertSwimmingLaneRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *InsertSwimmingLaneRequest) GetLogicalRegionId() *string {
	return s.LogicalRegionId
}

func (s *InsertSwimmingLaneRequest) GetName() *string {
	return s.Name
}

func (s *InsertSwimmingLaneRequest) GetTag() *string {
	return s.Tag
}

func (s *InsertSwimmingLaneRequest) SetAppInfos(v string) *InsertSwimmingLaneRequest {
	s.AppInfos = &v
	return s
}

func (s *InsertSwimmingLaneRequest) SetEnableRules(v bool) *InsertSwimmingLaneRequest {
	s.EnableRules = &v
	return s
}

func (s *InsertSwimmingLaneRequest) SetEntryRules(v string) *InsertSwimmingLaneRequest {
	s.EntryRules = &v
	return s
}

func (s *InsertSwimmingLaneRequest) SetGroupId(v int64) *InsertSwimmingLaneRequest {
	s.GroupId = &v
	return s
}

func (s *InsertSwimmingLaneRequest) SetLogicalRegionId(v string) *InsertSwimmingLaneRequest {
	s.LogicalRegionId = &v
	return s
}

func (s *InsertSwimmingLaneRequest) SetName(v string) *InsertSwimmingLaneRequest {
	s.Name = &v
	return s
}

func (s *InsertSwimmingLaneRequest) SetTag(v string) *InsertSwimmingLaneRequest {
	s.Tag = &v
	return s
}

func (s *InsertSwimmingLaneRequest) Validate() error {
	return dara.Validate(s)
}
