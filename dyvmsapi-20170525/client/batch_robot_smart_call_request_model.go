// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRobotSmartCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalledNumber(v string) *BatchRobotSmartCallRequest
	GetCalledNumber() *string
	SetCalledShowNumber(v string) *BatchRobotSmartCallRequest
	GetCalledShowNumber() *string
	SetCorpName(v string) *BatchRobotSmartCallRequest
	GetCorpName() *string
	SetDialogId(v string) *BatchRobotSmartCallRequest
	GetDialogId() *string
	SetEarlyMediaAsr(v bool) *BatchRobotSmartCallRequest
	GetEarlyMediaAsr() *bool
	SetIsSelfLine(v bool) *BatchRobotSmartCallRequest
	GetIsSelfLine() *bool
	SetOwnerId(v int64) *BatchRobotSmartCallRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *BatchRobotSmartCallRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BatchRobotSmartCallRequest
	GetResourceOwnerId() *int64
	SetScheduleCall(v bool) *BatchRobotSmartCallRequest
	GetScheduleCall() *bool
	SetScheduleTime(v int64) *BatchRobotSmartCallRequest
	GetScheduleTime() *int64
	SetTaskName(v string) *BatchRobotSmartCallRequest
	GetTaskName() *string
	SetTtsParam(v string) *BatchRobotSmartCallRequest
	GetTtsParam() *string
	SetTtsParamHead(v string) *BatchRobotSmartCallRequest
	GetTtsParamHead() *string
}

type BatchRobotSmartCallRequest struct {
	// The called number. Only mobile phone numbers in the Chinese mainland are supported.
	//
	// You can set up to 1,000 called numbers and separate the numbers with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The number displayed to called parties, which must be a number you purchased. You can view the numbers you purchased in the [Voice Messaging Service console](https://dyvms.console.aliyun.com/dyvms.htm#/number/normal).
	//
	// You can set up to 100 numbers and separate the numbers with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 222
	CalledShowNumber *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	// The company name, which must be the same as the **company name*	- specified on the [qualification management page](https://dyvms.console.aliyun.com/dyvms.htm#/corp/normal).
	//
	// > This parameter is optional if **isSelfLine*	- is set to **true**.
	//
	// example:
	//
	// Alibaba
	CorpName *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	// The ID of the robot or communication script that is used to initiate a call.
	//
	// You can obtain the **communication script ID*	- from the [Communication script management](https://dyvms.console.aliyun.com/dyvms.htm#/smart-call/saas/robot/list) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	DialogId *string `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	// The speech recognition identifier of early media. The default value is **false**, which means that the speech recognition identifier of early media is not enabled.
	//
	// Set the parameter to **true*	- if you want to enable the speech recognition identifier of early media.
	//
	// example:
	//
	// true
	EarlyMediaAsr *bool `json:"EarlyMediaAsr,omitempty" xml:"EarlyMediaAsr,omitempty"`
	// Specifies whether to call the self-managed line. Default value: **false**.
	//
	// example:
	//
	// true
	IsSelfLine           *bool   `json:"IsSelfLine,omitempty" xml:"IsSelfLine,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether the call is scheduled. If you set this parameter to **true**, the **ScheduleTime*	- parameter is required.
	//
	// example:
	//
	// true
	ScheduleCall *bool `json:"ScheduleCall,omitempty" xml:"ScheduleCall,omitempty"`
	// The preset call time. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// >  This parameter is required only when **ScheduleCall*	- is set to **true**.
	//
	// example:
	//
	// 12
	ScheduleTime *int64 `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The task name. The task name can be up to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Batch Tasks
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The variable value of the TTS template, in the JSON format.
	//
	// The variable value must correspond to a number. The TtsParam parameter must be used together with the TtsParamHead parameter.
	//
	// example:
	//
	// [{"number":"1390000****","params":[“Miss li”,"miss wang","Mr.li"]}]
	TtsParam *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	// The call tasks with variables, in the JSON format.
	//
	// The parameter value is a list of variable names. The TtsParamHead parameter must be used together with the TtsParam parameter.
	//
	// example:
	//
	// ["name1","name2","name3"]
	TtsParamHead *string `json:"TtsParamHead,omitempty" xml:"TtsParamHead,omitempty"`
}

func (s BatchRobotSmartCallRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchRobotSmartCallRequest) GoString() string {
	return s.String()
}

func (s *BatchRobotSmartCallRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *BatchRobotSmartCallRequest) GetCalledShowNumber() *string {
	return s.CalledShowNumber
}

func (s *BatchRobotSmartCallRequest) GetCorpName() *string {
	return s.CorpName
}

func (s *BatchRobotSmartCallRequest) GetDialogId() *string {
	return s.DialogId
}

func (s *BatchRobotSmartCallRequest) GetEarlyMediaAsr() *bool {
	return s.EarlyMediaAsr
}

func (s *BatchRobotSmartCallRequest) GetIsSelfLine() *bool {
	return s.IsSelfLine
}

func (s *BatchRobotSmartCallRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchRobotSmartCallRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BatchRobotSmartCallRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BatchRobotSmartCallRequest) GetScheduleCall() *bool {
	return s.ScheduleCall
}

func (s *BatchRobotSmartCallRequest) GetScheduleTime() *int64 {
	return s.ScheduleTime
}

func (s *BatchRobotSmartCallRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *BatchRobotSmartCallRequest) GetTtsParam() *string {
	return s.TtsParam
}

func (s *BatchRobotSmartCallRequest) GetTtsParamHead() *string {
	return s.TtsParamHead
}

func (s *BatchRobotSmartCallRequest) SetCalledNumber(v string) *BatchRobotSmartCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetCalledShowNumber(v string) *BatchRobotSmartCallRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetCorpName(v string) *BatchRobotSmartCallRequest {
	s.CorpName = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetDialogId(v string) *BatchRobotSmartCallRequest {
	s.DialogId = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetEarlyMediaAsr(v bool) *BatchRobotSmartCallRequest {
	s.EarlyMediaAsr = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetIsSelfLine(v bool) *BatchRobotSmartCallRequest {
	s.IsSelfLine = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetOwnerId(v int64) *BatchRobotSmartCallRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetResourceOwnerAccount(v string) *BatchRobotSmartCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetResourceOwnerId(v int64) *BatchRobotSmartCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetScheduleCall(v bool) *BatchRobotSmartCallRequest {
	s.ScheduleCall = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetScheduleTime(v int64) *BatchRobotSmartCallRequest {
	s.ScheduleTime = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetTaskName(v string) *BatchRobotSmartCallRequest {
	s.TaskName = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetTtsParam(v string) *BatchRobotSmartCallRequest {
	s.TtsParam = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetTtsParamHead(v string) *BatchRobotSmartCallRequest {
	s.TtsParamHead = &v
	return s
}

func (s *BatchRobotSmartCallRequest) Validate() error {
	return dara.Validate(s)
}
