// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertWebhookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAlertWebhookResponseBody
	GetRequestId() *string
}

type UpdateAlertWebhookResponseBody struct {
	// example:
	//
	// 8A33DBEA-*****-*****-*****-*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAlertWebhookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertWebhookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlertWebhookResponseBody) SetRequestId(v string) *UpdateAlertWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlertWebhookResponseBody) Validate() error {
	return dara.Validate(s)
}
