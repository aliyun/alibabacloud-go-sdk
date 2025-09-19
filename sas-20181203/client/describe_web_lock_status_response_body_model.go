// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCount(v int32) *DescribeWebLockStatusResponseBody
	GetAuthCount() *int32
	SetBindCount(v int32) *DescribeWebLockStatusResponseBody
	GetBindCount() *int32
	SetBlockCount(v int32) *DescribeWebLockStatusResponseBody
	GetBlockCount() *int32
	SetDirCount(v int32) *DescribeWebLockStatusResponseBody
	GetDirCount() *int32
	SetExpireTime(v int64) *DescribeWebLockStatusResponseBody
	GetExpireTime() *int64
	SetRequestId(v string) *DescribeWebLockStatusResponseBody
	GetRequestId() *string
	SetWhiteCount(v int32) *DescribeWebLockStatusResponseBody
	GetWhiteCount() *int32
}

type DescribeWebLockStatusResponseBody struct {
	// The total quota that you purchase for web tamper proofing.
	//
	// example:
	//
	// 32
	AuthCount *int32 `json:"AuthCount,omitempty" xml:"AuthCount,omitempty"`
	// The associated tamper proofing quota.
	//
	// example:
	//
	// 2
	BindCount *int32 `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	// The number of blocked processes.
	//
	// example:
	//
	// 48
	BlockCount *int32 `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
	// The number of protected directories.
	//
	// example:
	//
	// 2
	DirCount *int32 `json:"DirCount,omitempty" xml:"DirCount,omitempty"`
	// The timestamp generated when the quota for tamper proofing expires. Unit: millisecond.
	//
	// example:
	//
	// 1688090851000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF46038
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of processes in the whitelist.
	//
	// example:
	//
	// 6
	WhiteCount *int32 `json:"WhiteCount,omitempty" xml:"WhiteCount,omitempty"`
}

func (s DescribeWebLockStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebLockStatusResponseBody) GetAuthCount() *int32 {
	return s.AuthCount
}

func (s *DescribeWebLockStatusResponseBody) GetBindCount() *int32 {
	return s.BindCount
}

func (s *DescribeWebLockStatusResponseBody) GetBlockCount() *int32 {
	return s.BlockCount
}

func (s *DescribeWebLockStatusResponseBody) GetDirCount() *int32 {
	return s.DirCount
}

func (s *DescribeWebLockStatusResponseBody) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeWebLockStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebLockStatusResponseBody) GetWhiteCount() *int32 {
	return s.WhiteCount
}

func (s *DescribeWebLockStatusResponseBody) SetAuthCount(v int32) *DescribeWebLockStatusResponseBody {
	s.AuthCount = &v
	return s
}

func (s *DescribeWebLockStatusResponseBody) SetBindCount(v int32) *DescribeWebLockStatusResponseBody {
	s.BindCount = &v
	return s
}

func (s *DescribeWebLockStatusResponseBody) SetBlockCount(v int32) *DescribeWebLockStatusResponseBody {
	s.BlockCount = &v
	return s
}

func (s *DescribeWebLockStatusResponseBody) SetDirCount(v int32) *DescribeWebLockStatusResponseBody {
	s.DirCount = &v
	return s
}

func (s *DescribeWebLockStatusResponseBody) SetExpireTime(v int64) *DescribeWebLockStatusResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeWebLockStatusResponseBody) SetRequestId(v string) *DescribeWebLockStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebLockStatusResponseBody) SetWhiteCount(v int32) *DescribeWebLockStatusResponseBody {
	s.WhiteCount = &v
	return s
}

func (s *DescribeWebLockStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
