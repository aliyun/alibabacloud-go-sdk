// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBizEntityInfoByVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizEntityInfo(v *GetBizEntityInfoByVersionResponseBodyBizEntityInfo) *GetBizEntityInfoByVersionResponseBody
	GetBizEntityInfo() *GetBizEntityInfoByVersionResponseBodyBizEntityInfo
	SetCode(v string) *GetBizEntityInfoByVersionResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetBizEntityInfoByVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetBizEntityInfoByVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBizEntityInfoByVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBizEntityInfoByVersionResponseBody
	GetSuccess() *bool
}

type GetBizEntityInfoByVersionResponseBody struct {
	// The creation result.
	BizEntityInfo *GetBizEntityInfoByVersionResponseBodyBizEntityInfo `json:"BizEntityInfo,omitempty" xml:"BizEntityInfo,omitempty" type:"Struct"`
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBizEntityInfoByVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBizEntityInfoByVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetBizEntityInfoByVersionResponseBody) GetBizEntityInfo() *GetBizEntityInfoByVersionResponseBodyBizEntityInfo {
	return s.BizEntityInfo
}

func (s *GetBizEntityInfoByVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBizEntityInfoByVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBizEntityInfoByVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBizEntityInfoByVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBizEntityInfoByVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBizEntityInfoByVersionResponseBody) SetBizEntityInfo(v *GetBizEntityInfoByVersionResponseBodyBizEntityInfo) *GetBizEntityInfoByVersionResponseBody {
	s.BizEntityInfo = v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBody) SetCode(v string) *GetBizEntityInfoByVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBody) SetHttpStatusCode(v int32) *GetBizEntityInfoByVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBody) SetMessage(v string) *GetBizEntityInfoByVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBody) SetRequestId(v string) *GetBizEntityInfoByVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBody) SetSuccess(v bool) *GetBizEntityInfoByVersionResponseBody {
	s.Success = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBody) Validate() error {
	if s.BizEntityInfo != nil {
		if err := s.BizEntityInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBizEntityInfoByVersionResponseBodyBizEntityInfo struct {
	// The business object.
	BizObject *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject `json:"BizObject,omitempty" xml:"BizObject,omitempty" type:"Struct"`
	// The business process.
	BizProcess *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess `json:"BizProcess,omitempty" xml:"BizProcess,omitempty" type:"Struct"`
	// The business unit ID of the business object.
	//
	// example:
	//
	// 6798087749072704
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The ID of the subject domain to which the entity belongs.
	//
	// example:
	//
	// 20101011
	DataDomainId *int64 `json:"DataDomainId,omitempty" xml:"DataDomainId,omitempty"`
	// The type of the business entity. For more information, refer to the query business entity details operation.
	//
	// example:
	//
	// BIZ_OBJECT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetBizEntityInfoByVersionResponseBodyBizEntityInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBizEntityInfoByVersionResponseBodyBizEntityInfo) GoString() string {
	return s.String()
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfo) GetBizObject() *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	return s.BizObject
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfo) GetBizProcess() *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	return s.BizProcess
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfo) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfo) GetDataDomainId() *int64 {
	return s.DataDomainId
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfo) GetType() *string {
	return s.Type
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfo) SetBizObject(v *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) *GetBizEntityInfoByVersionResponseBodyBizEntityInfo {
	s.BizObject = v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfo) SetBizProcess(v *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) *GetBizEntityInfoByVersionResponseBodyBizEntityInfo {
	s.BizProcess = v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfo) SetBizUnitId(v int64) *GetBizEntityInfoByVersionResponseBodyBizEntityInfo {
	s.BizUnitId = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfo) SetDataDomainId(v int64) *GetBizEntityInfoByVersionResponseBodyBizEntityInfo {
	s.DataDomainId = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfo) SetType(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfo {
	s.Type = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfo) Validate() error {
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

type GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject struct {
	// The approval flow ID.
	//
	// example:
	//
	// 221323121212
	ApprovalId *string `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	// The approval status. For more information, refer to the query business entity details operation.
	//
	// example:
	//
	// APPROVING
	ApprovalStatus *string `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	// The list of child business objects of the business object.
	ChildBizEntityIdList []*int64 `json:"ChildBizEntityIdList,omitempty" xml:"ChildBizEntityIdList,omitempty" type:"Repeated"`
	// The description of the business object.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the business object.
	//
	// example:
	//
	// create_object_name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-10-10 10:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2024-10-10 10:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The business object ID.
	//
	// example:
	//
	// 1011
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the user who last modified the business object.
	//
	// example:
	//
	// 30010010
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// The name of the user who last modified the business object.
	//
	// example:
	//
	// 张三
	LastModifierName *string `json:"LastModifierName,omitempty" xml:"LastModifierName,omitempty"`
	// The code name of the business object.
	//
	// example:
	//
	// create_object_code_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The online status. For more information, refer to the query business entity details operation.
	//
	// example:
	//
	// 1
	OnlineStatus *string `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// The owner of the business object.
	//
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The owner of the business object.
	//
	// example:
	//
	// 30010010
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The parent entity that the business object inherits from. Only common business objects support inheritance, and the parent entity must be online.
	//
	// example:
	//
	// 116306
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The list of associated online business entity IDs.
	RefBizEntityIdList []*int64 `json:"RefBizEntityIdList,omitempty" xml:"RefBizEntityIdList,omitempty" type:"Repeated"`
	// The number of associated logical dimension tables.
	//
	// example:
	//
	// 1
	RefDimTableCount *int32 `json:"RefDimTableCount,omitempty" xml:"RefDimTableCount,omitempty"`
	// The number of associated aggregate tables.
	//
	// example:
	//
	// 1
	RefSummaryTableCount *int32 `json:"RefSummaryTableCount,omitempty" xml:"RefSummaryTableCount,omitempty"`
	// The status. For more information, refer to the query business entity details operation.
	//
	// example:
	//
	// 100
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The object type of the business object. For more information, refer to the query business entity details operation.
	//
	// example:
	//
	// NORMAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) String() string {
	return dara.Prettify(s)
}

func (s GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GoString() string {
	return s.String()
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetChildBizEntityIdList() []*int64 {
	return s.ChildBizEntityIdList
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetDescription() *string {
	return s.Description
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetId() *int64 {
	return s.Id
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetLastModifierName() *string {
	return s.LastModifierName
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetName() *string {
	return s.Name
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetOnlineStatus() *string {
	return s.OnlineStatus
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetParentId() *int64 {
	return s.ParentId
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetRefBizEntityIdList() []*int64 {
	return s.RefBizEntityIdList
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetRefDimTableCount() *int32 {
	return s.RefDimTableCount
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetRefSummaryTableCount() *int32 {
	return s.RefSummaryTableCount
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetStatus() *string {
	return s.Status
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) GetType() *string {
	return s.Type
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetApprovalId(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.ApprovalId = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetApprovalStatus(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.ApprovalStatus = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetChildBizEntityIdList(v []*int64) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.ChildBizEntityIdList = v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetDescription(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.Description = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetDisplayName(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.DisplayName = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetGmtCreate(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.GmtCreate = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetGmtModified(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.GmtModified = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetId(v int64) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.Id = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetLastModifier(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.LastModifier = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetLastModifierName(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.LastModifierName = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetName(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.Name = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetOnlineStatus(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.OnlineStatus = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetOwnerName(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.OwnerName = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetOwnerUserId(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.OwnerUserId = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetParentId(v int64) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.ParentId = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetRefBizEntityIdList(v []*int64) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.RefBizEntityIdList = v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetRefDimTableCount(v int32) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.RefDimTableCount = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetRefSummaryTableCount(v int32) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.RefSummaryTableCount = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetStatus(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.Status = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) SetType(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject {
	s.Type = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizObject) Validate() error {
	return dara.Validate(s)
}

type GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess struct {
	// The approval flow ID.
	//
	// example:
	//
	// 221323121212
	ApprovalId *string `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	// The approval status. For more information, refer to the query business entity details operation.
	//
	// example:
	//
	// APPROVING
	ApprovalStatus *string `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	// The list of business event activity IDs contained in the business flow activity. This parameter is valid only when the type is a business flow activity.
	BizEventEntityIdList []*int64 `json:"BizEventEntityIdList,omitempty" xml:"BizEventEntityIdList,omitempty" type:"Repeated"`
	// The description of the business process.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the business process.
	//
	// example:
	//
	// 业务活动测试
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-10-10 10:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2024-10-10 10:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the business process is depended on by other entities.
	HasDependent *bool `json:"HasDependent,omitempty" xml:"HasDependent,omitempty"`
	// The business process ID.
	//
	// example:
	//
	// 1011
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the user who last modified the business object.
	//
	// example:
	//
	// 30010010
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// The name of the user who last modified the business object.
	//
	// example:
	//
	// 张三
	LastModifierName *string `json:"LastModifierName,omitempty" xml:"LastModifierName,omitempty"`
	// The name of the business process.
	//
	// example:
	//
	// create_process_code_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The online status. For more information, refer to the query business entity details operation.
	//
	// example:
	//
	// 1
	OnlineStatus *string `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// The owner of the business object.
	//
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The owner of the business object.
	//
	// example:
	//
	// 30010010
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The preceding business flow activities of the business flow activity.
	PreBizProcessIdList []*int64 `json:"PreBizProcessIdList,omitempty" xml:"PreBizProcessIdList,omitempty" type:"Repeated"`
	// The list of associated online business entity IDs.
	RefBizEntityIdList []*int64 `json:"RefBizEntityIdList,omitempty" xml:"RefBizEntityIdList,omitempty" type:"Repeated"`
	// The number of associated fact tables.
	//
	// example:
	//
	// 1
	RefFactTableCount *int32 `json:"RefFactTableCount,omitempty" xml:"RefFactTableCount,omitempty"`
	// The status. For more information, refer to the query business entity details operation.
	//
	// example:
	//
	// 100
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the business process. For more information, refer to the query business entity details operation.
	//
	// example:
	//
	// BIZ_EVENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) String() string {
	return dara.Prettify(s)
}

func (s GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GoString() string {
	return s.String()
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetBizEventEntityIdList() []*int64 {
	return s.BizEventEntityIdList
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetDescription() *string {
	return s.Description
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetHasDependent() *bool {
	return s.HasDependent
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetId() *int64 {
	return s.Id
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetLastModifierName() *string {
	return s.LastModifierName
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetName() *string {
	return s.Name
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetOnlineStatus() *string {
	return s.OnlineStatus
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetPreBizProcessIdList() []*int64 {
	return s.PreBizProcessIdList
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetRefBizEntityIdList() []*int64 {
	return s.RefBizEntityIdList
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetRefFactTableCount() *int32 {
	return s.RefFactTableCount
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetStatus() *string {
	return s.Status
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) GetType() *string {
	return s.Type
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetApprovalId(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.ApprovalId = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetApprovalStatus(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.ApprovalStatus = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetBizEventEntityIdList(v []*int64) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.BizEventEntityIdList = v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetDescription(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.Description = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetDisplayName(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.DisplayName = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetGmtCreate(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.GmtCreate = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetGmtModified(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.GmtModified = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetHasDependent(v bool) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.HasDependent = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetId(v int64) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.Id = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetLastModifier(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.LastModifier = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetLastModifierName(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.LastModifierName = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetName(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.Name = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetOnlineStatus(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.OnlineStatus = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetOwnerName(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.OwnerName = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetOwnerUserId(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.OwnerUserId = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetPreBizProcessIdList(v []*int64) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.PreBizProcessIdList = v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetRefBizEntityIdList(v []*int64) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.RefBizEntityIdList = v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetRefFactTableCount(v int32) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.RefFactTableCount = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetStatus(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.Status = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) SetType(v string) *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess {
	s.Type = &v
	return s
}

func (s *GetBizEntityInfoByVersionResponseBodyBizEntityInfoBizProcess) Validate() error {
	return dara.Validate(s)
}
