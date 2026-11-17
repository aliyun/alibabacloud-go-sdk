// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrialDialogueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrialDialogueResponseBody
	GetCode() *string
	SetData(v *TrialDialogueResponseBodyData) *TrialDialogueResponseBody
	GetData() *TrialDialogueResponseBodyData
	SetHttpStatusCode(v int32) *TrialDialogueResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *TrialDialogueResponseBody
	GetMessage() *string
	SetParams(v []*string) *TrialDialogueResponseBody
	GetParams() []*string
	SetRequestId(v string) *TrialDialogueResponseBody
	GetRequestId() *string
}

type TrialDialogueResponseBody struct {
	// The business processing status code. OK indicates success. When an exception occurs, the corresponding business error code is returned. The error types and mapping rules follow the current API configuration, and no error codes are added or modified in this update.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The robot reply, turn ID, time, and control commands carried by the current SSE event. A single turn may have multiple events.
	//
	// example:
	//
	// {"Answer":"Sure, let me introduce the handling process to you.","StreamEnd":true,"SessionId":"text-11111111-2222-4333-8444-555555555555","StreamId":"1","StartTime":1794560400000,"EndTime":1794560400123,"ControlParamsList":[{"Type":"Voice"}],"Extras":{"DebugInfo":"{}"}}
	Data *TrialDialogueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code recorded in the business response body. A normal return is 200. To determine whether the business is successful, you should also check Code. You cannot determine that the call is successful solely based on the presence of this field.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The business response or error description. It is usually empty or not returned upon success, and you cannot require a fixed text to be returned upon success. Upon failure, it is an error description used to assist troubleshooting, and should be combined with Code and Params to understand the cause of the error. The example is an error message when the experience link visitor token expires (corresponding to Code Expired.VisitorToken), and is for format reference only.
	//
	// example:
	//
	// VisitorToken: example-visitor-token is expired.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of parameters used for the localized display of error messages, which are not scenario variables or business results. A normal successful response is usually an empty array. When an error occurs, parameters corresponding to the error template may be returned.
	//
	// example:
	//
	// []
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The ID used to associate and troubleshoot this POP request. The value in the response corresponds to this request. System request parameters are injected by the gateway, and the caller does not need to fill in system fields.
	//
	// example:
	//
	// AAAAAAAA-BBBB-4CCC-8DDD-EEEEEEEEEEEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TrialDialogueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrialDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *TrialDialogueResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrialDialogueResponseBody) GetData() *TrialDialogueResponseBodyData {
	return s.Data
}

func (s *TrialDialogueResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *TrialDialogueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrialDialogueResponseBody) GetParams() []*string {
	return s.Params
}

func (s *TrialDialogueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrialDialogueResponseBody) SetCode(v string) *TrialDialogueResponseBody {
	s.Code = &v
	return s
}

func (s *TrialDialogueResponseBody) SetData(v *TrialDialogueResponseBodyData) *TrialDialogueResponseBody {
	s.Data = v
	return s
}

func (s *TrialDialogueResponseBody) SetHttpStatusCode(v int32) *TrialDialogueResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TrialDialogueResponseBody) SetMessage(v string) *TrialDialogueResponseBody {
	s.Message = &v
	return s
}

func (s *TrialDialogueResponseBody) SetParams(v []*string) *TrialDialogueResponseBody {
	s.Params = v
	return s
}

func (s *TrialDialogueResponseBody) SetRequestId(v string) *TrialDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrialDialogueResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TrialDialogueResponseBodyData struct {
	// The robot reply text in the current streaming event. It may be returned across multiple SSE events, or it may be empty in the end event or debug event. Do not equate missing text with session failure. Whether the content is incremental or cumulative is determined by the output convention of the dialogue engine. The current general callback directly forwards the engine content.
	//
	// example:
	//
	// Sure, let me introduce the handling process to you
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// The robot control commands attached to this reply. The current POP structure only expands the command Type and does not add other command fields in this document. When no control commands are generated, the text service supplements a Voice type command.
	//
	// example:
	//
	// [{"Type":"Voice"}]
	ControlParamsList []*TrialDialogueResponseBodyDataControlParamsList `json:"ControlParamsList,omitempty" xml:"ControlParamsList,omitempty" type:"Repeated"`
	// The UNIX timestamp of the current streaming reply callback event, in milliseconds. It is generated based on the current callback time. It may also be returned when StreamEnd=false, and should not be directly interpreted as the end time of the entire conversation or call.
	//
	// example:
	//
	// 1794560400123
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The extended information object attached to the robot reply, which is different from the JSON event array of the request parameter Extras. The current POP expands the DebugInfo field. Whether it is actually returned depends on the engine used and the current turn of reply.
	//
	// example:
	//
	// {"DebugInfo":"{}"}
	Extras *TrialDialogueResponseBodyDataExtras `json:"Extras,omitempty" xml:"Extras,omitempty" type:"Struct"`
	// The session ID used to uniquely associate this interaction. For a text conversation, it is the text- identifier returned when the conversation is started. For web voice, it is the ws- identifier. For real number outbound call tests, it is the test- identifier. Subsequent calls must use the actual return value.
	//
	// example:
	//
	// text-11111111-2222-4333-8444-555555555555
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The UNIX timestamp recorded when the streaming reply callback is entered for the first time in this turn, in milliseconds. Subsequent events in this turn use this time. It is not the session creation time, nor is it an update time in the general sense.
	//
	// example:
	//
	// 1794560400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates whether the streaming output for the current reply turn has ended. A value of true indicates that the streaming output has ended and the service has completed the corresponding SSE response. A value of false indicates that there may still be subsequent events in this turn. This field does not indicate that the entire conversation or voice call has ended.
	//
	// example:
	//
	// true
	StreamEnd *bool `json:"StreamEnd,omitempty" xml:"StreamEnd,omitempty"`
	// The ID of the conversation turn to which this robot reply belongs. The same turn of reply can contain multiple streaming events. The text conversation context starts from 0 and increments during subsequent Dialogue calls. The actual return value of the engine shall prevail.
	//
	// example:
	//
	// 1
	StreamId *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
}

