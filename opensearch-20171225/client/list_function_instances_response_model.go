// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFunctionInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFunctionInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListFunctionInstancesResponseBody) *ListFunctionInstancesResponse
	GetBody() *ListFunctionInstancesResponseBody
}

type ListFunctionInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFunctionInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFunctionInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFunctionInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFunctionInstancesResponse) GetBody() *ListFunctionInstancesResponseBody {
	return s.Body
}

func (s *ListFunctionInstancesResponse) SetHeaders(v map[string]*string) *ListFunctionInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionInstancesResponse) SetStatusCode(v int32) *ListFunctionInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionInstancesResponse) SetBody(v *ListFunctionInstancesResponseBody) *ListFunctionInstancesResponse {
	s.Body = v
	return s
}

func (s *ListFunctionInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
