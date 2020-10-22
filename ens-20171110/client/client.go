// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeEpnBandwitdhByInternetChargeTypeRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Isp             *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	EnsRegionId     *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	NetworkingModel *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty"`
}

func (s DescribeEpnBandwitdhByInternetChargeTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnBandwitdhByInternetChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetVersion(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.Version = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetStartTime(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetEndTime(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetIsp(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.Isp = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetEnsRegionId(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetNetworkingModel(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.NetworkingModel = &v
	return s
}

type DescribeEpnBandwitdhByInternetChargeTypeResponse struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty" require:"true"`
	BandwidthValue     *int64  `json:"BandwidthValue,omitempty" xml:"BandwidthValue,omitempty" require:"true"`
	TimeStamp          *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
}

func (s DescribeEpnBandwitdhByInternetChargeTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnBandwitdhByInternetChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponse) SetRequestId(v string) *DescribeEpnBandwitdhByInternetChargeTypeResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponse) SetInternetChargeType(v string) *DescribeEpnBandwitdhByInternetChargeTypeResponse {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponse) SetBandwidthValue(v int64) *DescribeEpnBandwitdhByInternetChargeTypeResponse {
	s.BandwidthValue = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeResponse) SetTimeStamp(v string) *DescribeEpnBandwitdhByInternetChargeTypeResponse {
	s.TimeStamp = &v
	return s
}

type DescribeEpnBandWidthDataRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	EnsRegionId     *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Period          *string `json:"Period,omitempty" xml:"Period,omitempty" require:"true"`
	Isp             *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	NetworkingModel *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty"`
	EPNInstanceId   *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
}

func (s DescribeEpnBandWidthDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnBandWidthDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandWidthDataRequest) SetVersion(v string) *DescribeEpnBandWidthDataRequest {
	s.Version = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetEnsRegionId(v string) *DescribeEpnBandWidthDataRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetInstanceId(v string) *DescribeEpnBandWidthDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetStartTime(v string) *DescribeEpnBandWidthDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetEndTime(v string) *DescribeEpnBandWidthDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetPeriod(v string) *DescribeEpnBandWidthDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetIsp(v string) *DescribeEpnBandWidthDataRequest {
	s.Isp = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetNetworkingModel(v string) *DescribeEpnBandWidthDataRequest {
	s.NetworkingModel = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetEPNInstanceId(v string) *DescribeEpnBandWidthDataRequest {
	s.EPNInstanceId = &v
	return s
}

type DescribeEpnBandWidthDataResponse struct {
	RequestId   *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	MonitorData *DescribeEpnBandWidthDataResponseMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" require:"true" type:"Struct"`
}

func (s DescribeEpnBandWidthDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnBandWidthDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandWidthDataResponse) SetRequestId(v string) *DescribeEpnBandWidthDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponse) SetMonitorData(v *DescribeEpnBandWidthDataResponseMonitorData) *DescribeEpnBandWidthDataResponse {
	s.MonitorData = v
	return s
}

type DescribeEpnBandWidthDataResponseMonitorData struct {
	MaxDownBandWidth     *int64                                                             `json:"MaxDownBandWidth,omitempty" xml:"MaxDownBandWidth,omitempty" require:"true"`
	MaxUpBandWidth       *int64                                                             `json:"MaxUpBandWidth,omitempty" xml:"MaxUpBandWidth,omitempty" require:"true"`
	BandWidthMonitorData []*DescribeEpnBandWidthDataResponseMonitorDataBandWidthMonitorData `json:"BandWidthMonitorData,omitempty" xml:"BandWidthMonitorData,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeEpnBandWidthDataResponseMonitorData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnBandWidthDataResponseMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandWidthDataResponseMonitorData) SetMaxDownBandWidth(v int64) *DescribeEpnBandWidthDataResponseMonitorData {
	s.MaxDownBandWidth = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseMonitorData) SetMaxUpBandWidth(v int64) *DescribeEpnBandWidthDataResponseMonitorData {
	s.MaxUpBandWidth = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseMonitorData) SetBandWidthMonitorData(v []*DescribeEpnBandWidthDataResponseMonitorDataBandWidthMonitorData) *DescribeEpnBandWidthDataResponseMonitorData {
	s.BandWidthMonitorData = v
	return s
}

type DescribeEpnBandWidthDataResponseMonitorDataBandWidthMonitorData struct {
	UpBandWidth   *int64  `json:"UpBandWidth,omitempty" xml:"UpBandWidth,omitempty" require:"true"`
	DownBandWidth *int64  `json:"DownBandWidth,omitempty" xml:"DownBandWidth,omitempty" require:"true"`
	InternetTX    *int64  `json:"InternetTX,omitempty" xml:"InternetTX,omitempty" require:"true"`
	InternetRX    *int64  `json:"InternetRX,omitempty" xml:"InternetRX,omitempty" require:"true"`
	TimeStamp     *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
}

func (s DescribeEpnBandWidthDataResponseMonitorDataBandWidthMonitorData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnBandWidthDataResponseMonitorDataBandWidthMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandWidthDataResponseMonitorDataBandWidthMonitorData) SetUpBandWidth(v int64) *DescribeEpnBandWidthDataResponseMonitorDataBandWidthMonitorData {
	s.UpBandWidth = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseMonitorDataBandWidthMonitorData) SetDownBandWidth(v int64) *DescribeEpnBandWidthDataResponseMonitorDataBandWidthMonitorData {
	s.DownBandWidth = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseMonitorDataBandWidthMonitorData) SetInternetTX(v int64) *DescribeEpnBandWidthDataResponseMonitorDataBandWidthMonitorData {
	s.InternetTX = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseMonitorDataBandWidthMonitorData) SetInternetRX(v int64) *DescribeEpnBandWidthDataResponseMonitorDataBandWidthMonitorData {
	s.InternetRX = &v
	return s
}

func (s *DescribeEpnBandWidthDataResponseMonitorDataBandWidthMonitorData) SetTimeStamp(v string) *DescribeEpnBandWidthDataResponseMonitorDataBandWidthMonitorData {
	s.TimeStamp = &v
	return s
}

type DescribeEpnMeasurementDataRequest struct {
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
}

func (s DescribeEpnMeasurementDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnMeasurementDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataRequest) SetVersion(v string) *DescribeEpnMeasurementDataRequest {
	s.Version = &v
	return s
}

func (s *DescribeEpnMeasurementDataRequest) SetStartDate(v string) *DescribeEpnMeasurementDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeEpnMeasurementDataRequest) SetEndDate(v string) *DescribeEpnMeasurementDataRequest {
	s.EndDate = &v
	return s
}

type DescribeEpnMeasurementDataResponse struct {
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	MeasurementDatas *DescribeEpnMeasurementDataResponseMeasurementDatas `json:"MeasurementDatas,omitempty" xml:"MeasurementDatas,omitempty" require:"true" type:"Struct"`
}

func (s DescribeEpnMeasurementDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponse) SetRequestId(v string) *DescribeEpnMeasurementDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponse) SetMeasurementDatas(v *DescribeEpnMeasurementDataResponseMeasurementDatas) *DescribeEpnMeasurementDataResponse {
	s.MeasurementDatas = v
	return s
}

type DescribeEpnMeasurementDataResponseMeasurementDatas struct {
	MeasurementData []*DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementData `json:"MeasurementData,omitempty" xml:"MeasurementData,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeEpnMeasurementDataResponseMeasurementDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponseMeasurementDatas) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponseMeasurementDatas) SetMeasurementData(v []*DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementData) *DescribeEpnMeasurementDataResponseMeasurementDatas {
	s.MeasurementData = v
	return s
}

type DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementData struct {
	ChargeModel       *string                                                                             `json:"ChargeModel,omitempty" xml:"ChargeModel,omitempty" require:"true"`
	CostCycle         *string                                                                             `json:"CostCycle,omitempty" xml:"CostCycle,omitempty" require:"true"`
	CostStartTime     *string                                                                             `json:"CostStartTime,omitempty" xml:"CostStartTime,omitempty" require:"true"`
	CostEndTime       *string                                                                             `json:"CostEndTime,omitempty" xml:"CostEndTime,omitempty" require:"true"`
	BandWidthFeeDatas *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatas `json:"BandWidthFeeDatas,omitempty" xml:"BandWidthFeeDatas,omitempty" require:"true" type:"Struct"`
}

func (s DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementData) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementData) SetChargeModel(v string) *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementData {
	s.ChargeModel = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementData) SetCostCycle(v string) *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementData {
	s.CostCycle = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementData) SetCostStartTime(v string) *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementData {
	s.CostStartTime = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementData) SetCostEndTime(v string) *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementData {
	s.CostEndTime = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementData) SetBandWidthFeeDatas(v *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatas) *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementData {
	s.BandWidthFeeDatas = v
	return s
}

type DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatas struct {
	BandWidthFeeData []*DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData `json:"BandWidthFeeData,omitempty" xml:"BandWidthFeeData,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatas) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatas) SetBandWidthFeeData(v []*DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatas {
	s.BandWidthFeeData = v
	return s
}

type DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData struct {
	CostVal  *int    `json:"CostVal,omitempty" xml:"CostVal,omitempty" require:"true"`
	CostCode *string `json:"CostCode,omitempty" xml:"CostCode,omitempty" require:"true"`
	CostName *string `json:"CostName,omitempty" xml:"CostName,omitempty" require:"true"`
	CostType *string `json:"CostType,omitempty" xml:"CostType,omitempty" require:"true"`
	IspLine  *string `json:"IspLine,omitempty" xml:"IspLine,omitempty" require:"true"`
}

func (s DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostVal(v int) *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostVal = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostCode(v string) *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostCode = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostName(v string) *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostName = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostType(v string) *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostType = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetIspLine(v string) *DescribeEpnMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.IspLine = &v
	return s
}

type DescribeNetworkInterfacesRequest struct {
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	EnsRegionId      *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	PageNumber       *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeNetworkInterfacesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkInterfacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesRequest) SetInstanceId(v string) *DescribeNetworkInterfacesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetVSwitchId(v string) *DescribeNetworkInterfacesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetEnsRegionId(v string) *DescribeNetworkInterfacesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetPrimaryIpAddress(v string) *DescribeNetworkInterfacesRequest {
	s.PrimaryIpAddress = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetPageNumber(v string) *DescribeNetworkInterfacesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetPageSize(v string) *DescribeNetworkInterfacesRequest {
	s.PageSize = &v
	return s
}

type DescribeNetworkInterfacesResponse struct {
	RequestId            *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount           *int                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber           *int                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize             *int                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	NetworkInterfaceSets *DescribeNetworkInterfacesResponseNetworkInterfaceSets `json:"NetworkInterfaceSets,omitempty" xml:"NetworkInterfaceSets,omitempty" require:"true" type:"Struct"`
}

func (s DescribeNetworkInterfacesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkInterfacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponse) SetRequestId(v string) *DescribeNetworkInterfacesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponse) SetTotalCount(v int) *DescribeNetworkInterfacesResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeNetworkInterfacesResponse) SetPageNumber(v int) *DescribeNetworkInterfacesResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkInterfacesResponse) SetPageSize(v int) *DescribeNetworkInterfacesResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworkInterfacesResponse) SetNetworkInterfaceSets(v *DescribeNetworkInterfacesResponseNetworkInterfaceSets) *DescribeNetworkInterfacesResponse {
	s.NetworkInterfaceSets = v
	return s
}

type DescribeNetworkInterfacesResponseNetworkInterfaceSets struct {
	NetworkInterfaceSet []*DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet `json:"NetworkInterfaceSet,omitempty" xml:"NetworkInterfaceSet,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeNetworkInterfacesResponseNetworkInterfaceSets) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseNetworkInterfaceSets) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseNetworkInterfaceSets) SetNetworkInterfaceSet(v []*DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet) *DescribeNetworkInterfacesResponseNetworkInterfaceSets {
	s.NetworkInterfaceSet = v
	return s
}

type DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet struct {
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	PrimaryIp          *string `json:"PrimaryIp,omitempty" xml:"PrimaryIp,omitempty" require:"true"`
	EnsRegionId        *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty" require:"true"`
	MacAddress         *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty" require:"true"`
	CreationTime       *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	PrimaryIpType      *string `json:"PrimaryIpType,omitempty" xml:"PrimaryIpType,omitempty" require:"true"`
}

func (s DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet) SetStatus(v string) *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet {
	s.Status = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet) SetPrimaryIp(v string) *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet {
	s.PrimaryIp = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet) SetEnsRegionId(v string) *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet) SetInstanceId(v string) *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet) SetVSwitchId(v string) *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet {
	s.VSwitchId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet) SetNetworkInterfaceId(v string) *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet) SetMacAddress(v string) *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet {
	s.MacAddress = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet) SetCreationTime(v string) *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet {
	s.CreationTime = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet) SetPrimaryIpType(v string) *DescribeNetworkInterfacesResponseNetworkInterfaceSetsNetworkInterfaceSet {
	s.PrimaryIpType = &v
	return s
}

type CreateEPInstanceRequest struct {
	EPNInstanceType         *string `json:"EPNInstanceType,omitempty" xml:"EPNInstanceType,omitempty" require:"true"`
	EPNInstanceName         *string `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty" require:"true"`
	InternetChargeType      *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty" require:"true"`
	NetworkingModel         *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty" require:"true"`
	InternetMaxBandwidthOut *int    `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty" require:"true"`
}

func (s CreateEPInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEPInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateEPInstanceRequest) SetEPNInstanceType(v string) *CreateEPInstanceRequest {
	s.EPNInstanceType = &v
	return s
}

func (s *CreateEPInstanceRequest) SetEPNInstanceName(v string) *CreateEPInstanceRequest {
	s.EPNInstanceName = &v
	return s
}

func (s *CreateEPInstanceRequest) SetInternetChargeType(v string) *CreateEPInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateEPInstanceRequest) SetNetworkingModel(v string) *CreateEPInstanceRequest {
	s.NetworkingModel = &v
	return s
}

func (s *CreateEPInstanceRequest) SetInternetMaxBandwidthOut(v int) *CreateEPInstanceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

type CreateEPInstanceResponse struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty" require:"true"`
}

func (s CreateEPInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEPInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateEPInstanceResponse) SetRequestId(v string) *CreateEPInstanceResponse {
	s.RequestId = &v
	return s
}

func (s *CreateEPInstanceResponse) SetEPNInstanceId(v string) *CreateEPInstanceResponse {
	s.EPNInstanceId = &v
	return s
}

type ModifyImageSharePermissionRequest struct {
	ImageId        *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	AddAccounts    *string `json:"AddAccounts,omitempty" xml:"AddAccounts,omitempty"`
	RemoveAccounts *string `json:"RemoveAccounts,omitempty" xml:"RemoveAccounts,omitempty"`
}

func (s ModifyImageSharePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageSharePermissionRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageSharePermissionRequest) SetImageId(v string) *ModifyImageSharePermissionRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetAddAccounts(v string) *ModifyImageSharePermissionRequest {
	s.AddAccounts = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetRemoveAccounts(v string) *ModifyImageSharePermissionRequest {
	s.RemoveAccounts = &v
	return s
}

type ModifyImageSharePermissionResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyImageSharePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageSharePermissionResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageSharePermissionResponse) SetRequestId(v string) *ModifyImageSharePermissionResponse {
	s.RequestId = &v
	return s
}

type AddNetworkInterfaceToInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Networks   *string `json:"Networks,omitempty" xml:"Networks,omitempty" require:"true"`
}

func (s AddNetworkInterfaceToInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddNetworkInterfaceToInstanceRequest) GoString() string {
	return s.String()
}

func (s *AddNetworkInterfaceToInstanceRequest) SetInstanceId(v string) *AddNetworkInterfaceToInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *AddNetworkInterfaceToInstanceRequest) SetNetworks(v string) *AddNetworkInterfaceToInstanceRequest {
	s.Networks = &v
	return s
}

type AddNetworkInterfaceToInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddNetworkInterfaceToInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddNetworkInterfaceToInstanceResponse) GoString() string {
	return s.String()
}

func (s *AddNetworkInterfaceToInstanceResponse) SetRequestId(v string) *AddNetworkInterfaceToInstanceResponse {
	s.RequestId = &v
	return s
}

type DescribeImageSharePermissionRequest struct {
	ImageId    *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	AliyunId   *int64  `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
}

func (s DescribeImageSharePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageSharePermissionRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionRequest) SetImageId(v string) *DescribeImageSharePermissionRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) SetPageNumber(v string) *DescribeImageSharePermissionRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) SetPageSize(v string) *DescribeImageSharePermissionRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) SetAliyunId(v int64) *DescribeImageSharePermissionRequest {
	s.AliyunId = &v
	return s
}

type DescribeImageSharePermissionResponse struct {
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount *int                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber *int                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	ImageId    *string                                       `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	Accounts   *DescribeImageSharePermissionResponseAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" require:"true" type:"Struct"`
}

func (s DescribeImageSharePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageSharePermissionResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionResponse) SetRequestId(v string) *DescribeImageSharePermissionResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeImageSharePermissionResponse) SetTotalCount(v int) *DescribeImageSharePermissionResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageSharePermissionResponse) SetPageNumber(v int) *DescribeImageSharePermissionResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeImageSharePermissionResponse) SetPageSize(v int) *DescribeImageSharePermissionResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeImageSharePermissionResponse) SetImageId(v string) *DescribeImageSharePermissionResponse {
	s.ImageId = &v
	return s
}

func (s *DescribeImageSharePermissionResponse) SetAccounts(v *DescribeImageSharePermissionResponseAccounts) *DescribeImageSharePermissionResponse {
	s.Accounts = v
	return s
}

type DescribeImageSharePermissionResponseAccounts struct {
	Account []*string `json:"Account,omitempty" xml:"Account,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeImageSharePermissionResponseAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageSharePermissionResponseAccounts) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionResponseAccounts) SetAccount(v []*string) *DescribeImageSharePermissionResponseAccounts {
	s.Account = v
	return s
}

type RemoveVSwitchesFromEpnInstanceRequest struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty" require:"true"`
	VSwitchesInfo *string `json:"VSwitchesInfo,omitempty" xml:"VSwitchesInfo,omitempty" require:"true"`
}

func (s RemoveVSwitchesFromEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveVSwitchesFromEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *RemoveVSwitchesFromEpnInstanceRequest) SetEPNInstanceId(v string) *RemoveVSwitchesFromEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *RemoveVSwitchesFromEpnInstanceRequest) SetVSwitchesInfo(v string) *RemoveVSwitchesFromEpnInstanceRequest {
	s.VSwitchesInfo = &v
	return s
}

type RemoveVSwitchesFromEpnInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RemoveVSwitchesFromEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveVSwitchesFromEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *RemoveVSwitchesFromEpnInstanceResponse) SetRequestId(v string) *RemoveVSwitchesFromEpnInstanceResponse {
	s.RequestId = &v
	return s
}

type DistApplicationDataRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	DistStrategy *string `json:"DistStrategy,omitempty" xml:"DistStrategy,omitempty"`
}

func (s DistApplicationDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DistApplicationDataRequest) GoString() string {
	return s.String()
}

func (s *DistApplicationDataRequest) SetAppId(v string) *DistApplicationDataRequest {
	s.AppId = &v
	return s
}

func (s *DistApplicationDataRequest) SetData(v string) *DistApplicationDataRequest {
	s.Data = &v
	return s
}

func (s *DistApplicationDataRequest) SetDistStrategy(v string) *DistApplicationDataRequest {
	s.DistStrategy = &v
	return s
}

type DistApplicationDataResponse struct {
	RequestId              *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DistInstanceTotalCount *int                                        `json:"DistInstanceTotalCount,omitempty" xml:"DistInstanceTotalCount,omitempty" require:"true"`
	DistResults            *DistApplicationDataResponseDistResults     `json:"DistResults,omitempty" xml:"DistResults,omitempty" require:"true" type:"Struct"`
	DistInstanceIds        *DistApplicationDataResponseDistInstanceIds `json:"DistInstanceIds,omitempty" xml:"DistInstanceIds,omitempty" require:"true" type:"Struct"`
}

func (s DistApplicationDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DistApplicationDataResponse) GoString() string {
	return s.String()
}

func (s *DistApplicationDataResponse) SetRequestId(v string) *DistApplicationDataResponse {
	s.RequestId = &v
	return s
}

func (s *DistApplicationDataResponse) SetDistInstanceTotalCount(v int) *DistApplicationDataResponse {
	s.DistInstanceTotalCount = &v
	return s
}

func (s *DistApplicationDataResponse) SetDistResults(v *DistApplicationDataResponseDistResults) *DistApplicationDataResponse {
	s.DistResults = v
	return s
}

func (s *DistApplicationDataResponse) SetDistInstanceIds(v *DistApplicationDataResponseDistInstanceIds) *DistApplicationDataResponse {
	s.DistInstanceIds = v
	return s
}

type DistApplicationDataResponseDistResults struct {
	DistResult []*DistApplicationDataResponseDistResultsDistResult `json:"DistResult,omitempty" xml:"DistResult,omitempty" require:"true" type:"Repeated"`
}

func (s DistApplicationDataResponseDistResults) String() string {
	return tea.Prettify(s)
}

func (s DistApplicationDataResponseDistResults) GoString() string {
	return s.String()
}

func (s *DistApplicationDataResponseDistResults) SetDistResult(v []*DistApplicationDataResponseDistResultsDistResult) *DistApplicationDataResponseDistResults {
	s.DistResult = v
	return s
}

type DistApplicationDataResponseDistResultsDistResult struct {
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	ResultDescrip *string `json:"ResultDescrip,omitempty" xml:"ResultDescrip,omitempty" require:"true"`
	ResultCode    *int    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s DistApplicationDataResponseDistResultsDistResult) String() string {
	return tea.Prettify(s)
}

func (s DistApplicationDataResponseDistResultsDistResult) GoString() string {
	return s.String()
}

func (s *DistApplicationDataResponseDistResultsDistResult) SetVersion(v string) *DistApplicationDataResponseDistResultsDistResult {
	s.Version = &v
	return s
}

func (s *DistApplicationDataResponseDistResultsDistResult) SetResultDescrip(v string) *DistApplicationDataResponseDistResultsDistResult {
	s.ResultDescrip = &v
	return s
}

func (s *DistApplicationDataResponseDistResultsDistResult) SetResultCode(v int) *DistApplicationDataResponseDistResultsDistResult {
	s.ResultCode = &v
	return s
}

func (s *DistApplicationDataResponseDistResultsDistResult) SetName(v string) *DistApplicationDataResponseDistResultsDistResult {
	s.Name = &v
	return s
}

type DistApplicationDataResponseDistInstanceIds struct {
	DistInstanceId []*string `json:"DistInstanceId,omitempty" xml:"DistInstanceId,omitempty" require:"true" type:"Repeated"`
}

func (s DistApplicationDataResponseDistInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s DistApplicationDataResponseDistInstanceIds) GoString() string {
	return s.String()
}

func (s *DistApplicationDataResponseDistInstanceIds) SetDistInstanceId(v []*string) *DistApplicationDataResponseDistInstanceIds {
	s.DistInstanceId = v
	return s
}

type JoinVSwitchesToEpnInstanceRequest struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty" require:"true"`
	VSwitchesInfo *string `json:"VSwitchesInfo,omitempty" xml:"VSwitchesInfo,omitempty" require:"true"`
}

func (s JoinVSwitchesToEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinVSwitchesToEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *JoinVSwitchesToEpnInstanceRequest) SetEPNInstanceId(v string) *JoinVSwitchesToEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *JoinVSwitchesToEpnInstanceRequest) SetVSwitchesInfo(v string) *JoinVSwitchesToEpnInstanceRequest {
	s.VSwitchesInfo = &v
	return s
}

type JoinVSwitchesToEpnInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s JoinVSwitchesToEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinVSwitchesToEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *JoinVSwitchesToEpnInstanceResponse) SetRequestId(v string) *JoinVSwitchesToEpnInstanceResponse {
	s.RequestId = &v
	return s
}

type DescribeApplicationRequest struct {
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	AppVersions         *string `json:"AppVersions,omitempty" xml:"AppVersions,omitempty"`
	Level               *string `json:"Level,omitempty" xml:"Level,omitempty"`
	OutDetailStatParams *string `json:"OutDetailStatParams,omitempty" xml:"OutDetailStatParams,omitempty"`
}

func (s DescribeApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationRequest) SetAppId(v string) *DescribeApplicationRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationRequest) SetAppVersions(v string) *DescribeApplicationRequest {
	s.AppVersions = &v
	return s
}

func (s *DescribeApplicationRequest) SetLevel(v string) *DescribeApplicationRequest {
	s.Level = &v
	return s
}

func (s *DescribeApplicationRequest) SetOutDetailStatParams(v string) *DescribeApplicationRequest {
	s.OutDetailStatParams = &v
	return s
}

type DescribeApplicationResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Application *string `json:"Application,omitempty" xml:"Application,omitempty" require:"true"`
}

func (s DescribeApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationResponse) SetRequestId(v string) *DescribeApplicationResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationResponse) SetApplication(v string) *DescribeApplicationResponse {
	s.Application = &v
	return s
}

type DescribeDataPushResultRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	DataNames    *string `json:"DataNames,omitempty" xml:"DataNames,omitempty"`
	DataVersions *string `json:"DataVersions,omitempty" xml:"DataVersions,omitempty"`
	MinDate      *string `json:"MinDate,omitempty" xml:"MinDate,omitempty"`
	MaxDate      *string `json:"MaxDate,omitempty" xml:"MaxDate,omitempty"`
	PageNumber   *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionIds    *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
}

func (s DescribeDataPushResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultRequest) SetAppId(v string) *DescribeDataPushResultRequest {
	s.AppId = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetDataNames(v string) *DescribeDataPushResultRequest {
	s.DataNames = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetDataVersions(v string) *DescribeDataPushResultRequest {
	s.DataVersions = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetMinDate(v string) *DescribeDataPushResultRequest {
	s.MinDate = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetMaxDate(v string) *DescribeDataPushResultRequest {
	s.MaxDate = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetPageNumber(v int) *DescribeDataPushResultRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetPageSize(v int) *DescribeDataPushResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetRegionIds(v string) *DescribeDataPushResultRequest {
	s.RegionIds = &v
	return s
}

type DescribeDataPushResultResponse struct {
	RequestId   *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount  *int                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber  *int                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize    *int                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	PushResults *DescribeDataPushResultResponsePushResults `json:"PushResults,omitempty" xml:"PushResults,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDataPushResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponse) SetRequestId(v string) *DescribeDataPushResultResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDataPushResultResponse) SetTotalCount(v int) *DescribeDataPushResultResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataPushResultResponse) SetPageNumber(v int) *DescribeDataPushResultResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataPushResultResponse) SetPageSize(v int) *DescribeDataPushResultResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeDataPushResultResponse) SetPushResults(v *DescribeDataPushResultResponsePushResults) *DescribeDataPushResultResponse {
	s.PushResults = v
	return s
}

type DescribeDataPushResultResponsePushResults struct {
	PushResult []*DescribeDataPushResultResponsePushResultsPushResult `json:"PushResult,omitempty" xml:"PushResult,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDataPushResultResponsePushResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultResponsePushResults) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponsePushResults) SetPushResult(v []*DescribeDataPushResultResponsePushResultsPushResult) *DescribeDataPushResultResponsePushResults {
	s.PushResult = v
	return s
}

type DescribeDataPushResultResponsePushResultsPushResult struct {
	Name        *string                                                         `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Version     *string                                                         `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	StatusStatS *DescribeDataPushResultResponsePushResultsPushResultStatusStatS `json:"StatusStatS,omitempty" xml:"StatusStatS,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDataPushResultResponsePushResultsPushResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultResponsePushResultsPushResult) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponsePushResultsPushResult) SetName(v string) *DescribeDataPushResultResponsePushResultsPushResult {
	s.Name = &v
	return s
}

func (s *DescribeDataPushResultResponsePushResultsPushResult) SetVersion(v string) *DescribeDataPushResultResponsePushResultsPushResult {
	s.Version = &v
	return s
}

func (s *DescribeDataPushResultResponsePushResultsPushResult) SetStatusStatS(v *DescribeDataPushResultResponsePushResultsPushResultStatusStatS) *DescribeDataPushResultResponsePushResultsPushResult {
	s.StatusStatS = v
	return s
}

type DescribeDataPushResultResponsePushResultsPushResultStatusStatS struct {
	StatusStat []*DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStat `json:"StatusStat,omitempty" xml:"StatusStat,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDataPushResultResponsePushResultsPushResultStatusStatS) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultResponsePushResultsPushResultStatusStatS) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponsePushResultsPushResultStatusStatS) SetStatusStat(v []*DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStat) *DescribeDataPushResultResponsePushResultsPushResultStatusStatS {
	s.StatusStat = v
	return s
}

type DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStat struct {
	Status        *string                                                                            `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	RegionIdCount *int                                                                               `json:"RegionIdCount,omitempty" xml:"RegionIdCount,omitempty" require:"true"`
	RegionIds     *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIds `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStat) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStat) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStat) SetStatus(v string) *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStat {
	s.Status = &v
	return s
}

func (s *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStat) SetRegionIdCount(v int) *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStat {
	s.RegionIdCount = &v
	return s
}

func (s *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStat) SetRegionIds(v *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIds) *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStat {
	s.RegionIds = v
	return s
}

type DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIds struct {
	RegionId []*DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIdsRegionId `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIds) SetRegionId(v []*DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIds {
	s.RegionId = v
	return s
}

type DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIdsRegionId struct {
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	StatusDescrip *string `json:"StatusDescrip,omitempty" xml:"StatusDescrip,omitempty" require:"true"`
}

func (s DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) SetStartTime(v string) *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIdsRegionId {
	s.StartTime = &v
	return s
}

func (s *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) SetUpdateTime(v string) *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIdsRegionId {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) SetRegionId(v string) *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIdsRegionId {
	s.RegionId = &v
	return s
}

func (s *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIdsRegionId) SetStatusDescrip(v string) *DescribeDataPushResultResponsePushResultsPushResultStatusStatSStatusStatRegionIdsRegionId {
	s.StatusDescrip = &v
	return s
}

type PushApplicationDataRequest struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Timeout      *int    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	PushStrategy *string `json:"PushStrategy,omitempty" xml:"PushStrategy,omitempty"`
}

func (s PushApplicationDataRequest) String() string {
	return tea.Prettify(s)
}

func (s PushApplicationDataRequest) GoString() string {
	return s.String()
}

func (s *PushApplicationDataRequest) SetData(v string) *PushApplicationDataRequest {
	s.Data = &v
	return s
}

func (s *PushApplicationDataRequest) SetAppId(v string) *PushApplicationDataRequest {
	s.AppId = &v
	return s
}

func (s *PushApplicationDataRequest) SetTimeout(v int) *PushApplicationDataRequest {
	s.Timeout = &v
	return s
}

func (s *PushApplicationDataRequest) SetPushStrategy(v string) *PushApplicationDataRequest {
	s.PushStrategy = &v
	return s
}

type PushApplicationDataResponse struct {
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PushResults *PushApplicationDataResponsePushResults `json:"PushResults,omitempty" xml:"PushResults,omitempty" require:"true" type:"Struct"`
}

func (s PushApplicationDataResponse) String() string {
	return tea.Prettify(s)
}

func (s PushApplicationDataResponse) GoString() string {
	return s.String()
}

func (s *PushApplicationDataResponse) SetRequestId(v string) *PushApplicationDataResponse {
	s.RequestId = &v
	return s
}

func (s *PushApplicationDataResponse) SetPushResults(v *PushApplicationDataResponsePushResults) *PushApplicationDataResponse {
	s.PushResults = v
	return s
}

type PushApplicationDataResponsePushResults struct {
	PushResult []*PushApplicationDataResponsePushResultsPushResult `json:"PushResult,omitempty" xml:"PushResult,omitempty" require:"true" type:"Repeated"`
}

func (s PushApplicationDataResponsePushResults) String() string {
	return tea.Prettify(s)
}

func (s PushApplicationDataResponsePushResults) GoString() string {
	return s.String()
}

func (s *PushApplicationDataResponsePushResults) SetPushResult(v []*PushApplicationDataResponsePushResultsPushResult) *PushApplicationDataResponsePushResults {
	s.PushResult = v
	return s
}

type PushApplicationDataResponsePushResultsPushResult struct {
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	ResultDescrip *string `json:"ResultDescrip,omitempty" xml:"ResultDescrip,omitempty" require:"true"`
	ResultCode    *int    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s PushApplicationDataResponsePushResultsPushResult) String() string {
	return tea.Prettify(s)
}

func (s PushApplicationDataResponsePushResultsPushResult) GoString() string {
	return s.String()
}

func (s *PushApplicationDataResponsePushResultsPushResult) SetVersion(v string) *PushApplicationDataResponsePushResultsPushResult {
	s.Version = &v
	return s
}

func (s *PushApplicationDataResponsePushResultsPushResult) SetResultDescrip(v string) *PushApplicationDataResponsePushResultsPushResult {
	s.ResultDescrip = &v
	return s
}

func (s *PushApplicationDataResponsePushResultsPushResult) SetResultCode(v int) *PushApplicationDataResponsePushResultsPushResult {
	s.ResultCode = &v
	return s
}

func (s *PushApplicationDataResponsePushResultsPushResult) SetName(v string) *PushApplicationDataResponsePushResultsPushResult {
	s.Name = &v
	return s
}

type UpgradeApplicationRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty" require:"true"`
	Timeout  *int    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpgradeApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpgradeApplicationRequest) SetAppId(v string) *UpgradeApplicationRequest {
	s.AppId = &v
	return s
}

func (s *UpgradeApplicationRequest) SetTemplate(v string) *UpgradeApplicationRequest {
	s.Template = &v
	return s
}

func (s *UpgradeApplicationRequest) SetTimeout(v int) *UpgradeApplicationRequest {
	s.Timeout = &v
	return s
}

type UpgradeApplicationResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpgradeApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeApplicationResponse) GoString() string {
	return s.String()
}

func (s *UpgradeApplicationResponse) SetRequestId(v string) *UpgradeApplicationResponse {
	s.RequestId = &v
	return s
}

type RescaleApplicationRequest struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RescaleType      *string `json:"RescaleType,omitempty" xml:"RescaleType,omitempty" require:"true"`
	RescaleLevel     *string `json:"RescaleLevel,omitempty" xml:"RescaleLevel,omitempty"`
	ResourceSelector *string `json:"ResourceSelector,omitempty" xml:"ResourceSelector,omitempty" require:"true"`
	ToAppVersion     *string `json:"ToAppVersion,omitempty" xml:"ToAppVersion,omitempty"`
	Timeout          *int    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s RescaleApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s RescaleApplicationRequest) GoString() string {
	return s.String()
}

func (s *RescaleApplicationRequest) SetAppId(v string) *RescaleApplicationRequest {
	s.AppId = &v
	return s
}

func (s *RescaleApplicationRequest) SetRescaleType(v string) *RescaleApplicationRequest {
	s.RescaleType = &v
	return s
}

func (s *RescaleApplicationRequest) SetRescaleLevel(v string) *RescaleApplicationRequest {
	s.RescaleLevel = &v
	return s
}

func (s *RescaleApplicationRequest) SetResourceSelector(v string) *RescaleApplicationRequest {
	s.ResourceSelector = &v
	return s
}

func (s *RescaleApplicationRequest) SetToAppVersion(v string) *RescaleApplicationRequest {
	s.ToAppVersion = &v
	return s
}

func (s *RescaleApplicationRequest) SetTimeout(v int) *RescaleApplicationRequest {
	s.Timeout = &v
	return s
}

type RescaleApplicationResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RescaleApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s RescaleApplicationResponse) GoString() string {
	return s.String()
}

