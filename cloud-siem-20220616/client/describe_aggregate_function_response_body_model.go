// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAggregateFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAggregateFunctionResponseBody
	GetCode() *int32
	SetData(v []*DescribeAggregateFunctionResponseBodyData) *DescribeAggregateFunctionResponseBody
	GetData() []*DescribeAggregateFunctionResponseBodyData
	SetMessage(v string) *DescribeAggregateFunctionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAggregateFunctionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAggregateFunctionResponseBody
	GetSuccess() *bool
}

type DescribeAggregateFunctionResponseBody struct {
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
	Data []*DescribeAggregateFunctionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s DescribeAggregateFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAggregateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAggregateFunctionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAggregateFunctionResponseBody) GetData() []*DescribeAggregateFunctionResponseBodyData {
	return s.Data
}

func (s *DescribeAggregateFunctionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAggregateFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAggregateFunctionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAggregateFunctionResponseBody) SetCode(v int32) *DescribeAggregateFunctionResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAggregateFunctionResponseBody) SetData(v []*DescribeAggregateFunctionResponseBodyData) *DescribeAggregateFunctionResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAggregateFunctionResponseBody) SetMessage(v string) *DescribeAggregateFunctionResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAggregateFunctionResponseBody) SetRequestId(v string) *DescribeAggregateFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAggregateFunctionResponseBody) SetSuccess(v bool) *DescribeAggregateFunctionResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAggregateFunctionResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAggregateFunctionResponseBodyData struct {
	// The aggregate function.
	//
	// example:
	//
	// count
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The display name of the aggregate function.
	//
	// example:
	//
	// Count
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s DescribeAggregateFunctionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAggregateFunctionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAggregateFunctionResponseBodyData) GetFunction() *string {
	return s.Function
}

func (s *DescribeAggregateFunctionResponseBodyData) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeAggregateFunctionResponseBodyData) SetFunction(v string) *DescribeAggregateFunctionResponseBodyData {
	s.Function = &v
	return s
}

func (s *DescribeAggregateFunctionResponseBodyData) SetFunctionName(v string) *DescribeAggregateFunctionResponseBodyData {
	s.FunctionName = &v
	return s
}

func (s *DescribeAggregateFunctionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
