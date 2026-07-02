// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAuthResponse
	GetStatusCode() *int32
	SetBody(v *QueryAuthResponseBody) *QueryAuthResponse
	GetBody() *QueryAuthResponseBody
}

type QueryAuthResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAuthResponse) GoString() string {
	return s.String()
}

func (s *QueryAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAuthResponse) GetBody() *QueryAuthResponseBody {
	return s.Body
}

func (s *QueryAuthResponse) SetHeaders(v map[string]*string) *QueryAuthResponse {
	s.Headers = v
	return s
}

func (s *QueryAuthResponse) SetStatusCode(v int32) *QueryAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAuthResponse) SetBody(v *QueryAuthResponseBody) *QueryAuthResponse {
	s.Body = v
	return s
}

func (s *QueryAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
