// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataSourceConnectivityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckCommand(v *CheckDataSourceConnectivityRequestCheckCommand) *CheckDataSourceConnectivityRequest
	GetCheckCommand() *CheckDataSourceConnectivityRequestCheckCommand
	SetOpTenantId(v int64) *CheckDataSourceConnectivityRequest
	GetOpTenantId() *int64
}

type CheckDataSourceConnectivityRequest struct {
	// This parameter is required.
	CheckCommand *CheckDataSourceConnectivityRequestCheckCommand `json:"CheckCommand,omitempty" xml:"CheckCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CheckDataSourceConnectivityRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDataSourceConnectivityRequest) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityRequest) GetCheckCommand() *CheckDataSourceConnectivityRequestCheckCommand {
	return s.CheckCommand
}

func (s *CheckDataSourceConnectivityRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CheckDataSourceConnectivityRequest) SetCheckCommand(v *CheckDataSourceConnectivityRequestCheckCommand) *CheckDataSourceConnectivityRequest {
	s.CheckCommand = v
	return s
}

func (s *CheckDataSourceConnectivityRequest) SetOpTenantId(v int64) *CheckDataSourceConnectivityRequest {
	s.OpTenantId = &v
	return s
}

func (s *CheckDataSourceConnectivityRequest) Validate() error {
	if s.CheckCommand != nil {
		if err := s.CheckCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckDataSourceConnectivityRequestCheckCommand struct {
	// This parameter is required.
	ConfigItemList []*CheckDataSourceConnectivityRequestCheckCommandConfigItemList `json:"ConfigItemList,omitempty" xml:"ConfigItemList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// MAX_COMPUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CheckDataSourceConnectivityRequestCheckCommand) String() string {
	return dara.Prettify(s)
}

func (s CheckDataSourceConnectivityRequestCheckCommand) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityRequestCheckCommand) GetConfigItemList() []*CheckDataSourceConnectivityRequestCheckCommandConfigItemList {
	return s.ConfigItemList
}

func (s *CheckDataSourceConnectivityRequestCheckCommand) GetType() *string {
	return s.Type
}

func (s *CheckDataSourceConnectivityRequestCheckCommand) SetConfigItemList(v []*CheckDataSourceConnectivityRequestCheckCommandConfigItemList) *CheckDataSourceConnectivityRequestCheckCommand {
	s.ConfigItemList = v
	return s
}

func (s *CheckDataSourceConnectivityRequestCheckCommand) SetType(v string) *CheckDataSourceConnectivityRequestCheckCommand {
	s.Type = &v
	return s
}

func (s *CheckDataSourceConnectivityRequestCheckCommand) Validate() error {
	if s.ConfigItemList != nil {
		for _, item := range s.ConfigItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckDataSourceConnectivityRequestCheckCommandConfigItemList struct {
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

func (s CheckDataSourceConnectivityRequestCheckCommandConfigItemList) String() string {
	return dara.Prettify(s)
}

func (s CheckDataSourceConnectivityRequestCheckCommandConfigItemList) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityRequestCheckCommandConfigItemList) GetKey() *string {
	return s.Key
}

func (s *CheckDataSourceConnectivityRequestCheckCommandConfigItemList) GetValue() *string {
	return s.Value
}

func (s *CheckDataSourceConnectivityRequestCheckCommandConfigItemList) SetKey(v string) *CheckDataSourceConnectivityRequestCheckCommandConfigItemList {
	s.Key = &v
	return s
}

func (s *CheckDataSourceConnectivityRequestCheckCommandConfigItemList) SetValue(v string) *CheckDataSourceConnectivityRequestCheckCommandConfigItemList {
	s.Value = &v
	return s
}

func (s *CheckDataSourceConnectivityRequestCheckCommandConfigItemList) Validate() error {
	return dara.Validate(s)
}
