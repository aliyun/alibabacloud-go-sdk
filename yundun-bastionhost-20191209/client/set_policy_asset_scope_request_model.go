// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyAssetScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v []*SetPolicyAssetScopeRequestDatabases) *SetPolicyAssetScopeRequest
	GetDatabases() []*SetPolicyAssetScopeRequestDatabases
	SetHostGroups(v []*SetPolicyAssetScopeRequestHostGroups) *SetPolicyAssetScopeRequest
	GetHostGroups() []*SetPolicyAssetScopeRequestHostGroups
	SetHosts(v []*SetPolicyAssetScopeRequestHosts) *SetPolicyAssetScopeRequest
	GetHosts() []*SetPolicyAssetScopeRequestHosts
	SetInstanceId(v string) *SetPolicyAssetScopeRequest
	GetInstanceId() *string
	SetPolicyId(v string) *SetPolicyAssetScopeRequest
	GetPolicyId() *string
	SetRegionId(v string) *SetPolicyAssetScopeRequest
	GetRegionId() *string
	SetScopeType(v string) *SetPolicyAssetScopeRequest
	GetScopeType() *string
}

type SetPolicyAssetScopeRequest struct {
	// The databases to which the control policy applies.
	//
	// >  This parameter is required if ScopeType is set to Database. You can specify up to 500 databases.
	Databases []*SetPolicyAssetScopeRequestDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The asset groups to which the control policy applies.
	//
	// > This parameter is required if ScopeType is set to HostGroup. You can specify up to 100 asset groups.
	HostGroups []*SetPolicyAssetScopeRequestHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
	// The hosts to which the control policy applies.
	//
	// > This parameter is required if ScopeType is set to Host. You can specify up to 500 hosts.
	Hosts []*SetPolicyAssetScopeRequestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The bastion host ID.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy that you want to modify.
	//
	// >  You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to query the control policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scope of assets to which the control policy applies. Valid values:
	//
	// 	- **All**: The control policy applies to all assets.
	//
	// 	- **Host**: The control policy applies to specified hosts.
	//
	// 	- **Database**: The control policy applies to specified databases.
	//
	// 	- **HostGroup**: The control policy applies to specified asset groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// All
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s SetPolicyAssetScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyAssetScopeRequest) GoString() string {
	return s.String()
}

func (s *SetPolicyAssetScopeRequest) GetDatabases() []*SetPolicyAssetScopeRequestDatabases {
	return s.Databases
}

func (s *SetPolicyAssetScopeRequest) GetHostGroups() []*SetPolicyAssetScopeRequestHostGroups {
	return s.HostGroups
}

func (s *SetPolicyAssetScopeRequest) GetHosts() []*SetPolicyAssetScopeRequestHosts {
	return s.Hosts
}

func (s *SetPolicyAssetScopeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPolicyAssetScopeRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *SetPolicyAssetScopeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetPolicyAssetScopeRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *SetPolicyAssetScopeRequest) SetDatabases(v []*SetPolicyAssetScopeRequestDatabases) *SetPolicyAssetScopeRequest {
	s.Databases = v
	return s
}

func (s *SetPolicyAssetScopeRequest) SetHostGroups(v []*SetPolicyAssetScopeRequestHostGroups) *SetPolicyAssetScopeRequest {
	s.HostGroups = v
	return s
}

func (s *SetPolicyAssetScopeRequest) SetHosts(v []*SetPolicyAssetScopeRequestHosts) *SetPolicyAssetScopeRequest {
	s.Hosts = v
	return s
}

func (s *SetPolicyAssetScopeRequest) SetInstanceId(v string) *SetPolicyAssetScopeRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPolicyAssetScopeRequest) SetPolicyId(v string) *SetPolicyAssetScopeRequest {
	s.PolicyId = &v
	return s
}

func (s *SetPolicyAssetScopeRequest) SetRegionId(v string) *SetPolicyAssetScopeRequest {
	s.RegionId = &v
	return s
}

func (s *SetPolicyAssetScopeRequest) SetScopeType(v string) *SetPolicyAssetScopeRequest {
	s.ScopeType = &v
	return s
}

func (s *SetPolicyAssetScopeRequest) Validate() error {
	return dara.Validate(s)
}

type SetPolicyAssetScopeRequestDatabases struct {
	// The scope of database accounts to which the control policy applies. Valid values:
	//
	// 	- **All**: The control policy applies to all database accounts of the database.
	//
	// 	- **AccountId**: The control policy applies to specified database accounts of the database.
	//
	// example:
	//
	// AccountId
	AccountScopeType *string `json:"AccountScopeType,omitempty" xml:"AccountScopeType,omitempty"`
	// The database accounts to which the control policy applies.
	//
	// >  This parameter is required if AccountScopeType is set to AccountId.
	DatabaseAccountIds []*string `json:"DatabaseAccountIds,omitempty" xml:"DatabaseAccountIds,omitempty" type:"Repeated"`
	// The database ID.
	//
	// example:
	//
	// 3
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
}

