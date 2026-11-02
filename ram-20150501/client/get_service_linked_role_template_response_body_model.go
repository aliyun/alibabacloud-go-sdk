// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceLinkedRoleTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetServiceLinkedRoleTemplateResponseBody
	GetRequestId() *string
	SetServiceLinkedRoleTemplate(v *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) *GetServiceLinkedRoleTemplateResponseBody
	GetServiceLinkedRoleTemplate() *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate
}

type GetServiceLinkedRoleTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service-linked role template.
	ServiceLinkedRoleTemplate *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate `json:"ServiceLinkedRoleTemplate,omitempty" xml:"ServiceLinkedRoleTemplate,omitempty" type:"Struct"`
}

func (s GetServiceLinkedRoleTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceLinkedRoleTemplateResponseBody) GetServiceLinkedRoleTemplate() *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate {
	return s.ServiceLinkedRoleTemplate
}

func (s *GetServiceLinkedRoleTemplateResponseBody) SetRequestId(v string) *GetServiceLinkedRoleTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponseBody) SetServiceLinkedRoleTemplate(v *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) *GetServiceLinkedRoleTemplateResponseBody {
	s.ServiceLinkedRoleTemplate = v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponseBody) Validate() error {
	if s.ServiceLinkedRoleTemplate != nil {
		if err := s.ServiceLinkedRoleTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate struct {
	// Indicates whether multiple roles are supported. Valid values:
	//
	// - true: Multiple roles are supported.
	//
	// - false: Multiple roles are not supported.
	//
	// example:
	//
	// false
	MultipleRolesAllowed *bool                                                                              `json:"MultipleRolesAllowed,omitempty" xml:"MultipleRolesAllowed,omitempty"`
	RoleDescriptions     *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions `json:"RoleDescriptions,omitempty" xml:"RoleDescriptions,omitempty" type:"Struct"`
	// The prefix of the role name.
	//
	// example:
	//
	// AliyunServiceRoleForPolarDB
	RoleNamePrefix *string `json:"RoleNamePrefix,omitempty" xml:"RoleNamePrefix,omitempty"`
	// The cloud service name.
	//
	// example:
	//
	// polardb.aliyuncs.com
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The name of the system policy attached to the role.
	//
	// example:
	//
	// AliyunServiceRolePolicyForPolarDB
	SystemPolicyName *string `json:"SystemPolicyName,omitempty" xml:"SystemPolicyName,omitempty"`
}

func (s GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) GetMultipleRolesAllowed() *bool {
	return s.MultipleRolesAllowed
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) GetRoleDescriptions() *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions {
	return s.RoleDescriptions
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) GetRoleNamePrefix() *string {
	return s.RoleNamePrefix
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) GetSystemPolicyName() *string {
	return s.SystemPolicyName
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) SetMultipleRolesAllowed(v bool) *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate {
	s.MultipleRolesAllowed = &v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) SetRoleDescriptions(v *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions) *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate {
	s.RoleDescriptions = v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) SetRoleNamePrefix(v string) *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate {
	s.RoleNamePrefix = &v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) SetServiceName(v string) *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate {
	s.ServiceName = &v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) SetSystemPolicyName(v string) *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate {
	s.SystemPolicyName = &v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplate) Validate() error {
	if s.RoleDescriptions != nil {
		if err := s.RoleDescriptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions struct {
	RoleDescription []*GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription `json:"RoleDescription,omitempty" xml:"RoleDescription,omitempty" type:"Repeated"`
}

func (s GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions) GetRoleDescription() []*GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription {
	return s.RoleDescription
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions) SetRoleDescription(v []*GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription) *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions {
	s.RoleDescription = v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptions) Validate() error {
	if s.RoleDescription != nil {
		for _, item := range s.RoleDescription {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Language    *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription) GetDescription() *string {
	return s.Description
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription) GetLanguage() *string {
	return s.Language
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription) SetDescription(v string) *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription {
	s.Description = &v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription) SetLanguage(v string) *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription {
	s.Language = &v
	return s
}

func (s *GetServiceLinkedRoleTemplateResponseBodyServiceLinkedRoleTemplateRoleDescriptionsRoleDescription) Validate() error {
	return dara.Validate(s)
}
