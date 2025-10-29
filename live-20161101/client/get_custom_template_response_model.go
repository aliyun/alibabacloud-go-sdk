// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomTemplateResponseBody) *GetCustomTemplateResponse
	GetBody() *GetCustomTemplateResponseBody
}

type GetCustomTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetCustomTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomTemplateResponse) GetBody() *GetCustomTemplateResponseBody {
	return s.Body
}

func (s *GetCustomTemplateResponse) SetHeaders(v map[string]*string) *GetCustomTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetCustomTemplateResponse) SetStatusCode(v int32) *GetCustomTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomTemplateResponse) SetBody(v *GetCustomTemplateResponseBody) *GetCustomTemplateResponse {
	s.Body = v
	return s
}

func (s *GetCustomTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
