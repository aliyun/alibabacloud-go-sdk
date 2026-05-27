// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestLargeModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestLargeModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestLargeModelResponse
	GetStatusCode() *int32
	SetBody(v *TestLargeModelResponseBody) *TestLargeModelResponse
	GetBody() *TestLargeModelResponseBody
}

type TestLargeModelResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestLargeModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestLargeModelResponse) String() string {
	return dara.Prettify(s)
}

func (s TestLargeModelResponse) GoString() string {
	return s.String()
}

func (s *TestLargeModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestLargeModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestLargeModelResponse) GetBody() *TestLargeModelResponseBody {
	return s.Body
}

func (s *TestLargeModelResponse) SetHeaders(v map[string]*string) *TestLargeModelResponse {
	s.Headers = v
	return s
}

func (s *TestLargeModelResponse) SetStatusCode(v int32) *TestLargeModelResponse {
	s.StatusCode = &v
	return s
}

func (s *TestLargeModelResponse) SetBody(v *TestLargeModelResponseBody) *TestLargeModelResponse {
	s.Body = v
	return s
}

func (s *TestLargeModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
