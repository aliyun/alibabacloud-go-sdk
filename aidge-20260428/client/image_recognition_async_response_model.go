// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRecognitionAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageRecognitionAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageRecognitionAsyncResponse
	GetStatusCode() *int32
	SetBody(v *ImageRecognitionAsyncResponseBody) *ImageRecognitionAsyncResponse
	GetBody() *ImageRecognitionAsyncResponseBody
}

type ImageRecognitionAsyncResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageRecognitionAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageRecognitionAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageRecognitionAsyncResponse) GoString() string {
	return s.String()
}

func (s *ImageRecognitionAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageRecognitionAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageRecognitionAsyncResponse) GetBody() *ImageRecognitionAsyncResponseBody {
	return s.Body
}

func (s *ImageRecognitionAsyncResponse) SetHeaders(v map[string]*string) *ImageRecognitionAsyncResponse {
	s.Headers = v
	return s
}

func (s *ImageRecognitionAsyncResponse) SetStatusCode(v int32) *ImageRecognitionAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageRecognitionAsyncResponse) SetBody(v *ImageRecognitionAsyncResponseBody) *ImageRecognitionAsyncResponse {
	s.Body = v
	return s
}

func (s *ImageRecognitionAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
