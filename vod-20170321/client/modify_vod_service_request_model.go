// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVodServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInternetChargeType(v string) *ModifyVodServiceRequest
	GetInternetChargeType() *string
	SetOwnerId(v int64) *ModifyVodServiceRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *ModifyVodServiceRequest
	GetSecurityToken() *string
}

type ModifyVodServiceRequest struct {
	// This parameter is required.
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken      *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyVodServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVodServiceRequest) GoString() string {
	return s.String()
}

func (s *ModifyVodServiceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *ModifyVodServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVodServiceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyVodServiceRequest) SetInternetChargeType(v string) *ModifyVodServiceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *ModifyVodServiceRequest) SetOwnerId(v int64) *ModifyVodServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVodServiceRequest) SetSecurityToken(v string) *ModifyVodServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyVodServiceRequest) Validate() error {
	return dara.Validate(s)
}
