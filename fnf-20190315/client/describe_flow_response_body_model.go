// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *DescribeFlowResponseBody
	GetCreatedTime() *string
	SetDefinition(v string) *DescribeFlowResponseBody
	GetDefinition() *string
	SetDescription(v string) *DescribeFlowResponseBody
	GetDescription() *string
	SetEnvironment(v *DescribeFlowResponseBodyEnvironment) *DescribeFlowResponseBody
	GetEnvironment() *DescribeFlowResponseBodyEnvironment
	SetExecutionMode(v string) *DescribeFlowResponseBody
	GetExecutionMode() *string
	SetId(v string) *DescribeFlowResponseBody
	GetId() *string
	SetLastModifiedTime(v string) *DescribeFlowResponseBody
	GetLastModifiedTime() *string
	SetName(v string) *DescribeFlowResponseBody
	GetName() *string
	SetRequestId(v string) *DescribeFlowResponseBody
	GetRequestId() *string
	SetRoleArn(v string) *DescribeFlowResponseBody
	GetRoleArn() *string
	SetType(v string) *DescribeFlowResponseBody
	GetType() *string
}

type DescribeFlowResponseBody struct {
	// The time when the flow was created.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The definition of the workflow. The definition must comply with the flow definition language (FDL) syntax. Considering compatibility, the system supports the flow definition specifications of the old version and new version.
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
	Description *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Environment *DescribeFlowResponseBodyEnvironment `json:"Environment,omitempty" xml:"Environment,omitempty" type:"Struct"`
	// The execution mode or the enumeration type. Valid values: Express and Standard. A value of Standard indicates an empty string.
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
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud resource name (ARN) of the authorized role on which the execution of the flow relies. During the execution of the flow, CloudFlow assumes the role to call API operations of relevant services.
	//
	// example:
	//
	// acs:ram::${accountID}:${role}
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The type of the workflow.
	//
	// example:
	//
	// FDL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeFlowResponseBody) GetDefinition() *string {
	return s.Definition
}

func (s *DescribeFlowResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeFlowResponseBody) GetEnvironment() *DescribeFlowResponseBodyEnvironment {
	return s.Environment
}

func (s *DescribeFlowResponseBody) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *DescribeFlowResponseBody) GetId() *string {
	return s.Id
}

func (s *DescribeFlowResponseBody) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *DescribeFlowResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFlowResponseBody) GetRoleArn() *string {
	return s.RoleArn
}

func (s *DescribeFlowResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeFlowResponseBody) SetCreatedTime(v string) *DescribeFlowResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeFlowResponseBody) SetDefinition(v string) *DescribeFlowResponseBody {
	s.Definition = &v
	return s
}

func (s *DescribeFlowResponseBody) SetDescription(v string) *DescribeFlowResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeFlowResponseBody) SetEnvironment(v *DescribeFlowResponseBodyEnvironment) *DescribeFlowResponseBody {
	s.Environment = v
	return s
}

func (s *DescribeFlowResponseBody) SetExecutionMode(v string) *DescribeFlowResponseBody {
	s.ExecutionMode = &v
	return s
}

func (s *DescribeFlowResponseBody) SetId(v string) *DescribeFlowResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeFlowResponseBody) SetLastModifiedTime(v string) *DescribeFlowResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *DescribeFlowResponseBody) SetName(v string) *DescribeFlowResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeFlowResponseBody) SetRequestId(v string) *DescribeFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowResponseBody) SetRoleArn(v string) *DescribeFlowResponseBody {
	s.RoleArn = &v
	return s
}

func (s *DescribeFlowResponseBody) SetType(v string) *DescribeFlowResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeFlowResponseBody) Validate() error {
	if s.Environment != nil {
		if err := s.Environment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFlowResponseBodyEnvironment struct {
	Variables []*DescribeFlowResponseBodyEnvironmentVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s DescribeFlowResponseBodyEnvironment) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowResponseBodyEnvironment) GoString() string {
	return s.String()
}

func (s *DescribeFlowResponseBodyEnvironment) GetVariables() []*DescribeFlowResponseBodyEnvironmentVariables {
	return s.Variables
}

func (s *DescribeFlowResponseBodyEnvironment) SetVariables(v []*DescribeFlowResponseBodyEnvironmentVariables) *DescribeFlowResponseBodyEnvironment {
	s.Variables = v
	return s
}

func (s *DescribeFlowResponseBodyEnvironment) Validate() error {
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

type DescribeFlowResponseBodyEnvironmentVariables struct {
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

func (s DescribeFlowResponseBodyEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowResponseBodyEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *DescribeFlowResponseBodyEnvironmentVariables) GetDescription() *string {
	return s.Description
}

func (s *DescribeFlowResponseBodyEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *DescribeFlowResponseBodyEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *DescribeFlowResponseBodyEnvironmentVariables) SetDescription(v string) *DescribeFlowResponseBodyEnvironmentVariables {
	s.Description = &v
	return s
}

func (s *DescribeFlowResponseBodyEnvironmentVariables) SetName(v string) *DescribeFlowResponseBodyEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *DescribeFlowResponseBodyEnvironmentVariables) SetValue(v string) *DescribeFlowResponseBodyEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *DescribeFlowResponseBodyEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}
