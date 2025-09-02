// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterConfigsResponseBody) *ListClusterConfigsResponse
	GetBody() *ListClusterConfigsResponseBody
}

type ListClusterConfigsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListClusterConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterConfigsResponse) GetBody() *ListClusterConfigsResponseBody {
	return s.Body
}

func (s *ListClusterConfigsResponse) SetHeaders(v map[string]*string) *ListClusterConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListClusterConfigsResponse) SetStatusCode(v int32) *ListClusterConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterConfigsResponse) SetBody(v *ListClusterConfigsResponseBody) *ListClusterConfigsResponse {
	s.Body = v
	return s
}

func (s *ListClusterConfigsResponse) Validate() error {
	return dara.Validate(s)
}
