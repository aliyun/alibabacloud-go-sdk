// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBusinessAccessPointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessAccessPoints(v []*ListBusinessAccessPointsResponseBodyBusinessAccessPoints) *ListBusinessAccessPointsResponseBody
	GetBusinessAccessPoints() []*ListBusinessAccessPointsResponseBodyBusinessAccessPoints
	SetRequestId(v string) *ListBusinessAccessPointsResponseBody
	GetRequestId() *string
}

type ListBusinessAccessPointsResponseBody struct {
	// The list of all access point information for Express Connect circuits.
	BusinessAccessPoints []*ListBusinessAccessPointsResponseBodyBusinessAccessPoints `json:"BusinessAccessPoints,omitempty" xml:"BusinessAccessPoints,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 611CB80C-B6A9-43DB-9E38-0B0AC3D9B58F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBusinessAccessPointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBusinessAccessPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBusinessAccessPointsResponseBody) GetBusinessAccessPoints() []*ListBusinessAccessPointsResponseBodyBusinessAccessPoints {
	return s.BusinessAccessPoints
}

func (s *ListBusinessAccessPointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBusinessAccessPointsResponseBody) SetBusinessAccessPoints(v []*ListBusinessAccessPointsResponseBodyBusinessAccessPoints) *ListBusinessAccessPointsResponseBody {
	s.BusinessAccessPoints = v
	return s
}

func (s *ListBusinessAccessPointsResponseBody) SetRequestId(v string) *ListBusinessAccessPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBusinessAccessPointsResponseBody) Validate() error {
	if s.BusinessAccessPoints != nil {
		for _, item := range s.BusinessAccessPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBusinessAccessPointsResponseBodyBusinessAccessPoints struct {
	// The ID of the Express Connect circuit access point.
	//
	// example:
	//
	// ap-cn-hangzhou-xs-B
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The name of the Express Connect circuit access point.
	//
	// example:
	//
	// 杭州-萧山-B
	AccessPointName *string `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	// The CloudBox instance ID.
	//
	// > This parameter is available only when the queried Express Connect circuit and access point are CloudBox Express Connect circuits and CloudBox access points.
	//
	// example:
	//
	// cb-****
	CloudBoxInstanceIds *string `json:"CloudBoxInstanceIds,omitempty" xml:"CloudBoxInstanceIds,omitempty"`
	// The latitude of the access point.
	//
	// example:
	//
	// 30.198416
	Latitude *float64 `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// The longitude of the access point.
	//
	// example:
	//
	// 120.247514
	Longitude *float64 `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	// The collection of optical module models supported by the current access point.
	OpticalModuleModels []*ListBusinessAccessPointsResponseBodyBusinessAccessPointsOpticalModuleModels `json:"OpticalModuleModels,omitempty" xml:"OpticalModuleModels,omitempty" type:"Repeated"`
	// The telecommunications service providers that support physical line access. Valid values:
	//
	// - **CT**: China Telecom.
	//
	// - **CU**: China Unicom.
	//
	// - **CM**: China Mobile.
	//
	// - **CO**: Other Chinese providers.
	//
	// - **Equinix**: Equinix.
	//
	// - **Other**: Other providers outside the Chinese mainland.
	//
	// example:
	//
	// CT
	SupportLineOperator *string `json:"SupportLineOperator,omitempty" xml:"SupportLineOperator,omitempty"`
	// The port types available for purchase at the Express Connect circuit access point. Valid values:
	//
	// - **100Base-T**: 100M Ethernet port.
	//
	// - **1000Base-T**: 1 GE electrical port.
	//
	// - **1000Base-LX**: GE single-mode optical port (10 km).
	//
	// - **10GBase-T**: 10 GE electrical port.
	//
	// - **10GBase-LR**: 10 GE single-mode optical port (10 km).
	//
	// - **40GBase-LR**: 40 GE single-mode optical port.
	//
	// - **100GBase-LR**: 100 GE single-mode optical port.
	//
	// >  The creation of 40GBase-LR and 100GBase-LR ports depends on the actual backend port availability. Contact your account manager for details.
	//
	// example:
	//
	// 1000Base-T
	SupportPortTypes *string `json:"SupportPortTypes,omitempty" xml:"SupportPortTypes,omitempty"`
}

func (s ListBusinessAccessPointsResponseBodyBusinessAccessPoints) String() string {
	return dara.Prettify(s)
}

