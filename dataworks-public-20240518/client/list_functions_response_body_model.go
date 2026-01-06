// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListFunctionsResponseBodyPagingInfo) *ListFunctionsResponseBody
	GetPagingInfo() *ListFunctionsResponseBodyPagingInfo
	SetRequestId(v string) *ListFunctionsResponseBody
	GetRequestId() *string
}

type ListFunctionsResponseBody struct {
	// The pagination information.
	PagingInfo *ListFunctionsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 89FB2BF0-EB00-5D03-9C34-05931001XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFunctionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBody) GetPagingInfo() *ListFunctionsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListFunctionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFunctionsResponseBody) SetPagingInfo(v *ListFunctionsResponseBodyPagingInfo) *ListFunctionsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListFunctionsResponseBody) SetRequestId(v string) *ListFunctionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFunctionsResponseBodyPagingInfo struct {
	// The function list.
	Functions []*ListFunctionsResponseBodyPagingInfoFunctions `json:"Functions,omitempty" xml:"Functions,omitempty" type:"Repeated"`
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
	// The total number of entries returned.
	//
	// example:
	//
	// 294
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfo) GetFunctions() []*ListFunctionsResponseBodyPagingInfoFunctions {
	return s.Functions
}

func (s *ListFunctionsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFunctionsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFunctionsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFunctionsResponseBodyPagingInfo) SetFunctions(v []*ListFunctionsResponseBodyPagingInfoFunctions) *ListFunctionsResponseBodyPagingInfo {
	s.Functions = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfo) SetPageNumber(v int32) *ListFunctionsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfo) SetPageSize(v int32) *ListFunctionsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfo) SetTotalCount(v int32) *ListFunctionsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfo) Validate() error {
	if s.Functions != nil {
		for _, item := range s.Functions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFunctionsResponseBodyPagingInfoFunctions struct {
	// The list of resource files for the ARM cluster.
	//
	// example:
	//
	// xxx.jar,yyy.jar
	ArmResource *string `json:"ArmResource,omitempty" xml:"ArmResource,omitempty"`
	// The fully qualified class name of the UDF.
	//
	// example:
	//
	// com.demo.Main
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The command description.
	//
	// example:
	//
	// testUdf(xx,yy)
	CommandDescription *string `json:"CommandDescription,omitempty" xml:"CommandDescription,omitempty"`
	// The timestamp when the UDF was created.
	//
	// example:
	//
	// 1655953028000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Data source information of the UDF.
	DataSource *ListFunctionsResponseBodyPagingInfoFunctionsDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The database name. This parameter is used only when the function type is EMR Function.
	//
	// example:
	//
	// odps_first
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The general description of the function.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Content of the nested function code
	//
	// example:
	//
	// print(\\"hello,world!\\")
	EmbeddedCode *string `json:"EmbeddedCode,omitempty" xml:"EmbeddedCode,omitempty"`
	// The nested code type.
	//
	// Valid values:
	//
	// 	- Python2
	//
	// 	- Python3
	//
	// 	- Java8
	//
	// 	- Java11
	//
	// 	- Java17
	//
	// example:
	//
	// Python2
	EmbeddedCodeType *string `json:"EmbeddedCodeType,omitempty" xml:"EmbeddedCodeType,omitempty"`
	// The nested resource type.
	//
	// Valid values:
	//
	// 	- File: General resource file.
	//
	// 	- Embedded: Embedded resource.
	//
	// example:
	//
	// File
	EmbeddedResourceType *string `json:"EmbeddedResourceType,omitempty" xml:"EmbeddedResourceType,omitempty"`
	// The example description.
	//
	// example:
	//
	// Example description >>> select tsetUdf(xx,yy);
	//
	// abc
	ExampleDescription *string `json:"ExampleDescription,omitempty" xml:"ExampleDescription,omitempty"`
	// The implementation code of the function and the list of resource files.
	//
	// example:
	//
	// xxx.jar,yyy.jar
	FileResource *string `json:"FileResource,omitempty" xml:"FileResource,omitempty"`
	// The unique identifier of the UDF.
	//
	// >  Prior to SDK version 8.0.0, this field is of type Long. In SDK version 8.0.0 and later, it is of type String. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. Compilation failures caused by the type change may occur only when you upgrade the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// example:
	//
	// 580667964888595XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1655953028000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The UDF name.
	//
	// example:
	//
	// Function name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the UDF.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The parameter description.
	//
	// example:
	//
	// xx: parameter information XXX
	//
	// yy: parameter information YYY
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The ID of the project to which the UDF belongs.
	//
	// example:
	//
	// 307XXX
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The return value description.
	//
	// example:
	//
	// The return value is a string.
	ReturnValueDescription *string `json:"ReturnValueDescription,omitempty" xml:"ReturnValueDescription,omitempty"`
	// The runtime resource group information.
	RuntimeResource *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// Script information of the UDF.
	Script *ListFunctionsResponseBodyPagingInfoFunctionsScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// The UDF type.
	//
	// Valid values:
	//
	// 	- Math: Mathematical operation functions
	//
	// 	- Aggregate: Aggregation functions
	//
	// 	- String: String processing functions
	//
	// 	- Date: Date functions
	//
	// 	- Analytic: Window functions
	//
	// 	- Other: Other functions
	//
	// example:
	//
	// MATH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctions) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctions) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetArmResource() *string {
	return s.ArmResource
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetClassName() *string {
	return s.ClassName
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetCommandDescription() *string {
	return s.CommandDescription
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetDataSource() *ListFunctionsResponseBodyPagingInfoFunctionsDataSource {
	return s.DataSource
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetDescription() *string {
	return s.Description
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetEmbeddedCode() *string {
	return s.EmbeddedCode
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetEmbeddedCodeType() *string {
	return s.EmbeddedCodeType
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetEmbeddedResourceType() *string {
	return s.EmbeddedResourceType
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetExampleDescription() *string {
	return s.ExampleDescription
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetFileResource() *string {
	return s.FileResource
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetId() *string {
	return s.Id
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetName() *string {
	return s.Name
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetOwner() *string {
	return s.Owner
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetReturnValueDescription() *string {
	return s.ReturnValueDescription
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetRuntimeResource() *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource {
	return s.RuntimeResource
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetScript() *ListFunctionsResponseBodyPagingInfoFunctionsScript {
	return s.Script
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetType() *string {
	return s.Type
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetArmResource(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ArmResource = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetClassName(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ClassName = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetCommandDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.CommandDescription = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetCreateTime(v int64) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.CreateTime = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetDataSource(v *ListFunctionsResponseBodyPagingInfoFunctionsDataSource) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.DataSource = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetDatabaseName(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.DatabaseName = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Description = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetEmbeddedCode(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.EmbeddedCode = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetEmbeddedCodeType(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.EmbeddedCodeType = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetEmbeddedResourceType(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.EmbeddedResourceType = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetExampleDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ExampleDescription = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetFileResource(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.FileResource = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetId(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Id = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetModifyTime(v int64) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ModifyTime = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetName(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Name = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetOwner(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Owner = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetParameterDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ParameterDescription = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetProjectId(v int64) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ProjectId = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetReturnValueDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ReturnValueDescription = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetRuntimeResource(v *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.RuntimeResource = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetScript(v *ListFunctionsResponseBodyPagingInfoFunctionsScript) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Script = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetType(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Type = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) Validate() error {
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			return err
		}
	}
	if s.RuntimeResource != nil {
		if err := s.RuntimeResource.Validate(); err != nil {
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

type ListFunctionsResponseBodyPagingInfoFunctionsDataSource struct {
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

func (s ListFunctionsResponseBodyPagingInfoFunctionsDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsDataSource) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsDataSource) GetName() *string {
	return s.Name
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsDataSource) GetType() *string {
	return s.Type
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsDataSource) SetName(v string) *ListFunctionsResponseBodyPagingInfoFunctionsDataSource {
	s.Name = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsDataSource) SetType(v string) *ListFunctionsResponseBodyPagingInfoFunctionsDataSource {
	s.Type = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsDataSource) Validate() error {
	return dara.Validate(s)
}

type ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource struct {
	// The runtime resource group ID.
	//
	// example:
	//
	// S_resgrop_xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) SetResourceGroupId(v string) *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type ListFunctionsResponseBodyPagingInfoFunctionsScript struct {
	// The ID of the script.
	//
	// >  This field is of type Long in SDK versions prior to 8.0.0, and of type String in SDK version 8.0.0 and later. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. However, compilation failures may occur due to the type change only when upgrading the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The script path.
	//
	// example:
	//
	// XXX/OpenAPI/function/function_name
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Runtime
	Runtime *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScript) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScript) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) GetId() *string {
	return s.Id
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) GetPath() *string {
	return s.Path
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) GetRuntime() *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime {
	return s.Runtime
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) SetId(v string) *ListFunctionsResponseBodyPagingInfoFunctionsScript {
	s.Id = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) SetPath(v string) *ListFunctionsResponseBodyPagingInfoFunctionsScript {
	s.Path = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) SetRuntime(v *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) *ListFunctionsResponseBodyPagingInfoFunctionsScript {
	s.Runtime = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) Validate() error {
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime struct {
	// Command
	//
	// example:
	//
	// ODPS_FUNCTION
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) GetCommand() *string {
	return s.Command
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) SetCommand(v string) *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime {
	s.Command = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) Validate() error {
	return dara.Validate(s)
}
