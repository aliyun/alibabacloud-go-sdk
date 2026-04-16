// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFunctionLockInfo interface {
	dara.Model
	String() string
	GoString() string
	SetLockedAt(v string) *FunctionLockInfo
	GetLockedAt() *string
	SetLockedBy(v string) *FunctionLockInfo
	GetLockedBy() *string
	SetLockedResources(v []*string) *FunctionLockInfo
	GetLockedResources() []*string
}

type FunctionLockInfo struct {
	// 锁定时间
	//
	// example:
	//
	// 2025-04-05T10:00:00Z
	LockedAt *string `json:"lockedAt,omitempty" xml:"lockedAt,omitempty"`
	// 锁定方名称
	//
	// example:
	//
	// AgentRun
	LockedBy *string `json:"lockedBy,omitempty" xml:"lockedBy,omitempty"`
	// 锁定的资源类型列表
	//
	// example:
	//
	// ["function", "trigger", "version", "alias"]
	LockedResources []*string `json:"lockedResources" xml:"lockedResources" type:"Repeated"`
}

func (s FunctionLockInfo) String() string {
	return dara.Prettify(s)
}

func (s FunctionLockInfo) GoString() string {
	return s.String()
}

func (s *FunctionLockInfo) GetLockedAt() *string {
	return s.LockedAt
}

func (s *FunctionLockInfo) GetLockedBy() *string {
	return s.LockedBy
}

func (s *FunctionLockInfo) GetLockedResources() []*string {
	return s.LockedResources
}

func (s *FunctionLockInfo) SetLockedAt(v string) *FunctionLockInfo {
	s.LockedAt = &v
	return s
}

func (s *FunctionLockInfo) SetLockedBy(v string) *FunctionLockInfo {
	s.LockedBy = &v
	return s
}

func (s *FunctionLockInfo) SetLockedResources(v []*string) *FunctionLockInfo {
	s.LockedResources = v
	return s
}

func (s *FunctionLockInfo) Validate() error {
	return dara.Validate(s)
}
