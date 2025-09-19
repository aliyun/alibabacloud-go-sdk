// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAssetSelectionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *CreateAssetSelectionConfigRequest
	GetBusinessType() *string
	SetPlatform(v string) *CreateAssetSelectionConfigRequest
	GetPlatform() *string
	SetTargetType(v string) *CreateAssetSelectionConfigRequest
	GetTargetType() *string
}

type CreateAssetSelectionConfigRequest struct {
	// The feature that you want to select for the asset. Valid values:
	//
	// 	- **VIRUS_SCAN_CYCLE_CONFIG**: virus detection and removal
	//
	// 	- **VIRUS_SCAN_ONCE_TASK**: one-time scan for viruses
	//
	// 	- **AGENTLESS_MALICIOUS_WHITE_LIST_[ID]**: a whitelist rule for alerts that are detected by using the agentless detection feature
	//
	// 	- **AGENTLESS_VUL_WHITE_LIST_[ID]**: a whitelist rule for vulnerabilities that are detected by using the agentless detection feature
	//
	// 	- **FILE_PROTECT_RULE_SWITCH_TYPE_[ID]**: core file protection
	//
	// This parameter is required.
	//
	// example:
	//
	// VIRUS_SCAN_CYCLE_CONFIG
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The operating system of the asset. Valid values:
	//
	// 	- **all**: all operating systems
	//
	// 	- **windows**: the Windows operating system
	//
	// 	- **linux**: the Linux operating system
	//
	// >  If you leave this parameter empty, the system automatically selects a value for the parameter based on the value of the **BusinessType*	- parameter.
	//
	// 	- If the BusinessType parameter is set to **VIRUS_SCAN_CYCLE_CONFIG**, the value of the Platform parameter is **all**.
	//
	// 	- If the BusinessType parameter is set to **VIRUS_SCAN_ONCE_TASK**, the value of the Platform parameter is **all**.
	//
	// 	- If the BusinessType parameter is set to **AGENTLESS_MALICIOUS_WHITE_LIST_[ID]**, the value of the Platform parameter is **all**.
	//
	// 	- If the BusinessType parameter is set to **AGENTLESS_VUL_WHITE_LIST_[ID]*	- the value of the Platform parameter is **all**.
	//
	// 	- If the BusinessType parameter is set to **FILE_PROTECT_RULE_SWITCH_TYPE_[ID]**, the value of the Platform parameter is **linux**.
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The dimension based on which you want to select the asset. Valid values:
	//
	// 	- **instance**: selects the asset by server.
	//
	// 	- **group**: selects the asset by group.
	//
	// 	- **vpc**: selects the asset by virtual private cloud (VPC).
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateAssetSelectionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAssetSelectionConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateAssetSelectionConfigRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateAssetSelectionConfigRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateAssetSelectionConfigRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateAssetSelectionConfigRequest) SetBusinessType(v string) *CreateAssetSelectionConfigRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateAssetSelectionConfigRequest) SetPlatform(v string) *CreateAssetSelectionConfigRequest {
	s.Platform = &v
	return s
}

func (s *CreateAssetSelectionConfigRequest) SetTargetType(v string) *CreateAssetSelectionConfigRequest {
	s.TargetType = &v
	return s
}

func (s *CreateAssetSelectionConfigRequest) Validate() error {
	return dara.Validate(s)
}
