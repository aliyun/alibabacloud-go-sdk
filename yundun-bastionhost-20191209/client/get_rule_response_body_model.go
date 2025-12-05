// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRuleResponseBody
	GetRequestId() *string
	SetRule(v *GetRuleResponseBodyRule) *GetRuleResponseBody
	GetRule() *GetRuleResponseBodyRule
}

type GetRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned information about the authorization rule.
	Rule *GetRuleResponseBodyRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
}

func (s GetRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRuleResponseBody) GetRule() *GetRuleResponseBodyRule {
	return s.Rule
}

func (s *GetRuleResponseBody) SetRequestId(v string) *GetRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleResponseBody) SetRule(v *GetRuleResponseBodyRule) *GetRuleResponseBody {
	s.Rule = v
	return s
}

func (s *GetRuleResponseBody) Validate() error {
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRuleResponseBodyRule struct {
	// The remarks of the authorization rule.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The databases on which permissions are granted by using the authorization rule.
	Databases []*GetRuleResponseBodyRuleDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The end time of the validity period of the authorization rule. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1698720972
	EffectiveEndTime *string `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The start time of the validity period of the authorization rule. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1687140883
	EffectiveStartTime *string `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The asset groups on which permissions are granted by using the authorization rule.
	HostGroups []*GetRuleResponseBodyRuleHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
	// The information about the hosts that the policy authorizes users to manage.
	Hosts []*GetRuleResponseBodyRuleHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The ID of the authorization rule.
	//
	// example:
	//
	// 3
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the authorization rule.
	//
	// example:
	//
	// rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The authorized user groups.
	UserGroups []*GetRuleResponseBodyRuleUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
	// The authorized users.
	Users []*GetRuleResponseBodyRuleUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s GetRuleResponseBodyRule) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyRule) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyRule) GetComment() *string {
	return s.Comment
}

func (s *GetRuleResponseBodyRule) GetDatabases() []*GetRuleResponseBodyRuleDatabases {
	return s.Databases
}

func (s *GetRuleResponseBodyRule) GetEffectiveEndTime() *string {
	return s.EffectiveEndTime
}

func (s *GetRuleResponseBodyRule) GetEffectiveStartTime() *string {
	return s.EffectiveStartTime
}

func (s *GetRuleResponseBodyRule) GetHostGroups() []*GetRuleResponseBodyRuleHostGroups {
	return s.HostGroups
}

func (s *GetRuleResponseBodyRule) GetHosts() []*GetRuleResponseBodyRuleHosts {
	return s.Hosts
}

func (s *GetRuleResponseBodyRule) GetRuleId() *string {
	return s.RuleId
}

func (s *GetRuleResponseBodyRule) GetRuleName() *string {
	return s.RuleName
}

func (s *GetRuleResponseBodyRule) GetUserGroups() []*GetRuleResponseBodyRuleUserGroups {
	return s.UserGroups
}

func (s *GetRuleResponseBodyRule) GetUsers() []*GetRuleResponseBodyRuleUsers {
	return s.Users
}

func (s *GetRuleResponseBodyRule) SetComment(v string) *GetRuleResponseBodyRule {
	s.Comment = &v
	return s
}

func (s *GetRuleResponseBodyRule) SetDatabases(v []*GetRuleResponseBodyRuleDatabases) *GetRuleResponseBodyRule {
	s.Databases = v
	return s
}

func (s *GetRuleResponseBodyRule) SetEffectiveEndTime(v string) *GetRuleResponseBodyRule {
	s.EffectiveEndTime = &v
	return s
}

func (s *GetRuleResponseBodyRule) SetEffectiveStartTime(v string) *GetRuleResponseBodyRule {
	s.EffectiveStartTime = &v
	return s
}

func (s *GetRuleResponseBodyRule) SetHostGroups(v []*GetRuleResponseBodyRuleHostGroups) *GetRuleResponseBodyRule {
	s.HostGroups = v
	return s
}

func (s *GetRuleResponseBodyRule) SetHosts(v []*GetRuleResponseBodyRuleHosts) *GetRuleResponseBodyRule {
	s.Hosts = v
	return s
}

func (s *GetRuleResponseBodyRule) SetRuleId(v string) *GetRuleResponseBodyRule {
	s.RuleId = &v
	return s
}

func (s *GetRuleResponseBodyRule) SetRuleName(v string) *GetRuleResponseBodyRule {
	s.RuleName = &v
	return s
}

func (s *GetRuleResponseBodyRule) SetUserGroups(v []*GetRuleResponseBodyRuleUserGroups) *GetRuleResponseBodyRule {
	s.UserGroups = v
	return s
}

func (s *GetRuleResponseBodyRule) SetUsers(v []*GetRuleResponseBodyRuleUsers) *GetRuleResponseBodyRule {
	s.Users = v
	return s
}

