// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantHonorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantHonorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantHonorResponse
	GetStatusCode() *int32
	SetBody(v *GrantHonorResponseBody) *GrantHonorResponse
	GetBody() *GrantHonorResponseBody
}

type GrantHonorResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantHonorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantHonorResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantHonorResponse) GoString() string {
	return s.String()
}

func (s *GrantHonorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantHonorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantHonorResponse) GetBody() *GrantHonorResponseBody {
	return s.Body
}

func (s *GrantHonorResponse) SetHeaders(v map[string]*string) *GrantHonorResponse {
	s.Headers = v
	return s
}

func (s *GrantHonorResponse) SetStatusCode(v int32) *GrantHonorResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantHonorResponse) SetBody(v *GrantHonorResponseBody) *GrantHonorResponse {
	s.Body = v
	return s
}

func (s *GrantHonorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
