// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoleAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetRoleAclResponseBodyData) *GetRoleAclResponseBody
	GetData() *GetRoleAclResponseBodyData
	SetErrorCode(v string) *GetRoleAclResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetRoleAclResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *GetRoleAclResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetRoleAclResponseBody
	GetRequestId() *string
}

type GetRoleAclResponseBody struct {
	// The returned data.
	Data *GetRoleAclResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// error message
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0a06dc0916696898838762018e9564
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRoleAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRoleAclResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBody) GetData() *GetRoleAclResponseBodyData {
	return s.Data
}

func (s *GetRoleAclResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetRoleAclResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetRoleAclResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetRoleAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRoleAclResponseBody) SetData(v *GetRoleAclResponseBodyData) *GetRoleAclResponseBody {
	s.Data = v
	return s
}

func (s *GetRoleAclResponseBody) SetErrorCode(v string) *GetRoleAclResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRoleAclResponseBody) SetErrorMsg(v string) *GetRoleAclResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetRoleAclResponseBody) SetHttpCode(v int32) *GetRoleAclResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetRoleAclResponseBody) SetRequestId(v string) *GetRoleAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoleAclResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRoleAclResponseBodyData struct {
	// The function.
	Function []*GetRoleAclResponseBodyDataFunction `json:"function,omitempty" xml:"function,omitempty" type:"Repeated"`
	// The instance.
	Instance []*GetRoleAclResponseBodyDataInstance `json:"instance,omitempty" xml:"instance,omitempty" type:"Repeated"`
	// The package.
	Package []*GetRoleAclResponseBodyDataPackage `json:"package,omitempty" xml:"package,omitempty" type:"Repeated"`
	// The project.
	Project []*GetRoleAclResponseBodyDataProject `json:"project,omitempty" xml:"project,omitempty" type:"Repeated"`
	// The resource.
	Resource []*GetRoleAclResponseBodyDataResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Repeated"`
	// The table.
	Table []*GetRoleAclResponseBodyDataTable `json:"table,omitempty" xml:"table,omitempty" type:"Repeated"`
}

func (s GetRoleAclResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRoleAclResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyData) GetFunction() []*GetRoleAclResponseBodyDataFunction {
	return s.Function
}

func (s *GetRoleAclResponseBodyData) GetInstance() []*GetRoleAclResponseBodyDataInstance {
	return s.Instance
}

func (s *GetRoleAclResponseBodyData) GetPackage() []*GetRoleAclResponseBodyDataPackage {
	return s.Package
}

func (s *GetRoleAclResponseBodyData) GetProject() []*GetRoleAclResponseBodyDataProject {
	return s.Project
}

func (s *GetRoleAclResponseBodyData) GetResource() []*GetRoleAclResponseBodyDataResource {
	return s.Resource
}

func (s *GetRoleAclResponseBodyData) GetTable() []*GetRoleAclResponseBodyDataTable {
	return s.Table
}

func (s *GetRoleAclResponseBodyData) SetFunction(v []*GetRoleAclResponseBodyDataFunction) *GetRoleAclResponseBodyData {
	s.Function = v
	return s
}

func (s *GetRoleAclResponseBodyData) SetInstance(v []*GetRoleAclResponseBodyDataInstance) *GetRoleAclResponseBodyData {
	s.Instance = v
	return s
}

func (s *GetRoleAclResponseBodyData) SetPackage(v []*GetRoleAclResponseBodyDataPackage) *GetRoleAclResponseBodyData {
	s.Package = v
	return s
}

func (s *GetRoleAclResponseBodyData) SetProject(v []*GetRoleAclResponseBodyDataProject) *GetRoleAclResponseBodyData {
	s.Project = v
	return s
}

func (s *GetRoleAclResponseBodyData) SetResource(v []*GetRoleAclResponseBodyDataResource) *GetRoleAclResponseBodyData {
	s.Resource = v
	return s
}

func (s *GetRoleAclResponseBodyData) SetTable(v []*GetRoleAclResponseBodyDataTable) *GetRoleAclResponseBodyData {
	s.Table = v
	return s
}

