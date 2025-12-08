// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVideoLivingFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectVideoLivingFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectVideoLivingFaceResponse
	GetStatusCode() *int32
	SetBody(v *DetectVideoLivingFaceResponseBody) *DetectVideoLivingFaceResponse
	GetBody() *DetectVideoLivingFaceResponseBody
}

type DetectVideoLivingFaceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectVideoLivingFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectVideoLivingFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoLivingFaceResponse) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectVideoLivingFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectVideoLivingFaceResponse) GetBody() *DetectVideoLivingFaceResponseBody {
	return s.Body
}

func (s *DetectVideoLivingFaceResponse) SetHeaders(v map[string]*string) *DetectVideoLivingFaceResponse {
	s.Headers = v
	return s
}

func (s *DetectVideoLivingFaceResponse) SetStatusCode(v int32) *DetectVideoLivingFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectVideoLivingFaceResponse) SetBody(v *DetectVideoLivingFaceResponseBody) *DetectVideoLivingFaceResponse {
	s.Body = v
	return s
}

func (s *DetectVideoLivingFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
