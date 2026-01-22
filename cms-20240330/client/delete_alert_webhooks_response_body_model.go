// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertWebhooksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAlertWebhooksResponseBody
	GetRequestId() *string
}

type DeleteAlertWebhooksResponseBody struct {
	// example:
	//
	// E5B1D3D4-BB28-5996-8AD2-***********
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAlertWebhooksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertWebhooksResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertWebhooksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlertWebhooksResponseBody) SetRequestId(v string) *DeleteAlertWebhooksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlertWebhooksResponseBody) Validate() error {
	return dara.Validate(s)
}
