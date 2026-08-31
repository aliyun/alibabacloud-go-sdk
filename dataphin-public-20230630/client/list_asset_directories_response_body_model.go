// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetDirectoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAssetDirectoriesResponseBody
	GetCode() *string
	SetData(v *ListAssetDirectoriesResponseBodyData) *ListAssetDirectoriesResponseBody
	GetData() *ListAssetDirectoriesResponseBodyData
	SetHttpStatusCode(v int32) *ListAssetDirectoriesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAssetDirectoriesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAssetDirectoriesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAssetDirectoriesResponseBody
	GetSuccess() *bool
}

type ListAssetDirectoriesResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The paginated result of asset topic folders.
	Data *ListAssetDirectoriesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAssetDirectoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAssetDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssetDirectoriesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAssetDirectoriesResponseBody) GetData() *ListAssetDirectoriesResponseBodyData {
	return s.Data
}

func (s *ListAssetDirectoriesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAssetDirectoriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAssetDirectoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAssetDirectoriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAssetDirectoriesResponseBody) SetCode(v string) *ListAssetDirectoriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAssetDirectoriesResponseBody) SetData(v *ListAssetDirectoriesResponseBodyData) *ListAssetDirectoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListAssetDirectoriesResponseBody) SetHttpStatusCode(v int32) *ListAssetDirectoriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAssetDirectoriesResponseBody) SetMessage(v string) *ListAssetDirectoriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAssetDirectoriesResponseBody) SetRequestId(v string) *ListAssetDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssetDirectoriesResponseBody) SetSuccess(v bool) *ListAssetDirectoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAssetDirectoriesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAssetDirectoriesResponseBodyData struct {
	// The folder list.
	DirectoryList []*ListAssetDirectoriesResponseBodyDataDirectoryList `json:"DirectoryList,omitempty" xml:"DirectoryList,omitempty" type:"Repeated"`
	// The topic ID.
	//
	// example:
	//
	// 471794724245
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	// The topic name.
	//
	// example:
	//
	// Data Governance
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// The total number of records that match the conditions.
	//
	// example:
	//
	// -165955346599
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAssetDirectoriesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAssetDirectoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAssetDirectoriesResponseBodyData) GetDirectoryList() []*ListAssetDirectoriesResponseBodyDataDirectoryList {
	return s.DirectoryList
}

func (s *ListAssetDirectoriesResponseBodyData) GetTopicId() *int64 {
	return s.TopicId
}

func (s *ListAssetDirectoriesResponseBodyData) GetTopicName() *string {
	return s.TopicName
}

func (s *ListAssetDirectoriesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAssetDirectoriesResponseBodyData) SetDirectoryList(v []*ListAssetDirectoriesResponseBodyDataDirectoryList) *ListAssetDirectoriesResponseBodyData {
	s.DirectoryList = v
	return s
}

func (s *ListAssetDirectoriesResponseBodyData) SetTopicId(v int64) *ListAssetDirectoriesResponseBodyData {
	s.TopicId = &v
	return s
}

