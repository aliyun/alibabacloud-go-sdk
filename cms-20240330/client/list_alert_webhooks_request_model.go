// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertWebhooksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListAlertWebhooksRequest
	GetName() *string
	SetPageNumber(v int64) *ListAlertWebhooksRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAlertWebhooksRequest
	GetPageSize() *int64
	SetWebhookIds(v []*string) *ListAlertWebhooksRequest
	GetWebhookIds() []*string
}

type ListAlertWebhooksRequest struct {
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize   *int64    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	WebhookIds []*string `json:"webhookIds,omitempty" xml:"webhookIds,omitempty" type:"Repeated"`
}

func (s ListAlertWebhooksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertWebhooksRequest) GoString() string {
	return s.String()
}

func (s *ListAlertWebhooksRequest) GetName() *string {
	return s.Name
}

func (s *ListAlertWebhooksRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAlertWebhooksRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAlertWebhooksRequest) GetWebhookIds() []*string {
	return s.WebhookIds
}

func (s *ListAlertWebhooksRequest) SetName(v string) *ListAlertWebhooksRequest {
	s.Name = &v
	return s
}

func (s *ListAlertWebhooksRequest) SetPageNumber(v int64) *ListAlertWebhooksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertWebhooksRequest) SetPageSize(v int64) *ListAlertWebhooksRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertWebhooksRequest) SetWebhookIds(v []*string) *ListAlertWebhooksRequest {
	s.WebhookIds = v
	return s
}

func (s *ListAlertWebhooksRequest) Validate() error {
	return dara.Validate(s)
}
