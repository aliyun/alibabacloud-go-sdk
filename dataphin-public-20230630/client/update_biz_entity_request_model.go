// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateBizEntityRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateBizEntityRequestUpdateCommand) *UpdateBizEntityRequest
	GetUpdateCommand() *UpdateBizEntityRequestUpdateCommand
}

type UpdateBizEntityRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The update request.
	//
	// This parameter is required.
	UpdateCommand *UpdateBizEntityRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateBizEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizEntityRequest) GoString() string {
	return s.String()
}

func (s *UpdateBizEntityRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateBizEntityRequest) GetUpdateCommand() *UpdateBizEntityRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateBizEntityRequest) SetOpTenantId(v int64) *UpdateBizEntityRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateBizEntityRequest) SetUpdateCommand(v *UpdateBizEntityRequestUpdateCommand) *UpdateBizEntityRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateBizEntityRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateBizEntityRequestUpdateCommand struct {
	// The business object.
	BizObject *UpdateBizEntityRequestUpdateCommandBizObject `json:"BizObject,omitempty" xml:"BizObject,omitempty" type:"Struct"`
	// The business process.
	BizProcess *UpdateBizEntityRequestUpdateCommandBizProcess `json:"BizProcess,omitempty" xml:"BizProcess,omitempty" type:"Struct"`
	// The ID of the business unit to which the business process belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6798087749072704
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The ID of the data domain to which the business process belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20101011
	DataDomainId *int64 `json:"DataDomainId,omitempty" xml:"DataDomainId,omitempty"`
	// The ID of the business entity.
	//
	// This parameter is required.
	//
	// example:
	//
	// 101001201
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of the business entity. For more information, refer to the create business entity operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// BIZ_OBJECT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateBizEntityRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizEntityRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateBizEntityRequestUpdateCommand) GetBizObject() *UpdateBizEntityRequestUpdateCommandBizObject {
	return s.BizObject
}

func (s *UpdateBizEntityRequestUpdateCommand) GetBizProcess() *UpdateBizEntityRequestUpdateCommandBizProcess {
	return s.BizProcess
}

func (s *UpdateBizEntityRequestUpdateCommand) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *UpdateBizEntityRequestUpdateCommand) GetDataDomainId() *int64 {
	return s.DataDomainId
}

func (s *UpdateBizEntityRequestUpdateCommand) GetId() *int64 {
	return s.Id
}

func (s *UpdateBizEntityRequestUpdateCommand) GetType() *string {
	return s.Type
}

func (s *UpdateBizEntityRequestUpdateCommand) SetBizObject(v *UpdateBizEntityRequestUpdateCommandBizObject) *UpdateBizEntityRequestUpdateCommand {
	s.BizObject = v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommand) SetBizProcess(v *UpdateBizEntityRequestUpdateCommandBizProcess) *UpdateBizEntityRequestUpdateCommand {
	s.BizProcess = v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommand) SetBizUnitId(v int64) *UpdateBizEntityRequestUpdateCommand {
	s.BizUnitId = &v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommand) SetDataDomainId(v int64) *UpdateBizEntityRequestUpdateCommand {
	s.DataDomainId = &v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommand) SetId(v int64) *UpdateBizEntityRequestUpdateCommand {
	s.Id = &v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommand) SetType(v string) *UpdateBizEntityRequestUpdateCommand {
	s.Type = &v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommand) Validate() error {
	if s.BizObject != nil {
		if err := s.BizObject.Validate(); err != nil {
			return err
		}
	}
	if s.BizProcess != nil {
		if err := s.BizProcess.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateBizEntityRequestUpdateCommandBizObject struct {
	// The description of the business object. The description can be up to 128 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the business object. The name can be up to 64 characters in length and can contain only Chinese characters, letters, digits, underscores, and hyphens.
	//
	// This parameter is required.
	//
	// example:
	//
	// create_object_name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The code name of the business object. The name can be up to 64 characters in length and can contain only letters, digits, and underscores. For the ADB_PG engine, the code name can be up to 40 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// create_object_code_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user ID of the business object owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30010010
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The inherited entity of the business object. Only common business objects support inherited entity configuration, and only online business objects can be inherited.
	//
	// example:
	//
	// 116306
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The list of associated online business entity IDs. If this parameter is not specified, the existing values are cleared.
	RefBizEntityIdList []*int64 `json:"RefBizEntityIdList,omitempty" xml:"RefBizEntityIdList,omitempty" type:"Repeated"`
}

func (s UpdateBizEntityRequestUpdateCommandBizObject) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizEntityRequestUpdateCommandBizObject) GoString() string {
	return s.String()
}

func (s *UpdateBizEntityRequestUpdateCommandBizObject) GetDescription() *string {
	return s.Description
}

func (s *UpdateBizEntityRequestUpdateCommandBizObject) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateBizEntityRequestUpdateCommandBizObject) GetName() *string {
	return s.Name
}

