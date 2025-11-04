// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecognitionEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRecognitionEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRecognitionEntityResponse
	GetStatusCode() *int32
	SetBody(v *CreateRecognitionEntityResponseBody) *CreateRecognitionEntityResponse
	GetBody() *CreateRecognitionEntityResponseBody
}

type CreateRecognitionEntityResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecognitionEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecognitionEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRecognitionEntityResponse) GoString() string {
	return s.String()
}

func (s *CreateRecognitionEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRecognitionEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRecognitionEntityResponse) GetBody() *CreateRecognitionEntityResponseBody {
	return s.Body
}

func (s *CreateRecognitionEntityResponse) SetHeaders(v map[string]*string) *CreateRecognitionEntityResponse {
	s.Headers = v
	return s
}

func (s *CreateRecognitionEntityResponse) SetStatusCode(v int32) *CreateRecognitionEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecognitionEntityResponse) SetBody(v *CreateRecognitionEntityResponseBody) *CreateRecognitionEntityResponse {
	s.Body = v
	return s
}

func (s *CreateRecognitionEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
