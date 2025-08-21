// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandWithdChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandWithTypeInfo(v string) *DescribeBandWithdChargeTypeResponseBody
	GetBandWithTypeInfo() *string
	SetChargeContractType(v string) *DescribeBandWithdChargeTypeResponseBody
	GetChargeContractType() *string
	SetChargeCycleInfo(v string) *DescribeBandWithdChargeTypeResponseBody
	GetChargeCycleInfo() *string
	SetCode(v int32) *DescribeBandWithdChargeTypeResponseBody
	GetCode() *int32
	SetRequestId(v string) *DescribeBandWithdChargeTypeResponseBody
	GetRequestId() *string
}

type DescribeBandWithdChargeTypeResponseBody struct {
	// The information about the metering method.
	//
	// example:
	//
	// { "BandWidthName": "Monthly peak bandwidth of 95", 			"BandWidthType": "95BandwidthByMonth" 		}
	BandWithTypeInfo *string `json:"BandWithTypeInfo,omitempty" xml:"BandWithTypeInfo,omitempty"`
	// The metering type.
	//
	// 	- ChargeByUnified: unified metering.
	//
	// 	- ChargeByGrade: differential metering.
	//
	// example:
	//
	// ChargeByGrade
	ChargeContractType *string `json:"ChargeContractType,omitempty" xml:"ChargeContractType,omitempty"`
	// The metering cycle. Currently, this parameter is empty in the response.
	//
	// example:
	//
	// {null}
	ChargeCycleInfo *string `json:"ChargeCycleInfo,omitempty" xml:"ChargeCycleInfo,omitempty"`
	// The returned service code. A value of 0 indicates that the operation was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 791049FD-49CE-4667-BD6C-F23094DEDA7A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBandWithdChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandWithdChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBandWithdChargeTypeResponseBody) GetBandWithTypeInfo() *string {
	return s.BandWithTypeInfo
}

func (s *DescribeBandWithdChargeTypeResponseBody) GetChargeContractType() *string {
	return s.ChargeContractType
}

func (s *DescribeBandWithdChargeTypeResponseBody) GetChargeCycleInfo() *string {
	return s.ChargeCycleInfo
}

func (s *DescribeBandWithdChargeTypeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeBandWithdChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBandWithdChargeTypeResponseBody) SetBandWithTypeInfo(v string) *DescribeBandWithdChargeTypeResponseBody {
	s.BandWithTypeInfo = &v
	return s
}

func (s *DescribeBandWithdChargeTypeResponseBody) SetChargeContractType(v string) *DescribeBandWithdChargeTypeResponseBody {
	s.ChargeContractType = &v
	return s
}

func (s *DescribeBandWithdChargeTypeResponseBody) SetChargeCycleInfo(v string) *DescribeBandWithdChargeTypeResponseBody {
	s.ChargeCycleInfo = &v
	return s
}

func (s *DescribeBandWithdChargeTypeResponseBody) SetCode(v int32) *DescribeBandWithdChargeTypeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBandWithdChargeTypeResponseBody) SetRequestId(v string) *DescribeBandWithdChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBandWithdChargeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
