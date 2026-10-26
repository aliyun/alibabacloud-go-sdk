// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ModifyRuleRequest
	GetComment() *string
	SetDatabases(v []*ModifyRuleRequestDatabases) *ModifyRuleRequest
	GetDatabases() []*ModifyRuleRequestDatabases
	SetEffectiveEndTime(v int64) *ModifyRuleRequest
	GetEffectiveEndTime() *int64
	SetEffectiveStartTime(v int64) *ModifyRuleRequest
	GetEffectiveStartTime() *int64
	SetHostGroups(v []*ModifyRuleRequestHostGroups) *ModifyRuleRequest
	GetHostGroups() []*ModifyRuleRequestHostGroups
	SetHosts(v []*ModifyRuleRequestHosts) *ModifyRuleRequest
	GetHosts() []*ModifyRuleRequestHosts
	SetInstanceId(v string) *ModifyRuleRequest
	GetInstanceId() *string
	SetProjectId(v int64) *ModifyRuleRequest
	GetProjectId() *int64
	SetRegionId(v string) *ModifyRuleRequest
	GetRegionId() *string
	SetRuleId(v string) *ModifyRuleRequest
	GetRuleId() *string
	SetRuleName(v string) *ModifyRuleRequest
	GetRuleName() *string
	SetUserGroupIds(v []*string) *ModifyRuleRequest
	GetUserGroupIds() []*string
	SetUserIds(v []*string) *ModifyRuleRequest
	GetUserIds() []*string
}

type ModifyRuleRequest struct {
	// The new remarks for the authorization rule. The remarks can be up to 500 characters in length.
	//
	// example:
	//
	// TEST
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The new list of database instances and database accounts associated with the authorization rule.
	Databases []*ModifyRuleRequestDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The new end time of the validity period for the authorization rule. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1672502400
	EffectiveEndTime *int64 `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The new start time of the validity period for the authorization rule. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1669630029
	EffectiveStartTime *int64 `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The new list of asset groups and accounts associated with the authorization rule.
	HostGroups []*ModifyRuleRequestHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
	// The new array of host IDs and host account IDs associated with the authorization rule.
	Hosts []*ModifyRuleRequestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The ID of the Bastionhost instance to which the authorization rule to be updated belongs. > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-zmb2y9ydw08
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance to which the authorization rule to be updated belongs. > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the authorization rule to be updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 68
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The new name of the authorization rule. The name must be 1 to 128 characters in length and cannot start with a special character. It can contain only periods (.), underscores (_), hyphens (-), single quotation marks (\\"), and spaces.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The new list of user group IDs associated with the authorization rule.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The new list of user IDs associated with the authorization rule.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s ModifyRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyRuleRequest) GetComment() *string {
	return s.Comment
}

func (s *ModifyRuleRequest) GetDatabases() []*ModifyRuleRequestDatabases {
	return s.Databases
}

func (s *ModifyRuleRequest) GetEffectiveEndTime() *int64 {
	return s.EffectiveEndTime
}

func (s *ModifyRuleRequest) GetEffectiveStartTime() *int64 {
	return s.EffectiveStartTime
}

func (s *ModifyRuleRequest) GetHostGroups() []*ModifyRuleRequestHostGroups {
	return s.HostGroups
}

func (s *ModifyRuleRequest) GetHosts() []*ModifyRuleRequestHosts {
	return s.Hosts
}

func (s *ModifyRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyRuleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ModifyRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *ModifyRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyRuleRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *ModifyRuleRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *ModifyRuleRequest) SetComment(v string) *ModifyRuleRequest {
	s.Comment = &v
	return s
}

func (s *ModifyRuleRequest) SetDatabases(v []*ModifyRuleRequestDatabases) *ModifyRuleRequest {
	s.Databases = v
	return s
}

func (s *ModifyRuleRequest) SetEffectiveEndTime(v int64) *ModifyRuleRequest {
	s.EffectiveEndTime = &v
	return s
}

func (s *ModifyRuleRequest) SetEffectiveStartTime(v int64) *ModifyRuleRequest {
	s.EffectiveStartTime = &v
	return s
}

func (s *ModifyRuleRequest) SetHostGroups(v []*ModifyRuleRequestHostGroups) *ModifyRuleRequest {
	s.HostGroups = v
	return s
}

func (s *ModifyRuleRequest) SetHosts(v []*ModifyRuleRequestHosts) *ModifyRuleRequest {
	s.Hosts = v
	return s
}

func (s *ModifyRuleRequest) SetInstanceId(v string) *ModifyRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyRuleRequest) SetProjectId(v int64) *ModifyRuleRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyRuleRequest) SetRegionId(v string) *ModifyRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRuleRequest) SetRuleId(v string) *ModifyRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyRuleRequest) SetRuleName(v string) *ModifyRuleRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyRuleRequest) SetUserGroupIds(v []*string) *ModifyRuleRequest {
	s.UserGroupIds = v
	return s
}

