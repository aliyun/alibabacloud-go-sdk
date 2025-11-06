// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeNebulaResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcubeNebulaResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcubeNebulaResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListMcubeNebulaResourcesResponseBody) *ListMcubeNebulaResourcesResponse
	GetBody() *ListMcubeNebulaResourcesResponseBody
}

type ListMcubeNebulaResourcesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcubeNebulaResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcubeNebulaResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeNebulaResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListMcubeNebulaResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcubeNebulaResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcubeNebulaResourcesResponse) GetBody() *ListMcubeNebulaResourcesResponseBody {
	return s.Body
}

func (s *ListMcubeNebulaResourcesResponse) SetHeaders(v map[string]*string) *ListMcubeNebulaResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListMcubeNebulaResourcesResponse) SetStatusCode(v int32) *ListMcubeNebulaResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponse) SetBody(v *ListMcubeNebulaResourcesResponseBody) *ListMcubeNebulaResourcesResponse {
	s.Body = v
	return s
}

func (s *ListMcubeNebulaResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
