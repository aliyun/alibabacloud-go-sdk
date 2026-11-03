// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAclCheckDetailStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckRecord(v *UpdateAclCheckDetailStatusResponseBodyCheckRecord) *UpdateAclCheckDetailStatusResponseBody
	GetCheckRecord() *UpdateAclCheckDetailStatusResponseBodyCheckRecord
	SetRequestId(v string) *UpdateAclCheckDetailStatusResponseBody
	GetRequestId() *string
}

type UpdateAclCheckDetailStatusResponseBody struct {
	// The ACL check record.
	CheckRecord *UpdateAclCheckDetailStatusResponseBodyCheckRecord `json:"CheckRecord,omitempty" xml:"CheckRecord,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 850A84D6************00090125EEB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAclCheckDetailStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclCheckDetailStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAclCheckDetailStatusResponseBody) GetCheckRecord() *UpdateAclCheckDetailStatusResponseBodyCheckRecord {
	return s.CheckRecord
}

func (s *UpdateAclCheckDetailStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAclCheckDetailStatusResponseBody) SetCheckRecord(v *UpdateAclCheckDetailStatusResponseBodyCheckRecord) *UpdateAclCheckDetailStatusResponseBody {
	s.CheckRecord = v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBody) SetRequestId(v string) *UpdateAclCheckDetailStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBody) Validate() error {
	if s.CheckRecord != nil {
		if err := s.CheckRecord.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAclCheckDetailStatusResponseBodyCheckRecord struct {
	// The list of ACL check results.
	Acls []*UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Repeated"`
	// The name of the ACL check. Valid values:
	//
	// - **PolicyHitCountZero**: policies with no traffic hits
	//
	// - **PolicySourceDestinationSame**: Invalid policy with the same source and destination
	//
	// - **PolicyDuplicate**: Duplicate or redundant policy
	//
	// - **PolicyConflict**: Policy that conflicts with business
	//
	// - **DefaultPolicyNotDeny**: default policy that does not use the Deny All whitelist mechanism
	//
	// - **PolicyPortHighRisk**: policies that allow traffic on high-risk ports
	//
	// - **PolicyTooLoose**: overly loose control policy
	//
	// - **AddressBookIpSeparated**: duplicate, overlapping, or scattered IP address books
	//
	// - **AddressBookPortSeparated**: duplicate, overlapping, or scattered port address books
	//
	// - **AddressBookDomainValid**: validity check for domain name address books
	//
	// example:
	//
	// PolicyHitCountZero
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// dwd_mysql_lingwan_faxing_chat_config_di
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timestamp of the last check. The value is a UNIX timestamp measured in seconds.
	//
	// example:
	//
	// 1724982259
	LastCheckTime *string `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
	// The risk level.
	//
	// example:
	//
	// High
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The total number of policies.
	//
	// example:
	//
	// 1
	PolicyTotalCount *int64 `json:"PolicyTotalCount,omitempty" xml:"PolicyTotalCount,omitempty"`
	// The assessment details of the ACL check record.
	//
	// example:
	//
	// Due to business offline or other reasons, the number of hits of the object policy in a period of time is 0.
	RecordAssessmentDetail *string `json:"RecordAssessmentDetail,omitempty" xml:"RecordAssessmentDetail,omitempty"`
	// The ID of the ACL check task.
	//
	// example:
	//
	// task-c92d4544ef7b6a42
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateAclCheckDetailStatusResponseBodyCheckRecord) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclCheckDetailStatusResponseBodyCheckRecord) GoString() string {
	return s.String()
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) GetAcls() []*UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls {
	return s.Acls
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) GetCheckName() *string {
	return s.CheckName
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) GetDescription() *string {
	return s.Description
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) GetLastCheckTime() *string {
	return s.LastCheckTime
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) GetLevel() *string {
	return s.Level
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) GetPolicyTotalCount() *int64 {
	return s.PolicyTotalCount
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) GetRecordAssessmentDetail() *string {
	return s.RecordAssessmentDetail
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) SetAcls(v []*UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls) *UpdateAclCheckDetailStatusResponseBodyCheckRecord {
	s.Acls = v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) SetCheckName(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecord {
	s.CheckName = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) SetDescription(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecord {
	s.Description = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) SetLastCheckTime(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecord {
	s.LastCheckTime = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) SetLevel(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecord {
	s.Level = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) SetPolicyTotalCount(v int64) *UpdateAclCheckDetailStatusResponseBodyCheckRecord {
	s.PolicyTotalCount = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) SetRecordAssessmentDetail(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecord {
	s.RecordAssessmentDetail = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) SetTaskId(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecord {
	s.TaskId = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecord) Validate() error {
	if s.Acls != nil {
		for _, item := range s.Acls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls struct {
	// The ACL check result.
	Acl *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl `json:"Acl,omitempty" xml:"Acl,omitempty" type:"Struct"`
	// The assessment details of the ACL policy.
	//
	// example:
	//
	// No traffic hit policy.
	AclAssessmentDetail *string `json:"AclAssessmentDetail,omitempty" xml:"AclAssessmentDetail,omitempty"`
	// The ACL check status.
	//
	// example:
	//
	// configuring
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
}

func (s UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls) GoString() string {
	return s.String()
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls) GetAcl() *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	return s.Acl
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls) GetAclAssessmentDetail() *string {
	return s.AclAssessmentDetail
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls) GetAclStatus() *string {
	return s.AclStatus
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls) SetAcl(v *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls {
	s.Acl = v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls) SetAclAssessmentDetail(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls {
	s.AclAssessmentDetail = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls) SetAclStatus(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls {
	s.AclStatus = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls) Validate() error {
	if s.Acl != nil {
		if err := s.Acl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl struct {
	// The action that Cloud Firewall performs on the traffic in the access control policy. Valid values:
	//
	// - **accept**: allows the traffic.
	//
	// - **drop**: blocks the traffic.
	//
	// - **log**: monitors the traffic.
	//
	// example:
	//
	// log
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy.
	//
	// example:
	//
	// 1e8ed1b2-cebc-4b95-a9cc-0cb7ce2c0c2c
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The list of addresses in the address book.
	AddressList []*string `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
	// The number of addresses in the address book.
	//
	// example:
	//
	// 1
	AddressListCount *int32 `json:"AddressListCount,omitempty" xml:"AddressListCount,omitempty"`
	// The ID of the application for the traffic in the access control policy.
	//
	// example:
	//
	// HTTP
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type supported by the access control policy. Valid values:
	//
	// - **FTP**
	//
	// - **HTTP**
	//
	// - **HTTPS**
	//
	// - **Memcache**
	//
	// - **MongoDB**
	//
	// - **MQTT**
	//
	// - **MySQL**
	//
	// - **RDP**
	//
	// - **Redis**
	//
	// - **SMTP**
	//
	// - **SMTPS**
	//
	// - **SSH**
	//
	// - **SSL_No_Cert**
	//
	// - **SSL**
	//
	// - **VNC**
	//
	// > The valid values of the application type depend on the protocol type specified by Proto. If Proto is set to TCP, ApplicationNameList supports all the preceding application types. If both ApplicationNameList and ApplicationName are specified, ApplicationNameList takes precedence.
	//
	// example:
	//
	// ANY
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The list of application types supported by the access control policy.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// Specifies whether to automatically add the public IP addresses of newly tagged ECS instances to the address book. Newly tagged ECS instances include newly purchased ECS instances with tags configured and ECS instances with modified tags.
	//
	// example:
	//
	// 0
	AutoAddTagEcs *int32 `json:"AutoAddTagEcs,omitempty" xml:"AutoAddTagEcs,omitempty"`
	// The time when the policy was created. The value is a UNIX timestamp measured in seconds.
	//
	// example:
	//
	// 1761062400
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the access control policy.
	//
	// example:
	//
	// test_policy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port of the traffic in the access control policy.
	//
	// example:
	//
	// 22/22
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// - **port**: port
	//
	// - **group**: port address book
	//
	// example:
	//
	// port
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The list of ports in the destination port address book.
	DestPortGroupPorts []*string `json:"DestPortGroupPorts,omitempty" xml:"DestPortGroupPorts,omitempty" type:"Repeated"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// - **port**: port
	//
	// - **group**: port address book
	//
	// example:
	//
	// port
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy. Fuzzy match is supported. The value of this parameter varies based on the destination type (DestinationType).
	//
	// - If DestinationType is set to `net`, the value is a CIDR block. Example: 10.0.3.0/24.
	//
	// - If DestinationType is set to `domain`, the value is a domain name. Example: aliyun.
	//
	// - If DestinationType is set to `group`, the value is the name of an address book. Example: db_group.
	//
	// - If DestinationType is set to `location`, the value is an area name. For more information about the area positional encoding, see AddControlPolicy. Example: ["BJ11", "ZB"].
	//
	// > If the Settings for this parameter are not specified, all types of destination addresses are queried.
	//
	// example:
	//
	// kms.cn-shanghai.aliyuncs.com
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The list of CIDR blocks in the destination address book of the access control policy.
	DestinationGroupCidrs []*string `json:"DestinationGroupCidrs,omitempty" xml:"DestinationGroupCidrs,omitempty" type:"Repeated"`
	// The type of the destination address book in the access control policy. Valid values:
	//
	// - **ip**: IP address book that contains one or more CIDR blocks
	//
	// - **tag**: ECS tag-based address book that contains the IP addresses of one or more ECS tags
	//
	// - **domain**: domain name address book that contains one or more domain names
	//
	// - **threat**: threat intelligence address book that contains one or more malicious IP addresses or domain names
	//
	// - **backsrc**: back-to-origin address book that contains the back-to-origin addresses of one or more Anti-DDoS or WAF instances
	//
	// example:
	//
	// domain
	DestinationGroupType *string `json:"DestinationGroupType,omitempty" xml:"DestinationGroupType,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// - **net**: destination CIDR block
	//
	// - **group**: destination address book
	//
	// - **domain**: destination domain name
	//
	// example:
	//
	// group
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The traffic direction controlled by the access control policy. Valid values:
	//
	// - **in**: inbound traffic
	//
	// - **out**: outbound traffic
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The DNS resolution result.
	//
	// example:
	//
	// 192.168.0.1/32
	DnsResult *string `json:"DnsResult,omitempty" xml:"DnsResult,omitempty"`
	// The timestamp of the DNS resolution. The value is a UNIX timestamp measured in seconds.
	//
	// example:
	//
	// 1579261141
	DnsResultTime *int64 `json:"DnsResultTime,omitempty" xml:"DnsResultTime,omitempty"`
	// The domain name resolution mode of the access control policy. Valid values:
	//
	// - **FQDN**: FQDN-based resolution
	//
	// - **DNS**: DNS-based dynamic resolution
	//
	// - **FQDN_AND_DNS**: FQDN- and DNS-based dynamic resolution
	//
	// example:
	//
	// 0
	DomainResolveType *int32 `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The end time of the query. The value is a UNIX timestamp measured in seconds.
	//
	// example:
	//
	// 1752754426
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the address book.
	//
	// example:
	//
	// subscribe
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the address book. Valid values:
	//
	// - **ip**: IP address book
	//
	// - **domain**: domain name address book
	//
	// - **port**: port address book
	//
	// - **tag**: ECS tag-based address book
	//
	// - **allCloud**: cloud service address book
	//
	// - **threat**: threat intelligence address book
	//
	// example:
	//
	// port
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The unique ID of the address book.
	//
	// > Value source: [List address books](~~DescribeAddressBook~~).
	//
	// example:
	//
	// 5a96a798-9b73-47f7-831e-1d7aa3c987a9
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The timestamp of the latest time of hits. The value is in the UNIX timestamp format measured in seconds.
	//
	// example:
	//
	// 1579261141
	HitLastTime *int64 `json:"HitLastTime,omitempty" xml:"HitLastTime,omitempty"`
	// The number of hits of the access control policy.
	//
	// example:
	//
	// 1
	HitTimes *int64 `json:"HitTimes,omitempty" xml:"HitTimes,omitempty"`
	// The IP version of the asset protected by Cloud Firewall. Valid values:
	//
	// - **4**: IPv4 (default).
	//
	// - **6**: IPv6.
	//
	// example:
	//
	// 4
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The time when the policy was last modified. The value is a UNIX timestamp measured in seconds.
	//
	// example:
	//
	// 1761062400
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the NAT gateway.
	//
	// example:
	//
	// ngw-gw85zno51npz7lgc04z89
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The priority of the access control policy.
	//
	// The priority starts from 1 and increases in ascending order. A smaller priority number indicates a higher priority.
	//
	// example:
	//
	// 1
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The protocol type of the traffic in the access control policy. Valid values:
	//
	// - **TCP**
	//
	// - **UDP**
	//
	// - **ICMP**
	//
	// - **ANY**: all protocol types
	//
	// > If this parameter is not specified, all protocol types are queried.
	//
	// example:
	//
	// ANY
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The number of times the address book is referenced.
	//
	// example:
	//
	// 1
	ReferenceCount *int32 `json:"ReferenceCount,omitempty" xml:"ReferenceCount,omitempty"`
	// The status of the access control policy. By default, a policy is enabled after it is created. Valid values:
	//
	// - **true**: The policy is enabled.
	//
	// - **false**: The policy is disabled.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The set of recurrence days for the Policy Validity Period of the access control policy.
	//
	// - If RepeatType is set to `Permanent`, `None`, or `Daily`, RepeatDays is an empty set. Example: [].
	//
	// - If RepeatType is set to `Weekly`, RepeatDays cannot be empty. Example: [0, 6].
	//
	// > If the Settings for RepeatType are `Weekly`, the values in RepeatDays cannot be duplicate.
	//
	// - If RepeatType is set to `Monthly`, RepeatDays cannot be empty. Example: [1, 31].
	//
	// > If the Settings for RepeatType are `Monthly`, the values in RepeatDays cannot be duplicate.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The recurrence end time of the policy validity period. The time must be in the 24-hour HH:mm format, such as 23:30. It must be on the hour or half hour, and at least 30 minutes later than the recurrence start time.
	//
	// > If RepeatType is set to `Permanent` or `None`, RepeatEndTime is empty. If RepeatType is set to `Daily`, `Weekly`, or `Monthly`, RepeatEndTime must be specified.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The recurrence start time of the policy validity period. The time must be in the 24-hour HH:mm format, such as 08:00. It must be on the hour or half hour, and at least 30 minutes earlier than the recurrence end time.
	//
	// > If RepeatType is set to `Permanent` or `None`, RepeatStartTime is empty. If RepeatType is set to `Daily`, `Weekly`, or `Monthly`, RepeatStartTime must be specified.
	//
	// example:
	//
	// 08:00
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type of the Policy Validity Period for the access control policy. Valid values:
	//
	// - **Permanent**: always (default)
	//
	// - **None**: specified single time
	//
	// - **Daily**: every day
	//
	// - **Weekly**: every week
	//
	// - **Monthly**: every month
	//
	// example:
	//
	// None
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Valid values:
	//
	// - If **SourceType*	- is set to `net`, the value is a source CIDR block. Example: 10.2.4.0/24.
	//
	// - If **SourceType*	- is set to `group`, the value is the name of a source address book. Example: db_group.
	//
	// example:
	//
	// 10.71.94.24
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The list of CIDR blocks in the source address book of the access control policy.
	SourceGroupCidrs []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	// The type of the source address book in the access control policy. Valid values:
	//
	// - **ip**: IP address book that contains one or more CIDR blocks
	//
	// - **tag**: ECS tag-based address book that contains the IP addresses of one or more ECS tags
	//
	// - **domain**: domain name address book that contains one or more domain names
	//
	// - **threat**: threat intelligence address book that contains one or more malicious IP addresses or domain names
	//
	// - **backsrc**: back-to-origin address book that contains the back-to-origin addresses of one or more Anti-DDoS or WAF instances
	//
	// example:
	//
	// ip
	SourceGroupType *string `json:"SourceGroupType,omitempty" xml:"SourceGroupType,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// - **net**: source CIDR block
	//
	// - **group**: source address book
	//
	// - **location**: source region
	//
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The number of access control policy quotas consumed. This value is the sum of the quotas consumed by each policy.
	//
	// The number of quotas consumed by a single policy = Number of source CIDR blocks × Number of destination addresses (number of CIDR blocks, regions, or domain names) × Number of applications × Number of port ranges.
	//
	// example:
	//
	// 10
	SpreadCnt *int32 `json:"SpreadCnt,omitempty" xml:"SpreadCnt,omitempty"`
	// The start time of the query. The value is a UNIX timestamp measured in seconds.
	//
	// example:
	//
	// 1736130347
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The list of ECS tags.
	TagList []*UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAclTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The relationship between multiple ECS tags.
	//
	// example:
	//
	// and
	TagRelation *string `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
	// The instance ID of the border virtual private cloud (VPC) firewalls.
	//
	// example:
	//
	// cen-cw4z051hr8x53qniv5
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GoString() string {
	return s.String()
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetAclAction() *string {
	return s.AclAction
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetAclUuid() *string {
	return s.AclUuid
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetAddressList() []*string {
	return s.AddressList
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetAddressListCount() *int32 {
	return s.AddressListCount
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetApplicationNameList() []*string {
	return s.ApplicationNameList
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetAutoAddTagEcs() *int32 {
	return s.AutoAddTagEcs
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetDescription() *string {
	return s.Description
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetDestPort() *string {
	return s.DestPort
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetDestPortGroup() *string {
	return s.DestPortGroup
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetDestPortGroupPorts() []*string {
	return s.DestPortGroupPorts
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetDestPortType() *string {
	return s.DestPortType
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetDestination() *string {
	return s.Destination
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetDestinationGroupCidrs() []*string {
	return s.DestinationGroupCidrs
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetDestinationGroupType() *string {
	return s.DestinationGroupType
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetDestinationType() *string {
	return s.DestinationType
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetDirection() *string {
	return s.Direction
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetDnsResult() *string {
	return s.DnsResult
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetDnsResultTime() *int64 {
	return s.DnsResultTime
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetDomainResolveType() *int32 {
	return s.DomainResolveType
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetEndTime() *int64 {
	return s.EndTime
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetGroupType() *string {
	return s.GroupType
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetGroupUuid() *string {
	return s.GroupUuid
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetHitLastTime() *int64 {
	return s.HitLastTime
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetHitTimes() *int64 {
	return s.HitTimes
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetIpVersion() *int32 {
	return s.IpVersion
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetOrder() *int32 {
	return s.Order
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetProto() *string {
	return s.Proto
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetReferenceCount() *int32 {
	return s.ReferenceCount
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetRelease() *string {
	return s.Release
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetRepeatDays() []*int64 {
	return s.RepeatDays
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetRepeatEndTime() *string {
	return s.RepeatEndTime
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetRepeatStartTime() *string {
	return s.RepeatStartTime
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetRepeatType() *string {
	return s.RepeatType
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetSource() *string {
	return s.Source
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetSourceGroupCidrs() []*string {
	return s.SourceGroupCidrs
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetSourceGroupType() *string {
	return s.SourceGroupType
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetSpreadCnt() *int32 {
	return s.SpreadCnt
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetStartTime() *int64 {
	return s.StartTime
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetTagList() []*UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAclTagList {
	return s.TagList
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetTagRelation() *string {
	return s.TagRelation
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetAclAction(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.AclAction = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetAclUuid(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.AclUuid = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetAddressList(v []*string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.AddressList = v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetAddressListCount(v int32) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.AddressListCount = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetApplicationId(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.ApplicationId = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetApplicationName(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.ApplicationName = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetApplicationNameList(v []*string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.ApplicationNameList = v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetAutoAddTagEcs(v int32) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.AutoAddTagEcs = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetCreateTime(v int64) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.CreateTime = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetDescription(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.Description = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetDestPort(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.DestPort = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetDestPortGroup(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.DestPortGroup = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetDestPortGroupPorts(v []*string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.DestPortGroupPorts = v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetDestPortType(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.DestPortType = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetDestination(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.Destination = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetDestinationGroupCidrs(v []*string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.DestinationGroupCidrs = v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetDestinationGroupType(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.DestinationGroupType = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetDestinationType(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.DestinationType = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetDirection(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.Direction = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetDnsResult(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.DnsResult = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetDnsResultTime(v int64) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.DnsResultTime = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetDomainResolveType(v int32) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.DomainResolveType = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetEndTime(v int64) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.EndTime = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetGroupName(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.GroupName = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetGroupType(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.GroupType = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetGroupUuid(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.GroupUuid = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetHitLastTime(v int64) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.HitLastTime = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetHitTimes(v int64) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.HitTimes = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetIpVersion(v int32) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.IpVersion = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetModifyTime(v int64) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.ModifyTime = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetNatGatewayId(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.NatGatewayId = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetOrder(v int32) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.Order = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetProto(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.Proto = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetReferenceCount(v int32) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.ReferenceCount = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetRelease(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.Release = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetRepeatDays(v []*int64) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.RepeatDays = v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetRepeatEndTime(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.RepeatEndTime = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetRepeatStartTime(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.RepeatStartTime = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetRepeatType(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.RepeatType = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetSource(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.Source = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetSourceGroupCidrs(v []*string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.SourceGroupCidrs = v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetSourceGroupType(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.SourceGroupType = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetSourceType(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.SourceType = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetSpreadCnt(v int32) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.SpreadCnt = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetStartTime(v int64) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.StartTime = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetTagList(v []*UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAclTagList) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.TagList = v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetTagRelation(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.TagRelation = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) SetVpcFirewallId(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl {
	s.VpcFirewallId = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl) Validate() error {
	if s.TagList != nil {
		for _, item := range s.TagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAclTagList struct {
	// The ECS tag key.
	//
	// example:
	//
	// produce
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The ECS tag key.
	//
	// example:
	//
	// tfTestAcc0
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAclTagList) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAclTagList) GoString() string {
	return s.String()
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAclTagList) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAclTagList) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAclTagList) SetTagKey(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAclTagList {
	s.TagKey = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAclTagList) SetTagValue(v string) *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAclTagList {
	s.TagValue = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAclTagList) Validate() error {
	return dara.Validate(s)
}
