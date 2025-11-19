// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVodTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVodTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVodTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetVodTemplateResponseBody) *GetVodTemplateResponse
	GetBody() *GetVodTemplateResponseBody
}

type GetVodTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVodTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVodTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVodTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetVodTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVodTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVodTemplateResponse) GetBody() *GetVodTemplateResponseBody {
	return s.Body
}

func (s *GetVodTemplateResponse) SetHeaders(v map[string]*string) *GetVodTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetVodTemplateResponse) SetStatusCode(v int32) *GetVodTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVodTemplateResponse) SetBody(v *GetVodTemplateResponseBody) *GetVodTemplateResponse {
	s.Body = v
	return s
}

func (s *GetVodTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
