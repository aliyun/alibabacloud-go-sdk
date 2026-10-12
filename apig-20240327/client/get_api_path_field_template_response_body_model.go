// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiPathFieldTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetApiPathFieldTemplateResponseBody
	GetCode() *string
	SetData(v *GetApiPathFieldTemplateResponseBodyData) *GetApiPathFieldTemplateResponseBody
	GetData() *GetApiPathFieldTemplateResponseBodyData
	SetMessage(v string) *GetApiPathFieldTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetApiPathFieldTemplateResponseBody
	GetRequestId() *string
}

type GetApiPathFieldTemplateResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetApiPathFieldTemplateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetApiPathFieldTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApiPathFieldTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetApiPathFieldTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetApiPathFieldTemplateResponseBody) GetData() *GetApiPathFieldTemplateResponseBodyData {
	return s.Data
}

func (s *GetApiPathFieldTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApiPathFieldTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApiPathFieldTemplateResponseBody) SetCode(v string) *GetApiPathFieldTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetApiPathFieldTemplateResponseBody) SetData(v *GetApiPathFieldTemplateResponseBodyData) *GetApiPathFieldTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetApiPathFieldTemplateResponseBody) SetMessage(v string) *GetApiPathFieldTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetApiPathFieldTemplateResponseBody) SetRequestId(v string) *GetApiPathFieldTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApiPathFieldTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApiPathFieldTemplateResponseBodyData struct {
	// The description text of the field template.
	//
	// example:
	//
	// OpenAI chat completion API..
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The set of fields defined by the field template, grouped by mapping key. Each group contains a set of template field items.
	FieldPaths map[string][]*DataFieldPathsValue `json:"fieldPaths,omitempty" xml:"fieldPaths,omitempty"`
	// The API request path corresponding to the field template.
	//
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The API protocol applicable to the field template, such as openai.
	//
	// example:
	//
	// openai
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The name of the field template.
	//
	// example:
	//
	// OpenAI Chat Completions
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s GetApiPathFieldTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetApiPathFieldTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApiPathFieldTemplateResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetApiPathFieldTemplateResponseBodyData) GetFieldPaths() map[string][]*DataFieldPathsValue {
	return s.FieldPaths
}

func (s *GetApiPathFieldTemplateResponseBodyData) GetPath() *string {
	return s.Path
}

func (s *GetApiPathFieldTemplateResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *GetApiPathFieldTemplateResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetApiPathFieldTemplateResponseBodyData) SetDescription(v string) *GetApiPathFieldTemplateResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetApiPathFieldTemplateResponseBodyData) SetFieldPaths(v map[string][]*DataFieldPathsValue) *GetApiPathFieldTemplateResponseBodyData {
	s.FieldPaths = v
	return s
}

func (s *GetApiPathFieldTemplateResponseBodyData) SetPath(v string) *GetApiPathFieldTemplateResponseBodyData {
	s.Path = &v
	return s
}

func (s *GetApiPathFieldTemplateResponseBodyData) SetProtocol(v string) *GetApiPathFieldTemplateResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetApiPathFieldTemplateResponseBodyData) SetTemplateName(v string) *GetApiPathFieldTemplateResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *GetApiPathFieldTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
