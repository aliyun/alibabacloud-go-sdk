// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBizEntitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBizEntitiesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListBizEntitiesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBizEntitiesResponseBody
	GetMessage() *string
	SetPageResult(v *ListBizEntitiesResponseBodyPageResult) *ListBizEntitiesResponseBody
	GetPageResult() *ListBizEntitiesResponseBodyPageResult
	SetRequestId(v string) *ListBizEntitiesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBizEntitiesResponseBody
	GetSuccess() *bool
}

type ListBizEntitiesResponseBody struct {
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
	// The paginated result.
	PageResult *ListBizEntitiesResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBizEntitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBizEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBizEntitiesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBizEntitiesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBizEntitiesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBizEntitiesResponseBody) GetPageResult() *ListBizEntitiesResponseBodyPageResult {
	return s.PageResult
}

func (s *ListBizEntitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBizEntitiesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBizEntitiesResponseBody) SetCode(v string) *ListBizEntitiesResponseBody {
	s.Code = &v
	return s
}

func (s *ListBizEntitiesResponseBody) SetHttpStatusCode(v int32) *ListBizEntitiesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBizEntitiesResponseBody) SetMessage(v string) *ListBizEntitiesResponseBody {
	s.Message = &v
	return s
}

func (s *ListBizEntitiesResponseBody) SetPageResult(v *ListBizEntitiesResponseBodyPageResult) *ListBizEntitiesResponseBody {
	s.PageResult = v
	return s
}

func (s *ListBizEntitiesResponseBody) SetRequestId(v string) *ListBizEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBizEntitiesResponseBody) SetSuccess(v bool) *ListBizEntitiesResponseBody {
	s.Success = &v
	return s
}

