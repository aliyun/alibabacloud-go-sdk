// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTicketsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaseId(v int64) *QueryTicketsShrinkRequest
	GetCaseId() *int64
	SetCaseStatus(v int32) *QueryTicketsShrinkRequest
	GetCaseStatus() *int32
	SetCaseType(v int32) *QueryTicketsShrinkRequest
	GetCaseType() *int32
	SetChannelId(v string) *QueryTicketsShrinkRequest
	GetChannelId() *string
	SetChannelType(v int32) *QueryTicketsShrinkRequest
	GetChannelType() *int32
	SetCurrentPage(v int32) *QueryTicketsShrinkRequest
	GetCurrentPage() *int32
	SetDealId(v int64) *QueryTicketsShrinkRequest
	GetDealId() *int64
	SetExtraShrink(v string) *QueryTicketsShrinkRequest
	GetExtraShrink() *string
	SetInstanceId(v string) *QueryTicketsShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *QueryTicketsShrinkRequest
	GetPageSize() *int32
	SetSrType(v int64) *QueryTicketsShrinkRequest
	GetSrType() *int64
	SetTaskStatus(v int32) *QueryTicketsShrinkRequest
	GetTaskStatus() *int32
	SetTouchId(v int64) *QueryTicketsShrinkRequest
	GetTouchId() *int64
}

type QueryTicketsShrinkRequest struct {
	// example:
	//
	// 223468****
	CaseId *int64 `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	// example:
	//
	// 2
	CaseStatus *int32 `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	// example:
	//
	// 1223
	CaseType *int32 `json:"CaseType,omitempty" xml:"CaseType,omitempty"`
	// example:
	//
	// 02acfefd3fa14049826ac1a89e1xxxxx
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1
	ChannelType *int32 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 123456
	DealId      *int64  `json:"DealId,omitempty" xml:"DealId,omitempty"`
	ExtraShrink *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 29506
	SrType *int64 `json:"SrType,omitempty" xml:"SrType,omitempty"`
	// example:
	//
	// 3
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 15030609
	TouchId *int64 `json:"TouchId,omitempty" xml:"TouchId,omitempty"`
}

func (s QueryTicketsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTicketsShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryTicketsShrinkRequest) GetCaseId() *int64 {
	return s.CaseId
}

func (s *QueryTicketsShrinkRequest) GetCaseStatus() *int32 {
	return s.CaseStatus
}

func (s *QueryTicketsShrinkRequest) GetCaseType() *int32 {
	return s.CaseType
}

func (s *QueryTicketsShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *QueryTicketsShrinkRequest) GetChannelType() *int32 {
	return s.ChannelType
}

func (s *QueryTicketsShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryTicketsShrinkRequest) GetDealId() *int64 {
	return s.DealId
}

func (s *QueryTicketsShrinkRequest) GetExtraShrink() *string {
	return s.ExtraShrink
}

func (s *QueryTicketsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryTicketsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTicketsShrinkRequest) GetSrType() *int64 {
	return s.SrType
}

func (s *QueryTicketsShrinkRequest) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *QueryTicketsShrinkRequest) GetTouchId() *int64 {
	return s.TouchId
}

func (s *QueryTicketsShrinkRequest) SetCaseId(v int64) *QueryTicketsShrinkRequest {
	s.CaseId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetCaseStatus(v int32) *QueryTicketsShrinkRequest {
	s.CaseStatus = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetCaseType(v int32) *QueryTicketsShrinkRequest {
	s.CaseType = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetChannelId(v string) *QueryTicketsShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetChannelType(v int32) *QueryTicketsShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetCurrentPage(v int32) *QueryTicketsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetDealId(v int64) *QueryTicketsShrinkRequest {
	s.DealId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetExtraShrink(v string) *QueryTicketsShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetInstanceId(v string) *QueryTicketsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetPageSize(v int32) *QueryTicketsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetSrType(v int64) *QueryTicketsShrinkRequest {
	s.SrType = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetTaskStatus(v int32) *QueryTicketsShrinkRequest {
	s.TaskStatus = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetTouchId(v int64) *QueryTicketsShrinkRequest {
	s.TouchId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