func (s *ListAssetDirectoriesResponseBodyData) SetTopicName(v string) *ListAssetDirectoriesResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *ListAssetDirectoriesResponseBodyData) SetTotalCount(v int64) *ListAssetDirectoriesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListAssetDirectoriesResponseBodyData) Validate() error {
	if s.DirectoryList != nil {
		for _, item := range s.DirectoryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAssetDirectoriesResponseBodyDataDirectoryList struct {
	// The folder description.
	//
	// example:
	//
	// Core metrics asset folder
	DirectoryDescription *string `json:"DirectoryDescription,omitempty" xml:"DirectoryDescription,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// 8223183275
	DirectoryId *int64 `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The folder name.
	//
	// example:
	//
	// Core Metrics
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// The display path.
	//
	// example:
	//
	// /Data Governance/Core Metrics
	FullPath *string `json:"FullPath,omitempty" xml:"FullPath,omitempty"`
	// The ID path from the top level to the current folder.
	FullPathIds []*int64 `json:"FullPathIds,omitempty" xml:"FullPathIds,omitempty" type:"Repeated"`
	// The name path from the top level to the current folder.
	FullPathNames []*string `json:"FullPathNames,omitempty" xml:"FullPathNames,omitempty" type:"Repeated"`
	// Indicates whether published direct child folders exist.
	HasChildren *bool `json:"HasChildren,omitempty" xml:"HasChildren,omitempty"`
	// The absolute level of the folder.
	//
	// example:
	//
	// 7120213
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The last modifier.
	Modifier *ListAssetDirectoriesResponseBodyDataDirectoryListModifier `json:"Modifier,omitempty" xml:"Modifier,omitempty" type:"Struct"`
	// The last modified time.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The parent folder ID.
	//
	// example:
	//
	// -292276281678
	ParentDirectoryId *int64 `json:"ParentDirectoryId,omitempty" xml:"ParentDirectoryId,omitempty"`
}

func (s ListAssetDirectoriesResponseBodyDataDirectoryList) String() string {
	return dara.Prettify(s)
}

func (s ListAssetDirectoriesResponseBodyDataDirectoryList) GoString() string {
	return s.String()
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) GetDirectoryDescription() *string {
	return s.DirectoryDescription
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) GetDirectoryId() *int64 {
	return s.DirectoryId
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) GetFullPath() *string {
	return s.FullPath
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) GetFullPathIds() []*int64 {
	return s.FullPathIds
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) GetFullPathNames() []*string {
	return s.FullPathNames
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) GetHasChildren() *bool {
	return s.HasChildren
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) GetLevel() *int32 {
	return s.Level
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) GetModifier() *ListAssetDirectoriesResponseBodyDataDirectoryListModifier {
	return s.Modifier
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) GetParentDirectoryId() *int64 {
	return s.ParentDirectoryId
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) SetDirectoryDescription(v string) *ListAssetDirectoriesResponseBodyDataDirectoryList {
	s.DirectoryDescription = &v
	return s
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) SetDirectoryId(v int64) *ListAssetDirectoriesResponseBodyDataDirectoryList {
	s.DirectoryId = &v
	return s
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) SetDirectoryName(v string) *ListAssetDirectoriesResponseBodyDataDirectoryList {
	s.DirectoryName = &v
	return s
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) SetFullPath(v string) *ListAssetDirectoriesResponseBodyDataDirectoryList {
	s.FullPath = &v
	return s
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) SetFullPathIds(v []*int64) *ListAssetDirectoriesResponseBodyDataDirectoryList {
	s.FullPathIds = v
	return s
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) SetFullPathNames(v []*string) *ListAssetDirectoriesResponseBodyDataDirectoryList {
	s.FullPathNames = v
	return s
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) SetHasChildren(v bool) *ListAssetDirectoriesResponseBodyDataDirectoryList {
	s.HasChildren = &v
	return s
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) SetLevel(v int32) *ListAssetDirectoriesResponseBodyDataDirectoryList {
	s.Level = &v
	return s
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) SetModifier(v *ListAssetDirectoriesResponseBodyDataDirectoryListModifier) *ListAssetDirectoriesResponseBodyDataDirectoryList {
	s.Modifier = v
	return s
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) SetModifyTime(v string) *ListAssetDirectoriesResponseBodyDataDirectoryList {
	s.ModifyTime = &v
	return s
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) SetParentDirectoryId(v int64) *ListAssetDirectoriesResponseBodyDataDirectoryList {
	s.ParentDirectoryId = &v
	return s
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryList) Validate() error {
	if s.Modifier != nil {
		if err := s.Modifier.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAssetDirectoriesResponseBodyDataDirectoryListModifier struct {
	// The user ID.
	//
	// example:
	//
	// 30001011
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// John
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListAssetDirectoriesResponseBodyDataDirectoryListModifier) String() string {
	return dara.Prettify(s)
}

func (s ListAssetDirectoriesResponseBodyDataDirectoryListModifier) GoString() string {
	return s.String()
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryListModifier) GetUserId() *string {
	return s.UserId
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryListModifier) GetUserName() *string {
	return s.UserName
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryListModifier) SetUserId(v string) *ListAssetDirectoriesResponseBodyDataDirectoryListModifier {
	s.UserId = &v
	return s
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryListModifier) SetUserName(v string) *ListAssetDirectoriesResponseBodyDataDirectoryListModifier {
	s.UserName = &v
	return s
}

func (s *ListAssetDirectoriesResponseBodyDataDirectoryListModifier) Validate() error {
	return dara.Validate(s)
}
