// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterAIDBClusterCustomModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreated(v bool) *RegisterAIDBClusterCustomModelResponseBody
	GetCreated() *bool
	SetDisplayModelName(v string) *RegisterAIDBClusterCustomModelResponseBody
	GetDisplayModelName() *string
	SetModelId(v int64) *RegisterAIDBClusterCustomModelResponseBody
	GetModelId() *int64
	SetModelName(v string) *RegisterAIDBClusterCustomModelResponseBody
	GetModelName() *string
	SetModelType(v string) *RegisterAIDBClusterCustomModelResponseBody
	GetModelType() *string
	SetOssPath(v string) *RegisterAIDBClusterCustomModelResponseBody
	GetOssPath() *string
	SetRequestId(v string) *RegisterAIDBClusterCustomModelResponseBody
	GetRequestId() *string
}

type RegisterAIDBClusterCustomModelResponseBody struct {
	// Indicates whether the registration is newly created. A value of false indicates that an existing registration was updated.
	//
	// example:
	//
	// true
	Created *bool `json:"Created,omitempty" xml:"Created,omitempty"`
	// The display name and initial client-facing invocation name.
	//
	// example:
	//
	// my-qwen3
	DisplayModelName *string `json:"DisplayModelName,omitempty" xml:"DisplayModelName,omitempty"`
	// The model registration ID.
	//
	// example:
	//
	// 123456
	ModelId *int64 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The custom model registration key.
	//
	// example:
	//
	// Qwen3-32B
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The model type.
	//
	// example:
	//
	// custom
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The normalized OSS path.
	//
	// example:
	//
	// /my-model-bucket/models/qwen3
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3AA6E0E4-1234-5678-90AB-1234567890AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterAIDBClusterCustomModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterAIDBClusterCustomModelResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterAIDBClusterCustomModelResponseBody) GetCreated() *bool {
	return s.Created
}

func (s *RegisterAIDBClusterCustomModelResponseBody) GetDisplayModelName() *string {
	return s.DisplayModelName
}

func (s *RegisterAIDBClusterCustomModelResponseBody) GetModelId() *int64 {
	return s.ModelId
}

func (s *RegisterAIDBClusterCustomModelResponseBody) GetModelName() *string {
	return s.ModelName
}

func (s *RegisterAIDBClusterCustomModelResponseBody) GetModelType() *string {
	return s.ModelType
}

func (s *RegisterAIDBClusterCustomModelResponseBody) GetOssPath() *string {
	return s.OssPath
}

func (s *RegisterAIDBClusterCustomModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterAIDBClusterCustomModelResponseBody) SetCreated(v bool) *RegisterAIDBClusterCustomModelResponseBody {
	s.Created = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelResponseBody) SetDisplayModelName(v string) *RegisterAIDBClusterCustomModelResponseBody {
	s.DisplayModelName = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelResponseBody) SetModelId(v int64) *RegisterAIDBClusterCustomModelResponseBody {
	s.ModelId = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelResponseBody) SetModelName(v string) *RegisterAIDBClusterCustomModelResponseBody {
	s.ModelName = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelResponseBody) SetModelType(v string) *RegisterAIDBClusterCustomModelResponseBody {
	s.ModelType = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelResponseBody) SetOssPath(v string) *RegisterAIDBClusterCustomModelResponseBody {
	s.OssPath = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelResponseBody) SetRequestId(v string) *RegisterAIDBClusterCustomModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelResponseBody) Validate() error {
	return dara.Validate(s)
}
