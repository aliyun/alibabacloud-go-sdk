// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactBuildTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactBuildType(v string) *GetArtifactBuildTaskResponseBody
	GetArtifactBuildType() *string
	SetArtifactCompression(v *GetArtifactBuildTaskResponseBodyArtifactCompression) *GetArtifactBuildTaskResponseBody
	GetArtifactCompression() *GetArtifactBuildTaskResponseBodyArtifactCompression
	SetBuildTaskId(v string) *GetArtifactBuildTaskResponseBody
	GetBuildTaskId() *string
	SetCode(v string) *GetArtifactBuildTaskResponseBody
	GetCode() *string
	SetEndTime(v int32) *GetArtifactBuildTaskResponseBody
	GetEndTime() *int32
	SetInstructions(v []*string) *GetArtifactBuildTaskResponseBody
	GetInstructions() []*string
	SetIsSuccess(v bool) *GetArtifactBuildTaskResponseBody
	GetIsSuccess() *bool
	SetPriority(v int32) *GetArtifactBuildTaskResponseBody
	GetPriority() *int32
	SetRequestId(v string) *GetArtifactBuildTaskResponseBody
	GetRequestId() *string
	SetSourceArtifact(v *GetArtifactBuildTaskResponseBodySourceArtifact) *GetArtifactBuildTaskResponseBody
	GetSourceArtifact() *GetArtifactBuildTaskResponseBodySourceArtifact
	SetStartTime(v int32) *GetArtifactBuildTaskResponseBody
	GetStartTime() *int32
	SetTargetArtifact(v *GetArtifactBuildTaskResponseBodyTargetArtifact) *GetArtifactBuildTaskResponseBody
	GetTargetArtifact() *GetArtifactBuildTaskResponseBodyTargetArtifact
	SetTaskStatus(v string) *GetArtifactBuildTaskResponseBody
	GetTaskStatus() *string
}

