// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHitRuleFluctuationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHitRuleFluctuationResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeHitRuleFluctuationResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeHitRuleFluctuationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeHitRuleFluctuationResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeHitRuleFluctuationResponseBodyResultObject) *DescribeHitRuleFluctuationResponseBody
	GetResultObject() []*DescribeHitRuleFluctuationResponseBodyResultObject
	SetSuccess(v bool) *DescribeHitRuleFluctuationResponseBody
	GetSuccess() *bool
}

type DescribeHitRuleFluctuationResponseBody struct {
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
	ResultObject []*DescribeHitRuleFluctuationResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeHitRuleFluctuationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHitRuleFluctuationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHitRuleFluctuationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHitRuleFluctuationResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeHitRuleFluctuationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHitRuleFluctuationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHitRuleFluctuationResponseBody) GetResultObject() []*DescribeHitRuleFluctuationResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeHitRuleFluctuationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeHitRuleFluctuationResponseBody) SetCode(v string) *DescribeHitRuleFluctuationResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHitRuleFluctuationResponseBody) SetHttpStatusCode(v string) *DescribeHitRuleFluctuationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeHitRuleFluctuationResponseBody) SetMessage(v string) *DescribeHitRuleFluctuationResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHitRuleFluctuationResponseBody) SetRequestId(v string) *DescribeHitRuleFluctuationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHitRuleFluctuationResponseBody) SetResultObject(v []*DescribeHitRuleFluctuationResponseBodyResultObject) *DescribeHitRuleFluctuationResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeHitRuleFluctuationResponseBody) SetSuccess(v bool) *DescribeHitRuleFluctuationResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHitRuleFluctuationResponseBody) Validate() error {
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

type DescribeHitRuleFluctuationResponseBodyResultObject struct {
	// Policy ID
	//
	// example:
	//
	// 115019
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Policy name
	//
	// example:
	//
	// 营销风险识别
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Today\\"s count
	//
	// example:
	//
	// 100
	TodayNum *int64 `json:"todayNum,omitempty" xml:"todayNum,omitempty"`
	// Count within seven days
	//
	// example:
	//
	// 600
	WithinSevenDayNum *string `json:"withinSevenDayNum,omitempty" xml:"withinSevenDayNum,omitempty"`
	// Count within thirty days
	//
	// example:
	//
	// 1200
	WithinThirtyDayNum *string `json:"withinThirtyDayNum,omitempty" xml:"withinThirtyDayNum,omitempty"`
	// Count within three days
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

func (s DescribeHitRuleFluctuationResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeHitRuleFluctuationResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeHitRuleFluctuationResponseBodyResultObject) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeHitRuleFluctuationResponseBodyResultObject) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeHitRuleFluctuationResponseBodyResultObject) GetTodayNum() *int64 {
	return s.TodayNum
}

func (s *DescribeHitRuleFluctuationResponseBodyResultObject) GetWithinSevenDayNum() *string {
	return s.WithinSevenDayNum
}

func (s *DescribeHitRuleFluctuationResponseBodyResultObject) GetWithinThirtyDayNum() *string {
	return s.WithinThirtyDayNum
}

func (s *DescribeHitRuleFluctuationResponseBodyResultObject) GetWithinThreeDayNum() *string {
	return s.WithinThreeDayNum
}

func (s *DescribeHitRuleFluctuationResponseBodyResultObject) GetYesterdayNum() *int64 {
	return s.YesterdayNum
}

func (s *DescribeHitRuleFluctuationResponseBodyResultObject) SetRuleId(v string) *DescribeHitRuleFluctuationResponseBodyResultObject {
	s.RuleId = &v
	return s
}

func (s *DescribeHitRuleFluctuationResponseBodyResultObject) SetRuleName(v string) *DescribeHitRuleFluctuationResponseBodyResultObject {
	s.RuleName = &v
	return s
}

func (s *DescribeHitRuleFluctuationResponseBodyResultObject) SetTodayNum(v int64) *DescribeHitRuleFluctuationResponseBodyResultObject {
	s.TodayNum = &v
	return s
}

func (s *DescribeHitRuleFluctuationResponseBodyResultObject) SetWithinSevenDayNum(v string) *DescribeHitRuleFluctuationResponseBodyResultObject {
	s.WithinSevenDayNum = &v
	return s
}

func (s *DescribeHitRuleFluctuationResponseBodyResultObject) SetWithinThirtyDayNum(v string) *DescribeHitRuleFluctuationResponseBodyResultObject {
	s.WithinThirtyDayNum = &v
	return s
}

func (s *DescribeHitRuleFluctuationResponseBodyResultObject) SetWithinThreeDayNum(v string) *DescribeHitRuleFluctuationResponseBodyResultObject {
	s.WithinThreeDayNum = &v
	return s
}

func (s *DescribeHitRuleFluctuationResponseBodyResultObject) SetYesterdayNum(v int64) *DescribeHitRuleFluctuationResponseBodyResultObject {
	s.YesterdayNum = &v
	return s
}

func (s *DescribeHitRuleFluctuationResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
