// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateCustVariableRequest
	GetLang() *string
	SetCondition(v string) *CreateCustVariableRequest
	GetCondition() *string
	SetCreateType(v string) *CreateCustVariableRequest
	GetCreateType() *string
	SetDescription(v string) *CreateCustVariableRequest
	GetDescription() *string
	SetEventCodes(v string) *CreateCustVariableRequest
	GetEventCodes() *string
	SetHistoryValueType(v string) *CreateCustVariableRequest
	GetHistoryValueType() *string
	SetObject(v string) *CreateCustVariableRequest
	GetObject() *string
	SetOutputs(v string) *CreateCustVariableRequest
	GetOutputs() *string
	SetRegId(v string) *CreateCustVariableRequest
	GetRegId() *string
	SetSubject(v string) *CreateCustVariableRequest
	GetSubject() *string
	SetTimeType(v string) *CreateCustVariableRequest
	GetTimeType() *string
	SetTitle(v string) *CreateCustVariableRequest
	GetTitle() *string
	SetTwCount(v int32) *CreateCustVariableRequest
	GetTwCount() *int32
	SetVelocityFC(v string) *CreateCustVariableRequest
	GetVelocityFC() *string
	SetVelocityTW(v string) *CreateCustVariableRequest
	GetVelocityTW() *string
}

type CreateCustVariableRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Condition value.
	//
	// example:
	//
	// {"relationship":"and","list":[{"deepCount":1,"left":{"hasRightVariable":true,"fieldType":"INT","functionName":"","leftVariableType":"NATIVE","name":"DEtest222","operatorCode":"equals"},"right":{"rightVariableType":"constant","name":"11","functionName":""},"operatorCode":"equals"}]}
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Description information.
	//
	// example:
	//
	// ip调用次数累计描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Event code
	//
	// This parameter is required.
	//
	// example:
	//
	// de_ahqhsw7665,de_agbzfi5134
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Value type
	//
	// example:
	//
	// EARLIEST
	HistoryValueType *string `json:"historyValueType,omitempty" xml:"historyValueType,omitempty"`
	// Accumulative object
	//
	// example:
	//
	// age
	Object *string `json:"object,omitempty" xml:"object,omitempty"`
	// Variable return type
	//
	// example:
	//
	// STRING
	Outputs *string `json:"outputs,omitempty" xml:"outputs,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Primary object
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// Time slice type
	//
	// This parameter is required.
	//
	// example:
	//
	// CURRENT
	TimeType *string `json:"timeType,omitempty" xml:"timeType,omitempty"`
	// Title.
	//
	// This parameter is required.
	//
	// example:
	//
	// ip调用次数累计
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Number of time units
	//
	// example:
	//
	// 1
	TwCount *int32 `json:"twCount,omitempty" xml:"twCount,omitempty"`
	// Variable type
	//
	// This parameter is required.
	//
	// example:
	//
	// DISTINCT
	VelocityFC *string `json:"velocityFC,omitempty" xml:"velocityFC,omitempty"`
	// Time slice unit
	//
	// This parameter is required.
	//
	// example:
	//
	// DAY_1
	VelocityTW *string `json:"velocityTW,omitempty" xml:"velocityTW,omitempty"`
}

func (s CreateCustVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustVariableRequest) GoString() string {
	return s.String()
}

func (s *CreateCustVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateCustVariableRequest) GetCondition() *string {
	return s.Condition
}

func (s *CreateCustVariableRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *CreateCustVariableRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCustVariableRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *CreateCustVariableRequest) GetHistoryValueType() *string {
	return s.HistoryValueType
}

func (s *CreateCustVariableRequest) GetObject() *string {
	return s.Object
}

func (s *CreateCustVariableRequest) GetOutputs() *string {
	return s.Outputs
}

func (s *CreateCustVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateCustVariableRequest) GetSubject() *string {
	return s.Subject
}

func (s *CreateCustVariableRequest) GetTimeType() *string {
	return s.TimeType
}

func (s *CreateCustVariableRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateCustVariableRequest) GetTwCount() *int32 {
	return s.TwCount
}

func (s *CreateCustVariableRequest) GetVelocityFC() *string {
	return s.VelocityFC
}

func (s *CreateCustVariableRequest) GetVelocityTW() *string {
	return s.VelocityTW
}

func (s *CreateCustVariableRequest) SetLang(v string) *CreateCustVariableRequest {
	s.Lang = &v
	return s
}

func (s *CreateCustVariableRequest) SetCondition(v string) *CreateCustVariableRequest {
	s.Condition = &v
	return s
}

func (s *CreateCustVariableRequest) SetCreateType(v string) *CreateCustVariableRequest {
	s.CreateType = &v
	return s
}

func (s *CreateCustVariableRequest) SetDescription(v string) *CreateCustVariableRequest {
	s.Description = &v
	return s
}

func (s *CreateCustVariableRequest) SetEventCodes(v string) *CreateCustVariableRequest {
	s.EventCodes = &v
	return s
}

func (s *CreateCustVariableRequest) SetHistoryValueType(v string) *CreateCustVariableRequest {
	s.HistoryValueType = &v
	return s
}

func (s *CreateCustVariableRequest) SetObject(v string) *CreateCustVariableRequest {
	s.Object = &v
	return s
}

func (s *CreateCustVariableRequest) SetOutputs(v string) *CreateCustVariableRequest {
	s.Outputs = &v
	return s
}

func (s *CreateCustVariableRequest) SetRegId(v string) *CreateCustVariableRequest {
	s.RegId = &v
	return s
}

func (s *CreateCustVariableRequest) SetSubject(v string) *CreateCustVariableRequest {
	s.Subject = &v
	return s
}

func (s *CreateCustVariableRequest) SetTimeType(v string) *CreateCustVariableRequest {
	s.TimeType = &v
	return s
}

func (s *CreateCustVariableRequest) SetTitle(v string) *CreateCustVariableRequest {
	s.Title = &v
	return s
}

func (s *CreateCustVariableRequest) SetTwCount(v int32) *CreateCustVariableRequest {
	s.TwCount = &v
	return s
}

func (s *CreateCustVariableRequest) SetVelocityFC(v string) *CreateCustVariableRequest {
	s.VelocityFC = &v
	return s
}

func (s *CreateCustVariableRequest) SetVelocityTW(v string) *CreateCustVariableRequest {
	s.VelocityTW = &v
	return s
}

func (s *CreateCustVariableRequest) Validate() error {
	return dara.Validate(s)
}
