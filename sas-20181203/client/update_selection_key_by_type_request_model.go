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
	SetClientToken(v string) *UpdateSelectionKeyByTypeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateSelectionKeyByTypeRequest
	GetDryRun() *bool
	SetSelectionKey(v string) *UpdateSelectionKeyByTypeRequest
	GetSelectionKey() *string
}

type UpdateSelectionKeyByTypeRequest struct {
	// The business type of the asset selection. Valid values:
	//
	// - **VIRUS_SCAN_CYCLE_CONFIG**: virus scan cycle configuration
	//
	// - **VIRUS_SCAN_ONCE_TASK**: one-time virus scan task
	//
	// - **AGENTLESS_MALICIOUS_WHITE_LIST_[ID]**: agentless detection alert whitelist rule
	//
	// - **AGENTLESS_VUL_WHITE_LIST_[ID]**: agentless detection vulnerability whitelist rule
	//
	// - **FILE_PROTECT_RULE_SWITCH_TYPE_[ID]**: core file protection
	//
	// example:
	//
	// VIRUS_SCAN_CYCLE_CONFIG
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The client token used to ensure the idempotence of the request. Use a different token for different requests. Only ASCII characters are supported. The token can be up to 64 characters in length.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run for this request. Valid values:
	//
	// - true: performs only a dry run without executing the actual operation.
	//
	// - false: executes the request normally.
	//
	// Default value: false.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
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

func (s *UpdateSelectionKeyByTypeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateSelectionKeyByTypeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateSelectionKeyByTypeRequest) GetSelectionKey() *string {
	return s.SelectionKey
}

func (s *UpdateSelectionKeyByTypeRequest) SetBusinessType(v string) *UpdateSelectionKeyByTypeRequest {
	s.BusinessType = &v
	return s
}

func (s *UpdateSelectionKeyByTypeRequest) SetClientToken(v string) *UpdateSelectionKeyByTypeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSelectionKeyByTypeRequest) SetDryRun(v bool) *UpdateSelectionKeyByTypeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateSelectionKeyByTypeRequest) SetSelectionKey(v string) *UpdateSelectionKeyByTypeRequest {
	s.SelectionKey = &v
	return s
}

func (s *UpdateSelectionKeyByTypeRequest) Validate() error {
	return dara.Validate(s)
}
