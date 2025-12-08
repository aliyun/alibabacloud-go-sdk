// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectInfraredLivingFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectInfraredLivingFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectInfraredLivingFaceResponse
	GetStatusCode() *int32
	SetBody(v *DetectInfraredLivingFaceResponseBody) *DetectInfraredLivingFaceResponse
	GetBody() *DetectInfraredLivingFaceResponseBody
}

type DetectInfraredLivingFaceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectInfraredLivingFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectInfraredLivingFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectInfraredLivingFaceResponse) GoString() string {
	return s.String()
}

func (s *DetectInfraredLivingFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectInfraredLivingFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectInfraredLivingFaceResponse) GetBody() *DetectInfraredLivingFaceResponseBody {
	return s.Body
}

func (s *DetectInfraredLivingFaceResponse) SetHeaders(v map[string]*string) *DetectInfraredLivingFaceResponse {
	s.Headers = v
	return s
}

func (s *DetectInfraredLivingFaceResponse) SetStatusCode(v int32) *DetectInfraredLivingFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectInfraredLivingFaceResponse) SetBody(v *DetectInfraredLivingFaceResponseBody) *DetectInfraredLivingFaceResponse {
	s.Body = v
	return s
}

func (s *DetectInfraredLivingFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
