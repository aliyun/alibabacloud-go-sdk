// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeUpgradeTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcubeUpgradeTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcubeUpgradeTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListMcubeUpgradeTasksResponseBody) *ListMcubeUpgradeTasksResponse
	GetBody() *ListMcubeUpgradeTasksResponseBody
}

type ListMcubeUpgradeTasksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcubeUpgradeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcubeUpgradeTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeUpgradeTasksResponse) GoString() string {
	return s.String()
}

func (s *ListMcubeUpgradeTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcubeUpgradeTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcubeUpgradeTasksResponse) GetBody() *ListMcubeUpgradeTasksResponseBody {
	return s.Body
}

func (s *ListMcubeUpgradeTasksResponse) SetHeaders(v map[string]*string) *ListMcubeUpgradeTasksResponse {
	s.Headers = v
	return s
}

func (s *ListMcubeUpgradeTasksResponse) SetStatusCode(v int32) *ListMcubeUpgradeTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponse) SetBody(v *ListMcubeUpgradeTasksResponseBody) *ListMcubeUpgradeTasksResponse {
	s.Body = v
	return s
}

func (s *ListMcubeUpgradeTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
