// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeOrderResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeOrderResponseBody) *RecognizeOrderResponse
	GetBody() *RecognizeOrderResponseBody
}

type RecognizeOrderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeOrderResponse) GoString() string {
	return s.String()
}

func (s *RecognizeOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeOrderResponse) GetBody() *RecognizeOrderResponseBody {
	return s.Body
}

func (s *RecognizeOrderResponse) SetHeaders(v map[string]*string) *RecognizeOrderResponse {
	s.Headers = v
	return s
}

func (s *RecognizeOrderResponse) SetStatusCode(v int32) *RecognizeOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeOrderResponse) SetBody(v *RecognizeOrderResponseBody) *RecognizeOrderResponse {
	s.Body = v
	return s
}

func (s *RecognizeOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
