// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleBarChartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeRuleBarChartResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeRuleBarChartResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeRuleBarChartResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeRuleBarChartResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeRuleBarChartResponseBodyResultObject) *DescribeRuleBarChartResponseBody
	GetResultObject() *DescribeRuleBarChartResponseBodyResultObject
	SetSuccess(v bool) *DescribeRuleBarChartResponseBody
	GetSuccess() *bool
}

type DescribeRuleBarChartResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// The input parameter data is not valid. order_storage_company_num component not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject *DescribeRuleBarChartResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
	// Whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeRuleBarChartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleBarChartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleBarChartResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeRuleBarChartResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeRuleBarChartResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeRuleBarChartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleBarChartResponseBody) GetResultObject() *DescribeRuleBarChartResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeRuleBarChartResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRuleBarChartResponseBody) SetCode(v string) *DescribeRuleBarChartResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRuleBarChartResponseBody) SetHttpStatusCode(v string) *DescribeRuleBarChartResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeRuleBarChartResponseBody) SetMessage(v string) *DescribeRuleBarChartResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRuleBarChartResponseBody) SetRequestId(v string) *DescribeRuleBarChartResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleBarChartResponseBody) SetResultObject(v *DescribeRuleBarChartResponseBodyResultObject) *DescribeRuleBarChartResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeRuleBarChartResponseBody) SetSuccess(v bool) *DescribeRuleBarChartResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRuleBarChartResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRuleBarChartResponseBodyResultObject struct {
	// Data list
	Series []*DescribeRuleBarChartResponseBodyResultObjectSeries `json:"series,omitempty" xml:"series,omitempty" type:"Repeated"`
	// yaxis related results.
	Yaxis *DescribeRuleBarChartResponseBodyResultObjectYaxis `json:"yaxis,omitempty" xml:"yaxis,omitempty" type:"Struct"`
}

func (s DescribeRuleBarChartResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleBarChartResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeRuleBarChartResponseBodyResultObject) GetSeries() []*DescribeRuleBarChartResponseBodyResultObjectSeries {
	return s.Series
}

func (s *DescribeRuleBarChartResponseBodyResultObject) GetYaxis() *DescribeRuleBarChartResponseBodyResultObjectYaxis {
	return s.Yaxis
}

func (s *DescribeRuleBarChartResponseBodyResultObject) SetSeries(v []*DescribeRuleBarChartResponseBodyResultObjectSeries) *DescribeRuleBarChartResponseBodyResultObject {
	s.Series = v
	return s
}

func (s *DescribeRuleBarChartResponseBodyResultObject) SetYaxis(v *DescribeRuleBarChartResponseBodyResultObjectYaxis) *DescribeRuleBarChartResponseBodyResultObject {
	s.Yaxis = v
	return s
}

func (s *DescribeRuleBarChartResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeRuleBarChartResponseBodyResultObjectSeries struct {
	// Response data.
	Data []*DescribeRuleBarChartResponseBodyResultObjectSeriesData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Bar chart type
	//
	// example:
	//
	// bar
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeRuleBarChartResponseBodyResultObjectSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleBarChartResponseBodyResultObjectSeries) GoString() string {
	return s.String()
}

func (s *DescribeRuleBarChartResponseBodyResultObjectSeries) GetData() []*DescribeRuleBarChartResponseBodyResultObjectSeriesData {
	return s.Data
}

func (s *DescribeRuleBarChartResponseBodyResultObjectSeries) GetType() *string {
	return s.Type
}

func (s *DescribeRuleBarChartResponseBodyResultObjectSeries) SetData(v []*DescribeRuleBarChartResponseBodyResultObjectSeriesData) *DescribeRuleBarChartResponseBodyResultObjectSeries {
	s.Data = v
	return s
}

func (s *DescribeRuleBarChartResponseBodyResultObjectSeries) SetType(v string) *DescribeRuleBarChartResponseBodyResultObjectSeries {
	s.Type = &v
	return s
}

func (s *DescribeRuleBarChartResponseBodyResultObjectSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeRuleBarChartResponseBodyResultObjectSeriesData struct {
	// Event name.
	//
	// example:
	//
	// 营销事件
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Number.
	//
	// example:
	//
	// 100
	Num *int64 `json:"num,omitempty" xml:"num,omitempty"`
	// Policy name
	//
	// example:
	//
	// 营销风险识别评分_高风险_拒绝
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeRuleBarChartResponseBodyResultObjectSeriesData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleBarChartResponseBodyResultObjectSeriesData) GoString() string {
	return s.String()
}

func (s *DescribeRuleBarChartResponseBodyResultObjectSeriesData) GetEventName() *string {
	return s.EventName
}

func (s *DescribeRuleBarChartResponseBodyResultObjectSeriesData) GetNum() *int64 {
	return s.Num
}

func (s *DescribeRuleBarChartResponseBodyResultObjectSeriesData) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeRuleBarChartResponseBodyResultObjectSeriesData) GetStatus() *string {
	return s.Status
}

func (s *DescribeRuleBarChartResponseBodyResultObjectSeriesData) SetEventName(v string) *DescribeRuleBarChartResponseBodyResultObjectSeriesData {
	s.EventName = &v
	return s
}

func (s *DescribeRuleBarChartResponseBodyResultObjectSeriesData) SetNum(v int64) *DescribeRuleBarChartResponseBodyResultObjectSeriesData {
	s.Num = &v
	return s
}

func (s *DescribeRuleBarChartResponseBodyResultObjectSeriesData) SetRuleName(v string) *DescribeRuleBarChartResponseBodyResultObjectSeriesData {
	s.RuleName = &v
	return s
}

func (s *DescribeRuleBarChartResponseBodyResultObjectSeriesData) SetStatus(v string) *DescribeRuleBarChartResponseBodyResultObjectSeriesData {
	s.Status = &v
	return s
}

func (s *DescribeRuleBarChartResponseBodyResultObjectSeriesData) Validate() error {
	return dara.Validate(s)
}

type DescribeRuleBarChartResponseBodyResultObjectYaxis struct {
	// yaxis data items
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeRuleBarChartResponseBodyResultObjectYaxis) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleBarChartResponseBodyResultObjectYaxis) GoString() string {
	return s.String()
}

func (s *DescribeRuleBarChartResponseBodyResultObjectYaxis) GetData() []*string {
	return s.Data
}

func (s *DescribeRuleBarChartResponseBodyResultObjectYaxis) SetData(v []*string) *DescribeRuleBarChartResponseBodyResultObjectYaxis {
	s.Data = v
	return s
}

func (s *DescribeRuleBarChartResponseBodyResultObjectYaxis) Validate() error {
	return dara.Validate(s)
}
