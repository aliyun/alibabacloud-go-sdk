// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitAuthVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitAuthVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitAuthVerifyResponse
	GetStatusCode() *int32
	SetBody(v *InitAuthVerifyResponseBody) *InitAuthVerifyResponse
	GetBody() *InitAuthVerifyResponseBody
}

type InitAuthVerifyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitAuthVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitAuthVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s InitAuthVerifyResponse) GoString() string {
	return s.String()
}

func (s *InitAuthVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitAuthVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitAuthVerifyResponse) GetBody() *InitAuthVerifyResponseBody {
	return s.Body
}

func (s *InitAuthVerifyResponse) SetHeaders(v map[string]*string) *InitAuthVerifyResponse {
	s.Headers = v
	return s
}

func (s *InitAuthVerifyResponse) SetStatusCode(v int32) *InitAuthVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *InitAuthVerifyResponse) SetBody(v *InitAuthVerifyResponseBody) *InitAuthVerifyResponse {
	s.Body = v
	return s
}

func (s *InitAuthVerifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
