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
	ComputeResource *GetComputeResourceResponseBodyComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
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
	// example:
	//
	// InstanceMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// example:
	//
	// 1698286929333
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1107550004253538
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// My Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 16738
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1698286929333
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 1107550004253538
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// example:
	//
	// MyCs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 52660
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 1107550004253538:cn-beijing:holo:hgprecn-cn-x0r3oun4k001:testdb
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
	// example:
	//
	// hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
