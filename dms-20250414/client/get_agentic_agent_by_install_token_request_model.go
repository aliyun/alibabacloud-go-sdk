// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticAgentByInstallTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstallToken(v string) *GetAgenticAgentByInstallTokenRequest
	GetInstallToken() *string
}

type GetAgenticAgentByInstallTokenRequest struct {
	// This parameter is required.
	InstallToken *string `json:"InstallToken,omitempty" xml:"InstallToken,omitempty"`
}

func (s GetAgenticAgentByInstallTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticAgentByInstallTokenRequest) GoString() string {
	return s.String()
}

func (s *GetAgenticAgentByInstallTokenRequest) GetInstallToken() *string {
	return s.InstallToken
}

func (s *GetAgenticAgentByInstallTokenRequest) SetInstallToken(v string) *GetAgenticAgentByInstallTokenRequest {
	s.InstallToken = &v
	return s
}

func (s *GetAgenticAgentByInstallTokenRequest) Validate() error {
	return dara.Validate(s)
}
