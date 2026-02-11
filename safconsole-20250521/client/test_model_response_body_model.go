// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *TestModelResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *TestModelResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *TestModelResponseBody
	GetRequestId() *string
	SetResultObject(v *TestModelResponseBodyResultObject) *TestModelResponseBody
	GetResultObject() *TestModelResponseBodyResultObject
	SetSuccess(v bool) *TestModelResponseBody
	GetSuccess() *bool
}

type TestModelResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	ResultObject *TestModelResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	// Indicates whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TestModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestModelResponseBody) GoString() string {
	return s.String()
}

func (s *TestModelResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *TestModelResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *TestModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestModelResponseBody) GetResultObject() *TestModelResponseBodyResultObject {
	return s.ResultObject
}

func (s *TestModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TestModelResponseBody) SetCode(v int64) *TestModelResponseBody {
	s.Code = &v
	return s
}

func (s *TestModelResponseBody) SetHttpStatusCode(v int64) *TestModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TestModelResponseBody) SetRequestId(v string) *TestModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestModelResponseBody) SetResultObject(v *TestModelResponseBodyResultObject) *TestModelResponseBody {
	s.ResultObject = v
	return s
}

func (s *TestModelResponseBody) SetSuccess(v bool) *TestModelResponseBody {
	s.Success = &v
	return s
}

func (s *TestModelResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TestModelResponseBodyResultObject struct {
	// Number of consistencies.
	//
	// example:
	//
	// 49.0
	ConsistencyCount *int64 `json:"ConsistencyCount,omitempty" xml:"ConsistencyCount,omitempty"`
	// Consistency rate.
	//
	// example:
	//
	// 98.0
	ConsistencyRate *float64 `json:"ConsistencyRate,omitempty" xml:"ConsistencyRate,omitempty"`
	// Test results.
	TestResult []*TestModelResponseBodyResultObjectTestResult `json:"TestResult,omitempty" xml:"TestResult,omitempty" type:"Repeated"`
	// Total number.
	//
	// example:
	//
	// 50
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s TestModelResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s TestModelResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *TestModelResponseBodyResultObject) GetConsistencyCount() *int64 {
	return s.ConsistencyCount
}

func (s *TestModelResponseBodyResultObject) GetConsistencyRate() *float64 {
	return s.ConsistencyRate
}

func (s *TestModelResponseBodyResultObject) GetTestResult() []*TestModelResponseBodyResultObjectTestResult {
	return s.TestResult
}

func (s *TestModelResponseBodyResultObject) GetTotal() *int64 {
	return s.Total
}

func (s *TestModelResponseBodyResultObject) SetConsistencyCount(v int64) *TestModelResponseBodyResultObject {
	s.ConsistencyCount = &v
	return s
}

func (s *TestModelResponseBodyResultObject) SetConsistencyRate(v float64) *TestModelResponseBodyResultObject {
	s.ConsistencyRate = &v
	return s
}

func (s *TestModelResponseBodyResultObject) SetTestResult(v []*TestModelResponseBodyResultObjectTestResult) *TestModelResponseBodyResultObject {
	s.TestResult = v
	return s
}

func (s *TestModelResponseBodyResultObject) SetTotal(v int64) *TestModelResponseBodyResultObject {
	s.Total = &v
	return s
}

func (s *TestModelResponseBodyResultObject) Validate() error {
	if s.TestResult != nil {
		for _, item := range s.TestResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TestModelResponseBodyResultObjectTestResult struct {
	// Actual output.
	//
	// example:
	//
	// 0.00
	ActualResult *string `json:"ActualResult,omitempty" xml:"ActualResult,omitempty"`
	// Whether it is consistent.
	//
	// example:
	//
	// true
	Consistency *bool `json:"Consistency,omitempty" xml:"Consistency,omitempty"`
	// Test output.
	//
	// example:
	//
	// 0.00
	TestResult *string `json:"TestResult,omitempty" xml:"TestResult,omitempty"`
	// Training output.
	//
	// example:
	//
	// 0.00
	TrainResult *string `json:"TrainResult,omitempty" xml:"TrainResult,omitempty"`
}

func (s TestModelResponseBodyResultObjectTestResult) String() string {
	return dara.Prettify(s)
}

func (s TestModelResponseBodyResultObjectTestResult) GoString() string {
	return s.String()
}

func (s *TestModelResponseBodyResultObjectTestResult) GetActualResult() *string {
	return s.ActualResult
}

func (s *TestModelResponseBodyResultObjectTestResult) GetConsistency() *bool {
	return s.Consistency
}

func (s *TestModelResponseBodyResultObjectTestResult) GetTestResult() *string {
	return s.TestResult
}

func (s *TestModelResponseBodyResultObjectTestResult) GetTrainResult() *string {
	return s.TrainResult
}

func (s *TestModelResponseBodyResultObjectTestResult) SetActualResult(v string) *TestModelResponseBodyResultObjectTestResult {
	s.ActualResult = &v
	return s
}

func (s *TestModelResponseBodyResultObjectTestResult) SetConsistency(v bool) *TestModelResponseBodyResultObjectTestResult {
	s.Consistency = &v
	return s
}

func (s *TestModelResponseBodyResultObjectTestResult) SetTestResult(v string) *TestModelResponseBodyResultObjectTestResult {
	s.TestResult = &v
	return s
}

func (s *TestModelResponseBodyResultObjectTestResult) SetTrainResult(v string) *TestModelResponseBodyResultObjectTestResult {
	s.TrainResult = &v
	return s
}

func (s *TestModelResponseBodyResultObjectTestResult) Validate() error {
	return dara.Validate(s)
}
