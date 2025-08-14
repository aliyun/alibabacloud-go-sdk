// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiNameListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeApiNameListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeApiNameListResponseBodyResultObject) *DescribeApiNameListResponseBody
	GetResultObject() []*DescribeApiNameListResponseBodyResultObject
}

type DescribeApiNameListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject []*DescribeApiNameListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeApiNameListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiNameListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiNameListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiNameListResponseBody) GetResultObject() []*DescribeApiNameListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeApiNameListResponseBody) SetRequestId(v string) *DescribeApiNameListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiNameListResponseBody) SetResultObject(v []*DescribeApiNameListResponseBodyResultObject) *DescribeApiNameListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeApiNameListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApiNameListResponseBodyResultObject struct {
	// API ID.
	//
	// example:
	//
	// 33
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// API name.
	//
	// example:
	//
	// ListAuditLog
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
}

func (s DescribeApiNameListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiNameListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeApiNameListResponseBodyResultObject) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiNameListResponseBodyResultObject) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApiNameListResponseBodyResultObject) SetApiId(v string) *DescribeApiNameListResponseBodyResultObject {
	s.ApiId = &v
	return s
}

func (s *DescribeApiNameListResponseBodyResultObject) SetApiName(v string) *DescribeApiNameListResponseBodyResultObject {
	s.ApiName = &v
	return s
}

func (s *DescribeApiNameListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