func (s *RescaleApplicationResponse) SetRequestId(v string) *RescaleApplicationResponse {
	s.RequestId = &v
	return s
}

type DeleteEpnInstanceRequest struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty" require:"true"`
}

func (s DeleteEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteEpnInstanceRequest) SetEPNInstanceId(v string) *DeleteEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

type DeleteEpnInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteEpnInstanceResponse) SetRequestId(v string) *DeleteEpnInstanceResponse {
	s.RequestId = &v
	return s
}

type ListApplicationsRequest struct {
	ClusterNames     *string `json:"ClusterNames,omitempty" xml:"ClusterNames,omitempty"`
	AppVersions      *string `json:"AppVersions,omitempty" xml:"AppVersions,omitempty"`
	Level            *string `json:"Level,omitempty" xml:"Level,omitempty"`
	OutAppInfoParams *string `json:"OutAppInfoParams,omitempty" xml:"OutAppInfoParams,omitempty"`
	PageNumber       *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	MinDate          *string `json:"MinDate,omitempty" xml:"MinDate,omitempty"`
	MaxDate          *string `json:"MaxDate,omitempty" xml:"MaxDate,omitempty"`
}

func (s ListApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequest) SetClusterNames(v string) *ListApplicationsRequest {
	s.ClusterNames = &v
	return s
}

func (s *ListApplicationsRequest) SetAppVersions(v string) *ListApplicationsRequest {
	s.AppVersions = &v
	return s
}

func (s *ListApplicationsRequest) SetLevel(v string) *ListApplicationsRequest {
	s.Level = &v
	return s
}

func (s *ListApplicationsRequest) SetOutAppInfoParams(v string) *ListApplicationsRequest {
	s.OutAppInfoParams = &v
	return s
}

func (s *ListApplicationsRequest) SetPageNumber(v int) *ListApplicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsRequest) SetPageSize(v int) *ListApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsRequest) SetMinDate(v string) *ListApplicationsRequest {
	s.MinDate = &v
	return s
}

func (s *ListApplicationsRequest) SetMaxDate(v string) *ListApplicationsRequest {
	s.MaxDate = &v
	return s
}

type ListApplicationsResponse struct {
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount   *int                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber   *int                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize     *int                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Applications *ListApplicationsResponseApplications `json:"Applications,omitempty" xml:"Applications,omitempty" require:"true" type:"Struct"`
}

func (s ListApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponse) SetRequestId(v string) *ListApplicationsResponse {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsResponse) SetTotalCount(v int) *ListApplicationsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationsResponse) SetPageNumber(v int) *ListApplicationsResponse {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsResponse) SetPageSize(v int) *ListApplicationsResponse {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsResponse) SetApplications(v *ListApplicationsResponseApplications) *ListApplicationsResponse {
	s.Applications = v
	return s
}

type ListApplicationsResponseApplications struct {
	Application []*ListApplicationsResponseApplicationsApplication `json:"Application,omitempty" xml:"Application,omitempty" require:"true" type:"Repeated"`
}

func (s ListApplicationsResponseApplications) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseApplications) SetApplication(v []*ListApplicationsResponseApplicationsApplication) *ListApplicationsResponseApplications {
	s.Application = v
	return s
}

type ListApplicationsResponseApplicationsApplication struct {
	ClusterName *string                                                 `json:"ClusterName,omitempty" xml:"ClusterName,omitempty" require:"true"`
	AppList     *ListApplicationsResponseApplicationsApplicationAppList `json:"AppList,omitempty" xml:"AppList,omitempty" require:"true" type:"Struct"`
}

func (s ListApplicationsResponseApplicationsApplication) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseApplicationsApplication) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseApplicationsApplication) SetClusterName(v string) *ListApplicationsResponseApplicationsApplication {
	s.ClusterName = &v
	return s
}

func (s *ListApplicationsResponseApplicationsApplication) SetAppList(v *ListApplicationsResponseApplicationsApplicationAppList) *ListApplicationsResponseApplicationsApplication {
	s.AppList = v
	return s
}

type ListApplicationsResponseApplicationsApplicationAppList struct {
	App []*ListApplicationsResponseApplicationsApplicationAppListApp `json:"App,omitempty" xml:"App,omitempty" require:"true" type:"Repeated"`
}

func (s ListApplicationsResponseApplicationsApplicationAppList) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseApplicationsApplicationAppList) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseApplicationsApplicationAppList) SetApp(v []*ListApplicationsResponseApplicationsApplicationAppListApp) *ListApplicationsResponseApplicationsApplicationAppList {
	s.App = v
	return s
}

type ListApplicationsResponseApplicationsApplicationAppListApp struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	AppInfo *string `json:"AppInfo,omitempty" xml:"AppInfo,omitempty" require:"true"`
}

func (s ListApplicationsResponseApplicationsApplicationAppListApp) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseApplicationsApplicationAppListApp) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseApplicationsApplicationAppListApp) SetAppId(v string) *ListApplicationsResponseApplicationsApplicationAppListApp {
	s.AppId = &v
	return s
}

func (s *ListApplicationsResponseApplicationsApplicationAppListApp) SetAppInfo(v string) *ListApplicationsResponseApplicationsApplicationAppListApp {
	s.AppInfo = &v
	return s
}

type DescribeServcieScheduleRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Uuid          *string `json:"Uuid,omitempty" xml:"Uuid,omitempty" require:"true"`
	PodConfigName *string `json:"PodConfigName,omitempty" xml:"PodConfigName,omitempty"`
}

func (s DescribeServcieScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServcieScheduleRequest) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleRequest) SetAppId(v string) *DescribeServcieScheduleRequest {
	s.AppId = &v
	return s
}

func (s *DescribeServcieScheduleRequest) SetUuid(v string) *DescribeServcieScheduleRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeServcieScheduleRequest) SetPodConfigName(v string) *DescribeServcieScheduleRequest {
	s.PodConfigName = &v
	return s
}

type DescribeServcieScheduleResponse struct {
	RequestId       *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	InstanceId      *string                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	InstanceIp      *string                                         `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty" require:"true"`
	InstancePort    *int                                            `json:"InstancePort,omitempty" xml:"InstancePort,omitempty" require:"true"`
	Index           *int                                            `json:"Index,omitempty" xml:"Index,omitempty" require:"true"`
	TcpPorts        *string                                         `json:"TcpPorts,omitempty" xml:"TcpPorts,omitempty" require:"true"`
	RequestRepeated *bool                                           `json:"RequestRepeated,omitempty" xml:"RequestRepeated,omitempty" require:"true"`
	PodAbstractInfo *DescribeServcieScheduleResponsePodAbstractInfo `json:"PodAbstractInfo,omitempty" xml:"PodAbstractInfo,omitempty" require:"true" type:"Struct"`
}

func (s DescribeServcieScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServcieScheduleResponse) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleResponse) SetRequestId(v string) *DescribeServcieScheduleResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeServcieScheduleResponse) SetInstanceId(v string) *DescribeServcieScheduleResponse {
	s.InstanceId = &v
	return s
}

func (s *DescribeServcieScheduleResponse) SetInstanceIp(v string) *DescribeServcieScheduleResponse {
	s.InstanceIp = &v
	return s
}

func (s *DescribeServcieScheduleResponse) SetInstancePort(v int) *DescribeServcieScheduleResponse {
	s.InstancePort = &v
	return s
}

func (s *DescribeServcieScheduleResponse) SetIndex(v int) *DescribeServcieScheduleResponse {
	s.Index = &v
	return s
}

func (s *DescribeServcieScheduleResponse) SetTcpPorts(v string) *DescribeServcieScheduleResponse {
	s.TcpPorts = &v
	return s
}

func (s *DescribeServcieScheduleResponse) SetRequestRepeated(v bool) *DescribeServcieScheduleResponse {
	s.RequestRepeated = &v
	return s
}

func (s *DescribeServcieScheduleResponse) SetPodAbstractInfo(v *DescribeServcieScheduleResponsePodAbstractInfo) *DescribeServcieScheduleResponse {
	s.PodAbstractInfo = v
	return s
}

type DescribeServcieScheduleResponsePodAbstractInfo struct {
	Name              *bool                                                            `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	ResourceScope     *bool                                                            `json:"ResourceScope,omitempty" xml:"ResourceScope,omitempty" require:"true"`
	ContainerService  *bool                                                            `json:"ContainerService,omitempty" xml:"ContainerService,omitempty" require:"true"`
	Namespace         *bool                                                            `json:"Namespace,omitempty" xml:"Namespace,omitempty" require:"true"`
	Status            *bool                                                            `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ContainerStatuses *DescribeServcieScheduleResponsePodAbstractInfoContainerStatuses `json:"ContainerStatuses,omitempty" xml:"ContainerStatuses,omitempty" require:"true" type:"Struct"`
}

func (s DescribeServcieScheduleResponsePodAbstractInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeServcieScheduleResponsePodAbstractInfo) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleResponsePodAbstractInfo) SetName(v bool) *DescribeServcieScheduleResponsePodAbstractInfo {
	s.Name = &v
	return s
}

func (s *DescribeServcieScheduleResponsePodAbstractInfo) SetResourceScope(v bool) *DescribeServcieScheduleResponsePodAbstractInfo {
	s.ResourceScope = &v
	return s
}

func (s *DescribeServcieScheduleResponsePodAbstractInfo) SetContainerService(v bool) *DescribeServcieScheduleResponsePodAbstractInfo {
	s.ContainerService = &v
	return s
}

func (s *DescribeServcieScheduleResponsePodAbstractInfo) SetNamespace(v bool) *DescribeServcieScheduleResponsePodAbstractInfo {
	s.Namespace = &v
	return s
}

func (s *DescribeServcieScheduleResponsePodAbstractInfo) SetStatus(v bool) *DescribeServcieScheduleResponsePodAbstractInfo {
	s.Status = &v
	return s
}

func (s *DescribeServcieScheduleResponsePodAbstractInfo) SetContainerStatuses(v *DescribeServcieScheduleResponsePodAbstractInfoContainerStatuses) *DescribeServcieScheduleResponsePodAbstractInfo {
	s.ContainerStatuses = v
	return s
}

type DescribeServcieScheduleResponsePodAbstractInfoContainerStatuses struct {
	ContainerStatus []*DescribeServcieScheduleResponsePodAbstractInfoContainerStatusesContainerStatus `json:"ContainerStatus,omitempty" xml:"ContainerStatus,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeServcieScheduleResponsePodAbstractInfoContainerStatuses) String() string {
	return tea.Prettify(s)
}

func (s DescribeServcieScheduleResponsePodAbstractInfoContainerStatuses) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleResponsePodAbstractInfoContainerStatuses) SetContainerStatus(v []*DescribeServcieScheduleResponsePodAbstractInfoContainerStatusesContainerStatus) *DescribeServcieScheduleResponsePodAbstractInfoContainerStatuses {
	s.ContainerStatus = v
	return s
}

type DescribeServcieScheduleResponsePodAbstractInfoContainerStatusesContainerStatus struct {
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty" require:"true"`
}

func (s DescribeServcieScheduleResponsePodAbstractInfoContainerStatusesContainerStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeServcieScheduleResponsePodAbstractInfoContainerStatusesContainerStatus) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleResponsePodAbstractInfoContainerStatusesContainerStatus) SetName(v string) *DescribeServcieScheduleResponsePodAbstractInfoContainerStatusesContainerStatus {
	s.Name = &v
	return s
}

func (s *DescribeServcieScheduleResponsePodAbstractInfoContainerStatusesContainerStatus) SetContainerId(v string) *DescribeServcieScheduleResponsePodAbstractInfoContainerStatusesContainerStatus {
	s.ContainerId = &v
	return s
}

type DeleteApplicationRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Timeout *int    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s DeleteApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRequest) SetAppId(v string) *DeleteApplicationRequest {
	s.AppId = &v
	return s
}

func (s *DeleteApplicationRequest) SetTimeout(v int) *DeleteApplicationRequest {
	s.Timeout = &v
	return s
}

type DeleteApplicationResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponse) SetRequestId(v string) *DeleteApplicationResponse {
	s.RequestId = &v
	return s
}

type ModifyEpnInstanceRequest struct {
	EPNInstanceId           *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty" require:"true"`
	EPNInstanceName         *string `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty"`
	NetworkingModel         *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty"`
	InternetMaxBandwidthOut *int    `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
}

func (s ModifyEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyEpnInstanceRequest) SetEPNInstanceId(v string) *ModifyEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *ModifyEpnInstanceRequest) SetEPNInstanceName(v string) *ModifyEpnInstanceRequest {
	s.EPNInstanceName = &v
	return s
}

func (s *ModifyEpnInstanceRequest) SetNetworkingModel(v string) *ModifyEpnInstanceRequest {
	s.NetworkingModel = &v
	return s
}

func (s *ModifyEpnInstanceRequest) SetInternetMaxBandwidthOut(v int) *ModifyEpnInstanceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

type ModifyEpnInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyEpnInstanceResponse) SetRequestId(v string) *ModifyEpnInstanceResponse {
	s.RequestId = &v
	return s
}

type RollbackApplicationRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	FromAppVersion *string `json:"FromAppVersion,omitempty" xml:"FromAppVersion,omitempty" require:"true"`
	ToAppVersion   *string `json:"ToAppVersion,omitempty" xml:"ToAppVersion,omitempty"`
	Timeout        *int    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s RollbackApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackApplicationRequest) GoString() string {
	return s.String()
}

func (s *RollbackApplicationRequest) SetAppId(v string) *RollbackApplicationRequest {
	s.AppId = &v
	return s
}

func (s *RollbackApplicationRequest) SetFromAppVersion(v string) *RollbackApplicationRequest {
	s.FromAppVersion = &v
	return s
}

func (s *RollbackApplicationRequest) SetToAppVersion(v string) *RollbackApplicationRequest {
	s.ToAppVersion = &v
	return s
}

func (s *RollbackApplicationRequest) SetTimeout(v int) *RollbackApplicationRequest {
	s.Timeout = &v
	return s
}

type RollbackApplicationResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RollbackApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackApplicationResponse) GoString() string {
	return s.String()
}

func (s *RollbackApplicationResponse) SetRequestId(v string) *RollbackApplicationResponse {
	s.RequestId = &v
	return s
}

type DescribeEpnInstanceAttributeRequest struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty" require:"true"`
}

func (s DescribeEpnInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeRequest) SetEPNInstanceId(v string) *DescribeEpnInstanceAttributeRequest {
	s.EPNInstanceId = &v
	return s
}

type DescribeEpnInstanceAttributeResponse struct {
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	EPNInstanceId   *string                                             `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty" require:"true"`
	EPNInstanceName *string                                             `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty" require:"true"`
	NetworkingModel *string                                             `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty" require:"true"`
	VSwitches       []*DescribeEpnInstanceAttributeResponseVSwitches    `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" require:"true" type:"Repeated"`
	Instances       []*DescribeEpnInstanceAttributeResponseInstances    `json:"Instances,omitempty" xml:"Instances,omitempty" require:"true" type:"Repeated"`
	ConfVersions    []*DescribeEpnInstanceAttributeResponseConfVersions `json:"ConfVersions,omitempty" xml:"ConfVersions,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeEpnInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeResponse) SetRequestId(v string) *DescribeEpnInstanceAttributeResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponse) SetEPNInstanceId(v string) *DescribeEpnInstanceAttributeResponse {
	s.EPNInstanceId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponse) SetEPNInstanceName(v string) *DescribeEpnInstanceAttributeResponse {
	s.EPNInstanceName = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponse) SetNetworkingModel(v string) *DescribeEpnInstanceAttributeResponse {
	s.NetworkingModel = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponse) SetVSwitches(v []*DescribeEpnInstanceAttributeResponseVSwitches) *DescribeEpnInstanceAttributeResponse {
	s.VSwitches = v
	return s
}

func (s *DescribeEpnInstanceAttributeResponse) SetInstances(v []*DescribeEpnInstanceAttributeResponseInstances) *DescribeEpnInstanceAttributeResponse {
	s.Instances = v
	return s
}

func (s *DescribeEpnInstanceAttributeResponse) SetConfVersions(v []*DescribeEpnInstanceAttributeResponseConfVersions) *DescribeEpnInstanceAttributeResponse {
	s.ConfVersions = v
	return s
}

type DescribeEpnInstanceAttributeResponseVSwitches struct {
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	CidrBlock   *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty" require:"true"`
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty" require:"true"`
}

func (s DescribeEpnInstanceAttributeResponseVSwitches) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstanceAttributeResponseVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeResponseVSwitches) SetVSwitchId(v string) *DescribeEpnInstanceAttributeResponseVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseVSwitches) SetEnsRegionId(v string) *DescribeEpnInstanceAttributeResponseVSwitches {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseVSwitches) SetCidrBlock(v string) *DescribeEpnInstanceAttributeResponseVSwitches {
	s.CidrBlock = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseVSwitches) SetVSwitchName(v string) *DescribeEpnInstanceAttributeResponseVSwitches {
	s.VSwitchName = &v
	return s
}

type DescribeEpnInstanceAttributeResponseInstances struct {
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	PublicIpAddress  *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty" require:"true"`
	EnsRegionId      *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	Isp              *string `json:"Isp,omitempty" xml:"Isp,omitempty" require:"true"`
	InstanceName     *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty" require:"true"`
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" require:"true"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s DescribeEpnInstanceAttributeResponseInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstanceAttributeResponseInstances) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeResponseInstances) SetInstanceId(v string) *DescribeEpnInstanceAttributeResponseInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseInstances) SetPublicIpAddress(v string) *DescribeEpnInstanceAttributeResponseInstances {
	s.PublicIpAddress = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseInstances) SetEnsRegionId(v string) *DescribeEpnInstanceAttributeResponseInstances {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseInstances) SetIsp(v string) *DescribeEpnInstanceAttributeResponseInstances {
	s.Isp = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseInstances) SetInstanceName(v string) *DescribeEpnInstanceAttributeResponseInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseInstances) SetPrivateIpAddress(v string) *DescribeEpnInstanceAttributeResponseInstances {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseInstances) SetStatus(v string) *DescribeEpnInstanceAttributeResponseInstances {
	s.Status = &v
	return s
}

type DescribeEpnInstanceAttributeResponseConfVersions struct {
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	ConfVersion *string `json:"ConfVersion,omitempty" xml:"ConfVersion,omitempty" require:"true"`
}

func (s DescribeEpnInstanceAttributeResponseConfVersions) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstanceAttributeResponseConfVersions) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeResponseConfVersions) SetEnsRegionId(v string) *DescribeEpnInstanceAttributeResponseConfVersions {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeResponseConfVersions) SetConfVersion(v string) *DescribeEpnInstanceAttributeResponseConfVersions {
	s.ConfVersion = &v
	return s
}

type RunServiceScheduleRequest struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Uuid             *string `json:"Uuid,omitempty" xml:"Uuid,omitempty" require:"true"`
	ClientIp         *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty" require:"true"`
	ServiceAction    *string `json:"ServiceAction,omitempty" xml:"ServiceAction,omitempty" require:"true"`
	PodConfigName    *string `json:"PodConfigName,omitempty" xml:"PodConfigName,omitempty"`
	PreLockedTimeout *int    `json:"PreLockedTimeout,omitempty" xml:"PreLockedTimeout,omitempty"`
	Directorys       *string `json:"Directorys,omitempty" xml:"Directorys,omitempty"`
	ServiceCommands  *string `json:"ServiceCommands,omitempty" xml:"ServiceCommands,omitempty"`
	ScheduleStrategy *string `json:"ScheduleStrategy,omitempty" xml:"ScheduleStrategy,omitempty"`
}

func (s RunServiceScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s RunServiceScheduleRequest) GoString() string {
	return s.String()
}

func (s *RunServiceScheduleRequest) SetAppId(v string) *RunServiceScheduleRequest {
	s.AppId = &v
	return s
}

func (s *RunServiceScheduleRequest) SetUuid(v string) *RunServiceScheduleRequest {
	s.Uuid = &v
	return s
}

func (s *RunServiceScheduleRequest) SetClientIp(v string) *RunServiceScheduleRequest {
	s.ClientIp = &v
	return s
}

func (s *RunServiceScheduleRequest) SetServiceAction(v string) *RunServiceScheduleRequest {
	s.ServiceAction = &v
	return s
}

func (s *RunServiceScheduleRequest) SetPodConfigName(v string) *RunServiceScheduleRequest {
	s.PodConfigName = &v
	return s
}

func (s *RunServiceScheduleRequest) SetPreLockedTimeout(v int) *RunServiceScheduleRequest {
	s.PreLockedTimeout = &v
	return s
}

func (s *RunServiceScheduleRequest) SetDirectorys(v string) *RunServiceScheduleRequest {
	s.Directorys = &v
	return s
}

func (s *RunServiceScheduleRequest) SetServiceCommands(v string) *RunServiceScheduleRequest {
	s.ServiceCommands = &v
	return s
}

func (s *RunServiceScheduleRequest) SetScheduleStrategy(v string) *RunServiceScheduleRequest {
	s.ScheduleStrategy = &v
	return s
}

type RunServiceScheduleResponse struct {
	RequestId       *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RequestRepeated *string                                   `json:"RequestRepeated,omitempty" xml:"RequestRepeated,omitempty" require:"true"`
	TcpPorts        *bool                                     `json:"TcpPorts,omitempty" xml:"TcpPorts,omitempty" require:"true"`
	InstanceId      *string                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	InstancePort    *int                                      `json:"InstancePort,omitempty" xml:"InstancePort,omitempty" require:"true"`
	InstanceIp      *string                                   `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty" require:"true"`
	Index           *int                                      `json:"Index,omitempty" xml:"Index,omitempty" require:"true"`
	CommandResults  *RunServiceScheduleResponseCommandResults `json:"CommandResults,omitempty" xml:"CommandResults,omitempty" require:"true" type:"Struct"`
}

func (s RunServiceScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s RunServiceScheduleResponse) GoString() string {
	return s.String()
}

func (s *RunServiceScheduleResponse) SetRequestId(v string) *RunServiceScheduleResponse {
	s.RequestId = &v
	return s
}

func (s *RunServiceScheduleResponse) SetRequestRepeated(v string) *RunServiceScheduleResponse {
	s.RequestRepeated = &v
	return s
}

func (s *RunServiceScheduleResponse) SetTcpPorts(v bool) *RunServiceScheduleResponse {
	s.TcpPorts = &v
	return s
}

func (s *RunServiceScheduleResponse) SetInstanceId(v string) *RunServiceScheduleResponse {
	s.InstanceId = &v
	return s
}

func (s *RunServiceScheduleResponse) SetInstancePort(v int) *RunServiceScheduleResponse {
	s.InstancePort = &v
	return s
}

func (s *RunServiceScheduleResponse) SetInstanceIp(v string) *RunServiceScheduleResponse {
	s.InstanceIp = &v
	return s
}

func (s *RunServiceScheduleResponse) SetIndex(v int) *RunServiceScheduleResponse {
	s.Index = &v
	return s
}

func (s *RunServiceScheduleResponse) SetCommandResults(v *RunServiceScheduleResponseCommandResults) *RunServiceScheduleResponse {
	s.CommandResults = v
	return s
}

type RunServiceScheduleResponseCommandResults struct {
	CommandResult []*RunServiceScheduleResponseCommandResultsCommandResult `json:"CommandResult,omitempty" xml:"CommandResult,omitempty" require:"true" type:"Repeated"`
}

func (s RunServiceScheduleResponseCommandResults) String() string {
	return tea.Prettify(s)
}

func (s RunServiceScheduleResponseCommandResults) GoString() string {
	return s.String()
}

func (s *RunServiceScheduleResponseCommandResults) SetCommandResult(v []*RunServiceScheduleResponseCommandResultsCommandResult) *RunServiceScheduleResponseCommandResults {
	s.CommandResult = v
	return s
}

type RunServiceScheduleResponseCommandResultsCommandResult struct {
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty" require:"true"`
	Command       *string `json:"Command,omitempty" xml:"Command,omitempty" require:"true"`
	ResultMsg     *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty" require:"true"`
}

func (s RunServiceScheduleResponseCommandResultsCommandResult) String() string {
	return tea.Prettify(s)
}

func (s RunServiceScheduleResponseCommandResultsCommandResult) GoString() string {
	return s.String()
}

func (s *RunServiceScheduleResponseCommandResultsCommandResult) SetContainerName(v string) *RunServiceScheduleResponseCommandResultsCommandResult {
	s.ContainerName = &v
	return s
}

func (s *RunServiceScheduleResponseCommandResultsCommandResult) SetCommand(v string) *RunServiceScheduleResponseCommandResultsCommandResult {
	s.Command = &v
	return s
}

func (s *RunServiceScheduleResponseCommandResultsCommandResult) SetResultMsg(v string) *RunServiceScheduleResponseCommandResultsCommandResult {
	s.ResultMsg = &v
	return s
}

type CreateApplicationRequest struct {
	Template *string `json:"Template,omitempty" xml:"Template,omitempty" require:"true"`
	Timeout  *int    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) SetTemplate(v string) *CreateApplicationRequest {
	s.Template = &v
	return s
}

func (s *CreateApplicationRequest) SetTimeout(v int) *CreateApplicationRequest {
	s.Timeout = &v
	return s
}

type CreateApplicationResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s CreateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponse) SetRequestId(v string) *CreateApplicationResponse {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationResponse) SetAppId(v string) *CreateApplicationResponse {
	s.AppId = &v
	return s
}

type CreateEpnInstanceRequest struct {
	EPNInstanceType         *string `json:"EPNInstanceType,omitempty" xml:"EPNInstanceType,omitempty" require:"true"`
	EPNInstanceName         *string `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty"`
	InternetChargeType      *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty" require:"true"`
	NetworkingModel         *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty" require:"true"`
	InternetMaxBandwidthOut *int    `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty" require:"true"`
}

func (s CreateEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateEpnInstanceRequest) SetEPNInstanceType(v string) *CreateEpnInstanceRequest {
	s.EPNInstanceType = &v
	return s
}

func (s *CreateEpnInstanceRequest) SetEPNInstanceName(v string) *CreateEpnInstanceRequest {
	s.EPNInstanceName = &v
	return s
}

func (s *CreateEpnInstanceRequest) SetInternetChargeType(v string) *CreateEpnInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateEpnInstanceRequest) SetNetworkingModel(v string) *CreateEpnInstanceRequest {
	s.NetworkingModel = &v
	return s
}

func (s *CreateEpnInstanceRequest) SetInternetMaxBandwidthOut(v int) *CreateEpnInstanceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

type CreateEpnInstanceResponse struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty" require:"true"`
}

func (s CreateEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateEpnInstanceResponse) SetRequestId(v string) *CreateEpnInstanceResponse {
	s.RequestId = &v
	return s
}

func (s *CreateEpnInstanceResponse) SetEPNInstanceId(v string) *CreateEpnInstanceResponse {
	s.EPNInstanceId = &v
	return s
}

type StopEpnInstanceRequest struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty" require:"true"`
}

func (s StopEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopEpnInstanceRequest) SetEPNInstanceId(v string) *StopEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

type StopEpnInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StopEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopEpnInstanceResponse) SetRequestId(v string) *StopEpnInstanceResponse {
	s.RequestId = &v
	return s
}

type DescribeDataDistResultRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	DataNames    *string `json:"DataNames,omitempty" xml:"DataNames,omitempty"`
	DataVersions *string `json:"DataVersions,omitempty" xml:"DataVersions,omitempty"`
	InstanceIds  *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	MinDate      *string `json:"MinDate,omitempty" xml:"MinDate,omitempty"`
	MaxDate      *string `json:"MaxDate,omitempty" xml:"MaxDate,omitempty"`
	PageNumber   *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDataDistResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultRequest) SetAppId(v string) *DescribeDataDistResultRequest {
	s.AppId = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetDataNames(v string) *DescribeDataDistResultRequest {
	s.DataNames = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetDataVersions(v string) *DescribeDataDistResultRequest {
	s.DataVersions = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetInstanceIds(v string) *DescribeDataDistResultRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetMinDate(v string) *DescribeDataDistResultRequest {
	s.MinDate = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetMaxDate(v string) *DescribeDataDistResultRequest {
	s.MaxDate = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetPageNumber(v int) *DescribeDataDistResultRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetPageSize(v int) *DescribeDataDistResultRequest {
	s.PageSize = &v
	return s
}

type DescribeDataDistResultResponse struct {
	RequestId   *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount  *int                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber  *int                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize    *int                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	DistResults *DescribeDataDistResultResponseDistResults `json:"DistResults,omitempty" xml:"DistResults,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDataDistResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponse) SetRequestId(v string) *DescribeDataDistResultResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDataDistResultResponse) SetTotalCount(v int) *DescribeDataDistResultResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataDistResultResponse) SetPageNumber(v int) *DescribeDataDistResultResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataDistResultResponse) SetPageSize(v int) *DescribeDataDistResultResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeDataDistResultResponse) SetDistResults(v *DescribeDataDistResultResponseDistResults) *DescribeDataDistResultResponse {
	s.DistResults = v
	return s
}

type DescribeDataDistResultResponseDistResults struct {
	DistResult []*DescribeDataDistResultResponseDistResultsDistResult `json:"DistResult,omitempty" xml:"DistResult,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDataDistResultResponseDistResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultResponseDistResults) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseDistResults) SetDistResult(v []*DescribeDataDistResultResponseDistResultsDistResult) *DescribeDataDistResultResponseDistResults {
	s.DistResult = v
	return s
}

type DescribeDataDistResultResponseDistResultsDistResult struct {
	Version     *string                                                         `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Name        *string                                                         `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	StatusStats *DescribeDataDistResultResponseDistResultsDistResultStatusStats `json:"StatusStats,omitempty" xml:"StatusStats,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDataDistResultResponseDistResultsDistResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultResponseDistResultsDistResult) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseDistResultsDistResult) SetVersion(v string) *DescribeDataDistResultResponseDistResultsDistResult {
	s.Version = &v
	return s
}

func (s *DescribeDataDistResultResponseDistResultsDistResult) SetName(v string) *DescribeDataDistResultResponseDistResultsDistResult {
	s.Name = &v
	return s
}

func (s *DescribeDataDistResultResponseDistResultsDistResult) SetStatusStats(v *DescribeDataDistResultResponseDistResultsDistResultStatusStats) *DescribeDataDistResultResponseDistResultsDistResult {
	s.StatusStats = v
	return s
}

type DescribeDataDistResultResponseDistResultsDistResultStatusStats struct {
	StatusStat []*DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStat `json:"StatusStat,omitempty" xml:"StatusStat,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDataDistResultResponseDistResultsDistResultStatusStats) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultResponseDistResultsDistResultStatusStats) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseDistResultsDistResultStatusStats) SetStatusStat(v []*DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStat) *DescribeDataDistResultResponseDistResultsDistResultStatusStats {
	s.StatusStat = v
	return s
}

type DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStat struct {
	Status        *string                                                                            `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	InstanceCount *string                                                                            `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty" require:"true"`
	Instances     *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstances `json:"Instances,omitempty" xml:"Instances,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStat) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStat) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStat) SetStatus(v string) *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStat {
	s.Status = &v
	return s
}

func (s *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStat) SetInstanceCount(v string) *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStat {
	s.InstanceCount = &v
	return s
}

func (s *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStat) SetInstances(v *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstances) *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStat {
	s.Instances = v
	return s
}

type DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstances struct {
	Instance []*DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstances) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstances) SetInstance(v []*DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstancesInstance) *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstances {
	s.Instance = v
	return s
}

type DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstancesInstance struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	StatusDescrip *string `json:"StatusDescrip,omitempty" xml:"StatusDescrip,omitempty" require:"true"`
}

func (s DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstancesInstance) SetInstanceId(v string) *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstancesInstance) SetStartTime(v string) *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstancesInstance {
	s.StartTime = &v
	return s
}

func (s *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstancesInstance) SetUpdateTime(v string) *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstancesInstance {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstancesInstance) SetStatusDescrip(v string) *DescribeDataDistResultResponseDistResultsDistResultStatusStatsStatusStatInstancesInstance {
	s.StatusDescrip = &v
	return s
}

type DescribeEpnInstancesRequest struct {
	EPNInstanceId   *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	EPNInstanceName *string `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty"`
	PageNumber      *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeEpnInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstancesRequest) SetEPNInstanceId(v string) *DescribeEpnInstancesRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *DescribeEpnInstancesRequest) SetEPNInstanceName(v string) *DescribeEpnInstancesRequest {
	s.EPNInstanceName = &v
	return s
}

func (s *DescribeEpnInstancesRequest) SetPageNumber(v int) *DescribeEpnInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEpnInstancesRequest) SetPageSize(v int) *DescribeEpnInstancesRequest {
	s.PageSize = &v
	return s
}

type DescribeEpnInstancesResponse struct {
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount   *int                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber   *int                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize     *int                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	EPNInstances *DescribeEpnInstancesResponseEPNInstances `json:"EPNInstances,omitempty" xml:"EPNInstances,omitempty" require:"true" type:"Struct"`
}

func (s DescribeEpnInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstancesResponse) SetRequestId(v string) *DescribeEpnInstancesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeEpnInstancesResponse) SetTotalCount(v int) *DescribeEpnInstancesResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeEpnInstancesResponse) SetPageNumber(v int) *DescribeEpnInstancesResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeEpnInstancesResponse) SetPageSize(v int) *DescribeEpnInstancesResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeEpnInstancesResponse) SetEPNInstances(v *DescribeEpnInstancesResponseEPNInstances) *DescribeEpnInstancesResponse {
	s.EPNInstances = v
	return s
}

type DescribeEpnInstancesResponseEPNInstances struct {
	EPNInstance []*DescribeEpnInstancesResponseEPNInstancesEPNInstance `json:"EPNInstance,omitempty" xml:"EPNInstance,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeEpnInstancesResponseEPNInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstancesResponseEPNInstances) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstancesResponseEPNInstances) SetEPNInstance(v []*DescribeEpnInstancesResponseEPNInstancesEPNInstance) *DescribeEpnInstancesResponseEPNInstances {
	s.EPNInstance = v
	return s
}

type DescribeEpnInstancesResponseEPNInstancesEPNInstance struct {
	EPNInstanceId           *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty" require:"true"`
	EPNInstanceName         *string `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty" require:"true"`
	NetworkingModel         *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty" require:"true"`
	ModifyTime              *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	EPNInstanceType         *string `json:"EPNInstanceType,omitempty" xml:"EPNInstanceType,omitempty" require:"true"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	InternetMaxBandwidthOut *int    `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty" require:"true"`
	CreationTime            *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	StartTime               *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime                 *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s DescribeEpnInstancesResponseEPNInstancesEPNInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeEpnInstancesResponseEPNInstancesEPNInstance) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstancesResponseEPNInstancesEPNInstance) SetEPNInstanceId(v string) *DescribeEpnInstancesResponseEPNInstancesEPNInstance {
	s.EPNInstanceId = &v
	return s
}

func (s *DescribeEpnInstancesResponseEPNInstancesEPNInstance) SetEPNInstanceName(v string) *DescribeEpnInstancesResponseEPNInstancesEPNInstance {
	s.EPNInstanceName = &v
	return s
}

func (s *DescribeEpnInstancesResponseEPNInstancesEPNInstance) SetNetworkingModel(v string) *DescribeEpnInstancesResponseEPNInstancesEPNInstance {
	s.NetworkingModel = &v
	return s
}

func (s *DescribeEpnInstancesResponseEPNInstancesEPNInstance) SetModifyTime(v string) *DescribeEpnInstancesResponseEPNInstancesEPNInstance {
	s.ModifyTime = &v
	return s
}

func (s *DescribeEpnInstancesResponseEPNInstancesEPNInstance) SetEPNInstanceType(v string) *DescribeEpnInstancesResponseEPNInstancesEPNInstance {
	s.EPNInstanceType = &v
	return s
}

func (s *DescribeEpnInstancesResponseEPNInstancesEPNInstance) SetStatus(v string) *DescribeEpnInstancesResponseEPNInstancesEPNInstance {
	s.Status = &v
	return s
}

func (s *DescribeEpnInstancesResponseEPNInstancesEPNInstance) SetInternetMaxBandwidthOut(v int) *DescribeEpnInstancesResponseEPNInstancesEPNInstance {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeEpnInstancesResponseEPNInstancesEPNInstance) SetCreationTime(v string) *DescribeEpnInstancesResponseEPNInstancesEPNInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeEpnInstancesResponseEPNInstancesEPNInstance) SetStartTime(v string) *DescribeEpnInstancesResponseEPNInstancesEPNInstance {
	s.StartTime = &v
	return s
}

func (s *DescribeEpnInstancesResponseEPNInstancesEPNInstance) SetEndTime(v string) *DescribeEpnInstancesResponseEPNInstancesEPNInstance {
	s.EndTime = &v
	return s
}

type RemovePublicIpsFromEpnInstanceRequest struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty" require:"true"`
	InstanceInfos *string `json:"InstanceInfos,omitempty" xml:"InstanceInfos,omitempty" require:"true"`
}

func (s RemovePublicIpsFromEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RemovePublicIpsFromEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *RemovePublicIpsFromEpnInstanceRequest) SetEPNInstanceId(v string) *RemovePublicIpsFromEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *RemovePublicIpsFromEpnInstanceRequest) SetInstanceInfos(v string) *RemovePublicIpsFromEpnInstanceRequest {
	s.InstanceInfos = &v
	return s
}

type RemovePublicIpsFromEpnInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RemovePublicIpsFromEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RemovePublicIpsFromEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *RemovePublicIpsFromEpnInstanceResponse) SetRequestId(v string) *RemovePublicIpsFromEpnInstanceResponse {
	s.RequestId = &v
	return s
}

type JoinPublicIpsToEpnInstanceRequest struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty" require:"true"`
	InstanceInfos *string `json:"InstanceInfos,omitempty" xml:"InstanceInfos,omitempty" require:"true"`
}

func (s JoinPublicIpsToEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinPublicIpsToEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *JoinPublicIpsToEpnInstanceRequest) SetEPNInstanceId(v string) *JoinPublicIpsToEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *JoinPublicIpsToEpnInstanceRequest) SetInstanceInfos(v string) *JoinPublicIpsToEpnInstanceRequest {
	s.InstanceInfos = &v
	return s
}

type JoinPublicIpsToEpnInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s JoinPublicIpsToEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinPublicIpsToEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *JoinPublicIpsToEpnInstanceResponse) SetRequestId(v string) *JoinPublicIpsToEpnInstanceResponse {
	s.RequestId = &v
	return s
}

type DescribeApplicationResourceSummaryRequest struct {
	Level        *string `json:"Level,omitempty" xml:"Level,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeApplicationResourceSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationResourceSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationResourceSummaryRequest) SetLevel(v string) *DescribeApplicationResourceSummaryRequest {
	s.Level = &v
	return s
}

func (s *DescribeApplicationResourceSummaryRequest) SetResourceType(v string) *DescribeApplicationResourceSummaryRequest {
	s.ResourceType = &v
	return s
}

type DescribeApplicationResourceSummaryResponse struct {
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ApplicationResource *string `json:"ApplicationResource,omitempty" xml:"ApplicationResource,omitempty" require:"true"`
}

func (s DescribeApplicationResourceSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationResourceSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationResourceSummaryResponse) SetRequestId(v string) *DescribeApplicationResourceSummaryResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationResourceSummaryResponse) SetApplicationResource(v string) *DescribeApplicationResourceSummaryResponse {
	s.ApplicationResource = &v
	return s
}

type StartEpnInstanceRequest struct {
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty" require:"true"`
}

func (s StartEpnInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartEpnInstanceRequest) SetEPNInstanceId(v string) *StartEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

type StartEpnInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StartEpnInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartEpnInstanceResponse) SetRequestId(v string) *StartEpnInstanceResponse {
	s.RequestId = &v
	return s
}

type DescribeExportImageInfoRequest struct {
	ImageId    *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName  *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeExportImageInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportImageInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeExportImageInfoRequest) SetImageId(v string) *DescribeExportImageInfoRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeExportImageInfoRequest) SetImageName(v string) *DescribeExportImageInfoRequest {
	s.ImageName = &v
	return s
}

func (s *DescribeExportImageInfoRequest) SetPageNumber(v int) *DescribeExportImageInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeExportImageInfoRequest) SetPageSize(v int) *DescribeExportImageInfoRequest {
	s.PageSize = &v
	return s
}

type DescribeExportImageInfoResponse struct {
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount *int                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber *int                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Images     *DescribeExportImageInfoResponseImages `json:"Images,omitempty" xml:"Images,omitempty" require:"true" type:"Struct"`
}

func (s DescribeExportImageInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportImageInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeExportImageInfoResponse) SetRequestId(v string) *DescribeExportImageInfoResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeExportImageInfoResponse) SetTotalCount(v int) *DescribeExportImageInfoResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeExportImageInfoResponse) SetPageNumber(v int) *DescribeExportImageInfoResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeExportImageInfoResponse) SetPageSize(v int) *DescribeExportImageInfoResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeExportImageInfoResponse) SetImages(v *DescribeExportImageInfoResponseImages) *DescribeExportImageInfoResponse {
	s.Images = v
	return s
}

type DescribeExportImageInfoResponseImages struct {
	Image []*DescribeExportImageInfoResponseImagesImage `json:"Image,omitempty" xml:"Image,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeExportImageInfoResponseImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportImageInfoResponseImages) GoString() string {
	return s.String()
}

func (s *DescribeExportImageInfoResponseImages) SetImage(v []*DescribeExportImageInfoResponseImagesImage) *DescribeExportImageInfoResponseImages {
	s.Image = v
	return s
}

type DescribeExportImageInfoResponseImagesImage struct {
	CreationTime      *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	ImageId           *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	ImageName         *string `json:"ImageName,omitempty" xml:"ImageName,omitempty" require:"true"`
	Architecture      *string `json:"Architecture,omitempty" xml:"Architecture,omitempty" require:"true"`
	ImageOwnerAlias   *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty" require:"true"`
	Platform          *string `json:"Platform,omitempty" xml:"Platform,omitempty" require:"true"`
	ImageExportStatus *string `json:"ImageExportStatus,omitempty" xml:"ImageExportStatus,omitempty" require:"true"`
	ExportedImageURL  *string `json:"ExportedImageURL,omitempty" xml:"ExportedImageURL,omitempty" require:"true"`
}

func (s DescribeExportImageInfoResponseImagesImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportImageInfoResponseImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeExportImageInfoResponseImagesImage) SetCreationTime(v string) *DescribeExportImageInfoResponseImagesImage {
	s.CreationTime = &v
	return s
}

func (s *DescribeExportImageInfoResponseImagesImage) SetImageId(v string) *DescribeExportImageInfoResponseImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeExportImageInfoResponseImagesImage) SetImageName(v string) *DescribeExportImageInfoResponseImagesImage {
	s.ImageName = &v
	return s
}

func (s *DescribeExportImageInfoResponseImagesImage) SetArchitecture(v string) *DescribeExportImageInfoResponseImagesImage {
	s.Architecture = &v
	return s
}

func (s *DescribeExportImageInfoResponseImagesImage) SetImageOwnerAlias(v string) *DescribeExportImageInfoResponseImagesImage {
	s.ImageOwnerAlias = &v
	return s
}

func (s *DescribeExportImageInfoResponseImagesImage) SetPlatform(v string) *DescribeExportImageInfoResponseImagesImage {
	s.Platform = &v
	return s
}

func (s *DescribeExportImageInfoResponseImagesImage) SetImageExportStatus(v string) *DescribeExportImageInfoResponseImagesImage {
	s.ImageExportStatus = &v
	return s
}

func (s *DescribeExportImageInfoResponseImagesImage) SetExportedImageURL(v string) *DescribeExportImageInfoResponseImagesImage {
	s.ExportedImageURL = &v
	return s
}

type DescribeVSwitchesRequest struct {
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	EnsRegionId   *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	VSwitchId     *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchName   *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	PageNumber    *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	OrderByParams *string `json:"OrderByParams,omitempty" xml:"OrderByParams,omitempty"`
}

func (s DescribeVSwitchesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesRequest) SetVersion(v string) *DescribeVSwitchesRequest {
	s.Version = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetEnsRegionId(v string) *DescribeVSwitchesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVSwitchId(v string) *DescribeVSwitchesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVSwitchName(v string) *DescribeVSwitchesRequest {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetPageNumber(v int) *DescribeVSwitchesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetPageSize(v int) *DescribeVSwitchesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetOrderByParams(v string) *DescribeVSwitchesRequest {
	s.OrderByParams = &v
	return s
}

type DescribeVSwitchesResponse struct {
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount *int                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber *int                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	VSwitches  *DescribeVSwitchesResponseVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" require:"true" type:"Struct"`
}

func (s DescribeVSwitchesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponse) SetRequestId(v string) *DescribeVSwitchesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVSwitchesResponse) SetTotalCount(v int) *DescribeVSwitchesResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeVSwitchesResponse) SetPageNumber(v int) *DescribeVSwitchesResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchesResponse) SetPageSize(v int) *DescribeVSwitchesResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchesResponse) SetVSwitches(v *DescribeVSwitchesResponseVSwitches) *DescribeVSwitchesResponse {
	s.VSwitches = v
	return s
}

type DescribeVSwitchesResponseVSwitches struct {
	VSwitch []*DescribeVSwitchesResponseVSwitchesVSwitch `json:"VSwitch,omitempty" xml:"VSwitch,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeVSwitchesResponseVSwitches) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponseVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseVSwitches) SetVSwitch(v []*DescribeVSwitchesResponseVSwitchesVSwitch) *DescribeVSwitchesResponseVSwitches {
	s.VSwitch = v
	return s
}

type DescribeVSwitchesResponseVSwitchesVSwitch struct {
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true"`
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty" require:"true"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	CidrBlock   *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty" require:"true"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	FreeIpCount *int64  `json:"FreeIpCount,omitempty" xml:"FreeIpCount,omitempty" require:"true"`
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty" require:"true"`
}

func (s DescribeVSwitchesResponseVSwitchesVSwitch) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponseVSwitchesVSwitch) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseVSwitchesVSwitch) SetVSwitchId(v string) *DescribeVSwitchesResponseVSwitchesVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesResponseVSwitchesVSwitch) SetVSwitchName(v string) *DescribeVSwitchesResponseVSwitchesVSwitch {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchesResponseVSwitchesVSwitch) SetStatus(v string) *DescribeVSwitchesResponseVSwitchesVSwitch {
	s.Status = &v
	return s
}

func (s *DescribeVSwitchesResponseVSwitchesVSwitch) SetCidrBlock(v string) *DescribeVSwitchesResponseVSwitchesVSwitch {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVSwitchesResponseVSwitchesVSwitch) SetEnsRegionId(v string) *DescribeVSwitchesResponseVSwitchesVSwitch {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeVSwitchesResponseVSwitchesVSwitch) SetFreeIpCount(v int64) *DescribeVSwitchesResponseVSwitchesVSwitch {
	s.FreeIpCount = &v
	return s
}

func (s *DescribeVSwitchesResponseVSwitchesVSwitch) SetCreatedTime(v string) *DescribeVSwitchesResponseVSwitchesVSwitch {
	s.CreatedTime = &v
	return s
}

type DeleteVSwitchRequest struct {
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true"`
}

func (s DeleteVSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVSwitchRequest) GoString() string {
	return s.String()
}

func (s *DeleteVSwitchRequest) SetVersion(v string) *DeleteVSwitchRequest {
	s.Version = &v
	return s
}

func (s *DeleteVSwitchRequest) SetVSwitchId(v string) *DeleteVSwitchRequest {
	s.VSwitchId = &v
	return s
}

type DeleteVSwitchResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteVSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVSwitchResponse) GoString() string {
	return s.String()
}

func (s *DeleteVSwitchResponse) SetRequestId(v string) *DeleteVSwitchResponse {
	s.RequestId = &v
	return s
}

type CreateVSwitchRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	CidrBlock   *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty" require:"true"`
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s CreateVSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVSwitchRequest) GoString() string {
	return s.String()
}

func (s *CreateVSwitchRequest) SetVersion(v string) *CreateVSwitchRequest {
	s.Version = &v
	return s
}

func (s *CreateVSwitchRequest) SetEnsRegionId(v string) *CreateVSwitchRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateVSwitchRequest) SetCidrBlock(v string) *CreateVSwitchRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateVSwitchRequest) SetVSwitchName(v string) *CreateVSwitchRequest {
	s.VSwitchName = &v
	return s
}

type CreateVSwitchResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true"`
}

func (s CreateVSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVSwitchResponse) GoString() string {
	return s.String()
}

func (s *CreateVSwitchResponse) SetRequestId(v string) *CreateVSwitchResponse {
	s.RequestId = &v
	return s
}

func (s *CreateVSwitchResponse) SetVSwitchId(v string) *CreateVSwitchResponse {
	s.VSwitchId = &v
	return s
}

type DescribeExportImageStatusRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
}

func (s DescribeExportImageStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportImageStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeExportImageStatusRequest) SetVersion(v string) *DescribeExportImageStatusRequest {
	s.Version = &v
	return s
}

func (s *DescribeExportImageStatusRequest) SetImageId(v string) *DescribeExportImageStatusRequest {
	s.ImageId = &v
	return s
}

type DescribeExportImageStatusResponse struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ImageExportStatus *string `json:"ImageExportStatus,omitempty" xml:"ImageExportStatus,omitempty" require:"true"`
}

func (s DescribeExportImageStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportImageStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeExportImageStatusResponse) SetRequestId(v string) *DescribeExportImageStatusResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeExportImageStatusResponse) SetImageExportStatus(v string) *DescribeExportImageStatusResponse {
	s.ImageExportStatus = &v
	return s
}

type ExportImageRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	OSSBucket   *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty" require:"true"`
	OSSRegionId *string `json:"OSSRegionId,omitempty" xml:"OSSRegionId,omitempty" require:"true"`
	OSSPrefix   *string `json:"OSSPrefix,omitempty" xml:"OSSPrefix,omitempty"`
	RoleName    *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ExportImageRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportImageRequest) GoString() string {
	return s.String()
}

func (s *ExportImageRequest) SetVersion(v string) *ExportImageRequest {
	s.Version = &v
	return s
}

func (s *ExportImageRequest) SetImageId(v string) *ExportImageRequest {
	s.ImageId = &v
	return s
}

func (s *ExportImageRequest) SetOSSBucket(v string) *ExportImageRequest {
	s.OSSBucket = &v
	return s
}

func (s *ExportImageRequest) SetOSSRegionId(v string) *ExportImageRequest {
	s.OSSRegionId = &v
	return s
}

func (s *ExportImageRequest) SetOSSPrefix(v string) *ExportImageRequest {
	s.OSSPrefix = &v
	return s
}

func (s *ExportImageRequest) SetRoleName(v string) *ExportImageRequest {
	s.RoleName = &v
	return s
}

type ExportImageResponse struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ExportedImageURL *string `json:"ExportedImageURL,omitempty" xml:"ExportedImageURL,omitempty" require:"true"`
}

func (s ExportImageResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportImageResponse) GoString() string {
	return s.String()
}

func (s *ExportImageResponse) SetRequestId(v string) *ExportImageResponse {
	s.RequestId = &v
	return s
}

func (s *ExportImageResponse) SetExportedImageURL(v string) *ExportImageResponse {
	s.ExportedImageURL = &v
	return s
}

type ImportKeyPairRequest struct {
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	KeyPairName   *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty" require:"true"`
	PublicKeyBody *string `json:"PublicKeyBody,omitempty" xml:"PublicKeyBody,omitempty" require:"true"`
}

func (s ImportKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyPairRequest) GoString() string {
	return s.String()
}

func (s *ImportKeyPairRequest) SetVersion(v string) *ImportKeyPairRequest {
	s.Version = &v
	return s
}

func (s *ImportKeyPairRequest) SetKeyPairName(v string) *ImportKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *ImportKeyPairRequest) SetPublicKeyBody(v string) *ImportKeyPairRequest {
	s.PublicKeyBody = &v
	return s
}

type ImportKeyPairResponse struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	KeyPairName        *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty" require:"true"`
	KeyPairFingerPrint *string `json:"KeyPairFingerPrint,omitempty" xml:"KeyPairFingerPrint,omitempty" require:"true"`
}

func (s ImportKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyPairResponse) GoString() string {
	return s.String()
}

func (s *ImportKeyPairResponse) SetRequestId(v string) *ImportKeyPairResponse {
	s.RequestId = &v
	return s
}

func (s *ImportKeyPairResponse) SetKeyPairName(v string) *ImportKeyPairResponse {
	s.KeyPairName = &v
	return s
}

func (s *ImportKeyPairResponse) SetKeyPairFingerPrint(v string) *ImportKeyPairResponse {
	s.KeyPairFingerPrint = &v
	return s
}

type DescribeKeyPairsRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	PageNumber  *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeKeyPairsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsRequest) SetVersion(v string) *DescribeKeyPairsRequest {
	s.Version = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetKeyPairName(v string) *DescribeKeyPairsRequest {
	s.KeyPairName = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetPageNumber(v string) *DescribeKeyPairsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetPageSize(v string) *DescribeKeyPairsRequest {
	s.PageSize = &v
	return s
}

type DescribeKeyPairsResponse struct {
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount *int                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber *int                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                              `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	KeyPairs   *DescribeKeyPairsResponseKeyPairs `json:"KeyPairs,omitempty" xml:"KeyPairs,omitempty" require:"true" type:"Struct"`
}

func (s DescribeKeyPairsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyPairsResponse) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponse) SetRequestId(v string) *DescribeKeyPairsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeKeyPairsResponse) SetTotalCount(v int) *DescribeKeyPairsResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeKeyPairsResponse) SetPageNumber(v int) *DescribeKeyPairsResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeKeyPairsResponse) SetPageSize(v int) *DescribeKeyPairsResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeKeyPairsResponse) SetKeyPairs(v *DescribeKeyPairsResponseKeyPairs) *DescribeKeyPairsResponse {
	s.KeyPairs = v
	return s
}

type DescribeKeyPairsResponseKeyPairs struct {
	KeyPair []*DescribeKeyPairsResponseKeyPairsKeyPair `json:"KeyPair,omitempty" xml:"KeyPair,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeKeyPairsResponseKeyPairs) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyPairsResponseKeyPairs) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseKeyPairs) SetKeyPair(v []*DescribeKeyPairsResponseKeyPairsKeyPair) *DescribeKeyPairsResponseKeyPairs {
	s.KeyPair = v
	return s
}

type DescribeKeyPairsResponseKeyPairsKeyPair struct {
	CreationTime       *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	KeyPairName        *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty" require:"true"`
	KeyPairFingerPrint *string `json:"KeyPairFingerPrint,omitempty" xml:"KeyPairFingerPrint,omitempty" require:"true"`
}

func (s DescribeKeyPairsResponseKeyPairsKeyPair) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyPairsResponseKeyPairsKeyPair) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseKeyPairsKeyPair) SetCreationTime(v string) *DescribeKeyPairsResponseKeyPairsKeyPair {
	s.CreationTime = &v
	return s
}

func (s *DescribeKeyPairsResponseKeyPairsKeyPair) SetKeyPairName(v string) *DescribeKeyPairsResponseKeyPairsKeyPair {
	s.KeyPairName = &v
	return s
}

func (s *DescribeKeyPairsResponseKeyPairsKeyPair) SetKeyPairFingerPrint(v string) *DescribeKeyPairsResponseKeyPairsKeyPair {
	s.KeyPairFingerPrint = &v
	return s
}

type DeleteKeyPairsRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty" require:"true"`
}

func (s DeleteKeyPairsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsRequest) SetVersion(v string) *DeleteKeyPairsRequest {
	s.Version = &v
	return s
}

func (s *DeleteKeyPairsRequest) SetKeyPairName(v string) *DeleteKeyPairsRequest {
	s.KeyPairName = &v
	return s
}

type DeleteKeyPairsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteKeyPairsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyPairsResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsResponse) SetRequestId(v string) *DeleteKeyPairsResponse {
	s.RequestId = &v
	return s
}

type CreateKeyPairRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty" require:"true"`
}

func (s CreateKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyPairRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyPairRequest) SetVersion(v string) *CreateKeyPairRequest {
	s.Version = &v
	return s
}

func (s *CreateKeyPairRequest) SetKeyPairName(v string) *CreateKeyPairRequest {
	s.KeyPairName = &v
	return s
}

type CreateKeyPairResponse struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	KeyPairId          *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty" require:"true"`
	PrivateKeyBody     *string `json:"PrivateKeyBody,omitempty" xml:"PrivateKeyBody,omitempty" require:"true"`
	KeyPairName        *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty" require:"true"`
	KeyPairFingerPrint *string `json:"KeyPairFingerPrint,omitempty" xml:"KeyPairFingerPrint,omitempty" require:"true"`
}

func (s CreateKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyPairResponse) GoString() string {
	return s.String()
}

func (s *CreateKeyPairResponse) SetRequestId(v string) *CreateKeyPairResponse {
	s.RequestId = &v
	return s
}

func (s *CreateKeyPairResponse) SetKeyPairId(v string) *CreateKeyPairResponse {
	s.KeyPairId = &v
	return s
}

func (s *CreateKeyPairResponse) SetPrivateKeyBody(v string) *CreateKeyPairResponse {
	s.PrivateKeyBody = &v
	return s
}

func (s *CreateKeyPairResponse) SetKeyPairName(v string) *CreateKeyPairResponse {
	s.KeyPairName = &v
	return s
}

func (s *CreateKeyPairResponse) SetKeyPairFingerPrint(v string) *CreateKeyPairResponse {
	s.KeyPairFingerPrint = &v
	return s
}

type ExportBillDetailDataRequest struct {
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
}

func (s ExportBillDetailDataRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportBillDetailDataRequest) GoString() string {
	return s.String()
}

func (s *ExportBillDetailDataRequest) SetVersion(v string) *ExportBillDetailDataRequest {
	s.Version = &v
	return s
}

func (s *ExportBillDetailDataRequest) SetStartDate(v string) *ExportBillDetailDataRequest {
	s.StartDate = &v
	return s
}

func (s *ExportBillDetailDataRequest) SetEndDate(v string) *ExportBillDetailDataRequest {
	s.EndDate = &v
	return s
}

type ExportBillDetailDataResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	FilePath  *string `json:"FilePath,omitempty" xml:"FilePath,omitempty" require:"true"`
}

func (s ExportBillDetailDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportBillDetailDataResponse) GoString() string {
	return s.String()
}

func (s *ExportBillDetailDataResponse) SetRequestId(v string) *ExportBillDetailDataResponse {
	s.RequestId = &v
	return s
}

func (s *ExportBillDetailDataResponse) SetFilePath(v string) *ExportBillDetailDataResponse {
	s.FilePath = &v
	return s
}

type DescribeEnsRegionIdResourceRequest struct {
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	OrderByParams *string `json:"OrderByParams,omitempty" xml:"OrderByParams,omitempty"`
	PageNumber    *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Isp           *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeEnsRegionIdResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionIdResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdResourceRequest) SetVersion(v string) *DescribeEnsRegionIdResourceRequest {
	s.Version = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetStartTime(v string) *DescribeEnsRegionIdResourceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetEndTime(v string) *DescribeEnsRegionIdResourceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetOrderByParams(v string) *DescribeEnsRegionIdResourceRequest {
	s.OrderByParams = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetPageNumber(v int) *DescribeEnsRegionIdResourceRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetPageSize(v string) *DescribeEnsRegionIdResourceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEnsRegionIdResourceRequest) SetIsp(v string) *DescribeEnsRegionIdResourceRequest {
	s.Isp = &v
	return s
}

type DescribeEnsRegionIdResourceResponse struct {
	RequestId            *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount           *int                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber           *int                                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize             *int                                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	EnsRegionIdResources *DescribeEnsRegionIdResourceResponseEnsRegionIdResources `json:"EnsRegionIdResources,omitempty" xml:"EnsRegionIdResources,omitempty" require:"true" type:"Struct"`
}

func (s DescribeEnsRegionIdResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionIdResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdResourceResponse) SetRequestId(v string) *DescribeEnsRegionIdResourceResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponse) SetTotalCount(v int) *DescribeEnsRegionIdResourceResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponse) SetPageNumber(v int) *DescribeEnsRegionIdResourceResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponse) SetPageSize(v int) *DescribeEnsRegionIdResourceResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponse) SetEnsRegionIdResources(v *DescribeEnsRegionIdResourceResponseEnsRegionIdResources) *DescribeEnsRegionIdResourceResponse {
	s.EnsRegionIdResources = v
	return s
}

