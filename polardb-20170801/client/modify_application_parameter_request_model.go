// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationParameterRequest
	GetApplicationId() *string
	SetParameterName(v string) *ModifyApplicationParameterRequest
	GetParameterName() *string
	SetParameterValue(v string) *ModifyApplicationParameterRequest
	GetParameterValue() *string
	SetParameters(v []*ModifyApplicationParameterRequestParameters) *ModifyApplicationParameterRequest
	GetParameters() []*ModifyApplicationParameterRequestParameters
}

type ModifyApplicationParameterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// name
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// example:
	//
	// value
	ParameterValue *string                                        `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	Parameters     []*ModifyApplicationParameterRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s ModifyApplicationParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationParameterRequest) GoString() string {
	return s.String()
}

func (s *ModifyApplicationParameterRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationParameterRequest) GetParameterName() *string {
	return s.ParameterName
}

func (s *ModifyApplicationParameterRequest) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *ModifyApplicationParameterRequest) GetParameters() []*ModifyApplicationParameterRequestParameters {
	return s.Parameters
}

func (s *ModifyApplicationParameterRequest) SetApplicationId(v string) *ModifyApplicationParameterRequest {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationParameterRequest) SetParameterName(v string) *ModifyApplicationParameterRequest {
	s.ParameterName = &v
	return s
}

func (s *ModifyApplicationParameterRequest) SetParameterValue(v string) *ModifyApplicationParameterRequest {
	s.ParameterValue = &v
	return s
}

func (s *ModifyApplicationParameterRequest) SetParameters(v []*ModifyApplicationParameterRequestParameters) *ModifyApplicationParameterRequest {
	s.Parameters = v
	return s
}

func (s *ModifyApplicationParameterRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyApplicationParameterRequestParameters struct {
	// example:
	//
	// name
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// example:
	//
	// value
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s ModifyApplicationParameterRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationParameterRequestParameters) GoString() string {
	return s.String()
}

func (s *ModifyApplicationParameterRequestParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *ModifyApplicationParameterRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *ModifyApplicationParameterRequestParameters) SetParameterName(v string) *ModifyApplicationParameterRequestParameters {
	s.ParameterName = &v
	return s
}

func (s *ModifyApplicationParameterRequestParameters) SetParameterValue(v string) *ModifyApplicationParameterRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *ModifyApplicationParameterRequestParameters) Validate() error {
	return dara.Validate(s)
}
