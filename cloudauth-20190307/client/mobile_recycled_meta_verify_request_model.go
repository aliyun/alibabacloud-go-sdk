// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileRecycledMetaVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobile(v string) *MobileRecycledMetaVerifyRequest
	GetMobile() *string
	SetParamType(v string) *MobileRecycledMetaVerifyRequest
	GetParamType() *string
	SetRegisterDate(v string) *MobileRecycledMetaVerifyRequest
	GetRegisterDate() *string
}

type MobileRecycledMetaVerifyRequest struct {
	// This parameter is required.
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// This parameter is required.
	RegisterDate *string `json:"RegisterDate,omitempty" xml:"RegisterDate,omitempty"`
}

func (s MobileRecycledMetaVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s MobileRecycledMetaVerifyRequest) GoString() string {
	return s.String()
}

func (s *MobileRecycledMetaVerifyRequest) GetMobile() *string {
	return s.Mobile
}

func (s *MobileRecycledMetaVerifyRequest) GetParamType() *string {
	return s.ParamType
}

func (s *MobileRecycledMetaVerifyRequest) GetRegisterDate() *string {
	return s.RegisterDate
}

func (s *MobileRecycledMetaVerifyRequest) SetMobile(v string) *MobileRecycledMetaVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *MobileRecycledMetaVerifyRequest) SetParamType(v string) *MobileRecycledMetaVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *MobileRecycledMetaVerifyRequest) SetRegisterDate(v string) *MobileRecycledMetaVerifyRequest {
	s.RegisterDate = &v
	return s
}

func (s *MobileRecycledMetaVerifyRequest) Validate() error {
	return dara.Validate(s)
}
