// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMonitorWebhooksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ListMonitorWebhooksRequest
	GetEnterpriseId() *int64
	SetName(v string) *ListMonitorWebhooksRequest
	GetName() *string
	SetOrderBy(v string) *ListMonitorWebhooksRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListMonitorWebhooksRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListMonitorWebhooksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMonitorWebhooksRequest
	GetPageSize() *int32
	SetWebhookType(v string) *ListMonitorWebhooksRequest
	GetWebhookType() *string
}

type ListMonitorWebhooksRequest struct {
	// example:
	//
	// 2
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// example:
	//
	// mask_detect
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// id
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// DINGDING
	WebhookType *string `json:"webhookType,omitempty" xml:"webhookType,omitempty"`
}

func (s ListMonitorWebhooksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMonitorWebhooksRequest) GoString() string {
	return s.String()
}

func (s *ListMonitorWebhooksRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListMonitorWebhooksRequest) GetName() *string {
	return s.Name
}

func (s *ListMonitorWebhooksRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMonitorWebhooksRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListMonitorWebhooksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMonitorWebhooksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMonitorWebhooksRequest) GetWebhookType() *string {
	return s.WebhookType
}

func (s *ListMonitorWebhooksRequest) SetEnterpriseId(v int64) *ListMonitorWebhooksRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListMonitorWebhooksRequest) SetName(v string) *ListMonitorWebhooksRequest {
	s.Name = &v
	return s
}

func (s *ListMonitorWebhooksRequest) SetOrderBy(v string) *ListMonitorWebhooksRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMonitorWebhooksRequest) SetOrderDirection(v string) *ListMonitorWebhooksRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListMonitorWebhooksRequest) SetPageNumber(v int32) *ListMonitorWebhooksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMonitorWebhooksRequest) SetPageSize(v int32) *ListMonitorWebhooksRequest {
	s.PageSize = &v
	return s
}

func (s *ListMonitorWebhooksRequest) SetWebhookType(v string) *ListMonitorWebhooksRequest {
	s.WebhookType = &v
	return s
}

func (s *ListMonitorWebhooksRequest) Validate() error {
	return dara.Validate(s)
}
