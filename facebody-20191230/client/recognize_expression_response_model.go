// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeExpressionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeExpressionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeExpressionResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeExpressionResponseBody) *RecognizeExpressionResponse
	GetBody() *RecognizeExpressionResponseBody
}

type RecognizeExpressionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeExpressionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeExpressionResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeExpressionResponse) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeExpressionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeExpressionResponse) GetBody() *RecognizeExpressionResponseBody {
	return s.Body
}

func (s *RecognizeExpressionResponse) SetHeaders(v map[string]*string) *RecognizeExpressionResponse {
	s.Headers = v
	return s
}

func (s *RecognizeExpressionResponse) SetStatusCode(v int32) *RecognizeExpressionResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeExpressionResponse) SetBody(v *RecognizeExpressionResponseBody) *RecognizeExpressionResponse {
	s.Body = v
	return s
}

func (s *RecognizeExpressionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
