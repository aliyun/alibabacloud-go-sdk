// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorGroupDynamicRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DeleteMonitorGroupDynamicRuleRequest
	GetCategory() *string
	SetGroupId(v int64) *DeleteMonitorGroupDynamicRuleRequest
	GetGroupId() *int64
	SetRegionId(v string) *DeleteMonitorGroupDynamicRuleRequest
	GetRegionId() *string
}

type DeleteMonitorGroupDynamicRuleRequest struct {
	// The service to which the rule applies. Valid values: ecs, rds, and slb.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId  *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteMonitorGroupDynamicRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupDynamicRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupDynamicRuleRequest) GetCategory() *string {
	return s.Category
}

func (s *DeleteMonitorGroupDynamicRuleRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DeleteMonitorGroupDynamicRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMonitorGroupDynamicRuleRequest) SetCategory(v string) *DeleteMonitorGroupDynamicRuleRequest {
	s.Category = &v
	return s
}

func (s *DeleteMonitorGroupDynamicRuleRequest) SetGroupId(v int64) *DeleteMonitorGroupDynamicRuleRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteMonitorGroupDynamicRuleRequest) SetRegionId(v string) *DeleteMonitorGroupDynamicRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMonitorGroupDynamicRuleRequest) Validate() error {
	return dara.Validate(s)
}
