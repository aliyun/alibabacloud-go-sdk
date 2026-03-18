// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccessLogAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAccessLogAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAccessLogAuthResponse
	GetStatusCode() *int32
	SetBody(v *CheckAccessLogAuthResponseBody) *CheckAccessLogAuthResponse
	GetBody() *CheckAccessLogAuthResponseBody
}

type CheckAccessLogAuthResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAccessLogAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAccessLogAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAccessLogAuthResponse) GoString() string {
	return s.String()
}

func (s *CheckAccessLogAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAccessLogAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAccessLogAuthResponse) GetBody() *CheckAccessLogAuthResponseBody {
	return s.Body
}

func (s *CheckAccessLogAuthResponse) SetHeaders(v map[string]*string) *CheckAccessLogAuthResponse {
	s.Headers = v
	return s
}

func (s *CheckAccessLogAuthResponse) SetStatusCode(v int32) *CheckAccessLogAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAccessLogAuthResponse) SetBody(v *CheckAccessLogAuthResponseBody) *CheckAccessLogAuthResponse {
	s.Body = v
	return s
}

func (s *CheckAccessLogAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
