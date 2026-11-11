// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudPermissionTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRequest(v bool) *CreateCloudPermissionTemplateRequest
	GetAllowRequest() *bool
	SetApplicableScope(v *CreateCloudPermissionTemplateRequestApplicableScope) *CreateCloudPermissionTemplateRequest
	GetApplicableScope() *CreateCloudPermissionTemplateRequestApplicableScope
	SetClientToken(v string) *CreateCloudPermissionTemplateRequest
	GetClientToken() *string
	SetCloudAccountId(v string) *CreateCloudPermissionTemplateRequest
	GetCloudAccountId() *string
	SetCloudAccountResources(v []*CreateCloudPermissionTemplateRequestCloudAccountResources) *CreateCloudPermissionTemplateRequest
	GetCloudAccountResources() []*CreateCloudPermissionTemplateRequestCloudAccountResources
	SetCloudPermissionTemplateName(v string) *CreateCloudPermissionTemplateRequest
	GetCloudPermissionTemplateName() *string
	SetCloudPermissionTemplateUsageType(v string) *CreateCloudPermissionTemplateRequest
	GetCloudPermissionTemplateUsageType() *string
	SetDescription(v string) *CreateCloudPermissionTemplateRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateCloudPermissionTemplateRequest
	GetInstanceId() *string
}

type CreateCloudPermissionTemplateRequest struct {
	// Specifies whether users are allowed to apply for the permission template. Default value: true.
	//
	// example:
	//
	// true
	AllowRequest *bool `json:"AllowRequest,omitempty" xml:"AllowRequest,omitempty"`
	// The applicable scope of the permission template, including the scope type and the specified users, groups, or organizational units.
	ApplicableScope *CreateCloudPermissionTemplateRequestApplicableScope `json:"ApplicableScope,omitempty" xml:"ApplicableScope,omitempty" type:"Struct"`
	// Ensures the idempotence of the request. You can generate the value from your client to ensure that the value is unique across different requests. The client token supports only ASCII characters and can be up to 64 characters in length.
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
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The list of cloud account resource policies.
	CloudAccountResources []*CreateCloudPermissionTemplateRequestCloudAccountResources `json:"CloudAccountResources,omitempty" xml:"CloudAccountResources,omitempty" type:"Repeated"`
	// The name of the permission template. The name can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	CloudPermissionTemplateName *string `json:"CloudPermissionTemplateName,omitempty" xml:"CloudPermissionTemplateName,omitempty"`
	// The usage type of the permission template. Valid values: user (privilege escalation for human users) and application (privilege escalation for M2M applications). Default value: user.
	//
	// example:
	//
	// user
	CloudPermissionTemplateUsageType *string `json:"CloudPermissionTemplateUsageType,omitempty" xml:"CloudPermissionTemplateUsageType,omitempty"`
	// The description of the permission template. The description can be up to 128 characters in length.
	//
	// example:
	//
	// this is a test.
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

func (s CreateCloudPermissionTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudPermissionTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudPermissionTemplateRequest) GetAllowRequest() *bool {
	return s.AllowRequest
}

func (s *CreateCloudPermissionTemplateRequest) GetApplicableScope() *CreateCloudPermissionTemplateRequestApplicableScope {
	return s.ApplicableScope
}

func (s *CreateCloudPermissionTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCloudPermissionTemplateRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *CreateCloudPermissionTemplateRequest) GetCloudAccountResources() []*CreateCloudPermissionTemplateRequestCloudAccountResources {
	return s.CloudAccountResources
}

func (s *CreateCloudPermissionTemplateRequest) GetCloudPermissionTemplateName() *string {
	return s.CloudPermissionTemplateName
}

func (s *CreateCloudPermissionTemplateRequest) GetCloudPermissionTemplateUsageType() *string {
	return s.CloudPermissionTemplateUsageType
}

func (s *CreateCloudPermissionTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCloudPermissionTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCloudPermissionTemplateRequest) SetAllowRequest(v bool) *CreateCloudPermissionTemplateRequest {
	s.AllowRequest = &v
	return s
}

func (s *CreateCloudPermissionTemplateRequest) SetApplicableScope(v *CreateCloudPermissionTemplateRequestApplicableScope) *CreateCloudPermissionTemplateRequest {
	s.ApplicableScope = v
	return s
}

func (s *CreateCloudPermissionTemplateRequest) SetClientToken(v string) *CreateCloudPermissionTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCloudPermissionTemplateRequest) SetCloudAccountId(v string) *CreateCloudPermissionTemplateRequest {
	s.CloudAccountId = &v
	return s
}

