// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSwimmingLaneShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppEntryRuleShrink(v string) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetAppEntryRuleShrink() *string
	SetCanaryModel(v int32) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetCanaryModel() *int32
	SetEnable(v bool) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetEnable() *bool
	SetGroupId(v int64) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetGroupId() *int64
	SetLaneId(v int64) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetLaneId() *int64
	SetLaneName(v string) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetLaneName() *string
	SetLaneTag(v string) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetLaneTag() *string
	SetMseGatewayEntryRuleShrink(v string) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetMseGatewayEntryRuleShrink() *string
	SetNamespaceId(v string) *CreateOrUpdateSwimmingLaneShrinkRequest
	GetNamespaceId() *string
}

type CreateOrUpdateSwimmingLaneShrinkRequest struct {
	AppEntryRuleShrink *string `json:"AppEntryRule,omitempty" xml:"AppEntryRule,omitempty"`
	// example:
	//
	// 0
	CanaryModel *int32 `json:"CanaryModel,omitempty" xml:"CanaryModel,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// b2a8a925-477a-eswa-b823-d5e22500****
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 13857
	LaneId *int64 `json:"LaneId,omitempty" xml:"LaneId,omitempty"`
	// example:
	//
	// test
	LaneName *string `json:"LaneName,omitempty" xml:"LaneName,omitempty"`
	// example:
	//
	// {"alicloud.service.tag":"g1"}
	LaneTag                   *string `json:"LaneTag,omitempty" xml:"LaneTag,omitempty"`
	MseGatewayEntryRuleShrink *string `json:"MseGatewayEntryRule,omitempty" xml:"MseGatewayEntryRule,omitempty"`
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetAppEntryRuleShrink() *string {
	return s.AppEntryRuleShrink
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetCanaryModel() *int32 {
	return s.CanaryModel
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetLaneId() *int64 {
	return s.LaneId
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetLaneName() *string {
	return s.LaneName
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetLaneTag() *string {
	return s.LaneTag
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetMseGatewayEntryRuleShrink() *string {
	return s.MseGatewayEntryRuleShrink
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetAppEntryRuleShrink(v string) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.AppEntryRuleShrink = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetCanaryModel(v int32) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.CanaryModel = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetEnable(v bool) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.Enable = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetGroupId(v int64) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetLaneId(v int64) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.LaneId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetLaneName(v string) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.LaneName = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetLaneTag(v string) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.LaneTag = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetMseGatewayEntryRuleShrink(v string) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.MseGatewayEntryRuleShrink = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) SetNamespaceId(v string) *CreateOrUpdateSwimmingLaneShrinkRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneShrinkRequest) Validate() error {
	return dara.Validate(s)
}
