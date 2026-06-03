// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGateVerifyStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationType(v int32) *QueryGateVerifyStatisticRequest
	GetAuthenticationType() *int32
	SetEndDate(v string) *QueryGateVerifyStatisticRequest
	GetEndDate() *string
	SetOsType(v string) *QueryGateVerifyStatisticRequest
	GetOsType() *string
	SetOwnerId(v int64) *QueryGateVerifyStatisticRequest
	GetOwnerId() *int64
	SetProdCode(v string) *QueryGateVerifyStatisticRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *QueryGateVerifyStatisticRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryGateVerifyStatisticRequest
	GetResourceOwnerId() *int64
	SetSceneCode(v string) *QueryGateVerifyStatisticRequest
	GetSceneCode() *string
	SetStartDate(v string) *QueryGateVerifyStatisticRequest
	GetStartDate() *string
}

type QueryGateVerifyStatisticRequest struct {
	// This parameter is required.
	AuthenticationType *int32 `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// This parameter is required.
	EndDate              *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OsType               *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SceneCode            *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// This parameter is required.
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryGateVerifyStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyStatisticRequest) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticRequest) GetAuthenticationType() *int32 {
	return s.AuthenticationType
}

func (s *QueryGateVerifyStatisticRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryGateVerifyStatisticRequest) GetOsType() *string {
	return s.OsType
}

func (s *QueryGateVerifyStatisticRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryGateVerifyStatisticRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QueryGateVerifyStatisticRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryGateVerifyStatisticRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryGateVerifyStatisticRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *QueryGateVerifyStatisticRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryGateVerifyStatisticRequest) SetAuthenticationType(v int32) *QueryGateVerifyStatisticRequest {
	s.AuthenticationType = &v
	return s
}

func (s *QueryGateVerifyStatisticRequest) SetEndDate(v string) *QueryGateVerifyStatisticRequest {
	s.EndDate = &v
	return s
}

func (s *QueryGateVerifyStatisticRequest) SetOsType(v string) *QueryGateVerifyStatisticRequest {
	s.OsType = &v
	return s
}

func (s *QueryGateVerifyStatisticRequest) SetOwnerId(v int64) *QueryGateVerifyStatisticRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryGateVerifyStatisticRequest) SetProdCode(v string) *QueryGateVerifyStatisticRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryGateVerifyStatisticRequest) SetResourceOwnerAccount(v string) *QueryGateVerifyStatisticRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryGateVerifyStatisticRequest) SetResourceOwnerId(v int64) *QueryGateVerifyStatisticRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryGateVerifyStatisticRequest) SetSceneCode(v string) *QueryGateVerifyStatisticRequest {
	s.SceneCode = &v
	return s
}

func (s *QueryGateVerifyStatisticRequest) SetStartDate(v string) *QueryGateVerifyStatisticRequest {
	s.StartDate = &v
	return s
}

func (s *QueryGateVerifyStatisticRequest) Validate() error {
	return dara.Validate(s)
}
