// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceGuardRiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *FaceGuardRiskRequest
	GetBizId() *string
	SetDeviceToken(v string) *FaceGuardRiskRequest
	GetDeviceToken() *string
	SetMerchantBizId(v string) *FaceGuardRiskRequest
	GetMerchantBizId() *string
	SetProductCode(v string) *FaceGuardRiskRequest
	GetProductCode() *string
	SetType(v string) *FaceGuardRiskRequest
	GetType() *string
}

type FaceGuardRiskRequest struct {
	// The customer business ID.
	//
	// example:
	//
	// LMALL20******001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The device token obtained from the Face Guard SDK.
	//
	// example:
	//
	// Tk9SSUQuMS*****************ZDNmNWY5NzQxOW1o
	DeviceToken *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	// The merchant-defined unique business identifier, used for subsequent troubleshooting. The value can contain letters and digits with a maximum length of 32 characters. Make sure the value is unique.
	//
	// example:
	//
	// 0c83ce0101d34eff886b1f7d1cdef67f
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The product code. Set the value to FACE_GUARD.
	//
	// example:
	//
	// FACE_GUARD
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s FaceGuardRiskRequest) String() string {
	return dara.Prettify(s)
}

func (s FaceGuardRiskRequest) GoString() string {
	return s.String()
}

func (s *FaceGuardRiskRequest) GetBizId() *string {
	return s.BizId
}

func (s *FaceGuardRiskRequest) GetDeviceToken() *string {
	return s.DeviceToken
}

func (s *FaceGuardRiskRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *FaceGuardRiskRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *FaceGuardRiskRequest) GetType() *string {
	return s.Type
}

func (s *FaceGuardRiskRequest) SetBizId(v string) *FaceGuardRiskRequest {
	s.BizId = &v
	return s
}

func (s *FaceGuardRiskRequest) SetDeviceToken(v string) *FaceGuardRiskRequest {
	s.DeviceToken = &v
	return s
}

func (s *FaceGuardRiskRequest) SetMerchantBizId(v string) *FaceGuardRiskRequest {
	s.MerchantBizId = &v
	return s
}

func (s *FaceGuardRiskRequest) SetProductCode(v string) *FaceGuardRiskRequest {
	s.ProductCode = &v
	return s
}

func (s *FaceGuardRiskRequest) SetType(v string) *FaceGuardRiskRequest {
	s.Type = &v
	return s
}

func (s *FaceGuardRiskRequest) Validate() error {
	return dara.Validate(s)
}
