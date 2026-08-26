// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomTemplate(v string) *CreateCustomTemplateRequest
	GetCustomTemplate() *string
	SetOwnerId(v int64) *CreateCustomTemplateRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateCustomTemplateRequest
	GetRegionId() *string
	SetTemplate(v string) *CreateCustomTemplateRequest
	GetTemplate() *string
}

type CreateCustomTemplateRequest struct {
	// The template configuration in JSON format. The value must be in the following format: {height:xxx,scale:xxx,gop:xxx,bframes:xxx,cdesc:xxx}. All fields are required. The call fails if any field is missing.
	//
	// > For more information about the parameters, see the **CustomTemplate details*	- table below.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"height":"1060","scale":"[16:9]","gop":"60","bframes":"30","cdesc":"h264"}
	CustomTemplate *string `json:"CustomTemplate,omitempty" xml:"CustomTemplate,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the template to add.
	//
	// > Record the template name after you create it. The name is required for subsequent operations, such as using, querying, and deleting the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestTemplate
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s CreateCustomTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomTemplateRequest) GetCustomTemplate() *string {
	return s.CustomTemplate
}

func (s *CreateCustomTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCustomTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCustomTemplateRequest) GetTemplate() *string {
	return s.Template
}

func (s *CreateCustomTemplateRequest) SetCustomTemplate(v string) *CreateCustomTemplateRequest {
	s.CustomTemplate = &v
	return s
}

func (s *CreateCustomTemplateRequest) SetOwnerId(v int64) *CreateCustomTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCustomTemplateRequest) SetRegionId(v string) *CreateCustomTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCustomTemplateRequest) SetTemplate(v string) *CreateCustomTemplateRequest {
	s.Template = &v
	return s
}

func (s *CreateCustomTemplateRequest) Validate() error {
	return dara.Validate(s)
}
