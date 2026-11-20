// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetXConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetXConfigResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetXConfigResponseBody
	GetCode() *string
	SetData(v *GetXConfigResponseBodyData) *GetXConfigResponseBody
	GetData() *GetXConfigResponseBodyData
	SetMessage(v string) *GetXConfigResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetXConfigResponseBody
	GetSuccess() *bool
}

type GetXConfigResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code. Valid values:
	//
	//
	//
	// - 0000: Success.
	//
	//
	//
	// Other values indicate failure.
	//
	// example:
	//
	// 0000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configuration of the work number.
	//
	// example:
	//
	// {"CallAbility":"0","GNFlag":"00","SmsAbility":"0","SmsSignMode":"0"}
	Data *GetXConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: Success.
	//
	// - false: Failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetXConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetXConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetXConfigResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetXConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetXConfigResponseBody) GetData() *GetXConfigResponseBodyData {
	return s.Data
}

func (s *GetXConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetXConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetXConfigResponseBody) SetAccessDeniedDetail(v string) *GetXConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetXConfigResponseBody) SetCode(v string) *GetXConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetXConfigResponseBody) SetData(v *GetXConfigResponseBodyData) *GetXConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetXConfigResponseBody) SetMessage(v string) *GetXConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetXConfigResponseBody) SetSuccess(v bool) *GetXConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetXConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetXConfigResponseBodyData struct {
	// Specifies whether the call capability is enabled. Valid values:
	//
	// - 0: Disabled.
	//
	// - 1: Enabled.
	//
	// example:
	//
	// 0
	CallAbility *string `json:"CallAbility,omitempty" xml:"CallAbility,omitempty"`
	// Specifies whether to pass through the caller ID to display the real caller. Valid values:
	//
	// - 00: Non-passthrough. The work number is displayed when calls are made to each other.
	//
	// - 10: Passthrough. When Customer A is the caller, the caller ID displayed to Employee B is the number of Customer A. When Employee B is the caller, the caller ID displayed to Customer A is the work number.
	//
	// Default value: 00.
	//
	// example:
	//
	// 00
	GNFlag *string `json:"GNFlag,omitempty" xml:"GNFlag,omitempty"`
	// The control parameters for enterprise profile rules.
	ReachJsons []*GetXConfigResponseBodyDataReachJsons `json:"ReachJsons,omitempty" xml:"ReachJsons,omitempty" type:"Repeated"`
	// The control parameters for sequence ringing.
	SequenceCalls []*GetXConfigResponseBodyDataSequenceCalls `json:"SequenceCalls,omitempty" xml:"SequenceCalls,omitempty" type:"Repeated"`
	// The end time of sequence ringing. Format: HH:mm:ss, such as 18:00:00.
	//
	// example:
	//
	// 09:00:00
	SequenceEndTime *string `json:"SequenceEndTime,omitempty" xml:"SequenceEndTime,omitempty"`
	// The start time of sequence ringing. Format: HH:mm:ss, such as 09:00:00.
	//
	// example:
	//
	// 09:00:00
	SequenceStartTime *string `json:"SequenceStartTime,omitempty" xml:"SequenceStartTime,omitempty"`
	// Specifies whether the SMS feature is enabled. Valid values:
	//
	// - 0: Disabled.
	//
	// - 1: Enabled.
	//
	// example:
	//
	// 0
	SmsAbility *string `json:"SmsAbility,omitempty" xml:"SmsAbility,omitempty"`
	// Specifies whether to pass through the caller ID to display the real caller. Valid values:
	//
	// - 00: Non-passthrough. The work number is displayed when calls are made to each other.
	//
	// - 10: Passthrough. When Customer A is the caller, the caller ID displayed to Employee B is the number of Customer A. When Employee B is the caller, the caller ID displayed to Customer A is the work number.
	//
	// Default value: 00.
	//
	// example:
	//
	// 0
	SmsSignMode *string `json:"SmsSignMode,omitempty" xml:"SmsSignMode,omitempty"`
}

