// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecognitionSampleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRecognitionSampleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRecognitionSampleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRecognitionSampleResponseBody) *DeleteRecognitionSampleResponse
	GetBody() *DeleteRecognitionSampleResponseBody
}

type DeleteRecognitionSampleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecognitionSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecognitionSampleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecognitionSampleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecognitionSampleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRecognitionSampleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRecognitionSampleResponse) GetBody() *DeleteRecognitionSampleResponseBody {
	return s.Body
}

func (s *DeleteRecognitionSampleResponse) SetHeaders(v map[string]*string) *DeleteRecognitionSampleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecognitionSampleResponse) SetStatusCode(v int32) *DeleteRecognitionSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecognitionSampleResponse) SetBody(v *DeleteRecognitionSampleResponseBody) *DeleteRecognitionSampleResponse {
	s.Body = v
	return s
}

func (s *DeleteRecognitionSampleResponse) Validate() error {
	return dara.Validate(s)
}
