// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUdfFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassName(v string) *CreateUdfFileRequest
	GetClassName() *string
	SetCmdDescription(v string) *CreateUdfFileRequest
	GetCmdDescription() *string
	SetCreateFolderIfNotExists(v bool) *CreateUdfFileRequest
	GetCreateFolderIfNotExists() *bool
	SetExample(v string) *CreateUdfFileRequest
	GetExample() *string
	SetFileFolderPath(v string) *CreateUdfFileRequest
	GetFileFolderPath() *string
	SetFileName(v string) *CreateUdfFileRequest
	GetFileName() *string
	SetFunctionType(v string) *CreateUdfFileRequest
	GetFunctionType() *string
	SetParameterDescription(v string) *CreateUdfFileRequest
	GetParameterDescription() *string
	SetProjectId(v int64) *CreateUdfFileRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *CreateUdfFileRequest
	GetProjectIdentifier() *string
	SetResources(v string) *CreateUdfFileRequest
	GetResources() *string
	SetReturnValue(v string) *CreateUdfFileRequest
	GetReturnValue() *string
	SetUdfDescription(v string) *CreateUdfFileRequest
	GetUdfDescription() *string
}

type CreateUdfFileRequest struct {
	// The name of the class in which the function is defined. This parameter corresponds to the Class Name parameter in the Register Function section of the configuration tab of the function in the DataWorks console.
	//
	// This parameter is required.
	//
	// example:
	//
	// com.alibaba.DataWorks.api.udf.StringConcat
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The syntax used for calling the function. This parameter corresponds to the Expression Syntax parameter in the Register Function section of the configuration tab of the function in the DataWorks console.
	//
	// example:
	//
	// StringConcat(String... substrs)
	CmdDescription *string `json:"CmdDescription,omitempty" xml:"CmdDescription,omitempty"`
	// Specifies whether to automatically create the directory that is specified by the FileFolderPath parameter if the directory does not exist. Valid values:
	//
	// 	- true: The system automatically creates the directory if the directory does not exist.
	//
	// 	- false: The system does not automatically create the directory if the directory does not exist. In this case, the call fails.
	//
	// example:
	//
	// false
	CreateFolderIfNotExists *bool `json:"CreateFolderIfNotExists,omitempty" xml:"CreateFolderIfNotExists,omitempty"`
	// The example for calling the function. This parameter corresponds to the Example parameter in the Register Function section of the configuration tab of the function in the DataWorks console.
	//
	// example:
	//
	// StringConcat(\\"a\\", \\"b\\", \\"c\\")
	Example *string `json:"Example,omitempty" xml:"Example,omitempty"`
	// The path of the folder in which the file for the function is stored.
	//
	// example:
	//
	// Business_process/First_Business_Process/function/string_processing
	FileFolderPath *string `json:"FileFolderPath,omitempty" xml:"FileFolderPath,omitempty"`
	// The name of the file for the function.
	//
	// This parameter is required.
	//
	// example:
	//
	// StringConcat
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The type of the function. Valid values: MATH, AGGREGATE, STRING, DATE, ANALYTIC, and OTHER. This parameter corresponds to the Function Type parameter in the Register Function section of the configuration tab of the function on the DataStudio page.
	//
	// This parameter is required.
	//
	// example:
	//
	// STRING
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The description of the input parameters of the function. This parameter corresponds to the Parameter Description parameter in the Register Function section of the configuration tab of the function on the DataStudio page.
	//
	// example:
	//
	// List of strings to be connected
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The ID of the DataWorks workspace. You can click the Workspace Manage icon in the upper-right corner of the DataStudio page to go to the Workspace Management page and view the workspace ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The unique identifier of the DataWorks workspace. You can click the identifier in the upper-left corner of the DataStudio page to switch to another workspace.
	//
	// You must specify either this parameter or the projectId parameter to determine the DataWorks workspace to which the operation is called.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// The names of the resources that are referenced by the function. This parameter corresponds to the Resources parameter in the Register Function section of the configuration tab of the function in the DataWorks console. Multiple resource names are separated by commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// string-concat-1.0.0.jar,commons-lang-2.6.jar
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The description of the return value of the function. This parameter corresponds to the Return Value parameter in the Register Function section of the configuration tab of the function on the DataStudio page.
	//
	// example:
	//
	// New strings generated by concatenating all strings before and after the input order
	ReturnValue *string `json:"ReturnValue,omitempty" xml:"ReturnValue,omitempty"`
	// The description of the function. This parameter corresponds to the Description parameter in the Register Function section of the configuration tab of the function on the DataStudio page.
	//
	// example:
	//
	// Concatenate several strings to generate a new string
	UdfDescription *string `json:"UdfDescription,omitempty" xml:"UdfDescription,omitempty"`
}