func (s GetXConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetXConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetXConfigResponseBodyData) GetCallAbility() *string {
	return s.CallAbility
}

func (s *GetXConfigResponseBodyData) GetGNFlag() *string {
	return s.GNFlag
}

func (s *GetXConfigResponseBodyData) GetReachJsons() []*GetXConfigResponseBodyDataReachJsons {
	return s.ReachJsons
}

func (s *GetXConfigResponseBodyData) GetSequenceCalls() []*GetXConfigResponseBodyDataSequenceCalls {
	return s.SequenceCalls
}

func (s *GetXConfigResponseBodyData) GetSequenceEndTime() *string {
	return s.SequenceEndTime
}

func (s *GetXConfigResponseBodyData) GetSequenceStartTime() *string {
	return s.SequenceStartTime
}

func (s *GetXConfigResponseBodyData) GetSmsAbility() *string {
	return s.SmsAbility
}

func (s *GetXConfigResponseBodyData) GetSmsSignMode() *string {
	return s.SmsSignMode
}

func (s *GetXConfigResponseBodyData) SetCallAbility(v string) *GetXConfigResponseBodyData {
	s.CallAbility = &v
	return s
}

func (s *GetXConfigResponseBodyData) SetGNFlag(v string) *GetXConfigResponseBodyData {
	s.GNFlag = &v
	return s
}

func (s *GetXConfigResponseBodyData) SetReachJsons(v []*GetXConfigResponseBodyDataReachJsons) *GetXConfigResponseBodyData {
	s.ReachJsons = v
	return s
}

func (s *GetXConfigResponseBodyData) SetSequenceCalls(v []*GetXConfigResponseBodyDataSequenceCalls) *GetXConfigResponseBodyData {
	s.SequenceCalls = v
	return s
}

func (s *GetXConfigResponseBodyData) SetSequenceEndTime(v string) *GetXConfigResponseBodyData {
	s.SequenceEndTime = &v
	return s
}

func (s *GetXConfigResponseBodyData) SetSequenceStartTime(v string) *GetXConfigResponseBodyData {
	s.SequenceStartTime = &v
	return s
}

func (s *GetXConfigResponseBodyData) SetSmsAbility(v string) *GetXConfigResponseBodyData {
	s.SmsAbility = &v
	return s
}

func (s *GetXConfigResponseBodyData) SetSmsSignMode(v string) *GetXConfigResponseBodyData {
	s.SmsSignMode = &v
	return s
}

