// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeUpgradePackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcubeUpgradePackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcubeUpgradePackagesResponse
	GetStatusCode() *int32
	SetBody(v *ListMcubeUpgradePackagesResponseBody) *ListMcubeUpgradePackagesResponse
	GetBody() *ListMcubeUpgradePackagesResponseBody
}

type ListMcubeUpgradePackagesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcubeUpgradePackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcubeUpgradePackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeUpgradePackagesResponse) GoString() string {
	return s.String()
}

func (s *ListMcubeUpgradePackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcubeUpgradePackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcubeUpgradePackagesResponse) GetBody() *ListMcubeUpgradePackagesResponseBody {
	return s.Body
}

func (s *ListMcubeUpgradePackagesResponse) SetHeaders(v map[string]*string) *ListMcubeUpgradePackagesResponse {
	s.Headers = v
	return s
}

func (s *ListMcubeUpgradePackagesResponse) SetStatusCode(v int32) *ListMcubeUpgradePackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponse) SetBody(v *ListMcubeUpgradePackagesResponseBody) *ListMcubeUpgradePackagesResponse {
	s.Body = v
	return s
}

func (s *ListMcubeUpgradePackagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
