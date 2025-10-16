// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCtrlInstanceMemberAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceMemberCount(v int64) *DescribeCtrlInstanceMemberAccountsResponseBody
	GetInstanceMemberCount() *int64
	SetMaxInstanceMemberNum(v int64) *DescribeCtrlInstanceMemberAccountsResponseBody
	GetMaxInstanceMemberNum() *int64
	SetRequestId(v string) *DescribeCtrlInstanceMemberAccountsResponseBody
	GetRequestId() *string
}

type DescribeCtrlInstanceMemberAccountsResponseBody struct {
	// example:
	//
	// 3
	InstanceMemberCount *int64 `json:"InstanceMemberCount,omitempty" xml:"InstanceMemberCount,omitempty"`
	// example:
	//
	// 1000
	MaxInstanceMemberNum *int64 `json:"MaxInstanceMemberNum,omitempty" xml:"MaxInstanceMemberNum,omitempty"`
	// example:
	//
	// 9CC69FDA-69F6-585B-9262-A306F425****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCtrlInstanceMemberAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCtrlInstanceMemberAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCtrlInstanceMemberAccountsResponseBody) GetInstanceMemberCount() *int64 {
	return s.InstanceMemberCount
}

func (s *DescribeCtrlInstanceMemberAccountsResponseBody) GetMaxInstanceMemberNum() *int64 {
	return s.MaxInstanceMemberNum
}

func (s *DescribeCtrlInstanceMemberAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCtrlInstanceMemberAccountsResponseBody) SetInstanceMemberCount(v int64) *DescribeCtrlInstanceMemberAccountsResponseBody {
	s.InstanceMemberCount = &v
	return s
}

func (s *DescribeCtrlInstanceMemberAccountsResponseBody) SetMaxInstanceMemberNum(v int64) *DescribeCtrlInstanceMemberAccountsResponseBody {
	s.MaxInstanceMemberNum = &v
	return s
}

func (s *DescribeCtrlInstanceMemberAccountsResponseBody) SetRequestId(v string) *DescribeCtrlInstanceMemberAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCtrlInstanceMemberAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}
