// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcubeUpgradePackageInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMcubeUpgradePackageInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMcubeUpgradePackageInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetMcubeUpgradePackageInfoResponseBody) *GetMcubeUpgradePackageInfoResponse
	GetBody() *GetMcubeUpgradePackageInfoResponseBody
}

type GetMcubeUpgradePackageInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcubeUpgradePackageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcubeUpgradePackageInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeUpgradePackageInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMcubeUpgradePackageInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcubeUpgradePackageInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMcubeUpgradePackageInfoResponse) GetBody() *GetMcubeUpgradePackageInfoResponseBody {
	return s.Body
}

func (s *GetMcubeUpgradePackageInfoResponse) SetHeaders(v map[string]*string) *GetMcubeUpgradePackageInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponse) SetStatusCode(v int32) *GetMcubeUpgradePackageInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponse) SetBody(v *GetMcubeUpgradePackageInfoResponseBody) *GetMcubeUpgradePackageInfoResponse {
	s.Body = v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
