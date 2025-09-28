// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGateVerifyStatisticPublicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationType(v int32) *QueryGateVerifyStatisticPublicRequest
	GetAuthenticationType() *int32
	SetEndDate(v string) *QueryGateVerifyStatisticPublicRequest
	GetEndDate() *string
	SetOsType(v string) *QueryGateVerifyStatisticPublicRequest
	GetOsType() *string
	SetOwnerId(v int64) *QueryGateVerifyStatisticPublicRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryGateVerifyStatisticPublicRequest
	GetResourceOwnerAccount() *string
	SetSceneCode(v string) *QueryGateVerifyStatisticPublicRequest
	GetSceneCode() *string
	SetStartDate(v string) *QueryGateVerifyStatisticPublicRequest
	GetStartDate() *string
}

type QueryGateVerifyStatisticPublicRequest struct {
	// The verification method. Valid values:
	//
	// 	- **1**: one-click logon
	//
	// 	- **2**: phone number verification, including the verification of the phone number used in HTML5 pages
	//
	// 	- **3**: SMS verification
	//
	// 	- **4**: facial recognition
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuthenticationType *int32 `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// The end date. Specify this parameter in the YYYYMMDD format. Example: 20220106.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20220106
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The type of the operating system. Valid values:
	//
	// 	- **Android**
	//
	// 	- **iOS**
	//
	// example:
	//
	// Android
	OsType               *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The service code.
	//
	// example:
	//
	// FC100000038194004
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The start date. Specify this parameter in the YYYYMMDD format. Example: 20220101.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryGateVerifyStatisticPublicRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyStatisticPublicRequest) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticPublicRequest) GetAuthenticationType() *int32 {
	return s.AuthenticationType
}

func (s *QueryGateVerifyStatisticPublicRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryGateVerifyStatisticPublicRequest) GetOsType() *string {
	return s.OsType
}

func (s *QueryGateVerifyStatisticPublicRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryGateVerifyStatisticPublicRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryGateVerifyStatisticPublicRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *QueryGateVerifyStatisticPublicRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryGateVerifyStatisticPublicRequest) SetAuthenticationType(v int32) *QueryGateVerifyStatisticPublicRequest {
	s.AuthenticationType = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetEndDate(v string) *QueryGateVerifyStatisticPublicRequest {
	s.EndDate = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetOsType(v string) *QueryGateVerifyStatisticPublicRequest {
	s.OsType = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetOwnerId(v int64) *QueryGateVerifyStatisticPublicRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetResourceOwnerAccount(v string) *QueryGateVerifyStatisticPublicRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetSceneCode(v string) *QueryGateVerifyStatisticPublicRequest {
	s.SceneCode = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetStartDate(v string) *QueryGateVerifyStatisticPublicRequest {
	s.StartDate = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) Validate() error {
	return dara.Validate(s)
}
