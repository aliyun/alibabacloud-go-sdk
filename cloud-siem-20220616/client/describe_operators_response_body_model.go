// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperatorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeOperatorsResponseBody
	GetCode() *int32
	SetData(v []*DescribeOperatorsResponseBodyData) *DescribeOperatorsResponseBody
	GetData() []*DescribeOperatorsResponseBodyData
	SetMessage(v string) *DescribeOperatorsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeOperatorsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeOperatorsResponseBody
	GetSuccess() *bool
}

type DescribeOperatorsResponseBody struct {
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
	Data []*DescribeOperatorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s DescribeOperatorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOperatorsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeOperatorsResponseBody) GetData() []*DescribeOperatorsResponseBodyData {
	return s.Data
}

func (s *DescribeOperatorsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeOperatorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOperatorsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeOperatorsResponseBody) SetCode(v int32) *DescribeOperatorsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeOperatorsResponseBody) SetData(v []*DescribeOperatorsResponseBodyData) *DescribeOperatorsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOperatorsResponseBody) SetMessage(v string) *DescribeOperatorsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeOperatorsResponseBody) SetRequestId(v string) *DescribeOperatorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOperatorsResponseBody) SetSuccess(v bool) *DescribeOperatorsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeOperatorsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeOperatorsResponseBodyData struct {
	// The position of the operator in the operator list.
	//
	// example:
	//
	// 3
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The operator.
	//
	// example:
	//
	// <=
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The description of the operator in Chinese.
	//
	// example:
	//
	// arger than or equal to
	OperatorDescCn *string `json:"OperatorDescCn,omitempty" xml:"OperatorDescCn,omitempty"`
	// The description of the operator in English.
	//
	// example:
	//
	// larger than or equal to
	OperatorDescEn *string `json:"OperatorDescEn,omitempty" xml:"OperatorDescEn,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// <=
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// The data types that are supported by the operator. The data types are separated by commas (,).
	//
	// example:
	//
	// varchar
	SupportDataType *string `json:"SupportDataType,omitempty" xml:"SupportDataType,omitempty"`
	// The scenarios that are supported by the operator. Multiple scenarios are separated by commas (,), such as AGGREGATE scenarios. By default, this parameter is empty.
	//
	// example:
	//
	// [AGGREGATE]
	SupportTag []*string `json:"SupportTag,omitempty" xml:"SupportTag,omitempty" type:"Repeated"`
}

func (s DescribeOperatorsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOperatorsResponseBodyData) GetIndex() *int32 {
	return s.Index
}

func (s *DescribeOperatorsResponseBodyData) GetOperator() *string {
	return s.Operator
}

func (s *DescribeOperatorsResponseBodyData) GetOperatorDescCn() *string {
	return s.OperatorDescCn
}

func (s *DescribeOperatorsResponseBodyData) GetOperatorDescEn() *string {
	return s.OperatorDescEn
}

func (s *DescribeOperatorsResponseBodyData) GetOperatorName() *string {
	return s.OperatorName
}

func (s *DescribeOperatorsResponseBodyData) GetSupportDataType() *string {
	return s.SupportDataType
}

func (s *DescribeOperatorsResponseBodyData) GetSupportTag() []*string {
	return s.SupportTag
}

func (s *DescribeOperatorsResponseBodyData) SetIndex(v int32) *DescribeOperatorsResponseBodyData {
	s.Index = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetOperator(v string) *DescribeOperatorsResponseBodyData {
	s.Operator = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetOperatorDescCn(v string) *DescribeOperatorsResponseBodyData {
	s.OperatorDescCn = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetOperatorDescEn(v string) *DescribeOperatorsResponseBodyData {
	s.OperatorDescEn = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetOperatorName(v string) *DescribeOperatorsResponseBodyData {
	s.OperatorName = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetSupportDataType(v string) *DescribeOperatorsResponseBodyData {
	s.SupportDataType = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetSupportTag(v []*string) *DescribeOperatorsResponseBodyData {
	s.SupportTag = v
	return s
}

func (s *DescribeOperatorsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
