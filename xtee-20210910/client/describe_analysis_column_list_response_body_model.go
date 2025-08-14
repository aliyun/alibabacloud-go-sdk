// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnalysisColumnListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAnalysisColumnListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeAnalysisColumnListResponseBodyResultObject) *DescribeAnalysisColumnListResponseBody
	GetResultObject() []*DescribeAnalysisColumnListResponseBodyResultObject
}

type DescribeAnalysisColumnListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject []*DescribeAnalysisColumnListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeAnalysisColumnListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalysisColumnListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAnalysisColumnListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAnalysisColumnListResponseBody) GetResultObject() []*DescribeAnalysisColumnListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeAnalysisColumnListResponseBody) SetRequestId(v string) *DescribeAnalysisColumnListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAnalysisColumnListResponseBody) SetResultObject(v []*DescribeAnalysisColumnListResponseBodyResultObject) *DescribeAnalysisColumnListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeAnalysisColumnListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAnalysisColumnListResponseBodyResultObject struct {
	// Event code
	//
	// example:
	//
	// de_aszbjb7236
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册风险
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Whether it is a default column.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// Variable name.
	//
	// example:
	//
	// age
	VariableName *string `json:"variableName,omitempty" xml:"variableName,omitempty"`
	// Variable title.
	//
	// example:
	//
	// 年龄
	VariableTitle *string `json:"variableTitle,omitempty" xml:"variableTitle,omitempty"`
	// Variable type.
	//
	// example:
	//
	// NATIVE
	VariableType *string `json:"variableType,omitempty" xml:"variableType,omitempty"`
}

func (s DescribeAnalysisColumnListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalysisColumnListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeAnalysisColumnListResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeAnalysisColumnListResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeAnalysisColumnListResponseBodyResultObject) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeAnalysisColumnListResponseBodyResultObject) GetVariableName() *string {
	return s.VariableName
}

func (s *DescribeAnalysisColumnListResponseBodyResultObject) GetVariableTitle() *string {
	return s.VariableTitle
}

func (s *DescribeAnalysisColumnListResponseBodyResultObject) GetVariableType() *string {
	return s.VariableType
}

func (s *DescribeAnalysisColumnListResponseBodyResultObject) SetEventCode(v string) *DescribeAnalysisColumnListResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeAnalysisColumnListResponseBodyResultObject) SetEventName(v string) *DescribeAnalysisColumnListResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeAnalysisColumnListResponseBodyResultObject) SetIsDefault(v bool) *DescribeAnalysisColumnListResponseBodyResultObject {
	s.IsDefault = &v
	return s
}

func (s *DescribeAnalysisColumnListResponseBodyResultObject) SetVariableName(v string) *DescribeAnalysisColumnListResponseBodyResultObject {
	s.VariableName = &v
	return s
}

func (s *DescribeAnalysisColumnListResponseBodyResultObject) SetVariableTitle(v string) *DescribeAnalysisColumnListResponseBodyResultObject {
	s.VariableTitle = &v
	return s
}

func (s *DescribeAnalysisColumnListResponseBodyResultObject) SetVariableType(v string) *DescribeAnalysisColumnListResponseBodyResultObject {
	s.VariableType = &v
	return s
}

func (s *DescribeAnalysisColumnListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
