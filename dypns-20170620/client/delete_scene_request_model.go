// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteSceneRequest
	GetOwnerId() *int64
	SetProdCode(v string) *DeleteSceneRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *DeleteSceneRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSceneRequest
	GetResourceOwnerId() *int64
	SetSceneCode(v string) *DeleteSceneRequest
	GetSceneCode() *string
}

type DeleteSceneRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// Dypns
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// FC100*******4085
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s DeleteSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSceneRequest) GoString() string {
	return s.String()
}

func (s *DeleteSceneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSceneRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *DeleteSceneRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSceneRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSceneRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DeleteSceneRequest) SetOwnerId(v int64) *DeleteSceneRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSceneRequest) SetProdCode(v string) *DeleteSceneRequest {
	s.ProdCode = &v
	return s
}

func (s *DeleteSceneRequest) SetResourceOwnerAccount(v string) *DeleteSceneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSceneRequest) SetResourceOwnerId(v int64) *DeleteSceneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSceneRequest) SetSceneCode(v string) *DeleteSceneRequest {
	s.SceneCode = &v
	return s
}

func (s *DeleteSceneRequest) Validate() error {
	return dara.Validate(s)
}
