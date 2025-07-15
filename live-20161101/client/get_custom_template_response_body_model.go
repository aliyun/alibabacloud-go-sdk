// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomTemplate(v string) *GetCustomTemplateResponseBody
	GetCustomTemplate() *string
	SetRequestId(v string) *GetCustomTemplateResponseBody
	GetRequestId() *string
	SetTemplate(v string) *GetCustomTemplateResponseBody
	GetTemplate() *string
}

type GetCustomTemplateResponseBody struct {
	// The configurations of the template.
	//
	// example:
	//
	// {height:1080,scale:[16:9],gop:60,bframes:30,cdesc:h264}
	CustomTemplate *string `json:"CustomTemplate,omitempty" xml:"CustomTemplate,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BC1E78D3-FA8B-4457-DEE2-6093E1232254
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// TestTemplate
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s GetCustomTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomTemplateResponseBody) GetCustomTemplate() *string {
	return s.CustomTemplate
}

func (s *GetCustomTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomTemplateResponseBody) GetTemplate() *string {
	return s.Template
}

func (s *GetCustomTemplateResponseBody) SetCustomTemplate(v string) *GetCustomTemplateResponseBody {
	s.CustomTemplate = &v
	return s
}

func (s *GetCustomTemplateResponseBody) SetRequestId(v string) *GetCustomTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomTemplateResponseBody) SetTemplate(v string) *GetCustomTemplateResponseBody {
	s.Template = &v
	return s
}

func (s *GetCustomTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
