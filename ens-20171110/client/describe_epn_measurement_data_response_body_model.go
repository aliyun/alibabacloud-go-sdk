// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEpnMeasurementDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMeasurementDatas(v *DescribeEpnMeasurementDataResponseBodyMeasurementDatas) *DescribeEpnMeasurementDataResponseBody
	GetMeasurementDatas() *DescribeEpnMeasurementDataResponseBodyMeasurementDatas
	SetRequestId(v string) *DescribeEpnMeasurementDataResponseBody
	GetRequestId() *string
}

type DescribeEpnMeasurementDataResponseBody struct {
	// The metering data returned.
	MeasurementDatas *DescribeEpnMeasurementDataResponseBodyMeasurementDatas `json:"MeasurementDatas,omitempty" xml:"MeasurementDatas,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A6328C33-6304-5291-8641-0A00A99D0DD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEpnMeasurementDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponseBody) GetMeasurementDatas() *DescribeEpnMeasurementDataResponseBodyMeasurementDatas {
	return s.MeasurementDatas
}

func (s *DescribeEpnMeasurementDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEpnMeasurementDataResponseBody) SetMeasurementDatas(v *DescribeEpnMeasurementDataResponseBodyMeasurementDatas) *DescribeEpnMeasurementDataResponseBody {
	s.MeasurementDatas = v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBody) SetRequestId(v string) *DescribeEpnMeasurementDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEpnMeasurementDataResponseBodyMeasurementDatas struct {
	MeasurementData []*DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData `json:"MeasurementData,omitempty" xml:"MeasurementData,omitempty" type:"Repeated"`
}

func (s DescribeEpnMeasurementDataResponseBodyMeasurementDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponseBodyMeasurementDatas) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatas) GetMeasurementData() []*DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	return s.MeasurementData
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatas) SetMeasurementData(v []*DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) *DescribeEpnMeasurementDataResponseBodyMeasurementDatas {
	s.MeasurementData = v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatas) Validate() error {
	return dara.Validate(s)
}

type DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData struct {
	// The bandwidth data returned.
	BandWidthFeeDatas *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas `json:"BandWidthFeeDatas,omitempty" xml:"BandWidthFeeDatas,omitempty" type:"Struct"`
	// The metering method. Valid values:
	//
	// 	- ChargeByUnified: unified metering.
	//
	// 	- ChargeByGrade: differential metering.
	//
	// example:
	//
	// ChargeByGrade
	ChargeModel *string `json:"ChargeModel,omitempty" xml:"ChargeModel,omitempty"`
	// The metering cycle.
	//
	// example:
	//
	// 2019-07-30
	CostCycle *string `json:"CostCycle,omitempty" xml:"CostCycle,omitempty"`
	// The end time of the metering cycle.
	//
	// example:
	//
	// 2019-07-30T16:00:00Z
	CostEndTime *string `json:"CostEndTime,omitempty" xml:"CostEndTime,omitempty"`
	// The start time of the metering cycle.
	//
	// example:
	//
	// 2019-07-29T16:00:00Z
	CostStartTime *string `json:"CostStartTime,omitempty" xml:"CostStartTime,omitempty"`
}

func (s DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) GetBandWidthFeeDatas() *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas {
	return s.BandWidthFeeDatas
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) GetChargeModel() *string {
	return s.ChargeModel
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) GetCostCycle() *string {
	return s.CostCycle
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) GetCostEndTime() *string {
	return s.CostEndTime
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) GetCostStartTime() *string {
	return s.CostStartTime
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetBandWidthFeeDatas(v *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.BandWidthFeeDatas = v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetChargeModel(v string) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.ChargeModel = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetCostCycle(v string) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.CostCycle = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetCostEndTime(v string) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.CostEndTime = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetCostStartTime(v string) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.CostStartTime = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementData) Validate() error {
	return dara.Validate(s)
}

type DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas struct {
	BandWidthFeeData []*DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData `json:"BandWidthFeeData,omitempty" xml:"BandWidthFeeData,omitempty" type:"Repeated"`
}

func (s DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) GetBandWidthFeeData() []*DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	return s.BandWidthFeeData
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) SetBandWidthFeeData(v []*DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas {
	s.BandWidthFeeData = v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) Validate() error {
	return dara.Validate(s)
}

type DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData struct {
	// The code of the billable item.
	//
	// example:
	//
	// cn-cmcc-1
	CostCode *string `json:"CostCode,omitempty" xml:"CostCode,omitempty"`
	// The name of the billable item.
	//
	// example:
	//
	// Beijing, Shanghai, and Guangzhou Mobile
	CostName *string `json:"CostName,omitempty" xml:"CostName,omitempty"`
	// Metering method
	//
	// 	- SpeedUp: bandwidth of intelligent acceleration
	//
	// 	- IntranetConnection: internal bandwidth
	//
	// example:
	//
	// SpeedUp
	CostType *string `json:"CostType,omitempty" xml:"CostType,omitempty"`
	// The value of the billable item.
	//
	// example:
	//
	// 16486
	CostVal *int32 `json:"CostVal,omitempty" xml:"CostVal,omitempty"`
	// This parameter is unavailable.
	//
	// example:
	//
	// This parameter is not currently in use.
	IspLine *string `json:"IspLine,omitempty" xml:"IspLine,omitempty"`
}

func (s DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) GetCostCode() *string {
	return s.CostCode
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) GetCostName() *string {
	return s.CostName
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) GetCostType() *string {
	return s.CostType
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) GetCostVal() *int32 {
	return s.CostVal
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) GetIspLine() *string {
	return s.IspLine
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostCode(v string) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostCode = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostName(v string) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostName = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostType(v string) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostType = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostVal(v int32) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostVal = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetIspLine(v string) *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.IspLine = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) Validate() error {
	return dara.Validate(s)
}