func (s *GetRoleAclResponseBodyData) Validate() error {
	if s.Function != nil {
		for _, item := range s.Function {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Package != nil {
		for _, item := range s.Package {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Project != nil {
		for _, item := range s.Project {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Resource != nil {
		for _, item := range s.Resource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Table != nil {
		for _, item := range s.Table {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRoleAclResponseBodyDataFunction struct {
	// The operations that were performed on the function.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the function.
	//
	// example:
	//
	// functionA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Schema name.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataFunction) String() string {
	return dara.Prettify(s)
}

func (s GetRoleAclResponseBodyDataFunction) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataFunction) GetActions() []*string {
	return s.Actions
}

func (s *GetRoleAclResponseBodyDataFunction) GetName() *string {
	return s.Name
}

func (s *GetRoleAclResponseBodyDataFunction) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetRoleAclResponseBodyDataFunction) SetActions(v []*string) *GetRoleAclResponseBodyDataFunction {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataFunction) SetName(v string) *GetRoleAclResponseBodyDataFunction {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataFunction) SetSchemaName(v string) *GetRoleAclResponseBodyDataFunction {
	s.SchemaName = &v
	return s
}

func (s *GetRoleAclResponseBodyDataFunction) Validate() error {
	return dara.Validate(s)
}

type GetRoleAclResponseBodyDataInstance struct {
	// The operations that were performed on the instance.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the instance.
	//
	// example:
	//
	// instanceA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Schema name.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataInstance) String() string {
	return dara.Prettify(s)
}

func (s GetRoleAclResponseBodyDataInstance) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataInstance) GetActions() []*string {
	return s.Actions
}

func (s *GetRoleAclResponseBodyDataInstance) GetName() *string {
	return s.Name
}

func (s *GetRoleAclResponseBodyDataInstance) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetRoleAclResponseBodyDataInstance) SetActions(v []*string) *GetRoleAclResponseBodyDataInstance {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataInstance) SetName(v string) *GetRoleAclResponseBodyDataInstance {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataInstance) SetSchemaName(v string) *GetRoleAclResponseBodyDataInstance {
	s.SchemaName = &v
	return s
}

func (s *GetRoleAclResponseBodyDataInstance) Validate() error {
	return dara.Validate(s)
}

type GetRoleAclResponseBodyDataPackage struct {
	// The operations that were performed on the package.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the package.
	//
	// example:
	//
	// packageA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Schema name.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataPackage) String() string {
	return dara.Prettify(s)
}

func (s GetRoleAclResponseBodyDataPackage) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataPackage) GetActions() []*string {
	return s.Actions
}

func (s *GetRoleAclResponseBodyDataPackage) GetName() *string {
	return s.Name
}

func (s *GetRoleAclResponseBodyDataPackage) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetRoleAclResponseBodyDataPackage) SetActions(v []*string) *GetRoleAclResponseBodyDataPackage {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataPackage) SetName(v string) *GetRoleAclResponseBodyDataPackage {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataPackage) SetSchemaName(v string) *GetRoleAclResponseBodyDataPackage {
	s.SchemaName = &v
	return s
}

func (s *GetRoleAclResponseBodyDataPackage) Validate() error {
	return dara.Validate(s)
}

type GetRoleAclResponseBodyDataProject struct {
	// The operations that were performed on the project.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the MaxCompute project.
	//
	// example:
	//
	// projectA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Schema name.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataProject) String() string {
	return dara.Prettify(s)
}

func (s GetRoleAclResponseBodyDataProject) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataProject) GetActions() []*string {
	return s.Actions
}

func (s *GetRoleAclResponseBodyDataProject) GetName() *string {
	return s.Name
}

func (s *GetRoleAclResponseBodyDataProject) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetRoleAclResponseBodyDataProject) SetActions(v []*string) *GetRoleAclResponseBodyDataProject {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataProject) SetName(v string) *GetRoleAclResponseBodyDataProject {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataProject) SetSchemaName(v string) *GetRoleAclResponseBodyDataProject {
	s.SchemaName = &v
	return s
}

func (s *GetRoleAclResponseBodyDataProject) Validate() error {
	return dara.Validate(s)
}

type GetRoleAclResponseBodyDataResource struct {
	// The operations that were performed on the resource.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the resource.
	//
	// example:
	//
	// resourceA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Schema name.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataResource) String() string {
	return dara.Prettify(s)
}

func (s GetRoleAclResponseBodyDataResource) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataResource) GetActions() []*string {
	return s.Actions
}

func (s *GetRoleAclResponseBodyDataResource) GetName() *string {
	return s.Name
}

func (s *GetRoleAclResponseBodyDataResource) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetRoleAclResponseBodyDataResource) SetActions(v []*string) *GetRoleAclResponseBodyDataResource {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataResource) SetName(v string) *GetRoleAclResponseBodyDataResource {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataResource) SetSchemaName(v string) *GetRoleAclResponseBodyDataResource {
	s.SchemaName = &v
	return s
}

func (s *GetRoleAclResponseBodyDataResource) Validate() error {
	return dara.Validate(s)
}

type GetRoleAclResponseBodyDataTable struct {
	// The operations that were performed on the table.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the table.
	//
	// example:
	//
	// tableA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Schema name.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataTable) String() string {
	return dara.Prettify(s)
}

func (s GetRoleAclResponseBodyDataTable) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataTable) GetActions() []*string {
	return s.Actions
}

func (s *GetRoleAclResponseBodyDataTable) GetName() *string {
	return s.Name
}

func (s *GetRoleAclResponseBodyDataTable) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetRoleAclResponseBodyDataTable) SetActions(v []*string) *GetRoleAclResponseBodyDataTable {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataTable) SetName(v string) *GetRoleAclResponseBodyDataTable {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataTable) SetSchemaName(v string) *GetRoleAclResponseBodyDataTable {
	s.SchemaName = &v
	return s
}

func (s *GetRoleAclResponseBodyDataTable) Validate() error {
	return dara.Validate(s)
}
