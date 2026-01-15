// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModuleServiceExistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeModuleServiceExistResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *DescribeModuleServiceExistResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *DescribeModuleServiceExistResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeModuleServiceExistResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *DescribeModuleServiceExistResponseBody
	GetSuccess() *bool
}

type DescribeModuleServiceExistResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeModuleServiceExistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModuleServiceExistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModuleServiceExistResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeModuleServiceExistResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DescribeModuleServiceExistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModuleServiceExistResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeModuleServiceExistResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeModuleServiceExistResponseBody) SetCode(v int64) *DescribeModuleServiceExistResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeModuleServiceExistResponseBody) SetHttpStatusCode(v int64) *DescribeModuleServiceExistResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeModuleServiceExistResponseBody) SetRequestId(v string) *DescribeModuleServiceExistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModuleServiceExistResponseBody) SetResultObject(v bool) *DescribeModuleServiceExistResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeModuleServiceExistResponseBody) SetSuccess(v bool) *DescribeModuleServiceExistResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeModuleServiceExistResponseBody) Validate() error {
	return dara.Validate(s)
}
