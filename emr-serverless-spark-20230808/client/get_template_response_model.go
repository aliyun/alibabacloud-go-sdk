// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetTemplateResponseBody) *GetTemplateResponse
	GetBody() *GetTemplateResponseBody
}

type GetTemplateResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTemplateResponse) GetBody() *GetTemplateResponseBody {
	return s.Body
}

func (s *GetTemplateResponse) SetHeaders(v map[string]*string) *GetTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateResponse) SetStatusCode(v int32) *GetTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateResponse) SetBody(v *GetTemplateResponseBody) *GetTemplateResponse {
	s.Body = v
	return s
}

func (s *GetTemplateResponse) Validate() error {
	return dara.Validate(s)
}
