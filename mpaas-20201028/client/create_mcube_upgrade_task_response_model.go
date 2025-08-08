// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeUpgradeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcubeUpgradeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcubeUpgradeTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcubeUpgradeTaskResponseBody) *CreateMcubeUpgradeTaskResponse
	GetBody() *CreateMcubeUpgradeTaskResponseBody
}

type CreateMcubeUpgradeTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcubeUpgradeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcubeUpgradeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeUpgradeTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateMcubeUpgradeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcubeUpgradeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcubeUpgradeTaskResponse) GetBody() *CreateMcubeUpgradeTaskResponseBody {
	return s.Body
}

func (s *CreateMcubeUpgradeTaskResponse) SetHeaders(v map[string]*string) *CreateMcubeUpgradeTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateMcubeUpgradeTaskResponse) SetStatusCode(v int32) *CreateMcubeUpgradeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcubeUpgradeTaskResponse) SetBody(v *CreateMcubeUpgradeTaskResponseBody) *CreateMcubeUpgradeTaskResponse {
	s.Body = v
	return s
}

func (s *CreateMcubeUpgradeTaskResponse) Validate() error {
	return dara.Validate(s)
}
