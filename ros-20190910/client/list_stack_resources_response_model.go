// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStackResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStackResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListStackResourcesResponseBody) *ListStackResourcesResponse
	GetBody() *ListStackResourcesResponseBody
}

type ListStackResourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStackResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStackResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStackResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListStackResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStackResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStackResourcesResponse) GetBody() *ListStackResourcesResponseBody {
	return s.Body
}

func (s *ListStackResourcesResponse) SetHeaders(v map[string]*string) *ListStackResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListStackResourcesResponse) SetStatusCode(v int32) *ListStackResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackResourcesResponse) SetBody(v *ListStackResourcesResponseBody) *ListStackResourcesResponse {
	s.Body = v
	return s
}

func (s *ListStackResourcesResponse) Validate() error {
	return dara.Validate(s)
}