func (s *GetXConfigResponseBodyData) Validate() error {
	if s.ReachJsons != nil {
		for _, item := range s.ReachJsons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SequenceCalls != nil {
		for _, item := range s.SequenceCalls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetXConfigResponseBodyDataReachJsons struct {
	// The call direction. Valid values:
	//
	// - 1: Employee B calls Customer A.
	//
	// - 2: Customer A calls Employee B.
	//
	// example:
	//
	// 1
	CallDir *string `json:"CallDir,omitempty" xml:"CallDir,omitempty"`
	// The call status. Valid values:
	//
	// - 1: Call ringing.
	//
	// - 2: Before connection.
	//
	// - 3: After connection.
	//
	// - 4: Call ended.
	//
	// - 5: Connected.
	//
	// - 6: Not connected.
	//
	// example:
	//
	// 1
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// The receiving direction. Valid values:
	//
	// - 1: Caller.
	//
	// - 2: Called party.
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
	// Enterprise profile - SMS rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule type. Valid values:
	//
	// - 1: Enterprise profile - SMS.
	//
	// - 2: Enterprise profile - Flash message.
	//
	// - 3: Enterprise profile - Video.
	//
	// - 4: Enterprise profile - Audio.
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

func (s GetXConfigResponseBodyDataReachJsons) String() string {
	return dara.Prettify(s)
}

func (s GetXConfigResponseBodyDataReachJsons) GoString() string {
	return s.String()
}

func (s *GetXConfigResponseBodyDataReachJsons) GetCallDir() *string {
	return s.CallDir
}

func (s *GetXConfigResponseBodyDataReachJsons) GetCallStatus() *string {
	return s.CallStatus
}

func (s *GetXConfigResponseBodyDataReachJsons) GetReceiveDir() *string {
	return s.ReceiveDir
}

func (s *GetXConfigResponseBodyDataReachJsons) GetRuleId() *string {
	return s.RuleId
}

func (s *GetXConfigResponseBodyDataReachJsons) GetRuleName() *string {
	return s.RuleName
}

func (s *GetXConfigResponseBodyDataReachJsons) GetRuleType() *string {
	return s.RuleType
}

func (s *GetXConfigResponseBodyDataReachJsons) GetTempId() *string {
	return s.TempId
}

func (s *GetXConfigResponseBodyDataReachJsons) SetCallDir(v string) *GetXConfigResponseBodyDataReachJsons {
	s.CallDir = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetCallStatus(v string) *GetXConfigResponseBodyDataReachJsons {
	s.CallStatus = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetReceiveDir(v string) *GetXConfigResponseBodyDataReachJsons {
	s.ReceiveDir = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetRuleId(v string) *GetXConfigResponseBodyDataReachJsons {
	s.RuleId = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetRuleName(v string) *GetXConfigResponseBodyDataReachJsons {
	s.RuleName = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetRuleType(v string) *GetXConfigResponseBodyDataReachJsons {
	s.RuleType = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) SetTempId(v string) *GetXConfigResponseBodyDataReachJsons {
	s.TempId = &v
	return s
}

func (s *GetXConfigResponseBodyDataReachJsons) Validate() error {
	return dara.Validate(s)
}

type GetXConfigResponseBodyDataSequenceCalls struct {
	// The play ID of the prompt tone for sequence ringing. The format is the same as callNoPlayCode.
	//
	// example:
	//
	// Sample value sample value sample value
	SequenceCallNoPlayCode *string `json:"SequenceCallNoPlayCode,omitempty" xml:"SequenceCallNoPlayCode,omitempty"`
	// The called number for sequence ringing.
	//
	// example:
	//
	// Sample value sample value sample value
	SequenceCalledNo *string `json:"SequenceCalledNo,omitempty" xml:"SequenceCalledNo,omitempty"`
	// The play ID for the caller and called party after the call is connected. The format is the same as calledPlayCode.
	//
	// example:
	//
	// Sample value sample value sample value
	SequenceCalledPlayCode *string `json:"SequenceCalledPlayCode,omitempty" xml:"SequenceCalledPlayCode,omitempty"`
}

func (s GetXConfigResponseBodyDataSequenceCalls) String() string {
	return dara.Prettify(s)
}

func (s GetXConfigResponseBodyDataSequenceCalls) GoString() string {
	return s.String()
}

func (s *GetXConfigResponseBodyDataSequenceCalls) GetSequenceCallNoPlayCode() *string {
	return s.SequenceCallNoPlayCode
}

func (s *GetXConfigResponseBodyDataSequenceCalls) GetSequenceCalledNo() *string {
	return s.SequenceCalledNo
}

func (s *GetXConfigResponseBodyDataSequenceCalls) GetSequenceCalledPlayCode() *string {
	return s.SequenceCalledPlayCode
}

func (s *GetXConfigResponseBodyDataSequenceCalls) SetSequenceCallNoPlayCode(v string) *GetXConfigResponseBodyDataSequenceCalls {
	s.SequenceCallNoPlayCode = &v
	return s
}

func (s *GetXConfigResponseBodyDataSequenceCalls) SetSequenceCalledNo(v string) *GetXConfigResponseBodyDataSequenceCalls {
	s.SequenceCalledNo = &v
	return s
}

func (s *GetXConfigResponseBodyDataSequenceCalls) SetSequenceCalledPlayCode(v string) *GetXConfigResponseBodyDataSequenceCalls {
	s.SequenceCalledPlayCode = &v
	return s
}

func (s *GetXConfigResponseBodyDataSequenceCalls) Validate() error {
	return dara.Validate(s)
}
