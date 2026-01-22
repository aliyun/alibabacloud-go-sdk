// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertWebhooksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListAlertWebhooksShrinkRequest
	GetName() *string
	SetPageNumber(v int64) *ListAlertWebhooksShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAlertWebhooksShrinkRequest
	GetPageSize() *int64
	SetWebhookIdsShrink(v string) *ListAlertWebhooksShrinkRequest
	GetWebhookIdsShrink() *string
}

type ListAlertWebhooksShrinkRequest struct {
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
	PageSize         *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	WebhookIdsShrink *string `json:"webhookIds,omitempty" xml:"webhookIds,omitempty"`
}

func (s ListAlertWebhooksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertWebhooksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAlertWebhooksShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListAlertWebhooksShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAlertWebhooksShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAlertWebhooksShrinkRequest) GetWebhookIdsShrink() *string {
	return s.WebhookIdsShrink
}

func (s *ListAlertWebhooksShrinkRequest) SetName(v string) *ListAlertWebhooksShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListAlertWebhooksShrinkRequest) SetPageNumber(v int64) *ListAlertWebhooksShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertWebhooksShrinkRequest) SetPageSize(v int64) *ListAlertWebhooksShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertWebhooksShrinkRequest) SetWebhookIdsShrink(v string) *ListAlertWebhooksShrinkRequest {
	s.WebhookIdsShrink = &v
	return s
}

func (s *ListAlertWebhooksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
