// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckRecord(v *DescribeAclCheckResponseBodyCheckRecord) *DescribeAclCheckResponseBody
	GetCheckRecord() *DescribeAclCheckResponseBodyCheckRecord
	SetRequestId(v string) *DescribeAclCheckResponseBody
	GetRequestId() *string
}

type DescribeAclCheckResponseBody struct {
	// The check record.
	CheckRecord *DescribeAclCheckResponseBodyCheckRecord `json:"CheckRecord,omitempty" xml:"CheckRecord,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 25E655B0-CAED-53D4-8054-F983126****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAclCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclCheckResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAclCheckResponseBody) GetCheckRecord() *DescribeAclCheckResponseBodyCheckRecord {
	return s.CheckRecord
}

func (s *DescribeAclCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAclCheckResponseBody) SetCheckRecord(v *DescribeAclCheckResponseBodyCheckRecord) *DescribeAclCheckResponseBody {
	s.CheckRecord = v
	return s
}

func (s *DescribeAclCheckResponseBody) SetRequestId(v string) *DescribeAclCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAclCheckResponseBody) Validate() error {
	if s.CheckRecord != nil {
		if err := s.CheckRecord.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAclCheckResponseBodyCheckRecord struct {
	// The total number of access control list (ACL) policies at the time of the check.
	//
	// example:
	//
	// 10
	AclTotalCount *int64 `json:"AclTotalCount,omitempty" xml:"AclTotalCount,omitempty"`
	// The list of ACL check results.
	Acls []*DescribeAclCheckResponseBodyCheckRecordAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Repeated"`
	// The name of the ACL check.
	//
	// example:
	//
	// PolicyHitCountZero
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// The description of the ACL check item.
	//
	// example:
	//
	// Due to business offline or other reasons, the number of hits of the object policy in a period of time is 0.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timestamp of the last check. Unit: seconds.
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
	// The assessment details of the ACL check.
	//
	// example:
	//
	// It is recommended to remove the invalid policy, while helping to save the specification.
	RecordAssessmentDetail *string `json:"RecordAssessmentDetail,omitempty" xml:"RecordAssessmentDetail,omitempty"`
	// The task ID.
	//
	// example:
	//
	// task-c92d4544ef7b6a42
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeAclCheckResponseBodyCheckRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclCheckResponseBodyCheckRecord) GoString() string {
	return s.String()
}

func (s *DescribeAclCheckResponseBodyCheckRecord) GetAclTotalCount() *int64 {
	return s.AclTotalCount
}

func (s *DescribeAclCheckResponseBodyCheckRecord) GetAcls() []*DescribeAclCheckResponseBodyCheckRecordAcls {
	return s.Acls
}

func (s *DescribeAclCheckResponseBodyCheckRecord) GetCheckName() *string {
	return s.CheckName
}

func (s *DescribeAclCheckResponseBodyCheckRecord) GetDescription() *string {
	return s.Description
}

func (s *DescribeAclCheckResponseBodyCheckRecord) GetLastCheckTime() *string {
	return s.LastCheckTime
}

func (s *DescribeAclCheckResponseBodyCheckRecord) GetLevel() *string {
	return s.Level
}

func (s *DescribeAclCheckResponseBodyCheckRecord) GetRecordAssessmentDetail() *string {
	return s.RecordAssessmentDetail
}

