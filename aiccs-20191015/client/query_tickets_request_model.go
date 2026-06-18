// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTicketsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaseId(v int64) *QueryTicketsRequest
	GetCaseId() *int64
	SetCaseStatus(v int32) *QueryTicketsRequest
	GetCaseStatus() *int32
	SetCaseType(v int32) *QueryTicketsRequest
	GetCaseType() *int32
	SetChannelId(v string) *QueryTicketsRequest
	GetChannelId() *string
	SetChannelType(v int32) *QueryTicketsRequest
	GetChannelType() *int32
	SetCurrentPage(v int32) *QueryTicketsRequest
	GetCurrentPage() *int32
	SetDealId(v int64) *QueryTicketsRequest
	GetDealId() *int64
	SetExtra(v map[string]interface{}) *QueryTicketsRequest
	GetExtra() map[string]interface{}
	SetInstanceId(v string) *QueryTicketsRequest
	GetInstanceId() *string
	SetPageSize(v int32) *QueryTicketsRequest
	GetPageSize() *int32
	SetSrType(v int64) *QueryTicketsRequest
	GetSrType() *int64
	SetTaskStatus(v int32) *QueryTicketsRequest
	GetTaskStatus() *int32
	SetTouchId(v int64) *QueryTicketsRequest
	GetTouchId() *int64
}

type QueryTicketsRequest struct {
	// Ticket ID.
	//
	// example:
	//
	// 223468****
	CaseId *int64 `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	// Ticket status code.
	//
	// example:
	//
	// 2
	CaseStatus *int32 `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	// Ticket type code (assigned by the system).
	//
	// example:
	//
	// 1223
	CaseType *int32 `json:"CaseType,omitempty" xml:"CaseType,omitempty"`
	// Channel ID.
	//
	// example:
	//
	// 02acfefd3fa14049826ac1a89e1xxxxx
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// Channel Type. Valid values:
	//
	// - **0**: Not filled in
	//
	// - **1**: Hotline
	//
	// - **2**: Online
	//
	// example:
	//
	// 1
	ChannelType *int32 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// Current page. Default Value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Assignee ID.
	//
	// example:
	//
	// 123456
	DealId *int64 `json:"DealId,omitempty" xml:"DealId,omitempty"`
	// Additional information.
	//
	// example:
	//
	// 无
	Extra map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// Instance ID.
	//
	// Log on to the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview) and view the instance ID in **Instance Management**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Page size. Default Value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Template ID.
	//
	// example:
	//
	// 29506
	SrType *int64 `json:"SrType,omitempty" xml:"SrType,omitempty"`
	// Job status.
	//
	// example:
	//
	// 3
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// Touch ID.
	//
	// example:
	//
	// 15030609
	TouchId *int64 `json:"TouchId,omitempty" xml:"TouchId,omitempty"`
}

func (s QueryTicketsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTicketsRequest) GoString() string {
	return s.String()
}

func (s *QueryTicketsRequest) GetCaseId() *int64 {
	return s.CaseId
}

func (s *QueryTicketsRequest) GetCaseStatus() *int32 {
	return s.CaseStatus
}

func (s *QueryTicketsRequest) GetCaseType() *int32 {
	return s.CaseType
}

func (s *QueryTicketsRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *QueryTicketsRequest) GetChannelType() *int32 {
	return s.ChannelType
}

func (s *QueryTicketsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryTicketsRequest) GetDealId() *int64 {
	return s.DealId
}

func (s *QueryTicketsRequest) GetExtra() map[string]interface{} {
	return s.Extra
}

func (s *QueryTicketsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryTicketsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTicketsRequest) GetSrType() *int64 {
	return s.SrType
}

func (s *QueryTicketsRequest) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *QueryTicketsRequest) GetTouchId() *int64 {
	return s.TouchId
}

func (s *QueryTicketsRequest) SetCaseId(v int64) *QueryTicketsRequest {
	s.CaseId = &v
	return s
}

func (s *QueryTicketsRequest) SetCaseStatus(v int32) *QueryTicketsRequest {
	s.CaseStatus = &v
	return s
}

func (s *QueryTicketsRequest) SetCaseType(v int32) *QueryTicketsRequest {
	s.CaseType = &v
	return s
}

func (s *QueryTicketsRequest) SetChannelId(v string) *QueryTicketsRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryTicketsRequest) SetChannelType(v int32) *QueryTicketsRequest {
	s.ChannelType = &v
	return s
}

func (s *QueryTicketsRequest) SetCurrentPage(v int32) *QueryTicketsRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryTicketsRequest) SetDealId(v int64) *QueryTicketsRequest {
	s.DealId = &v
	return s
}

func (s *QueryTicketsRequest) SetExtra(v map[string]interface{}) *QueryTicketsRequest {
	s.Extra = v
	return s
}

func (s *QueryTicketsRequest) SetInstanceId(v string) *QueryTicketsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTicketsRequest) SetPageSize(v int32) *QueryTicketsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTicketsRequest) SetSrType(v int64) *QueryTicketsRequest {
	s.SrType = &v
	return s
}

func (s *QueryTicketsRequest) SetTaskStatus(v int32) *QueryTicketsRequest {
	s.TaskStatus = &v
	return s
}

func (s *QueryTicketsRequest) SetTouchId(v int64) *QueryTicketsRequest {
	s.TouchId = &v
	return s
}

func (s *QueryTicketsRequest) Validate() error {
	return dara.Validate(s)
}
