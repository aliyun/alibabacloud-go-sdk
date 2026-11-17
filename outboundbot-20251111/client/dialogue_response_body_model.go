// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDialogueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DialogueResponseBody
	GetCode() *string
	SetData(v *DialogueResponseBodyData) *DialogueResponseBody
	GetData() *DialogueResponseBodyData
	SetHttpStatusCode(v int32) *DialogueResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DialogueResponseBody
	GetMessage() *string
	SetParams(v []*string) *DialogueResponseBody
	GetParams() []*string
	SetRequestId(v string) *DialogueResponseBody
	GetRequestId() *string
}

type DialogueResponseBody struct {
	// The business processing status code. OK indicates success. When an exception occurs, the corresponding business error code is returned. The error types and mapping rules follow the current API configuration. No error codes are added or modified in this operation.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The robot reply, round ID, time, and control instructions carried by the current SSE event. A single round may contain multiple events.
	//
	// example:
	//
	// {"Answer":"Sure, let me introduce the handling process to you.","StreamEnd":true,"SessionId":"text-11111111-2222-4333-8444-555555555555","StreamId":"1","StartTime":1794560400000,"EndTime":1794560400123,"ControlParamsList":[{"Type":"Voice"}],"Extras":{"DebugInfo":"{}"}}
	Data *DialogueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code recorded in the business response body. A normal return is 200. To determine whether the business operation is successful, you must also check Code. You cannot determine that the call is successful solely based on the presence of this field.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The business response or error description. It is usually empty or not returned upon success. A fixed text cannot be required upon success. Upon failure, it provides an error description to assist in troubleshooting. Combine it with Code and Params to understand the cause of the error. The example shows an error message when the instance does not exist (corresponding to the Code NotExists.InstanceId) and is for format reference only.
	//
	// example:
	//
	// Instance instance-example-001 does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The parameter list used for the localized display of error messages. It is not a scenario variable or business result. A normal successful response is usually an empty array. When an error occurs, parameters corresponding to the error template may be returned.
	//
	// example:
	//
	// []
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The ID used to associate and troubleshoot this POP request. The value in the response corresponds to the current request. System request parameters are injected by the gateway, and the caller does not need to fill in the system fields.
	//
	// example:
	//
	// AAAAAAAA-BBBB-4CCC-8DDD-EEEEEEEEEEEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DialogueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DialogueResponseBody) GoString() string {
	return s.String()
}

func (s *DialogueResponseBody) GetCode() *string {
	return s.Code
}

func (s *DialogueResponseBody) GetData() *DialogueResponseBodyData {
	return s.Data
}

func (s *DialogueResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DialogueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DialogueResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DialogueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DialogueResponseBody) SetCode(v string) *DialogueResponseBody {
	s.Code = &v
	return s
}

func (s *DialogueResponseBody) SetData(v *DialogueResponseBodyData) *DialogueResponseBody {
	s.Data = v
	return s
}

func (s *DialogueResponseBody) SetHttpStatusCode(v int32) *DialogueResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DialogueResponseBody) SetMessage(v string) *DialogueResponseBody {
	s.Message = &v
	return s
}

func (s *DialogueResponseBody) SetParams(v []*string) *DialogueResponseBody {
	s.Params = v
	return s
}

func (s *DialogueResponseBody) SetRequestId(v string) *DialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DialogueResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DialogueResponseBodyData struct {
	// The robot reply text in the current streaming event. It may be returned across multiple SSE events, or it may be empty in the end event or debug event. Do not equate the absence of text with session failure. Whether the content is incremental or cumulative is determined by the output convention of the conversation engine. The current general callback directly forwards the engine content.
	//
	// example:
	//
	// Sure, let me introduce the handling process to you
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// The robot control instructions attached to this reply. The current POP structure only expands the instruction Type and does not add other instruction fields in this document. When no control instruction is generated, the text service supplements a Voice type instruction.
	//
	// example:
	//
	// [{"Type":"Voice"}]
	ControlParamsList []*DialogueResponseBodyDataControlParamsList `json:"ControlParamsList,omitempty" xml:"ControlParamsList,omitempty" type:"Repeated"`
	// The UNIX timestamp of the current streaming reply callback event, in milliseconds. It is generated based on the current callback time. It may also be returned when StreamEnd is false. It should not be directly interpreted as the end time of the entire session or call.
	//
	// example:
	//
	// 1794560400123
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The extension information object attached to the robot reply, which is different from the JSON event array of the request parameter Extras. The current POP expands the DebugInfo field. Whether it is actually returned depends on the engine used and the current round of reply.
	//
	// example:
	//
	// {"DebugInfo":"{}"}
	Extras *DialogueResponseBodyDataExtras `json:"Extras,omitempty" xml:"Extras,omitempty" type:"Struct"`
	// The session ID used to uniquely associate this interaction. For a text session, it is the text- identifier returned when the session is initiated. For web voice, it is the ws- identifier. For real number outbound call tests, it is the test- identifier. Subsequent calls must use the actual returned value.
	//
	// example:
	//
	// text-11111111-2222-4333-8444-555555555555
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The UNIX timestamp recorded when the streaming reply callback is entered for the first time in the current round, in milliseconds. Subsequent events in the current round use this time. It is not the session creation time, nor is it an update time in the general sense.
	//
	// example:
	//
	// 1794560400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates whether the streaming output of the current reply round has ended. A value of true indicates that the streaming output has ended and the service has completed the corresponding SSE response. A value of false indicates that there may still be subsequent events in the current round. This field does not indicate that the entire session or voice call has ended.
	//
	// example:
	//
	// true
	StreamEnd *bool `json:"StreamEnd,omitempty" xml:"StreamEnd,omitempty"`
	// The ID of the conversation round to which the current robot reply belongs. A single round of reply may contain multiple streaming events. The text session context starts from 0 and increments upon subsequent Dialogue calls. The actual value returned by the engine shall prevail.
	//
	// example:
	//
	// 1
	StreamId *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
}

