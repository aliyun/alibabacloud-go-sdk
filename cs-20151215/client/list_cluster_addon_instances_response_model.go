// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterAddonInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterAddonInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterAddonInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterAddonInstancesResponseBody) *ListClusterAddonInstancesResponse
	GetBody() *ListClusterAddonInstancesResponseBody
}

type ListClusterAddonInstancesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterAddonInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterAddonInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterAddonInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterAddonInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterAddonInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterAddonInstancesResponse) GetBody() *ListClusterAddonInstancesResponseBody {
	return s.Body
}

func (s *ListClusterAddonInstancesResponse) SetHeaders(v map[string]*string) *ListClusterAddonInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListClusterAddonInstancesResponse) SetStatusCode(v int32) *ListClusterAddonInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterAddonInstancesResponse) SetBody(v *ListClusterAddonInstancesResponseBody) *ListClusterAddonInstancesResponse {
	s.Body = v
	return s
}

func (s *ListClusterAddonInstancesResponse) Validate() error {
	return dara.Validate(s)
}