func (s *GetRuleResponseBodyRule) Validate() error {
	if s.Databases != nil {
		for _, item := range s.Databases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HostGroups != nil {
		for _, item := range s.HostGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Hosts != nil {
		for _, item := range s.Hosts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserGroups != nil {
		for _, item := range s.UserGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRuleResponseBodyRuleDatabases struct {
	// The database accounts on which permissions are granted by using the authorization rule.
	DatabaseAccounts []*GetRuleResponseBodyRuleDatabasesDatabaseAccounts `json:"DatabaseAccounts,omitempty" xml:"DatabaseAccounts,omitempty" type:"Repeated"`
	// The ID of the database that the policy authorizes users to manage.
	//
	// example:
	//
	// 2
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
}

func (s GetRuleResponseBodyRuleDatabases) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyRuleDatabases) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyRuleDatabases) GetDatabaseAccounts() []*GetRuleResponseBodyRuleDatabasesDatabaseAccounts {
	return s.DatabaseAccounts
}

func (s *GetRuleResponseBodyRuleDatabases) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *GetRuleResponseBodyRuleDatabases) SetDatabaseAccounts(v []*GetRuleResponseBodyRuleDatabasesDatabaseAccounts) *GetRuleResponseBodyRuleDatabases {
	s.DatabaseAccounts = v
	return s
}

func (s *GetRuleResponseBodyRuleDatabases) SetDatabaseId(v string) *GetRuleResponseBodyRuleDatabases {
	s.DatabaseId = &v
	return s
}

func (s *GetRuleResponseBodyRuleDatabases) Validate() error {
	if s.DatabaseAccounts != nil {
		for _, item := range s.DatabaseAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRuleResponseBodyRuleDatabasesDatabaseAccounts struct {
	// The ID of the database account that the policy authorizes users to manage.
	//
	// example:
	//
	// 5
	DatabaseAccountId *string `json:"DatabaseAccountId,omitempty" xml:"DatabaseAccountId,omitempty"`
}

func (s GetRuleResponseBodyRuleDatabasesDatabaseAccounts) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyRuleDatabasesDatabaseAccounts) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyRuleDatabasesDatabaseAccounts) GetDatabaseAccountId() *string {
	return s.DatabaseAccountId
}

func (s *GetRuleResponseBodyRuleDatabasesDatabaseAccounts) SetDatabaseAccountId(v string) *GetRuleResponseBodyRuleDatabasesDatabaseAccounts {
	s.DatabaseAccountId = &v
	return s
}

func (s *GetRuleResponseBodyRuleDatabasesDatabaseAccounts) Validate() error {
	return dara.Validate(s)
}

type GetRuleResponseBodyRuleHostGroups struct {
	// The asset accounts on which permissions are granted by using the authorization rule.
	HostAccountNames []*string `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	// The ID of the asset group that the policy authorizes users to manage.
	//
	// example:
	//
	// 7
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s GetRuleResponseBodyRuleHostGroups) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyRuleHostGroups) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyRuleHostGroups) GetHostAccountNames() []*string {
	return s.HostAccountNames
}

func (s *GetRuleResponseBodyRuleHostGroups) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *GetRuleResponseBodyRuleHostGroups) SetHostAccountNames(v []*string) *GetRuleResponseBodyRuleHostGroups {
	s.HostAccountNames = v
	return s
}

func (s *GetRuleResponseBodyRuleHostGroups) SetHostGroupId(v string) *GetRuleResponseBodyRuleHostGroups {
	s.HostGroupId = &v
	return s
}

func (s *GetRuleResponseBodyRuleHostGroups) Validate() error {
	return dara.Validate(s)
}

type GetRuleResponseBodyRuleHosts struct {
	// The host accounts that the policy authorizes users to manage.
	HostAccounts []*GetRuleResponseBodyRuleHostsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The ID of the host that the policy authorizes users to manage.
	//
	// example:
	//
	// 52
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s GetRuleResponseBodyRuleHosts) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyRuleHosts) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyRuleHosts) GetHostAccounts() []*GetRuleResponseBodyRuleHostsHostAccounts {
	return s.HostAccounts
}

func (s *GetRuleResponseBodyRuleHosts) GetHostId() *string {
	return s.HostId
}

func (s *GetRuleResponseBodyRuleHosts) SetHostAccounts(v []*GetRuleResponseBodyRuleHostsHostAccounts) *GetRuleResponseBodyRuleHosts {
	s.HostAccounts = v
	return s
}

func (s *GetRuleResponseBodyRuleHosts) SetHostId(v string) *GetRuleResponseBodyRuleHosts {
	s.HostId = &v
	return s
}

func (s *GetRuleResponseBodyRuleHosts) Validate() error {
	if s.HostAccounts != nil {
		for _, item := range s.HostAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRuleResponseBodyRuleHostsHostAccounts struct {
	// The ID of the host account that the policy authorizes users to manage.
	//
	// example:
	//
	// 9
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
}

func (s GetRuleResponseBodyRuleHostsHostAccounts) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyRuleHostsHostAccounts) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyRuleHostsHostAccounts) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *GetRuleResponseBodyRuleHostsHostAccounts) SetHostAccountId(v string) *GetRuleResponseBodyRuleHostsHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *GetRuleResponseBodyRuleHostsHostAccounts) Validate() error {
	return dara.Validate(s)
}

type GetRuleResponseBodyRuleUserGroups struct {
	// The ID of the authorized user group.
	//
	// example:
	//
	// 2
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s GetRuleResponseBodyRuleUserGroups) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyRuleUserGroups) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyRuleUserGroups) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *GetRuleResponseBodyRuleUserGroups) SetUserGroupId(v string) *GetRuleResponseBodyRuleUserGroups {
	s.UserGroupId = &v
	return s
}

func (s *GetRuleResponseBodyRuleUserGroups) Validate() error {
	return dara.Validate(s)
}

type GetRuleResponseBodyRuleUsers struct {
	// The ID of the authorized user.
	//
	// example:
	//
	// 3
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetRuleResponseBodyRuleUsers) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyRuleUsers) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyRuleUsers) GetUserId() *string {
	return s.UserId
}

func (s *GetRuleResponseBodyRuleUsers) SetUserId(v string) *GetRuleResponseBodyRuleUsers {
	s.UserId = &v
	return s
}

func (s *GetRuleResponseBodyRuleUsers) Validate() error {
	return dara.Validate(s)
}
