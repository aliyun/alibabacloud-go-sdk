// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushResourcePlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBufferCnt(v int64) *PushResourcePlanRequest
	GetBufferCnt() *int64
	SetDemandCount(v int64) *PushResourcePlanRequest
	GetDemandCount() *int64
	SetDemandId(v string) *PushResourcePlanRequest
	GetDemandId() *string
	SetMethodList(v []*PushResourcePlanRequestMethodList) *PushResourcePlanRequest
	GetMethodList() []*PushResourcePlanRequestMethodList
	SetRequestId(v string) *PushResourcePlanRequest
	GetRequestId() *string
	SetRequireCnt(v int64) *PushResourcePlanRequest
	GetRequireCnt() *int64
	SetSubDemandId(v string) *PushResourcePlanRequest
	GetSubDemandId() *string
}

type PushResourcePlanRequest struct {
	BufferCnt   *int64                               `json:"bufferCnt,omitempty" xml:"bufferCnt,omitempty"`
	DemandCount *int64                               `json:"demandCount,omitempty" xml:"demandCount,omitempty"`
	DemandId    *string                              `json:"demandId,omitempty" xml:"demandId,omitempty"`
	MethodList  []*PushResourcePlanRequestMethodList `json:"methodList,omitempty" xml:"methodList,omitempty" type:"Repeated"`
	RequestId   *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	RequireCnt  *int64                               `json:"requireCnt,omitempty" xml:"requireCnt,omitempty"`
	SubDemandId *string                              `json:"subDemandId,omitempty" xml:"subDemandId,omitempty"`
}

func (s PushResourcePlanRequest) String() string {
	return dara.Prettify(s)
}

func (s PushResourcePlanRequest) GoString() string {
	return s.String()
}

func (s *PushResourcePlanRequest) GetBufferCnt() *int64 {
	return s.BufferCnt
}

func (s *PushResourcePlanRequest) GetDemandCount() *int64 {
	return s.DemandCount
}

func (s *PushResourcePlanRequest) GetDemandId() *string {
	return s.DemandId
}

func (s *PushResourcePlanRequest) GetMethodList() []*PushResourcePlanRequestMethodList {
	return s.MethodList
}

func (s *PushResourcePlanRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *PushResourcePlanRequest) GetRequireCnt() *int64 {
	return s.RequireCnt
}

func (s *PushResourcePlanRequest) GetSubDemandId() *string {
	return s.SubDemandId
}

func (s *PushResourcePlanRequest) SetBufferCnt(v int64) *PushResourcePlanRequest {
	s.BufferCnt = &v
	return s
}

func (s *PushResourcePlanRequest) SetDemandCount(v int64) *PushResourcePlanRequest {
	s.DemandCount = &v
	return s
}

func (s *PushResourcePlanRequest) SetDemandId(v string) *PushResourcePlanRequest {
	s.DemandId = &v
	return s
}

func (s *PushResourcePlanRequest) SetMethodList(v []*PushResourcePlanRequestMethodList) *PushResourcePlanRequest {
	s.MethodList = v
	return s
}

func (s *PushResourcePlanRequest) SetRequestId(v string) *PushResourcePlanRequest {
	s.RequestId = &v
	return s
}

func (s *PushResourcePlanRequest) SetRequireCnt(v int64) *PushResourcePlanRequest {
	s.RequireCnt = &v
	return s
}

func (s *PushResourcePlanRequest) SetSubDemandId(v string) *PushResourcePlanRequest {
	s.SubDemandId = &v
	return s
}

func (s *PushResourcePlanRequest) Validate() error {
	return dara.Validate(s)
}

