// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeActionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeActionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeActionResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeActionResponseBody) *RecognizeActionResponse
	GetBody() *RecognizeActionResponseBody
}

type RecognizeActionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeActionResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeActionResponse) GoString() string {
	return s.String()
}

func (s *RecognizeActionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeActionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeActionResponse) GetBody() *RecognizeActionResponseBody {
	return s.Body
}

func (s *RecognizeActionResponse) SetHeaders(v map[string]*string) *RecognizeActionResponse {
	s.Headers = v
	return s
}

func (s *RecognizeActionResponse) SetStatusCode(v int32) *RecognizeActionResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeActionResponse) SetBody(v *RecognizeActionResponseBody) *RecognizeActionResponse {
	s.Body = v
	return s
}

func (s *RecognizeActionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
