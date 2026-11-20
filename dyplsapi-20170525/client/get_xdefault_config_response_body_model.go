// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetXDefaultConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetXDefaultConfigResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetXDefaultConfigResponseBody
	GetCode() *string
	SetData(v *GetXDefaultConfigResponseBodyData) *GetXDefaultConfigResponseBody
	GetData() *GetXDefaultConfigResponseBodyData
	SetMessage(v string) *GetXDefaultConfigResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetXDefaultConfigResponseBody
	GetSuccess() *bool
}

type GetXDefaultConfigResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code. Valid values:
	//
	// - 0000: Success.
	//
	// - Other values: Failure.
	//
	// example:
	//
	// 0000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The default configurations of the work number.
	//
	// example:
	//
	// {"CallAbility":"0","GNFlag":"00","SmsAbility":"0","SmsSignMode":"0"}
	Data *GetXDefaultConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetXDefaultConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetXDefaultConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetXDefaultConfigResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetXDefaultConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetXDefaultConfigResponseBody) GetData() *GetXDefaultConfigResponseBodyData {
	return s.Data
}

func (s *GetXDefaultConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetXDefaultConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetXDefaultConfigResponseBody) SetAccessDeniedDetail(v string) *GetXDefaultConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetXDefaultConfigResponseBody) SetCode(v string) *GetXDefaultConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetXDefaultConfigResponseBody) SetData(v *GetXDefaultConfigResponseBodyData) *GetXDefaultConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetXDefaultConfigResponseBody) SetMessage(v string) *GetXDefaultConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetXDefaultConfigResponseBody) SetSuccess(v bool) *GetXDefaultConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetXDefaultConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetXDefaultConfigResponseBodyData struct {
	// The status of the call capability. Valid values: 0: disabled. 1: enabled.
	//
	// example:
	//
	// 0
	CallAbility *string `json:"CallAbility,omitempty" xml:"CallAbility,omitempty"`
	// Specifies whether to pass through the caller ID to display the real caller. Valid values: 00: non-passthrough. The work number is displayed when calls are made to each other. 10: passthrough. When Customer A is the caller, the caller ID displayed to Employee B is the number of Customer A. When Employee B is the caller, the caller ID displayed to Customer A is the work number. Default value: 00.
	//
	// example:
	//
	// 00
	GNFlag *string `json:"GNFlag,omitempty" xml:"GNFlag,omitempty"`
	// The enterprise business card rule control parameters.
	ReachJson []*GetXDefaultConfigResponseBodyDataReachJson `json:"ReachJson,omitempty" xml:"ReachJson,omitempty" type:"Repeated"`
	// The sequential ringing control parameters.
	SequenceCall []*GetXDefaultConfigResponseBodyDataSequenceCall `json:"SequenceCall,omitempty" xml:"SequenceCall,omitempty" type:"Repeated"`
	// The end time of sequential ringing. Format: HH:mm:ss, such as 18:00:00.
	//
	// example:
	//
	// 09:00:00
	SequenceEndTime *string `json:"SequenceEndTime,omitempty" xml:"SequenceEndTime,omitempty"`
	// The start time of sequential ringing. Format: HH:mm:ss, such as 09:00:00.
	//
	// example:
	//
	// 09:00:00
	SequenceStartTime *string `json:"SequenceStartTime,omitempty" xml:"SequenceStartTime,omitempty"`
	// The status of the SMS feature. Valid values: 0: disabled. 1: enabled.
	//
	// example:
	//
	// 0
	SmsAbility *string `json:"SmsAbility,omitempty" xml:"SmsAbility,omitempty"`
	// Specifies whether to pass through the caller ID to display the real caller. Valid values: 00: non-passthrough. The work number is displayed when calls are made to each other. 10: passthrough. When Customer A is the caller, the caller ID displayed to Employee B is the number of Customer A. When Employee B is the caller, the caller ID displayed to Customer A is the work number. Default value: 00.
	//
	// example:
	//
	// 0
	SmsSignMode *string `json:"SmsSignMode,omitempty" xml:"SmsSignMode,omitempty"`
}

func (s GetXDefaultConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetXDefaultConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetXDefaultConfigResponseBodyData) GetCallAbility() *string {
	return s.CallAbility
}

func (s *GetXDefaultConfigResponseBodyData) GetGNFlag() *string {
	return s.GNFlag
}

func (s *GetXDefaultConfigResponseBodyData) GetReachJson() []*GetXDefaultConfigResponseBodyDataReachJson {
	return s.ReachJson
}

func (s *GetXDefaultConfigResponseBodyData) GetSequenceCall() []*GetXDefaultConfigResponseBodyDataSequenceCall {
	return s.SequenceCall
}

func (s *GetXDefaultConfigResponseBodyData) GetSequenceEndTime() *string {
	return s.SequenceEndTime
}

func (s *GetXDefaultConfigResponseBodyData) GetSequenceStartTime() *string {
	return s.SequenceStartTime
}

func (s *GetXDefaultConfigResponseBodyData) GetSmsAbility() *string {
	return s.SmsAbility
}

func (s *GetXDefaultConfigResponseBodyData) GetSmsSignMode() *string {
	return s.SmsSignMode
}