func (s *CreateCloudPermissionTemplateRequest) SetCloudAccountResources(v []*CreateCloudPermissionTemplateRequestCloudAccountResources) *CreateCloudPermissionTemplateRequest {
	s.CloudAccountResources = v
	return s
}

func (s *CreateCloudPermissionTemplateRequest) SetCloudPermissionTemplateName(v string) *CreateCloudPermissionTemplateRequest {
	s.CloudPermissionTemplateName = &v
	return s
}

func (s *CreateCloudPermissionTemplateRequest) SetCloudPermissionTemplateUsageType(v string) *CreateCloudPermissionTemplateRequest {
	s.CloudPermissionTemplateUsageType = &v
	return s
}

func (s *CreateCloudPermissionTemplateRequest) SetDescription(v string) *CreateCloudPermissionTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateCloudPermissionTemplateRequest) SetInstanceId(v string) *CreateCloudPermissionTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCloudPermissionTemplateRequest) Validate() error {
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

type CreateCloudPermissionTemplateRequestApplicableScope struct {
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
	// The type of the applicable scope. Valid values: global and custom.
	//
	// example:
	//
	// global
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// The list of user IDs.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s CreateCloudPermissionTemplateRequestApplicableScope) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudPermissionTemplateRequestApplicableScope) GoString() string {
	return s.String()
}

func (s *CreateCloudPermissionTemplateRequestApplicableScope) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *CreateCloudPermissionTemplateRequestApplicableScope) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
}

func (s *CreateCloudPermissionTemplateRequestApplicableScope) GetScopeType() *string {
	return s.ScopeType
}

func (s *CreateCloudPermissionTemplateRequestApplicableScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *CreateCloudPermissionTemplateRequestApplicableScope) SetGroupIds(v []*string) *CreateCloudPermissionTemplateRequestApplicableScope {
	s.GroupIds = v
	return s
}

func (s *CreateCloudPermissionTemplateRequestApplicableScope) SetOrganizationalUnitIds(v []*string) *CreateCloudPermissionTemplateRequestApplicableScope {
	s.OrganizationalUnitIds = v
	return s
}

func (s *CreateCloudPermissionTemplateRequestApplicableScope) SetScopeType(v string) *CreateCloudPermissionTemplateRequestApplicableScope {
	s.ScopeType = &v
	return s
}

func (s *CreateCloudPermissionTemplateRequestApplicableScope) SetUserIds(v []*string) *CreateCloudPermissionTemplateRequestApplicableScope {
	s.UserIds = v
	return s
}

func (s *CreateCloudPermissionTemplateRequestApplicableScope) Validate() error {
	return dara.Validate(s)
}

type CreateCloudPermissionTemplateRequestCloudAccountResources struct {
	// The ID of the cloud account resource.
	//
	// example:
	//
	// car_01l8a1l906kqtkacdfp5dxxxxx
	CloudAccountResourceId *string `json:"CloudAccountResourceId,omitempty" xml:"CloudAccountResourceId,omitempty"`
}

func (s CreateCloudPermissionTemplateRequestCloudAccountResources) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudPermissionTemplateRequestCloudAccountResources) GoString() string {
	return s.String()
}

func (s *CreateCloudPermissionTemplateRequestCloudAccountResources) GetCloudAccountResourceId() *string {
	return s.CloudAccountResourceId
}

func (s *CreateCloudPermissionTemplateRequestCloudAccountResources) SetCloudAccountResourceId(v string) *CreateCloudPermissionTemplateRequestCloudAccountResources {
	s.CloudAccountResourceId = &v
	return s
}

func (s *CreateCloudPermissionTemplateRequestCloudAccountResources) Validate() error {
	return dara.Validate(s)
}
