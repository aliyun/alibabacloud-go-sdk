// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertNotifyTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAlertNotifyTemplateResponseBody
	GetRequestId() *string
}

type UpdateAlertNotifyTemplateResponseBody struct {
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAlertNotifyTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertNotifyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertNotifyTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlertNotifyTemplateResponseBody) SetRequestId(v string) *UpdateAlertNotifyTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlertNotifyTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
