// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceCodeNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeServiceCodeNameResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeServiceCodeNameResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeServiceCodeNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeServiceCodeNameResponseBody
	GetRequestId() *string
	SetResultObject(v string) *DescribeServiceCodeNameResponseBody
	GetResultObject() *string
}

type DescribeServiceCodeNameResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
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
	// Return result.
	//
	// example:
	//
	// true
	ResultObject *string `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s DescribeServiceCodeNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceCodeNameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceCodeNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeServiceCodeNameResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeServiceCodeNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeServiceCodeNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceCodeNameResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *DescribeServiceCodeNameResponseBody) SetCode(v string) *DescribeServiceCodeNameResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeServiceCodeNameResponseBody) SetHttpStatusCode(v string) *DescribeServiceCodeNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeServiceCodeNameResponseBody) SetMessage(v string) *DescribeServiceCodeNameResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeServiceCodeNameResponseBody) SetRequestId(v string) *DescribeServiceCodeNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceCodeNameResponseBody) SetResultObject(v string) *DescribeServiceCodeNameResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeServiceCodeNameResponseBody) Validate() error {
	return dara.Validate(s)
}