func (s ListBusinessAccessPointsResponseBodyBusinessAccessPoints) GoString() string {
	return s.String()
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) GetAccessPointName() *string {
	return s.AccessPointName
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) GetCloudBoxInstanceIds() *string {
	return s.CloudBoxInstanceIds
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) GetLatitude() *float64 {
	return s.Latitude
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) GetLongitude() *float64 {
	return s.Longitude
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) GetOpticalModuleModels() []*ListBusinessAccessPointsResponseBodyBusinessAccessPointsOpticalModuleModels {
	return s.OpticalModuleModels
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) GetSupportLineOperator() *string {
	return s.SupportLineOperator
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) GetSupportPortTypes() *string {
	return s.SupportPortTypes
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) SetAccessPointId(v string) *ListBusinessAccessPointsResponseBodyBusinessAccessPoints {
	s.AccessPointId = &v
	return s
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) SetAccessPointName(v string) *ListBusinessAccessPointsResponseBodyBusinessAccessPoints {
	s.AccessPointName = &v
	return s
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) SetCloudBoxInstanceIds(v string) *ListBusinessAccessPointsResponseBodyBusinessAccessPoints {
	s.CloudBoxInstanceIds = &v
	return s
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) SetLatitude(v float64) *ListBusinessAccessPointsResponseBodyBusinessAccessPoints {
	s.Latitude = &v
	return s
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) SetLongitude(v float64) *ListBusinessAccessPointsResponseBodyBusinessAccessPoints {
	s.Longitude = &v
	return s
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) SetOpticalModuleModels(v []*ListBusinessAccessPointsResponseBodyBusinessAccessPointsOpticalModuleModels) *ListBusinessAccessPointsResponseBodyBusinessAccessPoints {
	s.OpticalModuleModels = v
	return s
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) SetSupportLineOperator(v string) *ListBusinessAccessPointsResponseBodyBusinessAccessPoints {
	s.SupportLineOperator = &v
	return s
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) SetSupportPortTypes(v string) *ListBusinessAccessPointsResponseBodyBusinessAccessPoints {
	s.SupportPortTypes = &v
	return s
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) Validate() error {
	if s.OpticalModuleModels != nil {
		for _, item := range s.OpticalModuleModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBusinessAccessPointsResponseBodyBusinessAccessPointsOpticalModuleModels struct {
	// The optical module model supported by the Express Connect circuit access point. Valid values:
	//
	// 1000Base-LX :
	//
	// SFP-GE-LR-SM1310,10KM
	//
	// SFP-GE-ER-SM1310,40KM
	//
	// SFP-GE-ZR-SM1550,80KM
	//
	// 10GBase-LR :
	//
	// SFP-10G-LR-SM1310,10KM
	//
	// SFP-10G-ER-SM1550,40KM
	//
	// SFP-10G-ZR-SM1550,80KM
	//
	// 40GBase-LR :
	//
	// QSFP-40G-LR4-WDM1300,10KM
	//
	// QSFP-40G-ER4-WDM1300,40KM
	//
	// QSFP-40G-ZR4-WDM1300,80KM
	//
	// 100GBase-LR :
	//
	// QSFP28-100G-LR4-WDM1300,10KM
	//
	// QSFP28-100G-ER4-WDM1300,40KM
	//
	// QSFP28-100G-ZR4-WDM1300,80KM.
	//
	// example:
	//
	// SFP-GE-LR-SM1310,10KM
	OpticalModuleModel *string `json:"OpticalModuleModel,omitempty" xml:"OpticalModuleModel,omitempty"`
	// The port type supported by the optical module at the Express Connect circuit access point. Valid values:
	//
	// ● 1000Base-LX: GE single-mode optical port.
	//
	// ● 10GBase-LR: 10 GE single-mode optical port.
	//
	// ● 40GBase-LR: 40 GE single-mode optical port.
	//
	// ● 100GBase-LR: 100 GE single-mode optical port.
	//
	// example:
	//
	// 1000Base-LX
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
}

func (s ListBusinessAccessPointsResponseBodyBusinessAccessPointsOpticalModuleModels) String() string {
	return dara.Prettify(s)
}

func (s ListBusinessAccessPointsResponseBodyBusinessAccessPointsOpticalModuleModels) GoString() string {
	return s.String()
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPointsOpticalModuleModels) GetOpticalModuleModel() *string {
	return s.OpticalModuleModel
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPointsOpticalModuleModels) GetPortType() *string {
	return s.PortType
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPointsOpticalModuleModels) SetOpticalModuleModel(v string) *ListBusinessAccessPointsResponseBodyBusinessAccessPointsOpticalModuleModels {
	s.OpticalModuleModel = &v
	return s
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPointsOpticalModuleModels) SetPortType(v string) *ListBusinessAccessPointsResponseBodyBusinessAccessPointsOpticalModuleModels {
	s.PortType = &v
	return s
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPointsOpticalModuleModels) Validate() error {
	return dara.Validate(s)
}
