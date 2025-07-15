// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateParameterConstraintsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParameters(v string) *GetTemplateParameterConstraintsRequest
	GetParameters() *string
	SetRegionId(v string) *GetTemplateParameterConstraintsRequest
	GetRegionId() *string
	SetTemplateContent(v string) *GetTemplateParameterConstraintsRequest
	GetTemplateContent() *string
	SetTemplateName(v string) *GetTemplateParameterConstraintsRequest
	GetTemplateName() *string
	SetTemplateURL(v string) *GetTemplateParameterConstraintsRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *GetTemplateParameterConstraintsRequest
	GetTemplateVersion() *string
}

type GetTemplateParameterConstraintsRequest struct {
	// The information about the parameters.
	//
	// example:
	//
	// {\\"endDate\\": \\"2022-04-13T03:31:20Z\\", \\"Status\\": \\"Stopped\\"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The content of the template in the JSON or YAML format. This parameter is the same as the Content parameter that you can specify when you call the CreateTemplate operation. You can use this parameter to specify the tasks that you want to run. This way, you do not need to create a template before you start an execution. If you select an existing template, you do not need to specify this parameter.
	//
	// example:
	//
	// {
	//
	// 	"Description": "Example template, describe instances in some status",
	//
	// 	"FormatVersion": "OOS-2019-06-01",
	//
	// 	"Parameters": {},
	//
	// 	"Tasks": [{
	//
	// 		"Name": "describeInstances",
	//
	// 		"Action": "ACS::ExecuteAPI",
	//
	// 		"Description": "desc-en",
	//
	// 		"Properties": {
	//
	// 			"Service": "ECS",
	//
	// 			"API": "DescribeInstances",
	//
	// 			"Parameters": {
	//
	// 				"Status": "Running"
	//
	// 			}
	//
	// 		}
	//
	// 	}]
	//
	// }
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The URL that is used to store the content of the Operation Orchestration Service (OOS) template in the Alibaba Cloud Object Storage Service (OSS). Only the public-read URL is supported. You can use this parameter to specify the tasks that you want to run. This way, you do not need to create a template before you start an execution. If you select an existing template, you do not need to specify this parameter.
	//
	// example:
	//
	// http://oos-template.cn-hangzhou.oss.aliyun-inc.com/oos-test-template.json
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. The default value is the latest version of the template.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetTemplateParameterConstraintsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateParameterConstraintsRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsRequest) GetParameters() *string {
	return s.Parameters
}

func (s *GetTemplateParameterConstraintsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTemplateParameterConstraintsRequest) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *GetTemplateParameterConstraintsRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetTemplateParameterConstraintsRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *GetTemplateParameterConstraintsRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetTemplateParameterConstraintsRequest) SetParameters(v string) *GetTemplateParameterConstraintsRequest {
	s.Parameters = &v
	return s
}

func (s *GetTemplateParameterConstraintsRequest) SetRegionId(v string) *GetTemplateParameterConstraintsRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateParameterConstraintsRequest) SetTemplateContent(v string) *GetTemplateParameterConstraintsRequest {
	s.TemplateContent = &v
	return s
}

func (s *GetTemplateParameterConstraintsRequest) SetTemplateName(v string) *GetTemplateParameterConstraintsRequest {
	s.TemplateName = &v
	return s
}

func (s *GetTemplateParameterConstraintsRequest) SetTemplateURL(v string) *GetTemplateParameterConstraintsRequest {
	s.TemplateURL = &v
	return s
}

func (s *GetTemplateParameterConstraintsRequest) SetTemplateVersion(v string) *GetTemplateParameterConstraintsRequest {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateParameterConstraintsRequest) Validate() error {
	return dara.Validate(s)
}
