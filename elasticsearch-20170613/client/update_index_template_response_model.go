// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIndexTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIndexTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIndexTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIndexTemplateResponseBody) *UpdateIndexTemplateResponse
	GetBody() *UpdateIndexTemplateResponseBody
}

type UpdateIndexTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIndexTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIndexTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIndexTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateIndexTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIndexTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIndexTemplateResponse) GetBody() *UpdateIndexTemplateResponseBody {
	return s.Body
}

func (s *UpdateIndexTemplateResponse) SetHeaders(v map[string]*string) *UpdateIndexTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateIndexTemplateResponse) SetStatusCode(v int32) *UpdateIndexTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIndexTemplateResponse) SetBody(v *UpdateIndexTemplateResponseBody) *UpdateIndexTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateIndexTemplateResponse) Validate() error {
	return dara.Validate(s)
}
