// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionMetasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFunctionMetasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFunctionMetasResponse
	GetStatusCode() *int32
	SetBody(v *ListFunctionMetasResponseBody) *ListFunctionMetasResponse
	GetBody() *ListFunctionMetasResponseBody
}

type ListFunctionMetasResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFunctionMetasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFunctionMetasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionMetasResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionMetasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFunctionMetasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFunctionMetasResponse) GetBody() *ListFunctionMetasResponseBody {
	return s.Body
}

func (s *ListFunctionMetasResponse) SetHeaders(v map[string]*string) *ListFunctionMetasResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionMetasResponse) SetStatusCode(v int32) *ListFunctionMetasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionMetasResponse) SetBody(v *ListFunctionMetasResponseBody) *ListFunctionMetasResponse {
	s.Body = v
	return s
}

func (s *ListFunctionMetasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
