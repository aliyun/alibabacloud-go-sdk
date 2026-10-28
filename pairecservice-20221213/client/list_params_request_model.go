// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncrypted(v bool) *ListParamsRequest
	GetEncrypted() *bool
	SetEnvironment(v string) *ListParamsRequest
	GetEnvironment() *string
	SetInstanceId(v string) *ListParamsRequest
	GetInstanceId() *string
	SetName(v string) *ListParamsRequest
	GetName() *string
	SetPageNumber(v int32) *ListParamsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListParamsRequest
	GetPageSize() *int32
	SetSceneId(v string) *ListParamsRequest
	GetSceneId() *string
}

type ListParamsRequest struct {
	// Specifies whether the parameter is encrypted.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The environment to which the parameter belongs. Valid values:
	//
	// - Daily: daily environment
	//
	// - Pre: staging environment
	//
	// - Prod: production environment
	//
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// The instance ID. You can obtain this ID by calling the ListInstances operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// home
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The scene ID. You can obtain this ID by calling the ListScenes operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ListParamsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListParamsRequest) GoString() string {
	return s.String()
}

func (s *ListParamsRequest) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *ListParamsRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *ListParamsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListParamsRequest) GetName() *string {
	return s.Name
}

func (s *ListParamsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListParamsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListParamsRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *ListParamsRequest) SetEncrypted(v bool) *ListParamsRequest {
	s.Encrypted = &v
	return s
}

func (s *ListParamsRequest) SetEnvironment(v string) *ListParamsRequest {
	s.Environment = &v
	return s
}

func (s *ListParamsRequest) SetInstanceId(v string) *ListParamsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListParamsRequest) SetName(v string) *ListParamsRequest {
	s.Name = &v
	return s
}

func (s *ListParamsRequest) SetPageNumber(v int32) *ListParamsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListParamsRequest) SetPageSize(v int32) *ListParamsRequest {
	s.PageSize = &v
	return s
}

func (s *ListParamsRequest) SetSceneId(v string) *ListParamsRequest {
	s.SceneId = &v
	return s
}

func (s *ListParamsRequest) Validate() error {
	return dara.Validate(s)
}
