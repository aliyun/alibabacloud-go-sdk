// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertNotifyTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertNotifyTemplateId(v string) *CreateAlertNotifyTemplateResponseBody
	GetAlertNotifyTemplateId() *string
	SetRequestId(v string) *CreateAlertNotifyTemplateResponseBody
	GetRequestId() *string
}

type CreateAlertNotifyTemplateResponseBody struct {
	// The ID of the notification template.
	//
	// example:
	//
	// testId
	AlertNotifyTemplateId *string `json:"alertNotifyTemplateId,omitempty" xml:"alertNotifyTemplateId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8A33DBEA-*****-*****-*****-*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateAlertNotifyTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertNotifyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertNotifyTemplateResponseBody) GetAlertNotifyTemplateId() *string {
	return s.AlertNotifyTemplateId
}

func (s *CreateAlertNotifyTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAlertNotifyTemplateResponseBody) SetAlertNotifyTemplateId(v string) *CreateAlertNotifyTemplateResponseBody {
	s.AlertNotifyTemplateId = &v
	return s
}

func (s *CreateAlertNotifyTemplateResponseBody) SetRequestId(v string) *CreateAlertNotifyTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlertNotifyTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
