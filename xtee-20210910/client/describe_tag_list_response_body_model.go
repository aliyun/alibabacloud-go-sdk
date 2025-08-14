// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeTagListResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeTagListResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeTagListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeTagListResponseBody
	GetRequestId() *string
	SetResultObject(v string) *DescribeTagListResponseBody
	GetResultObject() *string
}

type DescribeTagListResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Result object.
	//
	// example:
	//
	// true
	ResultObject *string `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s DescribeTagListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeTagListResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeTagListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTagListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagListResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *DescribeTagListResponseBody) SetCode(v string) *DescribeTagListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTagListResponseBody) SetHttpStatusCode(v string) *DescribeTagListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeTagListResponseBody) SetMessage(v string) *DescribeTagListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTagListResponseBody) SetRequestId(v string) *DescribeTagListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagListResponseBody) SetResultObject(v string) *DescribeTagListResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeTagListResponseBody) Validate() error {
	return dara.Validate(s)
}
