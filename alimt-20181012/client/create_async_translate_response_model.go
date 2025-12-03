// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAsyncTranslateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAsyncTranslateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAsyncTranslateResponse
	GetStatusCode() *int32
	SetBody(v *CreateAsyncTranslateResponseBody) *CreateAsyncTranslateResponse
	GetBody() *CreateAsyncTranslateResponseBody
}

type CreateAsyncTranslateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAsyncTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAsyncTranslateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAsyncTranslateResponse) GoString() string {
	return s.String()
}

func (s *CreateAsyncTranslateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAsyncTranslateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAsyncTranslateResponse) GetBody() *CreateAsyncTranslateResponseBody {
	return s.Body
}

func (s *CreateAsyncTranslateResponse) SetHeaders(v map[string]*string) *CreateAsyncTranslateResponse {
	s.Headers = v
	return s
}

func (s *CreateAsyncTranslateResponse) SetStatusCode(v int32) *CreateAsyncTranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAsyncTranslateResponse) SetBody(v *CreateAsyncTranslateResponseBody) *CreateAsyncTranslateResponse {
	s.Body = v
	return s
}

func (s *CreateAsyncTranslateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
