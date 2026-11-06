// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudPermissionTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRequest(v bool) *UpdateCloudPermissionTemplateRequest
	GetAllowRequest() *bool
	SetApplicableScope(v *UpdateCloudPermissionTemplateRequestApplicableScope) *UpdateCloudPermissionTemplateRequest
	GetApplicableScope() *UpdateCloudPermissionTemplateRequestApplicableScope
	SetClientToken(v string) *UpdateCloudPermissionTemplateRequest
	GetClientToken() *string
	SetCloudAccountId(v string) *UpdateCloudPermissionTemplateRequest
	GetCloudAccountId() *string
	SetCloudAccountResources(v []*UpdateCloudPermissionTemplateRequestCloudAccountResources) *UpdateCloudPermissionTemplateRequest
	GetCloudAccountResources() []*UpdateCloudPermissionTemplateRequestCloudAccountResources
	SetCloudPermissionTemplateId(v string) *UpdateCloudPermissionTemplateRequest
	GetCloudPermissionTemplateId() *string
	SetCloudPermissionTemplateName(v string) *UpdateCloudPermissionTemplateRequest
	GetCloudPermissionTemplateName() *string
	SetDescription(v string) *UpdateCloudPermissionTemplateRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateCloudPermissionTemplateRequest
	GetInstanceId() *string
}

type UpdateCloudPermissionTemplateRequest struct {
	// Specifies whether users are allowed to apply.
	//
	// example:
	//
	// true
	AllowRequest *bool `json:"AllowRequest,omitempty" xml:"AllowRequest,omitempty"`
	// The applicable scope of the permission template, including the scope type and specified users, groups, or organizational units.
	ApplicableScope *UpdateCloudPermissionTemplateRequestApplicableScope `json:"ApplicableScope,omitempty" xml:"ApplicableScope,omitempty" type:"Struct"`
	// The idempotency token.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cloud account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac_mkv7rgt4d7i4u7zqtzev2mxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The list of cloud account resources.
	CloudAccountResources []*UpdateCloudPermissionTemplateRequestCloudAccountResources `json:"CloudAccountResources,omitempty" xml:"CloudAccountResources,omitempty" type:"Repeated"`
	// The permission template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpt_mkv7rgt4d7i4u7zqtzev2mxxxx
	CloudPermissionTemplateId *string `json:"CloudPermissionTemplateId,omitempty" xml:"CloudPermissionTemplateId,omitempty"`
	// The name of the permission template. The name can be up to 128 characters in length.
	//
	// example:
	//
	// dev-default
	CloudPermissionTemplateName *string `json:"CloudPermissionTemplateName,omitempty" xml:"CloudPermissionTemplateName,omitempty"`
	// The description of the permission template. The description can be up to 128 characters in length.
	//
	// example:
	//
	// instance_for_test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the IDaaS EIAM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateCloudPermissionTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudPermissionTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudPermissionTemplateRequest) GetAllowRequest() *bool {
	return s.AllowRequest
}

func (s *UpdateCloudPermissionTemplateRequest) GetApplicableScope() *UpdateCloudPermissionTemplateRequestApplicableScope {
	return s.ApplicableScope
}

func (s *UpdateCloudPermissionTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudPermissionTemplateRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *UpdateCloudPermissionTemplateRequest) GetCloudAccountResources() []*UpdateCloudPermissionTemplateRequestCloudAccountResources {
	return s.CloudAccountResources
}

func (s *UpdateCloudPermissionTemplateRequest) GetCloudPermissionTemplateId() *string {
	return s.CloudPermissionTemplateId
}

func (s *UpdateCloudPermissionTemplateRequest) GetCloudPermissionTemplateName() *string {
	return s.CloudPermissionTemplateName
}

func (s *UpdateCloudPermissionTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateCloudPermissionTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCloudPermissionTemplateRequest) SetAllowRequest(v bool) *UpdateCloudPermissionTemplateRequest {
	s.AllowRequest = &v
	return s
}

func (s *UpdateCloudPermissionTemplateRequest) SetApplicableScope(v *UpdateCloudPermissionTemplateRequestApplicableScope) *UpdateCloudPermissionTemplateRequest {
	s.ApplicableScope = v
	return s
}

