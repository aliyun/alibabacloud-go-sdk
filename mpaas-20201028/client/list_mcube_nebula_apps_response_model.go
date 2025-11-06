// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeNebulaAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcubeNebulaAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcubeNebulaAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListMcubeNebulaAppsResponseBody) *ListMcubeNebulaAppsResponse
	GetBody() *ListMcubeNebulaAppsResponseBody
}

type ListMcubeNebulaAppsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcubeNebulaAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcubeNebulaAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeNebulaAppsResponse) GoString() string {
	return s.String()
}

func (s *ListMcubeNebulaAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcubeNebulaAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcubeNebulaAppsResponse) GetBody() *ListMcubeNebulaAppsResponseBody {
	return s.Body
}

func (s *ListMcubeNebulaAppsResponse) SetHeaders(v map[string]*string) *ListMcubeNebulaAppsResponse {
	s.Headers = v
	return s
}

func (s *ListMcubeNebulaAppsResponse) SetStatusCode(v int32) *ListMcubeNebulaAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcubeNebulaAppsResponse) SetBody(v *ListMcubeNebulaAppsResponseBody) *ListMcubeNebulaAppsResponse {
	s.Body = v
	return s
}

func (s *ListMcubeNebulaAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
