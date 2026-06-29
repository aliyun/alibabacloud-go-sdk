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
	// The create request.
	//
	// This parameter is required.
	CreateCommand *CreateBizEntityRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
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
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBizEntityRequestCreateCommand struct {
	// The business object.
	BizObject *CreateBizEntityRequestCreateCommandBizObject `json:"BizObject,omitempty" xml:"BizObject,omitempty" type:"Struct"`
	// The business activity.
	BizProcess *CreateBizEntityRequestCreateCommandBizProcess `json:"BizProcess,omitempty" xml:"BizProcess,omitempty" type:"Struct"`
	// The ID of the business unit to which the business activity belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6798087749072704
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The ID of the data domain to which the business activity belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20101011
	DataDomainId *int64 `json:"DataDomainId,omitempty" xml:"DataDomainId,omitempty"`
	// The business type. Valid values:
	//
	// - BIZ_OBJECT: business object.
	//
	// - BIZ_PROCESS: business activity.
	//
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

type CreateBizEntityRequestCreateCommandBizObject struct {
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
	// The code name of the business object. The name can be up to 64 characters in length and can contain only letters, digits, and underscores. For ADB_PG engines, the code name can be up to 40 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// create_object_code_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user ID of the business object owner.
	//
	// example:
	//
	// 30010010
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The parent entity from which the business object inherits. Only common business objects support inheritance, and the parent entity must be an online business object.
	//
	// example:
	//
	// 116306
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The list of associated online business entity IDs.
	RefBizEntityIdList []*int64 `json:"RefBizEntityIdList,omitempty" xml:"RefBizEntityIdList,omitempty" type:"Repeated"`
	// The object type of the business object. Valid values:
	//
	// - NORMAL: common object.
	//
	// - ENUM: enumeration object.
	//
	// - VIRTUAL: virtual object.
	//
	// - HIERARCHY: hierarchy object.
	//
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
	// The list of business event activity IDs included in the business process activity. This parameter takes effect only when the current activity is a business process activity.
	BizEventEntityIdList []*int64 `json:"BizEventEntityIdList,omitempty" xml:"BizEventEntityIdList,omitempty" type:"Repeated"`
	// The description of the business activity. The description can be up to 128 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the business activity. The name can be up to 64 characters in length and can contain only Chinese characters, letters, digits, underscores, and hyphens.
	//
	// This parameter is required.
	//
	// example:
	//
	// create_process_name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The code name of the business activity. The name can be up to 64 characters in length and can contain only letters, digits, and underscores. For ADB_PG engines, the code name can be up to 40 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// create_process_code_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user ID of the business activity owner.
	//
	// example:
	//
	// 30010010
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The list of preceding business process activity IDs for the business process activity.
	PreBizProcessIdList []*int64 `json:"PreBizProcessIdList,omitempty" xml:"PreBizProcessIdList,omitempty" type:"Repeated"`
	// The list of associated online business entity IDs.
	RefBizEntityIdList []*int64 `json:"RefBizEntityIdList,omitempty" xml:"RefBizEntityIdList,omitempty" type:"Repeated"`
	// The type of the business activity. Valid values:
	//
	// - BIZ_EVENT: business event.
	//
	// - BIZ_SNAPSHOT: business snapshot.
	//
	// - BIZ_PROCESS: business process.
	//
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
