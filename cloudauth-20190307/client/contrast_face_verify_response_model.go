// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContrastFaceVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ContrastFaceVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ContrastFaceVerifyResponse
	GetStatusCode() *int32
	SetBody(v *ContrastFaceVerifyResponseBody) *ContrastFaceVerifyResponse
	GetBody() *ContrastFaceVerifyResponseBody
}

type ContrastFaceVerifyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContrastFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContrastFaceVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s ContrastFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ContrastFaceVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ContrastFaceVerifyResponse) GetBody() *ContrastFaceVerifyResponseBody {
	return s.Body
}

func (s *ContrastFaceVerifyResponse) SetHeaders(v map[string]*string) *ContrastFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *ContrastFaceVerifyResponse) SetStatusCode(v int32) *ContrastFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *ContrastFaceVerifyResponse) SetBody(v *ContrastFaceVerifyResponseBody) *ContrastFaceVerifyResponse {
	s.Body = v
	return s
}

func (s *ContrastFaceVerifyResponse) Validate() error {
	return dara.Validate(s)
}
