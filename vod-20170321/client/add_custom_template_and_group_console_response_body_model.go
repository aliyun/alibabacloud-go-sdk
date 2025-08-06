// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomTemplateAndGroupConsoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddCustomTemplateAndGroupConsoleResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddCustomTemplateAndGroupConsoleResponseBody
	GetResult() *bool
	SetTemplateGroup(v *AddCustomTemplateAndGroupConsoleResponseBodyTemplateGroup) *AddCustomTemplateAndGroupConsoleResponseBody
	GetTemplateGroup() *AddCustomTemplateAndGroupConsoleResponseBodyTemplateGroup
}

type AddCustomTemplateAndGroupConsoleResponseBody struct {
	RequestId     *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result        *bool                                                      `json:"Result,omitempty" xml:"Result,omitempty"`
	TemplateGroup *AddCustomTemplateAndGroupConsoleResponseBodyTemplateGroup `json:"TemplateGroup,omitempty" xml:"TemplateGroup,omitempty" type:"Struct"`
}

func (s AddCustomTemplateAndGroupConsoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCustomTemplateAndGroupConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomTemplateAndGroupConsoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCustomTemplateAndGroupConsoleResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddCustomTemplateAndGroupConsoleResponseBody) GetTemplateGroup() *AddCustomTemplateAndGroupConsoleResponseBodyTemplateGroup {
	return s.TemplateGroup
}

func (s *AddCustomTemplateAndGroupConsoleResponseBody) SetRequestId(v string) *AddCustomTemplateAndGroupConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCustomTemplateAndGroupConsoleResponseBody) SetResult(v bool) *AddCustomTemplateAndGroupConsoleResponseBody {
	s.Result = &v
	return s
}

func (s *AddCustomTemplateAndGroupConsoleResponseBody) SetTemplateGroup(v *AddCustomTemplateAndGroupConsoleResponseBodyTemplateGroup) *AddCustomTemplateAndGroupConsoleResponseBody {
	s.TemplateGroup = v
	return s
}

func (s *AddCustomTemplateAndGroupConsoleResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddCustomTemplateAndGroupConsoleResponseBodyTemplateGroup struct {
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
}

func (s AddCustomTemplateAndGroupConsoleResponseBodyTemplateGroup) String() string {
	return dara.Prettify(s)
}

func (s AddCustomTemplateAndGroupConsoleResponseBodyTemplateGroup) GoString() string {
	return s.String()
}

func (s *AddCustomTemplateAndGroupConsoleResponseBodyTemplateGroup) GetTemplateGroupId() *string {
	return s.TemplateGroupId
}

func (s *AddCustomTemplateAndGroupConsoleResponseBodyTemplateGroup) SetTemplateGroupId(v string) *AddCustomTemplateAndGroupConsoleResponseBodyTemplateGroup {
	s.TemplateGroupId = &v
	return s
}

func (s *AddCustomTemplateAndGroupConsoleResponseBodyTemplateGroup) Validate() error {
	return dara.Validate(s)
}
