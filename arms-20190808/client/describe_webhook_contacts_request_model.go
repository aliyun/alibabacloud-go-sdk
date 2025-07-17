// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebhookContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactIds(v string) *DescribeWebhookContactsRequest
	GetContactIds() *string
	SetPage(v int64) *DescribeWebhookContactsRequest
	GetPage() *int64
	SetSize(v int64) *DescribeWebhookContactsRequest
	GetSize() *int64
	SetWebhookName(v string) *DescribeWebhookContactsRequest
	GetWebhookName() *string
}

type DescribeWebhookContactsRequest struct {
	// The ID of the alert contact.
	//
	// example:
	//
	// 123
	ContactIds *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of alert contacts displayed on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The name of the webhook alert contact.
	//
	// example:
	//
	// Webhook name
	WebhookName *string `json:"WebhookName,omitempty" xml:"WebhookName,omitempty"`
}

func (s DescribeWebhookContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebhookContactsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebhookContactsRequest) GetContactIds() *string {
	return s.ContactIds
}

func (s *DescribeWebhookContactsRequest) GetPage() *int64 {
	return s.Page
}

func (s *DescribeWebhookContactsRequest) GetSize() *int64 {
	return s.Size
}

func (s *DescribeWebhookContactsRequest) GetWebhookName() *string {
	return s.WebhookName
}

func (s *DescribeWebhookContactsRequest) SetContactIds(v string) *DescribeWebhookContactsRequest {
	s.ContactIds = &v
	return s
}

func (s *DescribeWebhookContactsRequest) SetPage(v int64) *DescribeWebhookContactsRequest {
	s.Page = &v
	return s
}

func (s *DescribeWebhookContactsRequest) SetSize(v int64) *DescribeWebhookContactsRequest {
	s.Size = &v
	return s
}

func (s *DescribeWebhookContactsRequest) SetWebhookName(v string) *DescribeWebhookContactsRequest {
	s.WebhookName = &v
	return s
}

func (s *DescribeWebhookContactsRequest) Validate() error {
	return dara.Validate(s)
}
