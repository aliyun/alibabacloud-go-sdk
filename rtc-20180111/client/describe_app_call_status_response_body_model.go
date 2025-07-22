// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppCallStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAppCallStatusResponseBody
	GetRequestId() *string
	SetResult(v string) *DescribeAppCallStatusResponseBody
	GetResult() *string
}

type DescribeAppCallStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 20A6D1E3-1F5F-5440-A4F1-EC7831646FE4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// enable
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeAppCallStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppCallStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppCallStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppCallStatusResponseBody) GetResult() *string {
	return s.Result
}

func (s *DescribeAppCallStatusResponseBody) SetRequestId(v string) *DescribeAppCallStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppCallStatusResponseBody) SetResult(v string) *DescribeAppCallStatusResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeAppCallStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
