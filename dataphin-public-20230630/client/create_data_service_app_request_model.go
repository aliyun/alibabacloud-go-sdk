// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateDataServiceAppRequestCreateCommand) *CreateDataServiceAppRequest
	GetCreateCommand() *CreateDataServiceAppRequestCreateCommand
	SetOpTenantId(v int64) *CreateDataServiceAppRequest
	GetOpTenantId() *int64
}

type CreateDataServiceAppRequest struct {
	// This parameter is required.
	CreateCommand *CreateDataServiceAppRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateDataServiceAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceAppRequest) GoString() string {
	return s.String()
}

func (s *CreateDataServiceAppRequest) GetCreateCommand() *CreateDataServiceAppRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateDataServiceAppRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateDataServiceAppRequest) SetCreateCommand(v *CreateDataServiceAppRequestCreateCommand) *CreateDataServiceAppRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateDataServiceAppRequest) SetOpTenantId(v int64) *CreateDataServiceAppRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDataServiceAppRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataServiceAppRequestCreateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 200000000
	AppGroupId *int32 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	// example:
	//
	// 200000001
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 默认应用
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// abc123456789
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// This parameter is required.
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 数据分析
	Scenarios *string `json:"Scenarios,omitempty" xml:"Scenarios,omitempty"`
}

func (s CreateDataServiceAppRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceAppRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateDataServiceAppRequestCreateCommand) GetAppGroupId() *int32 {
	return s.AppGroupId
}

func (s *CreateDataServiceAppRequestCreateCommand) GetAppKey() *string {
	return s.AppKey
}

func (s *CreateDataServiceAppRequestCreateCommand) GetAppName() *string {
	return s.AppName
}

func (s *CreateDataServiceAppRequestCreateCommand) GetAppSecret() *string {
	return s.AppSecret
}

func (s *CreateDataServiceAppRequestCreateCommand) GetOwnerIds() []*string {
	return s.OwnerIds
}

func (s *CreateDataServiceAppRequestCreateCommand) GetScenarios() *string {
	return s.Scenarios
}

func (s *CreateDataServiceAppRequestCreateCommand) SetAppGroupId(v int32) *CreateDataServiceAppRequestCreateCommand {
	s.AppGroupId = &v
	return s
}

func (s *CreateDataServiceAppRequestCreateCommand) SetAppKey(v string) *CreateDataServiceAppRequestCreateCommand {
	s.AppKey = &v
	return s
}

func (s *CreateDataServiceAppRequestCreateCommand) SetAppName(v string) *CreateDataServiceAppRequestCreateCommand {
	s.AppName = &v
	return s
}

func (s *CreateDataServiceAppRequestCreateCommand) SetAppSecret(v string) *CreateDataServiceAppRequestCreateCommand {
	s.AppSecret = &v
	return s
}

func (s *CreateDataServiceAppRequestCreateCommand) SetOwnerIds(v []*string) *CreateDataServiceAppRequestCreateCommand {
	s.OwnerIds = v
	return s
}

func (s *CreateDataServiceAppRequestCreateCommand) SetScenarios(v string) *CreateDataServiceAppRequestCreateCommand {
	s.Scenarios = &v
	return s
}

func (s *CreateDataServiceAppRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}
