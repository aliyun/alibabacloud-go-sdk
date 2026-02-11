// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModuleStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeModuleStatusResponseBody
	GetCode() *int32
	SetHttpStatusCode(v int64) *DescribeModuleStatusResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *DescribeModuleStatusResponseBody
	GetRequestId() *string
	SetResultObject(v string) *DescribeModuleStatusResponseBody
	GetResultObject() *string
	SetSuccess(v bool) *DescribeModuleStatusResponseBody
	GetSuccess() *bool
}

type DescribeModuleStatusResponseBody struct {
	// Status code
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// Returned result.
	//
	// example:
	//
	// true
	ResultObject *string `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// Whether the operation was successful
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeModuleStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModuleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModuleStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeModuleStatusResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DescribeModuleStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModuleStatusResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *DescribeModuleStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeModuleStatusResponseBody) SetCode(v int32) *DescribeModuleStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeModuleStatusResponseBody) SetHttpStatusCode(v int64) *DescribeModuleStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeModuleStatusResponseBody) SetRequestId(v string) *DescribeModuleStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModuleStatusResponseBody) SetResultObject(v string) *DescribeModuleStatusResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeModuleStatusResponseBody) SetSuccess(v bool) *DescribeModuleStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeModuleStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
