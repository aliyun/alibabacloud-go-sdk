// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustVariableDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCustVariableDetailResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeCustVariableDetailResponseBodyResultObject) *DescribeCustVariableDetailResponseBody
	GetResultObject() []*DescribeCustVariableDetailResponseBodyResultObject
}

type DescribeCustVariableDetailResponseBody struct {
	// Request ID
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject []*DescribeCustVariableDetailResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeCustVariableDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustVariableDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustVariableDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustVariableDetailResponseBody) GetResultObject() []*DescribeCustVariableDetailResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeCustVariableDetailResponseBody) SetRequestId(v string) *DescribeCustVariableDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustVariableDetailResponseBody) SetResultObject(v []*DescribeCustVariableDetailResponseBodyResultObject) *DescribeCustVariableDetailResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeCustVariableDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCustVariableDetailResponseBodyResultObject struct {
	// Condition value.
	//
	// example:
	//
	// {"currentId":1,"deepCount":0,"list":[{"currentId":2,"deepCount":1,"left":{"code":"accountId","fieldType":"STRING","functionName":"","hasRightVariable":false,"name":"accountId"},"operatorCode":"isNotEmptyWrapper","parentId":1,"right":{"code":"\\"A\\nB\\nC\\"","functionName":"","name":"","rightVariableType":"constant"}}],"parentId":0,"relationship":"and"}
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// Data version.
	//
	// example:
	//
	// 1
	DataVersion *int64 `json:"dataVersion,omitempty" xml:"dataVersion,omitempty"`
	// Description information.
	//
	// example:
	//
	// 变量描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Event code.
	//
	// example:
	//
	// de_ahespg8137
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Value type
	//
	// example:
	//
	// EARLIEST
	HistoryValueType *string `json:"historyValueType,omitempty" xml:"historyValueType,omitempty"`
	// Primary key ID of the cumulative variable
	//
	// example:
	//
	// 2793
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Cumulative object
	//
	// example:
	//
	// DEpage
	Object *string `json:"object,omitempty" xml:"object,omitempty"`
	// Variable return type
	//
	// example:
	//
	// DOUBLE
	Outputs *string `json:"outputs,omitempty" xml:"outputs,omitempty"`
	// Main object
	//
	// example:
	//
	// DEpname
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// Time slice type
	//
	// example:
	//
	// NEAR
	TimeType *string `json:"timeType,omitempty" xml:"timeType,omitempty"`
	// Title.
	//
	// example:
	//
	// 累计变量求平均值
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Time count
	//
	// example:
	//
	// 1
	TwCount *string `json:"twCount,omitempty" xml:"twCount,omitempty"`
	// Variable type
	//
	// example:
	//
	// DISTINCT
	VelocityFC *string `json:"velocityFC,omitempty" xml:"velocityFC,omitempty"`
	// Time slice unit
	//
	// example:
	//
	// HOUR_1
	VelocityTW *string `json:"velocityTW,omitempty" xml:"velocityTW,omitempty"`
}

func (s DescribeCustVariableDetailResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustVariableDetailResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) GetCondition() *string {
	return s.Condition
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) GetDataVersion() *int64 {
	return s.DataVersion
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) GetHistoryValueType() *string {
	return s.HistoryValueType
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) GetObject() *string {
	return s.Object
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) GetOutputs() *string {
	return s.Outputs
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) GetSubject() *string {
	return s.Subject
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) GetTimeType() *string {
	return s.TimeType
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) GetTitle() *string {
	return s.Title
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) GetTwCount() *string {
	return s.TwCount
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) GetVelocityFC() *string {
	return s.VelocityFC
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) GetVelocityTW() *string {
	return s.VelocityTW
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) SetCondition(v string) *DescribeCustVariableDetailResponseBodyResultObject {
	s.Condition = &v
	return s
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) SetDataVersion(v int64) *DescribeCustVariableDetailResponseBodyResultObject {
	s.DataVersion = &v
	return s
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) SetDescription(v string) *DescribeCustVariableDetailResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) SetEventCodes(v string) *DescribeCustVariableDetailResponseBodyResultObject {
	s.EventCodes = &v
	return s
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) SetHistoryValueType(v string) *DescribeCustVariableDetailResponseBodyResultObject {
	s.HistoryValueType = &v
	return s
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) SetId(v int64) *DescribeCustVariableDetailResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) SetObject(v string) *DescribeCustVariableDetailResponseBodyResultObject {
	s.Object = &v
	return s
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) SetOutputs(v string) *DescribeCustVariableDetailResponseBodyResultObject {
	s.Outputs = &v
	return s
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) SetSubject(v string) *DescribeCustVariableDetailResponseBodyResultObject {
	s.Subject = &v
	return s
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) SetTimeType(v string) *DescribeCustVariableDetailResponseBodyResultObject {
	s.TimeType = &v
	return s
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) SetTitle(v string) *DescribeCustVariableDetailResponseBodyResultObject {
	s.Title = &v
	return s
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) SetTwCount(v string) *DescribeCustVariableDetailResponseBodyResultObject {
	s.TwCount = &v
	return s
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) SetVelocityFC(v string) *DescribeCustVariableDetailResponseBodyResultObject {
	s.VelocityFC = &v
	return s
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) SetVelocityTW(v string) *DescribeCustVariableDetailResponseBodyResultObject {
	s.VelocityTW = &v
	return s
}

func (s *DescribeCustVariableDetailResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
