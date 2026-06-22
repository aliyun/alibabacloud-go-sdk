// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSelectionKeyByTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *UpdateSelectionKeyByTypeRequest
	GetBusinessType() *string
	SetSelectionKey(v string) *UpdateSelectionKeyByTypeRequest
	GetSelectionKey() *string
}

type UpdateSelectionKeyByTypeRequest struct {
	// The business type of the asset selection. Valid values:
	//
	// - **VIRUS_SCAN_CYCLE_CONFIG**: virus scan configuration.
	//
	// - **VIRUS_SCAN_ONCE_TASK**: one-time virus scan.
	//
	// - **AGENTLESS_MALICIOUS_WHITE_LIST_[ID]**: agentless detection alert whitelisting rule.
	//
	// - **AGENTLESS_VUL_WHITE_LIST_[ID]**: agentless detection vulnerability whitelisting rule.
	//
	// - **FILE_PROTECT_RULE_SWITCH_TYPE_[ID]**: core file protection.
	//
	// example:
	//
	// VIRUS_SCAN_CYCLE_CONFIG
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The unique identifier of the asset selection.
	//
	// example:
	//
	// 614d179e-4776-4939-a04a-d842ce64****
	SelectionKey *string `json:"SelectionKey,omitempty" xml:"SelectionKey,omitempty"`
}

func (s UpdateSelectionKeyByTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSelectionKeyByTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSelectionKeyByTypeRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *UpdateSelectionKeyByTypeRequest) GetSelectionKey() *string {
	return s.SelectionKey
}

func (s *UpdateSelectionKeyByTypeRequest) SetBusinessType(v string) *UpdateSelectionKeyByTypeRequest {
	s.BusinessType = &v
	return s
}

func (s *UpdateSelectionKeyByTypeRequest) SetSelectionKey(v string) *UpdateSelectionKeyByTypeRequest {
	s.SelectionKey = &v
	return s
}

func (s *UpdateSelectionKeyByTypeRequest) Validate() error {
	return dara.Validate(s)
}
