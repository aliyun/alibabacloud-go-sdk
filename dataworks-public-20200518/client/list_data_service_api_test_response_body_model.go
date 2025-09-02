// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDataServiceApiTestResponseBodyData) *ListDataServiceApiTestResponseBody
	GetData() []*ListDataServiceApiTestResponseBodyData
	SetRequestId(v string) *ListDataServiceApiTestResponseBody
	GetRequestId() *string
}

type ListDataServiceApiTestResponseBody struct {
	// The list of test records.
	Data []*ListDataServiceApiTestResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// ESDAFWEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataServiceApiTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiTestResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiTestResponseBody) GetData() []*ListDataServiceApiTestResponseBodyData {
	return s.Data
}

func (s *ListDataServiceApiTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServiceApiTestResponseBody) SetData(v []*ListDataServiceApiTestResponseBodyData) *ListDataServiceApiTestResponseBody {
	s.Data = v
	return s
}

func (s *ListDataServiceApiTestResponseBody) SetRequestId(v string) *ListDataServiceApiTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataServiceApiTestResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApiTestResponseBodyData struct {
	// The ID of the DataService Studio API on which the test is performed.
	//
	// example:
	//
	// 2343
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The time that is consumed to complete the test.
	//
	// example:
	//
	// 10
	CostTime *int32 `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	// The time when the test was initiated.
	//
	// example:
	//
	// 1651824913000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The request parameters configured for the test.
	//
	// example:
	//
	// {"name":"test"}
	ParamMap *string `json:"ParamMap,omitempty" xml:"ParamMap,omitempty"`
	// The status code returned for the test. If the test is not complete, this parameter is not returned.
	//
	// example:
	//
	// 0
	RetCode *int64 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	// The result returned for the test.
	//
	// example:
	//
	// {"id":2}
	RetResult *string `json:"RetResult,omitempty" xml:"RetResult,omitempty"`
	// The status of the test. Valid values: RUNNING and FINISHED.
	//
	// example:
	//
	// FINISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the test.
	//
	// example:
	//
	// 123
	TestId *int64 `json:"TestId,omitempty" xml:"TestId,omitempty"`
}

func (s ListDataServiceApiTestResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiTestResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiTestResponseBodyData) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListDataServiceApiTestResponseBodyData) GetCostTime() *int32 {
	return s.CostTime
}

func (s *ListDataServiceApiTestResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataServiceApiTestResponseBodyData) GetParamMap() *string {
	return s.ParamMap
}

func (s *ListDataServiceApiTestResponseBodyData) GetRetCode() *int64 {
	return s.RetCode
}

func (s *ListDataServiceApiTestResponseBodyData) GetRetResult() *string {
	return s.RetResult
}

func (s *ListDataServiceApiTestResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListDataServiceApiTestResponseBodyData) GetTestId() *int64 {
	return s.TestId
}

func (s *ListDataServiceApiTestResponseBodyData) SetApiId(v int64) *ListDataServiceApiTestResponseBodyData {
	s.ApiId = &v
	return s
}

func (s *ListDataServiceApiTestResponseBodyData) SetCostTime(v int32) *ListDataServiceApiTestResponseBodyData {
	s.CostTime = &v
	return s
}

func (s *ListDataServiceApiTestResponseBodyData) SetCreateTime(v int64) *ListDataServiceApiTestResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListDataServiceApiTestResponseBodyData) SetParamMap(v string) *ListDataServiceApiTestResponseBodyData {
	s.ParamMap = &v
	return s
}

func (s *ListDataServiceApiTestResponseBodyData) SetRetCode(v int64) *ListDataServiceApiTestResponseBodyData {
	s.RetCode = &v
	return s
}

func (s *ListDataServiceApiTestResponseBodyData) SetRetResult(v string) *ListDataServiceApiTestResponseBodyData {
	s.RetResult = &v
	return s
}

func (s *ListDataServiceApiTestResponseBodyData) SetStatus(v string) *ListDataServiceApiTestResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListDataServiceApiTestResponseBodyData) SetTestId(v int64) *ListDataServiceApiTestResponseBodyData {
	s.TestId = &v
	return s
}

func (s *ListDataServiceApiTestResponseBodyData) Validate() error {
	return dara.Validate(s)
}
