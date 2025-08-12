// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagValueListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeTagValueListResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeTagValueListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeTagValueListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeTagValueListResponseBody
	GetSuccess() *bool
	SetTagValues(v *DescribeTagValueListResponseBodyTagValues) *DescribeTagValueListResponseBody
	GetTagValues() *DescribeTagValueListResponseBodyTagValues
}

type DescribeTagValueListResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B04B8CF3-4489-432D-83BA-6F128E4F2295
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
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TagValues *DescribeTagValueListResponseBodyTagValues `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Struct"`
}

func (s DescribeTagValueListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagValueListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagValueListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeTagValueListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTagValueListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagValueListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTagValueListResponseBody) GetTagValues() *DescribeTagValueListResponseBodyTagValues {
	return s.TagValues
}

func (s *DescribeTagValueListResponseBody) SetCode(v string) *DescribeTagValueListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTagValueListResponseBody) SetMessage(v string) *DescribeTagValueListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTagValueListResponseBody) SetRequestId(v string) *DescribeTagValueListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagValueListResponseBody) SetSuccess(v bool) *DescribeTagValueListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTagValueListResponseBody) SetTagValues(v *DescribeTagValueListResponseBodyTagValues) *DescribeTagValueListResponseBody {
	s.TagValues = v
	return s
}

func (s *DescribeTagValueListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTagValueListResponseBodyTagValues struct {
	TagValue []*string `json:"TagValue,omitempty" xml:"TagValue,omitempty" type:"Repeated"`
}

func (s DescribeTagValueListResponseBodyTagValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagValueListResponseBodyTagValues) GoString() string {
	return s.String()
}

func (s *DescribeTagValueListResponseBodyTagValues) GetTagValue() []*string {
	return s.TagValue
}

func (s *DescribeTagValueListResponseBodyTagValues) SetTagValue(v []*string) *DescribeTagValueListResponseBodyTagValues {
	s.TagValue = v
	return s
}

func (s *DescribeTagValueListResponseBodyTagValues) Validate() error {
	return dara.Validate(s)
}
