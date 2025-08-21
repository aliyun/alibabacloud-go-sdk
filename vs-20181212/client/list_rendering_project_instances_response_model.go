// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingProjectInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRenderingProjectInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRenderingProjectInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListRenderingProjectInstancesResponseBody) *ListRenderingProjectInstancesResponse
	GetBody() *ListRenderingProjectInstancesResponseBody
}

type ListRenderingProjectInstancesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRenderingProjectInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRenderingProjectInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingProjectInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListRenderingProjectInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRenderingProjectInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRenderingProjectInstancesResponse) GetBody() *ListRenderingProjectInstancesResponseBody {
	return s.Body
}

func (s *ListRenderingProjectInstancesResponse) SetHeaders(v map[string]*string) *ListRenderingProjectInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListRenderingProjectInstancesResponse) SetStatusCode(v int32) *ListRenderingProjectInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRenderingProjectInstancesResponse) SetBody(v *ListRenderingProjectInstancesResponseBody) *ListRenderingProjectInstancesResponse {
	s.Body = v
	return s
}

func (s *ListRenderingProjectInstancesResponse) Validate() error {
	return dara.Validate(s)
}
