// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeFaceResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeFaceResponseBody) *RecognizeFaceResponse
	GetBody() *RecognizeFaceResponseBody
}

type RecognizeFaceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeFaceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeFaceResponse) GetBody() *RecognizeFaceResponseBody {
	return s.Body
}

func (s *RecognizeFaceResponse) SetHeaders(v map[string]*string) *RecognizeFaceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeFaceResponse) SetStatusCode(v int32) *RecognizeFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeFaceResponse) SetBody(v *RecognizeFaceResponseBody) *RecognizeFaceResponse {
	s.Body = v
	return s
}

func (s *RecognizeFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
