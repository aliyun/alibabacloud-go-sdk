// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindVerificationResponse
	GetStatusCode() *int32
	SetBody(v *UnbindVerificationResponseBody) *UnbindVerificationResponse
	GetBody() *UnbindVerificationResponseBody
}

type UnbindVerificationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindVerificationResponse) GoString() string {
	return s.String()
}

func (s *UnbindVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindVerificationResponse) GetBody() *UnbindVerificationResponseBody {
	return s.Body
}

func (s *UnbindVerificationResponse) SetHeaders(v map[string]*string) *UnbindVerificationResponse {
	s.Headers = v
	return s
}

func (s *UnbindVerificationResponse) SetStatusCode(v int32) *UnbindVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindVerificationResponse) SetBody(v *UnbindVerificationResponseBody) *UnbindVerificationResponse {
	s.Body = v
	return s
}

func (s *UnbindVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
