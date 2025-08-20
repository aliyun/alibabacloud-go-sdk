// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackOperationRisksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListStackOperationRisksRequest
	GetClientToken() *string
	SetOperationType(v string) *ListStackOperationRisksRequest
	GetOperationType() *string
	SetRamRoleName(v string) *ListStackOperationRisksRequest
	GetRamRoleName() *string
	SetRegionId(v string) *ListStackOperationRisksRequest
	GetRegionId() *string
	SetRetainAllResources(v bool) *ListStackOperationRisksRequest
	GetRetainAllResources() *bool
	SetRetainResources(v []*string) *ListStackOperationRisksRequest
	GetRetainResources() []*string
	SetStackId(v string) *ListStackOperationRisksRequest
	GetStackId() *string
	SetTemplateBody(v string) *ListStackOperationRisksRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *ListStackOperationRisksRequest
	GetTemplateId() *string
	SetTemplateURL(v string) *ListStackOperationRisksRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *ListStackOperationRisksRequest
	GetTemplateVersion() *string
}

type ListStackOperationRisksRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The type of the operation of which you want to detect risks. Valid values:
	//
	// 	- DeleteStack: detects high risks that may arise in resources when you delete a stack.
	//
	// 	- CreateStack: detects the missing permissions when you fail to create a stack.
	//
	// example:
	//
	// DeleteStack
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The name of the RAM role.
	//
	// 	- If you specify a RAM role, ROS creates stacks based on the permissions that are granted to the RAM role and uses the credentials of the RAM role to call the API operations of Alibaba Cloud services.
	//
	// 	- If you do not specify a RAM role, ROS creates stacks based on the permissions of your Alibaba Cloud account.
	//
	// The name of the RAM role can be up to 64 bytes in length.
	//
	// example:
	//
	// test-role
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to retain all resources in the stack. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// > This parameter takes effect only if you set OperationType to DeleteStack.
	//
	// example:
	//
	// false
	RetainAllResources *bool `json:"RetainAllResources,omitempty" xml:"RetainAllResources,omitempty"`
	// The list of resources to retain.
	//
	// > This parameter takes effect only if you set OperationType to DeleteStack.
	//
	// example:
	//
	// WebServer
	RetainResources []*string `json:"RetainResources,omitempty" xml:"RetainResources,omitempty" type:"Repeated"`
	// The ID of the stack.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId      *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The ID of the template. This parameter applies to shared and private templates.
	//
	// > You must specify one of TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length. If you do not specify RegionId in the URL, the region ID of the stack is used.
	//
	// > You must specify one of TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	//
	// example:
	//
	// oss://ros-template/demo
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template.
	//
	// > This parameter takes effect only if you specify TemplateId.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s ListStackOperationRisksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStackOperationRisksRequest) GoString() string {
	return s.String()
}

func (s *ListStackOperationRisksRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListStackOperationRisksRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *ListStackOperationRisksRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *ListStackOperationRisksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStackOperationRisksRequest) GetRetainAllResources() *bool {
	return s.RetainAllResources
}

func (s *ListStackOperationRisksRequest) GetRetainResources() []*string {
	return s.RetainResources
}

func (s *ListStackOperationRisksRequest) GetStackId() *string {
	return s.StackId
}

func (s *ListStackOperationRisksRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *ListStackOperationRisksRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListStackOperationRisksRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *ListStackOperationRisksRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *ListStackOperationRisksRequest) SetClientToken(v string) *ListStackOperationRisksRequest {
	s.ClientToken = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetOperationType(v string) *ListStackOperationRisksRequest {
	s.OperationType = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetRamRoleName(v string) *ListStackOperationRisksRequest {
	s.RamRoleName = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetRegionId(v string) *ListStackOperationRisksRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetRetainAllResources(v bool) *ListStackOperationRisksRequest {
	s.RetainAllResources = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetRetainResources(v []*string) *ListStackOperationRisksRequest {
	s.RetainResources = v
	return s
}

func (s *ListStackOperationRisksRequest) SetStackId(v string) *ListStackOperationRisksRequest {
	s.StackId = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetTemplateBody(v string) *ListStackOperationRisksRequest {
	s.TemplateBody = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetTemplateId(v string) *ListStackOperationRisksRequest {
	s.TemplateId = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetTemplateURL(v string) *ListStackOperationRisksRequest {
	s.TemplateURL = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetTemplateVersion(v string) *ListStackOperationRisksRequest {
	s.TemplateVersion = &v
	return s
}

func (s *ListStackOperationRisksRequest) Validate() error {
	return dara.Validate(s)
}
