// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomScriptTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomScriptTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomScriptTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomScriptTemplateResponseBody) *CreateCustomScriptTemplateResponse
	GetBody() *CreateCustomScriptTemplateResponseBody
}

type CreateCustomScriptTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomScriptTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomScriptTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomScriptTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomScriptTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomScriptTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomScriptTemplateResponse) GetBody() *CreateCustomScriptTemplateResponseBody {
	return s.Body
}

func (s *CreateCustomScriptTemplateResponse) SetHeaders(v map[string]*string) *CreateCustomScriptTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomScriptTemplateResponse) SetStatusCode(v int32) *CreateCustomScriptTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomScriptTemplateResponse) SetBody(v *CreateCustomScriptTemplateResponseBody) *CreateCustomScriptTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateCustomScriptTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
