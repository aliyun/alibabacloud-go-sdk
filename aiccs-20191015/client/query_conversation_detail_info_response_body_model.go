// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConversationDetailInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryConversationDetailInfoResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryConversationDetailInfoResponseBody
	GetCode() *string
	SetData(v *QueryConversationDetailInfoResponseBodyData) *QueryConversationDetailInfoResponseBody
	GetData() *QueryConversationDetailInfoResponseBodyData
	SetMessage(v string) *QueryConversationDetailInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryConversationDetailInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryConversationDetailInfoResponseBody
	GetSuccess() *bool
}

type QueryConversationDetailInfoResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *QueryConversationDetailInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The status code message.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F2051E18-FF3F-5C08-8D24-6F150D2AF757
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryConversationDetailInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryConversationDetailInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConversationDetailInfoResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryConversationDetailInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryConversationDetailInfoResponseBody) GetData() *QueryConversationDetailInfoResponseBodyData {
	return s.Data
}

func (s *QueryConversationDetailInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryConversationDetailInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryConversationDetailInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryConversationDetailInfoResponseBody) SetAccessDeniedDetail(v string) *QueryConversationDetailInfoResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBody) SetCode(v string) *QueryConversationDetailInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBody) SetData(v *QueryConversationDetailInfoResponseBodyData) *QueryConversationDetailInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryConversationDetailInfoResponseBody) SetMessage(v string) *QueryConversationDetailInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBody) SetRequestId(v string) *QueryConversationDetailInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBody) SetSuccess(v bool) *QueryConversationDetailInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryConversationDetailInfoResponseBodyData struct {
	// The unique call ID.
	//
	// example:
	//
	// 1231231231213^11231231231
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The call result. Valid values:
	//
	// - `CALL_FORWARDING`: Call forwarding.
	//
	// - `INCOMING_CALL_BARRED`: Incoming call barred.
	//
	// - `CALL_REJECTED`: Call rejected.
	//
	// - `ANSWERED`: Answered by user.
	//
	// - `USER_BUSY`: Called party busy.
	//
	// - `POWERED_OFF`: Powered off.
	//
	// - `NO_USER_RESPONSE`: Out of service area.
	//
	// - `OPERATOR_BLOCK`: Blocked by carrier.
	//
	// - `OTHERS`: Other.
	//
	// - `SUSPEND`: Suspended.
	//
	// - `CANCEL`: Canceled by caller.
	//
	// - `INVALID_NUMBER`: Invalid number.
	//
	// - `UNAVAILABLE`: Temporarily unavailable.
	//
	// - `NETWORK_BUSY`: Network busy.
	//
	// - `NO_ANSWER`: No answer.
	//
	// example:
	//
	// ANSWERED
	CallResult *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	// The called number.
	//
	// example:
	//
	// 186******
	CalledPhone *string `json:"CalledPhone,omitempty" xml:"CalledPhone,omitempty"`
	// The caller number.
	//
	// example:
	//
	// 0571*******
	CallerPhone *string `json:"CallerPhone,omitempty" xml:"CallerPhone,omitempty"`
	// The conversation record. The structure is a JSON array in which entries are sorted by time. Example:
	//
	// ```json
	//
	// [
	//
	//     {
	//
	//         "content":"Conversation content",
	//
	//         "role":"Role", // Valid values: user, assistant
	//
	//     }
	//
	// ]
	//
	// ```
	//
	// example:
	//
	// [
	//
	//   {
	//
	//     "content": "111您好，年龄222，性别男，我这边是**汽车的官方顾问，我们新出了一款车型为**；**已经上市了，售价**万元起，**分钟破*台，您看要不了解一下？",
	//
	//     "role": "assistant"
	//
	//   },
	//
	//   {
	//
	//     "content": "<客户打断>哎，你是谁？",
	//
	//     "role": "user"
	//
	//   },
	//
	//   {
	//
	//     "content": "<客户打断>你再说一遍。",
	//
	//     "role": "user"
	//
	//   },
	//
	//   {
	//
	//     "content": "哎，我没听清。",
	//
	//     "role": "user"
	//
	//   },
	//
	//   {
	//
	//     "content": "你在说什么？",
	//
	//     "role": "user"
	//
	//   },
	//
	//   {
	//
	//     "content": "您好，",
	//
	//     "role": "assistant"
	//
	//   },
	//
	//   {
	//
	//     "content": "我是**汽车总部销售服务顾问。",
	//
	//     "role": "assistant"
	//
	//   },
	//
	//   {
	//
	//     "content": "我们最近推出了一款新车**，想了解一下您是否对这款车型感兴趣？",
	//
	//     "role": "assistant"
	//
	//   },
	//
	//   {
	//
	//     "content": "<客户打断>哎，那我是谁？",
	//
	//     "role": "user"
	//
	//   },
	//
	//   {
	//
	//     "content": "你在说什么呢？",
	//
	//     "role": "user"
	//
	//   },
	//
	//   {
	//
	//     "content": "抱歉打扰了，111先生。",
	//
	//     "role": "assistant"
	//
	//   },
	//
	//   {
	//
	//     "content": "祝您生活愉快！再见！",
	//
	//     "role": "assistant"
	//
	//   }
	//
	// ]
	ConversationRecord *string `json:"ConversationRecord,omitempty" xml:"ConversationRecord,omitempty"`
	// The duration of the call, in seconds. If the call was not connected, the value is 0.
	//
	// example:
	//
	// 16
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// The failure reason.
	//
	// example:
	//
	// 主动取消
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// The party that hung up. Valid values:
	//
	// - **0**: user.
	//
	// - **1**: assistant.
	//
	// example:
	//
	// 用户
	HangupDirection *string `json:"HangupDirection,omitempty" xml:"HangupDirection,omitempty"`
	// The primary intent.
	//
	// example:
	//
	// D
	MajorIntent *string `json:"MajorIntent,omitempty" xml:"MajorIntent,omitempty"`
	// The business-specific ID that is passed in. You can use this unique ID to associate the call with your business.
	//
	// example:
	//
	// bb3bc32d-54b8-49c4-80d3-61583417d22e
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// A list of output tags.
	OutputTags []*QueryConversationDetailInfoResponseBodyDataOutputTags `json:"OutputTags,omitempty" xml:"OutputTags,omitempty" type:"Repeated"`
	// The timestamp when the call was answered, in milliseconds.
	//
	// example:
	//
	// 1754617273000
	PickUpTime *int64 `json:"PickUpTime,omitempty" xml:"PickUpTime,omitempty"`
	// The download URL for the recording file. This parameter is returned only after the recording file is generated.
	//
	// example:
	//
	// https://********
	RecordingFileDownloadUrl *string `json:"RecordingFileDownloadUrl,omitempty" xml:"RecordingFileDownloadUrl,omitempty"`
	// The timestamp when the call ended, in milliseconds.
	//
	// example:
	//
	// 98
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The timestamp when the call was initiated, in milliseconds.
	//
	// example:
	//
	// 123123123123123
	StartCallTime *int64 `json:"StartCallTime,omitempty" xml:"StartCallTime,omitempty"`
	// The call status code. For more information, see [Call status codes](https://help.aliyun.com/document_detail/112804.html) for the voice service.
	//
	// example:
	//
	// 200005
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The status message returned by the carrier.
	//
	// example:
	//
	// 呼叫结束（双呼）
	StatusMsg *string `json:"StatusMsg,omitempty" xml:"StatusMsg,omitempty"`
	// A list of variables associated with the call task.
	Variables []*QueryConversationDetailInfoResponseBodyDataVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s QueryConversationDetailInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryConversationDetailInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryConversationDetailInfoResponseBodyData) GetCallId() *string {
	return s.CallId
}

