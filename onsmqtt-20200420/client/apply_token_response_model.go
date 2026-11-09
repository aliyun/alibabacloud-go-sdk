// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyTokenResponse
	GetStatusCode() *int32
	SetBody(v *ApplyTokenResponseBody) *ApplyTokenResponse
	GetBody() *ApplyTokenResponseBody
}

type ApplyTokenResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyTokenResponse) GoString() string {
	return s.String()
}

func (s *ApplyTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyTokenResponse) GetBody() *ApplyTokenResponseBody {
	return s.Body
}

func (s *ApplyTokenResponse) SetHeaders(v map[string]*string) *ApplyTokenResponse {
	s.Headers = v
	return s
}

func (s *ApplyTokenResponse) SetStatusCode(v int32) *ApplyTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyTokenResponse) SetBody(v *ApplyTokenResponseBody) *ApplyTokenResponse {
	s.Body = v
	return s
}

func (s *ApplyTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
