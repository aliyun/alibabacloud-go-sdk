// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeVINCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeVINCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeVINCodeResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeVINCodeResponseBody) *RecognizeVINCodeResponse
	GetBody() *RecognizeVINCodeResponseBody
}

type RecognizeVINCodeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeVINCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeVINCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVINCodeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVINCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeVINCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeVINCodeResponse) GetBody() *RecognizeVINCodeResponseBody {
	return s.Body
}

func (s *RecognizeVINCodeResponse) SetHeaders(v map[string]*string) *RecognizeVINCodeResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVINCodeResponse) SetStatusCode(v int32) *RecognizeVINCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVINCodeResponse) SetBody(v *RecognizeVINCodeResponseBody) *RecognizeVINCodeResponse {
	s.Body = v
	return s
}

func (s *RecognizeVINCodeResponse) Validate() error {
	return dara.Validate(s)
}
