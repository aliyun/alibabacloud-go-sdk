// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateRuleRequest
	GetComment() *string
	SetDatabases(v []*CreateRuleRequestDatabases) *CreateRuleRequest
	GetDatabases() []*CreateRuleRequestDatabases
	SetEffectiveEndTime(v int64) *CreateRuleRequest
	GetEffectiveEndTime() *int64
	SetEffectiveStartTime(v int64) *CreateRuleRequest
	GetEffectiveStartTime() *int64
	SetHostGroups(v []*CreateRuleRequestHostGroups) *CreateRuleRequest
	GetHostGroups() []*CreateRuleRequestHostGroups
	SetHosts(v []*CreateRuleRequestHosts) *CreateRuleRequest
	GetHosts() []*CreateRuleRequestHosts
	SetInstanceId(v string) *CreateRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateRuleRequest
	GetRegionId() *string
	SetRuleName(v string) *CreateRuleRequest
	GetRuleName() *string
	SetUserGroupIds(v []*string) *CreateRuleRequest
	GetUserGroupIds() []*string
	SetUserIds(v []*string) *CreateRuleRequest
	GetUserIds() []*string
}

type CreateRuleRequest struct {
	// The remarks of the authorization rule. The remarks can be up to 500 characters in length.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The information about the database that runs on your server.
	Databases []*CreateRuleRequestDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The end time of the validity period of the authorization rule. Specify a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1672502400
	EffectiveEndTime *int64 `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The start time of the validity period of the authorization rule. Specify a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1669630029
	EffectiveStartTime *int64 `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The information about the asset group that you want to authorize to manage.
	HostGroups []*CreateRuleRequestHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
	// The host information.
	Hosts []*CreateRuleRequestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The bastion host ID.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-5yd2ymfsa0e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the authorization rule. The name can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// An array that consists of user group IDs.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// An array that consists of user IDs.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s CreateRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateRuleRequest) GetDatabases() []*CreateRuleRequestDatabases {
	return s.Databases
}

func (s *CreateRuleRequest) GetEffectiveEndTime() *int64 {
	return s.EffectiveEndTime
}

func (s *CreateRuleRequest) GetEffectiveStartTime() *int64 {
	return s.EffectiveStartTime
}

func (s *CreateRuleRequest) GetHostGroups() []*CreateRuleRequestHostGroups {
	return s.HostGroups
}

func (s *CreateRuleRequest) GetHosts() []*CreateRuleRequestHosts {
	return s.Hosts
}

func (s *CreateRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateRuleRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateRuleRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *CreateRuleRequest) SetComment(v string) *CreateRuleRequest {
	s.Comment = &v
	return s
}

func (s *CreateRuleRequest) SetDatabases(v []*CreateRuleRequestDatabases) *CreateRuleRequest {
	s.Databases = v
	return s
}

func (s *CreateRuleRequest) SetEffectiveEndTime(v int64) *CreateRuleRequest {
	s.EffectiveEndTime = &v
	return s
}

func (s *CreateRuleRequest) SetEffectiveStartTime(v int64) *CreateRuleRequest {
	s.EffectiveStartTime = &v
	return s
}

func (s *CreateRuleRequest) SetHostGroups(v []*CreateRuleRequestHostGroups) *CreateRuleRequest {
	s.HostGroups = v
	return s
}

func (s *CreateRuleRequest) SetHosts(v []*CreateRuleRequestHosts) *CreateRuleRequest {
	s.Hosts = v
	return s
}

func (s *CreateRuleRequest) SetInstanceId(v string) *CreateRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRuleRequest) SetRegionId(v string) *CreateRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRuleRequest) SetRuleName(v string) *CreateRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateRuleRequest) SetUserGroupIds(v []*string) *CreateRuleRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateRuleRequest) SetUserIds(v []*string) *CreateRuleRequest {
	s.UserIds = v
	return s
}

func (s *CreateRuleRequest) Validate() error {
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

type CreateRuleRequestDatabases struct {
	// An array that consists of database account IDs.
	DatabaseAccountIds []*string `json:"DatabaseAccountIds,omitempty" xml:"DatabaseAccountIds,omitempty" type:"Repeated"`
	// The database ID.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "DatabaseId": "1"
	//
	//     }
	//
	// ]
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
}

func (s CreateRuleRequestDatabases) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestDatabases) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestDatabases) GetDatabaseAccountIds() []*string {
	return s.DatabaseAccountIds
}

func (s *CreateRuleRequestDatabases) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *CreateRuleRequestDatabases) SetDatabaseAccountIds(v []*string) *CreateRuleRequestDatabases {
	s.DatabaseAccountIds = v
	return s
}

func (s *CreateRuleRequestDatabases) SetDatabaseId(v string) *CreateRuleRequestDatabases {
	s.DatabaseId = &v
	return s
}

func (s *CreateRuleRequestDatabases) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestHostGroups struct {
	// An array that consists of asset account names.
	HostAccountNames []*string `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	// The asset group ID.
	//
	// example:
	//
	// [{"HostGroupId":"1"}]
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s CreateRuleRequestHostGroups) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestHostGroups) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestHostGroups) GetHostAccountNames() []*string {
	return s.HostAccountNames
}

func (s *CreateRuleRequestHostGroups) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *CreateRuleRequestHostGroups) SetHostAccountNames(v []*string) *CreateRuleRequestHostGroups {
	s.HostAccountNames = v
	return s
}

func (s *CreateRuleRequestHostGroups) SetHostGroupId(v string) *CreateRuleRequestHostGroups {
	s.HostGroupId = &v
	return s
}

func (s *CreateRuleRequestHostGroups) Validate() error {
	return dara.Validate(s)
}

type CreateRuleRequestHosts struct {
	// An array that consists of host account IDs.
	HostAccountIds []*string `json:"HostAccountIds,omitempty" xml:"HostAccountIds,omitempty" type:"Repeated"`
	// The host ID.
	//
	// example:
	//
	// [{"HostId":"1"}]
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s CreateRuleRequestHosts) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequestHosts) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestHosts) GetHostAccountIds() []*string {
	return s.HostAccountIds
}

func (s *CreateRuleRequestHosts) GetHostId() *string {
	return s.HostId
}

func (s *CreateRuleRequestHosts) SetHostAccountIds(v []*string) *CreateRuleRequestHosts {
	s.HostAccountIds = v
	return s
}

func (s *CreateRuleRequestHosts) SetHostId(v string) *CreateRuleRequestHosts {
	s.HostId = &v
	return s
}

func (s *CreateRuleRequestHosts) Validate() error {
	return dara.Validate(s)
}
