// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWuyingServerSceneUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExpireDurationHours(v int32) *GenerateWuyingServerSceneUrlResponseBody
	GetExpireDurationHours() *int32
	SetExpireTime(v string) *GenerateWuyingServerSceneUrlResponseBody
	GetExpireTime() *string
	SetRequestId(v string) *GenerateWuyingServerSceneUrlResponseBody
	GetRequestId() *string
	SetUrl(v string) *GenerateWuyingServerSceneUrlResponseBody
	GetUrl() *string
}

type GenerateWuyingServerSceneUrlResponseBody struct {
	ExpireDurationHours *int32  `json:"ExpireDurationHours,omitempty" xml:"ExpireDurationHours,omitempty"`
	ExpireTime          *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Url                 *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GenerateWuyingServerSceneUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateWuyingServerSceneUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateWuyingServerSceneUrlResponseBody) GetExpireDurationHours() *int32 {
	return s.ExpireDurationHours
}

func (s *GenerateWuyingServerSceneUrlResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GenerateWuyingServerSceneUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateWuyingServerSceneUrlResponseBody) GetUrl() *string {
	return s.Url
}

func (s *GenerateWuyingServerSceneUrlResponseBody) SetExpireDurationHours(v int32) *GenerateWuyingServerSceneUrlResponseBody {
	s.ExpireDurationHours = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlResponseBody) SetExpireTime(v string) *GenerateWuyingServerSceneUrlResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlResponseBody) SetRequestId(v string) *GenerateWuyingServerSceneUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlResponseBody) SetUrl(v string) *GenerateWuyingServerSceneUrlResponseBody {
	s.Url = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
