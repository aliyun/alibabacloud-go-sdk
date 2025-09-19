// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockInclusiveFileTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInclusiveFileType(v []*string) *DescribeWebLockInclusiveFileTypeResponseBody
	GetInclusiveFileType() []*string
	SetRequestId(v string) *DescribeWebLockInclusiveFileTypeResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWebLockInclusiveFileTypeResponseBody
	GetTotalCount() *int32
}

type DescribeWebLockInclusiveFileTypeResponseBody struct {
	// An array that consists of the types of files that can be protected by web tamper proofing.
	InclusiveFileType []*string `json:"InclusiveFileType,omitempty" xml:"InclusiveFileType,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// CE500770-42D3-442E-9DDD-156E0F9F3B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the types of files that can be protected by web tamper proofing.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWebLockInclusiveFileTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockInclusiveFileTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebLockInclusiveFileTypeResponseBody) GetInclusiveFileType() []*string {
	return s.InclusiveFileType
}

func (s *DescribeWebLockInclusiveFileTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebLockInclusiveFileTypeResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWebLockInclusiveFileTypeResponseBody) SetInclusiveFileType(v []*string) *DescribeWebLockInclusiveFileTypeResponseBody {
	s.InclusiveFileType = v
	return s
}

func (s *DescribeWebLockInclusiveFileTypeResponseBody) SetRequestId(v string) *DescribeWebLockInclusiveFileTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebLockInclusiveFileTypeResponseBody) SetTotalCount(v int32) *DescribeWebLockInclusiveFileTypeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebLockInclusiveFileTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
