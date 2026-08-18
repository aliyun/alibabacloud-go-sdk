// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSnapshot interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactDiskTotalSizeInB(v int64) *Snapshot
	GetArtifactDiskTotalSizeInB() *int64
	SetArtifactDiskUsedSizeInB(v int64) *Snapshot
	GetArtifactDiskUsedSizeInB() *int64
	SetArtifactMemCacheSizeInB(v int64) *Snapshot
	GetArtifactMemCacheSizeInB() *int64
	SetArtifactMemTotalSizeInB(v int64) *Snapshot
	GetArtifactMemTotalSizeInB() *int64
	SetArtifactMemUsedSizeInB(v int64) *Snapshot
	GetArtifactMemUsedSizeInB() *int64
	SetCpu(v int64) *Snapshot
	GetCpu() *int64
	SetCreatedTime(v string) *Snapshot
	GetCreatedTime() *string
	SetDescription(v string) *Snapshot
	GetDescription() *string
	SetDiskSizeMB(v int64) *Snapshot
	GetDiskSizeMB() *int64
	SetEnvs(v map[string]*string) *Snapshot
	GetEnvs() map[string]*string
	SetExpiredTime(v string) *Snapshot
	GetExpiredTime() *string
	SetFunctionName(v string) *Snapshot
	GetFunctionName() *string
	SetImageDigest(v string) *Snapshot
	GetImageDigest() *string
	SetImageRepository(v string) *Snapshot
	GetImageRepository() *string
	SetMemoryMB(v int64) *Snapshot
	GetMemoryMB() *int64
	SetOsType(v string) *Snapshot
	GetOsType() *string
	SetQualifier(v string) *Snapshot
	GetQualifier() *string
	SetReadyCommand(v string) *Snapshot
	GetReadyCommand() *string
	SetResolvedVersion(v string) *Snapshot
	GetResolvedVersion() *string
	SetSnapshotId(v string) *Snapshot
	GetSnapshotId() *string
	SetSourceSessionId(v string) *Snapshot
	GetSourceSessionId() *string
	SetStartCommand(v string) *Snapshot
	GetStartCommand() *string
	SetStatus(v string) *Snapshot
	GetStatus() *string
}

type Snapshot struct {
	// This parameter is required.
	//
	// example:
	//
	// 10737418240
	ArtifactDiskTotalSizeInB *int64 `json:"artifactDiskTotalSizeInB,omitempty" xml:"artifactDiskTotalSizeInB,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2147483648
	ArtifactDiskUsedSizeInB *int64 `json:"artifactDiskUsedSizeInB,omitempty" xml:"artifactDiskUsedSizeInB,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 268435456
	ArtifactMemCacheSizeInB *int64 `json:"artifactMemCacheSizeInB,omitempty" xml:"artifactMemCacheSizeInB,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3221225472
	ArtifactMemTotalSizeInB *int64 `json:"artifactMemTotalSizeInB,omitempty" xml:"artifactMemTotalSizeInB,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1073741824
	ArtifactMemUsedSizeInB *int64 `json:"artifactMemUsedSizeInB,omitempty" xml:"artifactMemUsedSizeInB,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Cpu *int64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2026-07-30T10:00:00Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Snapshot for production environment
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10240
	DiskSizeMB *int64 `json:"diskSizeMB,omitempty" xml:"diskSizeMB,omitempty"`
	// This parameter is required.
	Envs map[string]*string `json:"envs" xml:"envs"`
	// This parameter is required.
	//
	// example:
	//
	// 2026-08-29T10:00:00Z
	ExpiredTime *string `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-function
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sha256:0123456789abcdef
	ImageDigest *string `json:"imageDigest,omitempty" xml:"imageDigest,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/example/function
	ImageRepository *string `json:"imageRepository,omitempty" xml:"imageRepository,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3072
	MemoryMB *int64 `json:"memoryMB,omitempty" xml:"memoryMB,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// linux
	OsType *string `json:"osType,omitempty" xml:"osType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// production
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /code/ready.sh
	ReadyCommand *string `json:"readyCommand,omitempty" xml:"readyCommand,omitempty"`
	// example:
	//
	// 1
	ResolvedVersion *string `json:"resolvedVersion,omitempty" xml:"resolvedVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	SnapshotId *string `json:"snapshotId,omitempty" xml:"snapshotId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// custom-test-session-id
	SourceSessionId *string `json:"sourceSessionId,omitempty" xml:"sourceSessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /code/start.sh
	StartCommand *string `json:"startCommand,omitempty" xml:"startCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s Snapshot) String() string {
	return dara.Prettify(s)
}

func (s Snapshot) GoString() string {
	return s.String()
}

func (s *Snapshot) GetArtifactDiskTotalSizeInB() *int64 {
	return s.ArtifactDiskTotalSizeInB
}

func (s *Snapshot) GetArtifactDiskUsedSizeInB() *int64 {
	return s.ArtifactDiskUsedSizeInB
}

