// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBizEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateBizEntityRequestCreateCommand) *CreateBizEntityRequest
	GetCreateCommand() *CreateBizEntityRequestCreateCommand
	SetOpTenantId(v int64) *CreateBizEntityRequest
	GetOpTenantId() *int64
}

type CreateBizEntityRequest struct {
	// This parameter is required.
	CreateCommand *CreateBizEntityRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateBizEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBizEntityRequest) GoString() string {
	return s.String()
}

func (s *CreateBizEntityRequest) GetCreateCommand() *CreateBizEntityRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateBizEntityRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateBizEntityRequest) SetCreateCommand(v *CreateBizEntityRequestCreateCommand) *CreateBizEntityRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateBizEntityRequest) SetOpTenantId(v int64) *CreateBizEntityRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateBizEntityRequest) Validate() error {
	return dara.Validate(s)
}

type CreateBizEntityRequestCreateCommand struct {
	BizObject  *CreateBizEntityRequestCreateCommandBizObject  `json:"BizObject,omitempty" xml:"BizObject,omitempty" type:"Struct"`
	BizProcess *CreateBizEntityRequestCreateCommandBizProcess `json:"BizProcess,omitempty" xml:"BizProcess,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 6798087749072704
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20101011
	DataDomainId *int64 `json:"DataDomainId,omitempty" xml:"DataDomainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BIZ_OBJECT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateBizEntityRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateBizEntityRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateBizEntityRequestCreateCommand) GetBizObject() *CreateBizEntityRequestCreateCommandBizObject {
	return s.BizObject
}

func (s *CreateBizEntityRequestCreateCommand) GetBizProcess() *CreateBizEntityRequestCreateCommandBizProcess {
	return s.BizProcess
}

func (s *CreateBizEntityRequestCreateCommand) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *CreateBizEntityRequestCreateCommand) GetDataDomainId() *int64 {
	return s.DataDomainId
}

func (s *CreateBizEntityRequestCreateCommand) GetType() *string {
	return s.Type
}

func (s *CreateBizEntityRequestCreateCommand) SetBizObject(v *CreateBizEntityRequestCreateCommandBizObject) *CreateBizEntityRequestCreateCommand {
	s.BizObject = v
	return s
}

func (s *CreateBizEntityRequestCreateCommand) SetBizProcess(v *CreateBizEntityRequestCreateCommandBizProcess) *CreateBizEntityRequestCreateCommand {
	s.BizProcess = v
	return s
}

func (s *CreateBizEntityRequestCreateCommand) SetBizUnitId(v int64) *CreateBizEntityRequestCreateCommand {
	s.BizUnitId = &v
	return s
}

func (s *CreateBizEntityRequestCreateCommand) SetDataDomainId(v int64) *CreateBizEntityRequestCreateCommand {
	s.DataDomainId = &v
	return s
}

func (s *CreateBizEntityRequestCreateCommand) SetType(v string) *CreateBizEntityRequestCreateCommand {
	s.Type = &v
	return s
}

func (s *CreateBizEntityRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}

type CreateBizEntityRequestCreateCommandBizObject struct {
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// create_object_name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// create_object_code_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 30010010
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// example:
	//
	// 116306
	ParentId           *int64   `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	RefBizEntityIdList []*int64 `json:"RefBizEntityIdList,omitempty" xml:"RefBizEntityIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateBizEntityRequestCreateCommandBizObject) String() string {
	return dara.Prettify(s)
}

func (s CreateBizEntityRequestCreateCommandBizObject) GoString() string {
	return s.String()
}

func (s *CreateBizEntityRequestCreateCommandBizObject) GetDescription() *string {
	return s.Description
}

func (s *CreateBizEntityRequestCreateCommandBizObject) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateBizEntityRequestCreateCommandBizObject) GetName() *string {
	return s.Name
}

func (s *CreateBizEntityRequestCreateCommandBizObject) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *CreateBizEntityRequestCreateCommandBizObject) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateBizEntityRequestCreateCommandBizObject) GetRefBizEntityIdList() []*int64 {
	return s.RefBizEntityIdList
}

