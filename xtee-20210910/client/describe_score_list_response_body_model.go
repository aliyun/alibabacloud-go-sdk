// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScoreListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeScoreListResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeScoreListResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeScoreListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeScoreListResponseBody
	GetRequestId() *string
	SetResultObject(v string) *DescribeScoreListResponseBody
	GetResultObject() *string
}

type DescribeScoreListResponseBody struct {
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

func (s DescribeScoreListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScoreListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeScoreListResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeScoreListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeScoreListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScoreListResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *DescribeScoreListResponseBody) SetCode(v string) *DescribeScoreListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeScoreListResponseBody) SetHttpStatusCode(v string) *DescribeScoreListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeScoreListResponseBody) SetMessage(v string) *DescribeScoreListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeScoreListResponseBody) SetRequestId(v string) *DescribeScoreListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScoreListResponseBody) SetResultObject(v string) *DescribeScoreListResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeScoreListResponseBody) Validate() error {
	return dara.Validate(s)
}
