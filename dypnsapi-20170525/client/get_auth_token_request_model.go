// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v int32) *GetAuthTokenRequest
	GetBizType() *int32
	SetCmApiCode(v int32) *GetAuthTokenRequest
	GetCmApiCode() *int32
	SetCtApiCode(v int32) *GetAuthTokenRequest
	GetCtApiCode() *int32
	SetCuApiCode(v int32) *GetAuthTokenRequest
	GetCuApiCode() *int32
	SetOrigin(v string) *GetAuthTokenRequest
	GetOrigin() *string
	SetOwnerId(v int64) *GetAuthTokenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetAuthTokenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetAuthTokenRequest
	GetResourceOwnerId() *int64
	SetSceneCode(v string) *GetAuthTokenRequest
	GetSceneCode() *string
	SetUrl(v string) *GetAuthTokenRequest
	GetUrl() *string
	SetVersion(v string) *GetAuthTokenRequest
	GetVersion() *string
}

type GetAuthTokenRequest struct {
	BizType   *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CmApiCode *int32 `json:"CmApiCode,omitempty" xml:"CmApiCode,omitempty"`
	CtApiCode *int32 `json:"CtApiCode,omitempty" xml:"CtApiCode,omitempty"`
	CuApiCode *int32 `json:"CuApiCode,omitempty" xml:"CuApiCode,omitempty"`
	// The requested domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.aliyundoc.com
	Origin               *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SceneCode            *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The URL of the requested web page.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.aliyundoc.com/
	Url     *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAuthTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *GetAuthTokenRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *GetAuthTokenRequest) GetCmApiCode() *int32 {
	return s.CmApiCode
}

func (s *GetAuthTokenRequest) GetCtApiCode() *int32 {
	return s.CtApiCode
}

func (s *GetAuthTokenRequest) GetCuApiCode() *int32 {
	return s.CuApiCode
}

func (s *GetAuthTokenRequest) GetOrigin() *string {
	return s.Origin
}

func (s *GetAuthTokenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetAuthTokenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetAuthTokenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetAuthTokenRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *GetAuthTokenRequest) GetUrl() *string {
	return s.Url
}

func (s *GetAuthTokenRequest) GetVersion() *string {
	return s.Version
}

func (s *GetAuthTokenRequest) SetBizType(v int32) *GetAuthTokenRequest {
	s.BizType = &v
	return s
}

func (s *GetAuthTokenRequest) SetCmApiCode(v int32) *GetAuthTokenRequest {
	s.CmApiCode = &v
	return s
}

func (s *GetAuthTokenRequest) SetCtApiCode(v int32) *GetAuthTokenRequest {
	s.CtApiCode = &v
	return s
}

func (s *GetAuthTokenRequest) SetCuApiCode(v int32) *GetAuthTokenRequest {
	s.CuApiCode = &v
	return s
}

func (s *GetAuthTokenRequest) SetOrigin(v string) *GetAuthTokenRequest {
	s.Origin = &v
	return s
}

func (s *GetAuthTokenRequest) SetOwnerId(v int64) *GetAuthTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAuthTokenRequest) SetResourceOwnerAccount(v string) *GetAuthTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAuthTokenRequest) SetResourceOwnerId(v int64) *GetAuthTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAuthTokenRequest) SetSceneCode(v string) *GetAuthTokenRequest {
	s.SceneCode = &v
	return s
}

func (s *GetAuthTokenRequest) SetUrl(v string) *GetAuthTokenRequest {
	s.Url = &v
	return s
}

func (s *GetAuthTokenRequest) SetVersion(v string) *GetAuthTokenRequest {
	s.Version = &v
	return s
}

func (s *GetAuthTokenRequest) Validate() error {
	return dara.Validate(s)
}