func (s *CreateBizEntityRequestCreateCommandBizObject) GetType() *string {
	return s.Type
}

func (s *CreateBizEntityRequestCreateCommandBizObject) SetDescription(v string) *CreateBizEntityRequestCreateCommandBizObject {
	s.Description = &v
	return s
}

func (s *CreateBizEntityRequestCreateCommandBizObject) SetDisplayName(v string) *CreateBizEntityRequestCreateCommandBizObject {
	s.DisplayName = &v
	return s
}

func (s *CreateBizEntityRequestCreateCommandBizObject) SetName(v string) *CreateBizEntityRequestCreateCommandBizObject {
	s.Name = &v
	return s
}

func (s *CreateBizEntityRequestCreateCommandBizObject) SetOwnerUserId(v string) *CreateBizEntityRequestCreateCommandBizObject {
	s.OwnerUserId = &v
	return s
}

func (s *CreateBizEntityRequestCreateCommandBizObject) SetParentId(v int64) *CreateBizEntityRequestCreateCommandBizObject {
	s.ParentId = &v
	return s
}

func (s *CreateBizEntityRequestCreateCommandBizObject) SetRefBizEntityIdList(v []*int64) *CreateBizEntityRequestCreateCommandBizObject {
	s.RefBizEntityIdList = v
	return s
}

func (s *CreateBizEntityRequestCreateCommandBizObject) SetType(v string) *CreateBizEntityRequestCreateCommandBizObject {
	s.Type = &v
	return s
}

func (s *CreateBizEntityRequestCreateCommandBizObject) Validate() error {
	return dara.Validate(s)
}

type CreateBizEntityRequestCreateCommandBizProcess struct {
	BizEventEntityIdList []*int64 `json:"BizEventEntityIdList,omitempty" xml:"BizEventEntityIdList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// create_process_name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// create_process_code_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 30010010
	OwnerUserId         *string  `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	PreBizProcessIdList []*int64 `json:"PreBizProcessIdList,omitempty" xml:"PreBizProcessIdList,omitempty" type:"Repeated"`
	RefBizEntityIdList  []*int64 `json:"RefBizEntityIdList,omitempty" xml:"RefBizEntityIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// BIZ_EVENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateBizEntityRequestCreateCommandBizProcess) String() string {
	return dara.Prettify(s)
}

func (s CreateBizEntityRequestCreateCommandBizProcess) GoString() string {
	return s.String()
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) GetBizEventEntityIdList() []*int64 {
	return s.BizEventEntityIdList
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) GetDescription() *string {
	return s.Description
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) GetName() *string {
	return s.Name
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) GetPreBizProcessIdList() []*int64 {
	return s.PreBizProcessIdList
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) GetRefBizEntityIdList() []*int64 {
	return s.RefBizEntityIdList
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) GetType() *string {
	return s.Type
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) SetBizEventEntityIdList(v []*int64) *CreateBizEntityRequestCreateCommandBizProcess {
	s.BizEventEntityIdList = v
	return s
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) SetDescription(v string) *CreateBizEntityRequestCreateCommandBizProcess {
	s.Description = &v
	return s
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) SetDisplayName(v string) *CreateBizEntityRequestCreateCommandBizProcess {
	s.DisplayName = &v
	return s
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) SetName(v string) *CreateBizEntityRequestCreateCommandBizProcess {
	s.Name = &v
	return s
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) SetOwnerUserId(v string) *CreateBizEntityRequestCreateCommandBizProcess {
	s.OwnerUserId = &v
	return s
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) SetPreBizProcessIdList(v []*int64) *CreateBizEntityRequestCreateCommandBizProcess {
	s.PreBizProcessIdList = v
	return s
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) SetRefBizEntityIdList(v []*int64) *CreateBizEntityRequestCreateCommandBizProcess {
	s.RefBizEntityIdList = v
	return s
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) SetType(v string) *CreateBizEntityRequestCreateCommandBizProcess {
	s.Type = &v
	return s
}

func (s *CreateBizEntityRequestCreateCommandBizProcess) Validate() error {
	return dara.Validate(s)
}