func (s DialogueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DialogueResponseBodyData) GoString() string {
	return s.String()
}

func (s *DialogueResponseBodyData) GetAnswer() *string {
	return s.Answer
}

func (s *DialogueResponseBodyData) GetControlParamsList() []*DialogueResponseBodyDataControlParamsList {
	return s.ControlParamsList
}

func (s *DialogueResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DialogueResponseBodyData) GetExtras() *DialogueResponseBodyDataExtras {
	return s.Extras
}

func (s *DialogueResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *DialogueResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DialogueResponseBodyData) GetStreamEnd() *bool {
	return s.StreamEnd
}

func (s *DialogueResponseBodyData) GetStreamId() *string {
	return s.StreamId
}

func (s *DialogueResponseBodyData) SetAnswer(v string) *DialogueResponseBodyData {
	s.Answer = &v
	return s
}

func (s *DialogueResponseBodyData) SetControlParamsList(v []*DialogueResponseBodyDataControlParamsList) *DialogueResponseBodyData {
	s.ControlParamsList = v
	return s
}

func (s *DialogueResponseBodyData) SetEndTime(v int64) *DialogueResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DialogueResponseBodyData) SetExtras(v *DialogueResponseBodyDataExtras) *DialogueResponseBodyData {
	s.Extras = v
	return s
}

func (s *DialogueResponseBodyData) SetSessionId(v string) *DialogueResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *DialogueResponseBodyData) SetStartTime(v int64) *DialogueResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DialogueResponseBodyData) SetStreamEnd(v bool) *DialogueResponseBodyData {
	s.StreamEnd = &v
	return s
}

func (s *DialogueResponseBodyData) SetStreamId(v string) *DialogueResponseBodyData {
	s.StreamId = &v
	return s
}

func (s *DialogueResponseBodyData) Validate() error {
	if s.ControlParamsList != nil {
		for _, item := range s.ControlParamsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Extras != nil {
		if err := s.Extras.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DialogueResponseBodyDataControlParamsList struct {
	// The values declared in the control instruction type include Voice (voice reply), HangUp (hang up), Transfer (transfer), GatherDtmf (collect key presses), and SendSms (send SMS). This update supplements the previously omitted SendSms without changing the backend structure or runtime enumeration constraints.
	//
	// example:
	//
	// Voice
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DialogueResponseBodyDataControlParamsList) String() string {
	return dara.Prettify(s)
}

func (s DialogueResponseBodyDataControlParamsList) GoString() string {
	return s.String()
}

func (s *DialogueResponseBodyDataControlParamsList) GetType() *string {
	return s.Type
}

func (s *DialogueResponseBodyDataControlParamsList) SetType(v string) *DialogueResponseBodyDataControlParamsList {
	s.Type = &v
	return s
}

func (s *DialogueResponseBodyDataControlParamsList) Validate() error {
	return dara.Validate(s)
}

type DialogueResponseBodyDataExtras struct {
	// The debug information returned by the conversation engine, which is a JSON string. It is returned only when the engine provides and obtains the debug information, and may be returned with supplementary events after the reply ends. The internal structure is determined by the engine. The example empty object string does not declare any fixed debug fields.
	//
	// example:
	//
	// {}
	DebugInfo *string `json:"DebugInfo,omitempty" xml:"DebugInfo,omitempty"`
}

func (s DialogueResponseBodyDataExtras) String() string {
	return dara.Prettify(s)
}

func (s DialogueResponseBodyDataExtras) GoString() string {
	return s.String()
}

func (s *DialogueResponseBodyDataExtras) GetDebugInfo() *string {
	return s.DebugInfo
}

func (s *DialogueResponseBodyDataExtras) SetDebugInfo(v string) *DialogueResponseBodyDataExtras {
	s.DebugInfo = &v
	return s
}

func (s *DialogueResponseBodyDataExtras) Validate() error {
	return dara.Validate(s)
}