type DescribeEnsRegionIdResourceResponseEnsRegionIdResources struct {
	EnsRegionIdResource []*DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource `json:"EnsRegionIdResource,omitempty" xml:"EnsRegionIdResource,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeEnsRegionIdResourceResponseEnsRegionIdResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionIdResourceResponseEnsRegionIdResources) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdResourceResponseEnsRegionIdResources) SetEnsRegionIdResource(v []*DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource) *DescribeEnsRegionIdResourceResponseEnsRegionIdResources {
	s.EnsRegionIdResource = v
	return s
}

type DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource struct {
	Area              *string `json:"Area,omitempty" xml:"Area,omitempty" require:"true"`
	AreaCode          *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty" require:"true"`
	EnsRegionId       *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	EnsRegionIdName   *string `json:"EnsRegionIdName,omitempty" xml:"EnsRegionIdName,omitempty" require:"true"`
	VCpu              *int    `json:"VCpu,omitempty" xml:"VCpu,omitempty" require:"true"`
	InternetBandwidth *int    `json:"InternetBandwidth,omitempty" xml:"InternetBandwidth,omitempty" require:"true"`
	Isp               *string `json:"Isp,omitempty" xml:"Isp,omitempty" require:"true"`
	BizDate           *string `json:"BizDate,omitempty" xml:"BizDate,omitempty" require:"true"`
	InstanceCount     *int    `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty" require:"true"`
}

func (s DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource) SetArea(v string) *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource {
	s.Area = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource) SetAreaCode(v string) *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource {
	s.AreaCode = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource) SetEnsRegionId(v string) *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource) SetEnsRegionIdName(v string) *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource {
	s.EnsRegionIdName = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource) SetVCpu(v int) *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource {
	s.VCpu = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource) SetInternetBandwidth(v int) *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource {
	s.InternetBandwidth = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource) SetIsp(v string) *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource {
	s.Isp = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource) SetBizDate(v string) *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource {
	s.BizDate = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource) SetInstanceCount(v int) *DescribeEnsRegionIdResourceResponseEnsRegionIdResourcesEnsRegionIdResource {
	s.InstanceCount = &v
	return s
}

type DescribeBandwitdhByInternetChargeTypeRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Isp         *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
}

func (s DescribeBandwitdhByInternetChargeTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandwitdhByInternetChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) SetVersion(v string) *DescribeBandwitdhByInternetChargeTypeRequest {
	s.Version = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) SetStartTime(v string) *DescribeBandwitdhByInternetChargeTypeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) SetEndTime(v string) *DescribeBandwitdhByInternetChargeTypeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) SetIsp(v string) *DescribeBandwitdhByInternetChargeTypeRequest {
	s.Isp = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) SetEnsRegionId(v string) *DescribeBandwitdhByInternetChargeTypeRequest {
	s.EnsRegionId = &v
	return s
}

type DescribeBandwitdhByInternetChargeTypeResponse struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty" require:"true"`
	BandwidthValue     *int64  `json:"BandwidthValue,omitempty" xml:"BandwidthValue,omitempty" require:"true"`
	TimeStamp          *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
}

func (s DescribeBandwitdhByInternetChargeTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandwitdhByInternetChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeBandwitdhByInternetChargeTypeResponse) SetRequestId(v string) *DescribeBandwitdhByInternetChargeTypeResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeResponse) SetInternetChargeType(v string) *DescribeBandwitdhByInternetChargeTypeResponse {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeResponse) SetBandwidthValue(v int64) *DescribeBandwitdhByInternetChargeTypeResponse {
	s.BandwidthValue = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeResponse) SetTimeStamp(v string) *DescribeBandwitdhByInternetChargeTypeResponse {
	s.TimeStamp = &v
	return s
}

type AuthorizeSecurityGroupRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty" require:"true"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty" require:"true"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" require:"true"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Priority        *int    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	SourceCidrIp    *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty" require:"true"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s AuthorizeSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupRequest) SetVersion(v string) *AuthorizeSecurityGroupRequest {
	s.Version = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetIpProtocol(v string) *AuthorizeSecurityGroupRequest {
	s.IpProtocol = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPortRange(v string) *AuthorizeSecurityGroupRequest {
	s.PortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSecurityGroupId(v string) *AuthorizeSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPolicy(v string) *AuthorizeSecurityGroupRequest {
	s.Policy = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPriority(v int) *AuthorizeSecurityGroupRequest {
	s.Priority = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourceCidrIp(v string) *AuthorizeSecurityGroupRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourcePortRange(v string) *AuthorizeSecurityGroupRequest {
	s.SourcePortRange = &v
	return s
}

type AuthorizeSecurityGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AuthorizeSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupResponse) SetRequestId(v string) *AuthorizeSecurityGroupResponse {
	s.RequestId = &v
	return s
}

type RevokeSecurityGroupRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty" require:"true"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty" require:"true"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" require:"true"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Priority        *int    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	SourceCidrIp    *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty" require:"true"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s RevokeSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupRequest) SetVersion(v string) *RevokeSecurityGroupRequest {
	s.Version = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetIpProtocol(v string) *RevokeSecurityGroupRequest {
	s.IpProtocol = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetPortRange(v string) *RevokeSecurityGroupRequest {
	s.PortRange = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSecurityGroupId(v string) *RevokeSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetPolicy(v string) *RevokeSecurityGroupRequest {
	s.Policy = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetPriority(v int) *RevokeSecurityGroupRequest {
	s.Priority = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSourceCidrIp(v string) *RevokeSecurityGroupRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSourcePortRange(v string) *RevokeSecurityGroupRequest {
	s.SourcePortRange = &v
	return s
}

type RevokeSecurityGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RevokeSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupResponse) SetRequestId(v string) *RevokeSecurityGroupResponse {
	s.RequestId = &v
	return s
}

type DeleteSecurityGroupRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" require:"true"`
}

func (s DeleteSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupRequest) SetVersion(v string) *DeleteSecurityGroupRequest {
	s.Version = &v
	return s
}

func (s *DeleteSecurityGroupRequest) SetSecurityGroupId(v string) *DeleteSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

type DeleteSecurityGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupResponse) SetRequestId(v string) *DeleteSecurityGroupResponse {
	s.RequestId = &v
	return s
}

type DescribeSecurityGroupAttributeRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" require:"true"`
}

func (s DescribeSecurityGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeRequest) SetVersion(v string) *DescribeSecurityGroupAttributeRequest {
	s.Version = &v
	return s
}

func (s *DescribeSecurityGroupAttributeRequest) SetSecurityGroupId(v string) *DescribeSecurityGroupAttributeRequest {
	s.SecurityGroupId = &v
	return s
}

type DescribeSecurityGroupAttributeResponse struct {
	RequestId       *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SecurityGroupId *string                                            `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" require:"true"`
	Permissions     *DescribeSecurityGroupAttributeResponsePermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" require:"true" type:"Struct"`
}

func (s DescribeSecurityGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeResponse) SetRequestId(v string) *DescribeSecurityGroupAttributeResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponse) SetSecurityGroupId(v string) *DescribeSecurityGroupAttributeResponse {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponse) SetPermissions(v *DescribeSecurityGroupAttributeResponsePermissions) *DescribeSecurityGroupAttributeResponse {
	s.Permissions = v
	return s
}

type DescribeSecurityGroupAttributeResponsePermissions struct {
	Permission []*DescribeSecurityGroupAttributeResponsePermissionsPermission `json:"Permission,omitempty" xml:"Permission,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSecurityGroupAttributeResponsePermissions) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupAttributeResponsePermissions) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeResponsePermissions) SetPermission(v []*DescribeSecurityGroupAttributeResponsePermissionsPermission) *DescribeSecurityGroupAttributeResponsePermissions {
	s.Permission = v
	return s
}

type DescribeSecurityGroupAttributeResponsePermissionsPermission struct {
	DestCidrIp      *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty" require:"true"`
	SourceCidrIp    *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty" require:"true"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty" require:"true"`
	Priority        *int    `json:"Priority,omitempty" xml:"Priority,omitempty" require:"true"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty" require:"true"`
	Direction       *string `json:"Direction,omitempty" xml:"Direction,omitempty" require:"true"`
	CreationTime    *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty" require:"true"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty" require:"true"`
}

func (s DescribeSecurityGroupAttributeResponsePermissionsPermission) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupAttributeResponsePermissionsPermission) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeResponsePermissionsPermission) SetDestCidrIp(v string) *DescribeSecurityGroupAttributeResponsePermissionsPermission {
	s.DestCidrIp = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponsePermissionsPermission) SetSourceCidrIp(v string) *DescribeSecurityGroupAttributeResponsePermissionsPermission {
	s.SourceCidrIp = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponsePermissionsPermission) SetIpProtocol(v string) *DescribeSecurityGroupAttributeResponsePermissionsPermission {
	s.IpProtocol = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponsePermissionsPermission) SetPriority(v int) *DescribeSecurityGroupAttributeResponsePermissionsPermission {
	s.Priority = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponsePermissionsPermission) SetPolicy(v string) *DescribeSecurityGroupAttributeResponsePermissionsPermission {
	s.Policy = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponsePermissionsPermission) SetDirection(v string) *DescribeSecurityGroupAttributeResponsePermissionsPermission {
	s.Direction = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponsePermissionsPermission) SetCreationTime(v string) *DescribeSecurityGroupAttributeResponsePermissionsPermission {
	s.CreationTime = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponsePermissionsPermission) SetPortRange(v string) *DescribeSecurityGroupAttributeResponsePermissionsPermission {
	s.PortRange = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponsePermissionsPermission) SetSourcePortRange(v string) *DescribeSecurityGroupAttributeResponsePermissionsPermission {
	s.SourcePortRange = &v
	return s
}

type LeaveSecurityGroupRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" require:"true"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s LeaveSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s LeaveSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *LeaveSecurityGroupRequest) SetVersion(v string) *LeaveSecurityGroupRequest {
	s.Version = &v
	return s
}

func (s *LeaveSecurityGroupRequest) SetSecurityGroupId(v string) *LeaveSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *LeaveSecurityGroupRequest) SetInstanceId(v string) *LeaveSecurityGroupRequest {
	s.InstanceId = &v
	return s
}

type LeaveSecurityGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s LeaveSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s LeaveSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *LeaveSecurityGroupResponse) SetRequestId(v string) *LeaveSecurityGroupResponse {
	s.RequestId = &v
	return s
}

type JoinSecurityGroupRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" require:"true"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s JoinSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *JoinSecurityGroupRequest) SetVersion(v string) *JoinSecurityGroupRequest {
	s.Version = &v
	return s
}

func (s *JoinSecurityGroupRequest) SetSecurityGroupId(v string) *JoinSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *JoinSecurityGroupRequest) SetInstanceId(v string) *JoinSecurityGroupRequest {
	s.InstanceId = &v
	return s
}

type JoinSecurityGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s JoinSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *JoinSecurityGroupResponse) SetRequestId(v string) *JoinSecurityGroupResponse {
	s.RequestId = &v
	return s
}

type AuthorizeSecurityGroupEgressRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty" require:"true"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty" require:"true"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" require:"true"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Priority        *int    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	DestCidrIp      *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty" require:"true"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s AuthorizeSecurityGroupEgressRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeSecurityGroupEgressRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupEgressRequest) SetVersion(v string) *AuthorizeSecurityGroupEgressRequest {
	s.Version = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetIpProtocol(v string) *AuthorizeSecurityGroupEgressRequest {
	s.IpProtocol = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPortRange(v string) *AuthorizeSecurityGroupEgressRequest {
	s.PortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetSecurityGroupId(v string) *AuthorizeSecurityGroupEgressRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPolicy(v string) *AuthorizeSecurityGroupEgressRequest {
	s.Policy = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPriority(v int) *AuthorizeSecurityGroupEgressRequest {
	s.Priority = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetDestCidrIp(v string) *AuthorizeSecurityGroupEgressRequest {
	s.DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetSourcePortRange(v string) *AuthorizeSecurityGroupEgressRequest {
	s.SourcePortRange = &v
	return s
}

type AuthorizeSecurityGroupEgressResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AuthorizeSecurityGroupEgressResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeSecurityGroupEgressResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupEgressResponse) SetRequestId(v string) *AuthorizeSecurityGroupEgressResponse {
	s.RequestId = &v
	return s
}

type RevokeSecurityGroupEgressRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty" require:"true"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty" require:"true"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" require:"true"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Priority        *int    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	DestCidrIp      *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty" require:"true"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s RevokeSecurityGroupEgressRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeSecurityGroupEgressRequest) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupEgressRequest) SetVersion(v string) *RevokeSecurityGroupEgressRequest {
	s.Version = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetIpProtocol(v string) *RevokeSecurityGroupEgressRequest {
	s.IpProtocol = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetPortRange(v string) *RevokeSecurityGroupEgressRequest {
	s.PortRange = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetSecurityGroupId(v string) *RevokeSecurityGroupEgressRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetPolicy(v string) *RevokeSecurityGroupEgressRequest {
	s.Policy = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetPriority(v int) *RevokeSecurityGroupEgressRequest {
	s.Priority = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetDestCidrIp(v string) *RevokeSecurityGroupEgressRequest {
	s.DestCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetSourcePortRange(v string) *RevokeSecurityGroupEgressRequest {
	s.SourcePortRange = &v
	return s
}

type RevokeSecurityGroupEgressResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RevokeSecurityGroupEgressResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeSecurityGroupEgressResponse) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupEgressResponse) SetRequestId(v string) *RevokeSecurityGroupEgressResponse {
	s.RequestId = &v
	return s
}

type DescribeSecurityGroupsRequest struct {
	Version           *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	SecurityGroupId   *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	PageNumber        *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s DescribeSecurityGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsRequest) SetVersion(v string) *DescribeSecurityGroupsRequest {
	s.Version = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetSecurityGroupId(v string) *DescribeSecurityGroupsRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetPageNumber(v int) *DescribeSecurityGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetPageSize(v int) *DescribeSecurityGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetSecurityGroupName(v string) *DescribeSecurityGroupsRequest {
	s.SecurityGroupName = &v
	return s
}

type DescribeSecurityGroupsResponse struct {
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount     *int                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber     *int                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize       *int                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	SecurityGroups *DescribeSecurityGroupsResponseSecurityGroups `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" require:"true" type:"Struct"`
}

func (s DescribeSecurityGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponse) SetRequestId(v string) *DescribeSecurityGroupsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityGroupsResponse) SetTotalCount(v int) *DescribeSecurityGroupsResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecurityGroupsResponse) SetPageNumber(v int) *DescribeSecurityGroupsResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeSecurityGroupsResponse) SetPageSize(v int) *DescribeSecurityGroupsResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeSecurityGroupsResponse) SetSecurityGroups(v *DescribeSecurityGroupsResponseSecurityGroups) *DescribeSecurityGroupsResponse {
	s.SecurityGroups = v
	return s
}

type DescribeSecurityGroupsResponseSecurityGroups struct {
	SecurityGroup []*DescribeSecurityGroupsResponseSecurityGroupsSecurityGroup `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSecurityGroupsResponseSecurityGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsResponseSecurityGroups) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseSecurityGroups) SetSecurityGroup(v []*DescribeSecurityGroupsResponseSecurityGroupsSecurityGroup) *DescribeSecurityGroupsResponseSecurityGroups {
	s.SecurityGroup = v
	return s
}

type DescribeSecurityGroupsResponseSecurityGroupsSecurityGroup struct {
	SecurityGroupId   *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" require:"true"`
	CreationTime      *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty" require:"true"`
}

func (s DescribeSecurityGroupsResponseSecurityGroupsSecurityGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsResponseSecurityGroupsSecurityGroup) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseSecurityGroupsSecurityGroup) SetSecurityGroupId(v string) *DescribeSecurityGroupsResponseSecurityGroupsSecurityGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupsResponseSecurityGroupsSecurityGroup) SetCreationTime(v string) *DescribeSecurityGroupsResponseSecurityGroupsSecurityGroup {
	s.CreationTime = &v
	return s
}

func (s *DescribeSecurityGroupsResponseSecurityGroupsSecurityGroup) SetSecurityGroupName(v string) *DescribeSecurityGroupsResponseSecurityGroupsSecurityGroup {
	s.SecurityGroupName = &v
	return s
}

type CreateSecurityGroupRequest struct {
	Version           *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s CreateSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityGroupRequest) SetVersion(v string) *CreateSecurityGroupRequest {
	s.Version = &v
	return s
}

func (s *CreateSecurityGroupRequest) SetSecurityGroupName(v string) *CreateSecurityGroupRequest {
	s.SecurityGroupName = &v
	return s
}

type CreateSecurityGroupResponse struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" require:"true"`
}

func (s CreateSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityGroupResponse) SetRequestId(v string) *CreateSecurityGroupResponse {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityGroupResponse) SetSecurityGroupId(v string) *CreateSecurityGroupResponse {
	s.SecurityGroupId = &v
	return s
}

type DescribeEnsRegionIdIpv6InfoRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
}

func (s DescribeEnsRegionIdIpv6InfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionIdIpv6InfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdIpv6InfoRequest) SetVersion(v string) *DescribeEnsRegionIdIpv6InfoRequest {
	s.Version = &v
	return s
}

func (s *DescribeEnsRegionIdIpv6InfoRequest) SetEnsRegionId(v string) *DescribeEnsRegionIdIpv6InfoRequest {
	s.EnsRegionId = &v
	return s
}

type DescribeEnsRegionIdIpv6InfoResponse struct {
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SupportIpv6Info *DescribeEnsRegionIdIpv6InfoResponseSupportIpv6Info `json:"SupportIpv6Info,omitempty" xml:"SupportIpv6Info,omitempty" require:"true" type:"Struct"`
}

func (s DescribeEnsRegionIdIpv6InfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionIdIpv6InfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdIpv6InfoResponse) SetRequestId(v string) *DescribeEnsRegionIdIpv6InfoResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsRegionIdIpv6InfoResponse) SetSupportIpv6Info(v *DescribeEnsRegionIdIpv6InfoResponseSupportIpv6Info) *DescribeEnsRegionIdIpv6InfoResponse {
	s.SupportIpv6Info = v
	return s
}

type DescribeEnsRegionIdIpv6InfoResponseSupportIpv6Info struct {
	SupportIpv6 *bool   `json:"SupportIpv6,omitempty" xml:"SupportIpv6,omitempty" require:"true"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
}

func (s DescribeEnsRegionIdIpv6InfoResponseSupportIpv6Info) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionIdIpv6InfoResponseSupportIpv6Info) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdIpv6InfoResponseSupportIpv6Info) SetSupportIpv6(v bool) *DescribeEnsRegionIdIpv6InfoResponseSupportIpv6Info {
	s.SupportIpv6 = &v
	return s
}

func (s *DescribeEnsRegionIdIpv6InfoResponseSupportIpv6Info) SetEnsRegionId(v string) *DescribeEnsRegionIdIpv6InfoResponseSupportIpv6Info {
	s.EnsRegionId = &v
	return s
}

type DescribeCreatePrePaidInstanceResultRequest struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s DescribeCreatePrePaidInstanceResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCreatePrePaidInstanceResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeCreatePrePaidInstanceResultRequest) SetVersion(v string) *DescribeCreatePrePaidInstanceResultRequest {
	s.Version = &v
	return s
}

func (s *DescribeCreatePrePaidInstanceResultRequest) SetInstanceId(v string) *DescribeCreatePrePaidInstanceResultRequest {
	s.InstanceId = &v
	return s
}

type DescribeCreatePrePaidInstanceResultResponse struct {
	RequestId            *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	InstanceCreateResult *DescribeCreatePrePaidInstanceResultResponseInstanceCreateResult `json:"InstanceCreateResult,omitempty" xml:"InstanceCreateResult,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCreatePrePaidInstanceResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCreatePrePaidInstanceResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeCreatePrePaidInstanceResultResponse) SetRequestId(v string) *DescribeCreatePrePaidInstanceResultResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCreatePrePaidInstanceResultResponse) SetInstanceCreateResult(v *DescribeCreatePrePaidInstanceResultResponseInstanceCreateResult) *DescribeCreatePrePaidInstanceResultResponse {
	s.InstanceCreateResult = v
	return s
}

type DescribeCreatePrePaidInstanceResultResponseInstanceCreateResult struct {
	InstanceCreateStatus *string `json:"InstanceCreateStatus,omitempty" xml:"InstanceCreateStatus,omitempty" require:"true"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s DescribeCreatePrePaidInstanceResultResponseInstanceCreateResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeCreatePrePaidInstanceResultResponseInstanceCreateResult) GoString() string {
	return s.String()
}

func (s *DescribeCreatePrePaidInstanceResultResponseInstanceCreateResult) SetInstanceCreateStatus(v string) *DescribeCreatePrePaidInstanceResultResponseInstanceCreateResult {
	s.InstanceCreateStatus = &v
	return s
}

func (s *DescribeCreatePrePaidInstanceResultResponseInstanceCreateResult) SetInstanceId(v string) *DescribeCreatePrePaidInstanceResultResponseInstanceCreateResult {
	s.InstanceId = &v
	return s
}

type DescribePriceRequest struct {
	Version            *string                         `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	InstanceType       *string                         `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" require:"true"`
	EnsRegionId        *string                         `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	Period             *int                            `json:"Period,omitempty" xml:"Period,omitempty" require:"true"`
	SystemDisk         *DescribePriceRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	Quantity           *int                            `json:"Quantity,omitempty" xml:"Quantity,omitempty" require:"true"`
	DataDisk           []*DescribePriceRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	InternetChargeType *string                         `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty" require:"true"`
}

func (s DescribePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) SetVersion(v string) *DescribePriceRequest {
	s.Version = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceType(v string) *DescribePriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribePriceRequest) SetEnsRegionId(v string) *DescribePriceRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribePriceRequest) SetPeriod(v int) *DescribePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribePriceRequest) SetSystemDisk(v *DescribePriceRequestSystemDisk) *DescribePriceRequest {
	s.SystemDisk = v
	return s
}

func (s *DescribePriceRequest) SetQuantity(v int) *DescribePriceRequest {
	s.Quantity = &v
	return s
}

func (s *DescribePriceRequest) SetDataDisk(v []*DescribePriceRequestDataDisk) *DescribePriceRequest {
	s.DataDisk = v
	return s
}

func (s *DescribePriceRequest) SetInternetChargeType(v string) *DescribePriceRequest {
	s.InternetChargeType = &v
	return s
}

type DescribePriceRequestSystemDisk struct {
	Size *int `json:"Size,omitempty" xml:"Size,omitempty" require:"true"`
}

func (s DescribePriceRequestSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestSystemDisk) SetSize(v int) *DescribePriceRequestSystemDisk {
	s.Size = &v
	return s
}

type DescribePriceRequestDataDisk struct {
	Size *int `json:"Size,omitempty" xml:"Size,omitempty" require:"true"`
}

func (s DescribePriceRequestDataDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequestDataDisk) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestDataDisk) SetSize(v int) *DescribePriceRequestDataDisk {
	s.Size = &v
	return s
}

type DescribePriceResponse struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PriceInfo *DescribePriceResponsePriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" require:"true" type:"Struct"`
}

func (s DescribePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribePriceResponse) SetRequestId(v string) *DescribePriceResponse {
	s.RequestId = &v
	return s
}

func (s *DescribePriceResponse) SetPriceInfo(v *DescribePriceResponsePriceInfo) *DescribePriceResponse {
	s.PriceInfo = v
	return s
}

type DescribePriceResponsePriceInfo struct {
	Price *DescribePriceResponsePriceInfoPrice `json:"Price,omitempty" xml:"Price,omitempty" require:"true" type:"Struct"`
}

func (s DescribePriceResponsePriceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponsePriceInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponsePriceInfo) SetPrice(v *DescribePriceResponsePriceInfoPrice) *DescribePriceResponsePriceInfo {
	s.Price = v
	return s
}

type DescribePriceResponsePriceInfoPrice struct {
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty" require:"true"`
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty" require:"true"`
	TradePrice    *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty" require:"true"`
	Currency      *string  `json:"Currency,omitempty" xml:"Currency,omitempty" require:"true"`
}

func (s DescribePriceResponsePriceInfoPrice) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponsePriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribePriceResponsePriceInfoPrice) SetDiscountPrice(v float32) *DescribePriceResponsePriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePriceResponsePriceInfoPrice) SetOriginalPrice(v float32) *DescribePriceResponsePriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponsePriceInfoPrice) SetTradePrice(v float32) *DescribePriceResponsePriceInfoPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribePriceResponsePriceInfoPrice) SetCurrency(v string) *DescribePriceResponsePriceInfoPrice {
	s.Currency = &v
	return s
}

type ExportMeasurementDataRequest struct {
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
}

func (s ExportMeasurementDataRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportMeasurementDataRequest) GoString() string {
	return s.String()
}

func (s *ExportMeasurementDataRequest) SetVersion(v string) *ExportMeasurementDataRequest {
	s.Version = &v
	return s
}

func (s *ExportMeasurementDataRequest) SetStartDate(v string) *ExportMeasurementDataRequest {
	s.StartDate = &v
	return s
}

func (s *ExportMeasurementDataRequest) SetEndDate(v string) *ExportMeasurementDataRequest {
	s.EndDate = &v
	return s
}

type ExportMeasurementDataResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	FilePath  *string `json:"FilePath,omitempty" xml:"FilePath,omitempty" require:"true"`
}

func (s ExportMeasurementDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportMeasurementDataResponse) GoString() string {
	return s.String()
}

func (s *ExportMeasurementDataResponse) SetRequestId(v string) *ExportMeasurementDataResponse {
	s.RequestId = &v
	return s
}

func (s *ExportMeasurementDataResponse) SetFilePath(v string) *ExportMeasurementDataResponse {
	s.FilePath = &v
	return s
}

type DescribeMeasurementDataRequest struct {
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
}

func (s DescribeMeasurementDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataRequest) SetVersion(v string) *DescribeMeasurementDataRequest {
	s.Version = &v
	return s
}

func (s *DescribeMeasurementDataRequest) SetStartDate(v string) *DescribeMeasurementDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeMeasurementDataRequest) SetEndDate(v string) *DescribeMeasurementDataRequest {
	s.EndDate = &v
	return s
}

type DescribeMeasurementDataResponse struct {
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	MeasurementDatas *DescribeMeasurementDataResponseMeasurementDatas `json:"MeasurementDatas,omitempty" xml:"MeasurementDatas,omitempty" require:"true" type:"Struct"`
}

func (s DescribeMeasurementDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponse) SetRequestId(v string) *DescribeMeasurementDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeMeasurementDataResponse) SetMeasurementDatas(v *DescribeMeasurementDataResponseMeasurementDatas) *DescribeMeasurementDataResponse {
	s.MeasurementDatas = v
	return s
}

type DescribeMeasurementDataResponseMeasurementDatas struct {
	MeasurementData []*DescribeMeasurementDataResponseMeasurementDatasMeasurementData `json:"MeasurementData,omitempty" xml:"MeasurementData,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeMeasurementDataResponseMeasurementDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponseMeasurementDatas) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseMeasurementDatas) SetMeasurementData(v []*DescribeMeasurementDataResponseMeasurementDatasMeasurementData) *DescribeMeasurementDataResponseMeasurementDatas {
	s.MeasurementData = v
	return s
}

type DescribeMeasurementDataResponseMeasurementDatasMeasurementData struct {
	ChargeModel            *string                                                                               `json:"ChargeModel,omitempty" xml:"ChargeModel,omitempty" require:"true"`
	CostCycle              *string                                                                               `json:"CostCycle,omitempty" xml:"CostCycle,omitempty" require:"true"`
	CostStartTime          *string                                                                               `json:"CostStartTime,omitempty" xml:"CostStartTime,omitempty" require:"true"`
	CostEndTime            *string                                                                               `json:"CostEndTime,omitempty" xml:"CostEndTime,omitempty" require:"true"`
	BandWidthFeeDatas      *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatas      `json:"BandWidthFeeDatas,omitempty" xml:"BandWidthFeeDatas,omitempty" require:"true" type:"Struct"`
	ResourceFeeDataDetails *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetails `json:"ResourceFeeDataDetails,omitempty" xml:"ResourceFeeDataDetails,omitempty" require:"true" type:"Struct"`
	ResourceFeeData        *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeData        `json:"ResourceFeeData,omitempty" xml:"ResourceFeeData,omitempty" require:"true" type:"Struct"`
}

func (s DescribeMeasurementDataResponseMeasurementDatasMeasurementData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponseMeasurementDatasMeasurementData) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementData) SetChargeModel(v string) *DescribeMeasurementDataResponseMeasurementDatasMeasurementData {
	s.ChargeModel = &v
	return s
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementData) SetCostCycle(v string) *DescribeMeasurementDataResponseMeasurementDatasMeasurementData {
	s.CostCycle = &v
	return s
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementData) SetCostStartTime(v string) *DescribeMeasurementDataResponseMeasurementDatasMeasurementData {
	s.CostStartTime = &v
	return s
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementData) SetCostEndTime(v string) *DescribeMeasurementDataResponseMeasurementDatasMeasurementData {
	s.CostEndTime = &v
	return s
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementData) SetBandWidthFeeDatas(v *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatas) *DescribeMeasurementDataResponseMeasurementDatasMeasurementData {
	s.BandWidthFeeDatas = v
	return s
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementData) SetResourceFeeDataDetails(v *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetails) *DescribeMeasurementDataResponseMeasurementDatasMeasurementData {
	s.ResourceFeeDataDetails = v
	return s
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementData) SetResourceFeeData(v *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeData) *DescribeMeasurementDataResponseMeasurementDatasMeasurementData {
	s.ResourceFeeData = v
	return s
}

type DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatas struct {
	BandWidthFeeData []*DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData `json:"BandWidthFeeData,omitempty" xml:"BandWidthFeeData,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatas) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatas) SetBandWidthFeeData(v []*DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatas {
	s.BandWidthFeeData = v
	return s
}

type DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData struct {
	CostVal  *int    `json:"CostVal,omitempty" xml:"CostVal,omitempty" require:"true"`
	CostCode *string `json:"CostCode,omitempty" xml:"CostCode,omitempty" require:"true"`
	CostName *string `json:"CostName,omitempty" xml:"CostName,omitempty" require:"true"`
}

func (s DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostVal(v int) *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostVal = &v
	return s
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostCode(v string) *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostCode = &v
	return s
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData) SetCostName(v string) *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataBandWidthFeeDatasBandWidthFeeData {
	s.CostName = &v
	return s
}

type DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetails struct {
	ResourceFeeDataDetail []*DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail `json:"ResourceFeeDataDetail,omitempty" xml:"ResourceFeeDataDetail,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetails) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetails) SetResourceFeeDataDetail(v []*DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetails {
	s.ResourceFeeDataDetail = v
	return s
}

type DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail struct {
	CostVal      *int    `json:"CostVal,omitempty" xml:"CostVal,omitempty" require:"true"`
	CostCode     *string `json:"CostCode,omitempty" xml:"CostCode,omitempty" require:"true"`
	CostName     *string `json:"CostName,omitempty" xml:"CostName,omitempty" require:"true"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
}

func (s DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) SetCostVal(v int) *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail {
	s.CostVal = &v
	return s
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) SetCostCode(v string) *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail {
	s.CostCode = &v
	return s
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) SetCostName(v string) *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail {
	s.CostName = &v
	return s
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail) SetResourceType(v string) *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeDataDetailsResourceFeeDataDetail {
	s.ResourceType = &v
	return s
}

type DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeData struct {
	Storage *int `json:"Storage,omitempty" xml:"Storage,omitempty" require:"true"`
	Memory  *int `json:"Memory,omitempty" xml:"Memory,omitempty" require:"true"`
	Vcpu    *int `json:"Vcpu,omitempty" xml:"Vcpu,omitempty" require:"true"`
}

func (s DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeData) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeData) SetStorage(v int) *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeData {
	s.Storage = &v
	return s
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeData) SetMemory(v int) *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeData {
	s.Memory = &v
	return s
}

func (s *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeData) SetVcpu(v int) *DescribeMeasurementDataResponseMeasurementDatasMeasurementDataResourceFeeData {
	s.Vcpu = &v
	return s
}

type DescribeAvailableResourceInfoRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s DescribeAvailableResourceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoRequest) SetVersion(v string) *DescribeAvailableResourceInfoRequest {
	s.Version = &v
	return s
}