func (s *DescribeAclCheckResponseBodyCheckRecord) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeAclCheckResponseBodyCheckRecord) SetAclTotalCount(v int64) *DescribeAclCheckResponseBodyCheckRecord {
	s.AclTotalCount = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecord) SetAcls(v []*DescribeAclCheckResponseBodyCheckRecordAcls) *DescribeAclCheckResponseBodyCheckRecord {
	s.Acls = v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecord) SetCheckName(v string) *DescribeAclCheckResponseBodyCheckRecord {
	s.CheckName = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecord) SetDescription(v string) *DescribeAclCheckResponseBodyCheckRecord {
	s.Description = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecord) SetLastCheckTime(v string) *DescribeAclCheckResponseBodyCheckRecord {
	s.LastCheckTime = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecord) SetLevel(v string) *DescribeAclCheckResponseBodyCheckRecord {
	s.Level = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecord) SetRecordAssessmentDetail(v string) *DescribeAclCheckResponseBodyCheckRecord {
	s.RecordAssessmentDetail = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecord) SetTaskId(v string) *DescribeAclCheckResponseBodyCheckRecord {
	s.TaskId = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecord) Validate() error {
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

type DescribeAclCheckResponseBodyCheckRecordAcls struct {
	// The ACL check result.
	Acl *DescribeAclCheckResponseBodyCheckRecordAclsAcl `json:"Acl,omitempty" xml:"Acl,omitempty" type:"Struct"`
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
	// Pending
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
}

func (s DescribeAclCheckResponseBodyCheckRecordAcls) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclCheckResponseBodyCheckRecordAcls) GoString() string {
	return s.String()
}

func (s *DescribeAclCheckResponseBodyCheckRecordAcls) GetAcl() *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	return s.Acl
}

func (s *DescribeAclCheckResponseBodyCheckRecordAcls) GetAclAssessmentDetail() *string {
	return s.AclAssessmentDetail
}

func (s *DescribeAclCheckResponseBodyCheckRecordAcls) GetAclStatus() *string {
	return s.AclStatus
}

