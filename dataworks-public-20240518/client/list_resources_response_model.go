// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListResourcesResponseBody) *ListResourcesResponse
	GetBody() *ListResourcesResponseBody
}

type ListResourcesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourcesResponse) GetBody() *ListResourcesResponseBody {
	return s.Body
}

func (s *ListResourcesResponse) SetHeaders(v map[string]*string) *ListResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesResponse) SetStatusCode(v int32) *ListResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesResponse) SetBody(v *ListResourcesResponseBody) *ListResourcesResponse {
	s.Body = v
	return s
}

func (s *ListResourcesResponse) Validate() error {
	return dara.Validate(s)
}
