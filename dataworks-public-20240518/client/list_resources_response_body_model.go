// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListResourcesResponseBodyPagingInfo) *ListResourcesResponseBody
	GetPagingInfo() *ListResourcesResponseBodyPagingInfo
	SetRequestId(v string) *ListResourcesResponseBody
	GetRequestId() *string
}

type ListResourcesResponseBody struct {
	// The pagination information.
	PagingInfo *ListResourcesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 99EBE7CF-69C0-5089-BE3E-79563C31XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) GetPagingInfo() *ListResourcesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourcesResponseBody) SetPagingInfo(v *ListResourcesResponseBodyPagingInfo) *ListResourcesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourcesResponseBodyPagingInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The queried file resources.
	Resources []*ListResourcesResponseBodyPagingInfoResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 131
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourcesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourcesResponseBodyPagingInfo) GetResources() []*ListResourcesResponseBodyPagingInfoResources {
	return s.Resources
}

func (s *ListResourcesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourcesResponseBodyPagingInfo) SetPageNumber(v int32) *ListResourcesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfo) SetPageSize(v int32) *ListResourcesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfo) SetResources(v []*ListResourcesResponseBodyPagingInfoResources) *ListResourcesResponseBodyPagingInfo {
	s.Resources = v
	return s
}

func (s *ListResourcesResponseBodyPagingInfo) SetTotalCount(v int32) *ListResourcesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfo) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourcesResponseBodyPagingInfoResources struct {
	// The time when the file resource was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1724505917000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The information about the data source.
	DataSource *ListResourcesResponseBodyPagingInfoResourcesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The ID of the file resource.
	//
	// example:
	//
	// 631478864897630XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The times when the file resource was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1724505917000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the file resource.
	//
	// example:
	//
	// math.py
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the file resource.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID.
	//
	// example:
	//
	// 344247
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The script information.
	Script *ListResourcesResponseBodyPagingInfoResourcesScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// The path of the source of the file resource. If the SourecType parameter is set to Local, this parameter is left empty.
	//
	// example:
	//
	// XXX/unknown/ide/1/XXX/20240820200851_963a9da676de44ef8d06a6576a8c4d6a.py
	SourcePath *string `json:"SourcePath,omitempty" xml:"SourcePath,omitempty"`
	// The storage type of the source of the file resource.
	//
	// Valid values:
	//
	// 	- Local
	//
	// 	- Oss
	//
	// example:
	//
	// local
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The storage path of the destination of the file resource.
	//
	// example:
	//
	// XXX/unknown/ide/1/XXX/20240820200851_963a9da676de44ef8d06a6576a8c4d6a.py
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The storage type of the destination of the file resource.
	//
	// Valid values:
	//
	// 	- Gateway
	//
	// 	- Oss
	//
	// 	- Hdfs
	//
	// example:
	//
	// oss
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The type of the file resource.
	//
	// Valid values:
	//
	// 	- Python
	//
	// 	- Jar
	//
	// 	- Archive
	//
	// 	- File
	//
	// example:
	//
	// jar
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcesResponseBodyPagingInfoResources) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfoResources) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfoResources) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListResourcesResponseBodyPagingInfoResources) GetDataSource() *ListResourcesResponseBodyPagingInfoResourcesDataSource {
	return s.DataSource
}

func (s *ListResourcesResponseBodyPagingInfoResources) GetId() *int64 {
	return s.Id
}

func (s *ListResourcesResponseBodyPagingInfoResources) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListResourcesResponseBodyPagingInfoResources) GetName() *string {
	return s.Name
}

func (s *ListResourcesResponseBodyPagingInfoResources) GetOwner() *string {
	return s.Owner
}

func (s *ListResourcesResponseBodyPagingInfoResources) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListResourcesResponseBodyPagingInfoResources) GetScript() *ListResourcesResponseBodyPagingInfoResourcesScript {
	return s.Script
}

func (s *ListResourcesResponseBodyPagingInfoResources) GetSourcePath() *string {
	return s.SourcePath
}

func (s *ListResourcesResponseBodyPagingInfoResources) GetSourceType() *string {
	return s.SourceType
}

func (s *ListResourcesResponseBodyPagingInfoResources) GetTargetPath() *string {
	return s.TargetPath
}

func (s *ListResourcesResponseBodyPagingInfoResources) GetTargetType() *string {
	return s.TargetType
}

func (s *ListResourcesResponseBodyPagingInfoResources) GetType() *string {
	return s.Type
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetCreateTime(v int64) *ListResourcesResponseBodyPagingInfoResources {
	s.CreateTime = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetDataSource(v *ListResourcesResponseBodyPagingInfoResourcesDataSource) *ListResourcesResponseBodyPagingInfoResources {
	s.DataSource = v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetId(v int64) *ListResourcesResponseBodyPagingInfoResources {
	s.Id = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetModifyTime(v int64) *ListResourcesResponseBodyPagingInfoResources {
	s.ModifyTime = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetName(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetOwner(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.Owner = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetProjectId(v int64) *ListResourcesResponseBodyPagingInfoResources {
	s.ProjectId = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetScript(v *ListResourcesResponseBodyPagingInfoResourcesScript) *ListResourcesResponseBodyPagingInfoResources {
	s.Script = v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetSourcePath(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.SourcePath = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetSourceType(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.SourceType = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetTargetPath(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.TargetPath = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetTargetType(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.TargetType = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetType(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.Type = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) Validate() error {
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			return err
		}
	}
	if s.Script != nil {
		if err := s.Script.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourcesResponseBodyPagingInfoResourcesDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// odps_first
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// odps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcesResponseBodyPagingInfoResourcesDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfoResourcesDataSource) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfoResourcesDataSource) GetName() *string {
	return s.Name
}

func (s *ListResourcesResponseBodyPagingInfoResourcesDataSource) GetType() *string {
	return s.Type
}

func (s *ListResourcesResponseBodyPagingInfoResourcesDataSource) SetName(v string) *ListResourcesResponseBodyPagingInfoResourcesDataSource {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResourcesDataSource) SetType(v string) *ListResourcesResponseBodyPagingInfoResourcesDataSource {
	s.Type = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResourcesDataSource) Validate() error {
	return dara.Validate(s)
}

type ListResourcesResponseBodyPagingInfoResourcesScript struct {
	// The script ID.
	//
	// example:
	//
	// 123348864897630XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The script path.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The runtime.
	Runtime *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListResourcesResponseBodyPagingInfoResourcesScript) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfoResourcesScript) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScript) GetId() *int64 {
	return s.Id
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScript) GetPath() *string {
	return s.Path
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScript) GetRuntime() *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime {
	return s.Runtime
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScript) SetId(v int64) *ListResourcesResponseBodyPagingInfoResourcesScript {
	s.Id = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScript) SetPath(v string) *ListResourcesResponseBodyPagingInfoResourcesScript {
	s.Path = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScript) SetRuntime(v *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) *ListResourcesResponseBodyPagingInfoResourcesScript {
	s.Runtime = v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScript) Validate() error {
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourcesResponseBodyPagingInfoResourcesScriptRuntime struct {
	// The command used to distinguish file resource types.
	//
	// example:
	//
	// ODPS_PYTHON
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) GetCommand() *string {
	return s.Command
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) SetCommand(v string) *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime {
	s.Command = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) Validate() error {
	return dara.Validate(s)
}