func (s *QueryConversationDetailInfoResponseBodyData) GetCallResult() *string {
	return s.CallResult
}

func (s *QueryConversationDetailInfoResponseBodyData) GetCalledPhone() *string {
	return s.CalledPhone
}

func (s *QueryConversationDetailInfoResponseBodyData) GetCallerPhone() *string {
	return s.CallerPhone
}

func (s *QueryConversationDetailInfoResponseBodyData) GetConversationRecord() *string {
	return s.ConversationRecord
}

func (s *QueryConversationDetailInfoResponseBodyData) GetDuration() *int64 {
	return s.Duration
}

func (s *QueryConversationDetailInfoResponseBodyData) GetEncryptionType() *string {
	return s.EncryptionType
}

func (s *QueryConversationDetailInfoResponseBodyData) GetFailedReason() *string {
	return s.FailedReason
}

func (s *QueryConversationDetailInfoResponseBodyData) GetHangupDirection() *string {
	return s.HangupDirection
}

func (s *QueryConversationDetailInfoResponseBodyData) GetMajorIntent() *string {
	return s.MajorIntent
}

func (s *QueryConversationDetailInfoResponseBodyData) GetOutId() *string {
	return s.OutId
}

func (s *QueryConversationDetailInfoResponseBodyData) GetOutputTags() []*QueryConversationDetailInfoResponseBodyDataOutputTags {
	return s.OutputTags
}

