// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIDBClusterCustomModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleted(v bool) *DeleteAIDBClusterCustomModelResponseBody
	GetDeleted() *bool
	SetModelId(v int64) *DeleteAIDBClusterCustomModelResponseBody
	GetModelId() *int64
	SetModelName(v string) *DeleteAIDBClusterCustomModelResponseBody
	GetModelName() *string
	SetRequestId(v string) *DeleteAIDBClusterCustomModelResponseBody
	GetRequestId() *string
}

type DeleteAIDBClusterCustomModelResponseBody struct {
	// Indicates whether the deletion was successful.
	//
	// example:
	//
	// true
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// The ID of the deleted model registration.
	//
	// example:
	//
	// 123456
	ModelId *int64 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The key of the deleted custom model registration.
	//
	// example:
	//
	// Qwen3-32B
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3AA6E0E4-1234-5678-90AB-1234567890AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAIDBClusterCustomModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIDBClusterCustomModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAIDBClusterCustomModelResponseBody) GetDeleted() *bool {
	return s.Deleted
}

func (s *DeleteAIDBClusterCustomModelResponseBody) GetModelId() *int64 {
	return s.ModelId
}

func (s *DeleteAIDBClusterCustomModelResponseBody) GetModelName() *string {
	return s.ModelName
}

func (s *DeleteAIDBClusterCustomModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAIDBClusterCustomModelResponseBody) SetDeleted(v bool) *DeleteAIDBClusterCustomModelResponseBody {
	s.Deleted = &v
	return s
}

func (s *DeleteAIDBClusterCustomModelResponseBody) SetModelId(v int64) *DeleteAIDBClusterCustomModelResponseBody {
	s.ModelId = &v
	return s
}

func (s *DeleteAIDBClusterCustomModelResponseBody) SetModelName(v string) *DeleteAIDBClusterCustomModelResponseBody {
	s.ModelName = &v
	return s
}

func (s *DeleteAIDBClusterCustomModelResponseBody) SetRequestId(v string) *DeleteAIDBClusterCustomModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAIDBClusterCustomModelResponseBody) Validate() error {
	return dara.Validate(s)
}
