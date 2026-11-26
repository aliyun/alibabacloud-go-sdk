// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBillingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListBillingRequest
	GetAppToken() *string
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
	SetOpsUserCode(v string) *ListBillingRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListBillingRequest
	GetOpsUserId() *string
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
	// The platform application token issued by the console. If you specify this parameter, the request is executed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The unique identifier of the business. When bizType is set to LibraryChat, bizId refers to the ID of the document library.
	//
	// example:
	//
	// exampleBizId
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// The business type. Valid values: LlmChat (model Q&A) and LibraryChat (document library Q&A).
	//
	// example:
	//
	// string_value
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// The actual end timestamp of the live stream. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Specifies whether to filter out bills with zero credit consumption. Default value: true.
	//
	// example:
	//
	// true
	IgnoreZero *bool `json:"ignoreZero,omitempty" xml:"ignoreZero,omitempty"`
	// The operation item. Valid values:
	//
	// - start: indicates the creation of a task. This is the default value and usually does not need to be explicitly specified.
	//
	// - stop: stops a real-time meeting task. This corresponds to the creation of a real-time meeting and is set to stop to trigger a call after the meeting ends. This value is used in real-time meeting scenarios.
	//
	// Note: When you end a real-time recording, you must set this parameter to stop.
	//
	// example:
	//
	// string_value
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
	// The code of the proxy user, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The ID of the proxy user, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
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
	// The start time of the query. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The task status. A status of Running is returned immediately after the task is submitted.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tenant ID. This is a common parameter. For winnexo-cli, this parameter is explicitly passed by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The user ID, which is the WINNEXO platform user ID. This parameter is optional for filtering.
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

func (s *ListBillingRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *ListBillingRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListBillingRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *ListBillingRequest) SetAppToken(v string) *ListBillingRequest {
	s.AppToken = &v
	return s
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

func (s *ListBillingRequest) SetOpsUserCode(v string) *ListBillingRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListBillingRequest) SetOpsUserId(v string) *ListBillingRequest {
	s.OpsUserId = &v
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
