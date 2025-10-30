// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBizEntityInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizEntityInfo(v *GetBizEntityInfoResponseBodyBizEntityInfo) *GetBizEntityInfoResponseBody
	GetBizEntityInfo() *GetBizEntityInfoResponseBodyBizEntityInfo
	SetCode(v string) *GetBizEntityInfoResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetBizEntityInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetBizEntityInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBizEntityInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBizEntityInfoResponseBody
	GetSuccess() *bool
}

type GetBizEntityInfoResponseBody struct {
	BizEntityInfo *GetBizEntityInfoResponseBodyBizEntityInfo `json:"BizEntityInfo,omitempty" xml:"BizEntityInfo,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBizEntityInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBizEntityInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBizEntityInfoResponseBody) GetBizEntityInfo() *GetBizEntityInfoResponseBodyBizEntityInfo {
	return s.BizEntityInfo
}

func (s *GetBizEntityInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBizEntityInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBizEntityInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBizEntityInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBizEntityInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBizEntityInfoResponseBody) SetBizEntityInfo(v *GetBizEntityInfoResponseBodyBizEntityInfo) *GetBizEntityInfoResponseBody {
	s.BizEntityInfo = v
	return s
}

func (s *GetBizEntityInfoResponseBody) SetCode(v string) *GetBizEntityInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetBizEntityInfoResponseBody) SetHttpStatusCode(v int32) *GetBizEntityInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBizEntityInfoResponseBody) SetMessage(v string) *GetBizEntityInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetBizEntityInfoResponseBody) SetRequestId(v string) *GetBizEntityInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBizEntityInfoResponseBody) SetSuccess(v bool) *GetBizEntityInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetBizEntityInfoResponseBody) Validate() error {
	if s.BizEntityInfo != nil {
		if err := s.BizEntityInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBizEntityInfoResponseBodyBizEntityInfo struct {
	BizObject  *GetBizEntityInfoResponseBodyBizEntityInfoBizObject  `json:"BizObject,omitempty" xml:"BizObject,omitempty" type:"Struct"`
	BizProcess *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess `json:"BizProcess,omitempty" xml:"BizProcess,omitempty" type:"Struct"`
	// example:
	//
	// 6798087749072704
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// example:
	//
	// 20101011
	DataDomainId *int64 `json:"DataDomainId,omitempty" xml:"DataDomainId,omitempty"`
	// example:
	//
	// BIZ_OBJECT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetBizEntityInfoResponseBodyBizEntityInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBizEntityInfoResponseBodyBizEntityInfo) GoString() string {
	return s.String()
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfo) GetBizObject() *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	return s.BizObject
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfo) GetBizProcess() *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	return s.BizProcess
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfo) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfo) GetDataDomainId() *int64 {
	return s.DataDomainId
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfo) GetType() *string {
	return s.Type
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfo) SetBizObject(v *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) *GetBizEntityInfoResponseBodyBizEntityInfo {
	s.BizObject = v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfo) SetBizProcess(v *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) *GetBizEntityInfoResponseBodyBizEntityInfo {
	s.BizProcess = v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfo) SetBizUnitId(v int64) *GetBizEntityInfoResponseBodyBizEntityInfo {
	s.BizUnitId = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfo) SetDataDomainId(v int64) *GetBizEntityInfoResponseBodyBizEntityInfo {
	s.DataDomainId = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfo) SetType(v string) *GetBizEntityInfoResponseBodyBizEntityInfo {
	s.Type = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfo) Validate() error {
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

type GetBizEntityInfoResponseBodyBizEntityInfoBizObject struct {
	// example:
	//
	// 221323121212
	ApprovalId *string `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	// example:
	//
	// APPROVING
	ApprovalStatus       *string  `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	ChildBizEntityIdList []*int64 `json:"ChildBizEntityIdList,omitempty" xml:"ChildBizEntityIdList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// create_object_name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2024-10-10 10:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-10-10 10:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1011
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 30010010
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// 张三
	LastModifierName *string `json:"LastModifierName,omitempty" xml:"LastModifierName,omitempty"`
	// example:
	//
	// create_object_code_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// SUBMITTED
	OnlineStatus *string `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 30010010
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// example:
	//
	// 116306
	ParentId           *int64   `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	RefBizEntityIdList []*int64 `json:"RefBizEntityIdList,omitempty" xml:"RefBizEntityIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	RefDimTableCount *int32 `json:"RefDimTableCount,omitempty" xml:"RefDimTableCount,omitempty"`
	// example:
	//
	// 1
	RefSummaryTableCount *int32 `json:"RefSummaryTableCount,omitempty" xml:"RefSummaryTableCount,omitempty"`
	// example:
	//
	// SUBMITTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// NORMAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetBizEntityInfoResponseBodyBizEntityInfoBizObject) String() string {
	return dara.Prettify(s)
}

func (s GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GoString() string {
	return s.String()
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetChildBizEntityIdList() []*int64 {
	return s.ChildBizEntityIdList
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetDescription() *string {
	return s.Description
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetId() *int64 {
	return s.Id
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetLastModifierName() *string {
	return s.LastModifierName
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetName() *string {
	return s.Name
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetOnlineStatus() *string {
	return s.OnlineStatus
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetParentId() *int64 {
	return s.ParentId
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetRefBizEntityIdList() []*int64 {
	return s.RefBizEntityIdList
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetRefDimTableCount() *int32 {
	return s.RefDimTableCount
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetRefSummaryTableCount() *int32 {
	return s.RefSummaryTableCount
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetStatus() *string {
	return s.Status
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) GetType() *string {
	return s.Type
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetApprovalId(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.ApprovalId = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetApprovalStatus(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.ApprovalStatus = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetChildBizEntityIdList(v []*int64) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.ChildBizEntityIdList = v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetDescription(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.Description = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetDisplayName(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.DisplayName = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetGmtCreate(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.GmtCreate = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetGmtModified(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.GmtModified = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetId(v int64) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.Id = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetLastModifier(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.LastModifier = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetLastModifierName(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.LastModifierName = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetName(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.Name = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetOnlineStatus(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.OnlineStatus = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetOwnerName(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.OwnerName = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetOwnerUserId(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.OwnerUserId = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetParentId(v int64) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.ParentId = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetRefBizEntityIdList(v []*int64) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.RefBizEntityIdList = v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetRefDimTableCount(v int32) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.RefDimTableCount = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetRefSummaryTableCount(v int32) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.RefSummaryTableCount = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetStatus(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.Status = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) SetType(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizObject {
	s.Type = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizObject) Validate() error {
	return dara.Validate(s)
}

type GetBizEntityInfoResponseBodyBizEntityInfoBizProcess struct {
	// example:
	//
	// 221323121212
	ApprovalId *string `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	// example:
	//
	// APPROVING
	ApprovalStatus       *string  `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	BizEventEntityIdList []*int64 `json:"BizEventEntityIdList,omitempty" xml:"BizEventEntityIdList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 业务活动测试
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2024-10-10 10:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-10-10 10:00:00
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HasDependent *bool   `json:"HasDependent,omitempty" xml:"HasDependent,omitempty"`
	// example:
	//
	// 1011
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 30010010
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// 张三
	LastModifierName *string `json:"LastModifierName,omitempty" xml:"LastModifierName,omitempty"`
	// example:
	//
	// create_process_code_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// SUBMITTED
	OnlineStatus *string `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 30010010
	OwnerUserId         *string  `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	PreBizProcessIdList []*int64 `json:"PreBizProcessIdList,omitempty" xml:"PreBizProcessIdList,omitempty" type:"Repeated"`
	RefBizEntityIdList  []*int64 `json:"RefBizEntityIdList,omitempty" xml:"RefBizEntityIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	RefFactTableCount *int32 `json:"RefFactTableCount,omitempty" xml:"RefFactTableCount,omitempty"`
	// example:
	//
	// SUBMITTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// BIZ_EVENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) String() string {
	return dara.Prettify(s)
}

func (s GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GoString() string {
	return s.String()
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetBizEventEntityIdList() []*int64 {
	return s.BizEventEntityIdList
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetDescription() *string {
	return s.Description
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetHasDependent() *bool {
	return s.HasDependent
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetId() *int64 {
	return s.Id
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetLastModifierName() *string {
	return s.LastModifierName
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetName() *string {
	return s.Name
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetOnlineStatus() *string {
	return s.OnlineStatus
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetPreBizProcessIdList() []*int64 {
	return s.PreBizProcessIdList
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetRefBizEntityIdList() []*int64 {
	return s.RefBizEntityIdList
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetRefFactTableCount() *int32 {
	return s.RefFactTableCount
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetStatus() *string {
	return s.Status
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) GetType() *string {
	return s.Type
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetApprovalId(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.ApprovalId = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetApprovalStatus(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.ApprovalStatus = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetBizEventEntityIdList(v []*int64) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.BizEventEntityIdList = v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetDescription(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.Description = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetDisplayName(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.DisplayName = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetGmtCreate(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.GmtCreate = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetGmtModified(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.GmtModified = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetHasDependent(v bool) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.HasDependent = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetId(v int64) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.Id = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetLastModifier(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.LastModifier = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetLastModifierName(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.LastModifierName = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetName(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.Name = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetOnlineStatus(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.OnlineStatus = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetOwnerName(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.OwnerName = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetOwnerUserId(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.OwnerUserId = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetPreBizProcessIdList(v []*int64) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.PreBizProcessIdList = v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetRefBizEntityIdList(v []*int64) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.RefBizEntityIdList = v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetRefFactTableCount(v int32) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.RefFactTableCount = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetStatus(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.Status = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) SetType(v string) *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess {
	s.Type = &v
	return s
}

func (s *GetBizEntityInfoResponseBodyBizEntityInfoBizProcess) Validate() error {
	return dara.Validate(s)
}
