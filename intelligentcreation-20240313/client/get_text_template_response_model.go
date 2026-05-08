// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTextTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTextTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetTextTemplateResponseBody) *GetTextTemplateResponse
	GetBody() *GetTextTemplateResponseBody
}

type GetTextTemplateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTextTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTextTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTextTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTextTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTextTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTextTemplateResponse) GetBody() *GetTextTemplateResponseBody {
	return s.Body
}

func (s *GetTextTemplateResponse) SetHeaders(v map[string]*string) *GetTextTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTextTemplateResponse) SetStatusCode(v int32) *GetTextTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTextTemplateResponse) SetBody(v *GetTextTemplateResponseBody) *GetTextTemplateResponse {
	s.Body = v
	return s
}

func (s *GetTextTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
