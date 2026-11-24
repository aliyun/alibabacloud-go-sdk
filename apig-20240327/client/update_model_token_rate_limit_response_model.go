// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelTokenRateLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelTokenRateLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelTokenRateLimitResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModelTokenRateLimitResponseBody) *UpdateModelTokenRateLimitResponse
	GetBody() *UpdateModelTokenRateLimitResponseBody
}

type UpdateModelTokenRateLimitResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelTokenRateLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelTokenRateLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelTokenRateLimitResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelTokenRateLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelTokenRateLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelTokenRateLimitResponse) GetBody() *UpdateModelTokenRateLimitResponseBody {
	return s.Body
}

func (s *UpdateModelTokenRateLimitResponse) SetHeaders(v map[string]*string) *UpdateModelTokenRateLimitResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelTokenRateLimitResponse) SetStatusCode(v int32) *UpdateModelTokenRateLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelTokenRateLimitResponse) SetBody(v *UpdateModelTokenRateLimitResponseBody) *UpdateModelTokenRateLimitResponse {
	s.Body = v
	return s
}

func (s *UpdateModelTokenRateLimitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
