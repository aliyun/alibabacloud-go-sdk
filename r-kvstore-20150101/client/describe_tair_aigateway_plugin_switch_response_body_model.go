// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairAIGatewayPluginSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeTairAIGatewayPluginSwitchResponseBody
	GetInstanceId() *string
	SetParameters(v *DescribeTairAIGatewayPluginSwitchResponseBodyParameters) *DescribeTairAIGatewayPluginSwitchResponseBody
	GetParameters() *DescribeTairAIGatewayPluginSwitchResponseBodyParameters
	SetRequestId(v string) *DescribeTairAIGatewayPluginSwitchResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeTairAIGatewayPluginSwitchResponseBody
	GetTotalCount() *int32
}

type DescribeTairAIGatewayPluginSwitchResponseBody struct {
	// The instance ID, which is globally unique.
	//
	// example:
	//
	// tg-bp1zxszhcgatnx****
	InstanceId *string                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Parameters *DescribeTairAIGatewayPluginSwitchResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BBC1E3D6-7C88-4DF5-9A3D-0DB1E6D9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries for the plug-in switch.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTairAIGatewayPluginSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayPluginSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBody) GetParameters() *DescribeTairAIGatewayPluginSwitchResponseBodyParameters {
	return s.Parameters
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBody) SetInstanceId(v string) *DescribeTairAIGatewayPluginSwitchResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBody) SetParameters(v *DescribeTairAIGatewayPluginSwitchResponseBodyParameters) *DescribeTairAIGatewayPluginSwitchResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBody) SetRequestId(v string) *DescribeTairAIGatewayPluginSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBody) SetTotalCount(v int32) *DescribeTairAIGatewayPluginSwitchResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBody) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTairAIGatewayPluginSwitchResponseBodyParameters struct {
	Parameters []*DescribeTairAIGatewayPluginSwitchResponseBodyParametersParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s DescribeTairAIGatewayPluginSwitchResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayPluginSwitchResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBodyParameters) GetParameters() []*DescribeTairAIGatewayPluginSwitchResponseBodyParametersParameters {
	return s.Parameters
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBodyParameters) SetParameters(v []*DescribeTairAIGatewayPluginSwitchResponseBodyParametersParameters) *DescribeTairAIGatewayPluginSwitchResponseBodyParameters {
	s.Parameters = v
	return s
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBodyParameters) Validate() error {
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTairAIGatewayPluginSwitchResponseBodyParametersParameters struct {
	ParameterName  *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeTairAIGatewayPluginSwitchResponseBodyParametersParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayPluginSwitchResponseBodyParametersParameters) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBodyParametersParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBodyParametersParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBodyParametersParameters) SetParameterName(v string) *DescribeTairAIGatewayPluginSwitchResponseBodyParametersParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBodyParametersParameters) SetParameterValue(v string) *DescribeTairAIGatewayPluginSwitchResponseBodyParametersParameters {
	s.ParameterValue = &v
	return s
}

func (s *DescribeTairAIGatewayPluginSwitchResponseBodyParametersParameters) Validate() error {
	return dara.Validate(s)
}
