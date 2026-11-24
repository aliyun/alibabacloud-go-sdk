// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelTokenRateLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelTokenRateLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelTokenRateLimitResponse
	GetStatusCode() *int32
	SetBody(v *GetModelTokenRateLimitResponseBody) *GetModelTokenRateLimitResponse
	GetBody() *GetModelTokenRateLimitResponseBody
}

type GetModelTokenRateLimitResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelTokenRateLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelTokenRateLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelTokenRateLimitResponse) GoString() string {
	return s.String()
}

func (s *GetModelTokenRateLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelTokenRateLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelTokenRateLimitResponse) GetBody() *GetModelTokenRateLimitResponseBody {
	return s.Body
}

func (s *GetModelTokenRateLimitResponse) SetHeaders(v map[string]*string) *GetModelTokenRateLimitResponse {
	s.Headers = v
	return s
}

func (s *GetModelTokenRateLimitResponse) SetStatusCode(v int32) *GetModelTokenRateLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelTokenRateLimitResponse) SetBody(v *GetModelTokenRateLimitResponseBody) *GetModelTokenRateLimitResponse {
	s.Body = v
	return s
}

func (s *GetModelTokenRateLimitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
