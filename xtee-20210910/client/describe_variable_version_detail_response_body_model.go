// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableVersionDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVariableVersionDetailResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeVariableVersionDetailResponseBodyResultObject) *DescribeVariableVersionDetailResponseBody
	GetResultObject() *DescribeVariableVersionDetailResponseBodyResultObject
}

type DescribeVariableVersionDetailResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object.
	ResultObject *DescribeVariableVersionDetailResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeVariableVersionDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableVersionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVariableVersionDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVariableVersionDetailResponseBody) GetResultObject() *DescribeVariableVersionDetailResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeVariableVersionDetailResponseBody) SetRequestId(v string) *DescribeVariableVersionDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBody) SetResultObject(v *DescribeVariableVersionDetailResponseBodyResultObject) *DescribeVariableVersionDetailResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeVariableVersionDetailResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVariableVersionDetailResponseBodyResultObject struct {
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
	// Variable description information.
	//
	// example:
	//
	// 变量描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Event code.
	//
	// example:
	//
	// de_awkhwh0314
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Global cumulative. Not currently available externally.
	//
	// example:
	//
	// false
	Global *bool `json:"global,omitempty" xml:"global,omitempty"`
	// Historical value parameter type.
	//
	// example:
	//
	// COUNT
	HistoryValueType *string `json:"historyValueType,omitempty" xml:"historyValueType,omitempty"`
	// Primary key ID of the version.
	//
	// example:
	//
	// 3434
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Subordinate attribute (single selection, sourced from the event attribute list code).
	//
	// example:
	//
	// name
	Object *string `json:"object,omitempty" xml:"object,omitempty"`
	// Status.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Main attribute (multiple separated by commas, up to 5, sourced from the event attribute list code).
	//
	// example:
	//
	// age
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// Time slice type.
	//
	// example:
	//
	// CURRENT
	TimeType *string `json:"timeType,omitempty" xml:"timeType,omitempty"`
	// Title.
	//
	// example:
	//
	// 累计个数
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Top N.
	//
	// example:
	//
	// topN累计
	TopN *int64 `json:"topN,omitempty" xml:"topN,omitempty"`
	// Number of time slices.
	//
	// example:
	//
	// 1
	TwCount *int64 `json:"twCount,omitempty" xml:"twCount,omitempty"`
	// Cumulative indicator processing function.
	//
	// example:
	//
	// COUNT
	VelocityFC *string `json:"velocityFC,omitempty" xml:"velocityFC,omitempty"`
	// Cumulative indicator time window.
	//
	// example:
	//
	// MONTH_1
	VelocityTW *string `json:"velocityTW,omitempty" xml:"velocityTW,omitempty"`
}

func (s DescribeVariableVersionDetailResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableVersionDetailResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) GetCondition() *string {
	return s.Condition
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) GetDataVersion() *int64 {
	return s.DataVersion
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) GetGlobal() *bool {
	return s.Global
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) GetHistoryValueType() *string {
	return s.HistoryValueType
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) GetObject() *string {
	return s.Object
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) GetSubject() *string {
	return s.Subject
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) GetTimeType() *string {
	return s.TimeType
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) GetTitle() *string {
	return s.Title
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) GetTopN() *int64 {
	return s.TopN
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) GetTwCount() *int64 {
	return s.TwCount
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) GetVelocityFC() *string {
	return s.VelocityFC
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) GetVelocityTW() *string {
	return s.VelocityTW
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) SetCondition(v string) *DescribeVariableVersionDetailResponseBodyResultObject {
	s.Condition = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) SetDataVersion(v int64) *DescribeVariableVersionDetailResponseBodyResultObject {
	s.DataVersion = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) SetDescription(v string) *DescribeVariableVersionDetailResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) SetEventCodes(v string) *DescribeVariableVersionDetailResponseBodyResultObject {
	s.EventCodes = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) SetGlobal(v bool) *DescribeVariableVersionDetailResponseBodyResultObject {
	s.Global = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) SetHistoryValueType(v string) *DescribeVariableVersionDetailResponseBodyResultObject {
	s.HistoryValueType = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) SetId(v int64) *DescribeVariableVersionDetailResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) SetObject(v string) *DescribeVariableVersionDetailResponseBodyResultObject {
	s.Object = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) SetStatus(v string) *DescribeVariableVersionDetailResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) SetSubject(v string) *DescribeVariableVersionDetailResponseBodyResultObject {
	s.Subject = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) SetTimeType(v string) *DescribeVariableVersionDetailResponseBodyResultObject {
	s.TimeType = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) SetTitle(v string) *DescribeVariableVersionDetailResponseBodyResultObject {
	s.Title = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) SetTopN(v int64) *DescribeVariableVersionDetailResponseBodyResultObject {
	s.TopN = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) SetTwCount(v int64) *DescribeVariableVersionDetailResponseBodyResultObject {
	s.TwCount = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) SetVelocityFC(v string) *DescribeVariableVersionDetailResponseBodyResultObject {
	s.VelocityFC = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) SetVelocityTW(v string) *DescribeVariableVersionDetailResponseBodyResultObject {
	s.VelocityTW = &v
	return s
}

func (s *DescribeVariableVersionDetailResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
