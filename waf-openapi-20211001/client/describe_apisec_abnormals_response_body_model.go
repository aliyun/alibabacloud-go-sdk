// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecAbnormalsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeApisecAbnormalsResponseBodyData) *DescribeApisecAbnormalsResponseBody
	GetData() []*DescribeApisecAbnormalsResponseBodyData
	SetRequestId(v string) *DescribeApisecAbnormalsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeApisecAbnormalsResponseBody
	GetTotalCount() *int64
}

type DescribeApisecAbnormalsResponseBody struct {
	// The list of security risks.
	Data []*DescribeApisecAbnormalsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9469646C-F2CC-5F0F-8401-C53***4F46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of security risks returned.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisecAbnormalsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecAbnormalsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecAbnormalsResponseBody) GetData() []*DescribeApisecAbnormalsResponseBodyData {
	return s.Data
}

func (s *DescribeApisecAbnormalsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecAbnormalsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeApisecAbnormalsResponseBody) SetData(v []*DescribeApisecAbnormalsResponseBodyData) *DescribeApisecAbnormalsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApisecAbnormalsResponseBody) SetRequestId(v string) *DescribeApisecAbnormalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBody) SetTotalCount(v int64) *DescribeApisecAbnormalsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBody) Validate() error {
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

type DescribeApisecAbnormalsResponseBodyData struct {
	// The number of risk events that are associated with the security risk.
	//
	// example:
	//
	// 2
	AbnormalEventNumber *int64 `json:"AbnormalEventNumber,omitempty" xml:"AbnormalEventNumber,omitempty"`
	// The ID of the security risk.
	//
	// example:
	//
	// 7c1431f27ae7e9c8cc64095***68e
	AbnormalId *string `json:"AbnormalId,omitempty" xml:"AbnormalId,omitempty"`
	// The details of the security risk in JSON format. The JSON object contains the following fields:
	//
	// - **rule**: The detection rule that triggered the security risk.
	//
	// - **data_type**: The type of sensitive data.
	//
	// - **custom_rule_name**: The name of the custom rule.
	//
	// - **rule_name**: The name of the built-in rule.
	//
	// example:
	//
	// {
	//
	//     "data_type": ["1005","1004"],
	//
	//     "rule": {
	//
	//         "parent": "RiskType_Permission",
	//
	//         "code": "Risk_UnauthSensitive",
	//
	//         "level": "high",
	//
	//         "origin": "default",
	//
	//         "name": "Risk_UnauthSensitive"
	//
	//     }
	//
	// }
	AbnormalInfo *string `json:"AbnormalInfo,omitempty" xml:"AbnormalInfo,omitempty"`
	// The severity level of the security risk. Valid values:
	//
	// - **high**: High.
	//
	// - **medium**: Medium.
	//
	// - **low**: Low.
	//
	// example:
	//
	// high
	AbnormalLevel *string `json:"AbnormalLevel,omitempty" xml:"AbnormalLevel,omitempty"`
	// The type of the security risk.
	//
	// > Call [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) to query the supported risk types.
	//
	// example:
	//
	// LackOfSpeedLimit
	AbnormalTag *string `json:"AbnormalTag,omitempty" xml:"AbnormalTag,omitempty"`
	// The processing status of the security risk.
	//
	// example:
	//
	// unresolved
	AbnromalStatus *string `json:"AbnromalStatus,omitempty" xml:"AbnromalStatus,omitempty"`
	// The path of the API that is associated with the security risk.
	//
	// example:
	//
	// /api/login
	ApiFormat *string `json:"ApiFormat,omitempty" xml:"ApiFormat,omitempty"`
	// The ID of the API that is associated with the security risk.
	//
	// example:
	//
	// 09559c0d71ca2ffc996b81***836d8
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The business purpose of the API.
	//
	// > Call [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) to query the supported business purposes.
	//
	// example:
	//
	// SendMail
	ApiTag *string `json:"ApiTag,omitempty" xml:"ApiTag,omitempty"`
	// The time when the security risk was first detected. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1684252800
	DiscoverTime *int64 `json:"DiscoverTime,omitempty" xml:"DiscoverTime,omitempty"`
	// Deprecated
	//
	// The list of security risk samples.
	Examples []*string `json:"Examples,omitempty" xml:"Examples,omitempty" type:"Repeated"`
	// The time when the API was first discovered. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1701138088
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// Indicates whether the security risk is being followed. Valid values:
	//
	// - **1**: The security risk is being followed.
	//
	// - **0*	- (default): The security risk is not being followed.
	//
	// example:
	//
	// 0
	Follow *int64 `json:"Follow,omitempty" xml:"Follow,omitempty"`
	// The time when the security risk was marked as ignored. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1684252800
	IgnoreTime *int64 `json:"IgnoreTime,omitempty" xml:"IgnoreTime,omitempty"`
	// The most recent time when the API was accessed. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1684252800
	LastestTime *int64 `json:"LastestTime,omitempty" xml:"LastestTime,omitempty"`
	// The most recent time when the security risk was detected. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1684252800
	LatestDiscoverTime *int64 `json:"LatestDiscoverTime,omitempty" xml:"LatestDiscoverTime,omitempty"`
	// The domain name or IP address that the API resides on.
	//
	// example:
	//
	// a.aliyun.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The remarks for the security risk.
	//
	// example:
	//
	// Strict
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The source of the risk detection rule. Valid values:
	//
	// - **custom**: Custom rule.
	//
	// - **default**: Built-in rule.
	//
	// example:
	//
	// custom
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The handling status of the security risk. Valid values:
	//
	// - **toBeConfirmed**: To be confirmed.
	//
	// - **confirmed**: Confirmed.
	//
	// - **toBeFixed**: To be fixed.
	//
	// - **fixed**: Fixed (manually verified).
	//
	// - **ignored**: Ignored.
	//
	// - **toBeVerified**: To be verified by the system.
	//
	// - **notFixed**: Verification failed.
	//
	// - **systemFixed**: Fixed (verified by the system).
	//
	// example:
	//
	// Confirmed
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s DescribeApisecAbnormalsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecAbnormalsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetAbnormalEventNumber() *int64 {
	return s.AbnormalEventNumber
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetAbnormalId() *string {
	return s.AbnormalId
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetAbnormalInfo() *string {
	return s.AbnormalInfo
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetAbnormalLevel() *string {
	return s.AbnormalLevel
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetAbnormalTag() *string {
	return s.AbnormalTag
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetAbnromalStatus() *string {
	return s.AbnromalStatus
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetApiFormat() *string {
	return s.ApiFormat
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetApiTag() *string {
	return s.ApiTag
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetDiscoverTime() *int64 {
	return s.DiscoverTime
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetExamples() []*string {
	return s.Examples
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetFollow() *int64 {
	return s.Follow
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetIgnoreTime() *int64 {
	return s.IgnoreTime
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetLastestTime() *int64 {
	return s.LastestTime
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetLatestDiscoverTime() *int64 {
	return s.LatestDiscoverTime
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetMatchedHost() *string {
	return s.MatchedHost
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetOrigin() *string {
	return s.Origin
}

func (s *DescribeApisecAbnormalsResponseBodyData) GetUserStatus() *string {
	return s.UserStatus
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetAbnormalEventNumber(v int64) *DescribeApisecAbnormalsResponseBodyData {
	s.AbnormalEventNumber = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetAbnormalId(v string) *DescribeApisecAbnormalsResponseBodyData {
	s.AbnormalId = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetAbnormalInfo(v string) *DescribeApisecAbnormalsResponseBodyData {
	s.AbnormalInfo = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetAbnormalLevel(v string) *DescribeApisecAbnormalsResponseBodyData {
	s.AbnormalLevel = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetAbnormalTag(v string) *DescribeApisecAbnormalsResponseBodyData {
	s.AbnormalTag = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetAbnromalStatus(v string) *DescribeApisecAbnormalsResponseBodyData {
	s.AbnromalStatus = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetApiFormat(v string) *DescribeApisecAbnormalsResponseBodyData {
	s.ApiFormat = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetApiId(v string) *DescribeApisecAbnormalsResponseBodyData {
	s.ApiId = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetApiTag(v string) *DescribeApisecAbnormalsResponseBodyData {
	s.ApiTag = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetDiscoverTime(v int64) *DescribeApisecAbnormalsResponseBodyData {
	s.DiscoverTime = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetExamples(v []*string) *DescribeApisecAbnormalsResponseBodyData {
	s.Examples = v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetFirstTime(v int64) *DescribeApisecAbnormalsResponseBodyData {
	s.FirstTime = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetFollow(v int64) *DescribeApisecAbnormalsResponseBodyData {
	s.Follow = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetIgnoreTime(v int64) *DescribeApisecAbnormalsResponseBodyData {
	s.IgnoreTime = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetLastestTime(v int64) *DescribeApisecAbnormalsResponseBodyData {
	s.LastestTime = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetLatestDiscoverTime(v int64) *DescribeApisecAbnormalsResponseBodyData {
	s.LatestDiscoverTime = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetMatchedHost(v string) *DescribeApisecAbnormalsResponseBodyData {
	s.MatchedHost = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetNote(v string) *DescribeApisecAbnormalsResponseBodyData {
	s.Note = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetOrigin(v string) *DescribeApisecAbnormalsResponseBodyData {
	s.Origin = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) SetUserStatus(v string) *DescribeApisecAbnormalsResponseBodyData {
	s.UserStatus = &v
	return s
}

func (s *DescribeApisecAbnormalsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
