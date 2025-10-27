// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *UpdateFlowResponseBody
	GetCreatedTime() *string
	SetDefinition(v string) *UpdateFlowResponseBody
	GetDefinition() *string
	SetDescription(v string) *UpdateFlowResponseBody
	GetDescription() *string
	SetEnvironment(v *UpdateFlowResponseBodyEnvironment) *UpdateFlowResponseBody
	GetEnvironment() *UpdateFlowResponseBodyEnvironment
	SetExternalStorageLocation(v string) *UpdateFlowResponseBody
	GetExternalStorageLocation() *string
	SetId(v string) *UpdateFlowResponseBody
	GetId() *string
	SetLastModifiedTime(v string) *UpdateFlowResponseBody
	GetLastModifiedTime() *string
	SetName(v string) *UpdateFlowResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateFlowResponseBody
	GetRequestId() *string
	SetRoleArn(v string) *UpdateFlowResponseBody
	GetRoleArn() *string
	SetType(v string) *UpdateFlowResponseBody
	GetType() *string
}

type UpdateFlowResponseBody struct {
	// The time when the flow was created.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The flow definition, which follows the FDL syntax standard. Considering compatibility, the system supports the two flow definition specifications.
	//
	// example:
	//
	// version: v1.0\\ntype: flow\\nname: test\\nsteps:\\n  - type: pass\\n    name: mypass
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The description of the flow.
	//
	// example:
	//
	// test definition
	Description *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Environment *UpdateFlowResponseBodyEnvironment `json:"Environment,omitempty" xml:"Environment,omitempty" type:"Struct"`
	// The path of the external storage.
	//
	// example:
	//
	// /path
	ExternalStorageLocation *string `json:"ExternalStorageLocation,omitempty" xml:"ExternalStorageLocation,omitempty"`
	// The unique ID of the flow.
	//
	// example:
	//
	// e589e092-e2c0-4dee-b306-3574ddfdddf5****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the flow was last modified.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The name of the flow.
	//
	// example:
	//
	// flow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud resource name (ARN) of the authorized role on which the execution of the flow relies. During the execution of the flow, the flow execution engine assumes the role to call API operations of relevant services.
	//
	// example:
	//
	// acs:ram::${accountID}:${role}
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The type of the flow.
	//
	// example:
	//
	// FDL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlowResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *UpdateFlowResponseBody) GetDefinition() *string {
	return s.Definition
}

func (s *UpdateFlowResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateFlowResponseBody) GetEnvironment() *UpdateFlowResponseBodyEnvironment {
	return s.Environment
}

func (s *UpdateFlowResponseBody) GetExternalStorageLocation() *string {
	return s.ExternalStorageLocation
}

func (s *UpdateFlowResponseBody) GetId() *string {
	return s.Id
}

func (s *UpdateFlowResponseBody) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *UpdateFlowResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFlowResponseBody) GetRoleArn() *string {
	return s.RoleArn
}

func (s *UpdateFlowResponseBody) GetType() *string {
	return s.Type
}

func (s *UpdateFlowResponseBody) SetCreatedTime(v string) *UpdateFlowResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *UpdateFlowResponseBody) SetDefinition(v string) *UpdateFlowResponseBody {
	s.Definition = &v
	return s
}

func (s *UpdateFlowResponseBody) SetDescription(v string) *UpdateFlowResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateFlowResponseBody) SetEnvironment(v *UpdateFlowResponseBodyEnvironment) *UpdateFlowResponseBody {
	s.Environment = v
	return s
}

func (s *UpdateFlowResponseBody) SetExternalStorageLocation(v string) *UpdateFlowResponseBody {
	s.ExternalStorageLocation = &v
	return s
}

func (s *UpdateFlowResponseBody) SetId(v string) *UpdateFlowResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateFlowResponseBody) SetLastModifiedTime(v string) *UpdateFlowResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *UpdateFlowResponseBody) SetName(v string) *UpdateFlowResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateFlowResponseBody) SetRequestId(v string) *UpdateFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlowResponseBody) SetRoleArn(v string) *UpdateFlowResponseBody {
	s.RoleArn = &v
	return s
}

func (s *UpdateFlowResponseBody) SetType(v string) *UpdateFlowResponseBody {
	s.Type = &v
	return s
}

func (s *UpdateFlowResponseBody) Validate() error {
	if s.Environment != nil {
		if err := s.Environment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateFlowResponseBodyEnvironment struct {
	Variables []*UpdateFlowResponseBodyEnvironmentVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s UpdateFlowResponseBodyEnvironment) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowResponseBodyEnvironment) GoString() string {
	return s.String()
}

func (s *UpdateFlowResponseBodyEnvironment) GetVariables() []*UpdateFlowResponseBodyEnvironmentVariables {
	return s.Variables
}

func (s *UpdateFlowResponseBodyEnvironment) SetVariables(v []*UpdateFlowResponseBodyEnvironmentVariables) *UpdateFlowResponseBodyEnvironment {
	s.Variables = v
	return s
}

func (s *UpdateFlowResponseBodyEnvironment) Validate() error {
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateFlowResponseBodyEnvironmentVariables struct {
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// key
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateFlowResponseBodyEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowResponseBodyEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *UpdateFlowResponseBodyEnvironmentVariables) GetDescription() *string {
	return s.Description
}

func (s *UpdateFlowResponseBodyEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *UpdateFlowResponseBodyEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *UpdateFlowResponseBodyEnvironmentVariables) SetDescription(v string) *UpdateFlowResponseBodyEnvironmentVariables {
	s.Description = &v
	return s
}

func (s *UpdateFlowResponseBodyEnvironmentVariables) SetName(v string) *UpdateFlowResponseBodyEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *UpdateFlowResponseBodyEnvironmentVariables) SetValue(v string) *UpdateFlowResponseBodyEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *UpdateFlowResponseBodyEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}
