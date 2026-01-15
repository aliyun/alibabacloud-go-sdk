// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestModelResponse
	GetStatusCode() *int32
	SetBody(v *TestModelResponseBody) *TestModelResponse
	GetBody() *TestModelResponseBody
}

type TestModelResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestModelResponse) String() string {
	return dara.Prettify(s)
}

func (s TestModelResponse) GoString() string {
	return s.String()
}

func (s *TestModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestModelResponse) GetBody() *TestModelResponseBody {
	return s.Body
}

func (s *TestModelResponse) SetHeaders(v map[string]*string) *TestModelResponse {
	s.Headers = v
	return s
}

func (s *TestModelResponse) SetStatusCode(v int32) *TestModelResponse {
	s.StatusCode = &v
	return s
}

func (s *TestModelResponse) SetBody(v *TestModelResponseBody) *TestModelResponse {
	s.Body = v
	return s
}

func (s *TestModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
