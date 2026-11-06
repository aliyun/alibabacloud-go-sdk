// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateSceneKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v string) *GenerateSceneKeyRequest
	GetDesktopId() *string
	SetExpireTime(v int32) *GenerateSceneKeyRequest
	GetExpireTime() *int32
	SetRegionId(v string) *GenerateSceneKeyRequest
	GetRegionId() *string
	SetScene(v string) *GenerateSceneKeyRequest
	GetScene() *string
}

type GenerateSceneKeyRequest struct {
	// The cloud desktop ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The validity period of the key, in hours. A value of 0 indicates that the key never expires. If this parameter is not specified, the default server value is used.
	//
	// example:
	//
	// 6
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scenario name.
	//
	// example:
	//
	// AgentGateway
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s GenerateSceneKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateSceneKeyRequest) GoString() string {
	return s.String()
}

func (s *GenerateSceneKeyRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *GenerateSceneKeyRequest) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *GenerateSceneKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateSceneKeyRequest) GetScene() *string {
	return s.Scene
}

func (s *GenerateSceneKeyRequest) SetDesktopId(v string) *GenerateSceneKeyRequest {
	s.DesktopId = &v
	return s
}

func (s *GenerateSceneKeyRequest) SetExpireTime(v int32) *GenerateSceneKeyRequest {
	s.ExpireTime = &v
	return s
}

func (s *GenerateSceneKeyRequest) SetRegionId(v string) *GenerateSceneKeyRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateSceneKeyRequest) SetScene(v string) *GenerateSceneKeyRequest {
	s.Scene = &v
	return s
}

func (s *GenerateSceneKeyRequest) Validate() error {
	return dara.Validate(s)
}
