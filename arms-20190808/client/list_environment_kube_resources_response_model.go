// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentKubeResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnvironmentKubeResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnvironmentKubeResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListEnvironmentKubeResourcesResponseBody) *ListEnvironmentKubeResourcesResponse
	GetBody() *ListEnvironmentKubeResourcesResponseBody
}

type ListEnvironmentKubeResourcesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnvironmentKubeResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnvironmentKubeResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentKubeResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentKubeResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnvironmentKubeResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnvironmentKubeResourcesResponse) GetBody() *ListEnvironmentKubeResourcesResponseBody {
	return s.Body
}

func (s *ListEnvironmentKubeResourcesResponse) SetHeaders(v map[string]*string) *ListEnvironmentKubeResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentKubeResourcesResponse) SetStatusCode(v int32) *ListEnvironmentKubeResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvironmentKubeResourcesResponse) SetBody(v *ListEnvironmentKubeResourcesResponseBody) *ListEnvironmentKubeResourcesResponse {
	s.Body = v
	return s
}

func (s *ListEnvironmentKubeResourcesResponse) Validate() error {
	return dara.Validate(s)
}
