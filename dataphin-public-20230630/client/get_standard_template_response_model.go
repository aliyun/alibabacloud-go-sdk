// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStandardTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStandardTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetStandardTemplateResponseBody) *GetStandardTemplateResponse
	GetBody() *GetStandardTemplateResponseBody
}

type GetStandardTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStandardTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStandardTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStandardTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStandardTemplateResponse) GetBody() *GetStandardTemplateResponseBody {
	return s.Body
}

func (s *GetStandardTemplateResponse) SetHeaders(v map[string]*string) *GetStandardTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetStandardTemplateResponse) SetStatusCode(v int32) *GetStandardTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStandardTemplateResponse) SetBody(v *GetStandardTemplateResponseBody) *GetStandardTemplateResponse {
	s.Body = v
	return s
}

func (s *GetStandardTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
