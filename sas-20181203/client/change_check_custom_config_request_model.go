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
	// > You can call the [ListCheckResult](~~ListCheckResult~~) operation to obtain the check item ID.
	//
	// example:
	//
	// 76
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The list of custom configuration items for the check item.
	CustomConfigs []*ChangeCheckCustomConfigRequestCustomConfigs `json:"CustomConfigs,omitempty" xml:"CustomConfigs,omitempty" type:"Repeated"`
	// The region of the Security Center instance. Valid values:
	//
	// - **cn-hangzhou:*	- China
	//
	// - **ap-southeast-1:*	- Singapore
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The repair parameters supported by the repair feature of the check item.
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
	if s.CustomConfigs != nil {
		for _, item := range s.CustomConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RepairConfigs != nil {
		for _, item := range s.RepairConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChangeCheckCustomConfigRequestCustomConfigs struct {
	// The name of the custom configuration item, which is unique within the check item.
	//
	// example:
	//
	// SessionTimeMax
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation type of the custom configuration item. Set this parameter to DELETE only when deleting a configuration item. You do not need to specify this parameter for create or update operations.
	//
	// example:
	//
	// DELETE
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The user-configured value string of the custom configuration item.
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
	// The ID of the repair flow used during the repair process.
	//
	// example:
	//
	// ascgrmscyjgs*********
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The name of the repair parameter, which is unique within the check item.
	//
	// example:
	//
	// Port
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation type of the custom configuration item. Set this parameter to DELETE only when deleting a configuration item. You do not need to specify this parameter for create or update operations.
	//
	// example:
	//
	// DELETE
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The user-configured value string of the repair configuration item.
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