func (s *UpdateBizEntityRequestUpdateCommandBizObject) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *UpdateBizEntityRequestUpdateCommandBizObject) GetParentId() *int64 {
	return s.ParentId
}

func (s *UpdateBizEntityRequestUpdateCommandBizObject) GetRefBizEntityIdList() []*int64 {
	return s.RefBizEntityIdList
}

func (s *UpdateBizEntityRequestUpdateCommandBizObject) SetDescription(v string) *UpdateBizEntityRequestUpdateCommandBizObject {
	s.Description = &v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommandBizObject) SetDisplayName(v string) *UpdateBizEntityRequestUpdateCommandBizObject {
	s.DisplayName = &v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommandBizObject) SetName(v string) *UpdateBizEntityRequestUpdateCommandBizObject {
	s.Name = &v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommandBizObject) SetOwnerUserId(v string) *UpdateBizEntityRequestUpdateCommandBizObject {
	s.OwnerUserId = &v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommandBizObject) SetParentId(v int64) *UpdateBizEntityRequestUpdateCommandBizObject {
	s.ParentId = &v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommandBizObject) SetRefBizEntityIdList(v []*int64) *UpdateBizEntityRequestUpdateCommandBizObject {
	s.RefBizEntityIdList = v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommandBizObject) Validate() error {
	return dara.Validate(s)
}

type UpdateBizEntityRequestUpdateCommandBizProcess struct {
	// The list of business event activity IDs contained in the business flow activity. This parameter is valid only when the current entity is a business flow activity.
	BizEventEntityIdList []*int64 `json:"BizEventEntityIdList,omitempty" xml:"BizEventEntityIdList,omitempty" type:"Repeated"`
	// The description of the business process. The description can be up to 128 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the business process. The name can be up to 64 characters in length and can contain only Chinese characters, letters, digits, underscores, and hyphens.
	//
	// This parameter is required.
	//
	// example:
	//
	// create_process_name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The code name of the business process. The name can be up to 64 characters in length and can contain only letters, digits, and underscores. For the ADB_PG engine, the code name can be up to 40 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// create_process_code_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user ID of the business process owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30010010
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The preceding business flow activities of the business flow activity.
	PreBizProcessIdList []*int64 `json:"PreBizProcessIdList,omitempty" xml:"PreBizProcessIdList,omitempty" type:"Repeated"`
	// The list of associated online business entity IDs. If this parameter is not specified, the existing values are cleared.
	RefBizEntityIdList []*int64 `json:"RefBizEntityIdList,omitempty" xml:"RefBizEntityIdList,omitempty" type:"Repeated"`
}

func (s UpdateBizEntityRequestUpdateCommandBizProcess) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizEntityRequestUpdateCommandBizProcess) GoString() string {
	return s.String()
}

func (s *UpdateBizEntityRequestUpdateCommandBizProcess) GetBizEventEntityIdList() []*int64 {
	return s.BizEventEntityIdList
}

func (s *UpdateBizEntityRequestUpdateCommandBizProcess) GetDescription() *string {
	return s.Description
}

func (s *UpdateBizEntityRequestUpdateCommandBizProcess) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateBizEntityRequestUpdateCommandBizProcess) GetName() *string {
	return s.Name
}

func (s *UpdateBizEntityRequestUpdateCommandBizProcess) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *UpdateBizEntityRequestUpdateCommandBizProcess) GetPreBizProcessIdList() []*int64 {
	return s.PreBizProcessIdList
}

func (s *UpdateBizEntityRequestUpdateCommandBizProcess) GetRefBizEntityIdList() []*int64 {
	return s.RefBizEntityIdList
}

func (s *UpdateBizEntityRequestUpdateCommandBizProcess) SetBizEventEntityIdList(v []*int64) *UpdateBizEntityRequestUpdateCommandBizProcess {
	s.BizEventEntityIdList = v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommandBizProcess) SetDescription(v string) *UpdateBizEntityRequestUpdateCommandBizProcess {
	s.Description = &v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommandBizProcess) SetDisplayName(v string) *UpdateBizEntityRequestUpdateCommandBizProcess {
	s.DisplayName = &v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommandBizProcess) SetName(v string) *UpdateBizEntityRequestUpdateCommandBizProcess {
	s.Name = &v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommandBizProcess) SetOwnerUserId(v string) *UpdateBizEntityRequestUpdateCommandBizProcess {
	s.OwnerUserId = &v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommandBizProcess) SetPreBizProcessIdList(v []*int64) *UpdateBizEntityRequestUpdateCommandBizProcess {
	s.PreBizProcessIdList = v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommandBizProcess) SetRefBizEntityIdList(v []*int64) *UpdateBizEntityRequestUpdateCommandBizProcess {
	s.RefBizEntityIdList = v
	return s
}

func (s *UpdateBizEntityRequestUpdateCommandBizProcess) Validate() error {
	return dara.Validate(s)
}
