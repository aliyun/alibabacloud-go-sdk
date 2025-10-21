// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelInstanceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelInstanceId(v int64) *GetModelInstanceInfoRequest
	GetModelInstanceId() *int64
	SetSceneType(v int32) *GetModelInstanceInfoRequest
	GetSceneType() *int32
}

type GetModelInstanceInfoRequest struct {
	// example:
	//
	// 123
	ModelInstanceId *int64 `json:"ModelInstanceId,omitempty" xml:"ModelInstanceId,omitempty"`
	// example:
	//
	// 2
	SceneType *int32 `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
}

func (s GetModelInstanceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelInstanceInfoRequest) GoString() string {
	return s.String()
}

func (s *GetModelInstanceInfoRequest) GetModelInstanceId() *int64 {
	return s.ModelInstanceId
}

func (s *GetModelInstanceInfoRequest) GetSceneType() *int32 {
	return s.SceneType
}

func (s *GetModelInstanceInfoRequest) SetModelInstanceId(v int64) *GetModelInstanceInfoRequest {
	s.ModelInstanceId = &v
	return s
}

func (s *GetModelInstanceInfoRequest) SetSceneType(v int32) *GetModelInstanceInfoRequest {
	s.SceneType = &v
	return s
}

func (s *GetModelInstanceInfoRequest) Validate() error {
	return dara.Validate(s)
}
