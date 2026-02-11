// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAlertTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableAlertTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableAlertTemplateResponseBody
	GetSuccess() *bool
}

type DisableAlertTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableAlertTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableAlertTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAlertTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableAlertTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableAlertTemplateResponseBody) SetRequestId(v string) *DisableAlertTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAlertTemplateResponseBody) SetSuccess(v bool) *DisableAlertTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DisableAlertTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
