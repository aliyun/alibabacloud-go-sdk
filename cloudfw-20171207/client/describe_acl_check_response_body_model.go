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
	CheckRecord *DescribeAclCheckResponseBodyCheckRecord `json:"CheckRecord,omitempty" xml:"CheckRecord,omitempty" type:"Struct"`
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
	// example:
	//
	// 10
	AclTotalCount *int64                                         `json:"AclTotalCount,omitempty" xml:"AclTotalCount,omitempty"`
	Acls          []*DescribeAclCheckResponseBodyCheckRecordAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Repeated"`
	// example:
	//
	// PolicyHitCountZero
	CheckName   *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1724982259
	LastCheckTime *string `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
	// example:
	//
	// High
	Level                  *string `json:"Level,omitempty" xml:"Level,omitempty"`
	RecordAssessmentDetail *string `json:"RecordAssessmentDetail,omitempty" xml:"RecordAssessmentDetail,omitempty"`
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
	Acl                 *DescribeAclCheckResponseBodyCheckRecordAclsAcl `json:"Acl,omitempty" xml:"Acl,omitempty" type:"Struct"`
	AclAssessmentDetail *string                                         `json:"AclAssessmentDetail,omitempty" xml:"AclAssessmentDetail,omitempty"`
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
	// example:
	//
	// log
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// example:
	//
	// 997b38e0-01fa-4db7-8d30-02ebf6fdb747
	AclUuid     *string   `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	AddressList []*string `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	AddressListCount *int32                                                     `json:"AddressListCount,omitempty" xml:"AddressListCount,omitempty"`
	Addresses        []*DescribeAclCheckResponseBodyCheckRecordAclsAclAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// example:
	//
	// plugin_idp4_ciam
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
	// 80/80
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// example:
	//
	// my_port_group
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
	// domain
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// 192.0.XX.XX
	DnsResult *string `json:"DnsResult,omitempty" xml:"DnsResult,omitempty"`
	// example:
	//
	// 1579261141
	DnsResultTime *int64 `json:"DnsResultTime,omitempty" xml:"DnsResultTime,omitempty"`
	// example:
	//
	// FQDN
	DomainResolveType *int32 `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// example:
	//
	// 1758334822
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// ip
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// example:
	//
	// b91d86c3-2b52-4534-aae9-8d0339b12a48
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
	// ngw-2ze4w62zbdkwjmoqeokgl
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// example:
	//
	// 1
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// TCP
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
	// Permanent
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// example:
	//
	// 172.28.7.167
	Source           *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceGroupCidrs []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	// example:
	//
	// ip
	SourceGroupType *string `json:"SourceGroupType,omitempty" xml:"SourceGroupType,omitempty"`
	// example:
	//
	// group
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 10
	SpreadCnt *int32 `json:"SpreadCnt,omitempty" xml:"SpreadCnt,omitempty"`
	// example:
	//
	// 1730318400
	StartTime *int64                                                   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TagList   []*DescribeAclCheckResponseBodyCheckRecordAclsAclTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// example:
	//
	// or
	TagRelation *string `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
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
	// example:
	//
	// 192.0.XX.XX/32
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Note    *string `json:"Note,omitempty" xml:"Note,omitempty"`
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
	// example:
	//
	// ss
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
