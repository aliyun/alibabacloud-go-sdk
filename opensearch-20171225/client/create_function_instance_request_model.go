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
	SetCron(v string) *CreateFunctionInstanceRequest
	GetCron() *string
	SetDescription(v string) *CreateFunctionInstanceRequest
	GetDescription() *string
	SetFunctionType(v string) *CreateFunctionInstanceRequest
	GetFunctionType() *string
	SetInstanceName(v string) *CreateFunctionInstanceRequest
	GetInstanceName() *string
	SetModelType(v string) *CreateFunctionInstanceRequest
	GetModelType() *string
	SetUsageParameters(v []*CreateFunctionInstanceRequestUsageParameters) *CreateFunctionInstanceRequest
	GetUsageParameters() []*CreateFunctionInstanceRequestUsageParameters
}

type CreateFunctionInstanceRequest struct {
	// The parameters used to create the instance.
	//
	// example:
	//
	// [   { "name": "param1", "value": "val1"   } ]
	CreateParameters []*CreateFunctionInstanceRequestCreateParameters `json:"createParameters,omitempty" xml:"createParameters,omitempty" type:"Repeated"`
	// The CRON expression used to schedule periodic training, in the format of Minutes Hours DayofMonth Month DayofWeek. The default value is empty, which specifies that no periodic training is performed. A value of 0 for DayofWeek specifies Sunday.
	//
	// example:
	//
	// 0 0 ? 	- 0,1,2,3,4,5,6
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// The description.
	//
	// example:
	//
	// test instance
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The feature type.
	//
	// 	- Default value: PAAS. Training is required before you can use the feature.
	//
	// example:
	//
	// PAAS
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	// The instance name. The name must be 1 to 30 characters in length and can contain letters, digits, and underscores (_). The name is case-sensitive and must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ctr_test
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The model type. The value varies based on the model.
	//
	// 	- Click-through rate (CTR) model: tf_checkpoint
	//
	// 	- Popularity model: pop
	//
	// 	- Category model: offline_inference
	//
	// 	- Hotword model: offline_inference
	//
	// 	- Hint model: offline_inference
	//
	// 	- Hotword model for real-time top searches: near_realtime
	//
	// 	- Personalized hint model: near_realtime
	//
	// 	- Drop-down suggestion model: offline_inference
	//
	// 	- Tokenization model: text
	//
	// 	- Term weight model: tf_checkpoint
	//
	// 	- Synonym model: offline_inference
	//
	// This parameter is required.
	//
	// example:
	//
	// tf_checkpoint
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The parameters used to use the instance.
	UsageParameters []*CreateFunctionInstanceRequestUsageParameters `json:"usageParameters,omitempty" xml:"usageParameters,omitempty" type:"Repeated"`
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

func (s *CreateFunctionInstanceRequest) GetCron() *string {
	return s.Cron
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

func (s *CreateFunctionInstanceRequest) GetUsageParameters() []*CreateFunctionInstanceRequestUsageParameters {
	return s.UsageParameters
}

func (s *CreateFunctionInstanceRequest) SetCreateParameters(v []*CreateFunctionInstanceRequestCreateParameters) *CreateFunctionInstanceRequest {
	s.CreateParameters = v
	return s
}

func (s *CreateFunctionInstanceRequest) SetCron(v string) *CreateFunctionInstanceRequest {
	s.Cron = &v
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

func (s *CreateFunctionInstanceRequest) SetUsageParameters(v []*CreateFunctionInstanceRequestUsageParameters) *CreateFunctionInstanceRequest {
	s.UsageParameters = v
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
	if s.UsageParameters != nil {
		for _, item := range s.UsageParameters {
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
	// title_field
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// title
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

type CreateFunctionInstanceRequestUsageParameters struct {
	// The parameter name.
	//
	// example:
	//
	// allow_dict_id
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// 123
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateFunctionInstanceRequestUsageParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionInstanceRequestUsageParameters) GoString() string {
	return s.String()
}

func (s *CreateFunctionInstanceRequestUsageParameters) GetName() *string {
	return s.Name
}

func (s *CreateFunctionInstanceRequestUsageParameters) GetValue() *string {
	return s.Value
}

func (s *CreateFunctionInstanceRequestUsageParameters) SetName(v string) *CreateFunctionInstanceRequestUsageParameters {
	s.Name = &v
	return s
}

func (s *CreateFunctionInstanceRequestUsageParameters) SetValue(v string) *CreateFunctionInstanceRequestUsageParameters {
	s.Value = &v
	return s
}

func (s *CreateFunctionInstanceRequestUsageParameters) Validate() error {
	return dara.Validate(s)
}
