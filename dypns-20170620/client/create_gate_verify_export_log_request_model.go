// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGateVerifyExportLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationType(v int32) *CreateGateVerifyExportLogRequest
	GetAuthenticationType() *int32
	SetOsType(v string) *CreateGateVerifyExportLogRequest
	GetOsType() *string
	SetOwnerId(v int64) *CreateGateVerifyExportLogRequest
	GetOwnerId() *int64
	SetQueryMonth(v string) *CreateGateVerifyExportLogRequest
	GetQueryMonth() *string
	SetResourceOwnerAccount(v string) *CreateGateVerifyExportLogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateGateVerifyExportLogRequest
	GetResourceOwnerId() *int64
	SetSceneCode(v string) *CreateGateVerifyExportLogRequest
	GetSceneCode() *string
}

type CreateGateVerifyExportLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuthenticationType *int32 `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// example:
	//
	// Android
	OsType  *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202211
	QueryMonth           *string `json:"QueryMonth,omitempty" xml:"QueryMonth,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// FC100*******4175
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s CreateGateVerifyExportLogRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGateVerifyExportLogRequest) GoString() string {
	return s.String()
}

func (s *CreateGateVerifyExportLogRequest) GetAuthenticationType() *int32 {
	return s.AuthenticationType
}

func (s *CreateGateVerifyExportLogRequest) GetOsType() *string {
	return s.OsType
}

func (s *CreateGateVerifyExportLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateGateVerifyExportLogRequest) GetQueryMonth() *string {
	return s.QueryMonth
}

func (s *CreateGateVerifyExportLogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateGateVerifyExportLogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateGateVerifyExportLogRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *CreateGateVerifyExportLogRequest) SetAuthenticationType(v int32) *CreateGateVerifyExportLogRequest {
	s.AuthenticationType = &v
	return s
}

func (s *CreateGateVerifyExportLogRequest) SetOsType(v string) *CreateGateVerifyExportLogRequest {
	s.OsType = &v
	return s
}

func (s *CreateGateVerifyExportLogRequest) SetOwnerId(v int64) *CreateGateVerifyExportLogRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGateVerifyExportLogRequest) SetQueryMonth(v string) *CreateGateVerifyExportLogRequest {
	s.QueryMonth = &v
	return s
}

func (s *CreateGateVerifyExportLogRequest) SetResourceOwnerAccount(v string) *CreateGateVerifyExportLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateGateVerifyExportLogRequest) SetResourceOwnerId(v int64) *CreateGateVerifyExportLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateGateVerifyExportLogRequest) SetSceneCode(v string) *CreateGateVerifyExportLogRequest {
	s.SceneCode = &v
	return s
}

func (s *CreateGateVerifyExportLogRequest) Validate() error {
	return dara.Validate(s)
}
