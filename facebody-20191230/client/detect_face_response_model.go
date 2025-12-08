// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectFaceResponse
	GetStatusCode() *int32
	SetBody(v *DetectFaceResponseBody) *DetectFaceResponse
	GetBody() *DetectFaceResponseBody
}

type DetectFaceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectFaceResponse) GoString() string {
	return s.String()
}

func (s *DetectFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectFaceResponse) GetBody() *DetectFaceResponseBody {
	return s.Body
}

func (s *DetectFaceResponse) SetHeaders(v map[string]*string) *DetectFaceResponse {
	s.Headers = v
	return s
}

func (s *DetectFaceResponse) SetStatusCode(v int32) *DetectFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectFaceResponse) SetBody(v *DetectFaceResponseBody) *DetectFaceResponse {
	s.Body = v
	return s
}

func (s *DetectFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
