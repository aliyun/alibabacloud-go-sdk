// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeasurementDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMeasurementDatas(v *DescribeMeasurementDataResponseBodyMeasurementDatas) *DescribeMeasurementDataResponseBody
	GetMeasurementDatas() *DescribeMeasurementDataResponseBodyMeasurementDatas
	SetRequestId(v string) *DescribeMeasurementDataResponseBody
	GetRequestId() *string
}

type DescribeMeasurementDataResponseBody struct {
	// The metering data returned.
	MeasurementDatas *DescribeMeasurementDataResponseBodyMeasurementDatas `json:"MeasurementDatas,omitempty" xml:"MeasurementDatas,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 125B04C7-3D0D-4245-AF96-14E3758E3F06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeasurementDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeasurementDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseBody) GetMeasurementDatas() *DescribeMeasurementDataResponseBodyMeasurementDatas {
	return s.MeasurementDatas
}

func (s *DescribeMeasurementDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMeasurementDataResponseBody) SetMeasurementDatas(v *DescribeMeasurementDataResponseBodyMeasurementDatas) *DescribeMeasurementDataResponseBody {
	s.MeasurementDatas = v
	return s
}

func (s *DescribeMeasurementDataResponseBody) SetRequestId(v string) *DescribeMeasurementDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMeasurementDataResponseBody) Validate() error {
	if s.MeasurementDatas != nil {
		if err := s.MeasurementDatas.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMeasurementDataResponseBodyMeasurementDatas struct {
	MeasurementData []*DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData `json:"MeasurementData,omitempty" xml:"MeasurementData,omitempty" type:"Repeated"`
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatas) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatas) GetMeasurementData() []*DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	return s.MeasurementData
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatas) SetMeasurementData(v []*DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) *DescribeMeasurementDataResponseBodyMeasurementDatas {
	s.MeasurementData = v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatas) Validate() error {
	if s.MeasurementData != nil {
		for _, item := range s.MeasurementData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData struct {
	// The bandwidth data returned.
	BandWidthFeeDatas *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas `json:"BandWidthFeeDatas,omitempty" xml:"BandWidthFeeDatas,omitempty" type:"Struct"`
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
	// The information about computing resources.
	ResourceFeeData *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData `json:"ResourceFeeData,omitempty" xml:"ResourceFeeData,omitempty" type:"Struct"`
	// Details of the computing resources.
	ResourceFeeDataDetails *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails `json:"ResourceFeeDataDetails,omitempty" xml:"ResourceFeeDataDetails,omitempty" type:"Struct"`
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) GetBandWidthFeeDatas() *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas {
	return s.BandWidthFeeDatas
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) GetChargeModel() *string {
	return s.ChargeModel
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) GetCostCycle() *string {
	return s.CostCycle
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) GetCostEndTime() *string {
	return s.CostEndTime
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) GetCostStartTime() *string {
	return s.CostStartTime
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) GetResourceFeeData() *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData {
	return s.ResourceFeeData
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) GetResourceFeeDataDetails() *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails {
	return s.ResourceFeeDataDetails
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetBandWidthFeeDatas(v *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.BandWidthFeeDatas = v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetChargeModel(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.ChargeModel = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetCostCycle(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.CostCycle = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetCostEndTime(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.CostEndTime = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetCostStartTime(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.CostStartTime = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetResourceFeeData(v *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.ResourceFeeData = v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) SetResourceFeeDataDetails(v *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData {
	s.ResourceFeeDataDetails = v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementData) Validate() error {
	if s.BandWidthFeeDatas != nil {
		if err := s.BandWidthFeeDatas.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceFeeData != nil {
		if err := s.ResourceFeeData.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceFeeDataDetails != nil {
		if err := s.ResourceFeeDataDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas struct {
	BandWidthFeeData []*DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData `json:"BandWidthFeeData,omitempty" xml:"BandWidthFeeData,omitempty" type:"Repeated"`
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) GetBandWidthFeeData() []*DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	return s.BandWidthFeeData
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) SetBandWidthFeeData(v []*DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas {
	s.BandWidthFeeData = v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatas) Validate() error {
	if s.BandWidthFeeData != nil {
		for _, item := range s.BandWidthFeeData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData struct {
	// The code of the bandwidth plan.
	//
	// example:
	//
	// cn-cmcc-1
	CostCode *string `json:"CostCode,omitempty" xml:"CostCode,omitempty"`
	// The name of the bandwidth plan.
	//
	// example:
	//
	// Beijing, Shanghai, and Guangzhou Mobile
	CostName *string `json:"CostName,omitempty" xml:"CostName,omitempty"`
	// The bandwidth consumption. Unit: bit/second.
	//
	// example:
	//
	// 16486
	CostVal *int32 `json:"CostVal,omitempty" xml:"CostVal,omitempty"`
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) GetCostCode() *string {
	return s.CostCode
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) GetCostName() *string {
	return s.CostName
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) GetCostVal() *int32 {
	return s.CostVal
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostCode(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostCode = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostName(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostName = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostVal(v int32) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostVal = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) Validate() error {
	return dara.Validate(s)
}

type DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData struct {
	// The memory size. Unit: GB.
	//
	// example:
	//
	// 24
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The storage capacity. Unit: GB.
	//
	// example:
	//
	// 60
	Storage *int32 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 12
	Vcpu *int32 `json:"Vcpu,omitempty" xml:"Vcpu,omitempty"`
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData) GetStorage() *int32 {
	return s.Storage
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData) GetVcpu() *int32 {
	return s.Vcpu
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData) SetMemory(v int32) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData {
	s.Memory = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData) SetStorage(v int32) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData {
	s.Storage = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData) SetVcpu(v int32) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData {
	s.Vcpu = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeData) Validate() error {
	return dara.Validate(s)
}

type DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails struct {
	ResourceFeeDataDetail []*DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail `json:"ResourceFeeDataDetail,omitempty" xml:"ResourceFeeDataDetail,omitempty" type:"Repeated"`
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails) GetResourceFeeDataDetail() []*DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail {
	return s.ResourceFeeDataDetail
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails) SetResourceFeeDataDetail(v []*DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails {
	s.ResourceFeeDataDetail = v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetails) Validate() error {
	if s.ResourceFeeDataDetail != nil {
		for _, item := range s.ResourceFeeDataDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail struct {
	// The code of the resource.
	//
	// example:
	//
	// vCPU
	CostCode *string `json:"CostCode,omitempty" xml:"CostCode,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// vCPU
	CostName *string `json:"CostName,omitempty" xml:"CostName,omitempty"`
	// The consumption of the resource.
	//
	// 	- Memory unit: GB.
	//
	// 	- CPU unit: vCPU.
	//
	// 	- Storage unit: GB.
	//
	// example:
	//
	// 55
	CostVal *int32 `json:"CostVal,omitempty" xml:"CostVal,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// vCPU
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) GetCostCode() *string {
	return s.CostCode
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) GetCostName() *string {
	return s.CostName
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) GetCostVal() *int32 {
	return s.CostVal
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) SetCostCode(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail {
	s.CostCode = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) SetCostName(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail {
	s.CostName = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) SetCostVal(v int32) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail {
	s.CostVal = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) SetResourceType(v string) *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail {
	s.ResourceType = &v
	return s
}

func (s *DescribeMeasurementDataResponseBodyMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) Validate() error {
	return dara.Validate(s)
}
