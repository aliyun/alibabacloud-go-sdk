// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppResourceReserveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDurationTime(v string) *GetEdgeContainerAppResourceReserveResponseBody
	GetDurationTime() *string
	SetEnable(v bool) *GetEdgeContainerAppResourceReserveResponseBody
	GetEnable() *bool
	SetForever(v bool) *GetEdgeContainerAppResourceReserveResponseBody
	GetForever() *bool
	SetRequestId(v string) *GetEdgeContainerAppResourceReserveResponseBody
	GetRequestId() *string
	SetReserveSet(v []*GetEdgeContainerAppResourceReserveResponseBodyReserveSet) *GetEdgeContainerAppResourceReserveResponseBody
	GetReserveSet() []*GetEdgeContainerAppResourceReserveResponseBodyReserveSet
}

type GetEdgeContainerAppResourceReserveResponseBody struct {
	// The end time of the reservation. The input is UTC time. It takes +8 hours to enter Beijing time. For example, if the current time is 2006-01-02 06:04:05 , you need to enter "2006-01-02T14:04:05Z".
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
	// Whether to enable resource reservation permanently.
	//
	// example:
	//
	// true
	Forever *bool `json:"Forever,omitempty" xml:"Forever,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Reserved resource list.
	ReserveSet []*GetEdgeContainerAppResourceReserveResponseBodyReserveSet `json:"ReserveSet,omitempty" xml:"ReserveSet,omitempty" type:"Repeated"`
}

func (s GetEdgeContainerAppResourceReserveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResourceReserveResponseBody) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) GetDurationTime() *string {
	return s.DurationTime
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) GetEnable() *bool {
	return s.Enable
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) GetForever() *bool {
	return s.Forever
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) GetReserveSet() []*GetEdgeContainerAppResourceReserveResponseBodyReserveSet {
	return s.ReserveSet
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) SetDurationTime(v string) *GetEdgeContainerAppResourceReserveResponseBody {
	s.DurationTime = &v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) SetEnable(v bool) *GetEdgeContainerAppResourceReserveResponseBody {
	s.Enable = &v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) SetForever(v bool) *GetEdgeContainerAppResourceReserveResponseBody {
	s.Forever = &v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) SetRequestId(v string) *GetEdgeContainerAppResourceReserveResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) SetReserveSet(v []*GetEdgeContainerAppResourceReserveResponseBodyReserveSet) *GetEdgeContainerAppResourceReserveResponseBody {
	s.ReserveSet = v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetEdgeContainerAppResourceReserveResponseBodyReserveSet struct {
	// The following ISPs are supported. You do not need to enter this field for overseas and special administrative regions. ISP:
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
	// 	- Northwest: xibei
	//
	// 	- Southwest: xinan
	//
	// 	- Northeast China: dongbei
	//
	// Special Administrative Regions and Overseas:
	//
	// 	- Taiwan, China: tw
	//
	// 	- Macau China: mo
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
	// 	- Argentina: ar
	//
	// 	- Australia: au
	//
	// 	- Brazil: br
	//
	// 	- Colombia: co
	//
	// 	- Germany: de
	//
	// 	- UK: gb
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

func (s GetEdgeContainerAppResourceReserveResponseBodyReserveSet) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResourceReserveResponseBodyReserveSet) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResourceReserveResponseBodyReserveSet) GetIsp() *string {
	return s.Isp
}

func (s *GetEdgeContainerAppResourceReserveResponseBodyReserveSet) GetRegion() *string {
	return s.Region
}

func (s *GetEdgeContainerAppResourceReserveResponseBodyReserveSet) GetReplicas() *int32 {
	return s.Replicas
}

func (s *GetEdgeContainerAppResourceReserveResponseBodyReserveSet) SetIsp(v string) *GetEdgeContainerAppResourceReserveResponseBodyReserveSet {
	s.Isp = &v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponseBodyReserveSet) SetRegion(v string) *GetEdgeContainerAppResourceReserveResponseBodyReserveSet {
	s.Region = &v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponseBodyReserveSet) SetReplicas(v int32) *GetEdgeContainerAppResourceReserveResponseBodyReserveSet {
	s.Replicas = &v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponseBodyReserveSet) Validate() error {
	return dara.Validate(s)
}
