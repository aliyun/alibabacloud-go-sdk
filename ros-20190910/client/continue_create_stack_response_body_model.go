// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinueCreateStackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDryRunResult(v *ContinueCreateStackResponseBodyDryRunResult) *ContinueCreateStackResponseBody
	GetDryRunResult() *ContinueCreateStackResponseBodyDryRunResult
	SetRequestId(v string) *ContinueCreateStackResponseBody
	GetRequestId() *string
	SetStackId(v string) *ContinueCreateStackResponseBody
	GetStackId() *string
}

type ContinueCreateStackResponseBody struct {
	// The validation result.
	DryRunResult *ContinueCreateStackResponseBodyDryRunResult `json:"DryRunResult,omitempty" xml:"DryRunResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The stack ID.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s ContinueCreateStackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ContinueCreateStackResponseBody) GoString() string {
	return s.String()
}

func (s *ContinueCreateStackResponseBody) GetDryRunResult() *ContinueCreateStackResponseBodyDryRunResult {
	return s.DryRunResult
}

func (s *ContinueCreateStackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ContinueCreateStackResponseBody) GetStackId() *string {
	return s.StackId
}

func (s *ContinueCreateStackResponseBody) SetDryRunResult(v *ContinueCreateStackResponseBodyDryRunResult) *ContinueCreateStackResponseBody {
	s.DryRunResult = v
	return s
}

func (s *ContinueCreateStackResponseBody) SetRequestId(v string) *ContinueCreateStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContinueCreateStackResponseBody) SetStackId(v string) *ContinueCreateStackResponseBody {
	s.StackId = &v
	return s
}

func (s *ContinueCreateStackResponseBody) Validate() error {
	return dara.Validate(s)
}

type ContinueCreateStackResponseBodyDryRunResult struct {
	// The parameters that can be modified.
	ParametersAllowedToBeModified []*string `json:"ParametersAllowedToBeModified,omitempty" xml:"ParametersAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters that can be modified under specific conditions.
	ParametersConditionallyAllowedToBeModified []*string `json:"ParametersConditionallyAllowedToBeModified,omitempty" xml:"ParametersConditionallyAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters that cannot be modified.
	ParametersNotAllowedToBeModified []*string `json:"ParametersNotAllowedToBeModified,omitempty" xml:"ParametersNotAllowedToBeModified,omitempty" type:"Repeated"`
}

func (s ContinueCreateStackResponseBodyDryRunResult) String() string {
	return dara.Prettify(s)
}

func (s ContinueCreateStackResponseBodyDryRunResult) GoString() string {
	return s.String()
}

func (s *ContinueCreateStackResponseBodyDryRunResult) GetParametersAllowedToBeModified() []*string {
	return s.ParametersAllowedToBeModified
}

func (s *ContinueCreateStackResponseBodyDryRunResult) GetParametersConditionallyAllowedToBeModified() []*string {
	return s.ParametersConditionallyAllowedToBeModified
}

func (s *ContinueCreateStackResponseBodyDryRunResult) GetParametersNotAllowedToBeModified() []*string {
	return s.ParametersNotAllowedToBeModified
}

func (s *ContinueCreateStackResponseBodyDryRunResult) SetParametersAllowedToBeModified(v []*string) *ContinueCreateStackResponseBodyDryRunResult {
	s.ParametersAllowedToBeModified = v
	return s
}

func (s *ContinueCreateStackResponseBodyDryRunResult) SetParametersConditionallyAllowedToBeModified(v []*string) *ContinueCreateStackResponseBodyDryRunResult {
	s.ParametersConditionallyAllowedToBeModified = v
	return s
}

func (s *ContinueCreateStackResponseBodyDryRunResult) SetParametersNotAllowedToBeModified(v []*string) *ContinueCreateStackResponseBodyDryRunResult {
	s.ParametersNotAllowedToBeModified = v
	return s
}

func (s *ContinueCreateStackResponseBodyDryRunResult) Validate() error {
	return dara.Validate(s)
}