func (s *ListBizEntitiesResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBizEntitiesResponseBodyPageResult struct {
	// The list of business entities.
	BizEntityList []*ListBizEntitiesResponseBodyPageResultBizEntityList `json:"BizEntityList,omitempty" xml:"BizEntityList,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 66
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBizEntitiesResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListBizEntitiesResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListBizEntitiesResponseBodyPageResult) GetBizEntityList() []*ListBizEntitiesResponseBodyPageResultBizEntityList {
	return s.BizEntityList
}

func (s *ListBizEntitiesResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBizEntitiesResponseBodyPageResult) SetBizEntityList(v []*ListBizEntitiesResponseBodyPageResultBizEntityList) *ListBizEntitiesResponseBodyPageResult {
	s.BizEntityList = v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResult) SetTotalCount(v int32) *ListBizEntitiesResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResult) Validate() error {
	if s.BizEntityList != nil {
		for _, item := range s.BizEntityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBizEntitiesResponseBodyPageResultBizEntityList struct {
	// The list of parent business process IDs. This parameter is valid only when the entity type is business process.
	BelongToBizEntityIdList []*int64 `json:"BelongToBizEntityIdList,omitempty" xml:"BelongToBizEntityIdList,omitempty" type:"Repeated"`
	// The ID of the business unit to which the business process belongs.
	//
	// example:
	//
	// 6798087749072704
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The list of child business object IDs. This parameter is valid only when the entity type is business object.
	ChildBizEntityIdList []*int64 `json:"ChildBizEntityIdList,omitempty" xml:"ChildBizEntityIdList,omitempty" type:"Repeated"`
	// The ID of the data domain to which the business process belongs.
	//
	// example:
	//
	// 20101011
	DataDomainId *int64 `json:"DataDomainId,omitempty" xml:"DataDomainId,omitempty"`
	// The description of the business entity.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the business entity.
	//
	// example:
	//
	// 业务活动测试
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the business entity was created.
	//
	// example:
	//
	// 2024-10-10 10:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the business entity was last modified.
	//
	// example:
	//
	// 2024-10-10 10:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the entity is a child business object.
	HasChildBizEntity *bool `json:"HasChildBizEntity,omitempty" xml:"HasChildBizEntity,omitempty"`
	// The ID of the business entity.
	//
	// example:
	//
	// 12121111
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the user who last modified the business entity.
	//
	// example:
	//
	// 30010010
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// The name of the user who last modified the business entity.
	//
	// example:
	//
	// 张三
	LastModifierName *string `json:"LastModifierName,omitempty" xml:"LastModifierName,omitempty"`
	// Indicates whether the entity is a sublevel of a hierarchy dimension.
	LevelSubBizObject *bool `json:"LevelSubBizObject,omitempty" xml:"LevelSubBizObject,omitempty"`
	// The name of the business entity.
	//
	// example:
	//
	// create_process_code_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The online status of the business entity. Valid values:
	//
	// - SUBMITTED: not online.
	//
	// - APPROVING: pending approval.
	//
	// - PUBLISHED: online.
	//
	// - REJECT: failed to go online.
	//
	// example:
	//
	// SUBMITTED
	OnlineStatus *string `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// The ID of the business entity owner.
	//
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The ID of the business entity owner.
	//
	// example:
	//
	// 30010010
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The list of associated central business entity IDs.
	RefBizEntityIdList []*int64 `json:"RefBizEntityIdList,omitempty" xml:"RefBizEntityIdList,omitempty" type:"Repeated"`
	// The number of associated logical tables.
	//
	// example:
	//
	// 1
	RefTableCount *int32 `json:"RefTableCount,omitempty" xml:"RefTableCount,omitempty"`
	// The status of the business entity. Valid values:
	//
	// - DRAFT: draft or not online.
	//
	// - SUBMITTING: being submitted.
	//
	// - SUBMITTED: submitted.
	//
	// - DEVELOPING: under development.
	//
	// - PUBLISHING: being published.
	//
	// - PUBLISHED: published.
	//
	// example:
	//
	// SUBMITTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subtype of the business entity.
	//
	// example:
	//
	// NORMAL
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// The list of subsequent business process IDs. This parameter is valid only when the entity type is business process.
	SuffixBizEntityIdList []*int64 `json:"SuffixBizEntityIdList,omitempty" xml:"SuffixBizEntityIdList,omitempty" type:"Repeated"`
	// The type of the business entity. Valid values:
	//
	// - BIZ_OBJECT: business object.
	//
	// - BIZ_PROCESS: business process.
	//
	// example:
	//
	// BIZ_OBJECT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListBizEntitiesResponseBodyPageResultBizEntityList) String() string {
	return dara.Prettify(s)
}

func (s ListBizEntitiesResponseBodyPageResultBizEntityList) GoString() string {
	return s.String()
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetBelongToBizEntityIdList() []*int64 {
	return s.BelongToBizEntityIdList
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetChildBizEntityIdList() []*int64 {
	return s.ChildBizEntityIdList
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetDataDomainId() *int64 {
	return s.DataDomainId
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetDescription() *string {
	return s.Description
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetHasChildBizEntity() *bool {
	return s.HasChildBizEntity
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetId() *int64 {
	return s.Id
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetLastModifier() *string {
	return s.LastModifier
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetLastModifierName() *string {
	return s.LastModifierName
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetLevelSubBizObject() *bool {
	return s.LevelSubBizObject
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetName() *string {
	return s.Name
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetOnlineStatus() *string {
	return s.OnlineStatus
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetRefBizEntityIdList() []*int64 {
	return s.RefBizEntityIdList
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetRefTableCount() *int32 {
	return s.RefTableCount
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetStatus() *string {
	return s.Status
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetSubType() *string {
	return s.SubType
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetSuffixBizEntityIdList() []*int64 {
	return s.SuffixBizEntityIdList
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) GetType() *string {
	return s.Type
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetBelongToBizEntityIdList(v []*int64) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.BelongToBizEntityIdList = v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetBizUnitId(v int64) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.BizUnitId = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetChildBizEntityIdList(v []*int64) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.ChildBizEntityIdList = v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetDataDomainId(v int64) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.DataDomainId = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetDescription(v string) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.Description = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetDisplayName(v string) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.DisplayName = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetGmtCreate(v string) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.GmtCreate = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetGmtModified(v string) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.GmtModified = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetHasChildBizEntity(v bool) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.HasChildBizEntity = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetId(v int64) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.Id = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetLastModifier(v string) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.LastModifier = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetLastModifierName(v string) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.LastModifierName = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetLevelSubBizObject(v bool) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.LevelSubBizObject = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetName(v string) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.Name = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetOnlineStatus(v string) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.OnlineStatus = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetOwnerName(v string) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.OwnerName = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetOwnerUserId(v string) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.OwnerUserId = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetRefBizEntityIdList(v []*int64) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.RefBizEntityIdList = v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetRefTableCount(v int32) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.RefTableCount = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetStatus(v string) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.Status = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetSubType(v string) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.SubType = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetSuffixBizEntityIdList(v []*int64) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.SuffixBizEntityIdList = v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) SetType(v string) *ListBizEntitiesResponseBodyPageResultBizEntityList {
	s.Type = &v
	return s
}

func (s *ListBizEntitiesResponseBodyPageResultBizEntityList) Validate() error {
	return dara.Validate(s)
}
