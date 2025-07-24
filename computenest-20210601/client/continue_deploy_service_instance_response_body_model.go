// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinueDeployServiceInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDryRunResult(v *ContinueDeployServiceInstanceResponseBodyDryRunResult) *ContinueDeployServiceInstanceResponseBody
	GetDryRunResult() *ContinueDeployServiceInstanceResponseBodyDryRunResult
	SetRequestId(v string) *ContinueDeployServiceInstanceResponseBody
	GetRequestId() *string
	SetServiceInstanceId(v string) *ContinueDeployServiceInstanceResponseBody
	GetServiceInstanceId() *string
}

type ContinueDeployServiceInstanceResponseBody struct {
	// The dry run result.
	DryRunResult *ContinueDeployServiceInstanceResponseBodyDryRunResult `json:"DryRunResult,omitempty" xml:"DryRunResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ContinueDeployServiceInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ContinueDeployServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceResponseBody) GetDryRunResult() *ContinueDeployServiceInstanceResponseBodyDryRunResult {
	return s.DryRunResult
}

func (s *ContinueDeployServiceInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ContinueDeployServiceInstanceResponseBody) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *ContinueDeployServiceInstanceResponseBody) SetDryRunResult(v *ContinueDeployServiceInstanceResponseBodyDryRunResult) *ContinueDeployServiceInstanceResponseBody {
	s.DryRunResult = v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBody) SetRequestId(v string) *ContinueDeployServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBody) SetServiceInstanceId(v string) *ContinueDeployServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type ContinueDeployServiceInstanceResponseBodyDryRunResult struct {
	// The parameters that can be modified. The operation that is performed to modify the parameters does not cause a validation error.
	//
	// > This parameter is returned only if DryRun is set to true.
	ParametersAllowedToBeModified []*string `json:"ParametersAllowedToBeModified,omitempty" xml:"ParametersAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters that can be modified under specific conditions. The new values of the parameters determine whether the operation that is performed to modify the parameters causes a validation error.
	//
	// > This parameter is returned only if DryRun is set to true.
	ParametersConditionallyAllowedToBeModified []*string `json:"ParametersConditionallyAllowedToBeModified,omitempty" xml:"ParametersConditionallyAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters that cannot be modified. The operation that is performed to modify the parameters causes a validation error.
	//
	// > This parameter is returned only if DryRun is set to true.
	ParametersNotAllowedToBeModified []*string `json:"ParametersNotAllowedToBeModified,omitempty" xml:"ParametersNotAllowedToBeModified,omitempty" type:"Repeated"`
}

func (s ContinueDeployServiceInstanceResponseBodyDryRunResult) String() string {
	return dara.Prettify(s)
}

func (s ContinueDeployServiceInstanceResponseBodyDryRunResult) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) GetParametersAllowedToBeModified() []*string {
	return s.ParametersAllowedToBeModified
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) GetParametersConditionallyAllowedToBeModified() []*string {
	return s.ParametersConditionallyAllowedToBeModified
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) GetParametersNotAllowedToBeModified() []*string {
	return s.ParametersNotAllowedToBeModified
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) SetParametersAllowedToBeModified(v []*string) *ContinueDeployServiceInstanceResponseBodyDryRunResult {
	s.ParametersAllowedToBeModified = v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) SetParametersConditionallyAllowedToBeModified(v []*string) *ContinueDeployServiceInstanceResponseBodyDryRunResult {
	s.ParametersConditionallyAllowedToBeModified = v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) SetParametersNotAllowedToBeModified(v []*string) *ContinueDeployServiceInstanceResponseBodyDryRunResult {
	s.ParametersNotAllowedToBeModified = v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) Validate() error {
	return dara.Validate(s)
}
