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
	SetClientToken(v string) *CreateAssetSelectionConfigRequest
	GetClientToken() *string
	SetPlatform(v string) *CreateAssetSelectionConfigRequest
	GetPlatform() *string
	SetTargetType(v string) *CreateAssetSelectionConfigRequest
	GetTargetType() *string
}

type CreateAssetSelectionConfigRequest struct {
	// The business type of the asset selection. Valid values:
	//
	// - **VIRUS_SCAN_CYCLE_CONFIG**: trojan scan configuration.
	//
	// - **VIRUS_SCAN_ONCE_TASK**: trojan scan one-time scan.
	//
	// - **AGENTLESS_MALICIOUS_WHITE_LIST_[ID]**: agentless detection alert whitelisting rule.
	//
	// - **AGENTLESS_VUL_WHITE_LIST_[ID]**: agentless detection vulnerability whitelisting rule.
	//
	// - **FILE_PROTECT_RULE_SWITCH_TYPE_[ID]**: core file protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// VIRUS_SCAN_CYCLE_CONFIG
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Different requests must use different tokens. The token supports only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The operating system of the target asset. Valid values:
	//
	// - **all**: all operating systems.
	//
	// - **windows**: Windows operating system.
	//
	// - **linux**: Linux operating system.
	//
	// > If this parameter is left empty, the default value is determined based on the **BusinessType*	- value.
	//
	// >- **VIRUS_SCAN_CYCLE_CONFIG**: the value is **all**.
	//
	// >- **VIRUS_SCAN_ONCE_TASK**: the value is **all**.
	//
	// >- **AGENTLESS_MALICIOUS_WHITE_LIST_[ID]**: the value is **all**.
	//
	// >- **AGENTLESS_VUL_WHITE_LIST_[ID]**: the value is **all**.
	//
	// >- **FILE_PROTECT_RULE_SWITCH_TYPE_[ID]**: the value is **linux**.
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The target asset type. Valid values:
	//
	// - **all_instance**: all servers.
	//
	// - **instance**: select by server.
	//
	// - **group**: select by group.
	//
	// - **vpc**: select by VPC.
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

func (s *CreateAssetSelectionConfigRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *CreateAssetSelectionConfigRequest) SetClientToken(v string) *CreateAssetSelectionConfigRequest {
	s.ClientToken = &v
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
