// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartCache interface {
	dara.Model
	String() string
	GoString() string
	SetCacheWorkerNum(v int64) *SmartCache
	GetCacheWorkerNum() *int64
	SetCacheWorkerSize(v int64) *SmartCache
	GetCacheWorkerSize() *int64
	SetDescription(v string) *SmartCache
	GetDescription() *string
	SetDisplayName(v string) *SmartCache
	GetDisplayName() *string
	SetDuration(v string) *SmartCache
	GetDuration() *string
	SetEndpoint(v string) *SmartCache
	GetEndpoint() *string
	SetFileSystemId(v string) *SmartCache
	GetFileSystemId() *string
	SetGmtCreateTime(v string) *SmartCache
	GetGmtCreateTime() *string
	SetGmtModifyTime(v string) *SmartCache
	GetGmtModifyTime() *string
	SetMountPath(v string) *SmartCache
	GetMountPath() *string
	SetOptions(v string) *SmartCache
	GetOptions() *string
	SetPath(v string) *SmartCache
	GetPath() *string
	SetSmartCacheId(v string) *SmartCache
	GetSmartCacheId() *string
	SetStatus(v string) *SmartCache
	GetStatus() *string
	SetType(v string) *SmartCache
	GetType() *string
	SetUserId(v string) *SmartCache
	GetUserId() *string
}

type SmartCache struct {
	// example:
	//
	// 10
	CacheWorkerNum *int64 `json:"CacheWorkerNum,omitempty" xml:"CacheWorkerNum,omitempty"`
	// example:
	//
	// 100
	CacheWorkerSize *int64 `json:"CacheWorkerSize,omitempty" xml:"CacheWorkerSize,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 123456
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// oss-cn-beijing-internal.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T23:36:01Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// example:
	//
	// /root/data/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// example:
	//
	// {"num_threads": 32}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// oss://buc/path/to/dir
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// smartcache-20210114104214-vf9lowjt3pso
	SmartCacheId *string `json:"SmartCacheId,omitempty" xml:"SmartCacheId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// oss
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 189xxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SmartCache) String() string {
	return dara.Prettify(s)
}

func (s SmartCache) GoString() string {
	return s.String()
}

func (s *SmartCache) GetCacheWorkerNum() *int64 {
	return s.CacheWorkerNum
}

func (s *SmartCache) GetCacheWorkerSize() *int64 {
	return s.CacheWorkerSize
}

func (s *SmartCache) GetDescription() *string {
	return s.Description
}

func (s *SmartCache) GetDisplayName() *string {
	return s.DisplayName
}

func (s *SmartCache) GetDuration() *string {
	return s.Duration
}

func (s *SmartCache) GetEndpoint() *string {
	return s.Endpoint
}

func (s *SmartCache) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *SmartCache) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *SmartCache) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *SmartCache) GetMountPath() *string {
	return s.MountPath
}

func (s *SmartCache) GetOptions() *string {
	return s.Options
}

func (s *SmartCache) GetPath() *string {
	return s.Path
}

func (s *SmartCache) GetSmartCacheId() *string {
	return s.SmartCacheId
}

func (s *SmartCache) GetStatus() *string {
	return s.Status
}

func (s *SmartCache) GetType() *string {
	return s.Type
}

func (s *SmartCache) GetUserId() *string {
	return s.UserId
}

func (s *SmartCache) SetCacheWorkerNum(v int64) *SmartCache {
	s.CacheWorkerNum = &v
	return s
}

func (s *SmartCache) SetCacheWorkerSize(v int64) *SmartCache {
	s.CacheWorkerSize = &v
	return s
}

func (s *SmartCache) SetDescription(v string) *SmartCache {
	s.Description = &v
	return s
}

func (s *SmartCache) SetDisplayName(v string) *SmartCache {
	s.DisplayName = &v
	return s
}

func (s *SmartCache) SetDuration(v string) *SmartCache {
	s.Duration = &v
	return s
}

func (s *SmartCache) SetEndpoint(v string) *SmartCache {
	s.Endpoint = &v
	return s
}

func (s *SmartCache) SetFileSystemId(v string) *SmartCache {
	s.FileSystemId = &v
	return s
}

func (s *SmartCache) SetGmtCreateTime(v string) *SmartCache {
	s.GmtCreateTime = &v
	return s
}

func (s *SmartCache) SetGmtModifyTime(v string) *SmartCache {
	s.GmtModifyTime = &v
	return s
}

func (s *SmartCache) SetMountPath(v string) *SmartCache {
	s.MountPath = &v
	return s
}

func (s *SmartCache) SetOptions(v string) *SmartCache {
	s.Options = &v
	return s
}

func (s *SmartCache) SetPath(v string) *SmartCache {
	s.Path = &v
	return s
}

func (s *SmartCache) SetSmartCacheId(v string) *SmartCache {
	s.SmartCacheId = &v
	return s
}

func (s *SmartCache) SetStatus(v string) *SmartCache {
	s.Status = &v
	return s
}

func (s *SmartCache) SetType(v string) *SmartCache {
	s.Type = &v
	return s
}

func (s *SmartCache) SetUserId(v string) *SmartCache {
	s.UserId = &v
	return s
}

func (s *SmartCache) Validate() error {
	return dara.Validate(s)
}
