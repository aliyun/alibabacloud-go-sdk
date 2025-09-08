// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DescribeAuthResponseBody
	GetData() *bool
	SetRequestId(v string) *DescribeAuthResponseBody
	GetRequestId() *string
}

type DescribeAuthResponseBody struct {
	// Indicates whether the SIEM system is granted the required permissions. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F539347-7D9A-51EA-8ABF-5D5507045C5C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuthResponseBody) GetData() *bool {
	return s.Data
}

func (s *DescribeAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuthResponseBody) SetData(v bool) *DescribeAuthResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeAuthResponseBody) SetRequestId(v string) *DescribeAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
