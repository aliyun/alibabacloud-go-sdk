// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCheckCustomConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *VerifyCheckCustomConfigShrinkRequest
	GetCheckId() *int64
	SetCustomCheckConfigShrink(v string) *VerifyCheckCustomConfigShrinkRequest
	GetCustomCheckConfigShrink() *string
	SetCustomConfigs(v []*VerifyCheckCustomConfigShrinkRequestCustomConfigs) *VerifyCheckCustomConfigShrinkRequest
	GetCustomConfigs() []*VerifyCheckCustomConfigShrinkRequestCustomConfigs
	SetRepairConfigs(v []*VerifyCheckCustomConfigShrinkRequestRepairConfigs) *VerifyCheckCustomConfigShrinkRequest
	GetRepairConfigs() []*VerifyCheckCustomConfigShrinkRequestRepairConfigs
	SetType(v string) *VerifyCheckCustomConfigShrinkRequest
	GetType() *string
}

type VerifyCheckCustomConfigShrinkRequest struct {
	// Check item ID.
	//
	// example:
	//
	// 76
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// Custom check item to validate input parameters.
	CustomCheckConfigShrink *string `json:"CustomCheckConfig,omitempty" xml:"CustomCheckConfig,omitempty"`
	// List of custom configuration items for the check item.
	CustomConfigs []*VerifyCheckCustomConfigShrinkRequestCustomConfigs `json:"CustomConfigs,omitempty" xml:"CustomConfigs,omitempty" type:"Repeated"`
	// Repair parameters supported by the check item\\"s repair function.
	RepairConfigs []*VerifyCheckCustomConfigShrinkRequestRepairConfigs `json:"RepairConfigs,omitempty" xml:"RepairConfigs,omitempty" type:"Repeated"`
	// Situation Awareness parameter validation types:
	//
	// - **REPAIR_CONFIG**: Repair and custom parameter validation (default)
	//
	// - **CHECK_ITEM_CONFIG**: Custom check item validation
	//
	// example:
	//
	// REPAIR_CONFIG
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s VerifyCheckCustomConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckCustomConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *VerifyCheckCustomConfigShrinkRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *VerifyCheckCustomConfigShrinkRequest) GetCustomCheckConfigShrink() *string {
	return s.CustomCheckConfigShrink
}

func (s *VerifyCheckCustomConfigShrinkRequest) GetCustomConfigs() []*VerifyCheckCustomConfigShrinkRequestCustomConfigs {
	return s.CustomConfigs
}

func (s *VerifyCheckCustomConfigShrinkRequest) GetRepairConfigs() []*VerifyCheckCustomConfigShrinkRequestRepairConfigs {
	return s.RepairConfigs
}

func (s *VerifyCheckCustomConfigShrinkRequest) GetType() *string {
	return s.Type
}

func (s *VerifyCheckCustomConfigShrinkRequest) SetCheckId(v int64) *VerifyCheckCustomConfigShrinkRequest {
	s.CheckId = &v
	return s
}

func (s *VerifyCheckCustomConfigShrinkRequest) SetCustomCheckConfigShrink(v string) *VerifyCheckCustomConfigShrinkRequest {
	s.CustomCheckConfigShrink = &v
	return s
}

func (s *VerifyCheckCustomConfigShrinkRequest) SetCustomConfigs(v []*VerifyCheckCustomConfigShrinkRequestCustomConfigs) *VerifyCheckCustomConfigShrinkRequest {
	s.CustomConfigs = v
	return s
}

func (s *VerifyCheckCustomConfigShrinkRequest) SetRepairConfigs(v []*VerifyCheckCustomConfigShrinkRequestRepairConfigs) *VerifyCheckCustomConfigShrinkRequest {
	s.RepairConfigs = v
	return s
}

func (s *VerifyCheckCustomConfigShrinkRequest) SetType(v string) *VerifyCheckCustomConfigShrinkRequest {
	s.Type = &v
	return s
}

