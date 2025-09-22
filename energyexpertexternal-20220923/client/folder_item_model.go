// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFolderItem interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentLevel(v int32) *FolderItem
	GetCurrentLevel() *int32
	SetDocCount(v int32) *FolderItem
	GetDocCount() *int32
	SetFolderDefault(v int32) *FolderItem
	GetFolderDefault() *int32
	SetFolderId(v string) *FolderItem
	GetFolderId() *string
	SetFolderName(v string) *FolderItem
	GetFolderName() *string
	SetFolderNum(v int32) *FolderItem
	GetFolderNum() *int32
	SetOssDomain(v string) *FolderItem
	GetOssDomain() *string
	SetOssPath(v string) *FolderItem
	GetOssPath() *string
	SetOssUpdateBy(v string) *FolderItem
	GetOssUpdateBy() *string
	SetParentFolderId(v string) *FolderItem
	GetParentFolderId() *string
	SetResourcePath(v string) *FolderItem
	GetResourcePath() *string
	SetStorageType(v int32) *FolderItem
	GetStorageType() *int32
	SetSubFolderList(v []*FolderItem) *FolderItem
	GetSubFolderList() []*FolderItem
	SetSyncParsingStatus(v int32) *FolderItem
	GetSyncParsingStatus() *int32
	SetSyncStatus(v int32) *FolderItem
	GetSyncStatus() *int32
	SetTaskId(v int64) *FolderItem
	GetTaskId() *int64
}

type FolderItem struct {
	CurrentLevel      *int32        `json:"currentLevel,omitempty" xml:"currentLevel,omitempty"`
	DocCount          *int32        `json:"docCount,omitempty" xml:"docCount,omitempty"`
	FolderDefault     *int32        `json:"folderDefault,omitempty" xml:"folderDefault,omitempty"`
	FolderId          *string       `json:"folderId,omitempty" xml:"folderId,omitempty"`
	FolderName        *string       `json:"folderName,omitempty" xml:"folderName,omitempty"`
	FolderNum         *int32        `json:"folderNum,omitempty" xml:"folderNum,omitempty"`
	OssDomain         *string       `json:"ossDomain,omitempty" xml:"ossDomain,omitempty"`
	OssPath           *string       `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	OssUpdateBy       *string       `json:"ossUpdateBy,omitempty" xml:"ossUpdateBy,omitempty"`
	ParentFolderId    *string       `json:"parentFolderId,omitempty" xml:"parentFolderId,omitempty"`
	ResourcePath      *string       `json:"resourcePath,omitempty" xml:"resourcePath,omitempty"`
	StorageType       *int32        `json:"storageType,omitempty" xml:"storageType,omitempty"`
	SubFolderList     []*FolderItem `json:"subFolderList,omitempty" xml:"subFolderList,omitempty" type:"Repeated"`
	SyncParsingStatus *int32        `json:"syncParsingStatus,omitempty" xml:"syncParsingStatus,omitempty"`
	SyncStatus        *int32        `json:"syncStatus,omitempty" xml:"syncStatus,omitempty"`
	TaskId            *int64        `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s FolderItem) String() string {
	return dara.Prettify(s)
}

func (s FolderItem) GoString() string {
	return s.String()
}

func (s *FolderItem) GetCurrentLevel() *int32 {
	return s.CurrentLevel
}

func (s *FolderItem) GetDocCount() *int32 {
	return s.DocCount
}

func (s *FolderItem) GetFolderDefault() *int32 {
	return s.FolderDefault
}

func (s *FolderItem) GetFolderId() *string {
	return s.FolderId
}

func (s *FolderItem) GetFolderName() *string {
	return s.FolderName
}

func (s *FolderItem) GetFolderNum() *int32 {
	return s.FolderNum
}

func (s *FolderItem) GetOssDomain() *string {
	return s.OssDomain
}

func (s *FolderItem) GetOssPath() *string {
	return s.OssPath
}

func (s *FolderItem) GetOssUpdateBy() *string {
	return s.OssUpdateBy
}

func (s *FolderItem) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *FolderItem) GetResourcePath() *string {
	return s.ResourcePath
}

func (s *FolderItem) GetStorageType() *int32 {
	return s.StorageType
}

func (s *FolderItem) GetSubFolderList() []*FolderItem {
	return s.SubFolderList
}

func (s *FolderItem) GetSyncParsingStatus() *int32 {
	return s.SyncParsingStatus
}

func (s *FolderItem) GetSyncStatus() *int32 {
	return s.SyncStatus
}

func (s *FolderItem) GetTaskId() *int64 {
	return s.TaskId
}

func (s *FolderItem) SetCurrentLevel(v int32) *FolderItem {
	s.CurrentLevel = &v
	return s
}

func (s *FolderItem) SetDocCount(v int32) *FolderItem {
	s.DocCount = &v
	return s
}

func (s *FolderItem) SetFolderDefault(v int32) *FolderItem {
	s.FolderDefault = &v
	return s
}

func (s *FolderItem) SetFolderId(v string) *FolderItem {
	s.FolderId = &v
	return s
}

func (s *FolderItem) SetFolderName(v string) *FolderItem {
	s.FolderName = &v
	return s
}

func (s *FolderItem) SetFolderNum(v int32) *FolderItem {
	s.FolderNum = &v
	return s
}

func (s *FolderItem) SetOssDomain(v string) *FolderItem {
	s.OssDomain = &v
	return s
}

func (s *FolderItem) SetOssPath(v string) *FolderItem {
	s.OssPath = &v
	return s
}

func (s *FolderItem) SetOssUpdateBy(v string) *FolderItem {
	s.OssUpdateBy = &v
	return s
}

func (s *FolderItem) SetParentFolderId(v string) *FolderItem {
	s.ParentFolderId = &v
	return s
}

func (s *FolderItem) SetResourcePath(v string) *FolderItem {
	s.ResourcePath = &v
	return s
}

func (s *FolderItem) SetStorageType(v int32) *FolderItem {
	s.StorageType = &v
	return s
}

func (s *FolderItem) SetSubFolderList(v []*FolderItem) *FolderItem {
	s.SubFolderList = v
	return s
}

func (s *FolderItem) SetSyncParsingStatus(v int32) *FolderItem {
	s.SyncParsingStatus = &v
	return s
}

func (s *FolderItem) SetSyncStatus(v int32) *FolderItem {
	s.SyncStatus = &v
	return s
}

func (s *FolderItem) SetTaskId(v int64) *FolderItem {
	s.TaskId = &v
	return s
}

func (s *FolderItem) Validate() error {
	return dara.Validate(s)
}