func (s *QueryConversationDetailInfoResponseBodyData) GetPickUpTime() *int64 {
	return s.PickUpTime
}

func (s *QueryConversationDetailInfoResponseBodyData) GetRecordingFileDownloadUrl() *string {
	return s.RecordingFileDownloadUrl
}

func (s *QueryConversationDetailInfoResponseBodyData) GetReleaseTime() *int64 {
	return s.ReleaseTime
}

func (s *QueryConversationDetailInfoResponseBodyData) GetStartCallTime() *int64 {
	return s.StartCallTime
}

func (s *QueryConversationDetailInfoResponseBodyData) GetStatusCode() *string {
	return s.StatusCode
}

func (s *QueryConversationDetailInfoResponseBodyData) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *QueryConversationDetailInfoResponseBodyData) GetVariables() []*QueryConversationDetailInfoResponseBodyDataVariables {
	return s.Variables
}

func (s *QueryConversationDetailInfoResponseBodyData) SetCallId(v string) *QueryConversationDetailInfoResponseBodyData {
	s.CallId = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetCallResult(v string) *QueryConversationDetailInfoResponseBodyData {
	s.CallResult = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetCalledPhone(v string) *QueryConversationDetailInfoResponseBodyData {
	s.CalledPhone = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetCallerPhone(v string) *QueryConversationDetailInfoResponseBodyData {
	s.CallerPhone = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetConversationRecord(v string) *QueryConversationDetailInfoResponseBodyData {
	s.ConversationRecord = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetDuration(v int64) *QueryConversationDetailInfoResponseBodyData {
	s.Duration = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetEncryptionType(v string) *QueryConversationDetailInfoResponseBodyData {
	s.EncryptionType = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetFailedReason(v string) *QueryConversationDetailInfoResponseBodyData {
	s.FailedReason = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetHangupDirection(v string) *QueryConversationDetailInfoResponseBodyData {
	s.HangupDirection = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetMajorIntent(v string) *QueryConversationDetailInfoResponseBodyData {
	s.MajorIntent = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetOutId(v string) *QueryConversationDetailInfoResponseBodyData {
	s.OutId = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetOutputTags(v []*QueryConversationDetailInfoResponseBodyDataOutputTags) *QueryConversationDetailInfoResponseBodyData {
	s.OutputTags = v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetPickUpTime(v int64) *QueryConversationDetailInfoResponseBodyData {
	s.PickUpTime = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetRecordingFileDownloadUrl(v string) *QueryConversationDetailInfoResponseBodyData {
	s.RecordingFileDownloadUrl = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetReleaseTime(v int64) *QueryConversationDetailInfoResponseBodyData {
	s.ReleaseTime = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetStartCallTime(v int64) *QueryConversationDetailInfoResponseBodyData {
	s.StartCallTime = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetStatusCode(v string) *QueryConversationDetailInfoResponseBodyData {
	s.StatusCode = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetStatusMsg(v string) *QueryConversationDetailInfoResponseBodyData {
	s.StatusMsg = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) SetVariables(v []*QueryConversationDetailInfoResponseBodyDataVariables) *QueryConversationDetailInfoResponseBodyData {
	s.Variables = v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyData) Validate() error {
	if s.OutputTags != nil {
		for _, item := range s.OutputTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryConversationDetailInfoResponseBodyDataOutputTags struct {
	// The tag ID.
	//
	// example:
	//
	// 9ca2*****************************
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The tag description.
	//
	// example:
	//
	// 评估客户对车型的兴趣和购买可能性
	OutputTagDescription *string `json:"OutputTagDescription,omitempty" xml:"OutputTagDescription,omitempty"`
	// The tag name.
	//
	// example:
	//
	// 客户意向度
	OutputTagName *string `json:"OutputTagName,omitempty" xml:"OutputTagName,omitempty"`
	// The tag value.
	//
	// example:
	//
	// ["高（非常积极，大概率转化）"]
	OutputTagValue *string `json:"OutputTagValue,omitempty" xml:"OutputTagValue,omitempty"`
}

func (s QueryConversationDetailInfoResponseBodyDataOutputTags) String() string {
	return dara.Prettify(s)
}

func (s QueryConversationDetailInfoResponseBodyDataOutputTags) GoString() string {
	return s.String()
}

func (s *QueryConversationDetailInfoResponseBodyDataOutputTags) GetId() *string {
	return s.Id
}

func (s *QueryConversationDetailInfoResponseBodyDataOutputTags) GetOutputTagDescription() *string {
	return s.OutputTagDescription
}

func (s *QueryConversationDetailInfoResponseBodyDataOutputTags) GetOutputTagName() *string {
	return s.OutputTagName
}

func (s *QueryConversationDetailInfoResponseBodyDataOutputTags) GetOutputTagValue() *string {
	return s.OutputTagValue
}

func (s *QueryConversationDetailInfoResponseBodyDataOutputTags) SetId(v string) *QueryConversationDetailInfoResponseBodyDataOutputTags {
	s.Id = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyDataOutputTags) SetOutputTagDescription(v string) *QueryConversationDetailInfoResponseBodyDataOutputTags {
	s.OutputTagDescription = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyDataOutputTags) SetOutputTagName(v string) *QueryConversationDetailInfoResponseBodyDataOutputTags {
	s.OutputTagName = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyDataOutputTags) SetOutputTagValue(v string) *QueryConversationDetailInfoResponseBodyDataOutputTags {
	s.OutputTagValue = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyDataOutputTags) Validate() error {
	return dara.Validate(s)
}

type QueryConversationDetailInfoResponseBodyDataVariables struct {
	// The variable ID.
	//
	// example:
	//
	// 22
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The variable key.
	//
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The variable name.
	//
	// example:
	//
	// 姓名
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the variable is required. Valid values:
	//
	// - `true`: The variable is required.
	//
	// - `false`: The variable is optional.
	//
	// example:
	//
	// false
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// 示例值
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The variable value.
	//
	// example:
	//
	// 张三
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryConversationDetailInfoResponseBodyDataVariables) String() string {
	return dara.Prettify(s)
}

func (s QueryConversationDetailInfoResponseBodyDataVariables) GoString() string {
	return s.String()
}

func (s *QueryConversationDetailInfoResponseBodyDataVariables) GetId() *string {
	return s.Id
}

func (s *QueryConversationDetailInfoResponseBodyDataVariables) GetKey() *string {
	return s.Key
}

func (s *QueryConversationDetailInfoResponseBodyDataVariables) GetName() *string {
	return s.Name
}

func (s *QueryConversationDetailInfoResponseBodyDataVariables) GetRequired() *bool {
	return s.Required
}

func (s *QueryConversationDetailInfoResponseBodyDataVariables) GetSource() *string {
	return s.Source
}

func (s *QueryConversationDetailInfoResponseBodyDataVariables) GetValue() *string {
	return s.Value
}

func (s *QueryConversationDetailInfoResponseBodyDataVariables) SetId(v string) *QueryConversationDetailInfoResponseBodyDataVariables {
	s.Id = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyDataVariables) SetKey(v string) *QueryConversationDetailInfoResponseBodyDataVariables {
	s.Key = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyDataVariables) SetName(v string) *QueryConversationDetailInfoResponseBodyDataVariables {
	s.Name = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyDataVariables) SetRequired(v bool) *QueryConversationDetailInfoResponseBodyDataVariables {
	s.Required = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyDataVariables) SetSource(v string) *QueryConversationDetailInfoResponseBodyDataVariables {
	s.Source = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyDataVariables) SetValue(v string) *QueryConversationDetailInfoResponseBodyDataVariables {
	s.Value = &v
	return s
}

func (s *QueryConversationDetailInfoResponseBodyDataVariables) Validate() error {
	return dara.Validate(s)
}
