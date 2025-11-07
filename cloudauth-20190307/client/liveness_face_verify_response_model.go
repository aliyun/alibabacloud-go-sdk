// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLivenessFaceVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LivenessFaceVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LivenessFaceVerifyResponse
	GetStatusCode() *int32
	SetBody(v *LivenessFaceVerifyResponseBody) *LivenessFaceVerifyResponse
	GetBody() *LivenessFaceVerifyResponseBody
}

type LivenessFaceVerifyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LivenessFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LivenessFaceVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s LivenessFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LivenessFaceVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LivenessFaceVerifyResponse) GetBody() *LivenessFaceVerifyResponseBody {
	return s.Body
}

func (s *LivenessFaceVerifyResponse) SetHeaders(v map[string]*string) *LivenessFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *LivenessFaceVerifyResponse) SetStatusCode(v int32) *LivenessFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *LivenessFaceVerifyResponse) SetBody(v *LivenessFaceVerifyResponseBody) *LivenessFaceVerifyResponse {
	s.Body = v
	return s
}

func (s *LivenessFaceVerifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
