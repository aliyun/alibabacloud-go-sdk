// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInitDigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeInitDigResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeInitDigResponseBodyResultObject) *DescribeInitDigResponseBody
	GetResultObject() *DescribeInitDigResponseBodyResultObject
}

type DescribeInitDigResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information
	ResultObject *DescribeInitDigResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DescribeInitDigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInitDigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInitDigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInitDigResponseBody) GetResultObject() *DescribeInitDigResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeInitDigResponseBody) SetRequestId(v string) *DescribeInitDigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInitDigResponseBody) SetResultObject(v *DescribeInitDigResponseBodyResultObject) *DescribeInitDigResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeInitDigResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInitDigResponseBodyResultObject struct {
	// Prompt information
	//
	// example:
	//
	// 阿里云内部测试
	TipInfo *string `json:"TipInfo,omitempty" xml:"TipInfo,omitempty"`
}

func (s DescribeInitDigResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeInitDigResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeInitDigResponseBodyResultObject) GetTipInfo() *string {
	return s.TipInfo
}

func (s *DescribeInitDigResponseBodyResultObject) SetTipInfo(v string) *DescribeInitDigResponseBodyResultObject {
	s.TipInfo = &v
	return s
}

func (s *DescribeInitDigResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
