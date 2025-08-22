// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEncryptRoutineUidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribeEncryptRoutineUidResponseBody
	GetContent() *string
	SetRequestId(v string) *DescribeEncryptRoutineUidResponseBody
	GetRequestId() *string
}

type DescribeEncryptRoutineUidResponseBody struct {
	// The returned ciphertext, which contains the Alibaba Cloud account ID, timestamp, and time to live (TTL).
	//
	// example:
	//
	// XXXXXj20p4UB/xgdOH5LtXXXXXX
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4DBA68F5-04A9-406B-B1E4-F2CB635E103F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEncryptRoutineUidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptRoutineUidResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEncryptRoutineUidResponseBody) GetContent() *string {
	return s.Content
}

func (s *DescribeEncryptRoutineUidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEncryptRoutineUidResponseBody) SetContent(v string) *DescribeEncryptRoutineUidResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeEncryptRoutineUidResponseBody) SetRequestId(v string) *DescribeEncryptRoutineUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEncryptRoutineUidResponseBody) Validate() error {
	return dara.Validate(s)
}