type DescribeAvailableResourceInfoResponse struct {
	RequestId        *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SupportResources *DescribeAvailableResourceInfoResponseSupportResources `json:"SupportResources,omitempty" xml:"SupportResources,omitempty" require:"true" type:"Struct"`
	Images           *DescribeAvailableResourceInfoResponseImages           `json:"Images,omitempty" xml:"Images,omitempty" require:"true" type:"Struct"`
}

func (s DescribeAvailableResourceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponse) SetRequestId(v string) *DescribeAvailableResourceInfoResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponse) SetSupportResources(v *DescribeAvailableResourceInfoResponseSupportResources) *DescribeAvailableResourceInfoResponse {
	s.SupportResources = v
	return s
}

func (s *DescribeAvailableResourceInfoResponse) SetImages(v *DescribeAvailableResourceInfoResponseImages) *DescribeAvailableResourceInfoResponse {
	s.Images = v
	return s
}

type DescribeAvailableResourceInfoResponseSupportResources struct {
	SupportResource []*DescribeAvailableResourceInfoResponseSupportResourcesSupportResource `json:"SupportResource,omitempty" xml:"SupportResource,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseSupportResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseSupportResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseSupportResources) SetSupportResource(v []*DescribeAvailableResourceInfoResponseSupportResourcesSupportResource) *DescribeAvailableResourceInfoResponseSupportResources {
	s.SupportResource = v
	return s
}

type DescribeAvailableResourceInfoResponseSupportResourcesSupportResource struct {
	DataDiskMinSize     *int                                                                                     `json:"DataDiskMinSize,omitempty" xml:"DataDiskMinSize,omitempty" require:"true"`
	DataDiskMaxSize     *int                                                                                     `json:"DataDiskMaxSize,omitempty" xml:"DataDiskMaxSize,omitempty" require:"true"`
	SystemDiskMinSize   *int                                                                                     `json:"SystemDiskMinSize,omitempty" xml:"SystemDiskMinSize,omitempty" require:"true"`
	SystemDiskMaxSize   *int                                                                                     `json:"SystemDiskMaxSize,omitempty" xml:"SystemDiskMaxSize,omitempty" require:"true"`
	EnsRegionIdsExtends *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtends `json:"EnsRegionIdsExtends,omitempty" xml:"EnsRegionIdsExtends,omitempty" require:"true" type:"Struct"`
	EnsRegionIds        *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIds        `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty" require:"true" type:"Struct"`
	InstanceSpeces      *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceInstanceSpeces      `json:"InstanceSpeces,omitempty" xml:"InstanceSpeces,omitempty" require:"true" type:"Struct"`
	BandwidthTypes      *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceBandwidthTypes      `json:"BandwidthTypes,omitempty" xml:"BandwidthTypes,omitempty" require:"true" type:"Struct"`
}

func (s DescribeAvailableResourceInfoResponseSupportResourcesSupportResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseSupportResourcesSupportResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResource) SetDataDiskMinSize(v int) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResource {
	s.DataDiskMinSize = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResource) SetDataDiskMaxSize(v int) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResource {
	s.DataDiskMaxSize = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResource) SetSystemDiskMinSize(v int) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResource {
	s.SystemDiskMinSize = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResource) SetSystemDiskMaxSize(v int) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResource {
	s.SystemDiskMaxSize = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResource) SetEnsRegionIdsExtends(v *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtends) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResource {
	s.EnsRegionIdsExtends = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResource) SetEnsRegionIds(v *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIds) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResource {
	s.EnsRegionIds = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResource) SetInstanceSpeces(v *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceInstanceSpeces) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResource {
	s.InstanceSpeces = v
	return s
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResource) SetBandwidthTypes(v *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceBandwidthTypes) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResource {
	s.BandwidthTypes = v
	return s
}

type DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtends struct {
	EnsRegionId []*DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtends) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtends) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtends) SetEnsRegionId(v []*DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtends {
	s.EnsRegionId = v
	return s
}

type DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId struct {
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	EnName      *string `json:"EnName,omitempty" xml:"EnName,omitempty" require:"true"`
	Area        *string `json:"Area,omitempty" xml:"Area,omitempty" require:"true"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty" require:"true"`
}

func (s DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) SetEnsRegionId(v string) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) SetName(v string) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	s.Name = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) SetEnName(v string) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	s.EnName = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) SetArea(v string) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	s.Area = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId) SetProvince(v string) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIdsExtendsEnsRegionId {
	s.Province = &v
	return s
}

type DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIds struct {
	// EnsRegionId
	EnsRegionId []*string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIds) SetEnsRegionId(v []*string) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceEnsRegionIds {
	s.EnsRegionId = v
	return s
}

type DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceInstanceSpeces struct {
	// InstanceSpec
	InstanceSpec []*string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceInstanceSpeces) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceInstanceSpeces) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceInstanceSpeces) SetInstanceSpec(v []*string) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceInstanceSpeces {
	s.InstanceSpec = v
	return s
}

type DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceBandwidthTypes struct {
	// BandwidthType
	BandwidthType []*string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceBandwidthTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceBandwidthTypes) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceBandwidthTypes) SetBandwidthType(v []*string) *DescribeAvailableResourceInfoResponseSupportResourcesSupportResourceBandwidthTypes {
	s.BandwidthType = v
	return s
}

type DescribeAvailableResourceInfoResponseImages struct {
	Image []*DescribeAvailableResourceInfoResponseImagesImage `json:"Image,omitempty" xml:"Image,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAvailableResourceInfoResponseImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseImages) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseImages) SetImage(v []*DescribeAvailableResourceInfoResponseImagesImage) *DescribeAvailableResourceInfoResponseImages {
	s.Image = v
	return s
}

type DescribeAvailableResourceInfoResponseImagesImage struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty" require:"true"`
	ImageSize *int    `json:"ImageSize,omitempty" xml:"ImageSize,omitempty" require:"true"`
}

func (s DescribeAvailableResourceInfoResponseImagesImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponseImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponseImagesImage) SetImageId(v string) *DescribeAvailableResourceInfoResponseImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseImagesImage) SetImageName(v string) *DescribeAvailableResourceInfoResponseImagesImage {
	s.ImageName = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponseImagesImage) SetImageSize(v int) *DescribeAvailableResourceInfoResponseImagesImage {
	s.ImageSize = &v
	return s
}

type DescribePrePaidInstanceStockRequest struct {
	Version        *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	EnsRegionId    *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	SystemDiskSize *int    `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" require:"true"`
	DataDiskSize   *int    `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" require:"true"`
	InstanceSpec   *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" require:"true"`
}

func (s DescribePrePaidInstanceStockRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePrePaidInstanceStockRequest) GoString() string {
	return s.String()
}

func (s *DescribePrePaidInstanceStockRequest) SetVersion(v string) *DescribePrePaidInstanceStockRequest {
	s.Version = &v
	return s
}

func (s *DescribePrePaidInstanceStockRequest) SetEnsRegionId(v string) *DescribePrePaidInstanceStockRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribePrePaidInstanceStockRequest) SetSystemDiskSize(v int) *DescribePrePaidInstanceStockRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribePrePaidInstanceStockRequest) SetDataDiskSize(v int) *DescribePrePaidInstanceStockRequest {
	s.DataDiskSize = &v
	return s
}

func (s *DescribePrePaidInstanceStockRequest) SetInstanceSpec(v string) *DescribePrePaidInstanceStockRequest {
	s.InstanceSpec = &v
	return s
}

type DescribePrePaidInstanceStockResponse struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DataDiskSize   *int    `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" require:"true"`
	EnsRegionId    *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	Cores          *int    `json:"Cores,omitempty" xml:"Cores,omitempty" require:"true"`
	Memory         *int    `json:"Memory,omitempty" xml:"Memory,omitempty" require:"true"`
	AvaliableCount *int    `json:"AvaliableCount,omitempty" xml:"AvaliableCount,omitempty" require:"true"`
	InstanceSpec   *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" require:"true"`
	SystemDiskSize *int    `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" require:"true"`
}

func (s DescribePrePaidInstanceStockResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePrePaidInstanceStockResponse) GoString() string {
	return s.String()
}

func (s *DescribePrePaidInstanceStockResponse) SetRequestId(v string) *DescribePrePaidInstanceStockResponse {
	s.RequestId = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponse) SetDataDiskSize(v int) *DescribePrePaidInstanceStockResponse {
	s.DataDiskSize = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponse) SetEnsRegionId(v string) *DescribePrePaidInstanceStockResponse {
	s.EnsRegionId = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponse) SetCores(v int) *DescribePrePaidInstanceStockResponse {
	s.Cores = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponse) SetMemory(v int) *DescribePrePaidInstanceStockResponse {
	s.Memory = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponse) SetAvaliableCount(v int) *DescribePrePaidInstanceStockResponse {
	s.AvaliableCount = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponse) SetInstanceSpec(v string) *DescribePrePaidInstanceStockResponse {
	s.InstanceSpec = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponse) SetSystemDiskSize(v int) *DescribePrePaidInstanceStockResponse {
	s.SystemDiskSize = &v
	return s
}

type UnassociateEipAddressRequest struct {
	Version              *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	EnsRegionId          *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	Eip                  *string `json:"Eip,omitempty" xml:"Eip,omitempty" require:"true"`
	InstanceIdInternetIp *string `json:"InstanceIdInternetIp,omitempty" xml:"InstanceIdInternetIp,omitempty" require:"true"`
}

func (s UnassociateEipAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s UnassociateEipAddressRequest) GoString() string {
	return s.String()
}

func (s *UnassociateEipAddressRequest) SetVersion(v string) *UnassociateEipAddressRequest {
	s.Version = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetEnsRegionId(v string) *UnassociateEipAddressRequest {
	s.EnsRegionId = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetEip(v string) *UnassociateEipAddressRequest {
	s.Eip = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetInstanceIdInternetIp(v string) *UnassociateEipAddressRequest {
	s.InstanceIdInternetIp = &v
	return s
}

type UnassociateEipAddressResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UnassociateEipAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s UnassociateEipAddressResponse) GoString() string {
	return s.String()
}

func (s *UnassociateEipAddressResponse) SetRequestId(v string) *UnassociateEipAddressResponse {
	s.RequestId = &v
	return s
}

type ReleaseEipAddressRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	Eips        *string `json:"Eips,omitempty" xml:"Eips,omitempty" require:"true"`
}

func (s ReleaseEipAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseEipAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleaseEipAddressRequest) SetVersion(v string) *ReleaseEipAddressRequest {
	s.Version = &v
	return s
}

func (s *ReleaseEipAddressRequest) SetEnsRegionId(v string) *ReleaseEipAddressRequest {
	s.EnsRegionId = &v
	return s
}

func (s *ReleaseEipAddressRequest) SetEips(v string) *ReleaseEipAddressRequest {
	s.Eips = &v
	return s
}

type ReleaseEipAddressResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ReleaseEipAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseEipAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleaseEipAddressResponse) SetRequestId(v string) *ReleaseEipAddressResponse {
	s.RequestId = &v
	return s
}

type DescribeEipAddressesRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	Eips        *string `json:"Eips,omitempty" xml:"Eips,omitempty"`
}

func (s DescribeEipAddressesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEipAddressesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesRequest) SetVersion(v string) *DescribeEipAddressesRequest {
	s.Version = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetEnsRegionId(v string) *DescribeEipAddressesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetEips(v string) *DescribeEipAddressesRequest {
	s.Eips = &v
	return s
}

type DescribeEipAddressesResponse struct {
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	EipAddresses *DescribeEipAddressesResponseEipAddresses `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty" require:"true" type:"Struct"`
}

func (s DescribeEipAddressesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEipAddressesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponse) SetRequestId(v string) *DescribeEipAddressesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeEipAddressesResponse) SetEipAddresses(v *DescribeEipAddressesResponseEipAddresses) *DescribeEipAddressesResponse {
	s.EipAddresses = v
	return s
}

type DescribeEipAddressesResponseEipAddresses struct {
	EipAddress []*DescribeEipAddressesResponseEipAddressesEipAddress `json:"EipAddress,omitempty" xml:"EipAddress,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeEipAddressesResponseEipAddresses) String() string {
	return tea.Prettify(s)
}

func (s DescribeEipAddressesResponseEipAddresses) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseEipAddresses) SetEipAddress(v []*DescribeEipAddressesResponseEipAddressesEipAddress) *DescribeEipAddressesResponseEipAddresses {
	s.EipAddress = v
	return s
}

type DescribeEipAddressesResponseEipAddressesEipAddress struct {
	Eip                  *string `json:"Eip,omitempty" xml:"Eip,omitempty" require:"true"`
	InstanceIdInternetIp *string `json:"InstanceIdInternetIp,omitempty" xml:"InstanceIdInternetIp,omitempty" require:"true"`
}

func (s DescribeEipAddressesResponseEipAddressesEipAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeEipAddressesResponseEipAddressesEipAddress) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseEipAddressesEipAddress) SetEip(v string) *DescribeEipAddressesResponseEipAddressesEipAddress {
	s.Eip = &v
	return s
}

func (s *DescribeEipAddressesResponseEipAddressesEipAddress) SetInstanceIdInternetIp(v string) *DescribeEipAddressesResponseEipAddressesEipAddress {
	s.InstanceIdInternetIp = &v
	return s
}

type AssociateEipAddressRequest struct {
	Version              *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	EnsRegionId          *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	Eip                  *string `json:"Eip,omitempty" xml:"Eip,omitempty" require:"true"`
	InstanceIdInternetIp *string `json:"InstanceIdInternetIp,omitempty" xml:"InstanceIdInternetIp,omitempty" require:"true"`
}

func (s AssociateEipAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateEipAddressRequest) GoString() string {
	return s.String()
}

func (s *AssociateEipAddressRequest) SetVersion(v string) *AssociateEipAddressRequest {
	s.Version = &v
	return s
}

func (s *AssociateEipAddressRequest) SetEnsRegionId(v string) *AssociateEipAddressRequest {
	s.EnsRegionId = &v
	return s
}

func (s *AssociateEipAddressRequest) SetEip(v string) *AssociateEipAddressRequest {
	s.Eip = &v
	return s
}

func (s *AssociateEipAddressRequest) SetInstanceIdInternetIp(v string) *AssociateEipAddressRequest {
	s.InstanceIdInternetIp = &v
	return s
}

type AssociateEipAddressResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AssociateEipAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateEipAddressResponse) GoString() string {
	return s.String()
}

func (s *AssociateEipAddressResponse) SetRequestId(v string) *AssociateEipAddressResponse {
	s.RequestId = &v
	return s
}

type AllocateEipAddressRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	Count       *int    `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
	MinCount    *int    `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
}

func (s AllocateEipAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateEipAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressRequest) SetVersion(v string) *AllocateEipAddressRequest {
	s.Version = &v
	return s
}

func (s *AllocateEipAddressRequest) SetEnsRegionId(v string) *AllocateEipAddressRequest {
	s.EnsRegionId = &v
	return s
}

func (s *AllocateEipAddressRequest) SetCount(v int) *AllocateEipAddressRequest {
	s.Count = &v
	return s
}

func (s *AllocateEipAddressRequest) SetMinCount(v int) *AllocateEipAddressRequest {
	s.MinCount = &v
	return s
}

type AllocateEipAddressResponse struct {
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizStatusCode *string                                 `json:"BizStatusCode,omitempty" xml:"BizStatusCode,omitempty" require:"true"`
	EipAddresses  *AllocateEipAddressResponseEipAddresses `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty" require:"true" type:"Struct"`
}

func (s AllocateEipAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocateEipAddressResponse) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressResponse) SetRequestId(v string) *AllocateEipAddressResponse {
	s.RequestId = &v
	return s
}

func (s *AllocateEipAddressResponse) SetBizStatusCode(v string) *AllocateEipAddressResponse {
	s.BizStatusCode = &v
	return s
}

func (s *AllocateEipAddressResponse) SetEipAddresses(v *AllocateEipAddressResponseEipAddresses) *AllocateEipAddressResponse {
	s.EipAddresses = v
	return s
}

type AllocateEipAddressResponseEipAddresses struct {
	EipAddress []*AllocateEipAddressResponseEipAddressesEipAddress `json:"EipAddress,omitempty" xml:"EipAddress,omitempty" require:"true" type:"Repeated"`
}

func (s AllocateEipAddressResponseEipAddresses) String() string {
	return tea.Prettify(s)
}

func (s AllocateEipAddressResponseEipAddresses) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressResponseEipAddresses) SetEipAddress(v []*AllocateEipAddressResponseEipAddressesEipAddress) *AllocateEipAddressResponseEipAddresses {
	s.EipAddress = v
	return s
}

type AllocateEipAddressResponseEipAddressesEipAddress struct {
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty" require:"true"`
}

func (s AllocateEipAddressResponseEipAddressesEipAddress) String() string {
	return tea.Prettify(s)
}

func (s AllocateEipAddressResponseEipAddressesEipAddress) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressResponseEipAddressesEipAddress) SetEip(v string) *AllocateEipAddressResponseEipAddressesEipAddress {
	s.Eip = &v
	return s
}

type ReleasePostPaidInstanceRequest struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s ReleasePostPaidInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleasePostPaidInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleasePostPaidInstanceRequest) SetVersion(v string) *ReleasePostPaidInstanceRequest {
	s.Version = &v
	return s
}

func (s *ReleasePostPaidInstanceRequest) SetInstanceId(v string) *ReleasePostPaidInstanceRequest {
	s.InstanceId = &v
	return s
}

type ReleasePostPaidInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ReleasePostPaidInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleasePostPaidInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleasePostPaidInstanceResponse) SetRequestId(v string) *ReleasePostPaidInstanceResponse {
	s.RequestId = &v
	return s
}

type ReleasePrePaidInstanceRequest struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s ReleasePrePaidInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleasePrePaidInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleasePrePaidInstanceRequest) SetVersion(v string) *ReleasePrePaidInstanceRequest {
	s.Version = &v
	return s
}

func (s *ReleasePrePaidInstanceRequest) SetInstanceId(v string) *ReleasePrePaidInstanceRequest {
	s.InstanceId = &v
	return s
}

type ReleasePrePaidInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ReleasePrePaidInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleasePrePaidInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleasePrePaidInstanceResponse) SetRequestId(v string) *ReleasePrePaidInstanceResponse {
	s.RequestId = &v
	return s
}

type AttachEnsInstancesRequest struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Scripts    *string `json:"Scripts,omitempty" xml:"Scripts,omitempty" require:"true"`
}

func (s AttachEnsInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachEnsInstancesRequest) GoString() string {
	return s.String()
}

func (s *AttachEnsInstancesRequest) SetVersion(v string) *AttachEnsInstancesRequest {
	s.Version = &v
	return s
}

func (s *AttachEnsInstancesRequest) SetInstanceId(v string) *AttachEnsInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachEnsInstancesRequest) SetScripts(v string) *AttachEnsInstancesRequest {
	s.Scripts = &v
	return s
}

type AttachEnsInstancesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AttachEnsInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachEnsInstancesResponse) GoString() string {
	return s.String()
}

func (s *AttachEnsInstancesResponse) SetRequestId(v string) *AttachEnsInstancesResponse {
	s.RequestId = &v
	return s
}

type DescribeReservedResourceRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s DescribeReservedResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceRequest) SetVersion(v string) *DescribeReservedResourceRequest {
	s.Version = &v
	return s
}

type DescribeReservedResourceResponse struct {
	RequestId        *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code             *int                                              `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Images           *DescribeReservedResourceResponseImages           `json:"Images,omitempty" xml:"Images,omitempty" require:"true" type:"Struct"`
	SupportResources *DescribeReservedResourceResponseSupportResources `json:"SupportResources,omitempty" xml:"SupportResources,omitempty" require:"true" type:"Struct"`
}

func (s DescribeReservedResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponse) SetRequestId(v string) *DescribeReservedResourceResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeReservedResourceResponse) SetCode(v int) *DescribeReservedResourceResponse {
	s.Code = &v
	return s
}

func (s *DescribeReservedResourceResponse) SetImages(v *DescribeReservedResourceResponseImages) *DescribeReservedResourceResponse {
	s.Images = v
	return s
}

func (s *DescribeReservedResourceResponse) SetSupportResources(v *DescribeReservedResourceResponseSupportResources) *DescribeReservedResourceResponse {
	s.SupportResources = v
	return s
}

type DescribeReservedResourceResponseImages struct {
	Image []*DescribeReservedResourceResponseImagesImage `json:"Image,omitempty" xml:"Image,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeReservedResourceResponseImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceResponseImages) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseImages) SetImage(v []*DescribeReservedResourceResponseImagesImage) *DescribeReservedResourceResponseImages {
	s.Image = v
	return s
}

type DescribeReservedResourceResponseImagesImage struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty" require:"true"`
}

func (s DescribeReservedResourceResponseImagesImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceResponseImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseImagesImage) SetImageId(v string) *DescribeReservedResourceResponseImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeReservedResourceResponseImagesImage) SetImageName(v string) *DescribeReservedResourceResponseImagesImage {
	s.ImageName = &v
	return s
}

type DescribeReservedResourceResponseSupportResources struct {
	SupportResource []*DescribeReservedResourceResponseSupportResourcesSupportResource `json:"SupportResource,omitempty" xml:"SupportResource,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeReservedResourceResponseSupportResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceResponseSupportResources) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseSupportResources) SetSupportResource(v []*DescribeReservedResourceResponseSupportResourcesSupportResource) *DescribeReservedResourceResponseSupportResources {
	s.SupportResource = v
	return s
}

type DescribeReservedResourceResponseSupportResourcesSupportResource struct {
	EnsRegionId           *string                                                                         `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	SupportResourcesCount *string                                                                         `json:"SupportResourcesCount,omitempty" xml:"SupportResourcesCount,omitempty" require:"true"`
	InstanceSpec          *string                                                                         `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" require:"true"`
	SystemDiskSizes       *DescribeReservedResourceResponseSupportResourcesSupportResourceSystemDiskSizes `json:"SystemDiskSizes,omitempty" xml:"SystemDiskSizes,omitempty" require:"true" type:"Struct"`
	DataDiskSizes         *DescribeReservedResourceResponseSupportResourcesSupportResourceDataDiskSizes   `json:"DataDiskSizes,omitempty" xml:"DataDiskSizes,omitempty" require:"true" type:"Struct"`
}

func (s DescribeReservedResourceResponseSupportResourcesSupportResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceResponseSupportResourcesSupportResource) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseSupportResourcesSupportResource) SetEnsRegionId(v string) *DescribeReservedResourceResponseSupportResourcesSupportResource {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeReservedResourceResponseSupportResourcesSupportResource) SetSupportResourcesCount(v string) *DescribeReservedResourceResponseSupportResourcesSupportResource {
	s.SupportResourcesCount = &v
	return s
}

func (s *DescribeReservedResourceResponseSupportResourcesSupportResource) SetInstanceSpec(v string) *DescribeReservedResourceResponseSupportResourcesSupportResource {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeReservedResourceResponseSupportResourcesSupportResource) SetSystemDiskSizes(v *DescribeReservedResourceResponseSupportResourcesSupportResourceSystemDiskSizes) *DescribeReservedResourceResponseSupportResourcesSupportResource {
	s.SystemDiskSizes = v
	return s
}

func (s *DescribeReservedResourceResponseSupportResourcesSupportResource) SetDataDiskSizes(v *DescribeReservedResourceResponseSupportResourcesSupportResourceDataDiskSizes) *DescribeReservedResourceResponseSupportResourcesSupportResource {
	s.DataDiskSizes = v
	return s
}

type DescribeReservedResourceResponseSupportResourcesSupportResourceSystemDiskSizes struct {
	// SystemDiskSize
	SystemDiskSize []*string `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeReservedResourceResponseSupportResourcesSupportResourceSystemDiskSizes) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceResponseSupportResourcesSupportResourceSystemDiskSizes) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseSupportResourcesSupportResourceSystemDiskSizes) SetSystemDiskSize(v []*string) *DescribeReservedResourceResponseSupportResourcesSupportResourceSystemDiskSizes {
	s.SystemDiskSize = v
	return s
}

type DescribeReservedResourceResponseSupportResourcesSupportResourceDataDiskSizes struct {
	// DataDiskSize
	DataDiskSize []*string `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeReservedResourceResponseSupportResourcesSupportResourceDataDiskSizes) String() string {
	return tea.Prettify(s)
}

func (s DescribeReservedResourceResponseSupportResourcesSupportResourceDataDiskSizes) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponseSupportResourcesSupportResourceDataDiskSizes) SetDataDiskSize(v []*string) *DescribeReservedResourceResponseSupportResourcesSupportResourceDataDiskSizes {
	s.DataDiskSize = v
	return s
}

type DescribeInstanceTypesRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s DescribeInstanceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesRequest) SetVersion(v string) *DescribeInstanceTypesRequest {
	s.Version = &v
	return s
}

type DescribeInstanceTypesResponse struct {
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code          *int                                        `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	InstanceTypes *DescribeInstanceTypesResponseInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" require:"true" type:"Struct"`
}

func (s DescribeInstanceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponse) SetRequestId(v string) *DescribeInstanceTypesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceTypesResponse) SetCode(v int) *DescribeInstanceTypesResponse {
	s.Code = &v
	return s
}

func (s *DescribeInstanceTypesResponse) SetInstanceTypes(v *DescribeInstanceTypesResponseInstanceTypes) *DescribeInstanceTypesResponse {
	s.InstanceTypes = v
	return s
}

type DescribeInstanceTypesResponseInstanceTypes struct {
	InstanceType []*DescribeInstanceTypesResponseInstanceTypesInstanceType `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstanceTypesResponseInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesResponseInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseInstanceTypes) SetInstanceType(v []*DescribeInstanceTypesResponseInstanceTypesInstanceType) *DescribeInstanceTypesResponseInstanceTypes {
	s.InstanceType = v
	return s
}

type DescribeInstanceTypesResponseInstanceTypesInstanceType struct {
	CpuCoreCount     *int    `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty" require:"true"`
	MemorySize       *int    `json:"MemorySize,omitempty" xml:"MemorySize,omitempty" require:"true"`
	InstanceTypeId   *string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty" require:"true"`
	InstanceTypeName *string `json:"InstanceTypeName,omitempty" xml:"InstanceTypeName,omitempty" require:"true"`
}

func (s DescribeInstanceTypesResponseInstanceTypesInstanceType) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesResponseInstanceTypesInstanceType) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseInstanceTypesInstanceType) SetCpuCoreCount(v int) *DescribeInstanceTypesResponseInstanceTypesInstanceType {
	s.CpuCoreCount = &v
	return s
}

func (s *DescribeInstanceTypesResponseInstanceTypesInstanceType) SetMemorySize(v int) *DescribeInstanceTypesResponseInstanceTypesInstanceType {
	s.MemorySize = &v
	return s
}

func (s *DescribeInstanceTypesResponseInstanceTypesInstanceType) SetInstanceTypeId(v string) *DescribeInstanceTypesResponseInstanceTypesInstanceType {
	s.InstanceTypeId = &v
	return s
}

func (s *DescribeInstanceTypesResponseInstanceTypesInstanceType) SetInstanceTypeName(v string) *DescribeInstanceTypesResponseInstanceTypesInstanceType {
	s.InstanceTypeName = &v
	return s
}

type CreateImageRequest struct {
	Product                *string `json:"product,omitempty" xml:"product,omitempty"`
	Version                *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	ImageName              *string `json:"ImageName,omitempty" xml:"ImageName,omitempty" require:"true"`
	DeleteAfterImageUpload *string `json:"DeleteAfterImageUpload,omitempty" xml:"DeleteAfterImageUpload,omitempty"`
}

func (s CreateImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageRequest) GoString() string {
	return s.String()
}

func (s *CreateImageRequest) SetProduct(v string) *CreateImageRequest {
	s.Product = &v
	return s
}

func (s *CreateImageRequest) SetVersion(v string) *CreateImageRequest {
	s.Version = &v
	return s
}

func (s *CreateImageRequest) SetInstanceId(v string) *CreateImageRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateImageRequest) SetImageName(v string) *CreateImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateImageRequest) SetDeleteAfterImageUpload(v string) *CreateImageRequest {
	s.DeleteAfterImageUpload = &v
	return s
}

type CreateImageResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
}

func (s CreateImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponse) GoString() string {
	return s.String()
}

func (s *CreateImageResponse) SetRequestId(v string) *CreateImageResponse {
	s.RequestId = &v
	return s
}

func (s *CreateImageResponse) SetCode(v int) *CreateImageResponse {
	s.Code = &v
	return s
}

type DescribeEnsNetSaleDistrictRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	NetLevelCode    *string `json:"NetLevelCode,omitempty" xml:"NetLevelCode,omitempty" require:"true"`
	NetDistrictCode *string `json:"NetDistrictCode,omitempty" xml:"NetDistrictCode,omitempty"`
}

func (s DescribeEnsNetSaleDistrictRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetSaleDistrictRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetSaleDistrictRequest) SetVersion(v string) *DescribeEnsNetSaleDistrictRequest {
	s.Version = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictRequest) SetNetLevelCode(v string) *DescribeEnsNetSaleDistrictRequest {
	s.NetLevelCode = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictRequest) SetNetDistrictCode(v string) *DescribeEnsNetSaleDistrictRequest {
	s.NetDistrictCode = &v
	return s
}

type DescribeEnsNetSaleDistrictResponse struct {
	RequestId       *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code            *int                                               `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	EnsNetDistricts *DescribeEnsNetSaleDistrictResponseEnsNetDistricts `json:"EnsNetDistricts,omitempty" xml:"EnsNetDistricts,omitempty" require:"true" type:"Struct"`
}

func (s DescribeEnsNetSaleDistrictResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetSaleDistrictResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetSaleDistrictResponse) SetRequestId(v string) *DescribeEnsNetSaleDistrictResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponse) SetCode(v int) *DescribeEnsNetSaleDistrictResponse {
	s.Code = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponse) SetEnsNetDistricts(v *DescribeEnsNetSaleDistrictResponseEnsNetDistricts) *DescribeEnsNetSaleDistrictResponse {
	s.EnsNetDistricts = v
	return s
}

type DescribeEnsNetSaleDistrictResponseEnsNetDistricts struct {
	EnsNetDistrict []*DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict `json:"EnsNetDistrict,omitempty" xml:"EnsNetDistrict,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeEnsNetSaleDistrictResponseEnsNetDistricts) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetSaleDistrictResponseEnsNetDistricts) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetSaleDistrictResponseEnsNetDistricts) SetEnsNetDistrict(v []*DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict) *DescribeEnsNetSaleDistrictResponseEnsNetDistricts {
	s.EnsNetDistrict = v
	return s
}

type DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict struct {
	NetDistrictCode       *string `json:"NetDistrictCode,omitempty" xml:"NetDistrictCode,omitempty" require:"true"`
	NetDistrictName       *string `json:"NetDistrictName,omitempty" xml:"NetDistrictName,omitempty" require:"true"`
	EnsRegionIdCount      *string `json:"EnsRegionIdCount,omitempty" xml:"EnsRegionIdCount,omitempty" require:"true"`
	NetDistrictLevel      *string `json:"NetDistrictLevel,omitempty" xml:"NetDistrictLevel,omitempty" require:"true"`
	NetDistrictFatherCode *string `json:"NetDistrictFatherCode,omitempty" xml:"NetDistrictFatherCode,omitempty" require:"true"`
	NetDistrictEnName     *string `json:"NetDistrictEnName,omitempty" xml:"NetDistrictEnName,omitempty" require:"true"`
	InstanceCount         *string `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty" require:"true"`
}

func (s DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict) SetNetDistrictCode(v string) *DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictCode = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict) SetNetDistrictName(v string) *DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictName = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict) SetEnsRegionIdCount(v string) *DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict {
	s.EnsRegionIdCount = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict) SetNetDistrictLevel(v string) *DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictLevel = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict) SetNetDistrictFatherCode(v string) *DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictFatherCode = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict) SetNetDistrictEnName(v string) *DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictEnName = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict) SetInstanceCount(v string) *DescribeEnsNetSaleDistrictResponseEnsNetDistrictsEnsNetDistrict {
	s.InstanceCount = &v
	return s
}

