// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeUpgradePackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcubeUpgradePackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcubeUpgradePackageResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcubeUpgradePackageResponseBody) *CreateMcubeUpgradePackageResponse
	GetBody() *CreateMcubeUpgradePackageResponseBody
}

type CreateMcubeUpgradePackageResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcubeUpgradePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcubeUpgradePackageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeUpgradePackageResponse) GoString() string {
	return s.String()
}

func (s *CreateMcubeUpgradePackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcubeUpgradePackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcubeUpgradePackageResponse) GetBody() *CreateMcubeUpgradePackageResponseBody {
	return s.Body
}

func (s *CreateMcubeUpgradePackageResponse) SetHeaders(v map[string]*string) *CreateMcubeUpgradePackageResponse {
	s.Headers = v
	return s
}

func (s *CreateMcubeUpgradePackageResponse) SetStatusCode(v int32) *CreateMcubeUpgradePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcubeUpgradePackageResponse) SetBody(v *CreateMcubeUpgradePackageResponseBody) *CreateMcubeUpgradePackageResponse {
	s.Body = v
	return s
}

func (s *CreateMcubeUpgradePackageResponse) Validate() error {
	return dara.Validate(s)
}
