// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestPreModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestPreModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestPreModelResponse
	GetStatusCode() *int32
	SetBody(v *TestPreModelResponseBody) *TestPreModelResponse
	GetBody() *TestPreModelResponseBody
}

type TestPreModelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestPreModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestPreModelResponse) String() string {
	return dara.Prettify(s)
}

func (s TestPreModelResponse) GoString() string {
	return s.String()
}

func (s *TestPreModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestPreModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestPreModelResponse) GetBody() *TestPreModelResponseBody {
	return s.Body
}

func (s *TestPreModelResponse) SetHeaders(v map[string]*string) *TestPreModelResponse {
	s.Headers = v
	return s
}

func (s *TestPreModelResponse) SetStatusCode(v int32) *TestPreModelResponse {
	s.StatusCode = &v
	return s
}

func (s *TestPreModelResponse) SetBody(v *TestPreModelResponseBody) *TestPreModelResponse {
	s.Body = v
	return s
}

func (s *TestPreModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