func (s TrialDialogueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TrialDialogueResponseBodyData) GoString() string {
	return s.String()
}

func (s *TrialDialogueResponseBodyData) GetAnswer() *string {
	return s.Answer
}

func (s *TrialDialogueResponseBodyData) GetControlParamsList() []*TrialDialogueResponseBodyDataControlParamsList {
	return s.ControlParamsList
}

func (s *TrialDialogueResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *TrialDialogueResponseBodyData) GetExtras() *TrialDialogueResponseBodyDataExtras {
	return s.Extras
}

func (s *TrialDialogueResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *TrialDialogueResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *TrialDialogueResponseBodyData) GetStreamEnd() *bool {
	return s.StreamEnd
}

func (s *TrialDialogueResponseBodyData) GetStreamId() *string {
	return s.StreamId
}

func (s *TrialDialogueResponseBodyData) SetAnswer(v string) *TrialDialogueResponseBodyData {
	s.Answer = &v
	return s
}

func (s *TrialDialogueResponseBodyData) SetControlParamsList(v []*TrialDialogueResponseBodyDataControlParamsList) *TrialDialogueResponseBodyData {
	s.ControlParamsList = v
	return s
}

func (s *TrialDialogueResponseBodyData) SetEndTime(v int64) *TrialDialogueResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *TrialDialogueResponseBodyData) SetExtras(v *TrialDialogueResponseBodyDataExtras) *TrialDialogueResponseBodyData {
	s.Extras = v
	return s
}

func (s *TrialDialogueResponseBodyData) SetSessionId(v string) *TrialDialogueResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *TrialDialogueResponseBodyData) SetStartTime(v int64) *TrialDialogueResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *TrialDialogueResponseBodyData) SetStreamEnd(v bool) *TrialDialogueResponseBodyData {
	s.StreamEnd = &v
	return s
}

func (s *TrialDialogueResponseBodyData) SetStreamId(v string) *TrialDialogueResponseBodyData {
	s.StreamId = &v
	return s
}

func (s *TrialDialogueResponseBodyData) Validate() error {
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

type TrialDialogueResponseBodyDataControlParamsList struct {
	// The values declared in the control command type include Voice (voice reply), HangUp (hang up), Transfer (transfer), GatherDtmf (collect key presses), and SendSms (send SMS). This update supplements SendSms, which was omitted in the original description, without changing the backend structure or runtime enumeration constraints.
	//
	// example:
	//
	// Voice
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s TrialDialogueResponseBodyDataControlParamsList) String() string {
	return dara.Prettify(s)
}

func (s TrialDialogueResponseBodyDataControlParamsList) GoString() string {
	return s.String()
}

func (s *TrialDialogueResponseBodyDataControlParamsList) GetType() *string {
	return s.Type
}

func (s *TrialDialogueResponseBodyDataControlParamsList) SetType(v string) *TrialDialogueResponseBodyDataControlParamsList {
	s.Type = &v
	return s
}

func (s *TrialDialogueResponseBodyDataControlParamsList) Validate() error {
	return dara.Validate(s)
}

type TrialDialogueResponseBodyDataExtras struct {
	// The debug information returned by the dialogue engine as a JSON string. It is returned only when the engine provides the debug information, and may be returned with supplementary events after the reply ends. The internal structure is determined by the engine. The example empty object string does not declare any fixed debug fields.
	//
	// example:
	//
	// {}
	DebugInfo *string `json:"DebugInfo,omitempty" xml:"DebugInfo,omitempty"`
}

func (s TrialDialogueResponseBodyDataExtras) String() string {
	return dara.Prettify(s)
}

func (s TrialDialogueResponseBodyDataExtras) GoString() string {
	return s.String()
}

func (s *TrialDialogueResponseBodyDataExtras) GetDebugInfo() *string {
	return s.DebugInfo
}

func (s *TrialDialogueResponseBodyDataExtras) SetDebugInfo(v string) *TrialDialogueResponseBodyDataExtras {
	s.DebugInfo = &v
	return s
}

func (s *TrialDialogueResponseBodyDataExtras) Validate() error {
	return dara.Validate(s)
}
