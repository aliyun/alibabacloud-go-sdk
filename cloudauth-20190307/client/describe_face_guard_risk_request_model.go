// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaceGuardRiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *DescribeFaceGuardRiskRequest
	GetBizId() *string
	SetDeviceToken(v string) *DescribeFaceGuardRiskRequest
	GetDeviceToken() *string
	SetOuterOrderNo(v string) *DescribeFaceGuardRiskRequest
	GetOuterOrderNo() *string
	SetProductCode(v string) *DescribeFaceGuardRiskRequest
	GetProductCode() *string
}

type DescribeFaceGuardRiskRequest struct {
	// example:
	//
	// aba9830f471a4335af4612c8adaa91b0
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// McozS1ZWRcRZStlERcZZo_QOytx5jcgZoZJEoRLOxxxxxxx
	DeviceToken *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c35xxxx
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// example:
	//
	// FACE_GUARD
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s DescribeFaceGuardRiskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaceGuardRiskRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaceGuardRiskRequest) GetBizId() *string {
	return s.BizId
}

func (s *DescribeFaceGuardRiskRequest) GetDeviceToken() *string {
	return s.DeviceToken
}

func (s *DescribeFaceGuardRiskRequest) GetOuterOrderNo() *string {
	return s.OuterOrderNo
}

func (s *DescribeFaceGuardRiskRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeFaceGuardRiskRequest) SetBizId(v string) *DescribeFaceGuardRiskRequest {
	s.BizId = &v
	return s
}

func (s *DescribeFaceGuardRiskRequest) SetDeviceToken(v string) *DescribeFaceGuardRiskRequest {
	s.DeviceToken = &v
	return s
}

func (s *DescribeFaceGuardRiskRequest) SetOuterOrderNo(v string) *DescribeFaceGuardRiskRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *DescribeFaceGuardRiskRequest) SetProductCode(v string) *DescribeFaceGuardRiskRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeFaceGuardRiskRequest) Validate() error {
	return dara.Validate(s)
}
