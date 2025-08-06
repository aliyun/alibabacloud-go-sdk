// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomTemplateAndGroupConsoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCustomTemplateAndGroupConsoleResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateCustomTemplateAndGroupConsoleResponseBody
	GetResult() *bool
}

type UpdateCustomTemplateAndGroupConsoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateCustomTemplateAndGroupConsoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomTemplateAndGroupConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomTemplateAndGroupConsoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomTemplateAndGroupConsoleResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateCustomTemplateAndGroupConsoleResponseBody) SetRequestId(v string) *UpdateCustomTemplateAndGroupConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomTemplateAndGroupConsoleResponseBody) SetResult(v bool) *UpdateCustomTemplateAndGroupConsoleResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateCustomTemplateAndGroupConsoleResponseBody) Validate() error {
	return dara.Validate(s)
}
