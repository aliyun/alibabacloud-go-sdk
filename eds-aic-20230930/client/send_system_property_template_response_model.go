// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSystemPropertyTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendSystemPropertyTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendSystemPropertyTemplateResponse
	GetStatusCode() *int32
	SetBody(v *SendSystemPropertyTemplateResponseBody) *SendSystemPropertyTemplateResponse
	GetBody() *SendSystemPropertyTemplateResponseBody
}

type SendSystemPropertyTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendSystemPropertyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendSystemPropertyTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s SendSystemPropertyTemplateResponse) GoString() string {
	return s.String()
}

func (s *SendSystemPropertyTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendSystemPropertyTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendSystemPropertyTemplateResponse) GetBody() *SendSystemPropertyTemplateResponseBody {
	return s.Body
}

func (s *SendSystemPropertyTemplateResponse) SetHeaders(v map[string]*string) *SendSystemPropertyTemplateResponse {
	s.Headers = v
	return s
}

func (s *SendSystemPropertyTemplateResponse) SetStatusCode(v int32) *SendSystemPropertyTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SendSystemPropertyTemplateResponse) SetBody(v *SendSystemPropertyTemplateResponseBody) *SendSystemPropertyTemplateResponse {
	s.Body = v
	return s
}

func (s *SendSystemPropertyTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
