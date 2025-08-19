// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitFaceVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitFaceVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitFaceVerifyResponse
	GetStatusCode() *int32
	SetBody(v *InitFaceVerifyResponseBody) *InitFaceVerifyResponse
	GetBody() *InitFaceVerifyResponseBody
}

type InitFaceVerifyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitFaceVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s InitFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitFaceVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitFaceVerifyResponse) GetBody() *InitFaceVerifyResponseBody {
	return s.Body
}

func (s *InitFaceVerifyResponse) SetHeaders(v map[string]*string) *InitFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *InitFaceVerifyResponse) SetStatusCode(v int32) *InitFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *InitFaceVerifyResponse) SetBody(v *InitFaceVerifyResponseBody) *InitFaceVerifyResponse {
	s.Body = v
	return s
}

func (s *InitFaceVerifyResponse) Validate() error {
	return dara.Validate(s)
}
