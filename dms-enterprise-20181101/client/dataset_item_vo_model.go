// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasetItemVO interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskList(v []*AsyncTaskVO) *DatasetItemVO
	GetAsyncTaskList() []*AsyncTaskVO
	SetDatasetStatus(v int32) *DatasetItemVO
	GetDatasetStatus() *int32
	SetDatasetType(v int32) *DatasetItemVO
	GetDatasetType() *int32
	SetDigest(v string) *DatasetItemVO
	GetDigest() *string
	SetFileSystem(v string) *DatasetItemVO
	GetFileSystem() *string
	SetId(v string) *DatasetItemVO
	GetId() *string
	SetKeyName(v string) *DatasetItemVO
	GetKeyName() *string
	SetMoreInfo(v string) *DatasetItemVO
	GetMoreInfo() *string
	SetPath(v string) *DatasetItemVO
	GetPath() *string
	SetProjectsLinked(v []*ProjectDetailsLiteVO) *DatasetItemVO
	GetProjectsLinked() []*ProjectDetailsLiteVO
	SetRecentTaskStatus(v int32) *DatasetItemVO
	GetRecentTaskStatus() *int32
	SetRemark(v string) *DatasetItemVO
	GetRemark() *string
	SetSchema(v string) *DatasetItemVO
	GetSchema() *string
	SetTableName(v string) *DatasetItemVO
	GetTableName() *string
	SetUrl(v string) *DatasetItemVO
	GetUrl() *string
}

type DatasetItemVO struct {
	AsyncTaskList    []*AsyncTaskVO          `json:"AsyncTaskList,omitempty" xml:"AsyncTaskList,omitempty" type:"Repeated"`
	DatasetStatus    *int32                  `json:"DatasetStatus,omitempty" xml:"DatasetStatus,omitempty"`
	DatasetType      *int32                  `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	Digest           *string                 `json:"Digest,omitempty" xml:"Digest,omitempty"`
	FileSystem       *string                 `json:"FileSystem,omitempty" xml:"FileSystem,omitempty"`
	Id               *string                 `json:"Id,omitempty" xml:"Id,omitempty"`
	KeyName          *string                 `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	MoreInfo         *string                 `json:"MoreInfo,omitempty" xml:"MoreInfo,omitempty"`
	Path             *string                 `json:"Path,omitempty" xml:"Path,omitempty"`
	ProjectsLinked   []*ProjectDetailsLiteVO `json:"ProjectsLinked,omitempty" xml:"ProjectsLinked,omitempty" type:"Repeated"`
	RecentTaskStatus *int32                  `json:"RecentTaskStatus,omitempty" xml:"RecentTaskStatus,omitempty"`
	Remark           *string                 `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Schema           *string                 `json:"Schema,omitempty" xml:"Schema,omitempty"`
	TableName        *string                 `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Url              *string                 `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DatasetItemVO) String() string {
	return dara.Prettify(s)
}

func (s DatasetItemVO) GoString() string {
	return s.String()
}

func (s *DatasetItemVO) GetAsyncTaskList() []*AsyncTaskVO {
	return s.AsyncTaskList
}

func (s *DatasetItemVO) GetDatasetStatus() *int32 {
	return s.DatasetStatus
}

func (s *DatasetItemVO) GetDatasetType() *int32 {
	return s.DatasetType
}

func (s *DatasetItemVO) GetDigest() *string {
	return s.Digest
}

func (s *DatasetItemVO) GetFileSystem() *string {
	return s.FileSystem
}

func (s *DatasetItemVO) GetId() *string {
	return s.Id
}

func (s *DatasetItemVO) GetKeyName() *string {
	return s.KeyName
}

func (s *DatasetItemVO) GetMoreInfo() *string {
	return s.MoreInfo
}

func (s *DatasetItemVO) GetPath() *string {
	return s.Path
}

func (s *DatasetItemVO) GetProjectsLinked() []*ProjectDetailsLiteVO {
	return s.ProjectsLinked
}

func (s *DatasetItemVO) GetRecentTaskStatus() *int32 {
	return s.RecentTaskStatus
}

func (s *DatasetItemVO) GetRemark() *string {
	return s.Remark
}

func (s *DatasetItemVO) GetSchema() *string {
	return s.Schema
}

func (s *DatasetItemVO) GetTableName() *string {
	return s.TableName
}

func (s *DatasetItemVO) GetUrl() *string {
	return s.Url
}

func (s *DatasetItemVO) SetAsyncTaskList(v []*AsyncTaskVO) *DatasetItemVO {
	s.AsyncTaskList = v
	return s
}

func (s *DatasetItemVO) SetDatasetStatus(v int32) *DatasetItemVO {
	s.DatasetStatus = &v
	return s
}

func (s *DatasetItemVO) SetDatasetType(v int32) *DatasetItemVO {
	s.DatasetType = &v
	return s
}

func (s *DatasetItemVO) SetDigest(v string) *DatasetItemVO {
	s.Digest = &v
	return s
}

func (s *DatasetItemVO) SetFileSystem(v string) *DatasetItemVO {
	s.FileSystem = &v
	return s
}

func (s *DatasetItemVO) SetId(v string) *DatasetItemVO {
	s.Id = &v
	return s
}

func (s *DatasetItemVO) SetKeyName(v string) *DatasetItemVO {
	s.KeyName = &v
	return s
}

func (s *DatasetItemVO) SetMoreInfo(v string) *DatasetItemVO {
	s.MoreInfo = &v
	return s
}

func (s *DatasetItemVO) SetPath(v string) *DatasetItemVO {
	s.Path = &v
	return s
}

func (s *DatasetItemVO) SetProjectsLinked(v []*ProjectDetailsLiteVO) *DatasetItemVO {
	s.ProjectsLinked = v
	return s
}

func (s *DatasetItemVO) SetRecentTaskStatus(v int32) *DatasetItemVO {
	s.RecentTaskStatus = &v
	return s
}

func (s *DatasetItemVO) SetRemark(v string) *DatasetItemVO {
	s.Remark = &v
	return s
}

func (s *DatasetItemVO) SetSchema(v string) *DatasetItemVO {
	s.Schema = &v
	return s
}

func (s *DatasetItemVO) SetTableName(v string) *DatasetItemVO {
	s.TableName = &v
	return s
}

func (s *DatasetItemVO) SetUrl(v string) *DatasetItemVO {
	s.Url = &v
	return s
}

func (s *DatasetItemVO) Validate() error {
	if s.AsyncTaskList != nil {
		for _, item := range s.AsyncTaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ProjectsLinked != nil {
		for _, item := range s.ProjectsLinked {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
