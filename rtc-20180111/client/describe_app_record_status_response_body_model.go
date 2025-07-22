// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppRecordStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAppRecordStatusResponseBody
	GetRequestId() *string
	SetResult(v string) *DescribeAppRecordStatusResponseBody
	GetResult() *string
}

type DescribeAppRecordStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// disable
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeAppRecordStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppRecordStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppRecordStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppRecordStatusResponseBody) GetResult() *string {
	return s.Result
}

func (s *DescribeAppRecordStatusResponseBody) SetRequestId(v string) *DescribeAppRecordStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppRecordStatusResponseBody) SetResult(v string) *DescribeAppRecordStatusResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeAppRecordStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
