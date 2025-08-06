// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTemplateConsoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomTemplateConsoleResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteCustomTemplateConsoleResponseBody
	GetResult() *bool
}

type DeleteCustomTemplateConsoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteCustomTemplateConsoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTemplateConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomTemplateConsoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomTemplateConsoleResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteCustomTemplateConsoleResponseBody) SetRequestId(v string) *DeleteCustomTemplateConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomTemplateConsoleResponseBody) SetResult(v bool) *DeleteCustomTemplateConsoleResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteCustomTemplateConsoleResponseBody) Validate() error {
	return dara.Validate(s)
}
