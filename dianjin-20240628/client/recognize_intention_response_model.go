// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeIntentionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeIntentionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeIntentionResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeIntentionResponseBody) *RecognizeIntentionResponse
	GetBody() *RecognizeIntentionResponseBody
}

type RecognizeIntentionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeIntentionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeIntentionResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIntentionResponse) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeIntentionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeIntentionResponse) GetBody() *RecognizeIntentionResponseBody {
	return s.Body
}

func (s *RecognizeIntentionResponse) SetHeaders(v map[string]*string) *RecognizeIntentionResponse {
	s.Headers = v
	return s
}

func (s *RecognizeIntentionResponse) SetStatusCode(v int32) *RecognizeIntentionResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeIntentionResponse) SetBody(v *RecognizeIntentionResponseBody) *RecognizeIntentionResponse {
	s.Body = v
	return s
}

func (s *RecognizeIntentionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
