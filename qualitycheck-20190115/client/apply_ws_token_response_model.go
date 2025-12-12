// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyWsTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyWsTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyWsTokenResponse
	GetStatusCode() *int32
	SetBody(v *ApplyWsTokenResponseBody) *ApplyWsTokenResponse
	GetBody() *ApplyWsTokenResponseBody
}

type ApplyWsTokenResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyWsTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyWsTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyWsTokenResponse) GoString() string {
	return s.String()
}

func (s *ApplyWsTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyWsTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyWsTokenResponse) GetBody() *ApplyWsTokenResponseBody {
	return s.Body
}

func (s *ApplyWsTokenResponse) SetHeaders(v map[string]*string) *ApplyWsTokenResponse {
	s.Headers = v
	return s
}

func (s *ApplyWsTokenResponse) SetStatusCode(v int32) *ApplyWsTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyWsTokenResponse) SetBody(v *ApplyWsTokenResponseBody) *ApplyWsTokenResponse {
	s.Body = v
	return s
}

func (s *ApplyWsTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
