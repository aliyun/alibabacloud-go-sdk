// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBillingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListBillingRequest
	GetBizId() *string
	SetBizType(v string) *ListBillingRequest
	GetBizType() *string
	SetEndTime(v string) *ListBillingRequest
	GetEndTime() *string
	SetIgnoreZero(v bool) *ListBillingRequest
	GetIgnoreZero() *bool
	SetOperation(v string) *ListBillingRequest
	GetOperation() *string
	SetPage(v int64) *ListBillingRequest
	GetPage() *int64
	SetPageSize(v int64) *ListBillingRequest
	GetPageSize() *int64
	SetStartTime(v string) *ListBillingRequest
	GetStartTime() *string
	SetStatus(v string) *ListBillingRequest
	GetStatus() *string
	SetTenantId(v string) *ListBillingRequest
	GetTenantId() *string
	SetWnUserId(v string) *ListBillingRequest
	GetWnUserId() *string
}

type ListBillingRequest struct {
	// The unique business identifier. When bizType is LibraryChat, bizId refers to the document library ID.
	//
	// example:
	//
	// exampleBizId
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// The business type. Currently supported values: model Q&A (LlmChat) and document library Q&A (LibraryChat).
	//
	// example:
	//
	// string_value
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// The actual end timestamp of the live stream, in milliseconds.
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Specifies whether to filter out bills with zero credit consumption. Default value: true (filter out).
	//
	// example:
	//
	// true
	IgnoreZero *bool `json:"ignoreZero,omitempty" xml:"ignoreZero,omitempty"`
	// The operation type. Valid values:
	//
	// - start: indicates task creation. This is the default value and does not need to be explicitly set in most cases.
	//
	// - stop: stops a real-time meeting task. This corresponds to the creation of a real-time meeting. After the meeting ends, set this to stop to trigger the call. This is used in real-time meeting scenarios.
	//
	// Note: When ending a real-time recording, you must set this parameter to stop.
	//
	// example:
	//
	// string_value
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page. Default value: 20. Minimum value: 1. Maximum value: 50.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The query start time. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The task status. The status is returned as Running upon submission.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass it explicitly with --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The user ID (WINNEXO platform user ID, optional filter).
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s ListBillingRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBillingRequest) GoString() string {
	return s.String()
}

func (s *ListBillingRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListBillingRequest) GetBizType() *string {
	return s.BizType
}

func (s *ListBillingRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListBillingRequest) GetIgnoreZero() *bool {
	return s.IgnoreZero
}

func (s *ListBillingRequest) GetOperation() *string {
	return s.Operation
}

func (s *ListBillingRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListBillingRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListBillingRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListBillingRequest) GetStatus() *string {
	return s.Status
}

func (s *ListBillingRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListBillingRequest) GetWnUserId() *string {
	return s.WnUserId
}

func (s *ListBillingRequest) SetBizId(v string) *ListBillingRequest {
	s.BizId = &v
	return s
}

func (s *ListBillingRequest) SetBizType(v string) *ListBillingRequest {
	s.BizType = &v
	return s
}

func (s *ListBillingRequest) SetEndTime(v string) *ListBillingRequest {
	s.EndTime = &v
	return s
}

func (s *ListBillingRequest) SetIgnoreZero(v bool) *ListBillingRequest {
	s.IgnoreZero = &v
	return s
}

func (s *ListBillingRequest) SetOperation(v string) *ListBillingRequest {
	s.Operation = &v
	return s
}

func (s *ListBillingRequest) SetPage(v int64) *ListBillingRequest {
	s.Page = &v
	return s
}

func (s *ListBillingRequest) SetPageSize(v int64) *ListBillingRequest {
	s.PageSize = &v
	return s
}

func (s *ListBillingRequest) SetStartTime(v string) *ListBillingRequest {
	s.StartTime = &v
	return s
}

func (s *ListBillingRequest) SetStatus(v string) *ListBillingRequest {
	s.Status = &v
	return s
}

func (s *ListBillingRequest) SetTenantId(v string) *ListBillingRequest {
	s.TenantId = &v
	return s
}

func (s *ListBillingRequest) SetWnUserId(v string) *ListBillingRequest {
	s.WnUserId = &v
	return s
}

func (s *ListBillingRequest) Validate() error {
	return dara.Validate(s)
}