type PushResourcePlanRequestMethodList struct {
	Azone            *string                                      `json:"azone,omitempty" xml:"azone,omitempty"`
	BufferCnt        *int64                                       `json:"bufferCnt,omitempty" xml:"bufferCnt,omitempty"`
	Cluster          *string                                      `json:"cluster,omitempty" xml:"cluster,omitempty"`
	Comment          *string                                      `json:"comment,omitempty" xml:"comment,omitempty"`
	ConvertHostCnt   *int64                                       `json:"convertHostCnt,omitempty" xml:"convertHostCnt,omitempty"`
	ConvertHostType  *string                                      `json:"convertHostType,omitempty" xml:"convertHostType,omitempty"`
	DataList         []*PushResourcePlanRequestMethodListDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	DenamdCount      *int64                                       `json:"denamdCount,omitempty" xml:"denamdCount,omitempty"`
	GapCnt           *int64                                       `json:"gapCnt,omitempty" xml:"gapCnt,omitempty"`
	PromiseDate      *string                                      `json:"promiseDate,omitempty" xml:"promiseDate,omitempty"`
	Region           *string                                      `json:"region,omitempty" xml:"region,omitempty"`
	ResourceMethodId *int64                                       `json:"resourceMethodId,omitempty" xml:"resourceMethodId,omitempty"`
	RoomCode         *string                                      `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
}

func (s PushResourcePlanRequestMethodList) String() string {
	return dara.Prettify(s)
}

func (s PushResourcePlanRequestMethodList) GoString() string {
	return s.String()
}

func (s *PushResourcePlanRequestMethodList) GetAzone() *string {
	return s.Azone
}

func (s *PushResourcePlanRequestMethodList) GetBufferCnt() *int64 {
	return s.BufferCnt
}

func (s *PushResourcePlanRequestMethodList) GetCluster() *string {
	return s.Cluster
}

func (s *PushResourcePlanRequestMethodList) GetComment() *string {
	return s.Comment
}

func (s *PushResourcePlanRequestMethodList) GetConvertHostCnt() *int64 {
	return s.ConvertHostCnt
}

func (s *PushResourcePlanRequestMethodList) GetConvertHostType() *string {
	return s.ConvertHostType
}

func (s *PushResourcePlanRequestMethodList) GetDataList() []*PushResourcePlanRequestMethodListDataList {
	return s.DataList
}

func (s *PushResourcePlanRequestMethodList) GetDenamdCount() *int64 {
	return s.DenamdCount
}

func (s *PushResourcePlanRequestMethodList) GetGapCnt() *int64 {
	return s.GapCnt
}

func (s *PushResourcePlanRequestMethodList) GetPromiseDate() *string {
	return s.PromiseDate
}

func (s *PushResourcePlanRequestMethodList) GetRegion() *string {
	return s.Region
}

func (s *PushResourcePlanRequestMethodList) GetResourceMethodId() *int64 {
	return s.ResourceMethodId
}

func (s *PushResourcePlanRequestMethodList) GetRoomCode() *string {
	return s.RoomCode
}

func (s *PushResourcePlanRequestMethodList) SetAzone(v string) *PushResourcePlanRequestMethodList {
	s.Azone = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetBufferCnt(v int64) *PushResourcePlanRequestMethodList {
	s.BufferCnt = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetCluster(v string) *PushResourcePlanRequestMethodList {
	s.Cluster = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetComment(v string) *PushResourcePlanRequestMethodList {
	s.Comment = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetConvertHostCnt(v int64) *PushResourcePlanRequestMethodList {
	s.ConvertHostCnt = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetConvertHostType(v string) *PushResourcePlanRequestMethodList {
	s.ConvertHostType = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetDataList(v []*PushResourcePlanRequestMethodListDataList) *PushResourcePlanRequestMethodList {
	s.DataList = v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetDenamdCount(v int64) *PushResourcePlanRequestMethodList {
	s.DenamdCount = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetGapCnt(v int64) *PushResourcePlanRequestMethodList {
	s.GapCnt = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetPromiseDate(v string) *PushResourcePlanRequestMethodList {
	s.PromiseDate = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetRegion(v string) *PushResourcePlanRequestMethodList {
	s.Region = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetResourceMethodId(v int64) *PushResourcePlanRequestMethodList {
	s.ResourceMethodId = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetRoomCode(v string) *PushResourcePlanRequestMethodList {
	s.RoomCode = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) Validate() error {
	return dara.Validate(s)
}

type PushResourcePlanRequestMethodListDataList struct {
	ClassZone       *string `json:"classZone,omitempty" xml:"classZone,omitempty"`
	ConvertHostType *string `json:"convertHostType,omitempty" xml:"convertHostType,omitempty"`
	LogicZone       *string `json:"logicZone,omitempty" xml:"logicZone,omitempty"`
	NetArch         *string `json:"netArch,omitempty" xml:"netArch,omitempty"`
	Nic             *string `json:"nic,omitempty" xml:"nic,omitempty"`
	Product3        *string `json:"product3,omitempty" xml:"product3,omitempty"`
	SafeZone        *string `json:"safeZone,omitempty" xml:"safeZone,omitempty"`
	Scenario        *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
	SupplyAmount    *int64  `json:"supplyAmount,omitempty" xml:"supplyAmount,omitempty"`
	// This parameter is required.
	SupplyDate     *string `json:"supplyDate,omitempty" xml:"supplyDate,omitempty"`
	SupplyType     *int64  `json:"supplyType,omitempty" xml:"supplyType,omitempty"`
	SupplyVmAmount *int32  `json:"supplyVmAmount,omitempty" xml:"supplyVmAmount,omitempty"`
}

func (s PushResourcePlanRequestMethodListDataList) String() string {
	return dara.Prettify(s)
}

func (s PushResourcePlanRequestMethodListDataList) GoString() string {
	return s.String()
}

func (s *PushResourcePlanRequestMethodListDataList) GetClassZone() *string {
	return s.ClassZone
}

func (s *PushResourcePlanRequestMethodListDataList) GetConvertHostType() *string {
	return s.ConvertHostType
}

func (s *PushResourcePlanRequestMethodListDataList) GetLogicZone() *string {
	return s.LogicZone
}

func (s *PushResourcePlanRequestMethodListDataList) GetNetArch() *string {
	return s.NetArch
}

func (s *PushResourcePlanRequestMethodListDataList) GetNic() *string {
	return s.Nic
}

func (s *PushResourcePlanRequestMethodListDataList) GetProduct3() *string {
	return s.Product3
}

func (s *PushResourcePlanRequestMethodListDataList) GetSafeZone() *string {
	return s.SafeZone
}

func (s *PushResourcePlanRequestMethodListDataList) GetScenario() *string {
	return s.Scenario
}

func (s *PushResourcePlanRequestMethodListDataList) GetSupplyAmount() *int64 {
	return s.SupplyAmount
}

func (s *PushResourcePlanRequestMethodListDataList) GetSupplyDate() *string {
	return s.SupplyDate
}

func (s *PushResourcePlanRequestMethodListDataList) GetSupplyType() *int64 {
	return s.SupplyType
}

func (s *PushResourcePlanRequestMethodListDataList) GetSupplyVmAmount() *int32 {
	return s.SupplyVmAmount
}

func (s *PushResourcePlanRequestMethodListDataList) SetClassZone(v string) *PushResourcePlanRequestMethodListDataList {
	s.ClassZone = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetConvertHostType(v string) *PushResourcePlanRequestMethodListDataList {
	s.ConvertHostType = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetLogicZone(v string) *PushResourcePlanRequestMethodListDataList {
	s.LogicZone = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetNetArch(v string) *PushResourcePlanRequestMethodListDataList {
	s.NetArch = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetNic(v string) *PushResourcePlanRequestMethodListDataList {
	s.Nic = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetProduct3(v string) *PushResourcePlanRequestMethodListDataList {
	s.Product3 = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetSafeZone(v string) *PushResourcePlanRequestMethodListDataList {
	s.SafeZone = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetScenario(v string) *PushResourcePlanRequestMethodListDataList {
	s.Scenario = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetSupplyAmount(v int64) *PushResourcePlanRequestMethodListDataList {
	s.SupplyAmount = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetSupplyDate(v string) *PushResourcePlanRequestMethodListDataList {
	s.SupplyDate = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetSupplyType(v int64) *PushResourcePlanRequestMethodListDataList {
	s.SupplyType = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetSupplyVmAmount(v int32) *PushResourcePlanRequestMethodListDataList {
	s.SupplyVmAmount = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) Validate() error {
	return dara.Validate(s)
}
