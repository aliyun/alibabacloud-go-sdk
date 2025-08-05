// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetPackageResponseBodyData) *GetPackageResponseBody
	GetData() *GetPackageResponseBodyData
	SetErrorCode(v string) *GetPackageResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetPackageResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *GetPackageResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetPackageResponseBody
	GetRequestId() *string
}

type GetPackageResponseBody struct {
	// The returned data.
	Data *GetPackageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// 040002
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// error message.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Indicates whether the request was successful. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0b57ff8316614119858417939e3e54
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPackageResponseBody) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBody) GetData() *GetPackageResponseBodyData {
	return s.Data
}

func (s *GetPackageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPackageResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetPackageResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPackageResponseBody) SetData(v *GetPackageResponseBodyData) *GetPackageResponseBody {
	s.Data = v
	return s
}

func (s *GetPackageResponseBody) SetErrorCode(v string) *GetPackageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPackageResponseBody) SetErrorMsg(v string) *GetPackageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetPackageResponseBody) SetHttpCode(v int32) *GetPackageResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetPackageResponseBody) SetRequestId(v string) *GetPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPackageResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPackageResponseBodyData struct {
	// The projects in which the package is installed.
	AllowedProjectList []*GetPackageResponseBodyDataAllowedProjectList `json:"allowedProjectList,omitempty" xml:"allowedProjectList,omitempty" type:"Repeated"`
	// The details of the resources that are included in the package.
	ResourceList *GetPackageResponseBodyDataResourceList `json:"resourceList,omitempty" xml:"resourceList,omitempty" type:"Struct"`
}

func (s GetPackageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPackageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyData) GetAllowedProjectList() []*GetPackageResponseBodyDataAllowedProjectList {
	return s.AllowedProjectList
}

func (s *GetPackageResponseBodyData) GetResourceList() *GetPackageResponseBodyDataResourceList {
	return s.ResourceList
}

func (s *GetPackageResponseBodyData) SetAllowedProjectList(v []*GetPackageResponseBodyDataAllowedProjectList) *GetPackageResponseBodyData {
	s.AllowedProjectList = v
	return s
}

func (s *GetPackageResponseBodyData) SetResourceList(v *GetPackageResponseBodyDataResourceList) *GetPackageResponseBodyData {
	s.ResourceList = v
	return s
}

func (s *GetPackageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetPackageResponseBodyDataAllowedProjectList struct {
	// The security level for sensitive data.
	//
	// example:
	//
	// 2
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The name of the MaxCompute project.
	//
	// example:
	//
	// proejctB
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s GetPackageResponseBodyDataAllowedProjectList) String() string {
	return dara.Prettify(s)
}

func (s GetPackageResponseBodyDataAllowedProjectList) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyDataAllowedProjectList) GetLabel() *string {
	return s.Label
}

func (s *GetPackageResponseBodyDataAllowedProjectList) GetProject() *string {
	return s.Project
}

func (s *GetPackageResponseBodyDataAllowedProjectList) SetLabel(v string) *GetPackageResponseBodyDataAllowedProjectList {
	s.Label = &v
	return s
}

func (s *GetPackageResponseBodyDataAllowedProjectList) SetProject(v string) *GetPackageResponseBodyDataAllowedProjectList {
	s.Project = &v
	return s
}

func (s *GetPackageResponseBodyDataAllowedProjectList) Validate() error {
	return dara.Validate(s)
}

type GetPackageResponseBodyDataResourceList struct {
	// The functions.
	Function []*GetPackageResponseBodyDataResourceListFunction `json:"function,omitempty" xml:"function,omitempty" type:"Repeated"`
	// The resources.
	Resource []*GetPackageResponseBodyDataResourceListResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Repeated"`
	// The tables.
	Table []*GetPackageResponseBodyDataResourceListTable `json:"table,omitempty" xml:"table,omitempty" type:"Repeated"`
}

func (s GetPackageResponseBodyDataResourceList) String() string {
	return dara.Prettify(s)
}

func (s GetPackageResponseBodyDataResourceList) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyDataResourceList) GetFunction() []*GetPackageResponseBodyDataResourceListFunction {
	return s.Function
}

func (s *GetPackageResponseBodyDataResourceList) GetResource() []*GetPackageResponseBodyDataResourceListResource {
	return s.Resource
}

func (s *GetPackageResponseBodyDataResourceList) GetTable() []*GetPackageResponseBodyDataResourceListTable {
	return s.Table
}

