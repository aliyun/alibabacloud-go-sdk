// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVulConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyVulConfigRequest
	GetClientToken() *string
	SetConfig(v string) *ModifyVulConfigRequest
	GetConfig() *string
	SetType(v string) *ModifyVulConfigRequest
	GetType() *string
}

type ModifyVulConfigRequest struct {
	// The client token that is used to ensure the idempotence of the request. Different requests should use different tokens. The token supports only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable or disable vulnerability detection. Valid values:
	//
	// - **on**: Enable vulnerability detection.
	//
	// - **off**: Disable vulnerability detection.
	//
	// > If the type is set to real risk, valid values:
	//
	// > - **real**: Real risk vulnerabilities.
	//
	// > - **all**: All vulnerabilities.
	//
	// example:
	//
	// on
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The type of vulnerability to modify. Valid values:
	//
	// - **cve**: Linux software vulnerability
	//
	// - **sys**: Windows system vulnerability
	//
	// - **cms**: Web-CMS vulnerability
	//
	// - **emg**: emergency vulnerability
	//
	// - **app**: application vulnerability
	//
	// - **yum**: YUM/APT source configuration
	//
	// - **scanMode**: real risk
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyVulConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVulConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyVulConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyVulConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyVulConfigRequest) GetType() *string {
	return s.Type
}

func (s *ModifyVulConfigRequest) SetClientToken(v string) *ModifyVulConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyVulConfigRequest) SetConfig(v string) *ModifyVulConfigRequest {
	s.Config = &v
	return s
}

func (s *ModifyVulConfigRequest) SetType(v string) *ModifyVulConfigRequest {
	s.Type = &v
	return s
}

func (s *ModifyVulConfigRequest) Validate() error {
	return dara.Validate(s)
}
