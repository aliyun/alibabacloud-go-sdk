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
	// The list of security events.
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
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApisecEventsResponseBodyData struct {
	// The total number of attacks in the security event.
	//
	// example:
	//
	// 10
	AllCnt *int64 `json:"AllCnt,omitempty" xml:"AllCnt,omitempty"`
	// The path of the API that is associated with the security event.
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
	// > Call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported business purposes.
	//
	// example:
	//
	// SendMail
	ApiTag *string `json:"ApiTag,omitempty" xml:"ApiTag,omitempty"`
	// The type of client that initiated the attack, such as a browser or automation tool.
	//
	// example:
	//
	// Chrome
	AttackClient *string `json:"AttackClient,omitempty" xml:"AttackClient,omitempty"`
	// Deprecated
	//
	// The attack count over time. The value is a JSON string in which each key is a UNIX timestamp in seconds and each value is the number of attacks at that time.
	//
	// example:
	//
	// {
	//
	//     "1717498320": 500,
	//
	//     "1717498380": 529,
	//
	//     "1717498440": 20
	//
	// }
	AttackCntInfo *string `json:"AttackCntInfo,omitempty" xml:"AttackCntInfo,omitempty"`
	// Deprecated
	//
	// The IP address of the attacker. 	Notice: This parameter is deprecated. Use the AttackIps parameter instead.
	//
	// example:
	//
	// 104.234.140.**
	AttackIp *string `json:"AttackIp,omitempty" xml:"AttackIp,omitempty"`
	// Deprecated
	//
	// The information about the attacker IP address. The value is a JSON string that contains the following fields:
	//
	// - **ip**: the IP address.
	//
	// - **country_id**: the country.
	//
	// - **region_id**: the region.
	//
	// - **cnt**: the number of attacks.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "ip": "72.*.*.119",
	//
	//         "country_id": "US",
	//
	//         "region_id": "",
	//
	//         "cnt": "2100"
	//
	//     }
	//
	// ]
	AttackIpInfo *string `json:"AttackIpInfo,omitempty" xml:"AttackIpInfo,omitempty"`
	// Deprecated
	//
	// The list of attacker IP addresses.
	AttackIps []*string `json:"AttackIps,omitempty" xml:"AttackIps,omitempty" type:"Repeated"`
	// The list of attackers that are associated with the security event.
	AttackerList []*string `json:"AttackerList,omitempty" xml:"AttackerList,omitempty" type:"Repeated"`
	// The end time of the event. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1683703260
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// The ID of the security event.
	//
	// example:
	//
	// c82cb276847e9c96f9597d9f4b0cdcff
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// Deprecated
	//
	// The details of the security event. The value is a JSON string that contains the following fields:
	//
	// - **ip_info**: the information about the attacker IP address. For more information, see the **AttackIpInfo*	- response parameter.
	//
	// - **rule_id**: the ID of the rule that corresponds to the event.
	//
	// - **rule_tag**: the information about the rule that corresponds to the event.
	//
	// example:
	//
	// {
	//
	//     "ip_info": [
	//
	//         {
	//
	//             "ip": "112.224.143.**",
	//
	//             "country_id": "CN",
	//
	//             "region_id": "-",
	//
	//             "cnt": "4"
	//
	//         }
	//
	//     ],
	//
	//     "rule_id": "837**",
	//
	//     "rule_tag": "interface returns a large amount of sensitive information"
	//
	// }
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// The severity level of the event. Valid values:
	//
	// - **high**: high severity.
	//
	// - **medium**: medium severity.
	//
	// - **low**: low severity.
	//
	// example:
	//
	// medium
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The event type.
	//
	// > Call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported event types.
	//
	// example:
	//
	// ObtainSensitiveUnauthorized
	EventTag *string `json:"EventTag,omitempty" xml:"EventTag,omitempty"`
	// Indicates whether the event is followed. Valid values:
	//
	// - **1**: The event is followed.
	//
	// - **0**: The event is not followed.
	//
	// example:
	//
	// 0
	Follow *int32 `json:"Follow,omitempty" xml:"Follow,omitempty"`
	// The domain name or IP address that is protected by WAF.
	//
	// example:
	//
	// a.***.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The remarks that are added to the security event.
	//
	// example:
	//
	// Notify
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The source of the event type. Valid values:
	//
	// - **custom**: a user-defined event type.
	//
	// - **default**: a built-in event type.
	//
	// example:
	//
	// custom
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The country where the attacker IP address is located.
	//
	// example:
	//
	// US
	RemoteCountry *string `json:"RemoteCountry,omitempty" xml:"RemoteCountry,omitempty"`
	// The region where the attacker IP address is located.
	//
	// example:
	//
	// 110000
	RemoteRegion *string `json:"RemoteRegion,omitempty" xml:"RemoteRegion,omitempty"`
	// Deprecated
	//
	// A sample of the API request data. The value is a JSON string.
	//
	// example:
	//
	// {}
	RequestData *string `json:"RequestData,omitempty" xml:"RequestData,omitempty"`
	// Deprecated
	//
	// A sample of the API response data. The value is a JSON string.
	//
	// example:
	//
	// {}
	ResponseData *string `json:"ResponseData,omitempty" xml:"ResponseData,omitempty"`
	// The start time of the event. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1683648000
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// The handling status of the event. Valid values:
	//
	// - **toBeConfirmed**: pending confirmation.
	//
	// - **confirmed**: confirmed but not yet handled.
	//
	// - **actioned**: handled.
	//
	// - **ignored**: ignored.
	//
	// example:
	//
	// toBeConfirmed
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

func (s *DescribeApisecEventsResponseBodyData) GetAttackerList() []*string {
	return s.AttackerList
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

func (s *DescribeApisecEventsResponseBodyData) SetAttackerList(v []*string) *DescribeApisecEventsResponseBodyData {
	s.AttackerList = v
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