func (s *GetPackageResponseBodyDataResourceList) SetFunction(v []*GetPackageResponseBodyDataResourceListFunction) *GetPackageResponseBodyDataResourceList {
	s.Function = v
	return s
}

func (s *GetPackageResponseBodyDataResourceList) SetResource(v []*GetPackageResponseBodyDataResourceListResource) *GetPackageResponseBodyDataResourceList {
	s.Resource = v
	return s
}

func (s *GetPackageResponseBodyDataResourceList) SetTable(v []*GetPackageResponseBodyDataResourceListTable) *GetPackageResponseBodyDataResourceList {
	s.Table = v
	return s
}

func (s *GetPackageResponseBodyDataResourceList) Validate() error {
	return dara.Validate(s)
}

type GetPackageResponseBodyDataResourceListFunction struct {
	// The operations that were performed on the function.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the function.
	//
	// example:
	//
	// function_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of schema.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetPackageResponseBodyDataResourceListFunction) String() string {
	return dara.Prettify(s)
}

func (s GetPackageResponseBodyDataResourceListFunction) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyDataResourceListFunction) GetActions() []*string {
	return s.Actions
}

func (s *GetPackageResponseBodyDataResourceListFunction) GetName() *string {
	return s.Name
}

func (s *GetPackageResponseBodyDataResourceListFunction) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetPackageResponseBodyDataResourceListFunction) SetActions(v []*string) *GetPackageResponseBodyDataResourceListFunction {
	s.Actions = v
	return s
}

func (s *GetPackageResponseBodyDataResourceListFunction) SetName(v string) *GetPackageResponseBodyDataResourceListFunction {
	s.Name = &v
	return s
}

func (s *GetPackageResponseBodyDataResourceListFunction) SetSchemaName(v string) *GetPackageResponseBodyDataResourceListFunction {
	s.SchemaName = &v
	return s
}

func (s *GetPackageResponseBodyDataResourceListFunction) Validate() error {
	return dara.Validate(s)
}

type GetPackageResponseBodyDataResourceListResource struct {
	// The operations that were performed on the resource.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the resource.
	//
	// example:
	//
	// res_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of schema.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetPackageResponseBodyDataResourceListResource) String() string {
	return dara.Prettify(s)
}

func (s GetPackageResponseBodyDataResourceListResource) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyDataResourceListResource) GetActions() []*string {
	return s.Actions
}

func (s *GetPackageResponseBodyDataResourceListResource) GetName() *string {
	return s.Name
}

func (s *GetPackageResponseBodyDataResourceListResource) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetPackageResponseBodyDataResourceListResource) SetActions(v []*string) *GetPackageResponseBodyDataResourceListResource {
	s.Actions = v
	return s
}

func (s *GetPackageResponseBodyDataResourceListResource) SetName(v string) *GetPackageResponseBodyDataResourceListResource {
	s.Name = &v
	return s
}

func (s *GetPackageResponseBodyDataResourceListResource) SetSchemaName(v string) *GetPackageResponseBodyDataResourceListResource {
	s.SchemaName = &v
	return s
}

func (s *GetPackageResponseBodyDataResourceListResource) Validate() error {
	return dara.Validate(s)
}

type GetPackageResponseBodyDataResourceListTable struct {
	// The operations that were performed on the table.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the table.
	//
	// example:
	//
	// dim_odps
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of schema.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetPackageResponseBodyDataResourceListTable) String() string {
	return dara.Prettify(s)
}

func (s GetPackageResponseBodyDataResourceListTable) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyDataResourceListTable) GetActions() []*string {
	return s.Actions
}

func (s *GetPackageResponseBodyDataResourceListTable) GetName() *string {
	return s.Name
}

func (s *GetPackageResponseBodyDataResourceListTable) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetPackageResponseBodyDataResourceListTable) SetActions(v []*string) *GetPackageResponseBodyDataResourceListTable {
	s.Actions = v
	return s
}

func (s *GetPackageResponseBodyDataResourceListTable) SetName(v string) *GetPackageResponseBodyDataResourceListTable {
	s.Name = &v
	return s
}

func (s *GetPackageResponseBodyDataResourceListTable) SetSchemaName(v string) *GetPackageResponseBodyDataResourceListTable {
	s.SchemaName = &v
	return s
}

func (s *GetPackageResponseBodyDataResourceListTable) Validate() error {
	return dara.Validate(s)
}