func (s *UpdateCloudPermissionTemplateRequest) SetClientToken(v string) *UpdateCloudPermissionTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudPermissionTemplateRequest) SetCloudAccountId(v string) *UpdateCloudPermissionTemplateRequest {
	s.CloudAccountId = &v
	return s
}

func (s *UpdateCloudPermissionTemplateRequest) SetCloudAccountResources(v []*UpdateCloudPermissionTemplateRequestCloudAccountResources) *UpdateCloudPermissionTemplateRequest {
	s.CloudAccountResources = v
	return s
}

func (s *UpdateCloudPermissionTemplateRequest) SetCloudPermissionTemplateId(v string) *UpdateCloudPermissionTemplateRequest {
	s.CloudPermissionTemplateId = &v
	return s
}

func (s *UpdateCloudPermissionTemplateRequest) SetCloudPermissionTemplateName(v string) *UpdateCloudPermissionTemplateRequest {
	s.CloudPermissionTemplateName = &v
	return s
}

func (s *UpdateCloudPermissionTemplateRequest) SetDescription(v string) *UpdateCloudPermissionTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateCloudPermissionTemplateRequest) SetInstanceId(v string) *UpdateCloudPermissionTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCloudPermissionTemplateRequest) Validate() error {
	if s.ApplicableScope != nil {
		if err := s.ApplicableScope.Validate(); err != nil {
			return err
		}
	}
	if s.CloudAccountResources != nil {
		for _, item := range s.CloudAccountResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateCloudPermissionTemplateRequestApplicableScope struct {
	// The list of group IDs.
	//
	// example:
	//
	// group_d6sbsusdfa4h66ec3il7yxxxx
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The list of organizational unit IDs.
	//
	// example:
	//
	// ou_d6sbsuudfgf4h66ec3il7yxxxx
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
	// The application scope type. Valid values: global and custom.
	//
	// example:
	//
	// global
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// The list of account IDs.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s UpdateCloudPermissionTemplateRequestApplicableScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudPermissionTemplateRequestApplicableScope) GoString() string {
	return s.String()
}

func (s *UpdateCloudPermissionTemplateRequestApplicableScope) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *UpdateCloudPermissionTemplateRequestApplicableScope) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
}

func (s *UpdateCloudPermissionTemplateRequestApplicableScope) GetScopeType() *string {
	return s.ScopeType
}

func (s *UpdateCloudPermissionTemplateRequestApplicableScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *UpdateCloudPermissionTemplateRequestApplicableScope) SetGroupIds(v []*string) *UpdateCloudPermissionTemplateRequestApplicableScope {
	s.GroupIds = v
	return s
}

func (s *UpdateCloudPermissionTemplateRequestApplicableScope) SetOrganizationalUnitIds(v []*string) *UpdateCloudPermissionTemplateRequestApplicableScope {
	s.OrganizationalUnitIds = v
	return s
}

func (s *UpdateCloudPermissionTemplateRequestApplicableScope) SetScopeType(v string) *UpdateCloudPermissionTemplateRequestApplicableScope {
	s.ScopeType = &v
	return s
}

func (s *UpdateCloudPermissionTemplateRequestApplicableScope) SetUserIds(v []*string) *UpdateCloudPermissionTemplateRequestApplicableScope {
	s.UserIds = v
	return s
}

func (s *UpdateCloudPermissionTemplateRequestApplicableScope) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudPermissionTemplateRequestCloudAccountResources struct {
	// The cloud account resource ID.
	//
	// example:
	//
	// car_01lap4os3aaerckla05vftnsl8edo
	CloudAccountResourceId *string `json:"CloudAccountResourceId,omitempty" xml:"CloudAccountResourceId,omitempty"`
}

func (s UpdateCloudPermissionTemplateRequestCloudAccountResources) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudPermissionTemplateRequestCloudAccountResources) GoString() string {
	return s.String()
}

func (s *UpdateCloudPermissionTemplateRequestCloudAccountResources) GetCloudAccountResourceId() *string {
	return s.CloudAccountResourceId
}

func (s *UpdateCloudPermissionTemplateRequestCloudAccountResources) SetCloudAccountResourceId(v string) *UpdateCloudPermissionTemplateRequestCloudAccountResources {
	s.CloudAccountResourceId = &v
	return s
}

func (s *UpdateCloudPermissionTemplateRequestCloudAccountResources) Validate() error {
	return dara.Validate(s)
}
