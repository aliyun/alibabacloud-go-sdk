// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDataSetResponseBody
	GetCode() *string
	SetCount(v int32) *ListDataSetResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *ListDataSetResponseBody
	GetCurrentPage() *int32
	SetData(v *ListDataSetResponseBodyData) *ListDataSetResponseBody
	GetData() *ListDataSetResponseBodyData
	SetHttpStatusCode(v int32) *ListDataSetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDataSetResponseBody
	GetMessage() *string
	SetMessages(v *ListDataSetResponseBodyMessages) *ListDataSetResponseBody
	GetMessages() *ListDataSetResponseBodyMessages
	SetPageNumber(v int32) *ListDataSetResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataSetResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataSetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataSetResponseBody
	GetSuccess() *bool
}

type ListDataSetResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 23
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 10
	CurrentPage *int32                       `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        *ListDataSetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *ListDataSetResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 96138D8D-8D26-4E41-BFF4-77AED1088BBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSetResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDataSetResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListDataSetResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDataSetResponseBody) GetData() *ListDataSetResponseBodyData {
	return s.Data
}

func (s *ListDataSetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataSetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDataSetResponseBody) GetMessages() *ListDataSetResponseBodyMessages {
	return s.Messages
}

func (s *ListDataSetResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataSetResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataSetResponseBody) SetCode(v string) *ListDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataSetResponseBody) SetCount(v int32) *ListDataSetResponseBody {
	s.Count = &v
	return s
}

func (s *ListDataSetResponseBody) SetCurrentPage(v int32) *ListDataSetResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListDataSetResponseBody) SetData(v *ListDataSetResponseBodyData) *ListDataSetResponseBody {
	s.Data = v
	return s
}

func (s *ListDataSetResponseBody) SetHttpStatusCode(v int32) *ListDataSetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataSetResponseBody) SetMessage(v string) *ListDataSetResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataSetResponseBody) SetMessages(v *ListDataSetResponseBodyMessages) *ListDataSetResponseBody {
	s.Messages = v
	return s
}

func (s *ListDataSetResponseBody) SetPageNumber(v int32) *ListDataSetResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDataSetResponseBody) SetPageSize(v int32) *ListDataSetResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataSetResponseBody) SetRequestId(v string) *ListDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSetResponseBody) SetSuccess(v bool) *ListDataSetResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataSetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	if s.Messages != nil {
		if err := s.Messages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataSetResponseBodyData struct {
	Data []*ListDataSetResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListDataSetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataSetResponseBodyData) GetData() []*ListDataSetResponseBodyDataData {
	return s.Data
}

func (s *ListDataSetResponseBodyData) SetData(v []*ListDataSetResponseBodyDataData) *ListDataSetResponseBodyData {
	s.Data = v
	return s
}

func (s *ListDataSetResponseBodyData) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSetResponseBodyDataData struct {
	// example:
	//
	// 1
	AutoTranscoding *int32 `json:"AutoTranscoding,omitempty" xml:"AutoTranscoding,omitempty"`
	// example:
	//
	// 0
	ChannelId0 *int32 `json:"ChannelId0,omitempty" xml:"ChannelId0,omitempty"`
	// example:
	//
	// 1
	ChannelId1 *int32 `json:"ChannelId1,omitempty" xml:"ChannelId1,omitempty"`
	// example:
	//
	// 1
	ChannelType *int32 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// example:
	//
	// 2019-06-20T17:33Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 0
	CreateType *int32 `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	// example:
	//
	// 3
	DataSetType *int32 `json:"DataSetType,omitempty" xml:"DataSetType,omitempty"`
	// example:
	//
	// 0
	IsDelete *int32 `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	// example:
	//
	// filesFromLocal/ef7ff45c147a4a5e882c925f9a75ac43
	RoleConfigProp *string `json:"RoleConfigProp,omitempty" xml:"RoleConfigProp,omitempty"`
	// example:
	//
	// 1
	RoleConfigStatus *int32 `json:"RoleConfigStatus,omitempty" xml:"RoleConfigStatus,omitempty"`
	// example:
	//
	// xx
	RoleConfigTask *string `json:"RoleConfigTask,omitempty" xml:"RoleConfigTask,omitempty"`
	// example:
	//
	// “”
	SetBucketName *string `json:"SetBucketName,omitempty" xml:"SetBucketName,omitempty"`
	// example:
	//
	// “”
	SetDomain *string `json:"SetDomain,omitempty" xml:"SetDomain,omitempty"`
	// example:
	//
	// “”
	SetFolderName *string `json:"SetFolderName,omitempty" xml:"SetFolderName,omitempty"`
	// example:
	//
	// 1
	SetId   *int64  `json:"SetId,omitempty" xml:"SetId,omitempty"`
	SetName *string `json:"SetName,omitempty" xml:"SetName,omitempty"`
	// example:
	//
	// 1
	SetNumber *int32 `json:"SetNumber,omitempty" xml:"SetNumber,omitempty"`
	// example:
	//
	// “”
	SetRoleArn *string `json:"SetRoleArn,omitempty" xml:"SetRoleArn,omitempty"`
	// example:
	//
	// 1
	SetType *int32 `json:"SetType,omitempty" xml:"SetType,omitempty"`
	// example:
	//
	// 11
	SourceDataType *int32 `json:"SourceDataType,omitempty" xml:"SourceDataType,omitempty"`
	// example:
	//
	// xx
	SubDir *string `json:"SubDir,omitempty" xml:"SubDir,omitempty"`
	// example:
	//
	// 2022-05-10T10:34Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1
	UserGroup *string `json:"UserGroup,omitempty" xml:"UserGroup,omitempty"`
}

func (s ListDataSetResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ListDataSetResponseBodyDataData) GetAutoTranscoding() *int32 {
	return s.AutoTranscoding
}

func (s *ListDataSetResponseBodyDataData) GetChannelId0() *int32 {
	return s.ChannelId0
}

func (s *ListDataSetResponseBodyDataData) GetChannelId1() *int32 {
	return s.ChannelId1
}

func (s *ListDataSetResponseBodyDataData) GetChannelType() *int32 {
	return s.ChannelType
}

func (s *ListDataSetResponseBodyDataData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDataSetResponseBodyDataData) GetCreateType() *int32 {
	return s.CreateType
}

func (s *ListDataSetResponseBodyDataData) GetDataSetType() *int32 {
	return s.DataSetType
}

func (s *ListDataSetResponseBodyDataData) GetIsDelete() *int32 {
	return s.IsDelete
}

func (s *ListDataSetResponseBodyDataData) GetRoleConfigProp() *string {
	return s.RoleConfigProp
}

func (s *ListDataSetResponseBodyDataData) GetRoleConfigStatus() *int32 {
	return s.RoleConfigStatus
}

func (s *ListDataSetResponseBodyDataData) GetRoleConfigTask() *string {
	return s.RoleConfigTask
}

func (s *ListDataSetResponseBodyDataData) GetSetBucketName() *string {
	return s.SetBucketName
}

func (s *ListDataSetResponseBodyDataData) GetSetDomain() *string {
	return s.SetDomain
}

func (s *ListDataSetResponseBodyDataData) GetSetFolderName() *string {
	return s.SetFolderName
}

func (s *ListDataSetResponseBodyDataData) GetSetId() *int64 {
	return s.SetId
}

func (s *ListDataSetResponseBodyDataData) GetSetName() *string {
	return s.SetName
}

func (s *ListDataSetResponseBodyDataData) GetSetNumber() *int32 {
	return s.SetNumber
}

func (s *ListDataSetResponseBodyDataData) GetSetRoleArn() *string {
	return s.SetRoleArn
}

func (s *ListDataSetResponseBodyDataData) GetSetType() *int32 {
	return s.SetType
}

func (s *ListDataSetResponseBodyDataData) GetSourceDataType() *int32 {
	return s.SourceDataType
}

func (s *ListDataSetResponseBodyDataData) GetSubDir() *string {
	return s.SubDir
}

func (s *ListDataSetResponseBodyDataData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListDataSetResponseBodyDataData) GetUserGroup() *string {
	return s.UserGroup
}

func (s *ListDataSetResponseBodyDataData) SetAutoTranscoding(v int32) *ListDataSetResponseBodyDataData {
	s.AutoTranscoding = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetChannelId0(v int32) *ListDataSetResponseBodyDataData {
	s.ChannelId0 = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetChannelId1(v int32) *ListDataSetResponseBodyDataData {
	s.ChannelId1 = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetChannelType(v int32) *ListDataSetResponseBodyDataData {
	s.ChannelType = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetCreateTime(v string) *ListDataSetResponseBodyDataData {
	s.CreateTime = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetCreateType(v int32) *ListDataSetResponseBodyDataData {
	s.CreateType = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetDataSetType(v int32) *ListDataSetResponseBodyDataData {
	s.DataSetType = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetIsDelete(v int32) *ListDataSetResponseBodyDataData {
	s.IsDelete = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetRoleConfigProp(v string) *ListDataSetResponseBodyDataData {
	s.RoleConfigProp = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetRoleConfigStatus(v int32) *ListDataSetResponseBodyDataData {
	s.RoleConfigStatus = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetRoleConfigTask(v string) *ListDataSetResponseBodyDataData {
	s.RoleConfigTask = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSetBucketName(v string) *ListDataSetResponseBodyDataData {
	s.SetBucketName = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSetDomain(v string) *ListDataSetResponseBodyDataData {
	s.SetDomain = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSetFolderName(v string) *ListDataSetResponseBodyDataData {
	s.SetFolderName = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSetId(v int64) *ListDataSetResponseBodyDataData {
	s.SetId = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSetName(v string) *ListDataSetResponseBodyDataData {
	s.SetName = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSetNumber(v int32) *ListDataSetResponseBodyDataData {
	s.SetNumber = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSetRoleArn(v string) *ListDataSetResponseBodyDataData {
	s.SetRoleArn = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSetType(v int32) *ListDataSetResponseBodyDataData {
	s.SetType = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSourceDataType(v int32) *ListDataSetResponseBodyDataData {
	s.SourceDataType = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetSubDir(v string) *ListDataSetResponseBodyDataData {
	s.SubDir = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetUpdateTime(v string) *ListDataSetResponseBodyDataData {
	s.UpdateTime = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) SetUserGroup(v string) *ListDataSetResponseBodyDataData {
	s.UserGroup = &v
	return s
}

func (s *ListDataSetResponseBodyDataData) Validate() error {
	return dara.Validate(s)
}

type ListDataSetResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s ListDataSetResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *ListDataSetResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *ListDataSetResponseBodyMessages) SetMessage(v []*string) *ListDataSetResponseBodyMessages {
	s.Message = v
	return s
}

func (s *ListDataSetResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
