// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStackTemplateByResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateStackTemplateByResourcesRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateStackTemplateByResourcesRequest
	GetDryRun() *bool
	SetLogicalResourceId(v []*string) *UpdateStackTemplateByResourcesRequest
	GetLogicalResourceId() []*string
	SetRegionId(v string) *UpdateStackTemplateByResourcesRequest
	GetRegionId() *string
	SetStackId(v string) *UpdateStackTemplateByResourcesRequest
	GetStackId() *string
	SetTemplateFormat(v string) *UpdateStackTemplateByResourcesRequest
	GetTemplateFormat() *string
}

type UpdateStackTemplateByResourcesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests.
	//
	// The token can be up to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only preview the corrected template in this request. Default value: false. Valid values:
	//
	// 	- true: returns the content of the corrected template and does not correct the template. After Resource Orchestration Service (ROS) compares the corrected template with the original template, ROS determines whether to execute the correction.
	//
	// 	- false: corrects the template to eliminate drift.
	//
	// >  We recommend that you set the DryRun parameter to true to preview the corrected template. If the template content meets expectations, set the DryRun parameter to false to execute the correction.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The logical ID of resource.
	//
	// example:
	//
	// Vpc
	LogicalResourceId []*string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty" type:"Repeated"`
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The format of the returned template. Default value: JSON. Valid values:
	//
	// 	- JSON
	//
	// 	- YAML
	//
	// example:
	//
	// JSON
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
}

func (s UpdateStackTemplateByResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackTemplateByResourcesRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackTemplateByResourcesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateStackTemplateByResourcesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateStackTemplateByResourcesRequest) GetLogicalResourceId() []*string {
	return s.LogicalResourceId
}

func (s *UpdateStackTemplateByResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateStackTemplateByResourcesRequest) GetStackId() *string {
	return s.StackId
}

func (s *UpdateStackTemplateByResourcesRequest) GetTemplateFormat() *string {
	return s.TemplateFormat
}

func (s *UpdateStackTemplateByResourcesRequest) SetClientToken(v string) *UpdateStackTemplateByResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackTemplateByResourcesRequest) SetDryRun(v bool) *UpdateStackTemplateByResourcesRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateStackTemplateByResourcesRequest) SetLogicalResourceId(v []*string) *UpdateStackTemplateByResourcesRequest {
	s.LogicalResourceId = v
	return s
}

func (s *UpdateStackTemplateByResourcesRequest) SetRegionId(v string) *UpdateStackTemplateByResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStackTemplateByResourcesRequest) SetStackId(v string) *UpdateStackTemplateByResourcesRequest {
	s.StackId = &v
	return s
}

func (s *UpdateStackTemplateByResourcesRequest) SetTemplateFormat(v string) *UpdateStackTemplateByResourcesRequest {
	s.TemplateFormat = &v
	return s
}

func (s *UpdateStackTemplateByResourcesRequest) Validate() error {
	return dara.Validate(s)
}
