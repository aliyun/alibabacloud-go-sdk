// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAlertTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAlertTemplateResponseBody
	GetSuccess() *bool
}

type DeleteAlertTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAlertTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlertTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAlertTemplateResponseBody) SetRequestId(v string) *DeleteAlertTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlertTemplateResponseBody) SetSuccess(v bool) *DeleteAlertTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAlertTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
