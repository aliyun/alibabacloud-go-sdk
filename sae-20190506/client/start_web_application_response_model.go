// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartWebApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartWebApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartWebApplicationResponse
	GetStatusCode() *int32
	SetBody(v *WebApplicationBody) *StartWebApplicationResponse
	GetBody() *WebApplicationBody
}

type StartWebApplicationResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebApplicationBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartWebApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s StartWebApplicationResponse) GoString() string {
	return s.String()
}

func (s *StartWebApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartWebApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartWebApplicationResponse) GetBody() *WebApplicationBody {
	return s.Body
}

func (s *StartWebApplicationResponse) SetHeaders(v map[string]*string) *StartWebApplicationResponse {
	s.Headers = v
	return s
}

func (s *StartWebApplicationResponse) SetStatusCode(v int32) *StartWebApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *StartWebApplicationResponse) SetBody(v *WebApplicationBody) *StartWebApplicationResponse {
	s.Body = v
	return s
}

func (s *StartWebApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
