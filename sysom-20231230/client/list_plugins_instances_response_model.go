// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginsInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPluginsInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPluginsInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListPluginsInstancesResponseBody) *ListPluginsInstancesResponse
	GetBody() *ListPluginsInstancesResponseBody
}

type ListPluginsInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPluginsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPluginsInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListPluginsInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPluginsInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPluginsInstancesResponse) GetBody() *ListPluginsInstancesResponseBody {
	return s.Body
}

func (s *ListPluginsInstancesResponse) SetHeaders(v map[string]*string) *ListPluginsInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListPluginsInstancesResponse) SetStatusCode(v int32) *ListPluginsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPluginsInstancesResponse) SetBody(v *ListPluginsInstancesResponseBody) *ListPluginsInstancesResponse {
	s.Body = v
	return s
}

func (s *ListPluginsInstancesResponse) Validate() error {
	return dara.Validate(s)
}