type DescribeEnsNetDistrictRequest struct {
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	NetLevelCode    *string `json:"NetLevelCode,omitempty" xml:"NetLevelCode,omitempty" require:"true"`
	NetDistrictCode *string `json:"NetDistrictCode,omitempty" xml:"NetDistrictCode,omitempty"`
}

func (s DescribeEnsNetDistrictRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetDistrictRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetDistrictRequest) SetVersion(v string) *DescribeEnsNetDistrictRequest {
	s.Version = &v
	return s
}

func (s *DescribeEnsNetDistrictRequest) SetNetLevelCode(v string) *DescribeEnsNetDistrictRequest {
	s.NetLevelCode = &v
	return s
}

func (s *DescribeEnsNetDistrictRequest) SetNetDistrictCode(v string) *DescribeEnsNetDistrictRequest {
	s.NetDistrictCode = &v
	return s
}

type DescribeEnsNetDistrictResponse struct {
	RequestId       *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code            *int                                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	EnsNetDistricts *DescribeEnsNetDistrictResponseEnsNetDistricts `json:"EnsNetDistricts,omitempty" xml:"EnsNetDistricts,omitempty" require:"true" type:"Struct"`
}

func (s DescribeEnsNetDistrictResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetDistrictResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetDistrictResponse) SetRequestId(v string) *DescribeEnsNetDistrictResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsNetDistrictResponse) SetCode(v int) *DescribeEnsNetDistrictResponse {
	s.Code = &v
	return s
}

func (s *DescribeEnsNetDistrictResponse) SetEnsNetDistricts(v *DescribeEnsNetDistrictResponseEnsNetDistricts) *DescribeEnsNetDistrictResponse {
	s.EnsNetDistricts = v
	return s
}

type DescribeEnsNetDistrictResponseEnsNetDistricts struct {
	EnsNetDistrict []*DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict `json:"EnsNetDistrict,omitempty" xml:"EnsNetDistrict,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeEnsNetDistrictResponseEnsNetDistricts) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetDistrictResponseEnsNetDistricts) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetDistrictResponseEnsNetDistricts) SetEnsNetDistrict(v []*DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict) *DescribeEnsNetDistrictResponseEnsNetDistricts {
	s.EnsNetDistrict = v
	return s
}

type DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict struct {
	NetDistrictCode       *string `json:"NetDistrictCode,omitempty" xml:"NetDistrictCode,omitempty" require:"true"`
	NetDistrictName       *string `json:"NetDistrictName,omitempty" xml:"NetDistrictName,omitempty" require:"true"`
	NetDistrictFatherCode *string `json:"NetDistrictFatherCode,omitempty" xml:"NetDistrictFatherCode,omitempty" require:"true"`
	EnsRegionIdCount      *string `json:"EnsRegionIdCount,omitempty" xml:"EnsRegionIdCount,omitempty" require:"true"`
	NetDistrictLevel      *string `json:"NetDistrictLevel,omitempty" xml:"NetDistrictLevel,omitempty" require:"true"`
	NetDistrictEnName     *string `json:"NetDistrictEnName,omitempty" xml:"NetDistrictEnName,omitempty" require:"true"`
}

func (s DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict) SetNetDistrictCode(v string) *DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictCode = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict) SetNetDistrictName(v string) *DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictName = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict) SetNetDistrictFatherCode(v string) *DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictFatherCode = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict) SetEnsRegionIdCount(v string) *DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict {
	s.EnsRegionIdCount = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict) SetNetDistrictLevel(v string) *DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictLevel = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict) SetNetDistrictEnName(v string) *DescribeEnsNetDistrictResponseEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictEnName = &v
	return s
}

type PreCreateEnsServiceRequest struct {
	Version                 *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	EnsServiceName          *string `json:"EnsServiceName,omitempty" xml:"EnsServiceName,omitempty" require:"true"`
	ImageId                 *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	InstanceSpec            *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" require:"true"`
	SystemDiskSize          *string `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" require:"true"`
	DataDiskSize            *string `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	BandwidthType           *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty" require:"true"`
	InstanceBandwithdLimit  *string `json:"InstanceBandwithdLimit,omitempty" xml:"InstanceBandwithdLimit,omitempty" require:"true"`
	Password                *string `json:"Password,omitempty" xml:"Password,omitempty"`
	KeyPairName             *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	UserData                *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	NetLevel                *string `json:"NetLevel,omitempty" xml:"NetLevel,omitempty" require:"true"`
	SchedulingStrategy      *string `json:"SchedulingStrategy,omitempty" xml:"SchedulingStrategy,omitempty" require:"true"`
	SchedulingPriceStrategy *string `json:"SchedulingPriceStrategy,omitempty" xml:"SchedulingPriceStrategy,omitempty"`
	BuyResourcesDetail      *string `json:"BuyResourcesDetail,omitempty" xml:"BuyResourcesDetail,omitempty" require:"true"`
}

func (s PreCreateEnsServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s PreCreateEnsServiceRequest) GoString() string {
	return s.String()
}

func (s *PreCreateEnsServiceRequest) SetVersion(v string) *PreCreateEnsServiceRequest {
	s.Version = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetEnsServiceName(v string) *PreCreateEnsServiceRequest {
	s.EnsServiceName = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetImageId(v string) *PreCreateEnsServiceRequest {
	s.ImageId = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetInstanceSpec(v string) *PreCreateEnsServiceRequest {
	s.InstanceSpec = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetSystemDiskSize(v string) *PreCreateEnsServiceRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetDataDiskSize(v string) *PreCreateEnsServiceRequest {
	s.DataDiskSize = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetBandwidthType(v string) *PreCreateEnsServiceRequest {
	s.BandwidthType = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetInstanceBandwithdLimit(v string) *PreCreateEnsServiceRequest {
	s.InstanceBandwithdLimit = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetPassword(v string) *PreCreateEnsServiceRequest {
	s.Password = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetKeyPairName(v string) *PreCreateEnsServiceRequest {
	s.KeyPairName = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetUserData(v string) *PreCreateEnsServiceRequest {
	s.UserData = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetNetLevel(v string) *PreCreateEnsServiceRequest {
	s.NetLevel = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetSchedulingStrategy(v string) *PreCreateEnsServiceRequest {
	s.SchedulingStrategy = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetSchedulingPriceStrategy(v string) *PreCreateEnsServiceRequest {
	s.SchedulingPriceStrategy = &v
	return s
}

func (s *PreCreateEnsServiceRequest) SetBuyResourcesDetail(v string) *PreCreateEnsServiceRequest {
	s.BuyResourcesDetail = &v
	return s
}

type PreCreateEnsServiceResponse struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code               *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	EnsServiceId       *string `json:"EnsServiceId,omitempty" xml:"EnsServiceId,omitempty" require:"true"`
	NetLevel           *string `json:"NetLevel,omitempty" xml:"NetLevel,omitempty" require:"true"`
	BuyResourcesDetail *string `json:"BuyResourcesDetail,omitempty" xml:"BuyResourcesDetail,omitempty" require:"true"`
}

func (s PreCreateEnsServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s PreCreateEnsServiceResponse) GoString() string {
	return s.String()
}

func (s *PreCreateEnsServiceResponse) SetRequestId(v string) *PreCreateEnsServiceResponse {
	s.RequestId = &v
	return s
}

func (s *PreCreateEnsServiceResponse) SetCode(v int) *PreCreateEnsServiceResponse {
	s.Code = &v
	return s
}

func (s *PreCreateEnsServiceResponse) SetEnsServiceId(v string) *PreCreateEnsServiceResponse {
	s.EnsServiceId = &v
	return s
}

func (s *PreCreateEnsServiceResponse) SetNetLevel(v string) *PreCreateEnsServiceResponse {
	s.NetLevel = &v
	return s
}

func (s *PreCreateEnsServiceResponse) SetBuyResourcesDetail(v string) *PreCreateEnsServiceResponse {
	s.BuyResourcesDetail = &v
	return s
}

type DescribeBandWithdChargeTypeRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s DescribeBandWithdChargeTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandWithdChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeBandWithdChargeTypeRequest) SetVersion(v string) *DescribeBandWithdChargeTypeRequest {
	s.Version = &v
	return s
}

type DescribeBandWithdChargeTypeResponse struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code               *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	BandWithTypeInfo   *string `json:"BandWithTypeInfo,omitempty" xml:"BandWithTypeInfo,omitempty" require:"true"`
	ChargeCycleInfo    *string `json:"ChargeCycleInfo,omitempty" xml:"ChargeCycleInfo,omitempty" require:"true"`
	ChargeContractType *string `json:"ChargeContractType,omitempty" xml:"ChargeContractType,omitempty" require:"true"`
}

func (s DescribeBandWithdChargeTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandWithdChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeBandWithdChargeTypeResponse) SetRequestId(v string) *DescribeBandWithdChargeTypeResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeBandWithdChargeTypeResponse) SetCode(v int) *DescribeBandWithdChargeTypeResponse {
	s.Code = &v
	return s
}

func (s *DescribeBandWithdChargeTypeResponse) SetBandWithTypeInfo(v string) *DescribeBandWithdChargeTypeResponse {
	s.BandWithTypeInfo = &v
	return s
}

func (s *DescribeBandWithdChargeTypeResponse) SetChargeCycleInfo(v string) *DescribeBandWithdChargeTypeResponse {
	s.ChargeCycleInfo = &v
	return s
}

func (s *DescribeBandWithdChargeTypeResponse) SetChargeContractType(v string) *DescribeBandWithdChargeTypeResponse {
	s.ChargeContractType = &v
	return s
}

type ModifyImageAttributeRequest struct {
	Product   *string `json:"product,omitempty" xml:"product,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty" require:"true"`
}

func (s ModifyImageAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeRequest) SetProduct(v string) *ModifyImageAttributeRequest {
	s.Product = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetVersion(v string) *ModifyImageAttributeRequest {
	s.Version = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetImageId(v string) *ModifyImageAttributeRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetImageName(v string) *ModifyImageAttributeRequest {
	s.ImageName = &v
	return s
}

type ModifyImageAttributeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
}

func (s ModifyImageAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeResponse) SetRequestId(v string) *ModifyImageAttributeResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyImageAttributeResponse) SetCode(v int) *ModifyImageAttributeResponse {
	s.Code = &v
	return s
}

type CreateEnsServiceRequest struct {
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	EnsServiceId *string `json:"EnsServiceId,omitempty" xml:"EnsServiceId,omitempty" require:"true"`
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty" require:"true"`
}

func (s CreateEnsServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEnsServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateEnsServiceRequest) SetVersion(v string) *CreateEnsServiceRequest {
	s.Version = &v
	return s
}

func (s *CreateEnsServiceRequest) SetEnsServiceId(v string) *CreateEnsServiceRequest {
	s.EnsServiceId = &v
	return s
}

func (s *CreateEnsServiceRequest) SetOrderType(v string) *CreateEnsServiceRequest {
	s.OrderType = &v
	return s
}

type CreateEnsServiceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
}

func (s CreateEnsServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEnsServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateEnsServiceResponse) SetRequestId(v string) *CreateEnsServiceResponse {
	s.RequestId = &v
	return s
}

func (s *CreateEnsServiceResponse) SetCode(v int) *CreateEnsServiceResponse {
	s.Code = &v
	return s
}

type DescribeEnsNetLevelRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s DescribeEnsNetLevelRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetLevelRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetLevelRequest) SetVersion(v string) *DescribeEnsNetLevelRequest {
	s.Version = &v
	return s
}

type DescribeEnsNetLevelResponse struct {
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code         *int                                     `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	EnsNetLevels *DescribeEnsNetLevelResponseEnsNetLevels `json:"EnsNetLevels,omitempty" xml:"EnsNetLevels,omitempty" require:"true" type:"Struct"`
}

func (s DescribeEnsNetLevelResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetLevelResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetLevelResponse) SetRequestId(v string) *DescribeEnsNetLevelResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsNetLevelResponse) SetCode(v int) *DescribeEnsNetLevelResponse {
	s.Code = &v
	return s
}

func (s *DescribeEnsNetLevelResponse) SetEnsNetLevels(v *DescribeEnsNetLevelResponseEnsNetLevels) *DescribeEnsNetLevelResponse {
	s.EnsNetLevels = v
	return s
}

type DescribeEnsNetLevelResponseEnsNetLevels struct {
	EnsNetLevel []*DescribeEnsNetLevelResponseEnsNetLevelsEnsNetLevel `json:"EnsNetLevel,omitempty" xml:"EnsNetLevel,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeEnsNetLevelResponseEnsNetLevels) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetLevelResponseEnsNetLevels) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetLevelResponseEnsNetLevels) SetEnsNetLevel(v []*DescribeEnsNetLevelResponseEnsNetLevelsEnsNetLevel) *DescribeEnsNetLevelResponseEnsNetLevels {
	s.EnsNetLevel = v
	return s
}

type DescribeEnsNetLevelResponseEnsNetLevelsEnsNetLevel struct {
	EnsNetLevelCode *string `json:"EnsNetLevelCode,omitempty" xml:"EnsNetLevelCode,omitempty" require:"true"`
}

func (s DescribeEnsNetLevelResponseEnsNetLevelsEnsNetLevel) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsNetLevelResponseEnsNetLevelsEnsNetLevel) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetLevelResponseEnsNetLevelsEnsNetLevel) SetEnsNetLevelCode(v string) *DescribeEnsNetLevelResponseEnsNetLevelsEnsNetLevel {
	s.EnsNetLevelCode = &v
	return s
}

type DescribeInstanceSpecRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s DescribeInstanceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecRequest) SetVersion(v string) *DescribeInstanceSpecRequest {
	s.Version = &v
	return s
}

type DescribeInstanceSpecResponse struct {
	RequestId         *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code              *int                                       `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	DataDiskMinSize   *int                                       `json:"DataDiskMinSize,omitempty" xml:"DataDiskMinSize,omitempty" require:"true"`
	DataDiskMaxSize   *int                                       `json:"DataDiskMaxSize,omitempty" xml:"DataDiskMaxSize,omitempty" require:"true"`
	SystemDiskMaxSize *int                                       `json:"SystemDiskMaxSize,omitempty" xml:"SystemDiskMaxSize,omitempty" require:"true"`
	BandwidthLimit    *int                                       `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty" require:"true"`
	InstanceSpecs     *DescribeInstanceSpecResponseInstanceSpecs `json:"InstanceSpecs,omitempty" xml:"InstanceSpecs,omitempty" require:"true" type:"Struct"`
}

func (s DescribeInstanceSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecResponse) SetRequestId(v string) *DescribeInstanceSpecResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSpecResponse) SetCode(v int) *DescribeInstanceSpecResponse {
	s.Code = &v
	return s
}

func (s *DescribeInstanceSpecResponse) SetDataDiskMinSize(v int) *DescribeInstanceSpecResponse {
	s.DataDiskMinSize = &v
	return s
}

func (s *DescribeInstanceSpecResponse) SetDataDiskMaxSize(v int) *DescribeInstanceSpecResponse {
	s.DataDiskMaxSize = &v
	return s
}

func (s *DescribeInstanceSpecResponse) SetSystemDiskMaxSize(v int) *DescribeInstanceSpecResponse {
	s.SystemDiskMaxSize = &v
	return s
}

func (s *DescribeInstanceSpecResponse) SetBandwidthLimit(v int) *DescribeInstanceSpecResponse {
	s.BandwidthLimit = &v
	return s
}

func (s *DescribeInstanceSpecResponse) SetInstanceSpecs(v *DescribeInstanceSpecResponseInstanceSpecs) *DescribeInstanceSpecResponse {
	s.InstanceSpecs = v
	return s
}

type DescribeInstanceSpecResponseInstanceSpecs struct {
	InstanceSpec []*DescribeInstanceSpecResponseInstanceSpecsInstanceSpec `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstanceSpecResponseInstanceSpecs) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecResponseInstanceSpecs) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecResponseInstanceSpecs) SetInstanceSpec(v []*DescribeInstanceSpecResponseInstanceSpecsInstanceSpec) *DescribeInstanceSpecResponseInstanceSpecs {
	s.InstanceSpec = v
	return s
}

type DescribeInstanceSpecResponseInstanceSpecsInstanceSpec struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" require:"true"`
	Core         *string `json:"Core,omitempty" xml:"Core,omitempty" require:"true"`
	Memory       *string `json:"Memory,omitempty" xml:"Memory,omitempty" require:"true"`
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
}

func (s DescribeInstanceSpecResponseInstanceSpecsInstanceSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecResponseInstanceSpecsInstanceSpec) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecResponseInstanceSpecsInstanceSpec) SetInstanceType(v string) *DescribeInstanceSpecResponseInstanceSpecsInstanceSpec {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceSpecResponseInstanceSpecsInstanceSpec) SetCore(v string) *DescribeInstanceSpecResponseInstanceSpecsInstanceSpec {
	s.Core = &v
	return s
}

func (s *DescribeInstanceSpecResponseInstanceSpecsInstanceSpec) SetMemory(v string) *DescribeInstanceSpecResponseInstanceSpecsInstanceSpec {
	s.Memory = &v
	return s
}

func (s *DescribeInstanceSpecResponseInstanceSpecsInstanceSpec) SetDisplayName(v string) *DescribeInstanceSpecResponseInstanceSpecsInstanceSpec {
	s.DisplayName = &v
	return s
}

type DescribeInstanceAutoRenewAttributeRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" require:"true"`
}

func (s DescribeInstanceAutoRenewAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeRequest) SetVersion(v string) *DescribeInstanceAutoRenewAttributeRequest {
	s.Version = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeRequest) SetInstanceIds(v string) *DescribeInstanceAutoRenewAttributeRequest {
	s.InstanceIds = &v
	return s
}

type DescribeInstanceAutoRenewAttributeResponse struct {
	RequestId               *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code                    *int                                                               `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	InstanceRenewAttributes *DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributes `json:"InstanceRenewAttributes,omitempty" xml:"InstanceRenewAttributes,omitempty" require:"true" type:"Struct"`
}

func (s DescribeInstanceAutoRenewAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeResponse) SetRequestId(v string) *DescribeInstanceAutoRenewAttributeResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponse) SetCode(v int) *DescribeInstanceAutoRenewAttributeResponse {
	s.Code = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponse) SetInstanceRenewAttributes(v *DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributes) *DescribeInstanceAutoRenewAttributeResponse {
	s.InstanceRenewAttributes = v
	return s
}

type DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributes struct {
	InstanceRenewAttribute []*DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributesInstanceRenewAttribute `json:"InstanceRenewAttribute,omitempty" xml:"InstanceRenewAttribute,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributes) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributes) SetInstanceRenewAttribute(v []*DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributesInstanceRenewAttribute) *DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributes {
	s.InstanceRenewAttribute = v
	return s
}

type DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributesInstanceRenewAttribute struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	AutoRenewal *bool   `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty" require:"true"`
	Duration    *string `json:"Duration,omitempty" xml:"Duration,omitempty" require:"true"`
}

func (s DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributesInstanceRenewAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributesInstanceRenewAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributesInstanceRenewAttribute) SetInstanceId(v string) *DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributesInstanceRenewAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributesInstanceRenewAttribute) SetAutoRenewal(v bool) *DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributesInstanceRenewAttribute {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributesInstanceRenewAttribute) SetDuration(v string) *DescribeInstanceAutoRenewAttributeResponseInstanceRenewAttributesInstanceRenewAttribute {
	s.Duration = &v
	return s
}

type DescribeAvailableResourceRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s DescribeAvailableResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) SetVersion(v string) *DescribeAvailableResourceRequest {
	s.Version = &v
	return s
}

type DescribeAvailableResourceResponse struct {
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code             *int                                               `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Images           *DescribeAvailableResourceResponseImages           `json:"Images,omitempty" xml:"Images,omitempty" require:"true" type:"Struct"`
	SupportResources *DescribeAvailableResourceResponseSupportResources `json:"SupportResources,omitempty" xml:"SupportResources,omitempty" require:"true" type:"Struct"`
}

func (s DescribeAvailableResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponse) SetRequestId(v string) *DescribeAvailableResourceResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourceResponse) SetCode(v int) *DescribeAvailableResourceResponse {
	s.Code = &v
	return s
}

func (s *DescribeAvailableResourceResponse) SetImages(v *DescribeAvailableResourceResponseImages) *DescribeAvailableResourceResponse {
	s.Images = v
	return s
}

func (s *DescribeAvailableResourceResponse) SetSupportResources(v *DescribeAvailableResourceResponseSupportResources) *DescribeAvailableResourceResponse {
	s.SupportResources = v
	return s
}

type DescribeAvailableResourceResponseImages struct {
	Image []*DescribeAvailableResourceResponseImagesImage `json:"Image,omitempty" xml:"Image,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseImages) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseImages) SetImage(v []*DescribeAvailableResourceResponseImagesImage) *DescribeAvailableResourceResponseImages {
	s.Image = v
	return s
}

type DescribeAvailableResourceResponseImagesImage struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty" require:"true"`
}

func (s DescribeAvailableResourceResponseImagesImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseImagesImage) SetImageId(v string) *DescribeAvailableResourceResponseImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeAvailableResourceResponseImagesImage) SetImageName(v string) *DescribeAvailableResourceResponseImagesImage {
	s.ImageName = &v
	return s
}

type DescribeAvailableResourceResponseSupportResources struct {
	SupportResource []*DescribeAvailableResourceResponseSupportResourcesSupportResource `json:"SupportResource,omitempty" xml:"SupportResource,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseSupportResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseSupportResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseSupportResources) SetSupportResource(v []*DescribeAvailableResourceResponseSupportResourcesSupportResource) *DescribeAvailableResourceResponseSupportResources {
	s.SupportResource = v
	return s
}

type DescribeAvailableResourceResponseSupportResourcesSupportResource struct {
	DataDiskSize          *string `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" require:"true"`
	EnsRegionId           *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	SupportResourcesCount *string `json:"SupportResourcesCount,omitempty" xml:"SupportResourcesCount,omitempty" require:"true"`
	InstanceSpec          *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" require:"true"`
	SystemDiskSize        *string `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" require:"true"`
}

func (s DescribeAvailableResourceResponseSupportResourcesSupportResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseSupportResourcesSupportResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseSupportResourcesSupportResource) SetDataDiskSize(v string) *DescribeAvailableResourceResponseSupportResourcesSupportResource {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeAvailableResourceResponseSupportResourcesSupportResource) SetEnsRegionId(v string) *DescribeAvailableResourceResponseSupportResourcesSupportResource {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeAvailableResourceResponseSupportResourcesSupportResource) SetSupportResourcesCount(v string) *DescribeAvailableResourceResponseSupportResourcesSupportResource {
	s.SupportResourcesCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseSupportResourcesSupportResource) SetInstanceSpec(v string) *DescribeAvailableResourceResponseSupportResourcesSupportResource {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeAvailableResourceResponseSupportResourcesSupportResource) SetSystemDiskSize(v string) *DescribeAvailableResourceResponseSupportResourcesSupportResource {
	s.SystemDiskSize = &v
	return s
}

type CreateInstanceRequest struct {
	InstanceType       *string                          `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" require:"true"`
	EnsRegionId        *string                          `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	Password           *string                          `json:"Password,omitempty" xml:"Password,omitempty"`
	Period             *string                          `json:"Period,omitempty" xml:"Period,omitempty" require:"true"`
	ImageId            *string                          `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	SystemDisk         *CreateInstanceRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	Quantity           *string                          `json:"Quantity,omitempty" xml:"Quantity,omitempty" require:"true"`
	DataDisk           []*CreateInstanceRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	InternetChargeType *string                          `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	AutoRenewPeriod    *string                          `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	AutoRenew          *string                          `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	IpType             *string                          `json:"IpType,omitempty" xml:"IpType,omitempty"`
	KeyPairName        *string                          `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	UserData           *string                          `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VSwitchId          *string                          `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	PrivateIpAddress   *string                          `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	PaymentType        *string                          `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	InstanceName       *string                          `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	HostName           *string                          `json:"HostName,omitempty" xml:"HostName,omitempty"`
	UniqueSuffix       *bool                            `json:"UniqueSuffix,omitempty" xml:"UniqueSuffix,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetEnsRegionId(v string) *CreateInstanceRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateInstanceRequest) SetPassword(v string) *CreateInstanceRequest {
	s.Password = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v string) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetImageId(v string) *CreateInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateInstanceRequest) SetSystemDisk(v *CreateInstanceRequestSystemDisk) *CreateInstanceRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateInstanceRequest) SetQuantity(v string) *CreateInstanceRequest {
	s.Quantity = &v
	return s
}

func (s *CreateInstanceRequest) SetDataDisk(v []*CreateInstanceRequestDataDisk) *CreateInstanceRequest {
	s.DataDisk = v
	return s
}

func (s *CreateInstanceRequest) SetInternetChargeType(v string) *CreateInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenewPeriod(v string) *CreateInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenew(v string) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetIpType(v string) *CreateInstanceRequest {
	s.IpType = &v
	return s
}

func (s *CreateInstanceRequest) SetKeyPairName(v string) *CreateInstanceRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateInstanceRequest) SetUserData(v string) *CreateInstanceRequest {
	s.UserData = &v
	return s
}

func (s *CreateInstanceRequest) SetVSwitchId(v string) *CreateInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequest) SetPrivateIpAddress(v string) *CreateInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateInstanceRequest) SetPaymentType(v string) *CreateInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetHostName(v string) *CreateInstanceRequest {
	s.HostName = &v
	return s
}

func (s *CreateInstanceRequest) SetUniqueSuffix(v bool) *CreateInstanceRequest {
	s.UniqueSuffix = &v
	return s
}

type CreateInstanceRequestSystemDisk struct {
	Size *string `json:"Size,omitempty" xml:"Size,omitempty" require:"true"`
}

func (s CreateInstanceRequestSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestSystemDisk) SetSize(v string) *CreateInstanceRequestSystemDisk {
	s.Size = &v
	return s
}

type CreateInstanceRequestDataDisk struct {
	Size *string `json:"Size,omitempty" xml:"Size,omitempty" require:"true"`
}

func (s CreateInstanceRequestDataDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestDataDisk) SetSize(v string) *CreateInstanceRequestDataDisk {
	s.Size = &v
	return s
}

type CreateInstanceResponse struct {
	RequestId   *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code        *int                               `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	InstanceIds *CreateInstanceResponseInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" require:"true" type:"Struct"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetRequestId(v string) *CreateInstanceResponse {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponse) SetCode(v int) *CreateInstanceResponse {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponse) SetInstanceIds(v *CreateInstanceResponseInstanceIds) *CreateInstanceResponse {
	s.InstanceIds = v
	return s
}

type CreateInstanceResponseInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true" type:"Repeated"`
}

func (s CreateInstanceResponseInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseInstanceIds) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseInstanceIds) SetInstanceId(v []*string) *CreateInstanceResponseInstanceIds {
	s.InstanceId = v
	return s
}

type ReInitDiskRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	DiskId  *string `json:"DiskId,omitempty" xml:"DiskId,omitempty" require:"true"`
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
}

func (s ReInitDiskRequest) String() string {
	return tea.Prettify(s)
}

func (s ReInitDiskRequest) GoString() string {
	return s.String()
}

func (s *ReInitDiskRequest) SetVersion(v string) *ReInitDiskRequest {
	s.Version = &v
	return s
}

func (s *ReInitDiskRequest) SetDiskId(v string) *ReInitDiskRequest {
	s.DiskId = &v
	return s
}

func (s *ReInitDiskRequest) SetImageId(v string) *ReInitDiskRequest {
	s.ImageId = &v
	return s
}

type ReInitDiskResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
}

func (s ReInitDiskResponse) String() string {
	return tea.Prettify(s)
}

func (s ReInitDiskResponse) GoString() string {
	return s.String()
}

func (s *ReInitDiskResponse) SetRequestId(v string) *ReInitDiskResponse {
	s.RequestId = &v
	return s
}

func (s *ReInitDiskResponse) SetCode(v int) *ReInitDiskResponse {
	s.Code = &v
	return s
}

type DescribeImageInfosRequest struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	OsType  *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
}

func (s DescribeImageInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageInfosRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageInfosRequest) SetVersion(v string) *DescribeImageInfosRequest {
	s.Version = &v
	return s
}

func (s *DescribeImageInfosRequest) SetOsType(v string) *DescribeImageInfosRequest {
	s.OsType = &v
	return s
}

type DescribeImageInfosResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                              `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Images    *DescribeImageInfosResponseImages `json:"Images,omitempty" xml:"Images,omitempty" require:"true" type:"Struct"`
}

func (s DescribeImageInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageInfosResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageInfosResponse) SetRequestId(v string) *DescribeImageInfosResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeImageInfosResponse) SetCode(v int) *DescribeImageInfosResponse {
	s.Code = &v
	return s
}

func (s *DescribeImageInfosResponse) SetImages(v *DescribeImageInfosResponseImages) *DescribeImageInfosResponse {
	s.Images = v
	return s
}

type DescribeImageInfosResponseImages struct {
	Image []*DescribeImageInfosResponseImagesImage `json:"Image,omitempty" xml:"Image,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeImageInfosResponseImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageInfosResponseImages) GoString() string {
	return s.String()
}

func (s *DescribeImageInfosResponseImages) SetImage(v []*DescribeImageInfosResponseImagesImage) *DescribeImageInfosResponseImages {
	s.Image = v
	return s
}

type DescribeImageInfosResponseImagesImage struct {
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty" require:"true"`
	OSType       *string `json:"OSType,omitempty" xml:"OSType,omitempty" require:"true"`
	OSName       *string `json:"OSName,omitempty" xml:"OSName,omitempty" require:"true"`
	ImageSize    *string `json:"ImageSize,omitempty" xml:"ImageSize,omitempty" require:"true"`
}

func (s DescribeImageInfosResponseImagesImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageInfosResponseImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeImageInfosResponseImagesImage) SetImageId(v string) *DescribeImageInfosResponseImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeImageInfosResponseImagesImage) SetDescription(v string) *DescribeImageInfosResponseImagesImage {
	s.Description = &v
	return s
}

func (s *DescribeImageInfosResponseImagesImage) SetImageVersion(v string) *DescribeImageInfosResponseImagesImage {
	s.ImageVersion = &v
	return s
}

func (s *DescribeImageInfosResponseImagesImage) SetOSType(v string) *DescribeImageInfosResponseImagesImage {
	s.OSType = &v
	return s
}

func (s *DescribeImageInfosResponseImagesImage) SetOSName(v string) *DescribeImageInfosResponseImagesImage {
	s.OSName = &v
	return s
}

func (s *DescribeImageInfosResponseImagesImage) SetImageSize(v string) *DescribeImageInfosResponseImagesImage {
	s.ImageSize = &v
	return s
}

type DescribeUserBandWidthDataRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Period      *string `json:"Period,omitempty" xml:"Period,omitempty" require:"true"`
	Isp         *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeUserBandWidthDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBandWidthDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserBandWidthDataRequest) SetVersion(v string) *DescribeUserBandWidthDataRequest {
	s.Version = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetEnsRegionId(v string) *DescribeUserBandWidthDataRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetInstanceId(v string) *DescribeUserBandWidthDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetStartTime(v string) *DescribeUserBandWidthDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetEndTime(v string) *DescribeUserBandWidthDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetPeriod(v string) *DescribeUserBandWidthDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetIsp(v string) *DescribeUserBandWidthDataRequest {
	s.Isp = &v
	return s
}

type DescribeUserBandWidthDataResponse struct {
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code        *int                                          `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	MonitorData *DescribeUserBandWidthDataResponseMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" require:"true" type:"Struct"`
}

func (s DescribeUserBandWidthDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBandWidthDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserBandWidthDataResponse) SetRequestId(v string) *DescribeUserBandWidthDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeUserBandWidthDataResponse) SetCode(v int) *DescribeUserBandWidthDataResponse {
	s.Code = &v
	return s
}

func (s *DescribeUserBandWidthDataResponse) SetMonitorData(v *DescribeUserBandWidthDataResponseMonitorData) *DescribeUserBandWidthDataResponse {
	s.MonitorData = v
	return s
}

