// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAlertTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAlertTemplateResponseBody
	GetSuccess() *bool
}

type UpdateAlertTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAlertTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlertTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAlertTemplateResponseBody) SetRequestId(v string) *UpdateAlertTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlertTemplateResponseBody) SetSuccess(v bool) *UpdateAlertTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAlertTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
