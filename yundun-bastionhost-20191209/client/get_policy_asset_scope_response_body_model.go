// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyAssetScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetScope(v *GetPolicyAssetScopeResponseBodyAssetScope) *GetPolicyAssetScopeResponseBody
	GetAssetScope() *GetPolicyAssetScopeResponseBodyAssetScope
	SetRequestId(v string) *GetPolicyAssetScopeResponseBody
	GetRequestId() *string
}

type GetPolicyAssetScopeResponseBody struct {
	// The assets to which the control policy applies.
	AssetScope *GetPolicyAssetScopeResponseBodyAssetScope `json:"AssetScope,omitempty" xml:"AssetScope,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPolicyAssetScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyAssetScopeResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyAssetScopeResponseBody) GetAssetScope() *GetPolicyAssetScopeResponseBodyAssetScope {
	return s.AssetScope
}

func (s *GetPolicyAssetScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPolicyAssetScopeResponseBody) SetAssetScope(v *GetPolicyAssetScopeResponseBodyAssetScope) *GetPolicyAssetScopeResponseBody {
	s.AssetScope = v
	return s
}

func (s *GetPolicyAssetScopeResponseBody) SetRequestId(v string) *GetPolicyAssetScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolicyAssetScopeResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPolicyAssetScopeResponseBodyAssetScope struct {
	// The databases and database accounts to which the control policy applies.
	Databases []*GetPolicyAssetScopeResponseBodyAssetScopeDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The asset groups and asset accounts to which the control policy applies.
	HostGroups []*GetPolicyAssetScopeResponseBodyAssetScopeHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
	// The hosts and host accounts to which the control policy applies.
	Hosts []*GetPolicyAssetScopeResponseBodyAssetScopeHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The scope of assets to which the control policy applies.
	//
	// 	- If **All*	- is returned for this parameter, the control policy applies to all assets.
	//
	// 	- If no value is returned for this parameter, the control policy applies to the assets specified in the return values of Databases, HostGroups, and Hosts.
	//
	// example:
	//
	// All
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s GetPolicyAssetScopeResponseBodyAssetScope) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyAssetScopeResponseBodyAssetScope) GoString() string {
	return s.String()
}

func (s *GetPolicyAssetScopeResponseBodyAssetScope) GetDatabases() []*GetPolicyAssetScopeResponseBodyAssetScopeDatabases {
	return s.Databases
}

func (s *GetPolicyAssetScopeResponseBodyAssetScope) GetHostGroups() []*GetPolicyAssetScopeResponseBodyAssetScopeHostGroups {
	return s.HostGroups
}

func (s *GetPolicyAssetScopeResponseBodyAssetScope) GetHosts() []*GetPolicyAssetScopeResponseBodyAssetScopeHosts {
	return s.Hosts
}

func (s *GetPolicyAssetScopeResponseBodyAssetScope) GetScopeType() *string {
	return s.ScopeType
}

func (s *GetPolicyAssetScopeResponseBodyAssetScope) SetDatabases(v []*GetPolicyAssetScopeResponseBodyAssetScopeDatabases) *GetPolicyAssetScopeResponseBodyAssetScope {
	s.Databases = v
	return s
}

func (s *GetPolicyAssetScopeResponseBodyAssetScope) SetHostGroups(v []*GetPolicyAssetScopeResponseBodyAssetScopeHostGroups) *GetPolicyAssetScopeResponseBodyAssetScope {
	s.HostGroups = v
	return s
}

func (s *GetPolicyAssetScopeResponseBodyAssetScope) SetHosts(v []*GetPolicyAssetScopeResponseBodyAssetScopeHosts) *GetPolicyAssetScopeResponseBodyAssetScope {
	s.Hosts = v
	return s
}

func (s *GetPolicyAssetScopeResponseBodyAssetScope) SetScopeType(v string) *GetPolicyAssetScopeResponseBodyAssetScope {
	s.ScopeType = &v
	return s
}

func (s *GetPolicyAssetScopeResponseBodyAssetScope) Validate() error {
	return dara.Validate(s)
}

type GetPolicyAssetScopeResponseBodyAssetScopeDatabases struct {
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
	// The IDs of database accounts to which the control policy applies.
	DatabaseAccountIds []*string `json:"DatabaseAccountIds,omitempty" xml:"DatabaseAccountIds,omitempty" type:"Repeated"`
	// The database ID.
	//
	// example:
	//
	// 17
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
}

func (s GetPolicyAssetScopeResponseBodyAssetScopeDatabases) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyAssetScopeResponseBodyAssetScopeDatabases) GoString() string {
	return s.String()
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeDatabases) GetAccountScopeType() *string {
	return s.AccountScopeType
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeDatabases) GetDatabaseAccountIds() []*string {
	return s.DatabaseAccountIds
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeDatabases) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeDatabases) SetAccountScopeType(v string) *GetPolicyAssetScopeResponseBodyAssetScopeDatabases {
	s.AccountScopeType = &v
	return s
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeDatabases) SetDatabaseAccountIds(v []*string) *GetPolicyAssetScopeResponseBodyAssetScopeDatabases {
	s.DatabaseAccountIds = v
	return s
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeDatabases) SetDatabaseId(v string) *GetPolicyAssetScopeResponseBodyAssetScopeDatabases {
	s.DatabaseId = &v
	return s
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeDatabases) Validate() error {
	return dara.Validate(s)
}

type GetPolicyAssetScopeResponseBodyAssetScopeHostGroups struct {
	// The asset accounts to which the control policy applies.
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
	// 4
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s GetPolicyAssetScopeResponseBodyAssetScopeHostGroups) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyAssetScopeResponseBodyAssetScopeHostGroups) GoString() string {
	return s.String()
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeHostGroups) GetAccountNames() []*string {
	return s.AccountNames
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeHostGroups) GetAccountScopeType() *string {
	return s.AccountScopeType
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeHostGroups) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeHostGroups) SetAccountNames(v []*string) *GetPolicyAssetScopeResponseBodyAssetScopeHostGroups {
	s.AccountNames = v
	return s
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeHostGroups) SetAccountScopeType(v string) *GetPolicyAssetScopeResponseBodyAssetScopeHostGroups {
	s.AccountScopeType = &v
	return s
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeHostGroups) SetHostGroupId(v string) *GetPolicyAssetScopeResponseBodyAssetScopeHostGroups {
	s.HostGroupId = &v
	return s
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeHostGroups) Validate() error {
	return dara.Validate(s)
}

type GetPolicyAssetScopeResponseBodyAssetScopeHosts struct {
	// The scope of host accounts to which the control policy applies. Valid values:
	//
	// 	- **All**: The control policy applies to all accounts of the host.
	//
	// 	- **AccountId**: The control policy applies to specified accounts of the host.
	//
	// example:
	//
	// All
	AccountScopeType *string `json:"AccountScopeType,omitempty" xml:"AccountScopeType,omitempty"`
	// The host accounts to which the control policy applies.
	HostAccountIds []*string `json:"HostAccountIds,omitempty" xml:"HostAccountIds,omitempty" type:"Repeated"`
	// The host ID.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s GetPolicyAssetScopeResponseBodyAssetScopeHosts) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyAssetScopeResponseBodyAssetScopeHosts) GoString() string {
	return s.String()
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeHosts) GetAccountScopeType() *string {
	return s.AccountScopeType
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeHosts) GetHostAccountIds() []*string {
	return s.HostAccountIds
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeHosts) GetHostId() *string {
	return s.HostId
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeHosts) SetAccountScopeType(v string) *GetPolicyAssetScopeResponseBodyAssetScopeHosts {
	s.AccountScopeType = &v
	return s
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeHosts) SetHostAccountIds(v []*string) *GetPolicyAssetScopeResponseBodyAssetScopeHosts {
	s.HostAccountIds = v
	return s
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeHosts) SetHostId(v string) *GetPolicyAssetScopeResponseBodyAssetScopeHosts {
	s.HostId = &v
	return s
}

func (s *GetPolicyAssetScopeResponseBodyAssetScopeHosts) Validate() error {
	return dara.Validate(s)
}
