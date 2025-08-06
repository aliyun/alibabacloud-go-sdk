// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateGroupConsoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTemplateGroupConsoleResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteTemplateGroupConsoleResponseBody
	GetResult() *bool
}

type DeleteTemplateGroupConsoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteTemplateGroupConsoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateGroupConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateGroupConsoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTemplateGroupConsoleResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteTemplateGroupConsoleResponseBody) SetRequestId(v string) *DeleteTemplateGroupConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTemplateGroupConsoleResponseBody) SetResult(v bool) *DeleteTemplateGroupConsoleResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteTemplateGroupConsoleResponseBody) Validate() error {
	return dara.Validate(s)
}