func (s *DescribeAclCheckResponseBodyCheckRecordAcls) SetAcl(v *DescribeAclCheckResponseBodyCheckRecordAclsAcl) *DescribeAclCheckResponseBodyCheckRecordAcls {
	s.Acl = v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAcls) SetAclAssessmentDetail(v string) *DescribeAclCheckResponseBodyCheckRecordAcls {
	s.AclAssessmentDetail = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAcls) SetAclStatus(v string) *DescribeAclCheckResponseBodyCheckRecordAcls {
	s.AclStatus = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAcls) Validate() error {
	if s.Acl != nil {
		if err := s.Acl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAclCheckResponseBodyCheckRecordAclsAcl struct {
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
	// 997b38e0-01fa-4db7-8d30-02ebf6fdb747
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The address information in the address book.
	AddressList []*string `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
	// The number of addresses in the address book.
	//
	// example:
	//
	// 1
	AddressListCount *int32 `json:"AddressListCount,omitempty" xml:"AddressListCount,omitempty"`
	// The list of addresses and description notes.
	Addresses []*DescribeAclCheckResponseBodyCheckRecordAclsAclAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// The ID of the application for the traffic in the access control policy.
	//
	// example:
	//
	// plugin_idp4_ciam
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application types supported by the access control policies of VPC border firewalls. Use ApplicationNameList. Valid values:
	//
	// - **HTTP**
	//
	// - **HTTPS**
	//
	// - **MySQL**
	//
	// - **SMTP**
	//
	// - **SMTPS**
	//
	// - **RDP**
	//
	// - **VNC**
	//
	// - **SSH**
	//
	// - **Redis**
	//
	// - **MQTT**
	//
	// - **MongoDB**
	//
	// - **Memcache**
	//
	// - **SSL**
	//
	// - **ANY**: all application types.
	//
	// example:
	//
	// ANY
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The list of application types supported by the access control policy. Valid values:
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
	// - **SSL**
	//
	// - **VNC**
	//
	// - **ANY**: all application types.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// Specifies whether to automatically add the public IP addresses of newly matched ECS instances to the address book. Newly matched ECS instances refer to newly purchased ECS instances with tags configured or ECS instances with modified tags.
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
	// 80/80
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The type of the destination port for traffic access in the access control policy. Valid values:
	//
	// - **port**: port.
	//
	// - **group**: port address book.
	//
	// example:
	//
	// my_port_group
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The ports in the destination port address book.
	DestPortGroupPorts []*string `json:"DestPortGroupPorts,omitempty" xml:"DestPortGroupPorts,omitempty" type:"Repeated"`
	// The type of the destination port for traffic access in the access control policy. Valid values:
	//
	// - **port**: port.
	//
	// - **group**: port address book.
	//
	// example:
	//
	// port
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy. Fuzzy match is supported. The value varies based on the destination type (DestinationType).
	//
	// - If DestinationType is set to `net`, the destination address is a CIDR block. Example: 10.0.3.0/24.
	//
	// - If DestinationType is set to `domain`, the destination address is a domain name. Example: aliyun.
	//
	// - If DestinationType is set to `group`, the destination address is the name of an address book. Example: db_group.
	//
	// - If DestinationType is set to `location`, the destination address is an area name. For more information about area positional encoding, see AddControlPolicy. Example: ["BJ11", "ZB"].
	//
	// > If you do not specify this parameter, destination addresses of all types are queried.
	//
	// example:
	//
	// kms.cn-shanghai.aliyuncs.com
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The CIDR blocks in the destination address book of the access control policy.
	DestinationGroupCidrs []*string `json:"DestinationGroupCidrs,omitempty" xml:"DestinationGroupCidrs,omitempty" type:"Repeated"`
	// The type of the destination address book in the access control policy. Valid values:
	//
	// - **ip**: IP address book, which contains one or more CIDR blocks.
	//
	// - **tag**: ECS tag-based address book, which contains the IP addresses of one or more ECS tags.
	//
	// - **domain**: domain name address book, which contains one or more domain names.
	//
	// - **threat**: threat intelligence address book, which contains one or more malicious IP addresses or domain names.
	//
	// - **backsrc**: back-to-origin address book, which contains the back-to-origin IP addresses of one or more Anti-DDoS or WAF instances.
	//
	// example:
	//
	// domain
	DestinationGroupType *string `json:"DestinationGroupType,omitempty" xml:"DestinationGroupType,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// - **net**: destination CIDR block.
	//
	// - **group**: destination address book.
	//
	// - **domain**: destination domain name.
	//
	// - **location**: destination location.
	//
	// example:
	//
	// domain
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The traffic direction on the Internet. Valid values:
	//
	// - **in**: access control for inbound traffic.
	//
	// - **out**: access control for outbound traffic.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The DNS resolution result.
	//
	// example:
	//
	// 192.0.XX.XX
	DnsResult *string `json:"DnsResult,omitempty" xml:"DnsResult,omitempty"`
	// The timestamp of the DNS resolution. The value is a UNIX timestamp measured in seconds.
	//
	// example:
	//
	// 1579261141
	DnsResultTime *int64 `json:"DnsResultTime,omitempty" xml:"DnsResultTime,omitempty"`
	// The domain name resolution method of the access control policy. Valid values:
	//
	// 	- **FQDN**: FQDN-based resolution.
	//
	// 	- **DNS**: DNS-based dynamic resolution.
	//
	// 	- **FQDN_AND_DNS**: FQDN- and DNS-based dynamic resolution.
	//
	// example:
	//
	// FQDN
	DomainResolveType *int32 `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The end time of the policy validity period for the access control policy. The value is a UNIX timestamp measured in seconds. The time must be on the hour or half-hour, and must be at least 30 minutes later than the start time.
	//
	// > If RepeatType is set to Permanent, EndTime is empty. If RepeatType is set to None, Daily, Weekly, or Monthly, EndTime must be specified.
	//
	// example:
	//
	// 1758334822
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the address book.
	//
	// example:
	//
	// Zhong Kui Open White List
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the address book. Valid values:
	//
	// - **ip**: IP address book.
	//
	// - **domain**: domain name address book.
	//
	// - **port**: port address book.
	//
	// - **tag**: ECS tag-based address book.
	//
	// - **allCloud**: cloud service address book.
	//
	// - **threat**: threat intelligence address book.
	//
	// example:
	//
	// ip
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The unique ID of the address book.
	//
	// To delete an address book for access control, you must provide the unique ID of the address book. You can call the [DescribeAddressBook](https://help.aliyun.com/document_detail/138869.html) operation to obtain the ID.
	//
	// example:
	//
	// b91d86c3-2b52-4534-aae9-8d0339b12a48
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The UNIX timestamp of the latest hit. The value is a UNIX timestamp measured in seconds.
	//
	// example:
	//
	// 1579261141
	HitLastTime *int64 `json:"HitLastTime,omitempty" xml:"HitLastTime,omitempty"`
	// The number of hits for the access control policy.
	//
	// example:
	//
	// 1
	HitTimes *int64 `json:"HitTimes,omitempty" xml:"HitTimes,omitempty"`
	// The IP address version. Valid values:
	//
	// - **4**: IPv4.
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
	// ngw-2ze4w62zbdkwjmoqeokgl
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The priority of the access control policy.
	//
	// The priority starts from 1 and increases in ascending order. A smaller value indicates a higher priority.
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
	// - **ANY**: all protocol types.
	//
	// > If you do not specify this parameter, all protocol types are queried.
	//
	// example:
	//
	// TCP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The number of times the address book is referenced.
	//
	// example:
	//
	// 1
	ReferenceCount *int32 `json:"ReferenceCount,omitempty" xml:"ReferenceCount,omitempty"`
	// The status of the access control policy. Valid values:
	//
	// - **true**: The access control policy is enabled.
	//
	// - **false**: The access control policy is disabled.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The recurrence days of the policy validity period for the access control policy.
	//
	// > If RepeatType is set to Weekly, valid values are 0 to 6. A week starts on Sunday.
	//
	// > If RepeatType is set to Monthly, valid values are 1 to 31.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The recurrence end time of the policy validity period for the access control policy. Example: 23:30. The time must be on the hour or half-hour, and must be at least 30 minutes later than the recurrence start time.
	//
	// > If RepeatType is set to Permanent or None, RepeatEndTime is empty. If RepeatType is set to Daily, Weekly, or Monthly, RepeatEndTime must be specified.
	//
	// > The time is in the HH:mm format (24-hour clock). Examples: 08:00 and 23:30.
	//
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The recurrence start time of the policy validity period for the access control policy. Example: 08:00. The time must be on the hour or half-hour, and must be at least 30 minutes earlier than the recurrence end time.
	//
	// > If RepeatType is set to Permanent or None, RepeatStartTime is empty. If RepeatType is set to Daily, Weekly, or Monthly, RepeatStartTime must be specified.
	//
	// > The time is in the HH:mm format (24-hour clock). Examples: 08:00 and 23:30.
	//
	// example:
	//
	// 08:00
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type of the policy validity period for the access control policy. Valid values:
	//
	// - **Permanent*	- (default): always.
	//
	// - **None**: specified single time.
	//
	// - **Daily**: every day.
	//
	// - **Weekly**: every week.
	//
	// - **Monthly**: every month.
	//
	// example:
	//
	// Permanent
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Valid values:
	//
	// - If **SourceType*	- is set to `net`, the source address is a CIDR block. Example: 192.0.XX.XX/24.
	//
	// - If **SourceType*	- is set to `group`, the source address is the name of the source address book. Example: db_group.
	//
	// - If **SourceType*	- is set to `location`, the source address is an area. For more information about area positional encoding, see [AddControlPolicy](https://help.aliyun.com/document_detail/138867.html). Example: ["BJ11", "ZB"].
	//
	// example:
	//
	// 172.28.7.167
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The CIDR blocks in the source address book of the access control policy.
	SourceGroupCidrs []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	// The type of the source address book in the access control policy. Valid values:
	//
	// - **ip**: IP address book, which contains one or more CIDR blocks.
	//
	// - **tag**: ECS tag-based address book, which contains the IP addresses of one or more ECS tags.
	//
	// - **domain**: domain name address book, which contains one or more domain names.
	//
	// - **threat**: threat intelligence address book, which contains one or more malicious IP addresses or domain names.
	//
	// - **backsrc**: back-to-origin address book, which contains the back-to-origin IP addresses of one or more Anti-DDoS or WAF instances.
	//
	// example:
	//
	// ip
	SourceGroupType *string `json:"SourceGroupType,omitempty" xml:"SourceGroupType,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// - **net**: source CIDR block.
	//
	// - **group**: source address book.
	//
	// - **location**: source location.
	//
	// example:
	//
	// group
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The number of access control policy quotas consumed. This value is the sum of the quotas consumed by each policy.
	//
	// Number of quotas consumed by a single policy = Number of source addresses (CIDR blocks or locations) × Number of destination addresses (CIDR blocks, locations, or domain names) × Number of port ranges × Number of applications.
	//
	// example:
	//
	// 10
	SpreadCnt *int32 `json:"SpreadCnt,omitempty" xml:"SpreadCnt,omitempty"`
	// The start time of the query. The value is a UNIX timestamp measured in seconds.
	//
	// example:
	//
	// 1730318400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The list of ECS tags.
	TagList []*DescribeAclCheckResponseBodyCheckRecordAclsAclTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The relationship between multiple ECS labels. Valid values:
	//
	// - **and**: The logical AND relation is used. The public IP addresses of Elastic Compute Service (ECS) instances that match all specified labels are added to the address book.
	//
	// - **or**: The logical OR relation is used. The public IP addresses of Elastic Compute Service (ECS) instances that match any of the specified labels are added to the address book.
	//
	// example:
	//
	// or
	TagRelation *string `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
	// The instance ID of the VPC border firewall.
	//
	// example:
	//
	// vfw-925514970c2c4bcab222
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeAclCheckResponseBodyCheckRecordAclsAcl) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclCheckResponseBodyCheckRecordAclsAcl) GoString() string {
	return s.String()
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetAclAction() *string {
	return s.AclAction
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetAclUuid() *string {
	return s.AclUuid
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetAddressList() []*string {
	return s.AddressList
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetAddressListCount() *int32 {
	return s.AddressListCount
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetAddresses() []*DescribeAclCheckResponseBodyCheckRecordAclsAclAddresses {
	return s.Addresses
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetApplicationNameList() []*string {
	return s.ApplicationNameList
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetAutoAddTagEcs() *int32 {
	return s.AutoAddTagEcs
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetDescription() *string {
	return s.Description
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetDestPort() *string {
	return s.DestPort
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetDestPortGroup() *string {
	return s.DestPortGroup
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetDestPortGroupPorts() []*string {
	return s.DestPortGroupPorts
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetDestPortType() *string {
	return s.DestPortType
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetDestination() *string {
	return s.Destination
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetDestinationGroupCidrs() []*string {
	return s.DestinationGroupCidrs
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetDestinationGroupType() *string {
	return s.DestinationGroupType
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetDestinationType() *string {
	return s.DestinationType
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetDirection() *string {
	return s.Direction
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetDnsResult() *string {
	return s.DnsResult
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetDnsResultTime() *int64 {
	return s.DnsResultTime
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetDomainResolveType() *int32 {
	return s.DomainResolveType
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetGroupUuid() *string {
	return s.GroupUuid
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetHitLastTime() *int64 {
	return s.HitLastTime
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetHitTimes() *int64 {
	return s.HitTimes
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetIpVersion() *int32 {
	return s.IpVersion
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetOrder() *int32 {
	return s.Order
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetProto() *string {
	return s.Proto
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetReferenceCount() *int32 {
	return s.ReferenceCount
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetRelease() *string {
	return s.Release
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetRepeatDays() []*int64 {
	return s.RepeatDays
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetRepeatEndTime() *string {
	return s.RepeatEndTime
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetRepeatStartTime() *string {
	return s.RepeatStartTime
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetRepeatType() *string {
	return s.RepeatType
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetSource() *string {
	return s.Source
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetSourceGroupCidrs() []*string {
	return s.SourceGroupCidrs
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetSourceGroupType() *string {
	return s.SourceGroupType
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetSpreadCnt() *int32 {
	return s.SpreadCnt
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetTagList() []*DescribeAclCheckResponseBodyCheckRecordAclsAclTagList {
	return s.TagList
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetTagRelation() *string {
	return s.TagRelation
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetAclAction(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.AclAction = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetAclUuid(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.AclUuid = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetAddressList(v []*string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.AddressList = v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetAddressListCount(v int32) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.AddressListCount = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetAddresses(v []*DescribeAclCheckResponseBodyCheckRecordAclsAclAddresses) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.Addresses = v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetApplicationId(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.ApplicationId = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetApplicationName(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.ApplicationName = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetApplicationNameList(v []*string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.ApplicationNameList = v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetAutoAddTagEcs(v int32) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.AutoAddTagEcs = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetCreateTime(v int64) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.CreateTime = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetDescription(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.Description = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetDestPort(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.DestPort = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetDestPortGroup(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.DestPortGroup = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetDestPortGroupPorts(v []*string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.DestPortGroupPorts = v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetDestPortType(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.DestPortType = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetDestination(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.Destination = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetDestinationGroupCidrs(v []*string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.DestinationGroupCidrs = v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetDestinationGroupType(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.DestinationGroupType = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetDestinationType(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.DestinationType = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetDirection(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.Direction = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetDnsResult(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.DnsResult = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetDnsResultTime(v int64) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.DnsResultTime = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetDomainResolveType(v int32) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.DomainResolveType = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetEndTime(v int64) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.EndTime = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetGroupName(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.GroupName = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetGroupType(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.GroupType = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetGroupUuid(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.GroupUuid = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetHitLastTime(v int64) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.HitLastTime = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetHitTimes(v int64) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.HitTimes = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetIpVersion(v int32) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.IpVersion = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetModifyTime(v int64) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.ModifyTime = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetNatGatewayId(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetOrder(v int32) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.Order = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetProto(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.Proto = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetReferenceCount(v int32) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.ReferenceCount = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetRelease(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.Release = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetRepeatDays(v []*int64) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.RepeatDays = v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetRepeatEndTime(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.RepeatEndTime = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetRepeatStartTime(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.RepeatStartTime = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetRepeatType(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.RepeatType = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetSource(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.Source = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetSourceGroupCidrs(v []*string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.SourceGroupCidrs = v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetSourceGroupType(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.SourceGroupType = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetSourceType(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.SourceType = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetSpreadCnt(v int32) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.SpreadCnt = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetStartTime(v int64) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.StartTime = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetTagList(v []*DescribeAclCheckResponseBodyCheckRecordAclsAclTagList) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.TagList = v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetTagRelation(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.TagRelation = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) SetVpcFirewallId(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAcl {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAcl) Validate() error {
	if s.Addresses != nil {
		for _, item := range s.Addresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type DescribeAclCheckResponseBodyCheckRecordAclsAclAddresses struct {
	// The address information in the address book.
	//
	// example:
	//
	// 192.0.XX.XX/32
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The description note.
	//
	// example:
	//
	// Reviewed
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s DescribeAclCheckResponseBodyCheckRecordAclsAclAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclCheckResponseBodyCheckRecordAclsAclAddresses) GoString() string {
	return s.String()
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAclAddresses) GetAddress() *string {
	return s.Address
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAclAddresses) GetNote() *string {
	return s.Note
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAclAddresses) SetAddress(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAclAddresses {
	s.Address = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAclAddresses) SetNote(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAclAddresses {
	s.Note = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAclAddresses) Validate() error {
	return dara.Validate(s)
}

type DescribeAclCheckResponseBodyCheckRecordAclsAclTagList struct {
	// The key of the ECS tag.
	//
	// example:
	//
	// ss
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the ECS tag.
	//
	// example:
	//
	// tfTestAcc0
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeAclCheckResponseBodyCheckRecordAclsAclTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclCheckResponseBodyCheckRecordAclsAclTagList) GoString() string {
	return s.String()
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAclTagList) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAclTagList) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAclTagList) SetTagKey(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAclTagList {
	s.TagKey = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAclTagList) SetTagValue(v string) *DescribeAclCheckResponseBodyCheckRecordAclsAclTagList {
	s.TagValue = &v
	return s
}

func (s *DescribeAclCheckResponseBodyCheckRecordAclsAclTagList) Validate() error {
	return dara.Validate(s)
}
