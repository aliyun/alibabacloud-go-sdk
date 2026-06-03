// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyVerifySchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCmApiCode(v int64) *CopyVerifySchemeRequest
	GetCmApiCode() *int64
	SetCmNewFlag(v bool) *CopyVerifySchemeRequest
	GetCmNewFlag() *bool
	SetCtApiCode(v int64) *CopyVerifySchemeRequest
	GetCtApiCode() *int64
	SetCtNewFlag(v bool) *CopyVerifySchemeRequest
	GetCtNewFlag() *bool
	SetCuApiCode(v int64) *CopyVerifySchemeRequest
	GetCuApiCode() *int64
	SetCuNewFlag(v bool) *CopyVerifySchemeRequest
	GetCuNewFlag() *bool
	SetEmail(v string) *CopyVerifySchemeRequest
	GetEmail() *string
	SetModelSceneCode(v string) *CopyVerifySchemeRequest
	GetModelSceneCode() *string
	SetOwnerId(v int64) *CopyVerifySchemeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CopyVerifySchemeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CopyVerifySchemeRequest
	GetResourceOwnerId() *int64
}

type CopyVerifySchemeRequest struct {
	// native：1，web：5
	//
	// example:
	//
	// 1
	CmApiCode *int64 `json:"CmApiCode,omitempty" xml:"CmApiCode,omitempty"`
	// true，将在移动侧创建新的AppId，false将直接从模版方案中复制已经创建的移动AppId信息
	//
	// example:
	//
	// true
	CmNewFlag *bool `json:"CmNewFlag,omitempty" xml:"CmNewFlag,omitempty"`
	// native：3，web：8
	//
	// example:
	//
	// 3
	CtApiCode *int64 `json:"CtApiCode,omitempty" xml:"CtApiCode,omitempty"`
	// true，将在电信侧创建新的AppId，false将直接从模版方案中复制已经创建的电信AppId信息
	//
	// example:
	//
	// false
	CtNewFlag *bool `json:"CtNewFlag,omitempty" xml:"CtNewFlag,omitempty"`
	// native：9，web：10
	//
	// example:
	//
	// 9
	CuApiCode *int64 `json:"CuApiCode,omitempty" xml:"CuApiCode,omitempty"`
	// true，将在联通侧创建新的AppId，false将直接从模版方案中复制已经创建的联通AppId信息
	//
	// example:
	//
	// false
	CuNewFlag *bool `json:"CuNewFlag,omitempty" xml:"CuNewFlag,omitempty"`
	// example:
	//
	// ****@xxx.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FC1001212***32
	ModelSceneCode       *string `json:"ModelSceneCode,omitempty" xml:"ModelSceneCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CopyVerifySchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyVerifySchemeRequest) GoString() string {
	return s.String()
}

func (s *CopyVerifySchemeRequest) GetCmApiCode() *int64 {
	return s.CmApiCode
}

func (s *CopyVerifySchemeRequest) GetCmNewFlag() *bool {
	return s.CmNewFlag
}

func (s *CopyVerifySchemeRequest) GetCtApiCode() *int64 {
	return s.CtApiCode
}

func (s *CopyVerifySchemeRequest) GetCtNewFlag() *bool {
	return s.CtNewFlag
}

func (s *CopyVerifySchemeRequest) GetCuApiCode() *int64 {
	return s.CuApiCode
}

func (s *CopyVerifySchemeRequest) GetCuNewFlag() *bool {
	return s.CuNewFlag
}

func (s *CopyVerifySchemeRequest) GetEmail() *string {
	return s.Email
}

func (s *CopyVerifySchemeRequest) GetModelSceneCode() *string {
	return s.ModelSceneCode
}

func (s *CopyVerifySchemeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CopyVerifySchemeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CopyVerifySchemeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CopyVerifySchemeRequest) SetCmApiCode(v int64) *CopyVerifySchemeRequest {
	s.CmApiCode = &v
	return s
}

func (s *CopyVerifySchemeRequest) SetCmNewFlag(v bool) *CopyVerifySchemeRequest {
	s.CmNewFlag = &v
	return s
}

func (s *CopyVerifySchemeRequest) SetCtApiCode(v int64) *CopyVerifySchemeRequest {
	s.CtApiCode = &v
	return s
}

func (s *CopyVerifySchemeRequest) SetCtNewFlag(v bool) *CopyVerifySchemeRequest {
	s.CtNewFlag = &v
	return s
}

func (s *CopyVerifySchemeRequest) SetCuApiCode(v int64) *CopyVerifySchemeRequest {
	s.CuApiCode = &v
	return s
}

func (s *CopyVerifySchemeRequest) SetCuNewFlag(v bool) *CopyVerifySchemeRequest {
	s.CuNewFlag = &v
	return s
}

func (s *CopyVerifySchemeRequest) SetEmail(v string) *CopyVerifySchemeRequest {
	s.Email = &v
	return s
}

func (s *CopyVerifySchemeRequest) SetModelSceneCode(v string) *CopyVerifySchemeRequest {
	s.ModelSceneCode = &v
	return s
}

func (s *CopyVerifySchemeRequest) SetOwnerId(v int64) *CopyVerifySchemeRequest {
	s.OwnerId = &v
	return s
}

func (s *CopyVerifySchemeRequest) SetResourceOwnerAccount(v string) *CopyVerifySchemeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CopyVerifySchemeRequest) SetResourceOwnerId(v int64) *CopyVerifySchemeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CopyVerifySchemeRequest) Validate() error {
	return dara.Validate(s)
}
