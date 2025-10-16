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
	CheckRecord *UpdateAclCheckDetailStatusResponseBodyCheckRecord `json:"CheckRecord,omitempty" xml:"CheckRecord,omitempty" type:"Struct"`
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
	Acls []*UpdateAclCheckDetailStatusResponseBodyCheckRecordAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Repeated"`
	// example:
	//
	// PolicyHitCountZero
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// example:
	//
	// dwd_mysql_lingwan_faxing_chat_config_di
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1724982259
	LastCheckTime *string `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
	// example:
	//
	// High
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 1
	PolicyTotalCount       *int64  `json:"PolicyTotalCount,omitempty" xml:"PolicyTotalCount,omitempty"`
	RecordAssessmentDetail *string `json:"RecordAssessmentDetail,omitempty" xml:"RecordAssessmentDetail,omitempty"`
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
	Acl                 *UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAcl `json:"Acl,omitempty" xml:"Acl,omitempty" type:"Struct"`
	AclAssessmentDetail *string                                                   `json:"AclAssessmentDetail,omitempty" xml:"AclAssessmentDetail,omitempty"`
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
	// example:
	//
	// log
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// example:
	//
	// 1e8ed1b2-cebc-4b95-a9cc-0cb7ce2c0c2c
	AclUuid     *string   `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	AddressList []*string `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	AddressListCount *int32 `json:"AddressListCount,omitempty" xml:"AddressListCount,omitempty"`
	// example:
	//
	// HTTP
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// ANY
	ApplicationName     *string   `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	AutoAddTagEcs *int32 `json:"AutoAddTagEcs,omitempty" xml:"AutoAddTagEcs,omitempty"`
	// example:
	//
	// 1761062400
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// test_policy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 22/22
	DestPort           *string   `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	DestPortGroup      *string   `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	DestPortGroupPorts []*string `json:"DestPortGroupPorts,omitempty" xml:"DestPortGroupPorts,omitempty" type:"Repeated"`
	// example:
	//
	// port
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// example:
	//
	// kms.cn-shanghai.aliyuncs.com
	Destination           *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationGroupCidrs []*string `json:"DestinationGroupCidrs,omitempty" xml:"DestinationGroupCidrs,omitempty" type:"Repeated"`
	// example:
	//
	// domain
	DestinationGroupType *string `json:"DestinationGroupType,omitempty" xml:"DestinationGroupType,omitempty"`
	// example:
	//
	// group
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// 192.168.0.1/32
	DnsResult *string `json:"DnsResult,omitempty" xml:"DnsResult,omitempty"`
	// example:
	//
	// 1579261141
	DnsResultTime *int64 `json:"DnsResultTime,omitempty" xml:"DnsResultTime,omitempty"`
	// example:
	//
	// 0
	DomainResolveType *int32 `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// example:
	//
	// 1752754426
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// subscribe
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// port
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// example:
	//
	// 5a96a798-9b73-47f7-831e-1d7aa3c987a9
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// example:
	//
	// 1579261141
	HitLastTime *int64 `json:"HitLastTime,omitempty" xml:"HitLastTime,omitempty"`
	// example:
	//
	// 1
	HitTimes *int64 `json:"HitTimes,omitempty" xml:"HitTimes,omitempty"`
	// example:
	//
	// 4
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// example:
	//
	// 1761062400
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// ngw-gw85zno51npz7lgc04z89
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// example:
	//
	// desc
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// ANY
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// example:
	//
	// 1
	ReferenceCount *int32 `json:"ReferenceCount,omitempty" xml:"ReferenceCount,omitempty"`
	// example:
	//
	// true
	Release    *string  `json:"Release,omitempty" xml:"Release,omitempty"`
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// example:
	//
	// 23:30
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// example:
	//
	// 08:00
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// example:
	//
	// None
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// example:
	//
	// 10.71.94.24
	Source           *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceGroupCidrs []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	// example:
	//
	// ip
	SourceGroupType *string `json:"SourceGroupType,omitempty" xml:"SourceGroupType,omitempty"`
	// example:
	//
	// net
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 10
	SpreadCnt *int32 `json:"SpreadCnt,omitempty" xml:"SpreadCnt,omitempty"`
	// example:
	//
	// 1736130347
	StartTime *int64                                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TagList   []*UpdateAclCheckDetailStatusResponseBodyCheckRecordAclsAclTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// example:
	//
	// and
	TagRelation *string `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
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
	// example:
	//
	// produce
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
