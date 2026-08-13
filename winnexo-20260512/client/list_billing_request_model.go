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
	// 业务来源ID（可选筛选）
	//
	// example:
	//
	// exampleBizId
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// 业务来源类型（可选筛选）
	//
	// example:
	//
	// string_value
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 结束时间范围，ISO-8601 字符串，如 2026-08-05T16:30:00.000Z
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 是否过滤 credit 消耗为 0 的账单，默认 true（过滤）
	//
	// example:
	//
	// true
	IgnoreZero *bool `json:"ignoreZero,omitempty" xml:"ignoreZero,omitempty"`
	// 操作类型（可选筛选）
	//
	// example:
	//
	// string_value
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
	// 页码
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页条数
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 开始时间范围，ISO-8601 字符串，如 2026-08-05T16:30:00.000Z
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 状态（可选筛选）
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 用户ID（WINNEXO 平台用户ID，可选筛选）
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
