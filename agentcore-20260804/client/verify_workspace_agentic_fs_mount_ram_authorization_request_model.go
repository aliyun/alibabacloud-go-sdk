// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyWorkspaceAgenticFsMountRamAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServer(v string) *VerifyWorkspaceAgenticFsMountRamAuthorizationRequest
	GetServer() *string
}

type VerifyWorkspaceAgenticFsMountRamAuthorizationRequest struct {
	// The domain name of the target AccessPoint, obtained from the DomainName field of NAS ListAccessPoints. Do not include the protocol, port, or path.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-0123456789abcdef0.0123456789-vlm36.cn-hangzhou.nas.aliyuncs.com
	Server *string `json:"server,omitempty" xml:"server,omitempty"`
}

func (s VerifyWorkspaceAgenticFsMountRamAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyWorkspaceAgenticFsMountRamAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationRequest) GetServer() *string {
	return s.Server
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationRequest) SetServer(v string) *VerifyWorkspaceAgenticFsMountRamAuthorizationRequest {
	s.Server = &v
	return s
}

func (s *VerifyWorkspaceAgenticFsMountRamAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
