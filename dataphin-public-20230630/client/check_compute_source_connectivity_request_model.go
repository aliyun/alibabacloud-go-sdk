// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckComputeSourceConnectivityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckCommand(v *CheckComputeSourceConnectivityRequestCheckCommand) *CheckComputeSourceConnectivityRequest
	GetCheckCommand() *CheckComputeSourceConnectivityRequestCheckCommand
	SetOpTenantId(v int64) *CheckComputeSourceConnectivityRequest
	GetOpTenantId() *int64
}

type CheckComputeSourceConnectivityRequest struct {
	// This parameter is required.
	CheckCommand *CheckComputeSourceConnectivityRequestCheckCommand `json:"CheckCommand,omitempty" xml:"CheckCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CheckComputeSourceConnectivityRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckComputeSourceConnectivityRequest) GoString() string {
	return s.String()
}

func (s *CheckComputeSourceConnectivityRequest) GetCheckCommand() *CheckComputeSourceConnectivityRequestCheckCommand {
	return s.CheckCommand
}

func (s *CheckComputeSourceConnectivityRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CheckComputeSourceConnectivityRequest) SetCheckCommand(v *CheckComputeSourceConnectivityRequestCheckCommand) *CheckComputeSourceConnectivityRequest {
	s.CheckCommand = v
	return s
}

func (s *CheckComputeSourceConnectivityRequest) SetOpTenantId(v int64) *CheckComputeSourceConnectivityRequest {
	s.OpTenantId = &v
	return s
}

func (s *CheckComputeSourceConnectivityRequest) Validate() error {
	if s.CheckCommand != nil {
		if err := s.CheckCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckComputeSourceConnectivityRequestCheckCommand struct {
	// This parameter is required.
	ConfigList []*CheckComputeSourceConnectivityRequestCheckCommandConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// MAX_COMPUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CheckComputeSourceConnectivityRequestCheckCommand) String() string {
	return dara.Prettify(s)
}

func (s CheckComputeSourceConnectivityRequestCheckCommand) GoString() string {
	return s.String()
}

func (s *CheckComputeSourceConnectivityRequestCheckCommand) GetConfigList() []*CheckComputeSourceConnectivityRequestCheckCommandConfigList {
	return s.ConfigList
}

func (s *CheckComputeSourceConnectivityRequestCheckCommand) GetType() *string {
	return s.Type
}

func (s *CheckComputeSourceConnectivityRequestCheckCommand) SetConfigList(v []*CheckComputeSourceConnectivityRequestCheckCommandConfigList) *CheckComputeSourceConnectivityRequestCheckCommand {
	s.ConfigList = v
	return s
}

func (s *CheckComputeSourceConnectivityRequestCheckCommand) SetType(v string) *CheckComputeSourceConnectivityRequestCheckCommand {
	s.Type = &v
	return s
}

func (s *CheckComputeSourceConnectivityRequestCheckCommand) Validate() error {
	if s.ConfigList != nil {
		for _, item := range s.ConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckComputeSourceConnectivityRequestCheckCommandConfigList struct {
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CheckComputeSourceConnectivityRequestCheckCommandConfigList) String() string {
	return dara.Prettify(s)
}

func (s CheckComputeSourceConnectivityRequestCheckCommandConfigList) GoString() string {
	return s.String()
}

func (s *CheckComputeSourceConnectivityRequestCheckCommandConfigList) GetKey() *string {
	return s.Key
}

func (s *CheckComputeSourceConnectivityRequestCheckCommandConfigList) GetValue() *string {
	return s.Value
}

func (s *CheckComputeSourceConnectivityRequestCheckCommandConfigList) SetKey(v string) *CheckComputeSourceConnectivityRequestCheckCommandConfigList {
	s.Key = &v
	return s
}

func (s *CheckComputeSourceConnectivityRequestCheckCommandConfigList) SetValue(v string) *CheckComputeSourceConnectivityRequestCheckCommandConfigList {
	s.Value = &v
	return s
}

func (s *CheckComputeSourceConnectivityRequestCheckCommandConfigList) Validate() error {
	return dara.Validate(s)
}
