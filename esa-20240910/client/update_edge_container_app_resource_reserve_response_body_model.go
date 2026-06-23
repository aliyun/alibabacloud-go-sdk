// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEdgeContainerAppResourceReserveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDurationTime(v string) *UpdateEdgeContainerAppResourceReserveResponseBody
	GetDurationTime() *string
	SetEnable(v bool) *UpdateEdgeContainerAppResourceReserveResponseBody
	GetEnable() *bool
	SetForever(v bool) *UpdateEdgeContainerAppResourceReserveResponseBody
	GetForever() *bool
	SetRequestId(v string) *UpdateEdgeContainerAppResourceReserveResponseBody
	GetRequestId() *string
	SetReserveSet(v []*UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) *UpdateEdgeContainerAppResourceReserveResponseBody
	GetReserveSet() []*UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet
}

type UpdateEdgeContainerAppResourceReserveResponseBody struct {
	// The reservation end time. This parameter uses UTC time. To convert from UTC+8, add 8 hours. For example, if the current time is 2006-01-02 06:04:05 in UTC+8, enter "2006-01-02T14:04:05Z".
	//
	// example:
	//
	// 2006-01-02T15:04:05Z
	DurationTime *string `json:"DurationTime,omitempty" xml:"DurationTime,omitempty"`
	// Specifies whether to enable resource reservation.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Indicates whether reservation is permanently enabled.
	//
	// example:
	//
	// true
	Forever *bool `json:"Forever,omitempty" xml:"Forever,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AB799CF-562A-5CAF-A99E-4354053D814F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of reserved resources.
	ReserveSet []*UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet `json:"ReserveSet,omitempty" xml:"ReserveSet,omitempty" type:"Repeated"`
}

func (s UpdateEdgeContainerAppResourceReserveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeContainerAppResourceReserveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) GetDurationTime() *string {
	return s.DurationTime
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) GetForever() *bool {
	return s.Forever
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) GetReserveSet() []*UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet {
	return s.ReserveSet
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) SetDurationTime(v string) *UpdateEdgeContainerAppResourceReserveResponseBody {
	s.DurationTime = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) SetEnable(v bool) *UpdateEdgeContainerAppResourceReserveResponseBody {
	s.Enable = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) SetForever(v bool) *UpdateEdgeContainerAppResourceReserveResponseBody {
	s.Forever = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) SetRequestId(v string) *UpdateEdgeContainerAppResourceReserveResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) SetReserveSet(v []*UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) *UpdateEdgeContainerAppResourceReserveResponseBody {
	s.ReserveSet = v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) Validate() error {
	if s.ReserveSet != nil {
		for _, item := range s.ReserveSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet struct {
	// The Internet service provider (ISP). The following ISPs are supported. You do not need to specify an ISP for special administrative regions or areas outside China:
	//
	// - China Mobile: cmcc
	//
	// - China Telecom: chinanet
	//
	// - China Unicom: unicom.
	//
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The region information. For the Chinese mainland, you can specify a major region. For special administrative regions and areas outside China, you can specify a country or region. The following list shows the parameter mappings:
	//
	// Chinese mainland:
	//
	// - East China: huadong
	//
	// - South China: huanan
	//
	// - Central China: huazhong
	//
	// - North China: huabei
	//
	// - Northwest China: xibei
	//
	// - Southwest China: xinan
	//
	// - Northeast China: dongbei
	//
	// Special administrative regions and outside China:
	//
	// - Taiwan (China): tw
	//
	// - Macao (China): mo
	//
	// - Hong Kong (China): hk
	//
	// - Japan: jp
	//
	// - United States: us
	//
	// - Thailand: th
	//
	// - South Korea: kr
	//
	// - Russia: ru
	//
	// - Singapore: sg
	//
	// - France: fr
	//
	// - Spain: es
	//
	// - Italy: it
	//
	// - Sweden: se
	//
	// - United Arab Emirates: ae
	//
	// - Indonesia: id
	//
	// - Chile: cl
	//
	// - Philippines: ph
	//
	// - Malaysia: my
	//
	// - Vietnam: vn
	//
	// - Argentina: ar
	//
	// - Australia: au
	//
	// - Brazil: br
	//
	// - Colombia: co
	//
	// - Germany: de
	//
	// - United Kingdom: gb
	//
	// - Peru: pe
	//
	// - Saudi Arabia: sa
	//
	// - Netherlands: nl
	//
	// - South Africa: za.
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

func (s UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) GoString() string {
	return s.String()
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) GetIsp() *string {
	return s.Isp
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) GetRegion() *string {
	return s.Region
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) GetReplicas() *int32 {
	return s.Replicas
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) SetIsp(v string) *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet {
	s.Isp = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) SetRegion(v string) *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet {
	s.Region = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) SetReplicas(v int32) *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet {
	s.Replicas = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) Validate() error {
	return dara.Validate(s)
}
