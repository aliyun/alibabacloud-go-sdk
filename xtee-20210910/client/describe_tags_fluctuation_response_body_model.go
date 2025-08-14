// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsFluctuationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeTagsFluctuationResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeTagsFluctuationResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeTagsFluctuationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeTagsFluctuationResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeTagsFluctuationResponseBodyResultObject) *DescribeTagsFluctuationResponseBody
	GetResultObject() []*DescribeTagsFluctuationResponseBodyResultObject
	SetSuccess(v bool) *DescribeTagsFluctuationResponseBody
	GetSuccess() *bool
}

type DescribeTagsFluctuationResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// HTTP status code
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
	// Return object
	ResultObject []*DescribeTagsFluctuationResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Indicates whether the call was successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeTagsFluctuationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsFluctuationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsFluctuationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeTagsFluctuationResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeTagsFluctuationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTagsFluctuationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagsFluctuationResponseBody) GetResultObject() []*DescribeTagsFluctuationResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeTagsFluctuationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTagsFluctuationResponseBody) SetCode(v string) *DescribeTagsFluctuationResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTagsFluctuationResponseBody) SetHttpStatusCode(v string) *DescribeTagsFluctuationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeTagsFluctuationResponseBody) SetMessage(v string) *DescribeTagsFluctuationResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTagsFluctuationResponseBody) SetRequestId(v string) *DescribeTagsFluctuationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsFluctuationResponseBody) SetResultObject(v []*DescribeTagsFluctuationResponseBodyResultObject) *DescribeTagsFluctuationResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeTagsFluctuationResponseBody) SetSuccess(v bool) *DescribeTagsFluctuationResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTagsFluctuationResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTagsFluctuationResponseBodyResultObject struct {
	// Tag name
	//
	// example:
	//
	// accountId
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// Data for today
	//
	// example:
	//
	// 100
	TodayNum *int64 `json:"todayNum,omitempty" xml:"todayNum,omitempty"`
	// Data for the last seven days
	//
	// example:
	//
	// 600
	WithinSevenDayNum *string `json:"withinSevenDayNum,omitempty" xml:"withinSevenDayNum,omitempty"`
	// Data for the last thirty days
	//
	// example:
	//
	// 1200
	WithinThirtyDayNum *string `json:"withinThirtyDayNum,omitempty" xml:"withinThirtyDayNum,omitempty"`
	// Data for the last three days
	//
	// example:
	//
	// 300
	WithinThreeDayNum *string `json:"withinThreeDayNum,omitempty" xml:"withinThreeDayNum,omitempty"`
	// Data for yesterday
	//
	// example:
	//
	// 200
	YesterdayNum *int64 `json:"yesterdayNum,omitempty" xml:"yesterdayNum,omitempty"`
}

func (s DescribeTagsFluctuationResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsFluctuationResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeTagsFluctuationResponseBodyResultObject) GetTableName() *string {
	return s.TableName
}

func (s *DescribeTagsFluctuationResponseBodyResultObject) GetTodayNum() *int64 {
	return s.TodayNum
}

func (s *DescribeTagsFluctuationResponseBodyResultObject) GetWithinSevenDayNum() *string {
	return s.WithinSevenDayNum
}

func (s *DescribeTagsFluctuationResponseBodyResultObject) GetWithinThirtyDayNum() *string {
	return s.WithinThirtyDayNum
}

func (s *DescribeTagsFluctuationResponseBodyResultObject) GetWithinThreeDayNum() *string {
	return s.WithinThreeDayNum
}

func (s *DescribeTagsFluctuationResponseBodyResultObject) GetYesterdayNum() *int64 {
	return s.YesterdayNum
}

func (s *DescribeTagsFluctuationResponseBodyResultObject) SetTableName(v string) *DescribeTagsFluctuationResponseBodyResultObject {
	s.TableName = &v
	return s
}

func (s *DescribeTagsFluctuationResponseBodyResultObject) SetTodayNum(v int64) *DescribeTagsFluctuationResponseBodyResultObject {
	s.TodayNum = &v
	return s
}

func (s *DescribeTagsFluctuationResponseBodyResultObject) SetWithinSevenDayNum(v string) *DescribeTagsFluctuationResponseBodyResultObject {
	s.WithinSevenDayNum = &v
	return s
}

func (s *DescribeTagsFluctuationResponseBodyResultObject) SetWithinThirtyDayNum(v string) *DescribeTagsFluctuationResponseBodyResultObject {
	s.WithinThirtyDayNum = &v
	return s
}

func (s *DescribeTagsFluctuationResponseBodyResultObject) SetWithinThreeDayNum(v string) *DescribeTagsFluctuationResponseBodyResultObject {
	s.WithinThreeDayNum = &v
	return s
}

func (s *DescribeTagsFluctuationResponseBodyResultObject) SetYesterdayNum(v int64) *DescribeTagsFluctuationResponseBodyResultObject {
	s.YesterdayNum = &v
	return s
}

func (s *DescribeTagsFluctuationResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
