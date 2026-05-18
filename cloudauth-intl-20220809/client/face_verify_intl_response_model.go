// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceVerifyIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FaceVerifyIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FaceVerifyIntlResponse
	GetStatusCode() *int32
	SetBody(v *FaceVerifyIntlResponseBody) *FaceVerifyIntlResponse
	GetBody() *FaceVerifyIntlResponseBody
}

type FaceVerifyIntlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceVerifyIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceVerifyIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s FaceVerifyIntlResponse) GoString() string {
	return s.String()
}

func (s *FaceVerifyIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FaceVerifyIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FaceVerifyIntlResponse) GetBody() *FaceVerifyIntlResponseBody {
	return s.Body
}

func (s *FaceVerifyIntlResponse) SetHeaders(v map[string]*string) *FaceVerifyIntlResponse {
	s.Headers = v
	return s
}

func (s *FaceVerifyIntlResponse) SetStatusCode(v int32) *FaceVerifyIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceVerifyIntlResponse) SetBody(v *FaceVerifyIntlResponseBody) *FaceVerifyIntlResponse {
	s.Body = v
	return s
}

func (s *FaceVerifyIntlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
