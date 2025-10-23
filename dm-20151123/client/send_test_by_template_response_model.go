// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendTestByTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendTestByTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendTestByTemplateResponse
	GetStatusCode() *int32
	SetBody(v *SendTestByTemplateResponseBody) *SendTestByTemplateResponse
	GetBody() *SendTestByTemplateResponseBody
}

type SendTestByTemplateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendTestByTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendTestByTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s SendTestByTemplateResponse) GoString() string {
	return s.String()
}

func (s *SendTestByTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendTestByTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendTestByTemplateResponse) GetBody() *SendTestByTemplateResponseBody {
	return s.Body
}

func (s *SendTestByTemplateResponse) SetHeaders(v map[string]*string) *SendTestByTemplateResponse {
	s.Headers = v
	return s
}

func (s *SendTestByTemplateResponse) SetStatusCode(v int32) *SendTestByTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SendTestByTemplateResponse) SetBody(v *SendTestByTemplateResponseBody) *SendTestByTemplateResponse {
	s.Body = v
	return s
}

func (s *SendTestByTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
