// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDryRunResult(v *UpdateStackResponseBodyDryRunResult) *UpdateStackResponseBody
	GetDryRunResult() *UpdateStackResponseBodyDryRunResult
	SetRequestId(v string) *UpdateStackResponseBody
	GetRequestId() *string
	SetStackId(v string) *UpdateStackResponseBody
	GetStackId() *string
}

type UpdateStackResponseBody struct {
	// The dry run result. This parameter is returned only if DryRun is set to true.
	DryRunResult *UpdateStackResponseBodyDryRunResult `json:"DryRunResult,omitempty" xml:"DryRunResult,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the stack.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s UpdateStackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStackResponseBody) GetDryRunResult() *UpdateStackResponseBodyDryRunResult {
	return s.DryRunResult
}

func (s *UpdateStackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStackResponseBody) GetStackId() *string {
	return s.StackId
}

func (s *UpdateStackResponseBody) SetDryRunResult(v *UpdateStackResponseBodyDryRunResult) *UpdateStackResponseBody {
	s.DryRunResult = v
	return s
}

func (s *UpdateStackResponseBody) SetRequestId(v string) *UpdateStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStackResponseBody) SetStackId(v string) *UpdateStackResponseBody {
	s.StackId = &v
	return s
}

func (s *UpdateStackResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateStackResponseBodyDryRunResult struct {
	// The parameters that can be modified. If you change only values of the parameters in a stack template and use the template to update the stack, no validation errors are caused.
	ParametersAllowedToBeModified []*string `json:"ParametersAllowedToBeModified,omitempty" xml:"ParametersAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters whose changes cause service interruptions.
	//
	// > - This parameter is supported only for a small number of resource types.
	//
	// > - This parameter is valid only for updates on ROS stacks.
	ParametersCauseInterruptionIfModified []*string `json:"ParametersCauseInterruptionIfModified,omitempty" xml:"ParametersCauseInterruptionIfModified,omitempty" type:"Repeated"`
	// The parameters whose changes trigger replacement updates for resources.
	//
	// > -  This parameter can be returned only if ReplacementOption is set to Enabled.
	//
	// > -  This parameter is valid only for updates on ROS stacks.
	ParametersCauseReplacementIfModified []*string `json:"ParametersCauseReplacementIfModified,omitempty" xml:"ParametersCauseReplacementIfModified,omitempty" type:"Repeated"`
	// The parameters that can be modified under specific conditions. If you change only values of the parameters in a stack template and use the template to update the stack, the new values of the parameters determine whether validation errors are caused.
	ParametersConditionallyAllowedToBeModified []*string `json:"ParametersConditionallyAllowedToBeModified,omitempty" xml:"ParametersConditionallyAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters whose changes cause service interruptions under specific conditions.
	//
	// > - This parameter is supported only for a small number of resource types.
	//
	// > -  This parameter is valid only for updates on ROS stacks.
	ParametersConditionallyCauseInterruptionIfModified []*string `json:"ParametersConditionallyCauseInterruptionIfModified,omitempty" xml:"ParametersConditionallyCauseInterruptionIfModified,omitempty" type:"Repeated"`
	// The parameters whose changes trigger replacement updates for resources under specific conditions.
	//
	// > - This parameter can be returned only if ReplacementOption is set to Enabled.
	//
	// > - This parameter is valid only for updates on ROS stacks.
	ParametersConditionallyCauseReplacementIfModified []*string `json:"ParametersConditionallyCauseReplacementIfModified,omitempty" xml:"ParametersConditionallyCauseReplacementIfModified,omitempty" type:"Repeated"`
	// The parameters that cannot be modified. If you change only values of the parameters in a stack template and use the template to update the stack, validation errors are caused.
	ParametersNotAllowedToBeModified []*string `json:"ParametersNotAllowedToBeModified,omitempty" xml:"ParametersNotAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters that can be modified under uncertain conditions. If you change only values of the parameters in a stack template and use the template to update the stack, the actual running environment determines whether validation errors are caused.
	ParametersUncertainlyAllowedToBeModified []*string `json:"ParametersUncertainlyAllowedToBeModified,omitempty" xml:"ParametersUncertainlyAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters whose changes cause service interruptions under uncertain conditions.
	//
	// > - This parameter is supported only for a small number of resource types.
	//
	// > - This parameter is valid only for updates on ROS stacks.
	ParametersUncertainlyCauseInterruptionIfModified []*string `json:"ParametersUncertainlyCauseInterruptionIfModified,omitempty" xml:"ParametersUncertainlyCauseInterruptionIfModified,omitempty" type:"Repeated"`
	// The parameters whose changes trigger replacement updates for resources under uncertain conditions.
	//
	// > - This parameter can be returned only if ReplacementOption is set to Enabled.
	//
	// > - This parameter is valid only for updates on ROS stacks.
	ParametersUncertainlyCauseReplacementIfModified []*string `json:"ParametersUncertainlyCauseReplacementIfModified,omitempty" xml:"ParametersUncertainlyCauseReplacementIfModified,omitempty" type:"Repeated"`
}

func (s UpdateStackResponseBodyDryRunResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackResponseBodyDryRunResult) GoString() string {
	return s.String()
}

func (s *UpdateStackResponseBodyDryRunResult) GetParametersAllowedToBeModified() []*string {
	return s.ParametersAllowedToBeModified
}

func (s *UpdateStackResponseBodyDryRunResult) GetParametersCauseInterruptionIfModified() []*string {
	return s.ParametersCauseInterruptionIfModified
}

func (s *UpdateStackResponseBodyDryRunResult) GetParametersCauseReplacementIfModified() []*string {
	return s.ParametersCauseReplacementIfModified
}

func (s *UpdateStackResponseBodyDryRunResult) GetParametersConditionallyAllowedToBeModified() []*string {
	return s.ParametersConditionallyAllowedToBeModified
}

func (s *UpdateStackResponseBodyDryRunResult) GetParametersConditionallyCauseInterruptionIfModified() []*string {
	return s.ParametersConditionallyCauseInterruptionIfModified
}

func (s *UpdateStackResponseBodyDryRunResult) GetParametersConditionallyCauseReplacementIfModified() []*string {
	return s.ParametersConditionallyCauseReplacementIfModified
}

func (s *UpdateStackResponseBodyDryRunResult) GetParametersNotAllowedToBeModified() []*string {
	return s.ParametersNotAllowedToBeModified
}

func (s *UpdateStackResponseBodyDryRunResult) GetParametersUncertainlyAllowedToBeModified() []*string {
	return s.ParametersUncertainlyAllowedToBeModified
}

func (s *UpdateStackResponseBodyDryRunResult) GetParametersUncertainlyCauseInterruptionIfModified() []*string {
	return s.ParametersUncertainlyCauseInterruptionIfModified
}

func (s *UpdateStackResponseBodyDryRunResult) GetParametersUncertainlyCauseReplacementIfModified() []*string {
	return s.ParametersUncertainlyCauseReplacementIfModified
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersAllowedToBeModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersAllowedToBeModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersCauseInterruptionIfModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersCauseInterruptionIfModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersCauseReplacementIfModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersCauseReplacementIfModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersConditionallyAllowedToBeModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersConditionallyAllowedToBeModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersConditionallyCauseInterruptionIfModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersConditionallyCauseInterruptionIfModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersConditionallyCauseReplacementIfModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersConditionallyCauseReplacementIfModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersNotAllowedToBeModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersNotAllowedToBeModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersUncertainlyAllowedToBeModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersUncertainlyAllowedToBeModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersUncertainlyCauseInterruptionIfModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersUncertainlyCauseInterruptionIfModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersUncertainlyCauseReplacementIfModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersUncertainlyCauseReplacementIfModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) Validate() error {
	return dara.Validate(s)
}
