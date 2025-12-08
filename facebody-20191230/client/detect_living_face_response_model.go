// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectLivingFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectLivingFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectLivingFaceResponse
	GetStatusCode() *int32
	SetBody(v *DetectLivingFaceResponseBody) *DetectLivingFaceResponse
	GetBody() *DetectLivingFaceResponseBody
}

type DetectLivingFaceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectLivingFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectLivingFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectLivingFaceResponse) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectLivingFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectLivingFaceResponse) GetBody() *DetectLivingFaceResponseBody {
	return s.Body
}

func (s *DetectLivingFaceResponse) SetHeaders(v map[string]*string) *DetectLivingFaceResponse {
	s.Headers = v
	return s
}

func (s *DetectLivingFaceResponse) SetStatusCode(v int32) *DetectLivingFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectLivingFaceResponse) SetBody(v *DetectLivingFaceResponseBody) *DetectLivingFaceResponse {
	s.Body = v
	return s
}

func (s *DetectLivingFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
