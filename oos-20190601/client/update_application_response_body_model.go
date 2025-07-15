// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplication(v *UpdateApplicationResponseBodyApplication) *UpdateApplicationResponseBody
	GetApplication() *UpdateApplicationResponseBodyApplication
	SetRequestId(v string) *UpdateApplicationResponseBody
	GetRequestId() *string
}

type UpdateApplicationResponseBody struct {
	// The information about the application.
	Application *UpdateApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F1F00F41-D24C-5377-831B-C97F739CE1AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBody) GetApplication() *UpdateApplicationResponseBodyApplication {
	return s.Application
}

func (s *UpdateApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationResponseBody) SetApplication(v *UpdateApplicationResponseBodyApplication) *UpdateApplicationResponseBody {
	s.Application = v
	return s
}

func (s *UpdateApplicationResponseBody) SetRequestId(v string) *UpdateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationResponseBodyApplication struct {
	// The time when the application was created.
	//
	// example:
	//
	// 2021-09-07T09:17:46Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// test application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application name.
	//
	// example:
	//
	// My-Application
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The time when the application was updated.
	//
	// example:
	//
	// 2021-09-07T10:17:46Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s UpdateApplicationResponseBodyApplication) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBodyApplication) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *UpdateApplicationResponseBodyApplication) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationResponseBodyApplication) GetName() *string {
	return s.Name
}

func (s *UpdateApplicationResponseBodyApplication) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateApplicationResponseBodyApplication) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *UpdateApplicationResponseBodyApplication) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *UpdateApplicationResponseBodyApplication) SetCreatedDate(v string) *UpdateApplicationResponseBodyApplication {
	s.CreatedDate = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetDescription(v string) *UpdateApplicationResponseBodyApplication {
	s.Description = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetName(v string) *UpdateApplicationResponseBodyApplication {
	s.Name = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetResourceGroupId(v string) *UpdateApplicationResponseBodyApplication {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetTags(v map[string]interface{}) *UpdateApplicationResponseBodyApplication {
	s.Tags = v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetUpdatedDate(v string) *UpdateApplicationResponseBodyApplication {
	s.UpdatedDate = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) Validate() error {
	return dara.Validate(s)
}
