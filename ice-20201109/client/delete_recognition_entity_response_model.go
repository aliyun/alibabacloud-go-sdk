// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecognitionEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRecognitionEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRecognitionEntityResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRecognitionEntityResponseBody) *DeleteRecognitionEntityResponse
	GetBody() *DeleteRecognitionEntityResponseBody
}

type DeleteRecognitionEntityResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecognitionEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecognitionEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecognitionEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecognitionEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRecognitionEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRecognitionEntityResponse) GetBody() *DeleteRecognitionEntityResponseBody {
	return s.Body
}

func (s *DeleteRecognitionEntityResponse) SetHeaders(v map[string]*string) *DeleteRecognitionEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecognitionEntityResponse) SetStatusCode(v int32) *DeleteRecognitionEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecognitionEntityResponse) SetBody(v *DeleteRecognitionEntityResponseBody) *DeleteRecognitionEntityResponse {
	s.Body = v
	return s
}

func (s *DeleteRecognitionEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
