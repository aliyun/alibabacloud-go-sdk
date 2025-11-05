// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSolutionInstanceConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeSolutionInstanceConfigurationRequest
	GetClientToken() *string
	SetParameters(v []*DescribeSolutionInstanceConfigurationRequestParameters) *DescribeSolutionInstanceConfigurationRequest
	GetParameters() []*DescribeSolutionInstanceConfigurationRequestParameters
	SetRegionId(v string) *DescribeSolutionInstanceConfigurationRequest
	GetRegionId() *string
	SetSolutionId(v string) *DescribeSolutionInstanceConfigurationRequest
	GetSolutionId() *string
}

type DescribeSolutionInstanceConfigurationRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The parameters.
	Parameters []*DescribeSolutionInstanceConfigurationRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the solution.
	//
	// This parameter is required.
	//
	// example:
	//
	// sln-xxxxx
	SolutionId *string `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s DescribeSolutionInstanceConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSolutionInstanceConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DescribeSolutionInstanceConfigurationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeSolutionInstanceConfigurationRequest) GetParameters() []*DescribeSolutionInstanceConfigurationRequestParameters {
	return s.Parameters
}

func (s *DescribeSolutionInstanceConfigurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSolutionInstanceConfigurationRequest) GetSolutionId() *string {
	return s.SolutionId
}

func (s *DescribeSolutionInstanceConfigurationRequest) SetClientToken(v string) *DescribeSolutionInstanceConfigurationRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSolutionInstanceConfigurationRequest) SetParameters(v []*DescribeSolutionInstanceConfigurationRequestParameters) *DescribeSolutionInstanceConfigurationRequest {
	s.Parameters = v
	return s
}

func (s *DescribeSolutionInstanceConfigurationRequest) SetRegionId(v string) *DescribeSolutionInstanceConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSolutionInstanceConfigurationRequest) SetSolutionId(v string) *DescribeSolutionInstanceConfigurationRequest {
	s.SolutionId = &v
	return s
}

func (s *DescribeSolutionInstanceConfigurationRequest) Validate() error {
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

type DescribeSolutionInstanceConfigurationRequestParameters struct {
	// The key of the parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ***
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter.
	//
	// > The Parameters parameter is optional. If you specify Parameters, you must specify ParameterValue.
	//
	// This parameter is required.
	//
	// example:
	//
	// ***
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeSolutionInstanceConfigurationRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeSolutionInstanceConfigurationRequestParameters) GoString() string {
	return s.String()
}

func (s *DescribeSolutionInstanceConfigurationRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *DescribeSolutionInstanceConfigurationRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeSolutionInstanceConfigurationRequestParameters) SetParameterKey(v string) *DescribeSolutionInstanceConfigurationRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *DescribeSolutionInstanceConfigurationRequestParameters) SetParameterValue(v string) *DescribeSolutionInstanceConfigurationRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *DescribeSolutionInstanceConfigurationRequestParameters) Validate() error {
	return dara.Validate(s)
}
