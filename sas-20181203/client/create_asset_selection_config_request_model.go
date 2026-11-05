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
	SetDryRun(v bool) *CreateAssetSelectionConfigRequest
	GetDryRun() *bool
	SetPlatform(v string) *CreateAssetSelectionConfigRequest
	GetPlatform() *string
	SetTargetType(v string) *CreateAssetSelectionConfigRequest
	GetTargetType() *string
}

type CreateAssetSelectionConfigRequest struct {
	// The business type of the asset selection. Valid values:
	//
	// - **VIRUS_SCAN_CYCLE_CONFIG**: Virus scan configuration.
	//
	// - **VIRUS_SCAN_ONCE_TASK**: One-time virus scan.
	//
	// - **AGENTLESS_MALICIOUS_WHITE_LIST_[ID]**: Whitelist rule for agentless detection alerts.
	//
	// - **AGENTLESS_VUL_WHITE_LIST_[ID]**: Whitelist rule for agentless detection vulnerabilities.
	//
	// - **FILE_PROTECT_RULE_SWITCH_TYPE_[ID]**: Core file protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// VIRUS_SCAN_CYCLE_CONFIG
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use different tokens for different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - true: Prechecks the request without performing the actual operation.
	//
	// - false: Performs the request.
	//
	// Default value: false.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The operating system of the target asset. Valid values:
	//
	// - **all**: All operating systems.
	//
	// - **windows**: Windows operating system.
	//
	// - **linux**: Linux operating system.
	//
	// > If this parameter is left empty, the default value is matched based on the **BusinessType*	- value.
	//
	// > - **VIRUS_SCAN_CYCLE_CONFIG**: The value is **all**.
	//
	// > - **VIRUS_SCAN_ONCE_TASK**: The value is **all**.
	//
	// > - **AGENTLESS_MALICIOUS_WHITE_LIST_[ID]**: The value is **all**.
	//
	// > - **AGENTLESS_VUL_WHITE_LIST_[ID]**: The value is **all**.
	//
	// > - **FILE_PROTECT_RULE_SWITCH_TYPE_[ID]**: The value is **linux**.
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The target Asset Type. Valid values:
	//
	// - **all_instance**: All servers.
	//
	// - **instance**: By server.
	//
	// - **group**: By group.
	//
	// - **vpc**: By VPC.
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

func (s *CreateAssetSelectionConfigRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *CreateAssetSelectionConfigRequest) SetDryRun(v bool) *CreateAssetSelectionConfigRequest {
	s.DryRun = &v
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
