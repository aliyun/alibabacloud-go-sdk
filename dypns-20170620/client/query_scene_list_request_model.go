// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySceneListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *QuerySceneListRequest
	GetAuthType() *string
	SetOwnerId(v int64) *QuerySceneListRequest
	GetOwnerId() *int64
	SetProdCode(v string) *QuerySceneListRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *QuerySceneListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySceneListRequest
	GetResourceOwnerId() *int64
}

type QuerySceneListRequest struct {
	AuthType             *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QuerySceneListRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySceneListRequest) GoString() string {
	return s.String()
}

func (s *QuerySceneListRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *QuerySceneListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySceneListRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QuerySceneListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySceneListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySceneListRequest) SetAuthType(v string) *QuerySceneListRequest {
	s.AuthType = &v
	return s
}

func (s *QuerySceneListRequest) SetOwnerId(v int64) *QuerySceneListRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySceneListRequest) SetProdCode(v string) *QuerySceneListRequest {
	s.ProdCode = &v
	return s
}

func (s *QuerySceneListRequest) SetResourceOwnerAccount(v string) *QuerySceneListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySceneListRequest) SetResourceOwnerId(v int64) *QuerySceneListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySceneListRequest) Validate() error {
	return dara.Validate(s)
}
