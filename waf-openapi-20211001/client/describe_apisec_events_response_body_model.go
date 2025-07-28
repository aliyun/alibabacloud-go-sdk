// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeApisecEventsResponseBodyData) *DescribeApisecEventsResponseBody
	GetData() []*DescribeApisecEventsResponseBodyData
	SetRequestId(v string) *DescribeApisecEventsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeApisecEventsResponseBody
	GetTotalCount() *int64
}

type DescribeApisecEventsResponseBody struct {
	// The security events.
	Data []*DescribeApisecEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 12F4CC8F-7E9F-5E4D-BF7C-BD1EDDE0C282
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisecEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecEventsResponseBody) GetData() []*DescribeApisecEventsResponseBodyData {
	return s.Data
}

func (s *DescribeApisecEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecEventsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeApisecEventsResponseBody) SetData(v []*DescribeApisecEventsResponseBodyData) *DescribeApisecEventsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApisecEventsResponseBody) SetRequestId(v string) *DescribeApisecEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecEventsResponseBody) SetTotalCount(v int64) *DescribeApisecEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisecEventsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApisecEventsResponseBodyData struct {
	// The number of attacks.
	//
	// 	Notice: The parameter has been deprecated, please use the Attackips parameter.
	//
	// example:
	//
	// 10
	AllCnt *int64 `json:"AllCnt,omitempty" xml:"AllCnt,omitempty"`
	// The API.
	//
	// example:
	//
	// /apisec/v1/register.php
	ApiFormat *string `json:"ApiFormat,omitempty" xml:"ApiFormat,omitempty"`
	// The ID of the API that is associated with the security event.
	//
	// example:
	//
	// 2ecc1cf67b91853bc55545052ccf06a8
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The business purpose of the API.
	//
	// >  You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the business purposes of APIs.
	//
	// example:
	//
	// SendMail
	ApiTag *string `json:"ApiTag,omitempty" xml:"ApiTag,omitempty"`
	// The client that is attacked.
	//
	// example:
	//
	// Chrome
	AttackClient *string `json:"AttackClient,omitempty" xml:"AttackClient,omitempty"`
	// The information about the number of attacks. The value of this parameter is a JSON string that contains multiple parameters. Key indicates the timestamp in seconds, and Value indicates the number of attacks.
	//
	// example:
	//
	// {\\"1717498320\\":500,\\"1717498380\\":529,\\"1717498440\\":20,\\"1717498260\\":518,\\"1717498200\\":481,\\"1717498140\\":52}
	AttackCntInfo *string `json:"AttackCntInfo,omitempty" xml:"AttackCntInfo,omitempty"`
	// The source IP address of the attack.
	//
	// example:
	//
	// 104.234.140.33
	AttackIp *string `json:"AttackIp,omitempty" xml:"AttackIp,omitempty"`
	// The information about the attack source IP address. The value of this parameter is a JSON string that contains multiple parameters. The value includes the following parameters:
	//
	// 	- **ip**: the IP address
	//
	// 	- **country_id**: the country ID
	//
	// 	- **region_id**: the region ID
	//
	// 	- **cnt**: the number of attacks
	//
	// example:
	//
	// [{\\"ip\\":\\"72.*.*.119\\",\\"country_id\\":\\"US\\",\\"region_id\\":\\"\\",\\"cnt\\":\\"2100\\"}]
	AttackIpInfo *string `json:"AttackIpInfo,omitempty" xml:"AttackIpInfo,omitempty"`
	// The source IP addresses of the attacks.
	AttackIps []*string `json:"AttackIps,omitempty" xml:"AttackIps,omitempty" type:"Repeated"`
	// The end of the time range to query. This value is a UNIX timestamp in UTC. Unit: seconds.
	//
	// example:
	//
	// 1683703260
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// The ID of the event.
	//
	// example:
	//
	// c82cb276847e9c96f9597d9f4b0cdcff
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The details of the event. The value of this parameter is a JSON string that contains multiple parameters. The value includes the following parameters:
	//
	// 	- **ip_info**: the information about the attack source IP address. This parameter corresponds to the **AttackIpInfo*	- response parameter.
	//
	// 	- **rule_id**: the ID of the rule corresponding to the event.
	//
	// 	- **rule_tag**: the information about the rule corresponding to the event.
	//
	// example:
	//
	// {}
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// The severity level of the event. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// medium
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The type of the event.
	//
	// >  You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported event types.
	//
	// example:
	//
	// ObtainSensitiveUnauthorized
	EventTag *string `json:"EventTag,omitempty" xml:"EventTag,omitempty"`
	// Indicates whether the API is followed. Valid values:
	//
	// 	- **1**: The API is followed.
	//
	// 	- **0**: The API is not followed.
	//
	// example:
	//
	// 0
	Follow *int32 `json:"Follow,omitempty" xml:"Follow,omitempty"`
	// The domain name or IP address of the API.
	//
	// example:
	//
	// a.aliyun.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The remarks.
	//
	// example:
	//
	// Notified
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The source of the event type. Valid values:
	//
	// 	- **custom**
	//
	// 	- **default**
	//
	// example:
	//
	// custom
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The country to which the attack source IP address belongs.
	//
	// example:
	//
	// US
	RemoteCountry *string `json:"RemoteCountry,omitempty" xml:"RemoteCountry,omitempty"`
	// The region to which the attack source IP address belongs.
	//
	// example:
	//
	// 110000
	RemoteRegion *string `json:"RemoteRegion,omitempty" xml:"RemoteRegion,omitempty"`
	// The sample API request. The value of this parameter is a JSON string that contains multiple parameters.
	//
	// example:
	//
	// {}
	RequestData *string `json:"RequestData,omitempty" xml:"RequestData,omitempty"`
	// The sample API response. The value of this parameter is a JSON string that contains multiple parameters.
	//
	// example:
	//
	// {}
	ResponseData *string `json:"ResponseData,omitempty" xml:"ResponseData,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp in UTC. Unit: seconds.
	//
	// example:
	//
	// 1683648000
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// The event status. Valid values:
	//
	// 	- **toBeConfirmed**
	//
	// 	- **confirmed**
	//
	// 	- **ignored**
	//
	// example:
	//
	// Ignore
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s DescribeApisecEventsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApisecEventsResponseBodyData) GetAllCnt() *int64 {
	return s.AllCnt
}

func (s *DescribeApisecEventsResponseBodyData) GetApiFormat() *string {
	return s.ApiFormat
}

func (s *DescribeApisecEventsResponseBodyData) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisecEventsResponseBodyData) GetApiTag() *string {
	return s.ApiTag
}

func (s *DescribeApisecEventsResponseBodyData) GetAttackClient() *string {
	return s.AttackClient
}

func (s *DescribeApisecEventsResponseBodyData) GetAttackCntInfo() *string {
	return s.AttackCntInfo
}

func (s *DescribeApisecEventsResponseBodyData) GetAttackIp() *string {
	return s.AttackIp
}

func (s *DescribeApisecEventsResponseBodyData) GetAttackIpInfo() *string {
	return s.AttackIpInfo
}

func (s *DescribeApisecEventsResponseBodyData) GetAttackIps() []*string {
	return s.AttackIps
}

func (s *DescribeApisecEventsResponseBodyData) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribeApisecEventsResponseBodyData) GetEventId() *string {
	return s.EventId
}

func (s *DescribeApisecEventsResponseBodyData) GetEventInfo() *string {
	return s.EventInfo
}

func (s *DescribeApisecEventsResponseBodyData) GetEventLevel() *string {
	return s.EventLevel
}

func (s *DescribeApisecEventsResponseBodyData) GetEventTag() *string {
	return s.EventTag
}

func (s *DescribeApisecEventsResponseBodyData) GetFollow() *int32 {
	return s.Follow
}

func (s *DescribeApisecEventsResponseBodyData) GetMatchedHost() *string {
	return s.MatchedHost
}

func (s *DescribeApisecEventsResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *DescribeApisecEventsResponseBodyData) GetOrigin() *string {
	return s.Origin
}

func (s *DescribeApisecEventsResponseBodyData) GetRemoteCountry() *string {
	return s.RemoteCountry
}

func (s *DescribeApisecEventsResponseBodyData) GetRemoteRegion() *string {
	return s.RemoteRegion
}

func (s *DescribeApisecEventsResponseBodyData) GetRequestData() *string {
	return s.RequestData
}

func (s *DescribeApisecEventsResponseBodyData) GetResponseData() *string {
	return s.ResponseData
}

func (s *DescribeApisecEventsResponseBodyData) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeApisecEventsResponseBodyData) GetUserStatus() *string {
	return s.UserStatus
}

func (s *DescribeApisecEventsResponseBodyData) SetAllCnt(v int64) *DescribeApisecEventsResponseBodyData {
	s.AllCnt = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetApiFormat(v string) *DescribeApisecEventsResponseBodyData {
	s.ApiFormat = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetApiId(v string) *DescribeApisecEventsResponseBodyData {
	s.ApiId = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetApiTag(v string) *DescribeApisecEventsResponseBodyData {
	s.ApiTag = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetAttackClient(v string) *DescribeApisecEventsResponseBodyData {
	s.AttackClient = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetAttackCntInfo(v string) *DescribeApisecEventsResponseBodyData {
	s.AttackCntInfo = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetAttackIp(v string) *DescribeApisecEventsResponseBodyData {
	s.AttackIp = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetAttackIpInfo(v string) *DescribeApisecEventsResponseBodyData {
	s.AttackIpInfo = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetAttackIps(v []*string) *DescribeApisecEventsResponseBodyData {
	s.AttackIps = v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetEndTs(v int64) *DescribeApisecEventsResponseBodyData {
	s.EndTs = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetEventId(v string) *DescribeApisecEventsResponseBodyData {
	s.EventId = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetEventInfo(v string) *DescribeApisecEventsResponseBodyData {
	s.EventInfo = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetEventLevel(v string) *DescribeApisecEventsResponseBodyData {
	s.EventLevel = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetEventTag(v string) *DescribeApisecEventsResponseBodyData {
	s.EventTag = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetFollow(v int32) *DescribeApisecEventsResponseBodyData {
	s.Follow = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetMatchedHost(v string) *DescribeApisecEventsResponseBodyData {
	s.MatchedHost = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetNote(v string) *DescribeApisecEventsResponseBodyData {
	s.Note = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetOrigin(v string) *DescribeApisecEventsResponseBodyData {
	s.Origin = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetRemoteCountry(v string) *DescribeApisecEventsResponseBodyData {
	s.RemoteCountry = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetRemoteRegion(v string) *DescribeApisecEventsResponseBodyData {
	s.RemoteRegion = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetRequestData(v string) *DescribeApisecEventsResponseBodyData {
	s.RequestData = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetResponseData(v string) *DescribeApisecEventsResponseBodyData {
	s.ResponseData = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetStartTs(v int64) *DescribeApisecEventsResponseBodyData {
	s.StartTs = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) SetUserStatus(v string) *DescribeApisecEventsResponseBodyData {
	s.UserStatus = &v
	return s
}

func (s *DescribeApisecEventsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
