// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebAccessLogEmptyCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableCount(v int32) *DescribeWebAccessLogEmptyCountResponseBody
	GetAvailableCount() *int32
	SetRequestId(v string) *DescribeWebAccessLogEmptyCountResponseBody
	GetRequestId() *string
}

type DescribeWebAccessLogEmptyCountResponseBody struct {
	// The remaining quota that you can clear the Logstore.
	//
	// example:
	//
	// 10
	AvailableCount *int32 `json:"AvailableCount,omitempty" xml:"AvailableCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWebAccessLogEmptyCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAccessLogEmptyCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogEmptyCountResponseBody) GetAvailableCount() *int32 {
	return s.AvailableCount
}

func (s *DescribeWebAccessLogEmptyCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebAccessLogEmptyCountResponseBody) SetAvailableCount(v int32) *DescribeWebAccessLogEmptyCountResponseBody {
	s.AvailableCount = &v
	return s
}

func (s *DescribeWebAccessLogEmptyCountResponseBody) SetRequestId(v string) *DescribeWebAccessLogEmptyCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebAccessLogEmptyCountResponseBody) Validate() error {
	return dara.Validate(s)
}
