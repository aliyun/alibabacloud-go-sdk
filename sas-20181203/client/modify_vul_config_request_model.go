// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVulConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyVulConfigRequest
	GetConfig() *string
	SetType(v string) *ModifyVulConfigRequest
	GetType() *string
}

type ModifyVulConfigRequest struct {
	// Specifies whether to enable the vulnerability scan feature. Valid values:
	//
	// 	- **on**: enables the feature
	//
	// 	- **off**: disables the feature
	//
	// > Valid values when you set the Type parameter to scanMode:
	//
	// 	- **real**: displays only easily exploitable vulnerabilities.
	//
	// 	- **all**: displays all vulnerabilities.
	//
	// example:
	//
	// on
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: Linux software vulnerability
	//
	// 	- **sys**: Windows system vulnerability
	//
	// 	- **cms**: Web-CMS vulnerability
	//
	// 	- **emg**: urgent vulnerability
	//
	// 	- **app**: application vulnerability
	//
	// 	- **yum**: YUM and APT source configuration
	//
	// 	- **scanMode**: easily exploitable vulnerability
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

func (s *ModifyVulConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyVulConfigRequest) GetType() *string {
	return s.Type
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
