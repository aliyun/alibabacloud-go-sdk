// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProject interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *Project
	GetCreateTime() *string
	SetDatasetCount(v int64) *Project
	GetDatasetCount() *int64
	SetDatasetMaxBindCount(v int64) *Project
	GetDatasetMaxBindCount() *int64
	SetDatasetMaxEntityCount(v int64) *Project
	GetDatasetMaxEntityCount() *int64
	SetDatasetMaxFileCount(v int64) *Project
	GetDatasetMaxFileCount() *int64
	SetDatasetMaxRelationCount(v int64) *Project
	GetDatasetMaxRelationCount() *int64
	SetDatasetMaxTotalFileSize(v int64) *Project
	GetDatasetMaxTotalFileSize() *int64
	SetDescription(v string) *Project
	GetDescription() *string
	SetEngineConcurrency(v int64) *Project
	GetEngineConcurrency() *int64
	SetFileCount(v int64) *Project
	GetFileCount() *int64
	SetProjectMaxDatasetCount(v int64) *Project
	GetProjectMaxDatasetCount() *int64
	SetProjectName(v string) *Project
	GetProjectName() *string
	SetProjectQueriesPerSecond(v int64) *Project
	GetProjectQueriesPerSecond() *int64
	SetServiceRole(v string) *Project
	GetServiceRole() *string
	SetTags(v []*ProjectTags) *Project
	GetTags() []*ProjectTags
	SetTemplateId(v string) *Project
	GetTemplateId() *string
	SetTotalFileSize(v int64) *Project
	GetTotalFileSize() *int64
	SetUpdateTime(v string) *Project
	GetUpdateTime() *string
}

type Project struct {
	CreateTime              *string        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatasetCount            *int64         `json:"DatasetCount,omitempty" xml:"DatasetCount,omitempty"`
	DatasetMaxBindCount     *int64         `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	DatasetMaxEntityCount   *int64         `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	DatasetMaxFileCount     *int64         `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	DatasetMaxRelationCount *int64         `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	DatasetMaxTotalFileSize *int64         `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	Description             *string        `json:"Description,omitempty" xml:"Description,omitempty"`
	EngineConcurrency       *int64         `json:"EngineConcurrency,omitempty" xml:"EngineConcurrency,omitempty"`
	FileCount               *int64         `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	ProjectMaxDatasetCount  *int64         `json:"ProjectMaxDatasetCount,omitempty" xml:"ProjectMaxDatasetCount,omitempty"`
	ProjectName             *string        `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectQueriesPerSecond *int64         `json:"ProjectQueriesPerSecond,omitempty" xml:"ProjectQueriesPerSecond,omitempty"`
	ServiceRole             *string        `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	Tags                    []*ProjectTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateId              *string        `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TotalFileSize           *int64         `json:"TotalFileSize,omitempty" xml:"TotalFileSize,omitempty"`
	UpdateTime              *string        `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Project) String() string {
	return dara.Prettify(s)
}

func (s Project) GoString() string {
	return s.String()
}

func (s *Project) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Project) GetDatasetCount() *int64 {
	return s.DatasetCount
}

func (s *Project) GetDatasetMaxBindCount() *int64 {
	return s.DatasetMaxBindCount
}

func (s *Project) GetDatasetMaxEntityCount() *int64 {
	return s.DatasetMaxEntityCount
}

func (s *Project) GetDatasetMaxFileCount() *int64 {
	return s.DatasetMaxFileCount
}

func (s *Project) GetDatasetMaxRelationCount() *int64 {
	return s.DatasetMaxRelationCount
}

func (s *Project) GetDatasetMaxTotalFileSize() *int64 {
	return s.DatasetMaxTotalFileSize
}

func (s *Project) GetDescription() *string {
	return s.Description
}

func (s *Project) GetEngineConcurrency() *int64 {
	return s.EngineConcurrency
}

func (s *Project) GetFileCount() *int64 {
	return s.FileCount
}

func (s *Project) GetProjectMaxDatasetCount() *int64 {
	return s.ProjectMaxDatasetCount
}

func (s *Project) GetProjectName() *string {
	return s.ProjectName
}

func (s *Project) GetProjectQueriesPerSecond() *int64 {
	return s.ProjectQueriesPerSecond
}

func (s *Project) GetServiceRole() *string {
	return s.ServiceRole
}

func (s *Project) GetTags() []*ProjectTags {
	return s.Tags
}

func (s *Project) GetTemplateId() *string {
	return s.TemplateId
}

func (s *Project) GetTotalFileSize() *int64 {
	return s.TotalFileSize
}

func (s *Project) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *Project) SetCreateTime(v string) *Project {
	s.CreateTime = &v
	return s
}

func (s *Project) SetDatasetCount(v int64) *Project {
	s.DatasetCount = &v
	return s
}

func (s *Project) SetDatasetMaxBindCount(v int64) *Project {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *Project) SetDatasetMaxEntityCount(v int64) *Project {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *Project) SetDatasetMaxFileCount(v int64) *Project {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *Project) SetDatasetMaxRelationCount(v int64) *Project {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *Project) SetDatasetMaxTotalFileSize(v int64) *Project {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *Project) SetDescription(v string) *Project {
	s.Description = &v
	return s
}

func (s *Project) SetEngineConcurrency(v int64) *Project {
	s.EngineConcurrency = &v
	return s
}

func (s *Project) SetFileCount(v int64) *Project {
	s.FileCount = &v
	return s
}

func (s *Project) SetProjectMaxDatasetCount(v int64) *Project {
	s.ProjectMaxDatasetCount = &v
	return s
}

func (s *Project) SetProjectName(v string) *Project {
	s.ProjectName = &v
	return s
}

func (s *Project) SetProjectQueriesPerSecond(v int64) *Project {
	s.ProjectQueriesPerSecond = &v
	return s
}

func (s *Project) SetServiceRole(v string) *Project {
	s.ServiceRole = &v
	return s
}

func (s *Project) SetTags(v []*ProjectTags) *Project {
	s.Tags = v
	return s
}

func (s *Project) SetTemplateId(v string) *Project {
	s.TemplateId = &v
	return s
}

func (s *Project) SetTotalFileSize(v int64) *Project {
	s.TotalFileSize = &v
	return s
}

func (s *Project) SetUpdateTime(v string) *Project {
	s.UpdateTime = &v
	return s
}

func (s *Project) Validate() error {
	return dara.Validate(s)
}

type ProjectTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ProjectTags) String() string {
	return dara.Prettify(s)
}

func (s ProjectTags) GoString() string {
	return s.String()
}

func (s *ProjectTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ProjectTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ProjectTags) SetTagKey(v string) *ProjectTags {
	s.TagKey = &v
	return s
}

func (s *ProjectTags) SetTagValue(v string) *ProjectTags {
	s.TagValue = &v
	return s
}

func (s *ProjectTags) Validate() error {
	return dara.Validate(s)
}
