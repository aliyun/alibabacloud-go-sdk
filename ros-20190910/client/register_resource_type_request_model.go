// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterResourceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RegisterResourceTypeRequest
	GetClientToken() *string
	SetDescription(v string) *RegisterResourceTypeRequest
	GetDescription() *string
	SetEntityType(v string) *RegisterResourceTypeRequest
	GetEntityType() *string
	SetResourceType(v string) *RegisterResourceTypeRequest
	GetResourceType() *string
	SetTemplateBody(v string) *RegisterResourceTypeRequest
	GetTemplateBody() *string
	SetTemplateURL(v string) *RegisterResourceTypeRequest
	GetTemplateURL() *string
}

type RegisterResourceTypeRequest struct {
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
	// The description of the resource type. The description can be up to 512 characters in length.
	//
	// example:
	//
	// It is a demo.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The entity type. Set the value to Module.
	//
	// This parameter is required.
	//
	// example:
	//
	// Module
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// MODULE::MyOrganization::MyService::MyUsecase
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. The template body is used as the module content. If the length of the template body exceeds the upper limit, we recommend that you add parameters to the HTTP POST request body to prevent request failures caused by excessively long URLs.
	//
	//
	// > - This parameter takes effect only when EntityType is set to Module.
	//
	// > - You can specify only one of the TemplateBody and TemplateURL parameters.
	//
	// example:
	//
	// {"ROSTemplateFormatVersion":"2015-09-01"}
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length. The template body is used as the module content.
	//
	// > - If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// > -  This parameter takes effect only when EntityType is set to Module.
	//
	// > -  You can specify only one of the TemplateBody and TemplateURL parameters.
	//
	// The URL can be up to 1,024 bytes in length.
	//
	// example:
	//
	// oss://ros-template/demo
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
}

func (s RegisterResourceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *RegisterResourceTypeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RegisterResourceTypeRequest) GetDescription() *string {
	return s.Description
}

func (s *RegisterResourceTypeRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *RegisterResourceTypeRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *RegisterResourceTypeRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *RegisterResourceTypeRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *RegisterResourceTypeRequest) SetClientToken(v string) *RegisterResourceTypeRequest {
	s.ClientToken = &v
	return s
}

func (s *RegisterResourceTypeRequest) SetDescription(v string) *RegisterResourceTypeRequest {
	s.Description = &v
	return s
}

func (s *RegisterResourceTypeRequest) SetEntityType(v string) *RegisterResourceTypeRequest {
	s.EntityType = &v
	return s
}

func (s *RegisterResourceTypeRequest) SetResourceType(v string) *RegisterResourceTypeRequest {
	s.ResourceType = &v
	return s
}

func (s *RegisterResourceTypeRequest) SetTemplateBody(v string) *RegisterResourceTypeRequest {
	s.TemplateBody = &v
	return s
}

func (s *RegisterResourceTypeRequest) SetTemplateURL(v string) *RegisterResourceTypeRequest {
	s.TemplateURL = &v
	return s
}

func (s *RegisterResourceTypeRequest) Validate() error {
	return dara.Validate(s)
}
