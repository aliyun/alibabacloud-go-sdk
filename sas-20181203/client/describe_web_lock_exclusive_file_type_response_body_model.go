// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockExclusiveFileTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExclusiveFileType(v []*string) *DescribeWebLockExclusiveFileTypeResponseBody
	GetExclusiveFileType() []*string
	SetRequestId(v string) *DescribeWebLockExclusiveFileTypeResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWebLockExclusiveFileTypeResponseBody
	GetTotalCount() *int32
}

type DescribeWebLockExclusiveFileTypeResponseBody struct {
	// An array that consists of the types of the files that are excluded from web tamper proofing.
	ExclusiveFileType []*string `json:"ExclusiveFileType,omitempty" xml:"ExclusiveFileType,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9CCD7D51-5E81-5FF5-BD74-813DDD248430
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of types of the files that are excluded from web tamper proofing.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWebLockExclusiveFileTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockExclusiveFileTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebLockExclusiveFileTypeResponseBody) GetExclusiveFileType() []*string {
	return s.ExclusiveFileType
}

func (s *DescribeWebLockExclusiveFileTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebLockExclusiveFileTypeResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWebLockExclusiveFileTypeResponseBody) SetExclusiveFileType(v []*string) *DescribeWebLockExclusiveFileTypeResponseBody {
	s.ExclusiveFileType = v
	return s
}

func (s *DescribeWebLockExclusiveFileTypeResponseBody) SetRequestId(v string) *DescribeWebLockExclusiveFileTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebLockExclusiveFileTypeResponseBody) SetTotalCount(v int32) *DescribeWebLockExclusiveFileTypeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebLockExclusiveFileTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
