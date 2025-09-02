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
	// Specifies whether to enable the alerting feature. Valid values: true and false.
	//
	// example:
	//
	// true
	AlertEnabled *bool `json:"AlertEnabled,omitempty" xml:"AlertEnabled,omitempty"`
	// The alert margin threshold of the baseline. Unit: minutes.
	//
	// example:
	//
	// 30
	AlertMarginThreshold *int32 `json:"AlertMarginThreshold,omitempty" xml:"AlertMarginThreshold,omitempty"`
	// The alert settings of the baseline.
	AlertSettingsShrink *string `json:"AlertSettings,omitempty" xml:"AlertSettings,omitempty"`
	// The baseline ID. You can call the [ListBaselines](https://help.aliyun.com/document_detail/2261507.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000010800007
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The name of the baseline.
	//
	// example:
	//
	// BaselineName
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The type of the baseline. Valid values: DAILY and HOURLY.
	//
	// example:
	//
	// DAILY
	BaselineType *string `json:"BaselineType,omitempty" xml:"BaselineType,omitempty"`
	// Specifies whether to enable the baseline. Valid values: true and false.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ancestor nodes of nodes in the baseline. Separate the ancestor nodes with commas (,). If a large number of ancestor nodes exist, we recommend that you create a zero load node and configure the zero load node as the descendant node of nodes in the baseline to facilitate node management.
	//
	// example:
	//
	// 1,2,3
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// The settings of the committed completion time of the baseline.
	OvertimeSettingsShrink *string `json:"OvertimeSettings,omitempty" xml:"OvertimeSettings,omitempty"`
	// The ID of the Alibaba Cloud account used by the baseline owner.
	//
	// example:
	//
	// 3726346****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the baseline. Valid values: {1,3,5,7,8}.
	//
	// example:
	//
	// 7
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The workspace ID. You can call the [ListBaselines](https://help.aliyun.com/document_detail/2261507.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2043
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the node that you want to disassociate from the baseline. You can specify multiple node IDs. Separate multiple node IDs with commas (,).
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
