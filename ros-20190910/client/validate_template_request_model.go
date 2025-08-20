// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ValidateTemplateRequest
	GetClientToken() *string
	SetRegionId(v string) *ValidateTemplateRequest
	GetRegionId() *string
	SetTemplateBody(v string) *ValidateTemplateRequest
	GetTemplateBody() *string
	SetTemplateURL(v string) *ValidateTemplateRequest
	GetTemplateURL() *string
	SetUpdateInfoOptions(v []*string) *ValidateTemplateRequest
	GetUpdateInfoOptions() []*string
	SetValidationOption(v string) *ValidateTemplateRequest
	GetValidationOption() *string
}

type ValidateTemplateRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the template. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP web server or in an Object Storage Service (OSS) bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length.
	//
	// > If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// You can specify one of TemplateBody and TemplateURL, but not both of them. The URL can be up to 1,024 bytes in length.\\
	//
	// example:
	//
	// oss://ros/template/demo
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The options that are used to control the generation of information about the stack update. You can specify up to two options.
	UpdateInfoOptions []*string `json:"UpdateInfoOptions,omitempty" xml:"UpdateInfoOptions,omitempty" type:"Repeated"`
	// Specifies whether to enable additional validation for the template. Valid values:
	//
	// 	- None (default): does not enable additional validation.
	//
	// 	- EnableTerraformValidation: runs the `terraform validate` command in the Terraform CLI to enable additional validation for a Terraform template.
	//
	// 	- EnableFastTerraformValidation: runs a command that is similar to the `terraform validate` command in the Terraform CLI to enable additional validation for a Terraform template.
	//
	// > Compared with the EnableTerraformValidation method, the EnableFastTerraformValidation method validates a template at a faster speed but a lower integrity level.
	//
	// example:
	//
	// None
	ValidationOption *string `json:"ValidationOption,omitempty" xml:"ValidationOption,omitempty"`
}

func (s ValidateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateTemplateRequest) GoString() string {
	return s.String()
}

func (s *ValidateTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ValidateTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ValidateTemplateRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *ValidateTemplateRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *ValidateTemplateRequest) GetUpdateInfoOptions() []*string {
	return s.UpdateInfoOptions
}

func (s *ValidateTemplateRequest) GetValidationOption() *string {
	return s.ValidationOption
}

func (s *ValidateTemplateRequest) SetClientToken(v string) *ValidateTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *ValidateTemplateRequest) SetRegionId(v string) *ValidateTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *ValidateTemplateRequest) SetTemplateBody(v string) *ValidateTemplateRequest {
	s.TemplateBody = &v
	return s
}

func (s *ValidateTemplateRequest) SetTemplateURL(v string) *ValidateTemplateRequest {
	s.TemplateURL = &v
	return s
}

func (s *ValidateTemplateRequest) SetUpdateInfoOptions(v []*string) *ValidateTemplateRequest {
	s.UpdateInfoOptions = v
	return s
}

func (s *ValidateTemplateRequest) SetValidationOption(v string) *ValidateTemplateRequest {
	s.ValidationOption = &v
	return s
}

func (s *ValidateTemplateRequest) Validate() error {
	return dara.Validate(s)
}
