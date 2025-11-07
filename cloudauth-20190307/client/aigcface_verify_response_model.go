// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAIGCFaceVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AIGCFaceVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AIGCFaceVerifyResponse
	GetStatusCode() *int32
	SetBody(v *AIGCFaceVerifyResponseBody) *AIGCFaceVerifyResponse
	GetBody() *AIGCFaceVerifyResponseBody
}

type AIGCFaceVerifyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AIGCFaceVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AIGCFaceVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s AIGCFaceVerifyResponse) GoString() string {
	return s.String()
}

func (s *AIGCFaceVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AIGCFaceVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AIGCFaceVerifyResponse) GetBody() *AIGCFaceVerifyResponseBody {
	return s.Body
}

func (s *AIGCFaceVerifyResponse) SetHeaders(v map[string]*string) *AIGCFaceVerifyResponse {
	s.Headers = v
	return s
}

func (s *AIGCFaceVerifyResponse) SetStatusCode(v int32) *AIGCFaceVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *AIGCFaceVerifyResponse) SetBody(v *AIGCFaceVerifyResponseBody) *AIGCFaceVerifyResponse {
	s.Body = v
	return s
}

func (s *AIGCFaceVerifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
