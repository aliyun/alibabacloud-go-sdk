// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultTemplateGroupConsoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDefaultTemplateGroupConsoleResponseBody
	GetRequestId() *string
	SetResult(v bool) *SetDefaultTemplateGroupConsoleResponseBody
	GetResult() *bool
}

type SetDefaultTemplateGroupConsoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SetDefaultTemplateGroupConsoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultTemplateGroupConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultTemplateGroupConsoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDefaultTemplateGroupConsoleResponseBody) GetResult() *bool {
	return s.Result
}

func (s *SetDefaultTemplateGroupConsoleResponseBody) SetRequestId(v string) *SetDefaultTemplateGroupConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultTemplateGroupConsoleResponseBody) SetResult(v bool) *SetDefaultTemplateGroupConsoleResponseBody {
	s.Result = &v
	return s
}

func (s *SetDefaultTemplateGroupConsoleResponseBody) Validate() error {
	return dara.Validate(s)
}
