// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListRolesResponseBodyData) *ListRolesResponseBody
	GetData() *ListRolesResponseBodyData
	SetRequestId(v string) *ListRolesResponseBody
	GetRequestId() *string
}

type ListRolesResponseBody struct {
	// The returned data.
	Data *ListRolesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0a06dfe716686526652451361e80ae
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) GetData() *ListRolesResponseBodyData {
	return s.Data
}

func (s *ListRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRolesResponseBody) SetData(v *ListRolesResponseBodyData) *ListRolesResponseBody {
	s.Data = v
	return s
}

func (s *ListRolesResponseBody) SetRequestId(v string) *ListRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRolesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRolesResponseBodyData struct {
	// The MaxCompute project-level roles.
	Roles []*ListRolesResponseBodyDataRoles `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
}

func (s ListRolesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyData) GetRoles() []*ListRolesResponseBodyDataRoles {
	return s.Roles
}

func (s *ListRolesResponseBodyData) SetRoles(v []*ListRolesResponseBodyDataRoles) *ListRolesResponseBodyData {
	s.Roles = v
	return s
}

func (s *ListRolesResponseBodyData) Validate() error {
	if s.Roles != nil {
		for _, item := range s.Roles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRolesResponseBodyDataRoles struct {
	// The ACL-based permissions that are granted to the role.
	Acl *ListRolesResponseBodyDataRolesAcl `json:"acl,omitempty" xml:"acl,omitempty" type:"Struct"`
	// The name of the role.
	//
	// example:
	//
	// roleA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The policy that is attached to the role.
	//
	// example:
	//
	// {
	//
	//       "Statement": [
	//
	//             {
	//
	//                   "Action": [
	//
	//                         "odps:*"
	//
	//                   ],
	//
	//                   "Effect": "Allow",
	//
	//                   "Resource": [
	//
	//                         "acs:odps:*:projects/{projectname}/authorization/packages"
	//
	//                   ]
	//
	//             }
	//
	//       ],
	//
	//       "Version": "1"
	//
	// }
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// The type of the role.
	//
	// example:
	//
	// admin
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListRolesResponseBodyDataRoles) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyDataRoles) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRoles) GetAcl() *ListRolesResponseBodyDataRolesAcl {
	return s.Acl
}

func (s *ListRolesResponseBodyDataRoles) GetName() *string {
	return s.Name
}

func (s *ListRolesResponseBodyDataRoles) GetPolicy() *string {
	return s.Policy
}

func (s *ListRolesResponseBodyDataRoles) GetType() *string {
	return s.Type
}

func (s *ListRolesResponseBodyDataRoles) SetAcl(v *ListRolesResponseBodyDataRolesAcl) *ListRolesResponseBodyDataRoles {
	s.Acl = v
	return s
}

func (s *ListRolesResponseBodyDataRoles) SetName(v string) *ListRolesResponseBodyDataRoles {
	s.Name = &v
	return s
}

func (s *ListRolesResponseBodyDataRoles) SetPolicy(v string) *ListRolesResponseBodyDataRoles {
	s.Policy = &v
	return s
}

func (s *ListRolesResponseBodyDataRoles) SetType(v string) *ListRolesResponseBodyDataRoles {
	s.Type = &v
	return s
}

func (s *ListRolesResponseBodyDataRoles) Validate() error {
	if s.Acl != nil {
		if err := s.Acl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRolesResponseBodyDataRolesAcl struct {
	// The function.
	Function []*ListRolesResponseBodyDataRolesAclFunction `json:"function,omitempty" xml:"function,omitempty" type:"Repeated"`
	// The instance.
	Instance []*ListRolesResponseBodyDataRolesAclInstance `json:"instance,omitempty" xml:"instance,omitempty" type:"Repeated"`
	// The package.
	Package []*ListRolesResponseBodyDataRolesAclPackage `json:"package,omitempty" xml:"package,omitempty" type:"Repeated"`
	// The project.
	Project []*ListRolesResponseBodyDataRolesAclProject `json:"project,omitempty" xml:"project,omitempty" type:"Repeated"`
	// The resource.
	Resource []*ListRolesResponseBodyDataRolesAclResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Repeated"`
	// The table.
	Table []*ListRolesResponseBodyDataRolesAclTable `json:"table,omitempty" xml:"table,omitempty" type:"Repeated"`
}

func (s ListRolesResponseBodyDataRolesAcl) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAcl) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAcl) GetFunction() []*ListRolesResponseBodyDataRolesAclFunction {
	return s.Function
}

func (s *ListRolesResponseBodyDataRolesAcl) GetInstance() []*ListRolesResponseBodyDataRolesAclInstance {
	return s.Instance
}

func (s *ListRolesResponseBodyDataRolesAcl) GetPackage() []*ListRolesResponseBodyDataRolesAclPackage {
	return s.Package
}

func (s *ListRolesResponseBodyDataRolesAcl) GetProject() []*ListRolesResponseBodyDataRolesAclProject {
	return s.Project
}

