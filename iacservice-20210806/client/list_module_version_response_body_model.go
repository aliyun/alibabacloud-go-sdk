// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModuleVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListModuleVersionResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModuleVersionResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListModuleVersionResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListModuleVersionResponseBody
	GetTotalCount() *int32
	SetVersions(v []*ListModuleVersionResponseBodyVersions) *ListModuleVersionResponseBody
	GetVersions() []*ListModuleVersionResponseBodyVersions
}

type ListModuleVersionResponseBody struct {
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
	// 792171BB-1A68-5148-8B9B-C7C728E1E98B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 6
	TotalCount *int32                                   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	Versions   []*ListModuleVersionResponseBodyVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ListModuleVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModuleVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListModuleVersionResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModuleVersionResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModuleVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModuleVersionResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListModuleVersionResponseBody) GetVersions() []*ListModuleVersionResponseBodyVersions {
	return s.Versions
}

func (s *ListModuleVersionResponseBody) SetPageNumber(v int32) *ListModuleVersionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListModuleVersionResponseBody) SetPageSize(v int32) *ListModuleVersionResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListModuleVersionResponseBody) SetRequestId(v string) *ListModuleVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModuleVersionResponseBody) SetTotalCount(v int32) *ListModuleVersionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModuleVersionResponseBody) SetVersions(v []*ListModuleVersionResponseBodyVersions) *ListModuleVersionResponseBody {
	s.Versions = v
	return s
}

func (s *ListModuleVersionResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListModuleVersionResponseBodyVersions struct {
	// example:
	//
	// 2022-05-13T02:21:49Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// mod-55f1739d9050fffed3ec3a2c4a5e5
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// v3
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
}

func (s ListModuleVersionResponseBodyVersions) String() string {
	return dara.Prettify(s)
}

func (s ListModuleVersionResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *ListModuleVersionResponseBodyVersions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListModuleVersionResponseBodyVersions) GetDescription() *string {
	return s.Description
}

func (s *ListModuleVersionResponseBodyVersions) GetModuleId() *string {
	return s.ModuleId
}

func (s *ListModuleVersionResponseBodyVersions) GetModuleVersion() *string {
	return s.ModuleVersion
}

func (s *ListModuleVersionResponseBodyVersions) GetName() *string {
	return s.Name
}

func (s *ListModuleVersionResponseBodyVersions) GetSourcePath() *string {
	return s.SourcePath
}

func (s *ListModuleVersionResponseBodyVersions) SetCreateTime(v string) *ListModuleVersionResponseBodyVersions {
	s.CreateTime = &v
	return s
}

func (s *ListModuleVersionResponseBodyVersions) SetDescription(v string) *ListModuleVersionResponseBodyVersions {
	s.Description = &v
	return s
}

func (s *ListModuleVersionResponseBodyVersions) SetModuleId(v string) *ListModuleVersionResponseBodyVersions {
	s.ModuleId = &v
	return s
}

func (s *ListModuleVersionResponseBodyVersions) SetModuleVersion(v string) *ListModuleVersionResponseBodyVersions {
	s.ModuleVersion = &v
	return s
}

func (s *ListModuleVersionResponseBodyVersions) SetName(v string) *ListModuleVersionResponseBodyVersions {
	s.Name = &v
	return s
}

func (s *ListModuleVersionResponseBodyVersions) SetSourcePath(v string) *ListModuleVersionResponseBodyVersions {
	s.SourcePath = &v
	return s
}

func (s *ListModuleVersionResponseBodyVersions) Validate() error {
	return dara.Validate(s)
}
