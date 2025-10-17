// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShieldPrecheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ShieldPrecheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ShieldPrecheckResponse
	GetStatusCode() *int32
	SetBody(v *ShieldPrecheckResponseBody) *ShieldPrecheckResponse
	GetBody() *ShieldPrecheckResponseBody
}

type ShieldPrecheckResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShieldPrecheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ShieldPrecheckResponse) String() string {
	return dara.Prettify(s)
}

func (s ShieldPrecheckResponse) GoString() string {
	return s.String()
}

func (s *ShieldPrecheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ShieldPrecheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ShieldPrecheckResponse) GetBody() *ShieldPrecheckResponseBody {
	return s.Body
}

func (s *ShieldPrecheckResponse) SetHeaders(v map[string]*string) *ShieldPrecheckResponse {
	s.Headers = v
	return s
}

func (s *ShieldPrecheckResponse) SetStatusCode(v int32) *ShieldPrecheckResponse {
	s.StatusCode = &v
	return s
}

func (s *ShieldPrecheckResponse) SetBody(v *ShieldPrecheckResponseBody) *ShieldPrecheckResponse {
	s.Body = v
	return s
}

func (s *ShieldPrecheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
