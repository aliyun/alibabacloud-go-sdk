// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyResponse
	GetStatusCode() *int32
	SetBody(v *VerifyResponseBody) *VerifyResponse
	GetBody() *VerifyResponseBody
}

type VerifyResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyResponse) GoString() string {
	return s.String()
}

func (s *VerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyResponse) GetBody() *VerifyResponseBody {
	return s.Body
}

func (s *VerifyResponse) SetHeaders(v map[string]*string) *VerifyResponse {
	s.Headers = v
	return s
}

func (s *VerifyResponse) SetStatusCode(v int32) *VerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyResponse) SetBody(v *VerifyResponseBody) *VerifyResponse {
	s.Body = v
	return s
}

func (s *VerifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
