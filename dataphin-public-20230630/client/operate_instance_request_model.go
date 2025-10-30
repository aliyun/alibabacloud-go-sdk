// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *OperateInstanceRequest
	GetEnv() *string
	SetOpTenantId(v int64) *OperateInstanceRequest
	GetOpTenantId() *int64
	SetOperateCommand(v *OperateInstanceRequestOperateCommand) *OperateInstanceRequest
	GetOperateCommand() *OperateInstanceRequestOperateCommand
}

type OperateInstanceRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	OperateCommand *OperateInstanceRequestOperateCommand `json:"OperateCommand,omitempty" xml:"OperateCommand,omitempty" type:"Struct"`
}

func (s OperateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateInstanceRequest) GoString() string {
	return s.String()
}

func (s *OperateInstanceRequest) GetEnv() *string {
	return s.Env
}

func (s *OperateInstanceRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *OperateInstanceRequest) GetOperateCommand() *OperateInstanceRequestOperateCommand {
	return s.OperateCommand
}

func (s *OperateInstanceRequest) SetEnv(v string) *OperateInstanceRequest {
	s.Env = &v
	return s
}

func (s *OperateInstanceRequest) SetOpTenantId(v int64) *OperateInstanceRequest {
	s.OpTenantId = &v
	return s
}

func (s *OperateInstanceRequest) SetOperateCommand(v *OperateInstanceRequestOperateCommand) *OperateInstanceRequest {
	s.OperateCommand = v
	return s
}

func (s *OperateInstanceRequest) Validate() error {
	if s.OperateCommand != nil {
		if err := s.OperateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OperateInstanceRequestOperateCommand struct {
	// This parameter is required.
	InstanceIdList []*OperateInstanceRequestOperateCommandInstanceIdList `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// RERUN
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 132311
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s OperateInstanceRequestOperateCommand) String() string {
	return dara.Prettify(s)
}

func (s OperateInstanceRequestOperateCommand) GoString() string {
	return s.String()
}

func (s *OperateInstanceRequestOperateCommand) GetInstanceIdList() []*OperateInstanceRequestOperateCommandInstanceIdList {
	return s.InstanceIdList
}

func (s *OperateInstanceRequestOperateCommand) GetOperation() *string {
	return s.Operation
}

func (s *OperateInstanceRequestOperateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *OperateInstanceRequestOperateCommand) SetInstanceIdList(v []*OperateInstanceRequestOperateCommandInstanceIdList) *OperateInstanceRequestOperateCommand {
	s.InstanceIdList = v
	return s
}

func (s *OperateInstanceRequestOperateCommand) SetOperation(v string) *OperateInstanceRequestOperateCommand {
	s.Operation = &v
	return s
}

func (s *OperateInstanceRequestOperateCommand) SetProjectId(v int64) *OperateInstanceRequestOperateCommand {
	s.ProjectId = &v
	return s
}

func (s *OperateInstanceRequestOperateCommand) Validate() error {
	if s.InstanceIdList != nil {
		for _, item := range s.InstanceIdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OperateInstanceRequestOperateCommandInstanceIdList struct {
	FieldInstanceIdList []*string `json:"FieldInstanceIdList,omitempty" xml:"FieldInstanceIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// t_32111312
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s OperateInstanceRequestOperateCommandInstanceIdList) String() string {
	return dara.Prettify(s)
}

func (s OperateInstanceRequestOperateCommandInstanceIdList) GoString() string {
	return s.String()
}

func (s *OperateInstanceRequestOperateCommandInstanceIdList) GetFieldInstanceIdList() []*string {
	return s.FieldInstanceIdList
}

func (s *OperateInstanceRequestOperateCommandInstanceIdList) GetId() *string {
	return s.Id
}

func (s *OperateInstanceRequestOperateCommandInstanceIdList) SetFieldInstanceIdList(v []*string) *OperateInstanceRequestOperateCommandInstanceIdList {
	s.FieldInstanceIdList = v
	return s
}

func (s *OperateInstanceRequestOperateCommandInstanceIdList) SetId(v string) *OperateInstanceRequestOperateCommandInstanceIdList {
	s.Id = &v
	return s
}

func (s *OperateInstanceRequestOperateCommandInstanceIdList) Validate() error {
	return dara.Validate(s)
}
