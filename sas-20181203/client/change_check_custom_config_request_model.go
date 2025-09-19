// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCheckCustomConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *ChangeCheckCustomConfigRequest
	GetCheckId() *int64
	SetCustomConfigs(v []*ChangeCheckCustomConfigRequestCustomConfigs) *ChangeCheckCustomConfigRequest
	GetCustomConfigs() []*ChangeCheckCustomConfigRequestCustomConfigs
	SetRegionId(v string) *ChangeCheckCustomConfigRequest
	GetRegionId() *string
	SetRepairConfigs(v []*ChangeCheckCustomConfigRequestRepairConfigs) *ChangeCheckCustomConfigRequest
	GetRepairConfigs() []*ChangeCheckCustomConfigRequestRepairConfigs
}

type ChangeCheckCustomConfigRequest struct {
	// The ID of the check item.
	//
	// > You can call the [ListCheckResult](~~ListCheckResult~~) operation to query the IDs of check items.
	//
	// example:
	//
	// 76
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The custom configuration items of the check item.
	CustomConfigs []*ChangeCheckCustomConfigRequestCustomConfigs `json:"CustomConfigs,omitempty" xml:"CustomConfigs,omitempty" type:"Repeated"`
	// The region where the Security Center instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: International
	//
	// 	- **ap-southeast-1**: Singapore
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The parameters required for fixing risk items.
	RepairConfigs []*ChangeCheckCustomConfigRequestRepairConfigs `json:"RepairConfigs,omitempty" xml:"RepairConfigs,omitempty" type:"Repeated"`
}

func (s ChangeCheckCustomConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckCustomConfigRequest) GoString() string {
	return s.String()
}

func (s *ChangeCheckCustomConfigRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ChangeCheckCustomConfigRequest) GetCustomConfigs() []*ChangeCheckCustomConfigRequestCustomConfigs {
	return s.CustomConfigs
}

func (s *ChangeCheckCustomConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ChangeCheckCustomConfigRequest) GetRepairConfigs() []*ChangeCheckCustomConfigRequestRepairConfigs {
	return s.RepairConfigs
}

func (s *ChangeCheckCustomConfigRequest) SetCheckId(v int64) *ChangeCheckCustomConfigRequest {
	s.CheckId = &v
	return s
}

func (s *ChangeCheckCustomConfigRequest) SetCustomConfigs(v []*ChangeCheckCustomConfigRequestCustomConfigs) *ChangeCheckCustomConfigRequest {
	s.CustomConfigs = v
	return s
}

func (s *ChangeCheckCustomConfigRequest) SetRegionId(v string) *ChangeCheckCustomConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeCheckCustomConfigRequest) SetRepairConfigs(v []*ChangeCheckCustomConfigRequestRepairConfigs) *ChangeCheckCustomConfigRequest {
	s.RepairConfigs = v
	return s
}

func (s *ChangeCheckCustomConfigRequest) Validate() error {
	return dara.Validate(s)
}

type ChangeCheckCustomConfigRequestCustomConfigs struct {
	// The name of the custom configuration item. The name of a custom configuration item is unique in a check item.
	//
	// example:
	//
	// SessionTimeMax
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation that you want to perform on the custom configuration item. This parameter is required only if you want to delete the custom configuration item. To delete the custom configuration item, set the value to DELETE.
	//
	// example:
	//
	// DELETE
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The value of the custom configuration item. The value is a string.
	//
	// example:
	//
	// 13
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ChangeCheckCustomConfigRequestCustomConfigs) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckCustomConfigRequestCustomConfigs) GoString() string {
	return s.String()
}

func (s *ChangeCheckCustomConfigRequestCustomConfigs) GetName() *string {
	return s.Name
}

func (s *ChangeCheckCustomConfigRequestCustomConfigs) GetOperation() *string {
	return s.Operation
}

func (s *ChangeCheckCustomConfigRequestCustomConfigs) GetValue() *string {
	return s.Value
}

func (s *ChangeCheckCustomConfigRequestCustomConfigs) SetName(v string) *ChangeCheckCustomConfigRequestCustomConfigs {
	s.Name = &v
	return s
}

func (s *ChangeCheckCustomConfigRequestCustomConfigs) SetOperation(v string) *ChangeCheckCustomConfigRequestCustomConfigs {
	s.Operation = &v
	return s
}

func (s *ChangeCheckCustomConfigRequestCustomConfigs) SetValue(v string) *ChangeCheckCustomConfigRequestCustomConfigs {
	s.Value = &v
	return s
}

func (s *ChangeCheckCustomConfigRequestCustomConfigs) Validate() error {
	return dara.Validate(s)
}

type ChangeCheckCustomConfigRequestRepairConfigs struct {
	// The ID of the fixing process.
	//
	// example:
	//
	// ascgrmscyjgs*********
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The name of the parameter required for fixing a risk item, which is unique in a check item.
	//
	// example:
	//
	// Port
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation that you want to perform on the custom configuration item. This parameter is required only if you want to delete the custom configuration item. To delete the custom configuration item, set the value to DELETE.
	//
	// example:
	//
	// DELETE
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The value of the parameter required for fixing a risk item. The value is a string.
	//
	// example:
	//
	// 80
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ChangeCheckCustomConfigRequestRepairConfigs) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckCustomConfigRequestRepairConfigs) GoString() string {
	return s.String()
}

func (s *ChangeCheckCustomConfigRequestRepairConfigs) GetFlowId() *string {
	return s.FlowId
}

func (s *ChangeCheckCustomConfigRequestRepairConfigs) GetName() *string {
	return s.Name
}

func (s *ChangeCheckCustomConfigRequestRepairConfigs) GetOperation() *string {
	return s.Operation
}

func (s *ChangeCheckCustomConfigRequestRepairConfigs) GetValue() *string {
	return s.Value
}

func (s *ChangeCheckCustomConfigRequestRepairConfigs) SetFlowId(v string) *ChangeCheckCustomConfigRequestRepairConfigs {
	s.FlowId = &v
	return s
}

func (s *ChangeCheckCustomConfigRequestRepairConfigs) SetName(v string) *ChangeCheckCustomConfigRequestRepairConfigs {
	s.Name = &v
	return s
}

func (s *ChangeCheckCustomConfigRequestRepairConfigs) SetOperation(v string) *ChangeCheckCustomConfigRequestRepairConfigs {
	s.Operation = &v
	return s
}

func (s *ChangeCheckCustomConfigRequestRepairConfigs) SetValue(v string) *ChangeCheckCustomConfigRequestRepairConfigs {
	s.Value = &v
	return s
}

func (s *ChangeCheckCustomConfigRequestRepairConfigs) Validate() error {
	return dara.Validate(s)
}
