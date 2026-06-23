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
	// The application ID. You can call the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation to obtain the application ID.
	//
	// 	Notice: The AppId format is the app- prefix followed by a numeric suffix, with a total length of 20 to 64 characters (example: app-8806886***83794688). Call ListEdgeContainerApps to obtain an existing AppId, or call CreateEdgeContainerApp to create an application first.</notice>.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
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
	// Specifies whether to permanently enable reservation. Once enabled, you cannot set a reservation end time.
	//
	// example:
	//
	// true
	Forever *bool `json:"Forever,omitempty" xml:"Forever,omitempty"`
	// The list of reserved resources.
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

type UpdateEdgeContainerAppResourceReserveRequestReserveSet struct {
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
