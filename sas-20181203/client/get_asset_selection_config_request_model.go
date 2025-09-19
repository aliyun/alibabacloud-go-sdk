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
	// The feature that is selected for the asset. Valid values:
	//
	// 	- **VIRUS_SCAN_CYCLE_CONFIG**: virus detection and removal
	//
	// 	- **VIRUS_SCAN_ONCE_TASK**: one-time scan for viruses
	//
	// 	- **AGENTLESS_MALICIOUS_WHITE_LIST_[ID]**: a whitelist rule for alerts that are detected by using the agentless detection feature
	//
	// 	- **AGENTLESS_VUL_WHITE_LIST_[ID]**: a whitelist rule for vulnerabilities that are detected by using the agentless detection feature
	//
	// 	- **FILE_PROTECT_RULE_SWITCH_TYPE_[ID]**: core file protectioion
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
