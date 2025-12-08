// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizePublicFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizePublicFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizePublicFaceResponse
	GetStatusCode() *int32
	SetBody(v *RecognizePublicFaceResponseBody) *RecognizePublicFaceResponse
	GetBody() *RecognizePublicFaceResponseBody
}

type RecognizePublicFaceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizePublicFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizePublicFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizePublicFaceResponse) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizePublicFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizePublicFaceResponse) GetBody() *RecognizePublicFaceResponseBody {
	return s.Body
}

func (s *RecognizePublicFaceResponse) SetHeaders(v map[string]*string) *RecognizePublicFaceResponse {
	s.Headers = v
	return s
}

func (s *RecognizePublicFaceResponse) SetStatusCode(v int32) *RecognizePublicFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizePublicFaceResponse) SetBody(v *RecognizePublicFaceResponseBody) *RecognizePublicFaceResponse {
	s.Body = v
	return s
}

func (s *RecognizePublicFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
