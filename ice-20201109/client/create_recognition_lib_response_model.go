// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecognitionLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRecognitionLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRecognitionLibResponse
	GetStatusCode() *int32
	SetBody(v *CreateRecognitionLibResponseBody) *CreateRecognitionLibResponse
	GetBody() *CreateRecognitionLibResponseBody
}

type CreateRecognitionLibResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecognitionLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecognitionLibResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRecognitionLibResponse) GoString() string {
	return s.String()
}

func (s *CreateRecognitionLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRecognitionLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRecognitionLibResponse) GetBody() *CreateRecognitionLibResponseBody {
	return s.Body
}

func (s *CreateRecognitionLibResponse) SetHeaders(v map[string]*string) *CreateRecognitionLibResponse {
	s.Headers = v
	return s
}

func (s *CreateRecognitionLibResponse) SetStatusCode(v int32) *CreateRecognitionLibResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecognitionLibResponse) SetBody(v *CreateRecognitionLibResponseBody) *CreateRecognitionLibResponse {
	s.Body = v
	return s
}

func (s *CreateRecognitionLibResponse) Validate() error {
	return dara.Validate(s)
}
