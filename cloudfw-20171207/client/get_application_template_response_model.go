// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationTemplateResponseBody) *GetApplicationTemplateResponse
	GetBody() *GetApplicationTemplateResponseBody
}

type GetApplicationTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationTemplateResponse) GetBody() *GetApplicationTemplateResponseBody {
	return s.Body
}

func (s *GetApplicationTemplateResponse) SetHeaders(v map[string]*string) *GetApplicationTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationTemplateResponse) SetStatusCode(v int32) *GetApplicationTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationTemplateResponse) SetBody(v *GetApplicationTemplateResponseBody) *GetApplicationTemplateResponse {
	s.Body = v
	return s
}

func (s *GetApplicationTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
