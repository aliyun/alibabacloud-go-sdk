// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMenuPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeMenuPermissionResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeMenuPermissionResponseBody
	GetResultObject() *bool
}

type DescribeMenuPermissionResponseBody struct {
	// Request ID
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeMenuPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMenuPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMenuPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMenuPermissionResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeMenuPermissionResponseBody) SetRequestId(v string) *DescribeMenuPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMenuPermissionResponseBody) SetResultObject(v bool) *DescribeMenuPermissionResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeMenuPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
