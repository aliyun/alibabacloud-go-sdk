// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParamByEventCodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeParamByEventCodesResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeParamByEventCodesResponseBodyResultObject) *DescribeParamByEventCodesResponseBody
	GetResultObject() []*DescribeParamByEventCodesResponseBodyResultObject
}

type DescribeParamByEventCodesResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject []*DescribeParamByEventCodesResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeParamByEventCodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParamByEventCodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParamByEventCodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParamByEventCodesResponseBody) GetResultObject() []*DescribeParamByEventCodesResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeParamByEventCodesResponseBody) SetRequestId(v string) *DescribeParamByEventCodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParamByEventCodesResponseBody) SetResultObject(v []*DescribeParamByEventCodesResponseBodyResultObject) *DescribeParamByEventCodesResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeParamByEventCodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeParamByEventCodesResponseBodyResultObject struct {
	// Return code.
	//
	// example:
	//
	// age
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Name
	//
	// example:
	//
	// 年龄
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeParamByEventCodesResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeParamByEventCodesResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeParamByEventCodesResponseBodyResultObject) GetCode() *string {
	return s.Code
}

func (s *DescribeParamByEventCodesResponseBodyResultObject) GetName() *string {
	return s.Name
}

func (s *DescribeParamByEventCodesResponseBodyResultObject) SetCode(v string) *DescribeParamByEventCodesResponseBodyResultObject {
	s.Code = &v
	return s
}

func (s *DescribeParamByEventCodesResponseBodyResultObject) SetName(v string) *DescribeParamByEventCodesResponseBodyResultObject {
	s.Name = &v
	return s
}

func (s *DescribeParamByEventCodesResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
