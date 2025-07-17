// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebhookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *CreateWebhookResponseBody
	GetContactId() *string
	SetRequestId(v string) *CreateWebhookResponseBody
	GetRequestId() *string
}

type CreateWebhookResponseBody struct {
	// The ID of the contact for webhook alerts.
	//
	// example:
	//
	// 48716
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 16AF921B-8187-489F-9913-43C808B4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWebhookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWebhookResponseBody) GetContactId() *string {
	return s.ContactId
}

func (s *CreateWebhookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWebhookResponseBody) SetContactId(v string) *CreateWebhookResponseBody {
	s.ContactId = &v
	return s
}

func (s *CreateWebhookResponseBody) SetRequestId(v string) *CreateWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWebhookResponseBody) Validate() error {
	return dara.Validate(s)
}
