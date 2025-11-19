// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWuyingServerSceneUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *GenerateWuyingServerSceneUrlRequest
	GetClientId() *string
	SetClientIp(v string) *GenerateWuyingServerSceneUrlRequest
	GetClientIp() *string
	SetClientOS(v string) *GenerateWuyingServerSceneUrlRequest
	GetClientOS() *string
	SetClientType(v string) *GenerateWuyingServerSceneUrlRequest
	GetClientType() *string
	SetClientVersion(v string) *GenerateWuyingServerSceneUrlRequest
	GetClientVersion() *string
	SetEndUserId(v string) *GenerateWuyingServerSceneUrlRequest
	GetEndUserId() *string
	SetLoginRegionId(v string) *GenerateWuyingServerSceneUrlRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *GenerateWuyingServerSceneUrlRequest
	GetLoginToken() *string
	SetProductType(v string) *GenerateWuyingServerSceneUrlRequest
	GetProductType() *string
	SetScene(v string) *GenerateWuyingServerSceneUrlRequest
	GetScene() *string
	SetSessionId(v string) *GenerateWuyingServerSceneUrlRequest
	GetSessionId() *string
	SetUuid(v string) *GenerateWuyingServerSceneUrlRequest
	GetUuid() *string
	SetWuyingServerId(v string) *GenerateWuyingServerSceneUrlRequest
	GetWuyingServerId() *string
}

type GenerateWuyingServerSceneUrlRequest struct {
	ClientId      *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientIp      *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientOS      *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientType    *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	EndUserId     *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// This parameter is required.
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// This parameter is required.
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// This parameter is required.
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Uuid      *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// This parameter is required.
	WuyingServerId *string `json:"WuyingServerId,omitempty" xml:"WuyingServerId,omitempty"`
}

func (s GenerateWuyingServerSceneUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateWuyingServerSceneUrlRequest) GoString() string {
	return s.String()
}

func (s *GenerateWuyingServerSceneUrlRequest) GetClientId() *string {
	return s.ClientId
}

func (s *GenerateWuyingServerSceneUrlRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *GenerateWuyingServerSceneUrlRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *GenerateWuyingServerSceneUrlRequest) GetClientType() *string {
	return s.ClientType
}

func (s *GenerateWuyingServerSceneUrlRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *GenerateWuyingServerSceneUrlRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *GenerateWuyingServerSceneUrlRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *GenerateWuyingServerSceneUrlRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *GenerateWuyingServerSceneUrlRequest) GetProductType() *string {
	return s.ProductType
}

func (s *GenerateWuyingServerSceneUrlRequest) GetScene() *string {
	return s.Scene
}

func (s *GenerateWuyingServerSceneUrlRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GenerateWuyingServerSceneUrlRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GenerateWuyingServerSceneUrlRequest) GetWuyingServerId() *string {
	return s.WuyingServerId
}

func (s *GenerateWuyingServerSceneUrlRequest) SetClientId(v string) *GenerateWuyingServerSceneUrlRequest {
	s.ClientId = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlRequest) SetClientIp(v string) *GenerateWuyingServerSceneUrlRequest {
	s.ClientIp = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlRequest) SetClientOS(v string) *GenerateWuyingServerSceneUrlRequest {
	s.ClientOS = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlRequest) SetClientType(v string) *GenerateWuyingServerSceneUrlRequest {
	s.ClientType = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlRequest) SetClientVersion(v string) *GenerateWuyingServerSceneUrlRequest {
	s.ClientVersion = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlRequest) SetEndUserId(v string) *GenerateWuyingServerSceneUrlRequest {
	s.EndUserId = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlRequest) SetLoginRegionId(v string) *GenerateWuyingServerSceneUrlRequest {
	s.LoginRegionId = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlRequest) SetLoginToken(v string) *GenerateWuyingServerSceneUrlRequest {
	s.LoginToken = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlRequest) SetProductType(v string) *GenerateWuyingServerSceneUrlRequest {
	s.ProductType = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlRequest) SetScene(v string) *GenerateWuyingServerSceneUrlRequest {
	s.Scene = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlRequest) SetSessionId(v string) *GenerateWuyingServerSceneUrlRequest {
	s.SessionId = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlRequest) SetUuid(v string) *GenerateWuyingServerSceneUrlRequest {
	s.Uuid = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlRequest) SetWuyingServerId(v string) *GenerateWuyingServerSceneUrlRequest {
	s.WuyingServerId = &v
	return s
}

func (s *GenerateWuyingServerSceneUrlRequest) Validate() error {
	return dara.Validate(s)
}