func (s SetPolicyAssetScopeRequestDatabases) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyAssetScopeRequestDatabases) GoString() string {
	return s.String()
}

func (s *SetPolicyAssetScopeRequestDatabases) GetAccountScopeType() *string {
	return s.AccountScopeType
}

func (s *SetPolicyAssetScopeRequestDatabases) GetDatabaseAccountIds() []*string {
	return s.DatabaseAccountIds
}

func (s *SetPolicyAssetScopeRequestDatabases) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *SetPolicyAssetScopeRequestDatabases) SetAccountScopeType(v string) *SetPolicyAssetScopeRequestDatabases {
	s.AccountScopeType = &v
	return s
}

func (s *SetPolicyAssetScopeRequestDatabases) SetDatabaseAccountIds(v []*string) *SetPolicyAssetScopeRequestDatabases {
	s.DatabaseAccountIds = v
	return s
}

func (s *SetPolicyAssetScopeRequestDatabases) SetDatabaseId(v string) *SetPolicyAssetScopeRequestDatabases {
	s.DatabaseId = &v
	return s
}

func (s *SetPolicyAssetScopeRequestDatabases) Validate() error {
	return dara.Validate(s)
}

type SetPolicyAssetScopeRequestHostGroups struct {
	// The asset accounts to which the control policy applies.
	//
	// > This parameter is required if AccountScopeType is set to AccountName.
	AccountNames []*string `json:"AccountNames,omitempty" xml:"AccountNames,omitempty" type:"Repeated"`
	// The scope of asset accounts to which the control policy applies. Valid values:
	//
	// 	- **All**: The control policy applies to all accounts in the asset group.
	//
	// 	- **AccountName**: The control policy applies to specified accounts in the asset group.
	//
	// example:
	//
	// All
	AccountScopeType *string `json:"AccountScopeType,omitempty" xml:"AccountScopeType,omitempty"`
	// The asset group ID.
	//
	// example:
	//
	// 86
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s SetPolicyAssetScopeRequestHostGroups) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyAssetScopeRequestHostGroups) GoString() string {
	return s.String()
}

func (s *SetPolicyAssetScopeRequestHostGroups) GetAccountNames() []*string {
	return s.AccountNames
}

func (s *SetPolicyAssetScopeRequestHostGroups) GetAccountScopeType() *string {
	return s.AccountScopeType
}

func (s *SetPolicyAssetScopeRequestHostGroups) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *SetPolicyAssetScopeRequestHostGroups) SetAccountNames(v []*string) *SetPolicyAssetScopeRequestHostGroups {
	s.AccountNames = v
	return s
}

func (s *SetPolicyAssetScopeRequestHostGroups) SetAccountScopeType(v string) *SetPolicyAssetScopeRequestHostGroups {
	s.AccountScopeType = &v
	return s
}

func (s *SetPolicyAssetScopeRequestHostGroups) SetHostGroupId(v string) *SetPolicyAssetScopeRequestHostGroups {
	s.HostGroupId = &v
	return s
}

func (s *SetPolicyAssetScopeRequestHostGroups) Validate() error {
	return dara.Validate(s)
}

type SetPolicyAssetScopeRequestHosts struct {
	// The scope of host accounts to which the control policy applies. Valid values:
	//
	// 	- **All**: The control policy applies to all accounts of the host.
	//
	// 	- **AccountId**: The control policy applies specified accounts of the host.
	//
	// example:
	//
	// All
	AccountScopeType *string `json:"AccountScopeType,omitempty" xml:"AccountScopeType,omitempty"`
	// The host accounts to which the control policy applies.
	//
	// > This parameter is required if AccountScopeType is set to AccountId.
	HostAccountIds []*string `json:"HostAccountIds,omitempty" xml:"HostAccountIds,omitempty" type:"Repeated"`
	// The host ID.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s SetPolicyAssetScopeRequestHosts) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyAssetScopeRequestHosts) GoString() string {
	return s.String()
}

func (s *SetPolicyAssetScopeRequestHosts) GetAccountScopeType() *string {
	return s.AccountScopeType
}

func (s *SetPolicyAssetScopeRequestHosts) GetHostAccountIds() []*string {
	return s.HostAccountIds
}

func (s *SetPolicyAssetScopeRequestHosts) GetHostId() *string {
	return s.HostId
}

func (s *SetPolicyAssetScopeRequestHosts) SetAccountScopeType(v string) *SetPolicyAssetScopeRequestHosts {
	s.AccountScopeType = &v
	return s
}

func (s *SetPolicyAssetScopeRequestHosts) SetHostAccountIds(v []*string) *SetPolicyAssetScopeRequestHosts {
	s.HostAccountIds = v
	return s
}

func (s *SetPolicyAssetScopeRequestHosts) SetHostId(v string) *SetPolicyAssetScopeRequestHosts {
	s.HostId = &v
	return s
}

func (s *SetPolicyAssetScopeRequestHosts) Validate() error {
	return dara.Validate(s)
}
