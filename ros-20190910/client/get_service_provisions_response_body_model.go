// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceProvisionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetServiceProvisionsResponseBody
	GetRequestId() *string
	SetServiceProvisions(v []*GetServiceProvisionsResponseBodyServiceProvisions) *GetServiceProvisionsResponseBody
	GetServiceProvisions() []*GetServiceProvisionsResponseBodyServiceProvisions
}

type GetServiceProvisionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A9938C04-26AD-5EFA-B394-EA7866842DFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the services.
	ServiceProvisions []*GetServiceProvisionsResponseBodyServiceProvisions `json:"ServiceProvisions,omitempty" xml:"ServiceProvisions,omitempty" type:"Repeated"`
}

func (s GetServiceProvisionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvisionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceProvisionsResponseBody) GetServiceProvisions() []*GetServiceProvisionsResponseBodyServiceProvisions {
	return s.ServiceProvisions
}

func (s *GetServiceProvisionsResponseBody) SetRequestId(v string) *GetServiceProvisionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceProvisionsResponseBody) SetServiceProvisions(v []*GetServiceProvisionsResponseBodyServiceProvisions) *GetServiceProvisionsResponseBody {
	s.ServiceProvisions = v
	return s
}

func (s *GetServiceProvisionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetServiceProvisionsResponseBodyServiceProvisions struct {
	// Indicates whether automatic activation for the service is defined in the template. Valid values:
	//
	// 	- true: Automatic activation for the service is defined in the template.
	//
	// 	- false: Manual activation for the service is defined in the template.
	//
	// example:
	//
	// false
	AutoEnableService   *bool                                                                   `json:"AutoEnableService,omitempty" xml:"AutoEnableService,omitempty"`
	CommodityProvisions []*GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions `json:"CommodityProvisions,omitempty" xml:"CommodityProvisions,omitempty" type:"Repeated"`
	// The names of the services on which the service that is queried depends.
	DependentServiceNames []*string `json:"DependentServiceNames,omitempty" xml:"DependentServiceNames,omitempty" type:"Repeated"`
	// The URL that points to the activation page of the service.
	//
	// > This parameter is returned if Status is set to Disabled.
	//
	// example:
	//
	// https://common-buy.aliyun.com/?commodityCode=nas
	EnableURL *string `json:"EnableURL,omitempty" xml:"EnableURL,omitempty"`
	// The information about the RAM roles of the service. If this parameter is empty, no RAM role is associated with the service.
	RoleProvision *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision `json:"RoleProvision,omitempty" xml:"RoleProvision,omitempty" type:"Struct"`
	// The service name.
	//
	// example:
	//
	// EHPC
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The activation status of the service. Valid values:
	//
	// 	- Enabled: The service is activated.
	//
	// 	- Disabled: The service is not activated.
	//
	// 	- Unknown: The activation status of the service is unknown.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the service is in the Disabled or Unknown state.
	//
	// > This parameter is returned if Status is set to Disabled or Unknown.
	//
	// example:
	//
	// No permission.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisions) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisions) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) GetAutoEnableService() *bool {
	return s.AutoEnableService
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) GetCommodityProvisions() []*GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions {
	return s.CommodityProvisions
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) GetDependentServiceNames() []*string {
	return s.DependentServiceNames
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) GetEnableURL() *string {
	return s.EnableURL
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) GetRoleProvision() *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision {
	return s.RoleProvision
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) GetStatus() *string {
	return s.Status
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetAutoEnableService(v bool) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.AutoEnableService = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetCommodityProvisions(v []*GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.CommodityProvisions = v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetDependentServiceNames(v []*string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.DependentServiceNames = v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetEnableURL(v string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.EnableURL = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetRoleProvision(v *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.RoleProvision = v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetServiceName(v string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.ServiceName = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetStatus(v string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.Status = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetStatusReason(v string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.StatusReason = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) Validate() error {
	return dara.Validate(s)
}

type GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	EnableURL     *string `json:"EnableURL,omitempty" xml:"EnableURL,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions) GetEnableURL() *string {
	return s.EnableURL
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions) GetStatus() *string {
	return s.Status
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions) SetCommodityCode(v string) *GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions {
	s.CommodityCode = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions) SetEnableURL(v string) *GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions {
	s.EnableURL = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions) SetStatus(v string) *GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions {
	s.Status = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions) Validate() error {
	return dara.Validate(s)
}

type GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision struct {
	// The authorization URL of the RAM role.
	//
	// > This parameter is returned if Created is set to false.
	//
	// example:
	//
	// https://ehpc.console.aliyun.com/
	AuthorizationURL *string `json:"AuthorizationURL,omitempty" xml:"AuthorizationURL,omitempty"`
	// The RAM roles of the service.
	Roles []*GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) GetAuthorizationURL() *string {
	return s.AuthorizationURL
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) GetRoles() []*GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	return s.Roles
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) SetAuthorizationURL(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision {
	s.AuthorizationURL = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) SetRoles(v []*GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision {
	s.Roles = v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) Validate() error {
	return dara.Validate(s)
}

type GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles struct {
	// The information about the API operation that is used to create the RAM role.
	ApiForCreation *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation `json:"ApiForCreation,omitempty" xml:"ApiForCreation,omitempty" type:"Struct"`
	// Indicates whether the RAM role is created. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Created *bool `json:"Created,omitempty" xml:"Created,omitempty"`
	// The purpose for which the RAM role is used. Default value: Default. A value of Default indicates that the RAM role is the default role of the service.
	//
	// example:
	//
	// Default
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The name of the role.
	//
	// example:
	//
	// AliyunServiceRoleForEHPC
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) GetApiForCreation() *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	return s.ApiForCreation
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) GetCreated() *bool {
	return s.Created
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) GetFunction() *string {
	return s.Function
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) GetRoleName() *string {
	return s.RoleName
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) SetApiForCreation(v *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	s.ApiForCreation = v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) SetCreated(v bool) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	s.Created = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) SetFunction(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	s.Function = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) SetRoleName(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	s.RoleName = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) Validate() error {
	return dara.Validate(s)
}

type GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation struct {
	// The name of the API operation.
	//
	// example:
	//
	// CreateServiceLinkedRole
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The ID of the Alibaba Cloud service to which the API operation belongs.
	//
	// example:
	//
	// rds
	ApiProductId *string `json:"ApiProductId,omitempty" xml:"ApiProductId,omitempty"`
	// The type of the API operation. Valid values:
	//
	// 	- Open: public
	//
	// 	- Inner: private
	//
	// example:
	//
	// Open
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// The parameters of the API operation. If a parameter is a variable, use the ${Variable name} format. Only the following variable is supported: ${RegionId}.
	//
	// example:
	//
	// {   "ServiceLinkedRole": "AliyunServiceRoleForRdsPgsqlOnEcs",   "RegionId": "${RegionId}" }
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) GetApiName() *string {
	return s.ApiName
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) GetApiProductId() *string {
	return s.ApiProductId
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) GetApiType() *string {
	return s.ApiType
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) SetApiName(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	s.ApiName = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) SetApiProductId(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	s.ApiProductId = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) SetApiType(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	s.ApiType = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) SetParameters(v map[string]interface{}) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	s.Parameters = v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) Validate() error {
	return dara.Validate(s)
}
