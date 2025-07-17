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
	return dara.Validate(s)
}

type ListFunctionsResponseBodyPagingInfo struct {
	// The UDFs.
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
	return dara.Validate(s)
}

type ListFunctionsResponseBodyPagingInfoFunctions struct {
	// The file resources in an Advanced RISC Machines (ARM) cluster.
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
	// The description of the command.
	//
	// example:
	//
	// testUdf(xx,yy)
	CommandDescription *string `json:"CommandDescription,omitempty" xml:"CommandDescription,omitempty"`
	// The time when the UDF was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1655953028000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The data source information about the UDF.
	DataSource *ListFunctionsResponseBodyPagingInfoFunctionsDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The name of the database. This parameter is returned for E-MapReduce (EMR) functions.
	//
	// example:
	//
	// odps_first
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The overall description of the UDF.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The code of the embedded UDF.
	//
	// example:
	//
	// print(\\"hello,world!\\")
	EmbeddedCode *string `json:"EmbeddedCode,omitempty" xml:"EmbeddedCode,omitempty"`
	// The type of the nested code.
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
	// The type of the nested resource.
	//
	// Valid values:
	//
	// 	- File: general resources
	//
	// 	- Embedded: embedded resources
	//
	// example:
	//
	// File
	EmbeddedResourceType *string `json:"EmbeddedResourceType,omitempty" xml:"EmbeddedResourceType,omitempty"`
	// The description of the example.
	//
	// example:
	//
	// Example description >>> select tsetUdf(xx,yy);
	//
	// abc
	ExampleDescription *string `json:"ExampleDescription,omitempty" xml:"ExampleDescription,omitempty"`
	// The files resources.
	//
	// example:
	//
	// xxx.jar,yyy.jar
	FileResource *string `json:"FileResource,omitempty" xml:"FileResource,omitempty"`
	// The ID of the UDF.
	//
	// example:
	//
	// 580667964888595XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the UDF was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1655953028000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the UDF.
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
	// The description of the parameter.
	//
	// example:
	//
	// xx: parameter information XXX
	//
	// yy: parameter information YYY
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The ID of the workspace to which the UDF belongs.
	//
	// example:
	//
	// 307XXX
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The description of the return value.
	//
	// example:
	//
	// The return value is a string.
	ReturnValueDescription *string `json:"ReturnValueDescription,omitempty" xml:"ReturnValueDescription,omitempty"`
	// The information about the resource group used when you run the UDF.
	RuntimeResource *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The script information about the UDF.
	Script *ListFunctionsResponseBodyPagingInfoFunctionsScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// The UDF type.
	//
	// Valid values:
	//
	// 	- Math: mathematical operation function
	//
	// 	- Aggregate: aggregate function
	//
	// 	- String: string processing function
	//
	// 	- Date: date function
	//
	// 	- Analytic: window function
	//
	// 	- Other: other functions
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

func (s *ListFunctionsResponseBodyPagingInfoFunctions) GetId() *int64 {
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

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetId(v int64) *ListFunctionsResponseBodyPagingInfoFunctions {
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
	return dara.Validate(s)
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
	// The ID of the resource group used when you run the UDF.
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
	// The script ID.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The script path.
	//
	// example:
	//
	// XXX/OpenAPI/function/function_name
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The runtime.
	Runtime *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScript) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScript) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) GetId() *int64 {
	return s.Id
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) GetPath() *string {
	return s.Path
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) GetRuntime() *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime {
	return s.Runtime
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) SetId(v int64) *ListFunctionsResponseBodyPagingInfoFunctionsScript {
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
	return dara.Validate(s)
}

type ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime struct {
	// The command.
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