func (s *GetXDefaultConfigResponseBodyData) SetCallAbility(v string) *GetXDefaultConfigResponseBodyData {
	s.CallAbility = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetGNFlag(v string) *GetXDefaultConfigResponseBodyData {
	s.GNFlag = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetReachJson(v []*GetXDefaultConfigResponseBodyDataReachJson) *GetXDefaultConfigResponseBodyData {
	s.ReachJson = v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetSequenceCall(v []*GetXDefaultConfigResponseBodyDataSequenceCall) *GetXDefaultConfigResponseBodyData {
	s.SequenceCall = v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetSequenceEndTime(v string) *GetXDefaultConfigResponseBodyData {
	s.SequenceEndTime = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetSequenceStartTime(v string) *GetXDefaultConfigResponseBodyData {
	s.SequenceStartTime = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetSmsAbility(v string) *GetXDefaultConfigResponseBodyData {
	s.SmsAbility = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) SetSmsSignMode(v string) *GetXDefaultConfigResponseBodyData {
	s.SmsSignMode = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyData) Validate() error {
	if s.ReachJson != nil {
		for _, item := range s.ReachJson {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SequenceCall != nil {
		for _, item := range s.SequenceCall {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetXDefaultConfigResponseBodyDataReachJson struct {
	// The call direction. Valid values: 1: Employee B calls Customer A. 2: Customer A calls Employee B.
	//
	// example:
	//
	// 1
	CallDir *string `json:"CallDir,omitempty" xml:"CallDir,omitempty"`
	// The call status. Valid values: 1: ringing. 2: before connection. 3: after connection. 4: call ended. 5: connected. 6: not connected.
	//
	// example:
	//
	// 1
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// The receive direction. Valid values: 1: caller. 2: callee.
	//
	// example:
	//
	// 1
	ReceiveDir *string `json:"ReceiveDir,omitempty" xml:"ReceiveDir,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 345
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// Enterprise business card - SMS rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule type. Valid values: 1: enterprise business card - SMS. 2: enterprise business card - flash SMS. 3: enterprise business card - video. 4: enterprise business card - audio.
	//
	// example:
	//
	// 1
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The template ID.
	//
	// example:
	//
	// 12345
	TempId *string `json:"TempId,omitempty" xml:"TempId,omitempty"`
}

func (s GetXDefaultConfigResponseBodyDataReachJson) String() string {
	return dara.Prettify(s)
}

func (s GetXDefaultConfigResponseBodyDataReachJson) GoString() string {
	return s.String()
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) GetCallDir() *string {
	return s.CallDir
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) GetCallStatus() *string {
	return s.CallStatus
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) GetReceiveDir() *string {
	return s.ReceiveDir
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) GetRuleId() *string {
	return s.RuleId
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) GetRuleName() *string {
	return s.RuleName
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) GetRuleType() *string {
	return s.RuleType
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) GetTempId() *string {
	return s.TempId
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetCallDir(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.CallDir = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetCallStatus(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.CallStatus = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetReceiveDir(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.ReceiveDir = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetRuleId(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.RuleId = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetRuleName(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.RuleName = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetRuleType(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.RuleType = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) SetTempId(v string) *GetXDefaultConfigResponseBodyDataReachJson {
	s.TempId = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataReachJson) Validate() error {
	return dara.Validate(s)
}

type GetXDefaultConfigResponseBodyDataSequenceCall struct {
	// The play code for the sequential ringing prompt tone. The format is the same as callNoPlayCode.
	//
	// example:
	//
	// Sample value
	SequenceCallNoPlayCode *string `json:"SequenceCallNoPlayCode,omitempty" xml:"SequenceCallNoPlayCode,omitempty"`
	// The callee number for sequential ringing.
	//
	// example:
	//
	// Sample value sample value sample value
	SequenceCalledNo *string `json:"SequenceCalledNo,omitempty" xml:"SequenceCalledNo,omitempty"`
	// The play code for the caller and callee after the call is connected. The format is the same as calledPlayCode.
	//
	// example:
	//
	// Sample value sample value sample value
	SequenceCalledPlayCode *string `json:"SequenceCalledPlayCode,omitempty" xml:"SequenceCalledPlayCode,omitempty"`
}

func (s GetXDefaultConfigResponseBodyDataSequenceCall) String() string {
	return dara.Prettify(s)
}

func (s GetXDefaultConfigResponseBodyDataSequenceCall) GoString() string {
	return s.String()
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) GetSequenceCallNoPlayCode() *string {
	return s.SequenceCallNoPlayCode
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) GetSequenceCalledNo() *string {
	return s.SequenceCalledNo
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) GetSequenceCalledPlayCode() *string {
	return s.SequenceCalledPlayCode
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) SetSequenceCallNoPlayCode(v string) *GetXDefaultConfigResponseBodyDataSequenceCall {
	s.SequenceCallNoPlayCode = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) SetSequenceCalledNo(v string) *GetXDefaultConfigResponseBodyDataSequenceCall {
	s.SequenceCalledNo = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) SetSequenceCalledPlayCode(v string) *GetXDefaultConfigResponseBodyDataSequenceCall {
	s.SequenceCalledPlayCode = &v
	return s
}

func (s *GetXDefaultConfigResponseBodyDataSequenceCall) Validate() error {
	return dara.Validate(s)
}
