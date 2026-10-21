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
	// The one-time install token issued by IssueAgenticInstallToken. The token starts with dit_ followed by 40 URL-safe Base64 characters, making a total length of 44 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// dit_0123456789abcdefghijklmnopqrstuvwxyzABCD
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
