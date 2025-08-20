// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeSetId(v string) *GetTemplateSummaryRequest
	GetChangeSetId() *string
	SetClientToken(v string) *GetTemplateSummaryRequest
	GetClientToken() *string
	SetParameters(v []*GetTemplateSummaryRequestParameters) *GetTemplateSummaryRequest
	GetParameters() []*GetTemplateSummaryRequestParameters
	SetRegionId(v string) *GetTemplateSummaryRequest
	GetRegionId() *string
	SetStackGroupName(v string) *GetTemplateSummaryRequest
	GetStackGroupName() *string
	SetStackId(v string) *GetTemplateSummaryRequest
	GetStackId() *string
	SetTemplateBody(v string) *GetTemplateSummaryRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *GetTemplateSummaryRequest
	GetTemplateId() *string
	SetTemplateURL(v string) *GetTemplateSummaryRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *GetTemplateSummaryRequest
	GetTemplateVersion() *string
}

type GetTemplateSummaryRequest struct {
	// The ID of the change set.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, StackId, ChangeSetId, and StackGroupName.
	//
	// example:
	//
	// 1f6521a4-05af-4975-afe9-bc4b45ad****
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.\\
	//
	// The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (_).\\
	//
	// For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The parameters that are defined in the template.
	Parameters []*GetTemplateSummaryRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The region ID of the stack or stack group that uses the template. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter takes effect only when one of the following parameters are specified: StackId, ChangeSetId, and StackGroupName.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stack group.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, StackId, ChangeSetId, and StackGroupName.
	//
	// example:
	//
	// my-stack-group
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The stack ID.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, StackId, ChangeSetId, and StackGroupName.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.\\
	//
	// If the length of the template body exceeds the upper limit, we recommend that you add parameters to the HTTP POST request body to prevent request failures caused by excessively long URLs.\\
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, StackId, ChangeSetId, and StackGroupName.
	//
	// example:
	//
	// {"ROSTemplateFormatVersion":"2015-09-01"}
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The template ID. This parameter applies to shared and private templates.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, StackId, ChangeSetId, and StackGroupName.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length.
	//
	// > If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, StackId, ChangeSetId, and StackGroupName.
	//
	// The URL can be up to 1,024 bytes in length.
	//
	// example:
	//
	// oss://ros/template/demo
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. This parameter takes effect when TemplateId is specified.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetTemplateSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateSummaryRequest) GetChangeSetId() *string {
	return s.ChangeSetId
}

func (s *GetTemplateSummaryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetTemplateSummaryRequest) GetParameters() []*GetTemplateSummaryRequestParameters {
	return s.Parameters
}

func (s *GetTemplateSummaryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTemplateSummaryRequest) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *GetTemplateSummaryRequest) GetStackId() *string {
	return s.StackId
}

func (s *GetTemplateSummaryRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *GetTemplateSummaryRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateSummaryRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *GetTemplateSummaryRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetTemplateSummaryRequest) SetChangeSetId(v string) *GetTemplateSummaryRequest {
	s.ChangeSetId = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetClientToken(v string) *GetTemplateSummaryRequest {
	s.ClientToken = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetParameters(v []*GetTemplateSummaryRequestParameters) *GetTemplateSummaryRequest {
	s.Parameters = v
	return s
}

func (s *GetTemplateSummaryRequest) SetRegionId(v string) *GetTemplateSummaryRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetStackGroupName(v string) *GetTemplateSummaryRequest {
	s.StackGroupName = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetStackId(v string) *GetTemplateSummaryRequest {
	s.StackId = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetTemplateBody(v string) *GetTemplateSummaryRequest {
	s.TemplateBody = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetTemplateId(v string) *GetTemplateSummaryRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetTemplateURL(v string) *GetTemplateSummaryRequest {
	s.TemplateURL = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetTemplateVersion(v string) *GetTemplateSummaryRequest {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateSummaryRequest) Validate() error {
	return dara.Validate(s)
}

type GetTemplateSummaryRequestParameters struct {
	// The name of parameter N that is defined in the template. If you do not specify the name and value of a parameter, Resource Orchestration Service (ROS) uses the default name and value that are defined in the template.
	//
	// Maximum value of N: 200.
	//
	// > The Parameters parameter is optional. If you specify Parameters, you must specify both Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	//
	// example:
	//
	// InstanceId
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N that is defined in the template.
	//
	// Maximum value of N: 200.
	//
	// > The Parameters parameter is optional. If you specify Parameters, you must specify both Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	//
	// example:
	//
	// i-rotp2e20whfrs2****
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetTemplateSummaryRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateSummaryRequestParameters) GoString() string {
	return s.String()
}

func (s *GetTemplateSummaryRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetTemplateSummaryRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GetTemplateSummaryRequestParameters) SetParameterKey(v string) *GetTemplateSummaryRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetTemplateSummaryRequestParameters) SetParameterValue(v string) *GetTemplateSummaryRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetTemplateSummaryRequestParameters) Validate() error {
	return dara.Validate(s)
}