func (s CreateUdfFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUdfFileRequest) GoString() string {
	return s.String()
}

func (s *CreateUdfFileRequest) GetClassName() *string {
	return s.ClassName
}

func (s *CreateUdfFileRequest) GetCmdDescription() *string {
	return s.CmdDescription
}

func (s *CreateUdfFileRequest) GetCreateFolderIfNotExists() *bool {
	return s.CreateFolderIfNotExists
}

func (s *CreateUdfFileRequest) GetExample() *string {
	return s.Example
}

func (s *CreateUdfFileRequest) GetFileFolderPath() *string {
	return s.FileFolderPath
}

func (s *CreateUdfFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateUdfFileRequest) GetFunctionType() *string {
	return s.FunctionType
}

func (s *CreateUdfFileRequest) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *CreateUdfFileRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateUdfFileRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *CreateUdfFileRequest) GetResources() *string {
	return s.Resources
}

func (s *CreateUdfFileRequest) GetReturnValue() *string {
	return s.ReturnValue
}

func (s *CreateUdfFileRequest) GetUdfDescription() *string {
	return s.UdfDescription
}

func (s *CreateUdfFileRequest) SetClassName(v string) *CreateUdfFileRequest {
	s.ClassName = &v
	return s
}

func (s *CreateUdfFileRequest) SetCmdDescription(v string) *CreateUdfFileRequest {
	s.CmdDescription = &v
	return s
}

func (s *CreateUdfFileRequest) SetCreateFolderIfNotExists(v bool) *CreateUdfFileRequest {
	s.CreateFolderIfNotExists = &v
	return s
}

func (s *CreateUdfFileRequest) SetExample(v string) *CreateUdfFileRequest {
	s.Example = &v
	return s
}

func (s *CreateUdfFileRequest) SetFileFolderPath(v string) *CreateUdfFileRequest {
	s.FileFolderPath = &v
	return s
}

func (s *CreateUdfFileRequest) SetFileName(v string) *CreateUdfFileRequest {
	s.FileName = &v
	return s
}

func (s *CreateUdfFileRequest) SetFunctionType(v string) *CreateUdfFileRequest {
	s.FunctionType = &v
	return s
}

func (s *CreateUdfFileRequest) SetParameterDescription(v string) *CreateUdfFileRequest {
	s.ParameterDescription = &v
	return s
}

func (s *CreateUdfFileRequest) SetProjectId(v int64) *CreateUdfFileRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateUdfFileRequest) SetProjectIdentifier(v string) *CreateUdfFileRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *CreateUdfFileRequest) SetResources(v string) *CreateUdfFileRequest {
	s.Resources = &v
	return s
}

func (s *CreateUdfFileRequest) SetReturnValue(v string) *CreateUdfFileRequest {
	s.ReturnValue = &v
	return s
}

func (s *CreateUdfFileRequest) SetUdfDescription(v string) *CreateUdfFileRequest {
	s.UdfDescription = &v
	return s
}

func (s *CreateUdfFileRequest) Validate() error {
	return dara.Validate(s)
}
