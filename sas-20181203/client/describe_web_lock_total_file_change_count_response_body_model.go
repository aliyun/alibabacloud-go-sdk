// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockTotalFileChangeCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeWebLockTotalFileChangeCountResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeWebLockTotalFileChangeCountResponseBody
	GetTotalCount() *int64
}

type DescribeWebLockTotalFileChangeCountResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// E70074C8-DFB4-44C5-96C7-909DD231D68A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of times that the files protected by web tamper proofing are changed.
	//
	// example:
	//
	// 200
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWebLockTotalFileChangeCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockTotalFileChangeCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebLockTotalFileChangeCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebLockTotalFileChangeCountResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeWebLockTotalFileChangeCountResponseBody) SetRequestId(v string) *DescribeWebLockTotalFileChangeCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebLockTotalFileChangeCountResponseBody) SetTotalCount(v int64) *DescribeWebLockTotalFileChangeCountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebLockTotalFileChangeCountResponseBody) Validate() error {
	return dara.Validate(s)
}
