// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFunctionResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFunctionResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListFunctionResourcesResponseBody) *ListFunctionResourcesResponse
	GetBody() *ListFunctionResourcesResponseBody
}

type ListFunctionResourcesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFunctionResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFunctionResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFunctionResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFunctionResourcesResponse) GetBody() *ListFunctionResourcesResponseBody {
	return s.Body
}

func (s *ListFunctionResourcesResponse) SetHeaders(v map[string]*string) *ListFunctionResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionResourcesResponse) SetStatusCode(v int32) *ListFunctionResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionResourcesResponse) SetBody(v *ListFunctionResourcesResponseBody) *ListFunctionResourcesResponse {
	s.Body = v
	return s
}

func (s *ListFunctionResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
