// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBaselineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertEnabled(v bool) *UpdateBaselineShrinkRequest
	GetAlertEnabled() *bool
	SetAlertMarginThreshold(v int32) *UpdateBaselineShrinkRequest
	GetAlertMarginThreshold() *int32
	SetAlertSettingsShrink(v string) *UpdateBaselineShrinkRequest
	GetAlertSettingsShrink() *string
	SetBaselineId(v int64) *UpdateBaselineShrinkRequest
	GetBaselineId() *int64
	SetBaselineName(v string) *UpdateBaselineShrinkRequest
	GetBaselineName() *string
	SetBaselineType(v string) *UpdateBaselineShrinkRequest
	GetBaselineType() *string
	SetEnabled(v bool) *UpdateBaselineShrinkRequest
	GetEnabled() *bool
	SetNodeIds(v string) *UpdateBaselineShrinkRequest
	GetNodeIds() *string
	SetOvertimeSettingsShrink(v string) *UpdateBaselineShrinkRequest
	GetOvertimeSettingsShrink() *string
	SetOwner(v string) *UpdateBaselineShrinkRequest
	GetOwner() *string
	SetPriority(v int32) *UpdateBaselineShrinkRequest
	GetPriority() *int32
	SetProjectId(v int64) *UpdateBaselineShrinkRequest
	GetProjectId() *int64
	SetRemoveNodeIds(v string) *UpdateBaselineShrinkRequest
	GetRemoveNodeIds() *string
}

type UpdateBaselineShrinkRequest struct {
	// Specifies whether alerting is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	AlertEnabled *bool `json:"AlertEnabled,omitempty" xml:"AlertEnabled,omitempty"`
	// The baseline alert margin. Unit: minutes.
	//
	// example:
	//
	// 30
	AlertMarginThreshold *int32 `json:"AlertMarginThreshold,omitempty" xml:"AlertMarginThreshold,omitempty"`
	// The baseline alert configurations.
	AlertSettingsShrink *string `json:"AlertSettings,omitempty" xml:"AlertSettings,omitempty"`
	// The ID of the baseline. You can call [ListBaselines](https://help.aliyun.com/document_detail/2261507.html) to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000010800007
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The baseline name.
	//
	// example:
	//
	// BaselineName
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The baseline type. Valid values:
	//
	// - DAILY: daily baseline.
	//
	// - HOURLY: hourly baseline.
	//
	// example:
	//
	// DAILY
	BaselineType *string `json:"BaselineType,omitempty" xml:"BaselineType,omitempty"`
	// Specifies whether the baseline is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The list of upstream node IDs for the baseline, separated by commas. If there are many nodes, we recommend that you add a virtual node downstream for easier management.
	//
	// example:
	//
	// 1,2,3
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// The baseline committed time configurations.
	OvertimeSettingsShrink *string `json:"OvertimeSettings,omitempty" xml:"OvertimeSettings,omitempty"`
	// The Alibaba Cloud UID of the baseline owner.
	//
	// example:
	//
	// 3726346****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the baseline. Valid values: 1, 3, 5, 7, and 8.
	//
	// example:
	//
	// 7
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The project ID. You can call [ListBaselines](https://help.aliyun.com/document_detail/2261507.html) to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2043
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The IDs of nodes to remove from the baseline. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// 123,456
	RemoveNodeIds *string `json:"RemoveNodeIds,omitempty" xml:"RemoveNodeIds,omitempty"`
}

func (s UpdateBaselineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBaselineShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateBaselineShrinkRequest) GetAlertEnabled() *bool {
	return s.AlertEnabled
}

func (s *UpdateBaselineShrinkRequest) GetAlertMarginThreshold() *int32 {
	return s.AlertMarginThreshold
}

func (s *UpdateBaselineShrinkRequest) GetAlertSettingsShrink() *string {
	return s.AlertSettingsShrink
}

func (s *UpdateBaselineShrinkRequest) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *UpdateBaselineShrinkRequest) GetBaselineName() *string {
	return s.BaselineName
}

func (s *UpdateBaselineShrinkRequest) GetBaselineType() *string {
	return s.BaselineType
}

func (s *UpdateBaselineShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateBaselineShrinkRequest) GetNodeIds() *string {
	return s.NodeIds
}

func (s *UpdateBaselineShrinkRequest) GetOvertimeSettingsShrink() *string {
	return s.OvertimeSettingsShrink
}

func (s *UpdateBaselineShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateBaselineShrinkRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateBaselineShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateBaselineShrinkRequest) GetRemoveNodeIds() *string {
	return s.RemoveNodeIds
}

func (s *UpdateBaselineShrinkRequest) SetAlertEnabled(v bool) *UpdateBaselineShrinkRequest {
	s.AlertEnabled = &v
	return s
}

func (s *UpdateBaselineShrinkRequest) SetAlertMarginThreshold(v int32) *UpdateBaselineShrinkRequest {
	s.AlertMarginThreshold = &v
	return s
}

func (s *UpdateBaselineShrinkRequest) SetAlertSettingsShrink(v string) *UpdateBaselineShrinkRequest {
	s.AlertSettingsShrink = &v
	return s
}

func (s *UpdateBaselineShrinkRequest) SetBaselineId(v int64) *UpdateBaselineShrinkRequest {
	s.BaselineId = &v
	return s
}

func (s *UpdateBaselineShrinkRequest) SetBaselineName(v string) *UpdateBaselineShrinkRequest {
	s.BaselineName = &v
	return s
}

func (s *UpdateBaselineShrinkRequest) SetBaselineType(v string) *UpdateBaselineShrinkRequest {
	s.BaselineType = &v
	return s
}

func (s *UpdateBaselineShrinkRequest) SetEnabled(v bool) *UpdateBaselineShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateBaselineShrinkRequest) SetNodeIds(v string) *UpdateBaselineShrinkRequest {
	s.NodeIds = &v
	return s
}

func (s *UpdateBaselineShrinkRequest) SetOvertimeSettingsShrink(v string) *UpdateBaselineShrinkRequest {
	s.OvertimeSettingsShrink = &v
	return s
}

func (s *UpdateBaselineShrinkRequest) SetOwner(v string) *UpdateBaselineShrinkRequest {
	s.Owner = &v
	return s
}

func (s *UpdateBaselineShrinkRequest) SetPriority(v int32) *UpdateBaselineShrinkRequest {
	s.Priority = &v
	return s
}

func (s *UpdateBaselineShrinkRequest) SetProjectId(v int64) *UpdateBaselineShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateBaselineShrinkRequest) SetRemoveNodeIds(v string) *UpdateBaselineShrinkRequest {
	s.RemoveNodeIds = &v
	return s
}

func (s *UpdateBaselineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
