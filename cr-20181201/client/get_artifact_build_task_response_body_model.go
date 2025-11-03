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
	// The type of the artifact building task. Valid values:
	//
	// 	- `IMAGE_TO_ACCELERATED_IMAGE`: builds accelerated images for Container Service for Kubernetes (ACK) clusters.
	//
	// 	- `IMAGE_TO_ECI_ACCELERATED_IMAGE`: builds accelerated images for elastic container instances.
	//
	// example:
	//
	// IMAGE_TO_ACCELERATED_IMAGE
	ArtifactBuildType *string `json:"ArtifactBuildType,omitempty" xml:"ArtifactBuildType,omitempty"`
	// The ID of the artifact building task.
	//
	// example:
	//
	// i2a-1yu****
	BuildTaskId *string `json:"BuildTaskId,omitempty" xml:"BuildTaskId,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the artifact building task ends.
	//
	// example:
	//
	// 156871880
	EndTime      *int32    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Instructions []*string `json:"Instructions,omitempty" xml:"Instructions,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C4C7DD0C-C9D6-437A-A7EE-121EFD70D002
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the source artifact.
	SourceArtifact *GetArtifactBuildTaskResponseBodySourceArtifact `json:"SourceArtifact,omitempty" xml:"SourceArtifact,omitempty" type:"Struct"`
	// The time when the artifact building task starts.
	//
	// example:
	//
	// 156871881
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The artifact that is built in the task.
	TargetArtifact *GetArtifactBuildTaskResponseBodyTargetArtifact `json:"TargetArtifact,omitempty" xml:"TargetArtifact,omitempty" type:"Struct"`
	// The status of the artifact that is built in the task. Valid values:
	//
	// 	- `PENDING`: The artifact is being scheduled.
	//
	// 	- `BUILDING`: The artifact is being built.
	//
	// 	- `SUCCESS`: The artifact is built.
	//
	// 	- `FAILED`: The artifact fails to be built.
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

type GetArtifactBuildTaskResponseBodySourceArtifact struct {
	// The type of the artifact that is built in the task. The value can only be IMAGE.
	//
	// example:
	//
	// IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The ID of the repository to which the source artifact belongs. The repository can only be an image repository.
	//
	// example:
	//
	// cri-shac42yvqzvq****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The version of the artifact. The artifact can only be an image.
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

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) GetRepoId() *string {
	return s.RepoId
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) GetVersion() *string {
	return s.Version
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) SetArtifactType(v string) *GetArtifactBuildTaskResponseBodySourceArtifact {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) SetRepoId(v string) *GetArtifactBuildTaskResponseBodySourceArtifact {
	s.RepoId = &v
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
	// The type of the artifact that is built in the task. The value can only be IMAGE.
	//
	// example:
	//
	// IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The ID of the repository to which the artifact that is built in the task belongs. The repository can only be an image repository. The value is the same as the ID of the repository to which the source artifact belongs.
	//
	// example:
	//
	// crr-1234567
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The version of the artifact that is built in the task. The artifact can only be an image.
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

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) GetRepoId() *string {
	return s.RepoId
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) GetVersion() *string {
	return s.Version
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) SetArtifactType(v string) *GetArtifactBuildTaskResponseBodyTargetArtifact {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) SetRepoId(v string) *GetArtifactBuildTaskResponseBodyTargetArtifact {
	s.RepoId = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) SetVersion(v string) *GetArtifactBuildTaskResponseBodyTargetArtifact {
	s.Version = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) Validate() error {
	return dara.Validate(s)
}