func (s *Snapshot) GetArtifactMemCacheSizeInB() *int64 {
	return s.ArtifactMemCacheSizeInB
}

func (s *Snapshot) GetArtifactMemTotalSizeInB() *int64 {
	return s.ArtifactMemTotalSizeInB
}

func (s *Snapshot) GetArtifactMemUsedSizeInB() *int64 {
	return s.ArtifactMemUsedSizeInB
}

func (s *Snapshot) GetCpu() *int64 {
	return s.Cpu
}

func (s *Snapshot) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Snapshot) GetDescription() *string {
	return s.Description
}

func (s *Snapshot) GetDiskSizeMB() *int64 {
	return s.DiskSizeMB
}

func (s *Snapshot) GetEnvs() map[string]*string {
	return s.Envs
}

func (s *Snapshot) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *Snapshot) GetFunctionName() *string {
	return s.FunctionName
}

func (s *Snapshot) GetImageDigest() *string {
	return s.ImageDigest
}

func (s *Snapshot) GetImageRepository() *string {
	return s.ImageRepository
}

func (s *Snapshot) GetMemoryMB() *int64 {
	return s.MemoryMB
}

func (s *Snapshot) GetOsType() *string {
	return s.OsType
}

func (s *Snapshot) GetQualifier() *string {
	return s.Qualifier
}

func (s *Snapshot) GetReadyCommand() *string {
	return s.ReadyCommand
}

func (s *Snapshot) GetResolvedVersion() *string {
	return s.ResolvedVersion
}

func (s *Snapshot) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *Snapshot) GetSourceSessionId() *string {
	return s.SourceSessionId
}

func (s *Snapshot) GetStartCommand() *string {
	return s.StartCommand
}

func (s *Snapshot) GetStatus() *string {
	return s.Status
}

func (s *Snapshot) SetArtifactDiskTotalSizeInB(v int64) *Snapshot {
	s.ArtifactDiskTotalSizeInB = &v
	return s
}

func (s *Snapshot) SetArtifactDiskUsedSizeInB(v int64) *Snapshot {
	s.ArtifactDiskUsedSizeInB = &v
	return s
}

func (s *Snapshot) SetArtifactMemCacheSizeInB(v int64) *Snapshot {
	s.ArtifactMemCacheSizeInB = &v
	return s
}

func (s *Snapshot) SetArtifactMemTotalSizeInB(v int64) *Snapshot {
	s.ArtifactMemTotalSizeInB = &v
	return s
}

func (s *Snapshot) SetArtifactMemUsedSizeInB(v int64) *Snapshot {
	s.ArtifactMemUsedSizeInB = &v
	return s
}

func (s *Snapshot) SetCpu(v int64) *Snapshot {
	s.Cpu = &v
	return s
}

func (s *Snapshot) SetCreatedTime(v string) *Snapshot {
	s.CreatedTime = &v
	return s
}

func (s *Snapshot) SetDescription(v string) *Snapshot {
	s.Description = &v
	return s
}

func (s *Snapshot) SetDiskSizeMB(v int64) *Snapshot {
	s.DiskSizeMB = &v
	return s
}

func (s *Snapshot) SetEnvs(v map[string]*string) *Snapshot {
	s.Envs = v
	return s
}

func (s *Snapshot) SetExpiredTime(v string) *Snapshot {
	s.ExpiredTime = &v
	return s
}

func (s *Snapshot) SetFunctionName(v string) *Snapshot {
	s.FunctionName = &v
	return s
}

func (s *Snapshot) SetImageDigest(v string) *Snapshot {
	s.ImageDigest = &v
	return s
}

func (s *Snapshot) SetImageRepository(v string) *Snapshot {
	s.ImageRepository = &v
	return s
}

func (s *Snapshot) SetMemoryMB(v int64) *Snapshot {
	s.MemoryMB = &v
	return s
}

func (s *Snapshot) SetOsType(v string) *Snapshot {
	s.OsType = &v
	return s
}

func (s *Snapshot) SetQualifier(v string) *Snapshot {
	s.Qualifier = &v
	return s
}

func (s *Snapshot) SetReadyCommand(v string) *Snapshot {
	s.ReadyCommand = &v
	return s
}

func (s *Snapshot) SetResolvedVersion(v string) *Snapshot {
	s.ResolvedVersion = &v
	return s
}

func (s *Snapshot) SetSnapshotId(v string) *Snapshot {
	s.SnapshotId = &v
	return s
}

func (s *Snapshot) SetSourceSessionId(v string) *Snapshot {
	s.SourceSessionId = &v
	return s
}

func (s *Snapshot) SetStartCommand(v string) *Snapshot {
	s.StartCommand = &v
	return s
}

func (s *Snapshot) SetStatus(v string) *Snapshot {
	s.Status = &v
	return s
}

func (s *Snapshot) Validate() error {
	return dara.Validate(s)
}
