// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppLiveStreamStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAppLiveStreamStatusResponseBody
	GetRequestId() *string
	SetResult(v string) *DescribeAppLiveStreamStatusResponseBody
	GetResult() *string
}

type DescribeAppLiveStreamStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// disable
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeAppLiveStreamStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppLiveStreamStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppLiveStreamStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppLiveStreamStatusResponseBody) GetResult() *string {
	return s.Result
}

func (s *DescribeAppLiveStreamStatusResponseBody) SetRequestId(v string) *DescribeAppLiveStreamStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppLiveStreamStatusResponseBody) SetResult(v string) *DescribeAppLiveStreamStatusResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeAppLiveStreamStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
