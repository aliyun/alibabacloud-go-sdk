// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizeRuleTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeCustomizeRuleTestResponseBody
	GetCode() *int32
	SetData(v *DescribeCustomizeRuleTestResponseBodyData) *DescribeCustomizeRuleTestResponseBody
	GetData() *DescribeCustomizeRuleTestResponseBodyData
	SetMessage(v string) *DescribeCustomizeRuleTestResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCustomizeRuleTestResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCustomizeRuleTestResponseBody
	GetSuccess() *bool
}

type DescribeCustomizeRuleTestResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeCustomizeRuleTestResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomizeRuleTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeRuleTestResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeCustomizeRuleTestResponseBody) GetData() *DescribeCustomizeRuleTestResponseBodyData {
	return s.Data
}

func (s *DescribeCustomizeRuleTestResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCustomizeRuleTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomizeRuleTestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCustomizeRuleTestResponseBody) SetCode(v int32) *DescribeCustomizeRuleTestResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBody) SetData(v *DescribeCustomizeRuleTestResponseBodyData) *DescribeCustomizeRuleTestResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBody) SetMessage(v string) *DescribeCustomizeRuleTestResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBody) SetRequestId(v string) *DescribeCustomizeRuleTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBody) SetSuccess(v bool) *DescribeCustomizeRuleTestResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCustomizeRuleTestResponseBodyData struct {
	// The ID of the rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The historical data that is used in the simulation test.
	//
	// example:
	//
	// [{"key1":"value1","key2":"value2","key3":"value3","key4":"value4","key5":"value5"}]
	SimulateData *string `json:"SimulateData,omitempty" xml:"SimulateData,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- 0: The rule is in the initial state.
	//
	// 	- 10: The simulation data is tested.
	//
	// 	- 15: The business data is being tested.
	//
	// 	- 20: The business data test ends.
	//
	// 	- 100: The rule takes effect.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCustomizeRuleTestResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeRuleTestResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DescribeCustomizeRuleTestResponseBodyData) GetSimulateData() *string {
	return s.SimulateData
}

func (s *DescribeCustomizeRuleTestResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCustomizeRuleTestResponseBodyData) SetId(v int64) *DescribeCustomizeRuleTestResponseBodyData {
	s.Id = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBodyData) SetSimulateData(v string) *DescribeCustomizeRuleTestResponseBodyData {
	s.SimulateData = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBodyData) SetStatus(v int32) *DescribeCustomizeRuleTestResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBodyData) Validate() error {
	return dara.Validate(s)
}