type GetArtifactBuildTaskResponseBody struct {
	// The artifact build type. Valid values:
	//
	// - `IMAGE_TO_ACCELERATED_IMAGE`: Accelerated image creation optimized for ACK scenarios.
	//
	// - `IMAGE_TO_ECI_ACCELERATED_IMAGE`: Accelerated image artifact optimized for ECI scenarios.
	//
	// example:
	//
	// IMAGE_TO_ACCELERATED_IMAGE
	ArtifactBuildType *string `json:"ArtifactBuildType,omitempty" xml:"ArtifactBuildType,omitempty"`
	// The artifact compression parameters.
	ArtifactCompression *GetArtifactBuildTaskResponseBodyArtifactCompression `json:"ArtifactCompression,omitempty" xml:"ArtifactCompression,omitempty" type:"Struct"`
	// The ID of the artifact build task.
	//
	// example:
	//
	// i2a-1yu****
	BuildTaskId *string `json:"BuildTaskId,omitempty" xml:"BuildTaskId,omitempty"`
	// The return code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The end time. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1685415871
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The reserved field list of the artifact build task. The list elements should be empty.
	Instructions []*string `json:"Instructions,omitempty" xml:"Instructions,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// 3
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C4C7DD0C-C9D6-437A-A7EE-121EFD70D002
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source artifact.
	SourceArtifact *GetArtifactBuildTaskResponseBodySourceArtifact `json:"SourceArtifact,omitempty" xml:"SourceArtifact,omitempty" type:"Struct"`
	// The start time. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1685437471
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The target artifact.
	TargetArtifact *GetArtifactBuildTaskResponseBodyTargetArtifact `json:"TargetArtifact,omitempty" xml:"TargetArtifact,omitempty" type:"Struct"`
	// The artifact build status. Valid values:
	//
	// - `PENDING`: Scheduling in progress.
	//
	// - `BUILDING`: Building in progress.
	//
	// - `SUCCESS`: Build succeeded.
	//
	// - `FAILED`: Build failed.
	//
	// example:
	//
	// BUILDING
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetArtifactBuildTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactBuildTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildTaskResponseBody) GetArtifactBuildType() *string {
	return s.ArtifactBuildType
}

func (s *GetArtifactBuildTaskResponseBody) GetArtifactCompression() *GetArtifactBuildTaskResponseBodyArtifactCompression {
	return s.ArtifactCompression
}

func (s *GetArtifactBuildTaskResponseBody) GetBuildTaskId() *string {
	return s.BuildTaskId
}

func (s *GetArtifactBuildTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetArtifactBuildTaskResponseBody) GetEndTime() *int32 {
	return s.EndTime
}

func (s *GetArtifactBuildTaskResponseBody) GetInstructions() []*string {
	return s.Instructions
}

func (s *GetArtifactBuildTaskResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetArtifactBuildTaskResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *GetArtifactBuildTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetArtifactBuildTaskResponseBody) GetSourceArtifact() *GetArtifactBuildTaskResponseBodySourceArtifact {
	return s.SourceArtifact
}

func (s *GetArtifactBuildTaskResponseBody) GetStartTime() *int32 {
	return s.StartTime
}

func (s *GetArtifactBuildTaskResponseBody) GetTargetArtifact() *GetArtifactBuildTaskResponseBodyTargetArtifact {
	return s.TargetArtifact
}

func (s *GetArtifactBuildTaskResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetArtifactBuildTaskResponseBody) SetArtifactBuildType(v string) *GetArtifactBuildTaskResponseBody {
	s.ArtifactBuildType = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetArtifactCompression(v *GetArtifactBuildTaskResponseBodyArtifactCompression) *GetArtifactBuildTaskResponseBody {
	s.ArtifactCompression = v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetBuildTaskId(v string) *GetArtifactBuildTaskResponseBody {
	s.BuildTaskId = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetCode(v string) *GetArtifactBuildTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetEndTime(v int32) *GetArtifactBuildTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetInstructions(v []*string) *GetArtifactBuildTaskResponseBody {
	s.Instructions = v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetIsSuccess(v bool) *GetArtifactBuildTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetPriority(v int32) *GetArtifactBuildTaskResponseBody {
	s.Priority = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetRequestId(v string) *GetArtifactBuildTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetSourceArtifact(v *GetArtifactBuildTaskResponseBodySourceArtifact) *GetArtifactBuildTaskResponseBody {
	s.SourceArtifact = v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetStartTime(v int32) *GetArtifactBuildTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetTargetArtifact(v *GetArtifactBuildTaskResponseBodyTargetArtifact) *GetArtifactBuildTaskResponseBody {
	s.TargetArtifact = v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetTaskStatus(v string) *GetArtifactBuildTaskResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) Validate() error {
	if s.ArtifactCompression != nil {
		if err := s.ArtifactCompression.Validate(); err != nil {
			return err
		}
	}
	if s.SourceArtifact != nil {
		if err := s.SourceArtifact.Validate(); err != nil {
			return err
		}
	}
	if s.TargetArtifact != nil {
		if err := s.TargetArtifact.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetArtifactBuildTaskResponseBodyArtifactCompression struct {
	// The operating system and architecture.
	//
	// example:
	//
	// linux/arm64
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The number of layers to retain after compression.
	//
	// example:
	//
	// 10
	SquashKeepLayers *int32 `json:"SquashKeepLayers,omitempty" xml:"SquashKeepLayers,omitempty"`
	// The digest of the starting layer for compression.
	//
	// example:
	//
	// sha256:xxxxx
	StartLayerDigest *string `json:"StartLayerDigest,omitempty" xml:"StartLayerDigest,omitempty"`
}

func (s GetArtifactBuildTaskResponseBodyArtifactCompression) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactBuildTaskResponseBodyArtifactCompression) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildTaskResponseBodyArtifactCompression) GetPlatform() *string {
	return s.Platform
}

func (s *GetArtifactBuildTaskResponseBodyArtifactCompression) GetSquashKeepLayers() *int32 {
	return s.SquashKeepLayers
}

func (s *GetArtifactBuildTaskResponseBodyArtifactCompression) GetStartLayerDigest() *string {
	return s.StartLayerDigest
}

func (s *GetArtifactBuildTaskResponseBodyArtifactCompression) SetPlatform(v string) *GetArtifactBuildTaskResponseBodyArtifactCompression {
	s.Platform = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodyArtifactCompression) SetSquashKeepLayers(v int32) *GetArtifactBuildTaskResponseBodyArtifactCompression {
	s.SquashKeepLayers = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodyArtifactCompression) SetStartLayerDigest(v string) *GetArtifactBuildTaskResponseBodyArtifactCompression {
	s.StartLayerDigest = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodyArtifactCompression) Validate() error {
	return dara.Validate(s)
}

type GetArtifactBuildTaskResponseBodySourceArtifact struct {
	// The artifact type. Only IMAGE is supported.
	//
	// example:
	//
	// IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The number of artifact layers.
	//
	// example:
	//
	// 10
	LayerCount *int32 `json:"LayerCount,omitempty" xml:"LayerCount,omitempty"`
	// The repository ID. Only image repositories are supported.
	//
	// example:
	//
	// cri-shac42yvqzvq****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The artifact size, in bytes.
	//
	// example:
	//
	// 5
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The artifact version. Only image versions are supported.
	//
	// example:
	//
	// latest
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetArtifactBuildTaskResponseBodySourceArtifact) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactBuildTaskResponseBodySourceArtifact) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) GetLayerCount() *int32 {
	return s.LayerCount
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) GetRepoId() *string {
	return s.RepoId
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) GetSize() *int64 {
	return s.Size
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) GetVersion() *string {
	return s.Version
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) SetArtifactType(v string) *GetArtifactBuildTaskResponseBodySourceArtifact {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) SetLayerCount(v int32) *GetArtifactBuildTaskResponseBodySourceArtifact {
	s.LayerCount = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) SetRepoId(v string) *GetArtifactBuildTaskResponseBodySourceArtifact {
	s.RepoId = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) SetSize(v int64) *GetArtifactBuildTaskResponseBodySourceArtifact {
	s.Size = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) SetVersion(v string) *GetArtifactBuildTaskResponseBodySourceArtifact {
	s.Version = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) Validate() error {
	return dara.Validate(s)
}

type GetArtifactBuildTaskResponseBodyTargetArtifact struct {
	// The artifact type. Only IMAGE is supported.
	//
	// example:
	//
	// IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The number of artifact layers.
	//
	// example:
	//
	// 5
	LayerCount *int32 `json:"LayerCount,omitempty" xml:"LayerCount,omitempty"`
	// The repository ID. Only image repositories are supported. The repository ID of the target artifact must be the same as that of the source artifact.
	//
	// example:
	//
	// crr-1234567
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The artifact size, in bytes.
	//
	// example:
	//
	// 10
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The artifact version. Only images are supported.
	//
	// example:
	//
	// latest_accelerated
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetArtifactBuildTaskResponseBodyTargetArtifact) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactBuildTaskResponseBodyTargetArtifact) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) GetLayerCount() *int32 {
	return s.LayerCount
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) GetRepoId() *string {
	return s.RepoId
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) GetSize() *int64 {
	return s.Size
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) GetVersion() *string {
	return s.Version
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) SetArtifactType(v string) *GetArtifactBuildTaskResponseBodyTargetArtifact {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) SetLayerCount(v int32) *GetArtifactBuildTaskResponseBodyTargetArtifact {
	s.LayerCount = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) SetRepoId(v string) *GetArtifactBuildTaskResponseBodyTargetArtifact {
	s.RepoId = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) SetSize(v int64) *GetArtifactBuildTaskResponseBodyTargetArtifact {
	s.Size = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) SetVersion(v string) *GetArtifactBuildTaskResponseBodyTargetArtifact {
	s.Version = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) Validate() error {
	return dara.Validate(s)
}
