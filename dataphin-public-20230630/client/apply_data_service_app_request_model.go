// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyDataServiceAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyCommand(v *ApplyDataServiceAppRequestApplyCommand) *ApplyDataServiceAppRequest
	GetApplyCommand() *ApplyDataServiceAppRequestApplyCommand
	SetOpTenantId(v int64) *ApplyDataServiceAppRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ApplyDataServiceAppRequest
	GetProjectId() *int32
}

type ApplyDataServiceAppRequest struct {
	// This parameter is required.
	ApplyCommand *ApplyDataServiceAppRequestApplyCommand `json:"ApplyCommand,omitempty" xml:"ApplyCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102102
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ApplyDataServiceAppRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyDataServiceAppRequest) GoString() string {
	return s.String()
}

func (s *ApplyDataServiceAppRequest) GetApplyCommand() *ApplyDataServiceAppRequestApplyCommand {
	return s.ApplyCommand
}

func (s *ApplyDataServiceAppRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ApplyDataServiceAppRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ApplyDataServiceAppRequest) SetApplyCommand(v *ApplyDataServiceAppRequestApplyCommand) *ApplyDataServiceAppRequest {
	s.ApplyCommand = v
	return s
}

func (s *ApplyDataServiceAppRequest) SetOpTenantId(v int64) *ApplyDataServiceAppRequest {
	s.OpTenantId = &v
	return s
}

func (s *ApplyDataServiceAppRequest) SetProjectId(v int32) *ApplyDataServiceAppRequest {
	s.ProjectId = &v
	return s
}

func (s *ApplyDataServiceAppRequest) Validate() error {
	return dara.Validate(s)
}

type ApplyDataServiceAppRequestApplyCommand struct {
	// appId
	//
	// This parameter is required.
	//
	// example:
	//
	// 2011
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-30
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ApplyDataServiceAppRequestApplyCommand) String() string {
	return dara.Prettify(s)
}

func (s ApplyDataServiceAppRequestApplyCommand) GoString() string {
	return s.String()
}

func (s *ApplyDataServiceAppRequestApplyCommand) GetAppId() *int32 {
	return s.AppId
}

func (s *ApplyDataServiceAppRequestApplyCommand) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *ApplyDataServiceAppRequestApplyCommand) GetReason() *string {
	return s.Reason
}

func (s *ApplyDataServiceAppRequestApplyCommand) SetAppId(v int32) *ApplyDataServiceAppRequestApplyCommand {
	s.AppId = &v
	return s
}

func (s *ApplyDataServiceAppRequestApplyCommand) SetExpireDate(v string) *ApplyDataServiceAppRequestApplyCommand {
	s.ExpireDate = &v
	return s
}

func (s *ApplyDataServiceAppRequestApplyCommand) SetReason(v string) *ApplyDataServiceAppRequestApplyCommand {
	s.Reason = &v
	return s
}

func (s *ApplyDataServiceAppRequestApplyCommand) Validate() error {
	return dara.Validate(s)
}
