// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameters(v []*DescribeParametersResponseBodyParameters) *DescribeParametersResponseBody
	GetParameters() []*DescribeParametersResponseBodyParameters
	SetRequestId(v string) *DescribeParametersResponseBody
	GetRequestId() *string
}

type DescribeParametersResponseBody struct {
	// The queried configuration parameters.
	Parameters []*DescribeParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 62506167-D284-562A-B7C2-0A**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBody) GetParameters() []*DescribeParametersResponseBodyParameters {
	return s.Parameters
}

func (s *DescribeParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParametersResponseBody) SetParameters(v []*DescribeParametersResponseBodyParameters) *DescribeParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeParametersResponseBody) SetRequestId(v string) *DescribeParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParametersResponseBody) Validate() error {
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

type DescribeParametersResponseBodyParameters struct {
	// The current value of the configuration parameter.
	//
	// example:
	//
	// 10800000
	CurrentValue *string `json:"CurrentValue,omitempty" xml:"CurrentValue,omitempty"`
	// Indicates whether a restart is required for parameter modifications to take effect. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ForceRestartInstance *string `json:"ForceRestartInstance,omitempty" xml:"ForceRestartInstance,omitempty"`
	// Indicates whether the configuration parameter can be modified. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsChangeableConfig *string `json:"IsChangeableConfig,omitempty" xml:"IsChangeableConfig,omitempty"`
	// The valid values of the configuration parameter.
	//
	// example:
	//
	// [0-2147483647]
	OptionalRange *string `json:"OptionalRange,omitempty" xml:"OptionalRange,omitempty"`
	// The description of the configuration parameter.
	//
	// example:
	//
	// Sets the maximum allowed duration of any statement, A value of 0 turns off the timeout.
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the configuration parameter.
	//
	// example:
	//
	// statement_timeout
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The default value of the configuration parameter.
	//
	// example:
	//
	// 10800000
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeParametersResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyParameters) GetCurrentValue() *string {
	return s.CurrentValue
}

func (s *DescribeParametersResponseBodyParameters) GetForceRestartInstance() *string {
	return s.ForceRestartInstance
}

func (s *DescribeParametersResponseBodyParameters) GetIsChangeableConfig() *string {
	return s.IsChangeableConfig
}

func (s *DescribeParametersResponseBodyParameters) GetOptionalRange() *string {
	return s.OptionalRange
}

func (s *DescribeParametersResponseBodyParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *DescribeParametersResponseBodyParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeParametersResponseBodyParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeParametersResponseBodyParameters) SetCurrentValue(v string) *DescribeParametersResponseBodyParameters {
	s.CurrentValue = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetForceRestartInstance(v string) *DescribeParametersResponseBodyParameters {
	s.ForceRestartInstance = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetIsChangeableConfig(v string) *DescribeParametersResponseBodyParameters {
	s.IsChangeableConfig = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetOptionalRange(v string) *DescribeParametersResponseBodyParameters {
	s.OptionalRange = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetParameterDescription(v string) *DescribeParametersResponseBodyParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetParameterName(v string) *DescribeParametersResponseBodyParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetParameterValue(v string) *DescribeParametersResponseBodyParameters {
	s.ParameterValue = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) Validate() error {
	return dara.Validate(s)
}
