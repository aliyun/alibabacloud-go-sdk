// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEdgeContainerAppResourceReserveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateEdgeContainerAppResourceReserveRequest
	GetAppId() *string
	SetDurationTime(v string) *UpdateEdgeContainerAppResourceReserveRequest
	GetDurationTime() *string
	SetEnable(v bool) *UpdateEdgeContainerAppResourceReserveRequest
	GetEnable() *bool
	SetForever(v bool) *UpdateEdgeContainerAppResourceReserveRequest
	GetForever() *bool
	SetReserveSet(v []*UpdateEdgeContainerAppResourceReserveRequestReserveSet) *UpdateEdgeContainerAppResourceReserveRequest
	GetReserveSet() []*UpdateEdgeContainerAppResourceReserveRequestReserveSet
}

type UpdateEdgeContainerAppResourceReserveRequest struct {
	// The application ID, which can be obtained by calling the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time of the reservation. The input time is UTC. It takes +8 hours to enter Beijing time. For example, if the current time is 2006-01-02 06:04:05, you need to enter "2006-01-02T14:04:05Z".
	//
	// example:
	//
	// 2006-01-02T15:04:05Z
	DurationTime *string `json:"DurationTime,omitempty" xml:"DurationTime,omitempty"`
	// Whether to enable resource reservation.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Whether to permanently enable the reservation. Once it is enabled, you are not allowed to set the reservation deadline.
	//
	// example:
	//
	// true
	Forever *bool `json:"Forever,omitempty" xml:"Forever,omitempty"`
	// Reserved resource list.
	ReserveSet []*UpdateEdgeContainerAppResourceReserveRequestReserveSet `json:"ReserveSet,omitempty" xml:"ReserveSet,omitempty" type:"Repeated"`
}

func (s UpdateEdgeContainerAppResourceReserveRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeContainerAppResourceReserveRequest) GoString() string {
	return s.String()
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) GetDurationTime() *string {
	return s.DurationTime
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) GetForever() *bool {
	return s.Forever
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) GetReserveSet() []*UpdateEdgeContainerAppResourceReserveRequestReserveSet {
	return s.ReserveSet
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) SetAppId(v string) *UpdateEdgeContainerAppResourceReserveRequest {
	s.AppId = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) SetDurationTime(v string) *UpdateEdgeContainerAppResourceReserveRequest {
	s.DurationTime = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) SetEnable(v bool) *UpdateEdgeContainerAppResourceReserveRequest {
	s.Enable = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) SetForever(v bool) *UpdateEdgeContainerAppResourceReserveRequest {
	s.Forever = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) SetReserveSet(v []*UpdateEdgeContainerAppResourceReserveRequestReserveSet) *UpdateEdgeContainerAppResourceReserveRequest {
	s.ReserveSet = v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateEdgeContainerAppResourceReserveRequestReserveSet struct {
	// The ISP. The following types are supported. You do not need to enter the ISP in regions outside the Chinese mainland:
	//
	// 	- China Mobile: cmcc
	//
	// 	- China Telecom: chinanet
	//
	// 	- China Unicom: unicom
	//
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// Information about the region. The Chinese mainland supports the input of regions and special administrative regions, and the regions outside the Chinese mainland support the input of countries. The following is the corresponding parameter mapping:
	//
	// Chinese mainland:
	//
	// 	- East China: huadong
	//
	// 	- South China: huanan
	//
	// 	- Central China: huazhong
	//
	// 	- North China: huabei
	//
	// 	- Northwest China: xibei
	//
	// 	- Southwest China: xinan
	//
	// 	- Northeast China: dongbei
	//
	// Special Administrative Regions and overseas:
	//
	// 	- Taiwan, China: tw
	//
	// 	- Macau, China: mo
	//
	// 	- Hong Kong, China: hk
	//
	// 	- Japan: jp
	//
	// 	- United States: us
	//
	// 	- Thailand: th
	//
	// 	- Korea: kr
	//
	// 	- Russia: ru
	//
	// 	- Singapore: sg
	//
	// 	- France: fr
	//
	// 	- Spain: es
	//
	// 	- Italy: it
	//
	// 	- Sweden: se
	//
	// 	- UAE: ae
	//
	// 	- Indonesia: id
	//
	// 	- Chile: cl
	//
	// 	- Philippines: ph
	//
	// 	- Malaysia: my
	//
	// 	- Vietnam: vn
	//
	// 	- Argentina: AR
	//
	// 	- Australia: au
	//
	// 	- Brazil: br
	//
	// 	- Colombia: co
	//
	// 	- Germany: de
	//
	// 	- UK: GB
	//
	// 	- Peru: pe
	//
	// 	- Saudi Arabia: sa
	//
	// 	- Netherlands: nl
	//
	// 	- South Africa: za
	//
	// example:
	//
	// huazhong
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The number of container replicas.
	//
	// example:
	//
	// 1
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
}

func (s UpdateEdgeContainerAppResourceReserveRequestReserveSet) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeContainerAppResourceReserveRequestReserveSet) GoString() string {
	return s.String()
}

func (s *UpdateEdgeContainerAppResourceReserveRequestReserveSet) GetIsp() *string {
	return s.Isp
}

func (s *UpdateEdgeContainerAppResourceReserveRequestReserveSet) GetRegion() *string {
	return s.Region
}

func (s *UpdateEdgeContainerAppResourceReserveRequestReserveSet) GetReplicas() *int32 {
	return s.Replicas
}

func (s *UpdateEdgeContainerAppResourceReserveRequestReserveSet) SetIsp(v string) *UpdateEdgeContainerAppResourceReserveRequestReserveSet {
	s.Isp = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveRequestReserveSet) SetRegion(v string) *UpdateEdgeContainerAppResourceReserveRequestReserveSet {
	s.Region = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveRequestReserveSet) SetReplicas(v int32) *UpdateEdgeContainerAppResourceReserveRequestReserveSet {
	s.Replicas = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveRequestReserveSet) Validate() error {
	return dara.Validate(s)
}
