// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBlurFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BlurFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BlurFaceResponse
	GetStatusCode() *int32
	SetBody(v *BlurFaceResponseBody) *BlurFaceResponse
	GetBody() *BlurFaceResponseBody
}

type BlurFaceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BlurFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BlurFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s BlurFaceResponse) GoString() string {
	return s.String()
}

func (s *BlurFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BlurFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BlurFaceResponse) GetBody() *BlurFaceResponseBody {
	return s.Body
}

func (s *BlurFaceResponse) SetHeaders(v map[string]*string) *BlurFaceResponse {
	s.Headers = v
	return s
}

func (s *BlurFaceResponse) SetStatusCode(v int32) *BlurFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *BlurFaceResponse) SetBody(v *BlurFaceResponseBody) *BlurFaceResponse {
	s.Body = v
	return s
}

func (s *BlurFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
