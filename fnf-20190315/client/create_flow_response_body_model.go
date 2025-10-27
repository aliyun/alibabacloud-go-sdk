// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *CreateFlowResponseBody
	GetCreatedTime() *string
	SetDefinition(v string) *CreateFlowResponseBody
	GetDefinition() *string
	SetDescription(v string) *CreateFlowResponseBody
	GetDescription() *string
	SetEnvironment(v *CreateFlowResponseBodyEnvironment) *CreateFlowResponseBody
	GetEnvironment() *CreateFlowResponseBodyEnvironment
	SetExecutionMode(v string) *CreateFlowResponseBody
	GetExecutionMode() *string
	SetId(v string) *CreateFlowResponseBody
	GetId() *string
	SetLastModifiedTime(v string) *CreateFlowResponseBody
	GetLastModifiedTime() *string
	SetName(v string) *CreateFlowResponseBody
	GetName() *string
	SetRequestId(v string) *CreateFlowResponseBody
	GetRequestId() *string
	SetRoleArn(v string) *CreateFlowResponseBody
	GetRoleArn() *string
	SetType(v string) *CreateFlowResponseBody
	GetType() *string
}

type CreateFlowResponseBody struct {
	// The time when the flow was created.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// Considering compatibility, the system supports two flow definition specifications.
	//
	// example:
	//
	// version: v1.0\\ntype: flow\\nname: test\\nsteps:\\n - type: pass\\n name: mypass
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The description of the flow.
	//
	// example:
	//
	// test flow
	Description *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Environment *CreateFlowResponseBodyEnvironment `json:"Environment,omitempty" xml:"Environment,omitempty" type:"Struct"`
	// The execution mode. Valid values: Express and Standard. Considering compatibility, an empty string is equivalent to the Standard execution mode.
	//
	// example:
	//
	// Standard
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
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
	// The request ID. Each time an `HTTP status code` is returned, Serverless Workflow returns a value for the parameter.
	//
	// example:
	//
	// testRequestID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud resource name (ARN) of the authorized role on which the execution of the flow relies. During the execution of the flow, CloudFlow assumes the role to call API operations of relevant services.
	//
	// example:
	//
	// acs:ram:${region}:${accountID}:${role}
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The type of the flow.
	//
	// Valid value:
	//
	// 	- FDL
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// FDL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *CreateFlowResponseBody) GetDefinition() *string {
	return s.Definition
}

func (s *CreateFlowResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateFlowResponseBody) GetEnvironment() *CreateFlowResponseBodyEnvironment {
	return s.Environment
}

func (s *CreateFlowResponseBody) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *CreateFlowResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateFlowResponseBody) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *CreateFlowResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFlowResponseBody) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CreateFlowResponseBody) GetType() *string {
	return s.Type
}

func (s *CreateFlowResponseBody) SetCreatedTime(v string) *CreateFlowResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateFlowResponseBody) SetDefinition(v string) *CreateFlowResponseBody {
	s.Definition = &v
	return s
}

func (s *CreateFlowResponseBody) SetDescription(v string) *CreateFlowResponseBody {
	s.Description = &v
	return s
}

func (s *CreateFlowResponseBody) SetEnvironment(v *CreateFlowResponseBodyEnvironment) *CreateFlowResponseBody {
	s.Environment = v
	return s
}

func (s *CreateFlowResponseBody) SetExecutionMode(v string) *CreateFlowResponseBody {
	s.ExecutionMode = &v
	return s
}

func (s *CreateFlowResponseBody) SetId(v string) *CreateFlowResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFlowResponseBody) SetLastModifiedTime(v string) *CreateFlowResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *CreateFlowResponseBody) SetName(v string) *CreateFlowResponseBody {
	s.Name = &v
	return s
}

func (s *CreateFlowResponseBody) SetRequestId(v string) *CreateFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowResponseBody) SetRoleArn(v string) *CreateFlowResponseBody {
	s.RoleArn = &v
	return s
}

func (s *CreateFlowResponseBody) SetType(v string) *CreateFlowResponseBody {
	s.Type = &v
	return s
}

func (s *CreateFlowResponseBody) Validate() error {
	if s.Environment != nil {
		if err := s.Environment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateFlowResponseBodyEnvironment struct {
	Variables []*CreateFlowResponseBodyEnvironmentVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s CreateFlowResponseBodyEnvironment) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowResponseBodyEnvironment) GoString() string {
	return s.String()
}

func (s *CreateFlowResponseBodyEnvironment) GetVariables() []*CreateFlowResponseBodyEnvironmentVariables {
	return s.Variables
}

func (s *CreateFlowResponseBodyEnvironment) SetVariables(v []*CreateFlowResponseBodyEnvironmentVariables) *CreateFlowResponseBodyEnvironment {
	s.Variables = v
	return s
}

func (s *CreateFlowResponseBodyEnvironment) Validate() error {
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

type CreateFlowResponseBodyEnvironmentVariables struct {
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

func (s CreateFlowResponseBodyEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowResponseBodyEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *CreateFlowResponseBodyEnvironmentVariables) GetDescription() *string {
	return s.Description
}

func (s *CreateFlowResponseBodyEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *CreateFlowResponseBodyEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *CreateFlowResponseBodyEnvironmentVariables) SetDescription(v string) *CreateFlowResponseBodyEnvironmentVariables {
	s.Description = &v
	return s
}

func (s *CreateFlowResponseBodyEnvironmentVariables) SetName(v string) *CreateFlowResponseBodyEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *CreateFlowResponseBodyEnvironmentVariables) SetValue(v string) *CreateFlowResponseBodyEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *CreateFlowResponseBodyEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}
