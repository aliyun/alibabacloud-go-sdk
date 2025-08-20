// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ValidateTemplateResponseBody) *ValidateTemplateResponse
	GetBody() *ValidateTemplateResponseBody
}

type ValidateTemplateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateTemplateResponse) GoString() string {
	return s.String()
}

func (s *ValidateTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateTemplateResponse) GetBody() *ValidateTemplateResponseBody {
	return s.Body
}

func (s *ValidateTemplateResponse) SetHeaders(v map[string]*string) *ValidateTemplateResponse {
	s.Headers = v
	return s
}

func (s *ValidateTemplateResponse) SetStatusCode(v int32) *ValidateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateTemplateResponse) SetBody(v *ValidateTemplateResponseBody) *ValidateTemplateResponse {
	s.Body = v
	return s
}

func (s *ValidateTemplateResponse) Validate() error {
	return dara.Validate(s)
}
