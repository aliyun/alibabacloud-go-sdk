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
	// The list of access points.
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
	return dara.Validate(s)
}

type ListBusinessAccessPointsResponseBodyBusinessAccessPoints struct {
	// The ID of the access point.
	//
	// example:
	//
	// ap-cn-hangzhou-xs-B
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The name of the access point.
	//
	// example:
	//
	// hangzhou-xs-B
	AccessPointName *string `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	// The ID of the cloud box.
	//
	// >  You can query this parameter if the Express Connect circuits and access points are of the cloud box type.
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
	// The connectivity provider of the Express Connect circuit. Valid values:
	//
	// 	- **CT**: China Telecom.
	//
	// 	- **CU**: China Unicom.
	//
	// 	- **CM**: China Mobile.
	//
	// 	- **CO**: other connectivity providers in the Chinese mainland.
	//
	// 	- **Equinix**: Equinix.
	//
	// 	- **Other**: other connectivity providers outside the Chinese mainland.
	//
	// example:
	//
	// CT
	SupportLineOperator *string `json:"SupportLineOperator,omitempty" xml:"SupportLineOperator,omitempty"`
	// The port type supported by the access point. Valid values:
	//
	// 	- **100Base-T**: 100 Mbit/s copper Ethernet port
	//
	// 	- **1000Base-T**: 1,000 Mbit/s copper Ethernet port
	//
	// 	- **1000Base-LX**: 1,000 Mbit/s single-mode optical port (10 km)
	//
	// 	- **10GBase-T**: 10,000 Mbit/s copper Ethernet port
	//
	// 	- **10GBase-LR**: 10,000 Mbit/s single-mode optical port (10 km)
	//
	// 	- **40GBase-LR**: 40,000 Mbit/s single-mode optical port
	//
	// 	- **100GBase-LR**: 100,000 Mbit/s single-mode optical port
	//
	// >  To use ports 40GBase-LR and 100GBase-LR, you must first contact your account manager.
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

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) SetSupportLineOperator(v string) *ListBusinessAccessPointsResponseBodyBusinessAccessPoints {
	s.SupportLineOperator = &v
	return s
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) SetSupportPortTypes(v string) *ListBusinessAccessPointsResponseBodyBusinessAccessPoints {
	s.SupportPortTypes = &v
	return s
}

func (s *ListBusinessAccessPointsResponseBodyBusinessAccessPoints) Validate() error {
	return dara.Validate(s)
}
