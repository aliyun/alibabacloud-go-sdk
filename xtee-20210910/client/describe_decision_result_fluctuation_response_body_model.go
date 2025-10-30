// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDecisionResultFluctuationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDecisionResultFluctuationResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeDecisionResultFluctuationResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeDecisionResultFluctuationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDecisionResultFluctuationResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeDecisionResultFluctuationResponseBodyResultObject) *DescribeDecisionResultFluctuationResponseBody
	GetResultObject() []*DescribeDecisionResultFluctuationResponseBodyResultObject
	SetSuccess(v bool) *DescribeDecisionResultFluctuationResponseBody
	GetSuccess() *bool
}

type DescribeDecisionResultFluctuationResponseBody struct {
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
	// Error details
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
	ResultObject []*DescribeDecisionResultFluctuationResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeDecisionResultFluctuationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDecisionResultFluctuationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDecisionResultFluctuationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDecisionResultFluctuationResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeDecisionResultFluctuationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDecisionResultFluctuationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDecisionResultFluctuationResponseBody) GetResultObject() []*DescribeDecisionResultFluctuationResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeDecisionResultFluctuationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDecisionResultFluctuationResponseBody) SetCode(v string) *DescribeDecisionResultFluctuationResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDecisionResultFluctuationResponseBody) SetHttpStatusCode(v string) *DescribeDecisionResultFluctuationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDecisionResultFluctuationResponseBody) SetMessage(v string) *DescribeDecisionResultFluctuationResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDecisionResultFluctuationResponseBody) SetRequestId(v string) *DescribeDecisionResultFluctuationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDecisionResultFluctuationResponseBody) SetResultObject(v []*DescribeDecisionResultFluctuationResponseBodyResultObject) *DescribeDecisionResultFluctuationResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeDecisionResultFluctuationResponseBody) SetSuccess(v bool) *DescribeDecisionResultFluctuationResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDecisionResultFluctuationResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDecisionResultFluctuationResponseBodyResultObject struct {
	// Execution status.
	//
	// example:
	//
	// PASS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Today\\"s count
	//
	// example:
	//
	// 100
	TodayNum *int64 `json:"todayNum,omitempty" xml:"todayNum,omitempty"`
	// Count within the last seven days
	//
	// example:
	//
	// 600
	WithinSevenDayNum *string `json:"withinSevenDayNum,omitempty" xml:"withinSevenDayNum,omitempty"`
	// Count within the last thirty days
	//
	// example:
	//
	// 1200
	WithinThirtyDayNum *string `json:"withinThirtyDayNum,omitempty" xml:"withinThirtyDayNum,omitempty"`
	// Count within the last three days
	//
	// example:
	//
	// 300
	WithinThreeDayNum *string `json:"withinThreeDayNum,omitempty" xml:"withinThreeDayNum,omitempty"`
	// Yesterday\\"s count
	//
	// example:
	//
	// 200
	YesterdayNum *int64 `json:"yesterdayNum,omitempty" xml:"yesterdayNum,omitempty"`
}

func (s DescribeDecisionResultFluctuationResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeDecisionResultFluctuationResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeDecisionResultFluctuationResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DescribeDecisionResultFluctuationResponseBodyResultObject) GetTodayNum() *int64 {
	return s.TodayNum
}

func (s *DescribeDecisionResultFluctuationResponseBodyResultObject) GetWithinSevenDayNum() *string {
	return s.WithinSevenDayNum
}

func (s *DescribeDecisionResultFluctuationResponseBodyResultObject) GetWithinThirtyDayNum() *string {
	return s.WithinThirtyDayNum
}

func (s *DescribeDecisionResultFluctuationResponseBodyResultObject) GetWithinThreeDayNum() *string {
	return s.WithinThreeDayNum
}

func (s *DescribeDecisionResultFluctuationResponseBodyResultObject) GetYesterdayNum() *int64 {
	return s.YesterdayNum
}

func (s *DescribeDecisionResultFluctuationResponseBodyResultObject) SetStatus(v string) *DescribeDecisionResultFluctuationResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DescribeDecisionResultFluctuationResponseBodyResultObject) SetTodayNum(v int64) *DescribeDecisionResultFluctuationResponseBodyResultObject {
	s.TodayNum = &v
	return s
}

func (s *DescribeDecisionResultFluctuationResponseBodyResultObject) SetWithinSevenDayNum(v string) *DescribeDecisionResultFluctuationResponseBodyResultObject {
	s.WithinSevenDayNum = &v
	return s
}

func (s *DescribeDecisionResultFluctuationResponseBodyResultObject) SetWithinThirtyDayNum(v string) *DescribeDecisionResultFluctuationResponseBodyResultObject {
	s.WithinThirtyDayNum = &v
	return s
}

func (s *DescribeDecisionResultFluctuationResponseBodyResultObject) SetWithinThreeDayNum(v string) *DescribeDecisionResultFluctuationResponseBodyResultObject {
	s.WithinThreeDayNum = &v
	return s
}

func (s *DescribeDecisionResultFluctuationResponseBodyResultObject) SetYesterdayNum(v int64) *DescribeDecisionResultFluctuationResponseBodyResultObject {
	s.YesterdayNum = &v
	return s
}

func (s *DescribeDecisionResultFluctuationResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
