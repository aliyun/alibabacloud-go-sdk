// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSystemTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSystemTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSystemTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetSystemTemplateResponseBody) *GetSystemTemplateResponse
	GetBody() *GetSystemTemplateResponseBody
}

type GetSystemTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSystemTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSystemTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSystemTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetSystemTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSystemTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSystemTemplateResponse) GetBody() *GetSystemTemplateResponseBody {
	return s.Body
}

func (s *GetSystemTemplateResponse) SetHeaders(v map[string]*string) *GetSystemTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetSystemTemplateResponse) SetStatusCode(v int32) *GetSystemTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSystemTemplateResponse) SetBody(v *GetSystemTemplateResponseBody) *GetSystemTemplateResponse {
	s.Body = v
	return s
}

func (s *GetSystemTemplateResponse) Validate() error {
	return dara.Validate(s)
}
