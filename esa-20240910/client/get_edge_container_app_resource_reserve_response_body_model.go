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
	// The reservation expiration time. This parameter uses UTC time. To convert from UTC+8, add 8 hours. For example, if the current UTC+8 time is 2006-01-02 06:04:05, enter "2006-01-02T14:04:05Z".
	//
	// example:
	//
	// 2006-01-02T15:04:05Z
	DurationTime *string `json:"DurationTime,omitempty" xml:"DurationTime,omitempty"`
	// Indicates whether resource reservation is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Indicates whether the reservation is permanently enabled.
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
	// The list of reserved resources.
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

type GetEdgeContainerAppResourceReserveResponseBodyReserveSet struct {
	// The following Internet service providers (ISPs) are supported. This parameter is not required for regions outside the Chinese mainland.
	//
	// Valid values:
	//
	// - cmcc: China Mobile
	//
	// - chinanet: China Telecom
	//
	// - unicom: China Unicom.
	//
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The Chinese mainland:
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
