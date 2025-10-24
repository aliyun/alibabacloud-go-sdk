// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTemplateResponse
	GetStatusCode() *int32
	SetBody(v *AddTemplateResponseBody) *AddTemplateResponse
	GetBody() *AddTemplateResponseBody
}

type AddTemplateResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTemplateResponse) GetBody() *AddTemplateResponseBody {
	return s.Body
}

func (s *AddTemplateResponse) SetHeaders(v map[string]*string) *AddTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddTemplateResponse) SetStatusCode(v int32) *AddTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTemplateResponse) SetBody(v *AddTemplateResponseBody) *AddTemplateResponse {
	s.Body = v
	return s
}

func (s *AddTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
