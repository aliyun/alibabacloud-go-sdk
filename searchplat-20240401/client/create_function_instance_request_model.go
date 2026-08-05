// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFunctionInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateParameters(v []*CreateFunctionInstanceRequestCreateParameters) *CreateFunctionInstanceRequest
	GetCreateParameters() []*CreateFunctionInstanceRequestCreateParameters
	SetDescription(v string) *CreateFunctionInstanceRequest
	GetDescription() *string
	SetFunctionType(v string) *CreateFunctionInstanceRequest
	GetFunctionType() *string
	SetInstanceName(v string) *CreateFunctionInstanceRequest
	GetInstanceName() *string
	SetModelType(v string) *CreateFunctionInstanceRequest
	GetModelType() *string
}

type CreateFunctionInstanceRequest struct {
	// The creation parameters.
	CreateParameters []*CreateFunctionInstanceRequestCreateParameters `json:"createParameters,omitempty" xml:"createParameters,omitempty" type:"Repeated"`
	// The instance description.
	//
	// example:
	//
	// desc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The configuration type. Valid values:
	//
	// - PAAS
	//
	// - SAAS.
	//
	// example:
	//
	// PAAS
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	// The configuration or model name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The service ID. Valid values:
	//
	// - ops-query-analyze-nl2sql-001
	//
	// - ops-embedding-dim-reduction-001: vector dimension reduction.
	//
	// This parameter is required.
	//
	// example:
	//
	// ops-query-analyze-nl2sql-001
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
}

func (s CreateFunctionInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionInstanceRequest) GetCreateParameters() []*CreateFunctionInstanceRequestCreateParameters {
	return s.CreateParameters
}

func (s *CreateFunctionInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFunctionInstanceRequest) GetFunctionType() *string {
	return s.FunctionType
}

func (s *CreateFunctionInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateFunctionInstanceRequest) GetModelType() *string {
	return s.ModelType
}

func (s *CreateFunctionInstanceRequest) SetCreateParameters(v []*CreateFunctionInstanceRequestCreateParameters) *CreateFunctionInstanceRequest {
	s.CreateParameters = v
	return s
}

func (s *CreateFunctionInstanceRequest) SetDescription(v string) *CreateFunctionInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateFunctionInstanceRequest) SetFunctionType(v string) *CreateFunctionInstanceRequest {
	s.FunctionType = &v
	return s
}

func (s *CreateFunctionInstanceRequest) SetInstanceName(v string) *CreateFunctionInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateFunctionInstanceRequest) SetModelType(v string) *CreateFunctionInstanceRequest {
	s.ModelType = &v
	return s
}

func (s *CreateFunctionInstanceRequest) Validate() error {
	if s.CreateParameters != nil {
		for _, item := range s.CreateParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateFunctionInstanceRequestCreateParameters struct {
	// The parameter name.
	//
	// example:
	//
	// config
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// {\\"DDL\\": [{\\"table\\": \\"schools\\",\\"columns\\": [{\\"column\\": \\"class\\",\\"column_des\\": \\"班级\\",\\"type\\": \\"str\\",\\"example\\": [\\"高一3班\\",\\"火箭班\\"],\\"value_mapping\\": {}},{\\"column\\": \\"school\\",\\"column_des\\": \\"学校\\",\\"type\\": \\"str\\",\\"example\\": [\\"清华大学\\",\\"北京大学\\"],\\"value_mapping\\": {}}]},{\\"table\\": \\"students\\",\\"columns\\": [{\\"column\\": \\"name\\",\\"column_des\\": \\"姓名\\",\\"type\\": \\"int\\",\\"example\\": [10002,100001],\\"value_mapping\\": [[10002,100001],[\\"张三\\",\\"李四\\"]]}]}],\\"foreign keys\\":[\\"table.column_1=table2.column_2\\",\\"table.column_1=table2.column_2\\"],\\"UDF\\": [[\\"初始节点\\",\\"aa\\"],[\\" (sub_action >100095 or action = 0001) and station =100001\\",\\"bbb\\"]],\\"Fewshot\\": [{\\"query\\": \\"叫张三的学生有多少\\",\\"sql\\": \\"SELECT COUNT(*) FROM students WHERE name = 10002\\"}]}
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateFunctionInstanceRequestCreateParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionInstanceRequestCreateParameters) GoString() string {
	return s.String()
}

func (s *CreateFunctionInstanceRequestCreateParameters) GetName() *string {
	return s.Name
}

func (s *CreateFunctionInstanceRequestCreateParameters) GetValue() *string {
	return s.Value
}

func (s *CreateFunctionInstanceRequestCreateParameters) SetName(v string) *CreateFunctionInstanceRequestCreateParameters {
	s.Name = &v
	return s
}

func (s *CreateFunctionInstanceRequestCreateParameters) SetValue(v string) *CreateFunctionInstanceRequestCreateParameters {
	s.Value = &v
	return s
}

func (s *CreateFunctionInstanceRequestCreateParameters) Validate() error {
	return dara.Validate(s)
}
