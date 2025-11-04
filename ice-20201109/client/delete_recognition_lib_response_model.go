// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecognitionLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRecognitionLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRecognitionLibResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRecognitionLibResponseBody) *DeleteRecognitionLibResponse
	GetBody() *DeleteRecognitionLibResponseBody
}

type DeleteRecognitionLibResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecognitionLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecognitionLibResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecognitionLibResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecognitionLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRecognitionLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRecognitionLibResponse) GetBody() *DeleteRecognitionLibResponseBody {
	return s.Body
}

func (s *DeleteRecognitionLibResponse) SetHeaders(v map[string]*string) *DeleteRecognitionLibResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecognitionLibResponse) SetStatusCode(v int32) *DeleteRecognitionLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecognitionLibResponse) SetBody(v *DeleteRecognitionLibResponseBody) *DeleteRecognitionLibResponse {
	s.Body = v
	return s
}

func (s *DeleteRecognitionLibResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
