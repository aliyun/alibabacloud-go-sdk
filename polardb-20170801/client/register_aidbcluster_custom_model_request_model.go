// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterAIDBClusterCustomModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomOssBucketName(v string) *RegisterAIDBClusterCustomModelRequest
	GetCustomOssBucketName() *string
	SetCustomOssBucketPath(v string) *RegisterAIDBClusterCustomModelRequest
	GetCustomOssBucketPath() *string
	SetDBClusterId(v string) *RegisterAIDBClusterCustomModelRequest
	GetDBClusterId() *string
	SetDisplayModelName(v string) *RegisterAIDBClusterCustomModelRequest
	GetDisplayModelName() *string
	SetModelName(v string) *RegisterAIDBClusterCustomModelRequest
	GetModelName() *string
	SetRegionId(v string) *RegisterAIDBClusterCustomModelRequest
	GetRegionId() *string
}

type RegisterAIDBClusterCustomModelRequest struct {
	// The name of the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-model-bucket
	CustomOssBucketName *string `json:"CustomOssBucketName,omitempty" xml:"CustomOssBucketName,omitempty"`
	// The model path within the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// models/qwen3
	CustomOssBucketPath *string `json:"CustomOssBucketPath,omitempty" xml:"CustomOssBucketPath,omitempty"`
	// The ID of the PolarDB AI 3.0 logical instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pm-2ze4x2mwo81knj08a
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The display name shown in the candidate list and the initial client-facing invocation name.
	//
	// example:
	//
	// my-qwen3
	DisplayModelName *string `json:"DisplayModelName,omitempty" xml:"DisplayModelName,omitempty"`
	// The custom model registration key and model directory name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Qwen3-32B
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RegisterAIDBClusterCustomModelRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterAIDBClusterCustomModelRequest) GoString() string {
	return s.String()
}

func (s *RegisterAIDBClusterCustomModelRequest) GetCustomOssBucketName() *string {
	return s.CustomOssBucketName
}

func (s *RegisterAIDBClusterCustomModelRequest) GetCustomOssBucketPath() *string {
	return s.CustomOssBucketPath
}

func (s *RegisterAIDBClusterCustomModelRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RegisterAIDBClusterCustomModelRequest) GetDisplayModelName() *string {
	return s.DisplayModelName
}

func (s *RegisterAIDBClusterCustomModelRequest) GetModelName() *string {
	return s.ModelName
}

func (s *RegisterAIDBClusterCustomModelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RegisterAIDBClusterCustomModelRequest) SetCustomOssBucketName(v string) *RegisterAIDBClusterCustomModelRequest {
	s.CustomOssBucketName = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetCustomOssBucketPath(v string) *RegisterAIDBClusterCustomModelRequest {
	s.CustomOssBucketPath = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetDBClusterId(v string) *RegisterAIDBClusterCustomModelRequest {
	s.DBClusterId = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetDisplayModelName(v string) *RegisterAIDBClusterCustomModelRequest {
	s.DisplayModelName = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetModelName(v string) *RegisterAIDBClusterCustomModelRequest {
	s.ModelName = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetRegionId(v string) *RegisterAIDBClusterCustomModelRequest {
	s.RegionId = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) Validate() error {
	return dara.Validate(s)
}
