// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertWebhookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertWebhookId(v string) *CreateAlertWebhookResponseBody
	GetAlertWebhookId() *string
	SetRequestId(v string) *CreateAlertWebhookResponseBody
	GetRequestId() *string
}

type CreateAlertWebhookResponseBody struct {
	// example:
	//
	// test
	AlertWebhookId *string `json:"alertWebhookId,omitempty" xml:"alertWebhookId,omitempty"`
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateAlertWebhookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertWebhookResponseBody) GetAlertWebhookId() *string {
	return s.AlertWebhookId
}

func (s *CreateAlertWebhookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAlertWebhookResponseBody) SetAlertWebhookId(v string) *CreateAlertWebhookResponseBody {
	s.AlertWebhookId = &v
	return s
}

func (s *CreateAlertWebhookResponseBody) SetRequestId(v string) *CreateAlertWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlertWebhookResponseBody) Validate() error {
	return dara.Validate(s)
}
