// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateEstimateCostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetTemplateEstimateCostRequest
	GetClientToken() *string
	SetParameters(v []*GetTemplateEstimateCostRequestParameters) *GetTemplateEstimateCostRequest
	GetParameters() []*GetTemplateEstimateCostRequestParameters
	SetRegionId(v string) *GetTemplateEstimateCostRequest
	GetRegionId() *string
	SetStackId(v string) *GetTemplateEstimateCostRequest
	GetStackId() *string
	SetTemplateBody(v string) *GetTemplateEstimateCostRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *GetTemplateEstimateCostRequest
	GetTemplateId() *string
	SetTemplateScratchId(v string) *GetTemplateEstimateCostRequest
	GetTemplateScratchId() *string
	SetTemplateScratchRegionId(v string) *GetTemplateEstimateCostRequest
	GetTemplateScratchRegionId() *string
	SetTemplateURL(v string) *GetTemplateEstimateCostRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *GetTemplateEstimateCostRequest
	GetTemplateVersion() *string
}

type GetTemplateEstimateCostRequest struct {
	// The name of parameter N. If you do not specify the name and value of a parameter, ROS uses the default name and value that are specified in the template.
	//
	// Maximum value of N: 200.
	//
	// Examples:
	//
	// 	- Parameters.1.ParameterKey: `Name`
	//
	// 	- Parameters.2.ParameterKey: `Netmode`
	//
	// >  The Parameters parameter is optional. If you want to specify Parameters, you must specify both Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the scenario. The default value is the same as the value of the RegionId parameter.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	Parameters []*GetTemplateEstimateCostRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the template. This parameter applies to shared and private templates.
	//
	// >  You must specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The stack ID.
	//
	// This parameter is used to query the estimated price in a configuration change scenario.
	//
	// Assume that the specified stack contains only one Elastic Compute Service (ECS) instance and the instance type is ecs.s6-c1m2.large. You downgrade the instance type to ecs.s6-c1m1.small and specify a new ApsaraDB RDS instance in the template that is used for the price inquiry. The queried result is the sum of the downgrade refund of the ECS instance and the price of the new ApsaraDB RDS instance.
	//
	// example:
	//
	// c754d2a4-28f1-46df-b557-9586173a****
	StackId      *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The value of parameter N.
	//
	// Maximum value of N: 200.
	//
	// Examples:
	//
	// 	- Parameters.1.ParameterValue: `DemoEip`
	//
	// 	- Parameters.2.ParameterValue: `public`
	//
	// >  The Parameters parameter is optional. If you want to specify Parameters, you must specify both Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the scenario.
	//
	// example:
	//
	// ts-aa9c62feab844a6b****
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The region ID of the scenario. The default value is the same as the value of the RegionId parameter.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	TemplateScratchRegionId *string `json:"TemplateScratchRegionId,omitempty" xml:"TemplateScratchRegionId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests.
	//
	// The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// oss://ros-template/demo
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The ID of the scenario.
	//
	// For more information about how to query the IDs of scenarios, see [ListTemplateScratches](https://help.aliyun.com/document_detail/363050.html).
	//
	// >  You must specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetTemplateEstimateCostRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateEstimateCostRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateEstimateCostRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetTemplateEstimateCostRequest) GetParameters() []*GetTemplateEstimateCostRequestParameters {
	return s.Parameters
}

func (s *GetTemplateEstimateCostRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTemplateEstimateCostRequest) GetStackId() *string {
	return s.StackId
}

func (s *GetTemplateEstimateCostRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *GetTemplateEstimateCostRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateEstimateCostRequest) GetTemplateScratchId() *string {
	return s.TemplateScratchId
}

func (s *GetTemplateEstimateCostRequest) GetTemplateScratchRegionId() *string {
	return s.TemplateScratchRegionId
}

func (s *GetTemplateEstimateCostRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *GetTemplateEstimateCostRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetTemplateEstimateCostRequest) SetClientToken(v string) *GetTemplateEstimateCostRequest {
	s.ClientToken = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetParameters(v []*GetTemplateEstimateCostRequestParameters) *GetTemplateEstimateCostRequest {
	s.Parameters = v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetRegionId(v string) *GetTemplateEstimateCostRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetStackId(v string) *GetTemplateEstimateCostRequest {
	s.StackId = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetTemplateBody(v string) *GetTemplateEstimateCostRequest {
	s.TemplateBody = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetTemplateId(v string) *GetTemplateEstimateCostRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetTemplateScratchId(v string) *GetTemplateEstimateCostRequest {
	s.TemplateScratchId = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetTemplateScratchRegionId(v string) *GetTemplateEstimateCostRequest {
	s.TemplateScratchRegionId = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetTemplateURL(v string) *GetTemplateEstimateCostRequest {
	s.TemplateURL = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetTemplateVersion(v string) *GetTemplateEstimateCostRequest {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) Validate() error {
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

type GetTemplateEstimateCostRequestParameters struct {
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// Name
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// Details of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// DemoEip
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetTemplateEstimateCostRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateEstimateCostRequestParameters) GoString() string {
	return s.String()
}

func (s *GetTemplateEstimateCostRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetTemplateEstimateCostRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GetTemplateEstimateCostRequestParameters) SetParameterKey(v string) *GetTemplateEstimateCostRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetTemplateEstimateCostRequestParameters) SetParameterValue(v string) *GetTemplateEstimateCostRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetTemplateEstimateCostRequestParameters) Validate() error {
	return dara.Validate(s)
}
