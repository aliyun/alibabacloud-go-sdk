// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComputeResource(v *GetComputeResourceResponseBodyComputeResource) *GetComputeResourceResponseBody
	GetComputeResource() *GetComputeResourceResponseBodyComputeResource
	SetRequestId(v string) *GetComputeResourceResponseBody
	GetRequestId() *string
}

type GetComputeResourceResponseBody struct {
	// The details of the computing resource.
	ComputeResource *GetComputeResourceResponseBodyComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	// The request ID. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 9252F32F-D855-549E-8898-61CF5A733050
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetComputeResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetComputeResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetComputeResourceResponseBody) GetComputeResource() *GetComputeResourceResponseBodyComputeResource {
	return s.ComputeResource
}

func (s *GetComputeResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetComputeResourceResponseBody) SetComputeResource(v *GetComputeResourceResponseBodyComputeResource) *GetComputeResourceResponseBody {
	s.ComputeResource = v
	return s
}

func (s *GetComputeResourceResponseBody) SetRequestId(v string) *GetComputeResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetComputeResourceResponseBody) Validate() error {
	if s.ComputeResource != nil {
		if err := s.ComputeResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetComputeResourceResponseBodyComputeResource struct {
	// The specific connection configuration details for the computing resource, including the connection address, access identity, and environment information. envType, which specifies the computing resource environment, is a property of this object. Valid values:
	//
	// 	- Dev
	//
	// 	- Prod Different types of computing resources have different attribute specifications under various configuration modes (ConnectionPropertiesMode).
	//
	// example:
	//
	// {
	//
	// 	"envType": "Prod",
	//
	// 	"regionId": "cn-beijing",
	//
	//     "instanceId": "hgprecn-cn-x0r3oun4k001",
	//
	//     "database": "testdb",
	//
	//     "securityProtocol": "authTypeNone",
	//
	//     "authType": "Executor",
	//
	//     "authIdentity": "1107550004253538"
	//
	// }
	ConnectionProperties interface{} `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// The addition category of the computing resource. Different types will have different subtypes, each with corresponding parameter constraints. For instance:
	//
	// 	- InstanceMode: Instance mode
	//
	// 	- UrlMode: Connection String Mode
	//
	// 	- CdhMode: CDH mode
	//
	// example:
	//
	// InstanceMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// The creation time, in timestamp format.
	//
	// example:
	//
	// 1698286929333
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the creator.
	//
	// example:
	//
	// 1107550004253538
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The description of the computing resource.
	//
	// example:
	//
	// My Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the computing resource.
	//
	// example:
	//
	// 16738
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The last modification time, in timestamp format.
	//
	// example:
	//
	// 1698286929333
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the modifier.
	//
	// example:
	//
	// 1107550004253538
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The name of the computing resource.
	//
	// example:
	//
	// MyCs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the workspace to which the computing resource belongs.
	//
	// example:
	//
	// 52660
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The business unique key for the computing resource. For example, the format for Hologres is ${tenantOwnerId}:${regionId}:${type}:${instanceId}:${database}.
	//
	// example:
	//
	// 1107550004253538:cn-beijing:holo:hgprecn-cn-x0r3oun4k001:testdb
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
	// The type of the computing resource.
	//
	// example:
	//
	// hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Specifies whether it is the default computing resource.
	//
	// example:
	//
	// true
	WhetherDefault *bool `json:"WhetherDefault,omitempty" xml:"WhetherDefault,omitempty"`
}

func (s GetComputeResourceResponseBodyComputeResource) String() string {
	return dara.Prettify(s)
}

func (s GetComputeResourceResponseBodyComputeResource) GoString() string {
	return s.String()
}

func (s *GetComputeResourceResponseBodyComputeResource) GetConnectionProperties() interface{} {
	return s.ConnectionProperties
}

func (s *GetComputeResourceResponseBodyComputeResource) GetConnectionPropertiesMode() *string {
	return s.ConnectionPropertiesMode
}

func (s *GetComputeResourceResponseBodyComputeResource) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetComputeResourceResponseBodyComputeResource) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetComputeResourceResponseBodyComputeResource) GetDescription() *string {
	return s.Description
}

func (s *GetComputeResourceResponseBodyComputeResource) GetId() *int64 {
	return s.Id
}

func (s *GetComputeResourceResponseBodyComputeResource) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetComputeResourceResponseBodyComputeResource) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *GetComputeResourceResponseBodyComputeResource) GetName() *string {
	return s.Name
}

func (s *GetComputeResourceResponseBodyComputeResource) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetComputeResourceResponseBodyComputeResource) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *GetComputeResourceResponseBodyComputeResource) GetType() *string {
	return s.Type
}

func (s *GetComputeResourceResponseBodyComputeResource) GetWhetherDefault() *bool {
	return s.WhetherDefault
}

func (s *GetComputeResourceResponseBodyComputeResource) SetConnectionProperties(v interface{}) *GetComputeResourceResponseBodyComputeResource {
	s.ConnectionProperties = v
	return s
}

func (s *GetComputeResourceResponseBodyComputeResource) SetConnectionPropertiesMode(v string) *GetComputeResourceResponseBodyComputeResource {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *GetComputeResourceResponseBodyComputeResource) SetCreateTime(v int64) *GetComputeResourceResponseBodyComputeResource {
	s.CreateTime = &v
	return s
}

func (s *GetComputeResourceResponseBodyComputeResource) SetCreateUser(v string) *GetComputeResourceResponseBodyComputeResource {
	s.CreateUser = &v
	return s
}

func (s *GetComputeResourceResponseBodyComputeResource) SetDescription(v string) *GetComputeResourceResponseBodyComputeResource {
	s.Description = &v
	return s
}

func (s *GetComputeResourceResponseBodyComputeResource) SetId(v int64) *GetComputeResourceResponseBodyComputeResource {
	s.Id = &v
	return s
}

func (s *GetComputeResourceResponseBodyComputeResource) SetModifyTime(v int64) *GetComputeResourceResponseBodyComputeResource {
	s.ModifyTime = &v
	return s
}

func (s *GetComputeResourceResponseBodyComputeResource) SetModifyUser(v string) *GetComputeResourceResponseBodyComputeResource {
	s.ModifyUser = &v
	return s
}

func (s *GetComputeResourceResponseBodyComputeResource) SetName(v string) *GetComputeResourceResponseBodyComputeResource {
	s.Name = &v
	return s
}

func (s *GetComputeResourceResponseBodyComputeResource) SetProjectId(v int64) *GetComputeResourceResponseBodyComputeResource {
	s.ProjectId = &v
	return s
}

func (s *GetComputeResourceResponseBodyComputeResource) SetQualifiedName(v string) *GetComputeResourceResponseBodyComputeResource {
	s.QualifiedName = &v
	return s
}

func (s *GetComputeResourceResponseBodyComputeResource) SetType(v string) *GetComputeResourceResponseBodyComputeResource {
	s.Type = &v
	return s
}

func (s *GetComputeResourceResponseBodyComputeResource) SetWhetherDefault(v bool) *GetComputeResourceResponseBodyComputeResource {
	s.WhetherDefault = &v
	return s
}

func (s *GetComputeResourceResponseBodyComputeResource) Validate() error {
	return dara.Validate(s)
}
