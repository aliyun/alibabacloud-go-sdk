// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairAIGatewayParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeTairAIGatewayParametersResponseBody
	GetInstanceId() *string
	SetParameters(v *DescribeTairAIGatewayParametersResponseBodyParameters) *DescribeTairAIGatewayParametersResponseBody
	GetParameters() *DescribeTairAIGatewayParametersResponseBodyParameters
	SetRequestId(v string) *DescribeTairAIGatewayParametersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeTairAIGatewayParametersResponseBody
	GetTotalCount() *int32
}

type DescribeTairAIGatewayParametersResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// tg-bp1zxszhcgatnx****
	InstanceId *string                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Parameters *DescribeTairAIGatewayParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 90B82DB7-FB28-4CC2-ADBF-1F8659F3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of returned parameters.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTairAIGatewayParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayParametersResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTairAIGatewayParametersResponseBody) GetParameters() *DescribeTairAIGatewayParametersResponseBodyParameters {
	return s.Parameters
}

func (s *DescribeTairAIGatewayParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTairAIGatewayParametersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTairAIGatewayParametersResponseBody) SetInstanceId(v string) *DescribeTairAIGatewayParametersResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeTairAIGatewayParametersResponseBody) SetParameters(v *DescribeTairAIGatewayParametersResponseBodyParameters) *DescribeTairAIGatewayParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeTairAIGatewayParametersResponseBody) SetRequestId(v string) *DescribeTairAIGatewayParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTairAIGatewayParametersResponseBody) SetTotalCount(v int32) *DescribeTairAIGatewayParametersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTairAIGatewayParametersResponseBody) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTairAIGatewayParametersResponseBodyParameters struct {
	Parameters []*DescribeTairAIGatewayParametersResponseBodyParametersParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s DescribeTairAIGatewayParametersResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayParametersResponseBodyParameters) GetParameters() []*DescribeTairAIGatewayParametersResponseBodyParametersParameters {
	return s.Parameters
}

func (s *DescribeTairAIGatewayParametersResponseBodyParameters) SetParameters(v []*DescribeTairAIGatewayParametersResponseBodyParametersParameters) *DescribeTairAIGatewayParametersResponseBodyParameters {
	s.Parameters = v
	return s
}

func (s *DescribeTairAIGatewayParametersResponseBodyParameters) Validate() error {
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

type DescribeTairAIGatewayParametersResponseBodyParametersParameters struct {
	ParameterName  *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeTairAIGatewayParametersResponseBodyParametersParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayParametersResponseBodyParametersParameters) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayParametersResponseBodyParametersParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeTairAIGatewayParametersResponseBodyParametersParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeTairAIGatewayParametersResponseBodyParametersParameters) SetParameterName(v string) *DescribeTairAIGatewayParametersResponseBodyParametersParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeTairAIGatewayParametersResponseBodyParametersParameters) SetParameterValue(v string) *DescribeTairAIGatewayParametersResponseBodyParametersParameters {
	s.ParameterValue = &v
	return s
}

func (s *DescribeTairAIGatewayParametersResponseBodyParametersParameters) Validate() error {
	return dara.Validate(s)
}
