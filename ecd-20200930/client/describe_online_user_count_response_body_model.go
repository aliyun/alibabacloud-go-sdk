// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOnlineUserCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdAssignedUserCount(v int64) *DescribeOnlineUserCountResponseBody
	GetAdAssignedUserCount() *int64
	SetAssignedUserCount(v int64) *DescribeOnlineUserCountResponseBody
	GetAssignedUserCount() *int64
	SetOnlineUserCount(v int64) *DescribeOnlineUserCountResponseBody
	GetOnlineUserCount() *int64
	SetRequestId(v string) *DescribeOnlineUserCountResponseBody
	GetRequestId() *string
	SetSimpleAssignedUserCount(v int64) *DescribeOnlineUserCountResponseBody
	GetSimpleAssignedUserCount() *int64
}

type DescribeOnlineUserCountResponseBody struct {
	// example:
	//
	// 2
	AdAssignedUserCount *int64 `json:"AdAssignedUserCount,omitempty" xml:"AdAssignedUserCount,omitempty"`
	// example:
	//
	// 10
	AssignedUserCount *int64 `json:"AssignedUserCount,omitempty" xml:"AssignedUserCount,omitempty"`
	// example:
	//
	// 1
	OnlineUserCount *int64 `json:"OnlineUserCount,omitempty" xml:"OnlineUserCount,omitempty"`
	// example:
	//
	// 269BDB16-2CD8-4865-84BD-11C40BC21DB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8
	SimpleAssignedUserCount *int64 `json:"SimpleAssignedUserCount,omitempty" xml:"SimpleAssignedUserCount,omitempty"`
}

func (s DescribeOnlineUserCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnlineUserCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOnlineUserCountResponseBody) GetAdAssignedUserCount() *int64 {
	return s.AdAssignedUserCount
}

func (s *DescribeOnlineUserCountResponseBody) GetAssignedUserCount() *int64 {
	return s.AssignedUserCount
}

func (s *DescribeOnlineUserCountResponseBody) GetOnlineUserCount() *int64 {
	return s.OnlineUserCount
}

func (s *DescribeOnlineUserCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOnlineUserCountResponseBody) GetSimpleAssignedUserCount() *int64 {
	return s.SimpleAssignedUserCount
}

func (s *DescribeOnlineUserCountResponseBody) SetAdAssignedUserCount(v int64) *DescribeOnlineUserCountResponseBody {
	s.AdAssignedUserCount = &v
	return s
}

func (s *DescribeOnlineUserCountResponseBody) SetAssignedUserCount(v int64) *DescribeOnlineUserCountResponseBody {
	s.AssignedUserCount = &v
	return s
}

func (s *DescribeOnlineUserCountResponseBody) SetOnlineUserCount(v int64) *DescribeOnlineUserCountResponseBody {
	s.OnlineUserCount = &v
	return s
}

func (s *DescribeOnlineUserCountResponseBody) SetRequestId(v string) *DescribeOnlineUserCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOnlineUserCountResponseBody) SetSimpleAssignedUserCount(v int64) *DescribeOnlineUserCountResponseBody {
	s.SimpleAssignedUserCount = &v
	return s
}

func (s *DescribeOnlineUserCountResponseBody) Validate() error {
	return dara.Validate(s)
}