func (s *ListRolesResponseBodyDataRolesAcl) GetResource() []*ListRolesResponseBodyDataRolesAclResource {
	return s.Resource
}

func (s *ListRolesResponseBodyDataRolesAcl) GetTable() []*ListRolesResponseBodyDataRolesAclTable {
	return s.Table
}

func (s *ListRolesResponseBodyDataRolesAcl) SetFunction(v []*ListRolesResponseBodyDataRolesAclFunction) *ListRolesResponseBodyDataRolesAcl {
	s.Function = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAcl) SetInstance(v []*ListRolesResponseBodyDataRolesAclInstance) *ListRolesResponseBodyDataRolesAcl {
	s.Instance = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAcl) SetPackage(v []*ListRolesResponseBodyDataRolesAclPackage) *ListRolesResponseBodyDataRolesAcl {
	s.Package = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAcl) SetProject(v []*ListRolesResponseBodyDataRolesAclProject) *ListRolesResponseBodyDataRolesAcl {
	s.Project = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAcl) SetResource(v []*ListRolesResponseBodyDataRolesAclResource) *ListRolesResponseBodyDataRolesAcl {
	s.Resource = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAcl) SetTable(v []*ListRolesResponseBodyDataRolesAclTable) *ListRolesResponseBodyDataRolesAcl {
	s.Table = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAcl) Validate() error {
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

type ListRolesResponseBodyDataRolesAclFunction struct {
	// The operations that were performed on the function.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the function.
	//
	// example:
	//
	// functionA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclFunction) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclFunction) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclFunction) GetActions() []*string {
	return s.Actions
}

func (s *ListRolesResponseBodyDataRolesAclFunction) GetName() *string {
	return s.Name
}

func (s *ListRolesResponseBodyDataRolesAclFunction) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclFunction {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclFunction) SetName(v string) *ListRolesResponseBodyDataRolesAclFunction {
	s.Name = &v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclFunction) Validate() error {
	return dara.Validate(s)
}

type ListRolesResponseBodyDataRolesAclInstance struct {
	// The operations that were performed on the instance.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the instance.
	//
	// example:
	//
	// instanceA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclInstance) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclInstance) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclInstance) GetActions() []*string {
	return s.Actions
}

func (s *ListRolesResponseBodyDataRolesAclInstance) GetName() *string {
	return s.Name
}

func (s *ListRolesResponseBodyDataRolesAclInstance) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclInstance {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclInstance) SetName(v string) *ListRolesResponseBodyDataRolesAclInstance {
	s.Name = &v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclInstance) Validate() error {
	return dara.Validate(s)
}

type ListRolesResponseBodyDataRolesAclPackage struct {
	// The operations that were performed on the package.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the package.
	//
	// example:
	//
	// packageA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclPackage) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclPackage) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclPackage) GetActions() []*string {
	return s.Actions
}

func (s *ListRolesResponseBodyDataRolesAclPackage) GetName() *string {
	return s.Name
}

func (s *ListRolesResponseBodyDataRolesAclPackage) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclPackage {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclPackage) SetName(v string) *ListRolesResponseBodyDataRolesAclPackage {
	s.Name = &v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclPackage) Validate() error {
	return dara.Validate(s)
}

type ListRolesResponseBodyDataRolesAclProject struct {
	// The operations that were performed on the project.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the MaxCompute project.
	//
	// example:
	//
	// projectA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclProject) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclProject) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclProject) GetActions() []*string {
	return s.Actions
}

func (s *ListRolesResponseBodyDataRolesAclProject) GetName() *string {
	return s.Name
}

func (s *ListRolesResponseBodyDataRolesAclProject) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclProject {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclProject) SetName(v string) *ListRolesResponseBodyDataRolesAclProject {
	s.Name = &v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclProject) Validate() error {
	return dara.Validate(s)
}

type ListRolesResponseBodyDataRolesAclResource struct {
	// The operations that were performed on the resource.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the resource.
	//
	// example:
	//
	// resourceA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclResource) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclResource) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclResource) GetActions() []*string {
	return s.Actions
}

func (s *ListRolesResponseBodyDataRolesAclResource) GetName() *string {
	return s.Name
}

func (s *ListRolesResponseBodyDataRolesAclResource) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclResource {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclResource) SetName(v string) *ListRolesResponseBodyDataRolesAclResource {
	s.Name = &v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclResource) Validate() error {
	return dara.Validate(s)
}

type ListRolesResponseBodyDataRolesAclTable struct {
	// The operations that were performed on the table.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the table.
	//
	// example:
	//
	// tableA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclTable) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclTable) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclTable) GetActions() []*string {
	return s.Actions
}

func (s *ListRolesResponseBodyDataRolesAclTable) GetName() *string {
	return s.Name
}

func (s *ListRolesResponseBodyDataRolesAclTable) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclTable {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclTable) SetName(v string) *ListRolesResponseBodyDataRolesAclTable {
	s.Name = &v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclTable) Validate() error {
	return dara.Validate(s)
}
