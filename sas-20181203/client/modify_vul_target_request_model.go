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
	// The configuration target. This parameter is in JSON format and contains the following fields:
	//
	// - **vulType**: The vulnerability type. Valid values:
	//
	//      - **cve**: Linux software vulnerability.
	//
	//     - **sys**: Windows system vulnerability.
	//
	//     - **cms**: Web-CMS vulnerability.
	//
	//     - **emg**: Emergency vulnerability.
	//
	// example:
	//
	// {\\"vulType\\":\\"sys\\"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The operation target. This parameter is in JSON format and contains the following fields:
	//
	// - **target**: The UUID of the target machine.
	//
	// - **targetType**: The target type. Fixed value: uuid.
	//
	// - **flag**: The flag. Valid values:
	//
	//     - **add**: Selected.
	//
	//     - **del**: Deselected.
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
