// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBeginTrialSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BeginTrialSessionResponseBody
	GetCode() *string
	SetData(v *BeginTrialSessionResponseBodyData) *BeginTrialSessionResponseBody
	GetData() *BeginTrialSessionResponseBodyData
	SetHttpStatusCode(v int32) *BeginTrialSessionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BeginTrialSessionResponseBody
	GetMessage() *string
	SetParams(v []*string) *BeginTrialSessionResponseBody
	GetParams() []*string
	SetRequestId(v string) *BeginTrialSessionResponseBody
	GetRequestId() *string
}

type BeginTrialSessionResponseBody struct {
	// The business processing status code. OK indicates success. If an exception occurs, the corresponding business error code is returned. Error types and mapping rules follow the current API configuration. No error codes are added or modified in this operation.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The robot reply, round ID, time, and control commands carried by the current SSE event. A single round may contain multiple events.
	//
	// example:
	//
	// {"Answer":"Hello, how can I help you?","StreamEnd":true,"SessionId":"text-11111111-2222-4333-8444-555555555555","StreamId":"0","StartTime":1794560400000,"EndTime":1794560400123,"ControlParamsList":[{"Type":"Voice"}],"Extras":{"DebugInfo":"{}"}}
	Data *BeginTrialSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code recorded in the business response body. A normal return value is 200. To determine whether the business operation is successful, you must also check Code. You cannot determine the success of a call solely based on the presence of this field.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The business response or error description. This field is usually empty or not returned upon success, and no fixed text is returned. Upon failure, it provides an error description to assist in troubleshooting. Combine this field with Code and Params to understand the cause of the error. The example shows an error message when the trial link visitor token expires (corresponding to the Expired.VisitorToken code) and is for format reference only.
	//
	// example:
	//
	// VisitorToken: example-visitor-token is expired.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The parameter list used for the localized display of error messages. It is not a scenario variable or business result. A normal successful response is usually an empty array. If an error occurs, parameters corresponding to the error template may be returned.
	//
	// example:
	//
	// []
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The ID used to associate and troubleshoot this POP request. The value in the response corresponds to the current request. System request parameters are injected by the gateway, and the caller does not need to fill in system fields.
	//
	// example:
	//
	// AAAAAAAA-BBBB-4CCC-8DDD-EEEEEEEEEEEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BeginTrialSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BeginTrialSessionResponseBody) GoString() string {
	return s.String()
}

func (s *BeginTrialSessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *BeginTrialSessionResponseBody) GetData() *BeginTrialSessionResponseBodyData {
	return s.Data
}

func (s *BeginTrialSessionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BeginTrialSessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BeginTrialSessionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *BeginTrialSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BeginTrialSessionResponseBody) SetCode(v string) *BeginTrialSessionResponseBody {
	s.Code = &v
	return s
}

func (s *BeginTrialSessionResponseBody) SetData(v *BeginTrialSessionResponseBodyData) *BeginTrialSessionResponseBody {
	s.Data = v
	return s
}

func (s *BeginTrialSessionResponseBody) SetHttpStatusCode(v int32) *BeginTrialSessionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BeginTrialSessionResponseBody) SetMessage(v string) *BeginTrialSessionResponseBody {
	s.Message = &v
	return s
}

func (s *BeginTrialSessionResponseBody) SetParams(v []*string) *BeginTrialSessionResponseBody {
	s.Params = v
	return s
}

func (s *BeginTrialSessionResponseBody) SetRequestId(v string) *BeginTrialSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *BeginTrialSessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BeginTrialSessionResponseBodyData struct {
	// The robot reply text in the current streaming event. It may be returned across multiple SSE events, or it may be empty in an end event or debug event. Do not equate missing text with session failure. Whether the content is incremental or cumulative is determined by the output conventions of the dialogue engine. The current general callback directly forwards the engine content.
	//
	// example:
	//
	// Hello, how can I help you?
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// The robot control commands attached to this reply. The current POP structure only expands the command Type and does not add other command fields in this document. If no control command is generated, the text service supplements a Voice type command.
	//
	// example:
	//
	// [{"Type":"Voice"}]
	ControlParamsList []*BeginTrialSessionResponseBodyDataControlParamsList `json:"ControlParamsList,omitempty" xml:"ControlParamsList,omitempty" type:"Repeated"`
	// The UNIX timestamp of the current streaming reply callback event, in milliseconds. It is generated based on the current callback time. It may also be returned when StreamEnd=false. Do not directly interpret it as the end time of the entire session or call.
	//
	// example:
	//
	// 1794560400123
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The extended information object attached to the robot reply, which is different from the JSON event array of the request parameter Extras. The current POP expands the DebugInfo field. Whether it is actually returned depends on the engine used and the current round of reply.
	//
	// example:
	//
	// {"DebugInfo":"{}"}
	Extras *BeginTrialSessionResponseBodyDataExtras `json:"Extras,omitempty" xml:"Extras,omitempty" type:"Struct"`
	// The session ID used to uniquely associate this interaction. For a text session, it is the text- ID returned when the session is started. For web voice, it is the ws- ID. For a real number outbound call test, it is the test- ID. Subsequent calls must use the actual return value.
	//
	// example:
	//
	// text-11111111-2222-4333-8444-555555555555
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The UNIX timestamp recorded when the streaming reply callback is entered for the first time in the current round, in milliseconds. Subsequent events in the current round use this time. It is not the session creation time or a general update time.
	//
	// example:
	//
	// 1794560400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Specifies whether the streaming output of the current reply round has ended. Valid values:
	//
	// - true: The streaming output has ended, and the service has completed the corresponding SSE response.
	//
	// - false: Subsequent events may still occur in the current round.
	//
	// This field does not indicate that the entire session or voice call has ended.
	//
	// example:
	//
	// true
	StreamEnd *bool `json:"StreamEnd,omitempty" xml:"StreamEnd,omitempty"`
	// The ID of the dialogue round to which this robot reply belongs. A single round of reply may contain multiple streaming events. The text session context starts from 0 and increments during subsequent Dialogue calls. The actual return value from the engine shall prevail.
	//
	// example:
	//
	// 0
	StreamId *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
}