func (s *ModifyRuleRequest) SetUserIds(v []*string) *ModifyRuleRequest {
	s.UserIds = v
	return s
}

func (s *ModifyRuleRequest) Validate() error {
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
	return nil
}

type ModifyRuleRequestDatabases struct {
	// The list of database account IDs.
	DatabaseAccountIds []*string `json:"DatabaseAccountIds,omitempty" xml:"DatabaseAccountIds,omitempty" type:"Repeated"`
	// The database instance ID.
	//
	// example:
	//
	// 2
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
}

func (s ModifyRuleRequestDatabases) String() string {
	return dara.Prettify(s)
}

func (s ModifyRuleRequestDatabases) GoString() string {
	return s.String()
}

func (s *ModifyRuleRequestDatabases) GetDatabaseAccountIds() []*string {
	return s.DatabaseAccountIds
}

func (s *ModifyRuleRequestDatabases) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ModifyRuleRequestDatabases) SetDatabaseAccountIds(v []*string) *ModifyRuleRequestDatabases {
	s.DatabaseAccountIds = v
	return s
}

func (s *ModifyRuleRequestDatabases) SetDatabaseId(v string) *ModifyRuleRequestDatabases {
	s.DatabaseId = &v
	return s
}

func (s *ModifyRuleRequestDatabases) Validate() error {
	return dara.Validate(s)
}

type ModifyRuleRequestHostGroups struct {
	// The list of asset account names.
	HostAccountNames []*string `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	// The asset group ID.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s ModifyRuleRequestHostGroups) String() string {
	return dara.Prettify(s)
}

func (s ModifyRuleRequestHostGroups) GoString() string {
	return s.String()
}

func (s *ModifyRuleRequestHostGroups) GetHostAccountNames() []*string {
	return s.HostAccountNames
}

func (s *ModifyRuleRequestHostGroups) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *ModifyRuleRequestHostGroups) SetHostAccountNames(v []*string) *ModifyRuleRequestHostGroups {
	s.HostAccountNames = v
	return s
}

func (s *ModifyRuleRequestHostGroups) SetHostGroupId(v string) *ModifyRuleRequestHostGroups {
	s.HostGroupId = &v
	return s
}

func (s *ModifyRuleRequestHostGroups) Validate() error {
	return dara.Validate(s)
}

type ModifyRuleRequestHosts struct {
	// The list of host account IDs.
	HostAccountIds []*string `json:"HostAccountIds,omitempty" xml:"HostAccountIds,omitempty" type:"Repeated"`
	// The host ID.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s ModifyRuleRequestHosts) String() string {
	return dara.Prettify(s)
}

func (s ModifyRuleRequestHosts) GoString() string {
	return s.String()
}

func (s *ModifyRuleRequestHosts) GetHostAccountIds() []*string {
	return s.HostAccountIds
}

func (s *ModifyRuleRequestHosts) GetHostId() *string {
	return s.HostId
}

func (s *ModifyRuleRequestHosts) SetHostAccountIds(v []*string) *ModifyRuleRequestHosts {
	s.HostAccountIds = v
	return s
}

func (s *ModifyRuleRequestHosts) SetHostId(v string) *ModifyRuleRequestHosts {
	s.HostId = &v
	return s
}

func (s *ModifyRuleRequestHosts) Validate() error {
	return dara.Validate(s)
}
