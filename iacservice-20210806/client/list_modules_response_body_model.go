// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModules(v []*ListModulesResponseBodyModules) *ListModulesResponseBody
	GetModules() []*ListModulesResponseBodyModules
	SetPageNumber(v int32) *ListModulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListModulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListModulesResponseBody
	GetTotalCount() *int32
}

type ListModulesResponseBody struct {
	Modules []*ListModulesResponseBodyModules `json:"modules,omitempty" xml:"modules,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// C617E03B-3DD2-5F0C-A6CF-3028B499A2D5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 2790
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListModulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListModulesResponseBody) GetModules() []*ListModulesResponseBodyModules {
	return s.Modules
}

func (s *ListModulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListModulesResponseBody) SetModules(v []*ListModulesResponseBodyModules) *ListModulesResponseBody {
	s.Modules = v
	return s
}

func (s *ListModulesResponseBody) SetPageNumber(v int32) *ListModulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListModulesResponseBody) SetPageSize(v int32) *ListModulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListModulesResponseBody) SetRequestId(v string) *ListModulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModulesResponseBody) SetTotalCount(v int32) *ListModulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModulesResponseBody) Validate() error {
	if s.Modules != nil {
		for _, item := range s.Modules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModulesResponseBodyModules struct {
	// example:
	//
	// 2022-01-30T02:14:16Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// false
	DeletionProtection *bool `json:"deletionProtection,omitempty" xml:"deletionProtection,omitempty"`
	// example:
	//
	// description
	Description *string                                  `json:"description,omitempty" xml:"description,omitempty"`
	GroupInfo   *ListModulesResponseBodyModulesGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	// example:
	//
	// v1
	LatestVersion *string `json:"latestVersion,omitempty" xml:"latestVersion,omitempty"`
	// example:
	//
	// mod-518855d9a058cdbd3fd6951d59
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// Created
	Status *string                               `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*ListModulesResponseBodyModulesTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListModulesResponseBodyModules) String() string {
	return dara.Prettify(s)
}

func (s ListModulesResponseBodyModules) GoString() string {
	return s.String()
}

func (s *ListModulesResponseBodyModules) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListModulesResponseBodyModules) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ListModulesResponseBodyModules) GetDescription() *string {
	return s.Description
}

func (s *ListModulesResponseBodyModules) GetGroupInfo() *ListModulesResponseBodyModulesGroupInfo {
	return s.GroupInfo
}

func (s *ListModulesResponseBodyModules) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *ListModulesResponseBodyModules) GetModuleId() *string {
	return s.ModuleId
}

func (s *ListModulesResponseBodyModules) GetName() *string {
	return s.Name
}

func (s *ListModulesResponseBodyModules) GetSource() *string {
	return s.Source
}

func (s *ListModulesResponseBodyModules) GetStatus() *string {
	return s.Status
}

func (s *ListModulesResponseBodyModules) GetTags() []*ListModulesResponseBodyModulesTags {
	return s.Tags
}

func (s *ListModulesResponseBodyModules) SetCreateTime(v string) *ListModulesResponseBodyModules {
	s.CreateTime = &v
	return s
}

func (s *ListModulesResponseBodyModules) SetDeletionProtection(v bool) *ListModulesResponseBodyModules {
	s.DeletionProtection = &v
	return s
}

func (s *ListModulesResponseBodyModules) SetDescription(v string) *ListModulesResponseBodyModules {
	s.Description = &v
	return s
}

func (s *ListModulesResponseBodyModules) SetGroupInfo(v *ListModulesResponseBodyModulesGroupInfo) *ListModulesResponseBodyModules {
	s.GroupInfo = v
	return s
}

func (s *ListModulesResponseBodyModules) SetLatestVersion(v string) *ListModulesResponseBodyModules {
	s.LatestVersion = &v
	return s
}

func (s *ListModulesResponseBodyModules) SetModuleId(v string) *ListModulesResponseBodyModules {
	s.ModuleId = &v
	return s
}

func (s *ListModulesResponseBodyModules) SetName(v string) *ListModulesResponseBodyModules {
	s.Name = &v
	return s
}

func (s *ListModulesResponseBodyModules) SetSource(v string) *ListModulesResponseBodyModules {
	s.Source = &v
	return s
}

func (s *ListModulesResponseBodyModules) SetStatus(v string) *ListModulesResponseBodyModules {
	s.Status = &v
	return s
}

func (s *ListModulesResponseBodyModules) SetTags(v []*ListModulesResponseBodyModulesTags) *ListModulesResponseBodyModules {
	s.Tags = v
	return s
}

func (s *ListModulesResponseBodyModules) Validate() error {
	if s.GroupInfo != nil {
		if err := s.GroupInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModulesResponseBodyModulesGroupInfo struct {
	GroupId     *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupName   *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	ProjectId   *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s ListModulesResponseBodyModulesGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s ListModulesResponseBodyModulesGroupInfo) GoString() string {
	return s.String()
}

func (s *ListModulesResponseBodyModulesGroupInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *ListModulesResponseBodyModulesGroupInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *ListModulesResponseBodyModulesGroupInfo) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListModulesResponseBodyModulesGroupInfo) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListModulesResponseBodyModulesGroupInfo) SetGroupId(v string) *ListModulesResponseBodyModulesGroupInfo {
	s.GroupId = &v
	return s
}

func (s *ListModulesResponseBodyModulesGroupInfo) SetGroupName(v string) *ListModulesResponseBodyModulesGroupInfo {
	s.GroupName = &v
	return s
}

func (s *ListModulesResponseBodyModulesGroupInfo) SetProjectId(v string) *ListModulesResponseBodyModulesGroupInfo {
	s.ProjectId = &v
	return s
}

func (s *ListModulesResponseBodyModulesGroupInfo) SetProjectName(v string) *ListModulesResponseBodyModulesGroupInfo {
	s.ProjectName = &v
	return s
}

func (s *ListModulesResponseBodyModulesGroupInfo) Validate() error {
	return dara.Validate(s)
}

type ListModulesResponseBodyModulesTags struct {
	Key      *string `json:"key,omitempty" xml:"key,omitempty"`
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListModulesResponseBodyModulesTags) String() string {
	return dara.Prettify(s)
}

func (s ListModulesResponseBodyModulesTags) GoString() string {
	return s.String()
}

func (s *ListModulesResponseBodyModulesTags) GetKey() *string {
	return s.Key
}

func (s *ListModulesResponseBodyModulesTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListModulesResponseBodyModulesTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListModulesResponseBodyModulesTags) GetValue() *string {
	return s.Value
}

func (s *ListModulesResponseBodyModulesTags) SetKey(v string) *ListModulesResponseBodyModulesTags {
	s.Key = &v
	return s
}

func (s *ListModulesResponseBodyModulesTags) SetTagKey(v string) *ListModulesResponseBodyModulesTags {
	s.TagKey = &v
	return s
}

func (s *ListModulesResponseBodyModulesTags) SetTagValue(v string) *ListModulesResponseBodyModulesTags {
	s.TagValue = &v
	return s
}

func (s *ListModulesResponseBodyModulesTags) SetValue(v string) *ListModulesResponseBodyModulesTags {
	s.Value = &v
	return s
}

func (s *ListModulesResponseBodyModulesTags) Validate() error {
	return dara.Validate(s)
}
