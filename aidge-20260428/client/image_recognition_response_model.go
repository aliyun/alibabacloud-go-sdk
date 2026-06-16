// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRecognitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageRecognitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageRecognitionResponse
	GetStatusCode() *int32
	SetBody(v *ImageRecognitionResponseBody) *ImageRecognitionResponse
	GetBody() *ImageRecognitionResponseBody
}

type ImageRecognitionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageRecognitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageRecognitionResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageRecognitionResponse) GoString() string {
	return s.String()
}

func (s *ImageRecognitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageRecognitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageRecognitionResponse) GetBody() *ImageRecognitionResponseBody {
	return s.Body
}

func (s *ImageRecognitionResponse) SetHeaders(v map[string]*string) *ImageRecognitionResponse {
	s.Headers = v
	return s
}

func (s *ImageRecognitionResponse) SetStatusCode(v int32) *ImageRecognitionResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageRecognitionResponse) SetBody(v *ImageRecognitionResponseBody) *ImageRecognitionResponse {
	s.Body = v
	return s
}

func (s *ImageRecognitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
