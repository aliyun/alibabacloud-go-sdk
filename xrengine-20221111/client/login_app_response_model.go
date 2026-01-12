// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LoginAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LoginAppResponse
	GetStatusCode() *int32
	SetBody(v *LoginAppResponseBody) *LoginAppResponse
	GetBody() *LoginAppResponseBody
}

type LoginAppResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LoginAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LoginAppResponse) String() string {
	return dara.Prettify(s)
}

func (s LoginAppResponse) GoString() string {
	return s.String()
}

func (s *LoginAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LoginAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LoginAppResponse) GetBody() *LoginAppResponseBody {
	return s.Body
}

func (s *LoginAppResponse) SetHeaders(v map[string]*string) *LoginAppResponse {
	s.Headers = v
	return s
}

func (s *LoginAppResponse) SetStatusCode(v int32) *LoginAppResponse {
	s.StatusCode = &v
	return s
}

func (s *LoginAppResponse) SetBody(v *LoginAppResponseBody) *LoginAppResponse {
	s.Body = v
	return s
}

func (s *LoginAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