type DescribeUserBandWidthDataResponseMonitorData struct {
	MaxDownBandWidth     *string                                                             `json:"MaxDownBandWidth,omitempty" xml:"MaxDownBandWidth,omitempty" require:"true"`
	MaxUpBandWidth       *string                                                             `json:"MaxUpBandWidth,omitempty" xml:"MaxUpBandWidth,omitempty" require:"true"`
	BandWidthMonitorData []*DescribeUserBandWidthDataResponseMonitorDataBandWidthMonitorData `json:"BandWidthMonitorData,omitempty" xml:"BandWidthMonitorData,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeUserBandWidthDataResponseMonitorData) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBandWidthDataResponseMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeUserBandWidthDataResponseMonitorData) SetMaxDownBandWidth(v string) *DescribeUserBandWidthDataResponseMonitorData {
	s.MaxDownBandWidth = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseMonitorData) SetMaxUpBandWidth(v string) *DescribeUserBandWidthDataResponseMonitorData {
	s.MaxUpBandWidth = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseMonitorData) SetBandWidthMonitorData(v []*DescribeUserBandWidthDataResponseMonitorDataBandWidthMonitorData) *DescribeUserBandWidthDataResponseMonitorData {
	s.BandWidthMonitorData = v
	return s
}

type DescribeUserBandWidthDataResponseMonitorDataBandWidthMonitorData struct {
	UpBandWidth   *int    `json:"UpBandWidth,omitempty" xml:"UpBandWidth,omitempty" require:"true"`
	DownBandWidth *int    `json:"DownBandWidth,omitempty" xml:"DownBandWidth,omitempty" require:"true"`
	InternetTX    *int    `json:"InternetTX,omitempty" xml:"InternetTX,omitempty" require:"true"`
	InternetRX    *int    `json:"InternetRX,omitempty" xml:"InternetRX,omitempty" require:"true"`
	TimeStamp     *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
}

func (s DescribeUserBandWidthDataResponseMonitorDataBandWidthMonitorData) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBandWidthDataResponseMonitorDataBandWidthMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeUserBandWidthDataResponseMonitorDataBandWidthMonitorData) SetUpBandWidth(v int) *DescribeUserBandWidthDataResponseMonitorDataBandWidthMonitorData {
	s.UpBandWidth = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseMonitorDataBandWidthMonitorData) SetDownBandWidth(v int) *DescribeUserBandWidthDataResponseMonitorDataBandWidthMonitorData {
	s.DownBandWidth = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseMonitorDataBandWidthMonitorData) SetInternetTX(v int) *DescribeUserBandWidthDataResponseMonitorDataBandWidthMonitorData {
	s.InternetTX = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseMonitorDataBandWidthMonitorData) SetInternetRX(v int) *DescribeUserBandWidthDataResponseMonitorDataBandWidthMonitorData {
	s.InternetRX = &v
	return s
}

func (s *DescribeUserBandWidthDataResponseMonitorDataBandWidthMonitorData) SetTimeStamp(v string) *DescribeUserBandWidthDataResponseMonitorDataBandWidthMonitorData {
	s.TimeStamp = &v
	return s
}

type RebootInstanceRequest struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	ForceStop  *string `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
}

func (s RebootInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebootInstanceRequest) SetVersion(v string) *RebootInstanceRequest {
	s.Version = &v
	return s
}

func (s *RebootInstanceRequest) SetInstanceId(v string) *RebootInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RebootInstanceRequest) SetForceStop(v string) *RebootInstanceRequest {
	s.ForceStop = &v
	return s
}

type RebootInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
}

func (s RebootInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootInstanceResponse) GoString() string {
	return s.String()
}

func (s *RebootInstanceResponse) SetRequestId(v string) *RebootInstanceResponse {
	s.RequestId = &v
	return s
}

func (s *RebootInstanceResponse) SetCode(v int) *RebootInstanceResponse {
	s.Code = &v
	return s
}

type DescribeEnsRegionsRequest struct {
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
}

func (s DescribeEnsRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionsRequest) SetVersion(v string) *DescribeEnsRegionsRequest {
	s.Version = &v
	return s
}

func (s *DescribeEnsRegionsRequest) SetEnsRegionId(v string) *DescribeEnsRegionsRequest {
	s.EnsRegionId = &v
	return s
}

type DescribeEnsRegionsResponse struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code       *int                                  `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	EnsRegions *DescribeEnsRegionsResponseEnsRegions `json:"EnsRegions,omitempty" xml:"EnsRegions,omitempty" require:"true" type:"Struct"`
}

func (s DescribeEnsRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionsResponse) SetRequestId(v string) *DescribeEnsRegionsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsRegionsResponse) SetCode(v int) *DescribeEnsRegionsResponse {
	s.Code = &v
	return s
}

func (s *DescribeEnsRegionsResponse) SetEnsRegions(v *DescribeEnsRegionsResponseEnsRegions) *DescribeEnsRegionsResponse {
	s.EnsRegions = v
	return s
}

type DescribeEnsRegionsResponseEnsRegions struct {
	EnsRegions []*DescribeEnsRegionsResponseEnsRegionsEnsRegions `json:"EnsRegions,omitempty" xml:"EnsRegions,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeEnsRegionsResponseEnsRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionsResponseEnsRegions) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionsResponseEnsRegions) SetEnsRegions(v []*DescribeEnsRegionsResponseEnsRegionsEnsRegions) *DescribeEnsRegionsResponseEnsRegions {
	s.EnsRegions = v
	return s
}

type DescribeEnsRegionsResponseEnsRegionsEnsRegions struct {
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	EnName      *string `json:"EnName,omitempty" xml:"EnName,omitempty" require:"true"`
	Area        *string `json:"Area,omitempty" xml:"Area,omitempty" require:"true"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty" require:"true"`
}

func (s DescribeEnsRegionsResponseEnsRegionsEnsRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnsRegionsResponseEnsRegionsEnsRegions) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionsResponseEnsRegionsEnsRegions) SetEnsRegionId(v string) *DescribeEnsRegionsResponseEnsRegionsEnsRegions {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsRegionsResponseEnsRegionsEnsRegions) SetName(v string) *DescribeEnsRegionsResponseEnsRegionsEnsRegions {
	s.Name = &v
	return s
}

func (s *DescribeEnsRegionsResponseEnsRegionsEnsRegions) SetEnName(v string) *DescribeEnsRegionsResponseEnsRegionsEnsRegions {
	s.EnName = &v
	return s
}

func (s *DescribeEnsRegionsResponseEnsRegionsEnsRegions) SetArea(v string) *DescribeEnsRegionsResponseEnsRegionsEnsRegions {
	s.Area = &v
	return s
}

func (s *DescribeEnsRegionsResponseEnsRegionsEnsRegions) SetProvince(v string) *DescribeEnsRegionsResponseEnsRegionsEnsRegions {
	s.Province = &v
	return s
}

type StartInstanceRequest struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s StartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) SetVersion(v string) *StartInstanceRequest {
	s.Version = &v
	return s
}

func (s *StartInstanceRequest) SetInstanceId(v string) *StartInstanceRequest {
	s.InstanceId = &v
	return s
}

type StartInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
}

func (s StartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceResponse) SetRequestId(v string) *StartInstanceResponse {
	s.RequestId = &v
	return s
}

func (s *StartInstanceResponse) SetCode(v int) *StartInstanceResponse {
	s.Code = &v
	return s
}

type DescribeInstanceMonitorDataRequest struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s DescribeInstanceMonitorDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataRequest) SetVersion(v string) *DescribeInstanceMonitorDataRequest {
	s.Version = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetInstanceId(v string) *DescribeInstanceMonitorDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetStartTime(v string) *DescribeInstanceMonitorDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetEndTime(v string) *DescribeInstanceMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetPeriod(v string) *DescribeInstanceMonitorDataRequest {
	s.Period = &v
	return s
}

type DescribeInstanceMonitorDataResponse struct {
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code        *int                                            `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	MonitorData *DescribeInstanceMonitorDataResponseMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" require:"true" type:"Struct"`
}

func (s DescribeInstanceMonitorDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataResponse) SetRequestId(v string) *DescribeInstanceMonitorDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponse) SetCode(v int) *DescribeInstanceMonitorDataResponse {
	s.Code = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponse) SetMonitorData(v *DescribeInstanceMonitorDataResponseMonitorData) *DescribeInstanceMonitorDataResponse {
	s.MonitorData = v
	return s
}

type DescribeInstanceMonitorDataResponseMonitorData struct {
	InstanceMonitorData []*DescribeInstanceMonitorDataResponseMonitorDataInstanceMonitorData `json:"InstanceMonitorData,omitempty" xml:"InstanceMonitorData,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstanceMonitorDataResponseMonitorData) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMonitorDataResponseMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataResponseMonitorData) SetInstanceMonitorData(v []*DescribeInstanceMonitorDataResponseMonitorDataInstanceMonitorData) *DescribeInstanceMonitorDataResponseMonitorData {
	s.InstanceMonitorData = v
	return s
}

type DescribeInstanceMonitorDataResponseMonitorDataInstanceMonitorData struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Memory     *string `json:"Memory,omitempty" xml:"Memory,omitempty" require:"true"`
	CPU        *string `json:"CPU,omitempty" xml:"CPU,omitempty" require:"true"`
}

func (s DescribeInstanceMonitorDataResponseMonitorDataInstanceMonitorData) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMonitorDataResponseMonitorDataInstanceMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataResponseMonitorDataInstanceMonitorData) SetInstanceId(v string) *DescribeInstanceMonitorDataResponseMonitorDataInstanceMonitorData {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseMonitorDataInstanceMonitorData) SetMemory(v string) *DescribeInstanceMonitorDataResponseMonitorDataInstanceMonitorData {
	s.Memory = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseMonitorDataInstanceMonitorData) SetCPU(v string) *DescribeInstanceMonitorDataResponseMonitorDataInstanceMonitorData {
	s.CPU = &v
	return s
}

type DescribeInstancesRequest struct {
	Version              *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	EnsRegionId          *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	EnsRegionIds         *string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty"`
	InstanceIds          *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	InstanceName         *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ImageId              *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	PageNumber           *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	OrderByParams        *string `json:"OrderByParams,omitempty" xml:"OrderByParams,omitempty"`
	EnsServiceId         *string `json:"EnsServiceId,omitempty" xml:"EnsServiceId,omitempty"`
	InstanceResourceType *string `json:"InstanceResourceType,omitempty" xml:"InstanceResourceType,omitempty"`
	SearchKey            *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetVersion(v string) *DescribeInstancesRequest {
	s.Version = &v
	return s
}

func (s *DescribeInstancesRequest) SetEnsRegionId(v string) *DescribeInstancesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceId(v string) *DescribeInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesRequest) SetEnsRegionIds(v string) *DescribeInstancesRequest {
	s.EnsRegionIds = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceIds(v string) *DescribeInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceName(v string) *DescribeInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesRequest) SetImageId(v string) *DescribeInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v int) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v string) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetStatus(v string) *DescribeInstancesRequest {
	s.Status = &v
	return s
}

func (s *DescribeInstancesRequest) SetOrderByParams(v string) *DescribeInstancesRequest {
	s.OrderByParams = &v
	return s
}

func (s *DescribeInstancesRequest) SetEnsServiceId(v string) *DescribeInstancesRequest {
	s.EnsServiceId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceResourceType(v string) *DescribeInstancesRequest {
	s.InstanceResourceType = &v
	return s
}

func (s *DescribeInstancesRequest) SetSearchKey(v string) *DescribeInstancesRequest {
	s.SearchKey = &v
	return s
}

type DescribeInstancesResponse struct {
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code       *int                                `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	TotalCount *int                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber *int                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Instances  *DescribeInstancesResponseInstances `json:"Instances,omitempty" xml:"Instances,omitempty" require:"true" type:"Struct"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetRequestId(v string) *DescribeInstancesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponse) SetCode(v int) *DescribeInstancesResponse {
	s.Code = &v
	return s
}

func (s *DescribeInstancesResponse) SetTotalCount(v int) *DescribeInstancesResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponse) SetPageNumber(v int) *DescribeInstancesResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesResponse) SetPageSize(v int) *DescribeInstancesResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponse) SetInstances(v *DescribeInstancesResponseInstances) *DescribeInstancesResponse {
	s.Instances = v
	return s
}

type DescribeInstancesResponseInstances struct {
	Instance []*DescribeInstancesResponseInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstancesResponseInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseInstances) SetInstance(v []*DescribeInstancesResponseInstancesInstance) *DescribeInstancesResponseInstances {
	s.Instance = v
	return s
}

type DescribeInstancesResponseInstancesInstance struct {
	InstanceId              *string                                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	InstanceName            *string                                                       `json:"InstanceName,omitempty" xml:"InstanceName,omitempty" require:"true"`
	EnsRegionId             *string                                                       `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty" require:"true"`
	Cpu                     *string                                                       `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Memory                  *int                                                          `json:"Memory,omitempty" xml:"Memory,omitempty" require:"true"`
	Disk                    *int                                                          `json:"Disk,omitempty" xml:"Disk,omitempty" require:"true"`
	InternetMaxBandwidthIn  *int                                                          `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty" require:"true"`
	InternetMaxBandwidthOut *int                                                          `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty" require:"true"`
	CreationTime            *string                                                       `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	Status                  *string                                                       `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	HostName                *string                                                       `json:"HostName,omitempty" xml:"HostName,omitempty" require:"true"`
	ImageId                 *string                                                       `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	ExpiredTime             *string                                                       `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty" require:"true"`
	InstanceResourceType    *string                                                       `json:"InstanceResourceType,omitempty" xml:"InstanceResourceType,omitempty" require:"true"`
	SpecName                *string                                                       `json:"SpecName,omitempty" xml:"SpecName,omitempty" require:"true"`
	OSName                  *string                                                       `json:"OSName,omitempty" xml:"OSName,omitempty" require:"true"`
	DataDisk                *DescribeInstancesResponseInstancesInstanceDataDisk           `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" require:"true" type:"Struct"`
	PublicIpAddresses       *DescribeInstancesResponseInstancesInstancePublicIpAddresses  `json:"PublicIpAddresses,omitempty" xml:"PublicIpAddresses,omitempty" require:"true" type:"Struct"`
	PrivateIpAddresses      *DescribeInstancesResponseInstancesInstancePrivateIpAddresses `json:"PrivateIpAddresses,omitempty" xml:"PrivateIpAddresses,omitempty" require:"true" type:"Struct"`
	SystemDisk              *DescribeInstancesResponseInstancesInstanceSystemDisk         `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" require:"true" type:"Struct"`
	SecurityGroupIds        *DescribeInstancesResponseInstancesInstanceSecurityGroupIds   `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" require:"true" type:"Struct"`
	InnerIpAddress          *DescribeInstancesResponseInstancesInstanceInnerIpAddress     `json:"InnerIpAddress,omitempty" xml:"InnerIpAddress,omitempty" require:"true" type:"Struct"`
	PublicIpAddress         *DescribeInstancesResponseInstancesInstancePublicIpAddress    `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty" require:"true" type:"Struct"`
}

