// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallDetailByTaskIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCallDetailByTaskIdResponseBody
	GetCode() *string
	SetData(v string) *QueryCallDetailByTaskIdResponseBody
	GetData() *string
	SetMessage(v string) *QueryCallDetailByTaskIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCallDetailByTaskIdResponseBody
	GetRequestId() *string
}

type QueryCallDetailByTaskIdResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/112502.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The call details of the outbound robocall task, in the JSON format.
	//
	// 	- **startDate**: the time when the call was answered.
	//
	// 	- **stateDesc**: the reason why the call was hung up. If the status code of early media was returned, this parameter indicates the reason why the status code of early media was used.
	//
	// 	- **statusCode**: the status code.
	//
	// 	- **EndDate**: the time when the call was ended.
	//
	// 	- **calleeShowNumber**: the calling number displayed to the called party.
	//
	// 	- **callee**: the called number.
	//
	// 	- **duration**: the billing duration.
	//
	// 	- **gmtCreate**: the time when the outbound robocall task was created.
	//
	// 	- **hangupDirection**: the party who hung up.
	//
	// 	- **tags**: the call tags.
	//
	// 	- **dialogCount**: the number of conversation rounds in the call.
	//
	// 	- **sureCount**: the number of times that the robocall task was acknowledged.
	//
	// 	- **denyCount**: the number of times that the robocall task was denied.
	//
	// 	- **rejectCount**: the number of times that the robocall task was rejected.
	//
	// 	- **customCount**: the number of times that the robocall task was customized.
	//
	// 	- **knowledgeCount**: the number of times that the knowledge base was queried.
	//
	// 	- **recordFile**: the download URL of the recording file. The URL is valid only for 48 hours. The recording file must be downloaded in time.
	//
	// 	- **callId**: the call ID.
	//
	// 	- **recordStatus**: indicates whether a recording file was available. Valid values:
	//
	//     	- 1: A recording file was available.
	//
	//     	- 2: No recording file was available.
	//
	// 	- **knowledgeCommonCount**: the number of call failures caused by the common issues in the knowledge base.
	//
	// 	- **knowledgeBusinessCount**: the number of call failures caused by the business issues in the knowledge base.
	//
	// 	- **callee**: the called number.
	//
	// 	- **dialogDetail**: the conversation details. The value is a JSON array that contains the following parameters:
	//
	//     	- **role**: the role who spoke.
	//
	//     	- **content**: the content of the speech.
	//
	//     	- **time**: the start time of the speech.
	//
	// > The preceding parameters are for reference only. The actually returned parameters prevail.
	//
	// example:
	//
	// {"rejectCount":0,"dialogCount":3,"tags":"","startDate":"2019-03-27 10:34:54","gmtCreate":"2019-03-27 10:34:40","sureCount":0,"state":"200000","recordFile":"http://alicom-fc-record-biz.cn-hangzhou.oss.aliyun-inc.com/Freeswitch_RU_115987800002_02c3554f-ea24-422d-b1de-e671f455f21a_record.wav?OSSAccessKeyId=bypFNbGJVk73****&Signature=VWHOX%2FFhvvtSkxfMTw%2F5fdJUQuk%3D&Expires=1554382725","defaultCount":0,"endDate":"2019-03-27 10:35:09","calleeShowNumber":"1390000****","customCount":0,"callId":"1390000****","knowledgeCount":0,"recordStatus":1,"denyCount":0,"duration":16,"knowledgeCommonCount":0,"callee":"1390000****","knowledgeBusinessCount":0,"hangupDirection":1}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D86B61A8-F2EE-4E4C-9F05-08A4676FFD89
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCallDetailByTaskIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCallDetailByTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallDetailByTaskIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCallDetailByTaskIdResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryCallDetailByTaskIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCallDetailByTaskIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCallDetailByTaskIdResponseBody) SetCode(v string) *QueryCallDetailByTaskIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCallDetailByTaskIdResponseBody) SetData(v string) *QueryCallDetailByTaskIdResponseBody {
	s.Data = &v
	return s
}

func (s *QueryCallDetailByTaskIdResponseBody) SetMessage(v string) *QueryCallDetailByTaskIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCallDetailByTaskIdResponseBody) SetRequestId(v string) *QueryCallDetailByTaskIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCallDetailByTaskIdResponseBody) Validate() error {
	return dara.Validate(s)
}
