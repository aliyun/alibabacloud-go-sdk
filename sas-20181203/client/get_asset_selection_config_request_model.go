// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetSelectionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *GetAssetSelectionConfigRequest
	GetBusinessType() *string
}

type GetAssetSelectionConfigRequest struct {
	// The business type of the asset selection. Valid values:
	//
	// - **VIRUS_SCAN_CYCLE_CONFIG**: trojan scan configuration
	//
	// - **VIRUS_SCAN_ONCE_TASK**: trojan scan one-time scan
	//
	// - **AGENTLESS_MALICIOUS_WHITE_LIST_[ID]**: agentless detection alert whitelisting rule
	//
	// - **AGENTLESS_VUL_WHITE_LIST_[ID]**: agentless detection vulnerability whitelisting rule
	//
	// - **FILE_PROTECT_RULE_SWITCH_TYPE_[ID]**: core file protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// VIRUS_SCAN_CYCLE_CONFIG
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
}

func (s GetAssetSelectionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAssetSelectionConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAssetSelectionConfigRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetAssetSelectionConfigRequest) SetBusinessType(v string) *GetAssetSelectionConfigRequest {
	s.BusinessType = &v
	return s
}

func (s *GetAssetSelectionConfigRequest) Validate() error {
	return dara.Validate(s)
}