func (s DescribeInstancesResponseInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseInstancesInstance) SetInstanceId(v string) *DescribeInstancesResponseInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetInstanceName(v string) *DescribeInstancesResponseInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetEnsRegionId(v string) *DescribeInstancesResponseInstancesInstance {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetCpu(v string) *DescribeInstancesResponseInstancesInstance {
	s.Cpu = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetMemory(v int) *DescribeInstancesResponseInstancesInstance {
	s.Memory = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetDisk(v int) *DescribeInstancesResponseInstancesInstance {
	s.Disk = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetInternetMaxBandwidthIn(v int) *DescribeInstancesResponseInstancesInstance {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetInternetMaxBandwidthOut(v int) *DescribeInstancesResponseInstancesInstance {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetCreationTime(v string) *DescribeInstancesResponseInstancesInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetStatus(v string) *DescribeInstancesResponseInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetHostName(v string) *DescribeInstancesResponseInstancesInstance {
	s.HostName = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetImageId(v string) *DescribeInstancesResponseInstancesInstance {
	s.ImageId = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetExpiredTime(v string) *DescribeInstancesResponseInstancesInstance {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetInstanceResourceType(v string) *DescribeInstancesResponseInstancesInstance {
	s.InstanceResourceType = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetSpecName(v string) *DescribeInstancesResponseInstancesInstance {
	s.SpecName = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetOSName(v string) *DescribeInstancesResponseInstancesInstance {
	s.OSName = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetDataDisk(v *DescribeInstancesResponseInstancesInstanceDataDisk) *DescribeInstancesResponseInstancesInstance {
	s.DataDisk = v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetPublicIpAddresses(v *DescribeInstancesResponseInstancesInstancePublicIpAddresses) *DescribeInstancesResponseInstancesInstance {
	s.PublicIpAddresses = v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetPrivateIpAddresses(v *DescribeInstancesResponseInstancesInstancePrivateIpAddresses) *DescribeInstancesResponseInstancesInstance {
	s.PrivateIpAddresses = v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetSystemDisk(v *DescribeInstancesResponseInstancesInstanceSystemDisk) *DescribeInstancesResponseInstancesInstance {
	s.SystemDisk = v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetSecurityGroupIds(v *DescribeInstancesResponseInstancesInstanceSecurityGroupIds) *DescribeInstancesResponseInstancesInstance {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetInnerIpAddress(v *DescribeInstancesResponseInstancesInstanceInnerIpAddress) *DescribeInstancesResponseInstancesInstance {
	s.InnerIpAddress = v
	return s
}

func (s *DescribeInstancesResponseInstancesInstance) SetPublicIpAddress(v *DescribeInstancesResponseInstancesInstancePublicIpAddress) *DescribeInstancesResponseInstancesInstance {
	s.PublicIpAddress = v
	return s
}

type DescribeInstancesResponseInstancesInstanceDataDisk struct {
	DataDisk []*DescribeInstancesResponseInstancesInstanceDataDiskDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstancesResponseInstancesInstanceDataDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseInstancesInstanceDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseInstancesInstanceDataDisk) SetDataDisk(v []*DescribeInstancesResponseInstancesInstanceDataDiskDataDisk) *DescribeInstancesResponseInstancesInstanceDataDisk {
	s.DataDisk = v
	return s
}

type DescribeInstancesResponseInstancesInstanceDataDiskDataDisk struct {
	DeviceType *string `json:"device_type,omitempty" xml:"device_type,omitempty" require:"true"`
	Storage    *int    `json:"storage,omitempty" xml:"storage,omitempty" require:"true"`
	Uuid       *string `json:"uuid,omitempty" xml:"uuid,omitempty" require:"true"`
	DiskType   *string `json:"disk_type,omitempty" xml:"disk_type,omitempty" require:"true"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	Category   *string `json:"Category,omitempty" xml:"Category,omitempty" require:"true"`
	Size       *int    `json:"Size,omitempty" xml:"Size,omitempty" require:"true"`
	DiskId     *string `json:"DiskId,omitempty" xml:"DiskId,omitempty" require:"true"`
	DiskName   *string `json:"DiskName,omitempty" xml:"DiskName,omitempty" require:"true"`
}

func (s DescribeInstancesResponseInstancesInstanceDataDiskDataDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseInstancesInstanceDataDiskDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk) SetDeviceType(v string) *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk {
	s.DeviceType = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk) SetStorage(v int) *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk {
	s.Storage = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk) SetUuid(v string) *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk {
	s.Uuid = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk) SetDiskType(v string) *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk {
	s.DiskType = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk) SetName(v string) *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk {
	s.Name = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk) SetCategory(v string) *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk {
	s.Category = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk) SetSize(v int) *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk {
	s.Size = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk) SetDiskId(v string) *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk {
	s.DiskId = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk) SetDiskName(v string) *DescribeInstancesResponseInstancesInstanceDataDiskDataDisk {
	s.DiskName = &v
	return s
}

type DescribeInstancesResponseInstancesInstancePublicIpAddresses struct {
	PublicIpAddress []*DescribeInstancesResponseInstancesInstancePublicIpAddressesPublicIpAddress `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstancesResponseInstancesInstancePublicIpAddresses) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseInstancesInstancePublicIpAddresses) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseInstancesInstancePublicIpAddresses) SetPublicIpAddress(v []*DescribeInstancesResponseInstancesInstancePublicIpAddressesPublicIpAddress) *DescribeInstancesResponseInstancesInstancePublicIpAddresses {
	s.PublicIpAddress = v
	return s
}

type DescribeInstancesResponseInstancesInstancePublicIpAddressesPublicIpAddress struct {
	Ip      *string `json:"Ip,omitempty" xml:"Ip,omitempty" require:"true"`
	GateWay *string `json:"GateWay,omitempty" xml:"GateWay,omitempty" require:"true"`
	Isp     *string `json:"Isp,omitempty" xml:"Isp,omitempty" require:"true"`
}

func (s DescribeInstancesResponseInstancesInstancePublicIpAddressesPublicIpAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseInstancesInstancePublicIpAddressesPublicIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseInstancesInstancePublicIpAddressesPublicIpAddress) SetIp(v string) *DescribeInstancesResponseInstancesInstancePublicIpAddressesPublicIpAddress {
	s.Ip = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstancePublicIpAddressesPublicIpAddress) SetGateWay(v string) *DescribeInstancesResponseInstancesInstancePublicIpAddressesPublicIpAddress {
	s.GateWay = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstancePublicIpAddressesPublicIpAddress) SetIsp(v string) *DescribeInstancesResponseInstancesInstancePublicIpAddressesPublicIpAddress {
	s.Isp = &v
	return s
}

type DescribeInstancesResponseInstancesInstancePrivateIpAddresses struct {
	PrivateIpAddress []*DescribeInstancesResponseInstancesInstancePrivateIpAddressesPrivateIpAddress `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstancesResponseInstancesInstancePrivateIpAddresses) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseInstancesInstancePrivateIpAddresses) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseInstancesInstancePrivateIpAddresses) SetPrivateIpAddress(v []*DescribeInstancesResponseInstancesInstancePrivateIpAddressesPrivateIpAddress) *DescribeInstancesResponseInstancesInstancePrivateIpAddresses {
	s.PrivateIpAddress = v
	return s
}

type DescribeInstancesResponseInstancesInstancePrivateIpAddressesPrivateIpAddress struct {
	Ip      *string `json:"Ip,omitempty" xml:"Ip,omitempty" require:"true"`
	GateWay *string `json:"GateWay,omitempty" xml:"GateWay,omitempty" require:"true"`
	Isp     *string `json:"Isp,omitempty" xml:"Isp,omitempty" require:"true"`
}

func (s DescribeInstancesResponseInstancesInstancePrivateIpAddressesPrivateIpAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseInstancesInstancePrivateIpAddressesPrivateIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseInstancesInstancePrivateIpAddressesPrivateIpAddress) SetIp(v string) *DescribeInstancesResponseInstancesInstancePrivateIpAddressesPrivateIpAddress {
	s.Ip = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstancePrivateIpAddressesPrivateIpAddress) SetGateWay(v string) *DescribeInstancesResponseInstancesInstancePrivateIpAddressesPrivateIpAddress {
	s.GateWay = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstancePrivateIpAddressesPrivateIpAddress) SetIsp(v string) *DescribeInstancesResponseInstancesInstancePrivateIpAddressesPrivateIpAddress {
	s.Isp = &v
	return s
}

type DescribeInstancesResponseInstancesInstanceSystemDisk struct {
	DeviceType *string `json:"device_type,omitempty" xml:"device_type,omitempty" require:"true"`
	Storage    *int    `json:"storage,omitempty" xml:"storage,omitempty" require:"true"`
	Uuid       *string `json:"uuid,omitempty" xml:"uuid,omitempty" require:"true"`
	DiskType   *string `json:"disk_type,omitempty" xml:"disk_type,omitempty" require:"true"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	Category   *string `json:"Category,omitempty" xml:"Category,omitempty" require:"true"`
	Size       *int    `json:"Size,omitempty" xml:"Size,omitempty" require:"true"`
	DiskId     *string `json:"DiskId,omitempty" xml:"DiskId,omitempty" require:"true"`
	DiskName   *string `json:"DiskName,omitempty" xml:"DiskName,omitempty" require:"true"`
}

func (s DescribeInstancesResponseInstancesInstanceSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseInstancesInstanceSystemDisk) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseInstancesInstanceSystemDisk) SetDeviceType(v string) *DescribeInstancesResponseInstancesInstanceSystemDisk {
	s.DeviceType = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstanceSystemDisk) SetStorage(v int) *DescribeInstancesResponseInstancesInstanceSystemDisk {
	s.Storage = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstanceSystemDisk) SetUuid(v string) *DescribeInstancesResponseInstancesInstanceSystemDisk {
	s.Uuid = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstanceSystemDisk) SetDiskType(v string) *DescribeInstancesResponseInstancesInstanceSystemDisk {
	s.DiskType = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstanceSystemDisk) SetName(v string) *DescribeInstancesResponseInstancesInstanceSystemDisk {
	s.Name = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstanceSystemDisk) SetCategory(v string) *DescribeInstancesResponseInstancesInstanceSystemDisk {
	s.Category = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstanceSystemDisk) SetSize(v int) *DescribeInstancesResponseInstancesInstanceSystemDisk {
	s.Size = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstanceSystemDisk) SetDiskId(v string) *DescribeInstancesResponseInstancesInstanceSystemDisk {
	s.DiskId = &v
	return s
}

func (s *DescribeInstancesResponseInstancesInstanceSystemDisk) SetDiskName(v string) *DescribeInstancesResponseInstancesInstanceSystemDisk {
	s.DiskName = &v
	return s
}

type DescribeInstancesResponseInstancesInstanceSecurityGroupIds struct {
	// SecurityGroupId
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstancesResponseInstancesInstanceSecurityGroupIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseInstancesInstanceSecurityGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseInstancesInstanceSecurityGroupIds) SetSecurityGroupId(v []*string) *DescribeInstancesResponseInstancesInstanceSecurityGroupIds {
	s.SecurityGroupId = v
	return s
}

type DescribeInstancesResponseInstancesInstanceInnerIpAddress struct {
	// IpAddress
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstancesResponseInstancesInstanceInnerIpAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseInstancesInstanceInnerIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseInstancesInstanceInnerIpAddress) SetIpAddress(v []*string) *DescribeInstancesResponseInstancesInstanceInnerIpAddress {
	s.IpAddress = v
	return s
}

type DescribeInstancesResponseInstancesInstancePublicIpAddress struct {
	// IpAddress
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstancesResponseInstancesInstancePublicIpAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseInstancesInstancePublicIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseInstancesInstancePublicIpAddress) SetIpAddress(v []*string) *DescribeInstancesResponseInstancesInstancePublicIpAddress {
	s.IpAddress = v
	return s
}

type DescribeImagesRequest struct {
	Product     *string `json:"product,omitempty" xml:"product,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ImageName   *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	PageNumber  *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagesRequest) SetProduct(v string) *DescribeImagesRequest {
	s.Product = &v
	return s
}

func (s *DescribeImagesRequest) SetVersion(v string) *DescribeImagesRequest {
	s.Version = &v
	return s
}

func (s *DescribeImagesRequest) SetEnsRegionId(v string) *DescribeImagesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeImagesRequest) SetImageId(v string) *DescribeImagesRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeImagesRequest) SetStatus(v string) *DescribeImagesRequest {
	s.Status = &v
	return s
}

func (s *DescribeImagesRequest) SetImageName(v string) *DescribeImagesRequest {
	s.ImageName = &v
	return s
}

func (s *DescribeImagesRequest) SetPageNumber(v string) *DescribeImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeImagesRequest) SetPageSize(v string) *DescribeImagesRequest {
	s.PageSize = &v
	return s
}

type DescribeImagesResponse struct {
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code       *int                          `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	TotalCount *int                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber *int                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                          `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Images     *DescribeImagesResponseImages `json:"Images,omitempty" xml:"Images,omitempty" require:"true" type:"Struct"`
}

func (s DescribeImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponse) SetRequestId(v string) *DescribeImagesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeImagesResponse) SetCode(v int) *DescribeImagesResponse {
	s.Code = &v
	return s
}

func (s *DescribeImagesResponse) SetTotalCount(v int) *DescribeImagesResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeImagesResponse) SetPageNumber(v int) *DescribeImagesResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeImagesResponse) SetPageSize(v int) *DescribeImagesResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeImagesResponse) SetImages(v *DescribeImagesResponseImages) *DescribeImagesResponse {
	s.Images = v
	return s
}

type DescribeImagesResponseImages struct {
	Image []*DescribeImagesResponseImagesImage `json:"Image,omitempty" xml:"Image,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeImagesResponseImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponseImages) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseImages) SetImage(v []*DescribeImagesResponseImagesImage) *DescribeImagesResponseImages {
	s.Image = v
	return s
}

type DescribeImagesResponseImagesImage struct {
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	ImageName       *string `json:"ImageName,omitempty" xml:"ImageName,omitempty" require:"true"`
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty" require:"true"`
	Platform        *string `json:"Platform,omitempty" xml:"Platform,omitempty" require:"true"`
	Architecture    *string `json:"Architecture,omitempty" xml:"Architecture,omitempty" require:"true"`
	CreationTime    *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	ImageSize       *string `json:"ImageSize,omitempty" xml:"ImageSize,omitempty" require:"true"`
}

func (s DescribeImagesResponseImagesImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponseImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseImagesImage) SetImageId(v string) *DescribeImagesResponseImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeImagesResponseImagesImage) SetImageName(v string) *DescribeImagesResponseImagesImage {
	s.ImageName = &v
	return s
}

func (s *DescribeImagesResponseImagesImage) SetImageOwnerAlias(v string) *DescribeImagesResponseImagesImage {
	s.ImageOwnerAlias = &v
	return s
}

func (s *DescribeImagesResponseImagesImage) SetPlatform(v string) *DescribeImagesResponseImagesImage {
	s.Platform = &v
	return s
}

func (s *DescribeImagesResponseImagesImage) SetArchitecture(v string) *DescribeImagesResponseImagesImage {
	s.Architecture = &v
	return s
}

func (s *DescribeImagesResponseImagesImage) SetCreationTime(v string) *DescribeImagesResponseImagesImage {
	s.CreationTime = &v
	return s
}

func (s *DescribeImagesResponseImagesImage) SetImageSize(v string) *DescribeImagesResponseImagesImage {
	s.ImageSize = &v
	return s
}

type StopInstanceRequest struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	ForceStop  *string `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
}

func (s StopInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) SetVersion(v string) *StopInstanceRequest {
	s.Version = &v
	return s
}

func (s *StopInstanceRequest) SetInstanceId(v string) *StopInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopInstanceRequest) SetForceStop(v string) *StopInstanceRequest {
	s.ForceStop = &v
	return s
}

type StopInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
}

func (s StopInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopInstanceResponse) SetRequestId(v string) *StopInstanceResponse {
	s.RequestId = &v
	return s
}

func (s *StopInstanceResponse) SetCode(v int) *StopInstanceResponse {
	s.Code = &v
	return s
}

type ModifyInstanceAttributeRequest struct {
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s ModifyInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequest) SetVersion(v string) *ModifyInstanceAttributeRequest {
	s.Version = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceId(v string) *ModifyInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetPassword(v string) *ModifyInstanceAttributeRequest {
	s.Password = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceName(v string) *ModifyInstanceAttributeRequest {
	s.InstanceName = &v
	return s
}

type ModifyInstanceAttributeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
}

func (s ModifyInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeResponse) SetRequestId(v string) *ModifyInstanceAttributeResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceAttributeResponse) SetCode(v int) *ModifyInstanceAttributeResponse {
	s.Code = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ens"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) DescribeEpnBandwitdhByInternetChargeTypeWithOptions(request *DescribeEpnBandwitdhByInternetChargeTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeEpnBandwitdhByInternetChargeTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeEpnBandwitdhByInternetChargeTypeResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeEpnBandwitdhByInternetChargeType"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEpnBandwitdhByInternetChargeType(request *DescribeEpnBandwitdhByInternetChargeTypeRequest) (_result *DescribeEpnBandwitdhByInternetChargeTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEpnBandwitdhByInternetChargeTypeResponse{}
	_body, _err := client.DescribeEpnBandwitdhByInternetChargeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEpnBandWidthDataWithOptions(request *DescribeEpnBandWidthDataRequest, runtime *util.RuntimeOptions) (_result *DescribeEpnBandWidthDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeEpnBandWidthDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeEpnBandWidthData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEpnBandWidthData(request *DescribeEpnBandWidthDataRequest) (_result *DescribeEpnBandWidthDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEpnBandWidthDataResponse{}
	_body, _err := client.DescribeEpnBandWidthDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEpnMeasurementDataWithOptions(request *DescribeEpnMeasurementDataRequest, runtime *util.RuntimeOptions) (_result *DescribeEpnMeasurementDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeEpnMeasurementDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeEpnMeasurementData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEpnMeasurementData(request *DescribeEpnMeasurementDataRequest) (_result *DescribeEpnMeasurementDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEpnMeasurementDataResponse{}
	_body, _err := client.DescribeEpnMeasurementDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNetworkInterfacesWithOptions(request *DescribeNetworkInterfacesRequest, runtime *util.RuntimeOptions) (_result *DescribeNetworkInterfacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeNetworkInterfacesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeNetworkInterfaces"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNetworkInterfaces(request *DescribeNetworkInterfacesRequest) (_result *DescribeNetworkInterfacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNetworkInterfacesResponse{}
	_body, _err := client.DescribeNetworkInterfacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEPInstanceWithOptions(request *CreateEPInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateEPInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateEPInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("CreateEPInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEPInstance(request *CreateEPInstanceRequest) (_result *CreateEPInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEPInstanceResponse{}
	_body, _err := client.CreateEPInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyImageSharePermissionWithOptions(request *ModifyImageSharePermissionRequest, runtime *util.RuntimeOptions) (_result *ModifyImageSharePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyImageSharePermissionResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyImageSharePermission"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyImageSharePermission(request *ModifyImageSharePermissionRequest) (_result *ModifyImageSharePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyImageSharePermissionResponse{}
	_body, _err := client.ModifyImageSharePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddNetworkInterfaceToInstanceWithOptions(request *AddNetworkInterfaceToInstanceRequest, runtime *util.RuntimeOptions) (_result *AddNetworkInterfaceToInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddNetworkInterfaceToInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("AddNetworkInterfaceToInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddNetworkInterfaceToInstance(request *AddNetworkInterfaceToInstanceRequest) (_result *AddNetworkInterfaceToInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddNetworkInterfaceToInstanceResponse{}
	_body, _err := client.AddNetworkInterfaceToInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageSharePermissionWithOptions(request *DescribeImageSharePermissionRequest, runtime *util.RuntimeOptions) (_result *DescribeImageSharePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeImageSharePermissionResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeImageSharePermission"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImageSharePermission(request *DescribeImageSharePermissionRequest) (_result *DescribeImageSharePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageSharePermissionResponse{}
	_body, _err := client.DescribeImageSharePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveVSwitchesFromEpnInstanceWithOptions(request *RemoveVSwitchesFromEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *RemoveVSwitchesFromEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RemoveVSwitchesFromEpnInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("RemoveVSwitchesFromEpnInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveVSwitchesFromEpnInstance(request *RemoveVSwitchesFromEpnInstanceRequest) (_result *RemoveVSwitchesFromEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveVSwitchesFromEpnInstanceResponse{}
	_body, _err := client.RemoveVSwitchesFromEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DistApplicationDataWithOptions(request *DistApplicationDataRequest, runtime *util.RuntimeOptions) (_result *DistApplicationDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DistApplicationDataResponse{}
	_body, _err := client.DoRequest(tea.String("DistApplicationData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DistApplicationData(request *DistApplicationDataRequest) (_result *DistApplicationDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DistApplicationDataResponse{}
	_body, _err := client.DistApplicationDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinVSwitchesToEpnInstanceWithOptions(request *JoinVSwitchesToEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *JoinVSwitchesToEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &JoinVSwitchesToEpnInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("JoinVSwitchesToEpnInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinVSwitchesToEpnInstance(request *JoinVSwitchesToEpnInstanceRequest) (_result *JoinVSwitchesToEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinVSwitchesToEpnInstanceResponse{}
	_body, _err := client.JoinVSwitchesToEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApplicationWithOptions(request *DescribeApplicationRequest, runtime *util.RuntimeOptions) (_result *DescribeApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeApplicationResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeApplication"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApplication(request *DescribeApplicationRequest) (_result *DescribeApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApplicationResponse{}
	_body, _err := client.DescribeApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataPushResultWithOptions(request *DescribeDataPushResultRequest, runtime *util.RuntimeOptions) (_result *DescribeDataPushResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDataPushResultResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDataPushResult"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataPushResult(request *DescribeDataPushResultRequest) (_result *DescribeDataPushResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataPushResultResponse{}
	_body, _err := client.DescribeDataPushResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushApplicationDataWithOptions(request *PushApplicationDataRequest, runtime *util.RuntimeOptions) (_result *PushApplicationDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PushApplicationDataResponse{}
	_body, _err := client.DoRequest(tea.String("PushApplicationData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushApplicationData(request *PushApplicationDataRequest) (_result *PushApplicationDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushApplicationDataResponse{}
	_body, _err := client.PushApplicationDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeApplicationWithOptions(request *UpgradeApplicationRequest, runtime *util.RuntimeOptions) (_result *UpgradeApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpgradeApplicationResponse{}
	_body, _err := client.DoRequest(tea.String("UpgradeApplication"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeApplication(request *UpgradeApplicationRequest) (_result *UpgradeApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeApplicationResponse{}
	_body, _err := client.UpgradeApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RescaleApplicationWithOptions(request *RescaleApplicationRequest, runtime *util.RuntimeOptions) (_result *RescaleApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RescaleApplicationResponse{}
	_body, _err := client.DoRequest(tea.String("RescaleApplication"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RescaleApplication(request *RescaleApplicationRequest) (_result *RescaleApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RescaleApplicationResponse{}
	_body, _err := client.RescaleApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEpnInstanceWithOptions(request *DeleteEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteEpnInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteEpnInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEpnInstance(request *DeleteEpnInstanceRequest) (_result *DeleteEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEpnInstanceResponse{}
	_body, _err := client.DeleteEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplicationsWithOptions(request *ListApplicationsRequest, runtime *util.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListApplicationsResponse{}
	_body, _err := client.DoRequest(tea.String("ListApplications"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplications(request *ListApplicationsRequest) (_result *ListApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationsResponse{}
	_body, _err := client.ListApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServcieScheduleWithOptions(request *DescribeServcieScheduleRequest, runtime *util.RuntimeOptions) (_result *DescribeServcieScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeServcieScheduleResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeServcieSchedule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServcieSchedule(request *DescribeServcieScheduleRequest) (_result *DescribeServcieScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServcieScheduleResponse{}
	_body, _err := client.DescribeServcieScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, runtime *util.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteApplication"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApplication(request *DeleteApplicationRequest) (_result *DeleteApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DeleteApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyEpnInstanceWithOptions(request *ModifyEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *ModifyEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyEpnInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyEpnInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyEpnInstance(request *ModifyEpnInstanceRequest) (_result *ModifyEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyEpnInstanceResponse{}
	_body, _err := client.ModifyEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackApplicationWithOptions(request *RollbackApplicationRequest, runtime *util.RuntimeOptions) (_result *RollbackApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RollbackApplicationResponse{}
	_body, _err := client.DoRequest(tea.String("RollbackApplication"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackApplication(request *RollbackApplicationRequest) (_result *RollbackApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackApplicationResponse{}
	_body, _err := client.RollbackApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEpnInstanceAttributeWithOptions(request *DescribeEpnInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeEpnInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeEpnInstanceAttributeResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeEpnInstanceAttribute"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEpnInstanceAttribute(request *DescribeEpnInstanceAttributeRequest) (_result *DescribeEpnInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEpnInstanceAttributeResponse{}
	_body, _err := client.DescribeEpnInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunServiceScheduleWithOptions(request *RunServiceScheduleRequest, runtime *util.RuntimeOptions) (_result *RunServiceScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RunServiceScheduleResponse{}
	_body, _err := client.DoRequest(tea.String("RunServiceSchedule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunServiceSchedule(request *RunServiceScheduleRequest) (_result *RunServiceScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunServiceScheduleResponse{}
	_body, _err := client.RunServiceScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApplicationWithOptions(request *CreateApplicationRequest, runtime *util.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.DoRequest(tea.String("CreateApplication"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApplication(request *CreateApplicationRequest) (_result *CreateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CreateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEpnInstanceWithOptions(request *CreateEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateEpnInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("CreateEpnInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEpnInstance(request *CreateEpnInstanceRequest) (_result *CreateEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEpnInstanceResponse{}
	_body, _err := client.CreateEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopEpnInstanceWithOptions(request *StopEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *StopEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopEpnInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("StopEpnInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopEpnInstance(request *StopEpnInstanceRequest) (_result *StopEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopEpnInstanceResponse{}
	_body, _err := client.StopEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataDistResultWithOptions(request *DescribeDataDistResultRequest, runtime *util.RuntimeOptions) (_result *DescribeDataDistResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDataDistResultResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDataDistResult"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataDistResult(request *DescribeDataDistResultRequest) (_result *DescribeDataDistResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataDistResultResponse{}
	_body, _err := client.DescribeDataDistResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEpnInstancesWithOptions(request *DescribeEpnInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeEpnInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeEpnInstancesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeEpnInstances"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEpnInstances(request *DescribeEpnInstancesRequest) (_result *DescribeEpnInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEpnInstancesResponse{}
	_body, _err := client.DescribeEpnInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemovePublicIpsFromEpnInstanceWithOptions(request *RemovePublicIpsFromEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *RemovePublicIpsFromEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RemovePublicIpsFromEpnInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("RemovePublicIpsFromEpnInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemovePublicIpsFromEpnInstance(request *RemovePublicIpsFromEpnInstanceRequest) (_result *RemovePublicIpsFromEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemovePublicIpsFromEpnInstanceResponse{}
	_body, _err := client.RemovePublicIpsFromEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinPublicIpsToEpnInstanceWithOptions(request *JoinPublicIpsToEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *JoinPublicIpsToEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &JoinPublicIpsToEpnInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("JoinPublicIpsToEpnInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinPublicIpsToEpnInstance(request *JoinPublicIpsToEpnInstanceRequest) (_result *JoinPublicIpsToEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinPublicIpsToEpnInstanceResponse{}
	_body, _err := client.JoinPublicIpsToEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApplicationResourceSummaryWithOptions(request *DescribeApplicationResourceSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeApplicationResourceSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeApplicationResourceSummaryResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeApplicationResourceSummary"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApplicationResourceSummary(request *DescribeApplicationResourceSummaryRequest) (_result *DescribeApplicationResourceSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApplicationResourceSummaryResponse{}
	_body, _err := client.DescribeApplicationResourceSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartEpnInstanceWithOptions(request *StartEpnInstanceRequest, runtime *util.RuntimeOptions) (_result *StartEpnInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartEpnInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("StartEpnInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartEpnInstance(request *StartEpnInstanceRequest) (_result *StartEpnInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartEpnInstanceResponse{}
	_body, _err := client.StartEpnInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExportImageInfoWithOptions(request *DescribeExportImageInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeExportImageInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeExportImageInfoResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeExportImageInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExportImageInfo(request *DescribeExportImageInfoRequest) (_result *DescribeExportImageInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExportImageInfoResponse{}
	_body, _err := client.DescribeExportImageInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVSwitchesWithOptions(request *DescribeVSwitchesRequest, runtime *util.RuntimeOptions) (_result *DescribeVSwitchesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVSwitches"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVSwitches(request *DescribeVSwitchesRequest) (_result *DescribeVSwitchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.DescribeVSwitchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVSwitchWithOptions(request *DeleteVSwitchRequest, runtime *util.RuntimeOptions) (_result *DeleteVSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteVSwitchResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteVSwitch"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVSwitch(request *DeleteVSwitchRequest) (_result *DeleteVSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVSwitchResponse{}
	_body, _err := client.DeleteVSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVSwitchWithOptions(request *CreateVSwitchRequest, runtime *util.RuntimeOptions) (_result *CreateVSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateVSwitchResponse{}
	_body, _err := client.DoRequest(tea.String("CreateVSwitch"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVSwitch(request *CreateVSwitchRequest) (_result *CreateVSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVSwitchResponse{}
	_body, _err := client.CreateVSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExportImageStatusWithOptions(request *DescribeExportImageStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeExportImageStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeExportImageStatusResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeExportImageStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExportImageStatus(request *DescribeExportImageStatusRequest) (_result *DescribeExportImageStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExportImageStatusResponse{}
	_body, _err := client.DescribeExportImageStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportImageWithOptions(request *ExportImageRequest, runtime *util.RuntimeOptions) (_result *ExportImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ExportImageResponse{}
	_body, _err := client.DoRequest(tea.String("ExportImage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportImage(request *ExportImageRequest) (_result *ExportImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportImageResponse{}
	_body, _err := client.ExportImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportKeyPairWithOptions(request *ImportKeyPairRequest, runtime *util.RuntimeOptions) (_result *ImportKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ImportKeyPairResponse{}
	_body, _err := client.DoRequest(tea.String("ImportKeyPair"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportKeyPair(request *ImportKeyPairRequest) (_result *ImportKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportKeyPairResponse{}
	_body, _err := client.ImportKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeKeyPairsWithOptions(request *DescribeKeyPairsRequest, runtime *util.RuntimeOptions) (_result *DescribeKeyPairsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeKeyPairsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeKeyPairs"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeKeyPairs(request *DescribeKeyPairsRequest) (_result *DescribeKeyPairsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeKeyPairsResponse{}
	_body, _err := client.DescribeKeyPairsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteKeyPairsWithOptions(request *DeleteKeyPairsRequest, runtime *util.RuntimeOptions) (_result *DeleteKeyPairsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteKeyPairsResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteKeyPairs"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteKeyPairs(request *DeleteKeyPairsRequest) (_result *DeleteKeyPairsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteKeyPairsResponse{}
	_body, _err := client.DeleteKeyPairsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateKeyPairWithOptions(request *CreateKeyPairRequest, runtime *util.RuntimeOptions) (_result *CreateKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateKeyPairResponse{}
	_body, _err := client.DoRequest(tea.String("CreateKeyPair"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateKeyPair(request *CreateKeyPairRequest) (_result *CreateKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateKeyPairResponse{}
	_body, _err := client.CreateKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportBillDetailDataWithOptions(request *ExportBillDetailDataRequest, runtime *util.RuntimeOptions) (_result *ExportBillDetailDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ExportBillDetailDataResponse{}
	_body, _err := client.DoRequest(tea.String("ExportBillDetailData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportBillDetailData(request *ExportBillDetailDataRequest) (_result *ExportBillDetailDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportBillDetailDataResponse{}
	_body, _err := client.ExportBillDetailDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnsRegionIdResourceWithOptions(request *DescribeEnsRegionIdResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeEnsRegionIdResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeEnsRegionIdResourceResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeEnsRegionIdResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnsRegionIdResource(request *DescribeEnsRegionIdResourceRequest) (_result *DescribeEnsRegionIdResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnsRegionIdResourceResponse{}
	_body, _err := client.DescribeEnsRegionIdResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBandwitdhByInternetChargeTypeWithOptions(request *DescribeBandwitdhByInternetChargeTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeBandwitdhByInternetChargeTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeBandwitdhByInternetChargeTypeResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeBandwitdhByInternetChargeType"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBandwitdhByInternetChargeType(request *DescribeBandwitdhByInternetChargeTypeRequest) (_result *DescribeBandwitdhByInternetChargeTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBandwitdhByInternetChargeTypeResponse{}
	_body, _err := client.DescribeBandwitdhByInternetChargeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthorizeSecurityGroupWithOptions(request *AuthorizeSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *AuthorizeSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AuthorizeSecurityGroupResponse{}
	_body, _err := client.DoRequest(tea.String("AuthorizeSecurityGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthorizeSecurityGroup(request *AuthorizeSecurityGroupRequest) (_result *AuthorizeSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeSecurityGroupResponse{}
	_body, _err := client.AuthorizeSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeSecurityGroupWithOptions(request *RevokeSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *RevokeSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RevokeSecurityGroupResponse{}
	_body, _err := client.DoRequest(tea.String("RevokeSecurityGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeSecurityGroup(request *RevokeSecurityGroupRequest) (_result *RevokeSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeSecurityGroupResponse{}
	_body, _err := client.RevokeSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSecurityGroupWithOptions(request *DeleteSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteSecurityGroupResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteSecurityGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) (_result *DeleteSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecurityGroupResponse{}
	_body, _err := client.DeleteSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityGroupAttributeWithOptions(request *DescribeSecurityGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSecurityGroupAttributeResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSecurityGroupAttribute"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityGroupAttribute(request *DescribeSecurityGroupAttributeRequest) (_result *DescribeSecurityGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityGroupAttributeResponse{}
	_body, _err := client.DescribeSecurityGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LeaveSecurityGroupWithOptions(request *LeaveSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *LeaveSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &LeaveSecurityGroupResponse{}
	_body, _err := client.DoRequest(tea.String("LeaveSecurityGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LeaveSecurityGroup(request *LeaveSecurityGroupRequest) (_result *LeaveSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LeaveSecurityGroupResponse{}
	_body, _err := client.LeaveSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinSecurityGroupWithOptions(request *JoinSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *JoinSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &JoinSecurityGroupResponse{}
	_body, _err := client.DoRequest(tea.String("JoinSecurityGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinSecurityGroup(request *JoinSecurityGroupRequest) (_result *JoinSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinSecurityGroupResponse{}
	_body, _err := client.JoinSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthorizeSecurityGroupEgressWithOptions(request *AuthorizeSecurityGroupEgressRequest, runtime *util.RuntimeOptions) (_result *AuthorizeSecurityGroupEgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AuthorizeSecurityGroupEgressResponse{}
	_body, _err := client.DoRequest(tea.String("AuthorizeSecurityGroupEgress"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthorizeSecurityGroupEgress(request *AuthorizeSecurityGroupEgressRequest) (_result *AuthorizeSecurityGroupEgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeSecurityGroupEgressResponse{}
	_body, _err := client.AuthorizeSecurityGroupEgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeSecurityGroupEgressWithOptions(request *RevokeSecurityGroupEgressRequest, runtime *util.RuntimeOptions) (_result *RevokeSecurityGroupEgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RevokeSecurityGroupEgressResponse{}
	_body, _err := client.DoRequest(tea.String("RevokeSecurityGroupEgress"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeSecurityGroupEgress(request *RevokeSecurityGroupEgressRequest) (_result *RevokeSecurityGroupEgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeSecurityGroupEgressResponse{}
	_body, _err := client.RevokeSecurityGroupEgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityGroupsWithOptions(request *DescribeSecurityGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSecurityGroupsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSecurityGroups"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (_result *DescribeSecurityGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityGroupsResponse{}
	_body, _err := client.DescribeSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSecurityGroupWithOptions(request *CreateSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *CreateSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateSecurityGroupResponse{}
	_body, _err := client.DoRequest(tea.String("CreateSecurityGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) (_result *CreateSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSecurityGroupResponse{}
	_body, _err := client.CreateSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnsRegionIdIpv6InfoWithOptions(request *DescribeEnsRegionIdIpv6InfoRequest, runtime *util.RuntimeOptions) (_result *DescribeEnsRegionIdIpv6InfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeEnsRegionIdIpv6InfoResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeEnsRegionIdIpv6Info"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnsRegionIdIpv6Info(request *DescribeEnsRegionIdIpv6InfoRequest) (_result *DescribeEnsRegionIdIpv6InfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnsRegionIdIpv6InfoResponse{}
	_body, _err := client.DescribeEnsRegionIdIpv6InfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCreatePrePaidInstanceResultWithOptions(request *DescribeCreatePrePaidInstanceResultRequest, runtime *util.RuntimeOptions) (_result *DescribeCreatePrePaidInstanceResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCreatePrePaidInstanceResultResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCreatePrePaidInstanceResult"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCreatePrePaidInstanceResult(request *DescribeCreatePrePaidInstanceResultRequest) (_result *DescribeCreatePrePaidInstanceResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCreatePrePaidInstanceResultResponse{}
	_body, _err := client.DescribeCreatePrePaidInstanceResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePriceWithOptions(request *DescribePriceRequest, runtime *util.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribePriceResponse{}
	_body, _err := client.DoRequest(tea.String("DescribePrice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePrice(request *DescribePriceRequest) (_result *DescribePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePriceResponse{}
	_body, _err := client.DescribePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportMeasurementDataWithOptions(request *ExportMeasurementDataRequest, runtime *util.RuntimeOptions) (_result *ExportMeasurementDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ExportMeasurementDataResponse{}
	_body, _err := client.DoRequest(tea.String("ExportMeasurementData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportMeasurementData(request *ExportMeasurementDataRequest) (_result *ExportMeasurementDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportMeasurementDataResponse{}
	_body, _err := client.ExportMeasurementDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMeasurementDataWithOptions(request *DescribeMeasurementDataRequest, runtime *util.RuntimeOptions) (_result *DescribeMeasurementDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeMeasurementDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeMeasurementData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMeasurementData(request *DescribeMeasurementDataRequest) (_result *DescribeMeasurementDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMeasurementDataResponse{}
	_body, _err := client.DescribeMeasurementDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableResourceInfoWithOptions(request *DescribeAvailableResourceInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeAvailableResourceInfoResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeAvailableResourceInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableResourceInfo(request *DescribeAvailableResourceInfoRequest) (_result *DescribeAvailableResourceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableResourceInfoResponse{}
	_body, _err := client.DescribeAvailableResourceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePrePaidInstanceStockWithOptions(request *DescribePrePaidInstanceStockRequest, runtime *util.RuntimeOptions) (_result *DescribePrePaidInstanceStockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribePrePaidInstanceStockResponse{}
	_body, _err := client.DoRequest(tea.String("DescribePrePaidInstanceStock"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePrePaidInstanceStock(request *DescribePrePaidInstanceStockRequest) (_result *DescribePrePaidInstanceStockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePrePaidInstanceStockResponse{}
	_body, _err := client.DescribePrePaidInstanceStockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnassociateEipAddressWithOptions(request *UnassociateEipAddressRequest, runtime *util.RuntimeOptions) (_result *UnassociateEipAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UnassociateEipAddressResponse{}
	_body, _err := client.DoRequest(tea.String("UnassociateEipAddress"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnassociateEipAddress(request *UnassociateEipAddressRequest) (_result *UnassociateEipAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnassociateEipAddressResponse{}
	_body, _err := client.UnassociateEipAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseEipAddressWithOptions(request *ReleaseEipAddressRequest, runtime *util.RuntimeOptions) (_result *ReleaseEipAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ReleaseEipAddressResponse{}
	_body, _err := client.DoRequest(tea.String("ReleaseEipAddress"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseEipAddress(request *ReleaseEipAddressRequest) (_result *ReleaseEipAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseEipAddressResponse{}
	_body, _err := client.ReleaseEipAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEipAddressesWithOptions(request *DescribeEipAddressesRequest, runtime *util.RuntimeOptions) (_result *DescribeEipAddressesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeEipAddressesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeEipAddresses"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEipAddresses(request *DescribeEipAddressesRequest) (_result *DescribeEipAddressesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEipAddressesResponse{}
	_body, _err := client.DescribeEipAddressesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateEipAddressWithOptions(request *AssociateEipAddressRequest, runtime *util.RuntimeOptions) (_result *AssociateEipAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AssociateEipAddressResponse{}
	_body, _err := client.DoRequest(tea.String("AssociateEipAddress"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateEipAddress(request *AssociateEipAddressRequest) (_result *AssociateEipAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateEipAddressResponse{}
	_body, _err := client.AssociateEipAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AllocateEipAddressWithOptions(request *AllocateEipAddressRequest, runtime *util.RuntimeOptions) (_result *AllocateEipAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AllocateEipAddressResponse{}
	_body, _err := client.DoRequest(tea.String("AllocateEipAddress"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocateEipAddress(request *AllocateEipAddressRequest) (_result *AllocateEipAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocateEipAddressResponse{}
	_body, _err := client.AllocateEipAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleasePostPaidInstanceWithOptions(request *ReleasePostPaidInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleasePostPaidInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ReleasePostPaidInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("ReleasePostPaidInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleasePostPaidInstance(request *ReleasePostPaidInstanceRequest) (_result *ReleasePostPaidInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleasePostPaidInstanceResponse{}
	_body, _err := client.ReleasePostPaidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleasePrePaidInstanceWithOptions(request *ReleasePrePaidInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleasePrePaidInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ReleasePrePaidInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("ReleasePrePaidInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleasePrePaidInstance(request *ReleasePrePaidInstanceRequest) (_result *ReleasePrePaidInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleasePrePaidInstanceResponse{}
	_body, _err := client.ReleasePrePaidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachEnsInstancesWithOptions(request *AttachEnsInstancesRequest, runtime *util.RuntimeOptions) (_result *AttachEnsInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AttachEnsInstancesResponse{}
	_body, _err := client.DoRequest(tea.String("AttachEnsInstances"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachEnsInstances(request *AttachEnsInstancesRequest) (_result *AttachEnsInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachEnsInstancesResponse{}
	_body, _err := client.AttachEnsInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeReservedResourceWithOptions(request *DescribeReservedResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeReservedResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeReservedResourceResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeReservedResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeReservedResource(request *DescribeReservedResourceRequest) (_result *DescribeReservedResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeReservedResourceResponse{}
	_body, _err := client.DescribeReservedResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceTypesWithOptions(request *DescribeInstanceTypesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeInstanceTypesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeInstanceTypes"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceTypes(request *DescribeInstanceTypesRequest) (_result *DescribeInstanceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceTypesResponse{}
	_body, _err := client.DescribeInstanceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageWithOptions(request *CreateImageRequest, runtime *util.RuntimeOptions) (_result *CreateImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateImageResponse{}
	_body, _err := client.DoRequest(tea.String("CreateImage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImage(request *CreateImageRequest) (_result *CreateImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageResponse{}
	_body, _err := client.CreateImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnsNetSaleDistrictWithOptions(request *DescribeEnsNetSaleDistrictRequest, runtime *util.RuntimeOptions) (_result *DescribeEnsNetSaleDistrictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeEnsNetSaleDistrictResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeEnsNetSaleDistrict"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnsNetSaleDistrict(request *DescribeEnsNetSaleDistrictRequest) (_result *DescribeEnsNetSaleDistrictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnsNetSaleDistrictResponse{}
	_body, _err := client.DescribeEnsNetSaleDistrictWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnsNetDistrictWithOptions(request *DescribeEnsNetDistrictRequest, runtime *util.RuntimeOptions) (_result *DescribeEnsNetDistrictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeEnsNetDistrictResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeEnsNetDistrict"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnsNetDistrict(request *DescribeEnsNetDistrictRequest) (_result *DescribeEnsNetDistrictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnsNetDistrictResponse{}
	_body, _err := client.DescribeEnsNetDistrictWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PreCreateEnsServiceWithOptions(request *PreCreateEnsServiceRequest, runtime *util.RuntimeOptions) (_result *PreCreateEnsServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PreCreateEnsServiceResponse{}
	_body, _err := client.DoRequest(tea.String("PreCreateEnsService"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PreCreateEnsService(request *PreCreateEnsServiceRequest) (_result *PreCreateEnsServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PreCreateEnsServiceResponse{}
	_body, _err := client.PreCreateEnsServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBandWithdChargeTypeWithOptions(request *DescribeBandWithdChargeTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeBandWithdChargeTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeBandWithdChargeTypeResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeBandWithdChargeType"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBandWithdChargeType(request *DescribeBandWithdChargeTypeRequest) (_result *DescribeBandWithdChargeTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBandWithdChargeTypeResponse{}
	_body, _err := client.DescribeBandWithdChargeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyImageAttributeWithOptions(request *ModifyImageAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyImageAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyImageAttributeResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyImageAttribute"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyImageAttribute(request *ModifyImageAttributeRequest) (_result *ModifyImageAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyImageAttributeResponse{}
	_body, _err := client.ModifyImageAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEnsServiceWithOptions(request *CreateEnsServiceRequest, runtime *util.RuntimeOptions) (_result *CreateEnsServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateEnsServiceResponse{}
	_body, _err := client.DoRequest(tea.String("CreateEnsService"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEnsService(request *CreateEnsServiceRequest) (_result *CreateEnsServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEnsServiceResponse{}
	_body, _err := client.CreateEnsServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnsNetLevelWithOptions(request *DescribeEnsNetLevelRequest, runtime *util.RuntimeOptions) (_result *DescribeEnsNetLevelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeEnsNetLevelResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeEnsNetLevel"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnsNetLevel(request *DescribeEnsNetLevelRequest) (_result *DescribeEnsNetLevelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnsNetLevelResponse{}
	_body, _err := client.DescribeEnsNetLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceSpecWithOptions(request *DescribeInstanceSpecRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeInstanceSpecResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeInstanceSpec"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceSpec(request *DescribeInstanceSpecRequest) (_result *DescribeInstanceSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceSpecResponse{}
	_body, _err := client.DescribeInstanceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceAutoRenewAttributeWithOptions(request *DescribeInstanceAutoRenewAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceAutoRenewAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeInstanceAutoRenewAttributeResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeInstanceAutoRenewAttribute"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceAutoRenewAttribute(request *DescribeInstanceAutoRenewAttributeRequest) (_result *DescribeInstanceAutoRenewAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceAutoRenewAttributeResponse{}
	_body, _err := client.DescribeInstanceAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableResourceWithOptions(request *DescribeAvailableResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeAvailableResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableResource(request *DescribeAvailableResourceRequest) (_result *DescribeAvailableResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DescribeAvailableResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("CreateInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReInitDiskWithOptions(request *ReInitDiskRequest, runtime *util.RuntimeOptions) (_result *ReInitDiskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ReInitDiskResponse{}
	_body, _err := client.DoRequest(tea.String("ReInitDisk"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReInitDisk(request *ReInitDiskRequest) (_result *ReInitDiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReInitDiskResponse{}
	_body, _err := client.ReInitDiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageInfosWithOptions(request *DescribeImageInfosRequest, runtime *util.RuntimeOptions) (_result *DescribeImageInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeImageInfosResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeImageInfos"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImageInfos(request *DescribeImageInfosRequest) (_result *DescribeImageInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageInfosResponse{}
	_body, _err := client.DescribeImageInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserBandWidthDataWithOptions(request *DescribeUserBandWidthDataRequest, runtime *util.RuntimeOptions) (_result *DescribeUserBandWidthDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeUserBandWidthDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeUserBandWidthData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserBandWidthData(request *DescribeUserBandWidthDataRequest) (_result *DescribeUserBandWidthDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserBandWidthDataResponse{}
	_body, _err := client.DescribeUserBandWidthDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebootInstanceWithOptions(request *RebootInstanceRequest, runtime *util.RuntimeOptions) (_result *RebootInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RebootInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("RebootInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebootInstance(request *RebootInstanceRequest) (_result *RebootInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebootInstanceResponse{}
	_body, _err := client.RebootInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnsRegionsWithOptions(request *DescribeEnsRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeEnsRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeEnsRegionsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeEnsRegions"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnsRegions(request *DescribeEnsRegionsRequest) (_result *DescribeEnsRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnsRegionsResponse{}
	_body, _err := client.DescribeEnsRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("StartInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceMonitorDataWithOptions(request *DescribeInstanceMonitorDataRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceMonitorDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeInstanceMonitorDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeInstanceMonitorData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceMonitorData(request *DescribeInstanceMonitorDataRequest) (_result *DescribeInstanceMonitorDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceMonitorDataResponse{}
	_body, _err := client.DescribeInstanceMonitorDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeInstances"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImagesWithOptions(request *DescribeImagesRequest, runtime *util.RuntimeOptions) (_result *DescribeImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeImages"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImages(request *DescribeImagesRequest) (_result *DescribeImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DescribeImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopInstanceWithOptions(request *StopInstanceRequest, runtime *util.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("StopInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopInstance(request *StopInstanceRequest) (_result *StopInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceAttributeWithOptions(request *ModifyInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyInstanceAttribute"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-10"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) (_result *ModifyInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.ModifyInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
