// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsEmptyCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableCount(v int32) *DescribeSlsEmptyCountResponseBody
	GetAvailableCount() *int32
	SetRequestId(v string) *DescribeSlsEmptyCountResponseBody
	GetRequestId() *string
}

type DescribeSlsEmptyCountResponseBody struct {
	// example:
	//
	// 0
	AvailableCount *int32 `json:"AvailableCount,omitempty" xml:"AvailableCount,omitempty"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSlsEmptyCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsEmptyCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsEmptyCountResponseBody) GetAvailableCount() *int32 {
	return s.AvailableCount
}

func (s *DescribeSlsEmptyCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlsEmptyCountResponseBody) SetAvailableCount(v int32) *DescribeSlsEmptyCountResponseBody {
	s.AvailableCount = &v
	return s
}

func (s *DescribeSlsEmptyCountResponseBody) SetRequestId(v string) *DescribeSlsEmptyCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsEmptyCountResponseBody) Validate() error {
	return dara.Validate(s)
}
