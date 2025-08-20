// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateParameterConstraintsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetTemplateParameterConstraintsShrinkRequest
	GetClientToken() *string
	SetParameters(v []*GetTemplateParameterConstraintsShrinkRequestParameters) *GetTemplateParameterConstraintsShrinkRequest
	GetParameters() []*GetTemplateParameterConstraintsShrinkRequestParameters
	SetParametersKeyFilterShrink(v string) *GetTemplateParameterConstraintsShrinkRequest
	GetParametersKeyFilterShrink() *string
	SetParametersOrderShrink(v string) *GetTemplateParameterConstraintsShrinkRequest
	GetParametersOrderShrink() *string
	SetRegionId(v string) *GetTemplateParameterConstraintsShrinkRequest
	GetRegionId() *string
	SetStackId(v string) *GetTemplateParameterConstraintsShrinkRequest
	GetStackId() *string
	SetTemplateBody(v string) *GetTemplateParameterConstraintsShrinkRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *GetTemplateParameterConstraintsShrinkRequest
	GetTemplateId() *string
	SetTemplateURL(v string) *GetTemplateParameterConstraintsShrinkRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *GetTemplateParameterConstraintsShrinkRequest
	GetTemplateVersion() *string
}

type GetTemplateParameterConstraintsShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of parameter N in the template.
	Parameters []*GetTemplateParameterConstraintsShrinkRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The parameters whose values you want to query.
	ParametersKeyFilterShrink *string `json:"ParametersKeyFilter,omitempty" xml:"ParametersKeyFilter,omitempty"`
	// The order in which associated parameters are arranged.
	//
	// >  By default, the order of the associated parameters specified in the `Metadata` section of the template is used.
	ParametersOrderShrink *string `json:"ParametersOrder,omitempty" xml:"ParametersOrder,omitempty"`
	// The region ID of the template.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack.
	//
	// example:
	//
	// c754d2a4-28f1-46df-b557-9586173a****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The structure that contains the template body.
	//
	// The template body must be 1 to 524,288 bytes in length. If the length of the template body exceeds the upper limit, we recommend that you add parameters to the HTTP POST request body to prevent request failures caused by excessively long URLs.
	//
	// >  You must specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	//
	// example:
	//
	// {"Parameters":{"ZoneInfo":{"Type": "String"},"InstanceType": {"Type": "String"}},"ROSTemplateFormatVersion": "2015-09-01","Resources":{"ECS":{"Properties":{"ZoneId":{"Ref": "ZoneInfo"},"InstanceType": {"Ref": "InstanceType"}},"Type": "ALIYUN::ECS::Instance"}}}
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The ID of the template. This parameter applies to shared and private templates.
	//
	// >  You must specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/stack-policy/demo or oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length. If you do not specify the region ID of the OSS bucket, the value of the RegionId parameter is used.
	//
	// >  You must specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	//
	// example:
	//
	// oss://ros-template/demo
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. If you do not specify this parameter, the latest version is used.
	//
	// >  This parameter takes effect only if the TemplateId parameter is specified.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetTemplateParameterConstraintsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateParameterConstraintsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetTemplateParameterConstraintsShrinkRequest) GetParameters() []*GetTemplateParameterConstraintsShrinkRequestParameters {
	return s.Parameters
}

func (s *GetTemplateParameterConstraintsShrinkRequest) GetParametersKeyFilterShrink() *string {
	return s.ParametersKeyFilterShrink
}

func (s *GetTemplateParameterConstraintsShrinkRequest) GetParametersOrderShrink() *string {
	return s.ParametersOrderShrink
}

func (s *GetTemplateParameterConstraintsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTemplateParameterConstraintsShrinkRequest) GetStackId() *string {
	return s.StackId
}

func (s *GetTemplateParameterConstraintsShrinkRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *GetTemplateParameterConstraintsShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateParameterConstraintsShrinkRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *GetTemplateParameterConstraintsShrinkRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetClientToken(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetParameters(v []*GetTemplateParameterConstraintsShrinkRequestParameters) *GetTemplateParameterConstraintsShrinkRequest {
	s.Parameters = v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetParametersKeyFilterShrink(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.ParametersKeyFilterShrink = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetParametersOrderShrink(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.ParametersOrderShrink = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetRegionId(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetStackId(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.StackId = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetTemplateBody(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.TemplateBody = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetTemplateId(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetTemplateURL(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.TemplateURL = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetTemplateVersion(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type GetTemplateParameterConstraintsShrinkRequestParameters struct {
	// The name of parameter N in the template.
	//
	// >  The Parameters parameter is optional. If you specify the Parameters parameter, you must specify the Parameters.N.ParameterKey parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZoneInfo
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N in the template.
	//
	// >  The Parameters parameter is optional. If you specify the Parameters parameter, you must specify the Parameters.N.ParameterValue parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-h
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetTemplateParameterConstraintsShrinkRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateParameterConstraintsShrinkRequestParameters) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsShrinkRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetTemplateParameterConstraintsShrinkRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GetTemplateParameterConstraintsShrinkRequestParameters) SetParameterKey(v string) *GetTemplateParameterConstraintsShrinkRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequestParameters) SetParameterValue(v string) *GetTemplateParameterConstraintsShrinkRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequestParameters) Validate() error {
	return dara.Validate(s)
}