func (s *VerifyCheckCustomConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type VerifyCheckCustomConfigShrinkRequestCustomConfigs struct {
	// Name of the custom configuration item for the check item, unique within the same check item.
	//
	// example:
	//
	// IPList
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Operation type for the custom configuration item of the check item. Only pass DELETE when deleting; no need to pass for creation or update.
	//
	// example:
	//
	// DELETE
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// User-configured value string for the custom configuration item of the check item.
	//
	// example:
	//
	// 10.12.4.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s VerifyCheckCustomConfigShrinkRequestCustomConfigs) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckCustomConfigShrinkRequestCustomConfigs) GoString() string {
	return s.String()
}

func (s *VerifyCheckCustomConfigShrinkRequestCustomConfigs) GetName() *string {
	return s.Name
}

func (s *VerifyCheckCustomConfigShrinkRequestCustomConfigs) GetOperation() *string {
	return s.Operation
}

func (s *VerifyCheckCustomConfigShrinkRequestCustomConfigs) GetValue() *string {
	return s.Value
}

func (s *VerifyCheckCustomConfigShrinkRequestCustomConfigs) SetName(v string) *VerifyCheckCustomConfigShrinkRequestCustomConfigs {
	s.Name = &v
	return s
}

func (s *VerifyCheckCustomConfigShrinkRequestCustomConfigs) SetOperation(v string) *VerifyCheckCustomConfigShrinkRequestCustomConfigs {
	s.Operation = &v
	return s
}

func (s *VerifyCheckCustomConfigShrinkRequestCustomConfigs) SetValue(v string) *VerifyCheckCustomConfigShrinkRequestCustomConfigs {
	s.Value = &v
	return s
}

func (s *VerifyCheckCustomConfigShrinkRequestCustomConfigs) Validate() error {
	return dara.Validate(s)
}

type VerifyCheckCustomConfigShrinkRequestRepairConfigs struct {
	// ID of the repair process during the repair.
	//
	// example:
	//
	// 7fec0a3395b345c18f108ffc9fc0****
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// Name of the repair parameter for the check item, unique within the same check item.
	//
	// example:
	//
	// IPLists
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Operation type for the custom configuration item of the check item. Only pass DELETE when deleting; no need to pass for creation or update.
	//
	// example:
	//
	// DELETE
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// User-configured value string for the repair parameter of the check item.
	//
	// example:
	//
	// 172.26.49.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s VerifyCheckCustomConfigShrinkRequestRepairConfigs) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckCustomConfigShrinkRequestRepairConfigs) GoString() string {
	return s.String()
}

func (s *VerifyCheckCustomConfigShrinkRequestRepairConfigs) GetFlowId() *string {
	return s.FlowId
}

func (s *VerifyCheckCustomConfigShrinkRequestRepairConfigs) GetName() *string {
	return s.Name
}

func (s *VerifyCheckCustomConfigShrinkRequestRepairConfigs) GetOperation() *string {
	return s.Operation
}

func (s *VerifyCheckCustomConfigShrinkRequestRepairConfigs) GetValue() *string {
	return s.Value
}

func (s *VerifyCheckCustomConfigShrinkRequestRepairConfigs) SetFlowId(v string) *VerifyCheckCustomConfigShrinkRequestRepairConfigs {
	s.FlowId = &v
	return s
}

func (s *VerifyCheckCustomConfigShrinkRequestRepairConfigs) SetName(v string) *VerifyCheckCustomConfigShrinkRequestRepairConfigs {
	s.Name = &v
	return s
}

func (s *VerifyCheckCustomConfigShrinkRequestRepairConfigs) SetOperation(v string) *VerifyCheckCustomConfigShrinkRequestRepairConfigs {
	s.Operation = &v
	return s
}

func (s *VerifyCheckCustomConfigShrinkRequestRepairConfigs) SetValue(v string) *VerifyCheckCustomConfigShrinkRequestRepairConfigs {
	s.Value = &v
	return s
}

func (s *VerifyCheckCustomConfigShrinkRequestRepairConfigs) Validate() error {
	return dara.Validate(s)
}
