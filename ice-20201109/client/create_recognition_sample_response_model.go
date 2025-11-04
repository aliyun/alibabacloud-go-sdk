// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecognitionSampleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRecognitionSampleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRecognitionSampleResponse
	GetStatusCode() *int32
	SetBody(v *CreateRecognitionSampleResponseBody) *CreateRecognitionSampleResponse
	GetBody() *CreateRecognitionSampleResponseBody
}

type CreateRecognitionSampleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecognitionSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecognitionSampleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRecognitionSampleResponse) GoString() string {
	return s.String()
}

func (s *CreateRecognitionSampleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRecognitionSampleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRecognitionSampleResponse) GetBody() *CreateRecognitionSampleResponseBody {
	return s.Body
}

func (s *CreateRecognitionSampleResponse) SetHeaders(v map[string]*string) *CreateRecognitionSampleResponse {
	s.Headers = v
	return s
}

func (s *CreateRecognitionSampleResponse) SetStatusCode(v int32) *CreateRecognitionSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecognitionSampleResponse) SetBody(v *CreateRecognitionSampleResponseBody) *CreateRecognitionSampleResponse {
	s.Body = v
	return s
}

func (s *CreateRecognitionSampleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
