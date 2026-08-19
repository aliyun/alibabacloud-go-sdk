// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserDeliveryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *UpdateUserDeliveryTaskRequest
	GetBusinessType() *string
	SetDetails(v string) *UpdateUserDeliveryTaskRequest
	GetDetails() *string
	SetDiscardRate(v float32) *UpdateUserDeliveryTaskRequest
	GetDiscardRate() *float32
	SetFieldName(v string) *UpdateUserDeliveryTaskRequest
	GetFieldName() *string
	SetFilterVer(v string) *UpdateUserDeliveryTaskRequest
	GetFilterVer() *string
	SetTaskName(v string) *UpdateUserDeliveryTaskRequest
	GetTaskName() *string
}

type UpdateUserDeliveryTaskRequest struct {
	// The real-time log type. Valid values:
	//
	// - **dcdn_log_access_l1 (default)**: access logs.
	//
	// - **dcdn_log_er**: Edge Routine function logs.
	//
	// - **dcdn_log_waf**: security protection logs.
	//
	// - **dcdn_log_ipa**: Layer 4 acceleration logs.
	//
	// example:
	//
	// dcdn_log_er
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The list of ER PODs to configure.
	//
	// example:
	//
	// xxx,xxx
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// The discard rate. If not specified, the default value is 0.
	//
	// example:
	//
	// 0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// The selected fields, separated by commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ClientRequestID,ClientRequestHost
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The version of the filter rule.
	//
	// > Compatible with legacy filter rules. The default value is v1. Newly created tasks use v2.
	//
	// example:
	//
	// v2
	FilterVer *string `json:"FilterVer,omitempty" xml:"FilterVer,omitempty"`
	// The task name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_project
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateUserDeliveryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserDeliveryTaskRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *UpdateUserDeliveryTaskRequest) GetDetails() *string {
	return s.Details
}

func (s *UpdateUserDeliveryTaskRequest) GetDiscardRate() *float32 {
	return s.DiscardRate
}

func (s *UpdateUserDeliveryTaskRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *UpdateUserDeliveryTaskRequest) GetFilterVer() *string {
	return s.FilterVer
}

func (s *UpdateUserDeliveryTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateUserDeliveryTaskRequest) SetBusinessType(v string) *UpdateUserDeliveryTaskRequest {
	s.BusinessType = &v
	return s
}

func (s *UpdateUserDeliveryTaskRequest) SetDetails(v string) *UpdateUserDeliveryTaskRequest {
	s.Details = &v
	return s
}

func (s *UpdateUserDeliveryTaskRequest) SetDiscardRate(v float32) *UpdateUserDeliveryTaskRequest {
	s.DiscardRate = &v
	return s
}

func (s *UpdateUserDeliveryTaskRequest) SetFieldName(v string) *UpdateUserDeliveryTaskRequest {
	s.FieldName = &v
	return s
}

func (s *UpdateUserDeliveryTaskRequest) SetFilterVer(v string) *UpdateUserDeliveryTaskRequest {
	s.FilterVer = &v
	return s
}

func (s *UpdateUserDeliveryTaskRequest) SetTaskName(v string) *UpdateUserDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

func (s *UpdateUserDeliveryTaskRequest) Validate() error {
	return dara.Validate(s)
}