func (s BeginTrialSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BeginTrialSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *BeginTrialSessionResponseBodyData) GetAnswer() *string {
	return s.Answer
}

func (s *BeginTrialSessionResponseBodyData) GetControlParamsList() []*BeginTrialSessionResponseBodyDataControlParamsList {
	return s.ControlParamsList
}

func (s *BeginTrialSessionResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *BeginTrialSessionResponseBodyData) GetExtras() *BeginTrialSessionResponseBodyDataExtras {
	return s.Extras
}

func (s *BeginTrialSessionResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *BeginTrialSessionResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *BeginTrialSessionResponseBodyData) GetStreamEnd() *bool {
	return s.StreamEnd
}

func (s *BeginTrialSessionResponseBodyData) GetStreamId() *string {
	return s.StreamId
}

func (s *BeginTrialSessionResponseBodyData) SetAnswer(v string) *BeginTrialSessionResponseBodyData {
	s.Answer = &v
	return s
}

func (s *BeginTrialSessionResponseBodyData) SetControlParamsList(v []*BeginTrialSessionResponseBodyDataControlParamsList) *BeginTrialSessionResponseBodyData {
	s.ControlParamsList = v
	return s
}

func (s *BeginTrialSessionResponseBodyData) SetEndTime(v int64) *BeginTrialSessionResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *BeginTrialSessionResponseBodyData) SetExtras(v *BeginTrialSessionResponseBodyDataExtras) *BeginTrialSessionResponseBodyData {
	s.Extras = v
	return s
}

func (s *BeginTrialSessionResponseBodyData) SetSessionId(v string) *BeginTrialSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *BeginTrialSessionResponseBodyData) SetStartTime(v int64) *BeginTrialSessionResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *BeginTrialSessionResponseBodyData) SetStreamEnd(v bool) *BeginTrialSessionResponseBodyData {
	s.StreamEnd = &v
	return s
}

func (s *BeginTrialSessionResponseBodyData) SetStreamId(v string) *BeginTrialSessionResponseBodyData {
	s.StreamId = &v
	return s
}

func (s *BeginTrialSessionResponseBodyData) Validate() error {
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

type BeginTrialSessionResponseBodyDataControlParamsList struct {
	// The values declared in the control command type include Voice (voice reply), HangUp (hang up), Transfer (transfer), GatherDtmf (collect DTMF tones), and SendSms (send SMS message). SendSms, which was omitted in the original description, is added in this update without changing the backend structure or runtime enumeration constraints.
	//
	// example:
	//
	// Voice
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BeginTrialSessionResponseBodyDataControlParamsList) String() string {
	return dara.Prettify(s)
}

func (s BeginTrialSessionResponseBodyDataControlParamsList) GoString() string {
	return s.String()
}

func (s *BeginTrialSessionResponseBodyDataControlParamsList) GetType() *string {
	return s.Type
}

func (s *BeginTrialSessionResponseBodyDataControlParamsList) SetType(v string) *BeginTrialSessionResponseBodyDataControlParamsList {
	s.Type = &v
	return s
}

func (s *BeginTrialSessionResponseBodyDataControlParamsList) Validate() error {
	return dara.Validate(s)
}

type BeginTrialSessionResponseBodyDataExtras struct {
	// The debug information returned by the dialogue engine, which is a JSON string. It is returned only when the engine provides and obtains debug information, and may be returned with supplementary events after the reply ends. The internal structure is determined by the engine. The example empty object string does not declare any fixed debug fields.
	//
	// example:
	//
	// {}
	DebugInfo *string `json:"DebugInfo,omitempty" xml:"DebugInfo,omitempty"`
}

func (s BeginTrialSessionResponseBodyDataExtras) String() string {
	return dara.Prettify(s)
}

func (s BeginTrialSessionResponseBodyDataExtras) GoString() string {
	return s.String()
}

func (s *BeginTrialSessionResponseBodyDataExtras) GetDebugInfo() *string {
	return s.DebugInfo
}

func (s *BeginTrialSessionResponseBodyDataExtras) SetDebugInfo(v string) *BeginTrialSessionResponseBodyDataExtras {
	s.DebugInfo = &v
	return s
}

func (s *BeginTrialSessionResponseBodyDataExtras) Validate() error {
	return dara.Validate(s)
}
