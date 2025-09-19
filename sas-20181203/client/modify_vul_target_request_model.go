// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVulTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyVulTargetRequest
	GetConfig() *string
	SetTarget(v string) *ModifyVulTargetRequest
	GetTarget() *string
}

type ModifyVulTargetRequest struct {
	// The configurations. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **vulType**: the type of the vulnerabilities to scan. Valid values:
	//
	//     	- **cve**: Linux software vulnerabilities
	//
	//     	- **sys**: Windows system vulnerabilities
	//
	//     	- **cms**: Web-CMS vulnerabilities
	//
	//     	- **emg**: urgent vulnerabilities
	//
	// example:
	//
	// {\\"vulType\\":\\"sys\\"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The operation. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **target**: the UUID of the server.
	//
	// 	- **targetType**: the application scope of the operation. Set the value to uuid.
	//
	// 	- **flag**: the type of the operation. Valid values:
	//
	//     	- **add**: select
	//
	//     	- **del**: deselect
	//
	// example:
	//
	// [{\\"target\\": \\"9cd5c684-7201-4de5-ad2c-cea89a5e****\\", \\"targetType\\": \\"uuid\\", \\"flag\\": \\"add\\"}]
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyVulTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVulTargetRequest) GoString() string {
	return s.String()
}

func (s *ModifyVulTargetRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyVulTargetRequest) GetTarget() *string {
	return s.Target
}

func (s *ModifyVulTargetRequest) SetConfig(v string) *ModifyVulTargetRequest {
	s.Config = &v
	return s
}

func (s *ModifyVulTargetRequest) SetTarget(v string) *ModifyVulTargetRequest {
	s.Target = &v
	return s
}

func (s *ModifyVulTargetRequest) Validate() error {
	return dara.Validate(s)
}
