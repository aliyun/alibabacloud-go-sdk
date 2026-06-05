// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetAppTemplateResponseBody) *GetAppTemplateResponse
	GetBody() *GetAppTemplateResponseBody
}

type GetAppTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetAppTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppTemplateResponse) GetBody() *GetAppTemplateResponseBody {
	return s.Body
}

func (s *GetAppTemplateResponse) SetHeaders(v map[string]*string) *GetAppTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetAppTemplateResponse) SetStatusCode(v int32) *GetAppTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppTemplateResponse) SetBody(v *GetAppTemplateResponseBody) *GetAppTemplateResponse {
	s.Body = v
	return s
}

func (s *GetAppTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
