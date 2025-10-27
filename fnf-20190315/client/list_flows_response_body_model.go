// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlows(v []*ListFlowsResponseBodyFlows) *ListFlowsResponseBody
	GetFlows() []*ListFlowsResponseBodyFlows
	SetNextToken(v string) *ListFlowsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListFlowsResponseBody
	GetRequestId() *string
}

type ListFlowsResponseBody struct {
	// The details of flows.
	Flows []*ListFlowsResponseBodyFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	// The start key for the next query. This parameter is not returned if all results have been returned.
	//
	// example:
	//
	// flow_nextxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFlowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBody) GetFlows() []*ListFlowsResponseBodyFlows {
	return s.Flows
}

func (s *ListFlowsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFlowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlowsResponseBody) SetFlows(v []*ListFlowsResponseBodyFlows) *ListFlowsResponseBody {
	s.Flows = v
	return s
}

func (s *ListFlowsResponseBody) SetNextToken(v string) *ListFlowsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFlowsResponseBody) SetRequestId(v string) *ListFlowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowsResponseBody) Validate() error {
	if s.Flows != nil {
		for _, item := range s.Flows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFlowsResponseBodyFlows struct {
	// The time when the flow was created.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The definition of the flow. The definition must comply with the Flow Definition Language (FDL) syntax.
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
	Description *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Environment *ListFlowsResponseBodyFlowsEnvironment `json:"Environment,omitempty" xml:"Environment,omitempty" type:"Struct"`
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
	// e589e092-e2c0-4dee-b306-3574ddf5****
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
	// The Alibaba Cloud resource name (ARN) of the specified Resource Access Management (RAM) role that Serverless Workflow assumes to invoke resources when the flow is executed.
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

func (s ListFlowsResponseBodyFlows) String() string {
	return dara.Prettify(s)
}

func (s ListFlowsResponseBodyFlows) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBodyFlows) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListFlowsResponseBodyFlows) GetDefinition() *string {
	return s.Definition
}

func (s *ListFlowsResponseBodyFlows) GetDescription() *string {
	return s.Description
}

func (s *ListFlowsResponseBodyFlows) GetEnvironment() *ListFlowsResponseBodyFlowsEnvironment {
	return s.Environment
}

func (s *ListFlowsResponseBodyFlows) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *ListFlowsResponseBodyFlows) GetId() *string {
	return s.Id
}

func (s *ListFlowsResponseBodyFlows) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *ListFlowsResponseBodyFlows) GetName() *string {
	return s.Name
}

func (s *ListFlowsResponseBodyFlows) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ListFlowsResponseBodyFlows) GetType() *string {
	return s.Type
}

func (s *ListFlowsResponseBodyFlows) SetCreatedTime(v string) *ListFlowsResponseBodyFlows {
	s.CreatedTime = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetDefinition(v string) *ListFlowsResponseBodyFlows {
	s.Definition = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetDescription(v string) *ListFlowsResponseBodyFlows {
	s.Description = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetEnvironment(v *ListFlowsResponseBodyFlowsEnvironment) *ListFlowsResponseBodyFlows {
	s.Environment = v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetExecutionMode(v string) *ListFlowsResponseBodyFlows {
	s.ExecutionMode = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetId(v string) *ListFlowsResponseBodyFlows {
	s.Id = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetLastModifiedTime(v string) *ListFlowsResponseBodyFlows {
	s.LastModifiedTime = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetName(v string) *ListFlowsResponseBodyFlows {
	s.Name = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetRoleArn(v string) *ListFlowsResponseBodyFlows {
	s.RoleArn = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetType(v string) *ListFlowsResponseBodyFlows {
	s.Type = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) Validate() error {
	if s.Environment != nil {
		if err := s.Environment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFlowsResponseBodyFlowsEnvironment struct {
	Variables []*ListFlowsResponseBodyFlowsEnvironmentVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListFlowsResponseBodyFlowsEnvironment) String() string {
	return dara.Prettify(s)
}

func (s ListFlowsResponseBodyFlowsEnvironment) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBodyFlowsEnvironment) GetVariables() []*ListFlowsResponseBodyFlowsEnvironmentVariables {
	return s.Variables
}

func (s *ListFlowsResponseBodyFlowsEnvironment) SetVariables(v []*ListFlowsResponseBodyFlowsEnvironmentVariables) *ListFlowsResponseBodyFlowsEnvironment {
	s.Variables = v
	return s
}

func (s *ListFlowsResponseBodyFlowsEnvironment) Validate() error {
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

type ListFlowsResponseBodyFlowsEnvironmentVariables struct {
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

func (s ListFlowsResponseBodyFlowsEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s ListFlowsResponseBodyFlowsEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBodyFlowsEnvironmentVariables) GetDescription() *string {
	return s.Description
}

func (s *ListFlowsResponseBodyFlowsEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *ListFlowsResponseBodyFlowsEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *ListFlowsResponseBodyFlowsEnvironmentVariables) SetDescription(v string) *ListFlowsResponseBodyFlowsEnvironmentVariables {
	s.Description = &v
	return s
}

func (s *ListFlowsResponseBodyFlowsEnvironmentVariables) SetName(v string) *ListFlowsResponseBodyFlowsEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *ListFlowsResponseBodyFlowsEnvironmentVariables) SetValue(v string) *ListFlowsResponseBodyFlowsEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *ListFlowsResponseBodyFlowsEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}
