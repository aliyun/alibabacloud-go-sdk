// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeServiceListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeServiceListResponseBodyResultObject) *DescribeServiceListResponseBody
	GetResultObject() []*DescribeServiceListResponseBodyResultObject
}

type DescribeServiceListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject []*DescribeServiceListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeServiceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceListResponseBody) GetResultObject() []*DescribeServiceListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeServiceListResponseBody) SetRequestId(v string) *DescribeServiceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceListResponseBody) SetResultObject(v []*DescribeServiceListResponseBodyResultObject) *DescribeServiceListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeServiceListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceListResponseBodyResultObject struct {
	// Service code
	//
	// example:
	//
	// coupon_abuse_detection
	ServiceCode *string `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	// Service name.
	//
	// example:
	//
	// coupon_abuse_detection
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s DescribeServiceListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeServiceListResponseBodyResultObject) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeServiceListResponseBodyResultObject) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeServiceListResponseBodyResultObject) SetServiceCode(v string) *DescribeServiceListResponseBodyResultObject {
	s.ServiceCode = &v
	return s
}

func (s *DescribeServiceListResponseBodyResultObject) SetServiceName(v string) *DescribeServiceListResponseBodyResultObject {
	s.ServiceName = &v
	return s
}

func (s *DescribeServiceListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
