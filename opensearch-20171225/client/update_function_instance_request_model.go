// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunctionInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateParameters(v []*UpdateFunctionInstanceRequestCreateParameters) *UpdateFunctionInstanceRequest
	GetCreateParameters() []*UpdateFunctionInstanceRequestCreateParameters
	SetCron(v string) *UpdateFunctionInstanceRequest
	GetCron() *string
	SetDescription(v string) *UpdateFunctionInstanceRequest
	GetDescription() *string
	SetUsageParameters(v []*UpdateFunctionInstanceRequestUsageParameters) *UpdateFunctionInstanceRequest
	GetUsageParameters() []*UpdateFunctionInstanceRequestUsageParameters
}

type UpdateFunctionInstanceRequest struct {
	// The parameters that are used to create the instance.
	//
	// example:
	//
	// {             "name": "title_field",             "value": "title"         }
	CreateParameters []*UpdateFunctionInstanceRequestCreateParameters `json:"createParameters,omitempty" xml:"createParameters,omitempty" type:"Repeated"`
	// The cron expression used to schedule periodic training, in the format of (Minutes Hours DayofMonth Month DayofWeek). The default value is empty, which indicates that no periodic training is performed. DayofWeek 0 indicates Sunday.
	//
	// example:
	//
	// "0 3 ? 	- 0,1,3,5"
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// test instance
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The parameters that are used.
	UsageParameters []*UpdateFunctionInstanceRequestUsageParameters `json:"usageParameters,omitempty" xml:"usageParameters,omitempty" type:"Repeated"`
}

func (s UpdateFunctionInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInstanceRequest) GetCreateParameters() []*UpdateFunctionInstanceRequestCreateParameters {
	return s.CreateParameters
}

func (s *UpdateFunctionInstanceRequest) GetCron() *string {
	return s.Cron
}

func (s *UpdateFunctionInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateFunctionInstanceRequest) GetUsageParameters() []*UpdateFunctionInstanceRequestUsageParameters {
	return s.UsageParameters
}

func (s *UpdateFunctionInstanceRequest) SetCreateParameters(v []*UpdateFunctionInstanceRequestCreateParameters) *UpdateFunctionInstanceRequest {
	s.CreateParameters = v
	return s
}

func (s *UpdateFunctionInstanceRequest) SetCron(v string) *UpdateFunctionInstanceRequest {
	s.Cron = &v
	return s
}

func (s *UpdateFunctionInstanceRequest) SetDescription(v string) *UpdateFunctionInstanceRequest {
	s.Description = &v
	return s
}

func (s *UpdateFunctionInstanceRequest) SetUsageParameters(v []*UpdateFunctionInstanceRequestUsageParameters) *UpdateFunctionInstanceRequest {
	s.UsageParameters = v
	return s
}

func (s *UpdateFunctionInstanceRequest) Validate() error {
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

type UpdateFunctionInstanceRequestCreateParameters struct {
	// The name of the parameter.
	//
	// example:
	//
	// title_field
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// title
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateFunctionInstanceRequestCreateParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionInstanceRequestCreateParameters) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInstanceRequestCreateParameters) GetName() *string {
	return s.Name
}

func (s *UpdateFunctionInstanceRequestCreateParameters) GetValue() *string {
	return s.Value
}

func (s *UpdateFunctionInstanceRequestCreateParameters) SetName(v string) *UpdateFunctionInstanceRequestCreateParameters {
	s.Name = &v
	return s
}

func (s *UpdateFunctionInstanceRequestCreateParameters) SetValue(v string) *UpdateFunctionInstanceRequestCreateParameters {
	s.Value = &v
	return s
}

func (s *UpdateFunctionInstanceRequestCreateParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateFunctionInstanceRequestUsageParameters struct {
	// The name of the parameter.
	//
	// example:
	//
	// allow_dict_id
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// 123
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateFunctionInstanceRequestUsageParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionInstanceRequestUsageParameters) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInstanceRequestUsageParameters) GetName() *string {
	return s.Name
}

func (s *UpdateFunctionInstanceRequestUsageParameters) GetValue() *string {
	return s.Value
}

func (s *UpdateFunctionInstanceRequestUsageParameters) SetName(v string) *UpdateFunctionInstanceRequestUsageParameters {
	s.Name = &v
	return s
}

func (s *UpdateFunctionInstanceRequestUsageParameters) SetValue(v string) *UpdateFunctionInstanceRequestUsageParameters {
	s.Value = &v
	return s
}

func (s *UpdateFunctionInstanceRequestUsageParameters) Validate() error {
	return dara.Validate(s)
}
