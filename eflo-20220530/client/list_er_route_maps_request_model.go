// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListErRouteMapsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidrBlock(v string) *ListErRouteMapsRequest
	GetDestinationCidrBlock() *string
	SetEnablePage(v bool) *ListErRouteMapsRequest
	GetEnablePage() *bool
	SetErId(v string) *ListErRouteMapsRequest
	GetErId() *string
	SetErRouteMapId(v string) *ListErRouteMapsRequest
	GetErRouteMapId() *string
	SetErRouteMapNum(v int32) *ListErRouteMapsRequest
	GetErRouteMapNum() *int32
	SetPageNumber(v int32) *ListErRouteMapsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListErRouteMapsRequest
	GetPageSize() *int32
	SetReceptionInstanceId(v string) *ListErRouteMapsRequest
	GetReceptionInstanceId() *string
	SetReceptionInstanceName(v string) *ListErRouteMapsRequest
	GetReceptionInstanceName() *string
	SetReceptionInstanceType(v string) *ListErRouteMapsRequest
	GetReceptionInstanceType() *string
	SetRegionId(v string) *ListErRouteMapsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListErRouteMapsRequest
	GetResourceGroupId() *string
	SetRouteMapAction(v string) *ListErRouteMapsRequest
	GetRouteMapAction() *string
	SetTransmissionInstanceId(v string) *ListErRouteMapsRequest
	GetTransmissionInstanceId() *string
	SetTransmissionInstanceName(v string) *ListErRouteMapsRequest
	GetTransmissionInstanceName() *string
	SetTransmissionInstanceType(v string) *ListErRouteMapsRequest
	GetTransmissionInstanceType() *string
}

type ListErRouteMapsRequest struct {
	// Destination CIDR Block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to enable paged query.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// Elastic Router ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// routing policy ID
	//
	// example:
	//
	// er-rmap-uwglhzom
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	// Policy number (default for automatic creation is 3000; The value range of the policy number manually created by the user is 1001-2000)
	//
	// example:
	//
	// 1001
	ErRouteMapNum *int32 `json:"ErRouteMapNum,omitempty" xml:"ErRouteMapNum,omitempty"`
	// The page number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Receive Instance ID
	//
	// example:
	//
	// vpd-x2lohgpv
	ReceptionInstanceId *string `json:"ReceptionInstanceId,omitempty" xml:"ReceptionInstanceId,omitempty"`
	// Receive Instance Name
	//
	// example:
	//
	// vpd2
	ReceptionInstanceName *string `json:"ReceptionInstanceName,omitempty" xml:"ReceptionInstanceName,omitempty"`
	// The type of the received instance. Optional values:
	//
	// 	- **VPD**: Lingjun network segment.
	//
	// 	- **VCC**: Lingjun Connection.
	//
	// example:
	//
	// VPD
	ReceptionInstanceType *string `json:"ReceptionInstanceType,omitempty" xml:"ReceptionInstanceType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmzaq3ypaqkdy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Policy behavior; optional values:
	//
	// 	- **permit**: Allow
	//
	// 	- **deny**: Rejected
	//
	// example:
	//
	// deny
	RouteMapAction *string `json:"RouteMapAction,omitempty" xml:"RouteMapAction,omitempty"`
	// Release Instance ID
	//
	// example:
	//
	// vpd-xsdlg2xb
	TransmissionInstanceId *string `json:"TransmissionInstanceId,omitempty" xml:"TransmissionInstanceId,omitempty"`
	// Release Instance Name
	//
	// example:
	//
	// vpd1
	TransmissionInstanceName *string `json:"TransmissionInstanceName,omitempty" xml:"TransmissionInstanceName,omitempty"`
	// The type of the published instance. Optional values:
	//
	// 	- **VPD**: Lingjun network segment.
	//
	// 	- **VCC**: Lingjun Connection.
	//
	// example:
	//
	// VPD
	TransmissionInstanceType *string `json:"TransmissionInstanceType,omitempty" xml:"TransmissionInstanceType,omitempty"`
}

func (s ListErRouteMapsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListErRouteMapsRequest) GoString() string {
	return s.String()
}

func (s *ListErRouteMapsRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ListErRouteMapsRequest) GetEnablePage() *bool {
	return s.EnablePage
}

func (s *ListErRouteMapsRequest) GetErId() *string {
	return s.ErId
}

func (s *ListErRouteMapsRequest) GetErRouteMapId() *string {
	return s.ErRouteMapId
}

func (s *ListErRouteMapsRequest) GetErRouteMapNum() *int32 {
	return s.ErRouteMapNum
}

func (s *ListErRouteMapsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListErRouteMapsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListErRouteMapsRequest) GetReceptionInstanceId() *string {
	return s.ReceptionInstanceId
}

func (s *ListErRouteMapsRequest) GetReceptionInstanceName() *string {
	return s.ReceptionInstanceName
}

func (s *ListErRouteMapsRequest) GetReceptionInstanceType() *string {
	return s.ReceptionInstanceType
}

func (s *ListErRouteMapsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListErRouteMapsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListErRouteMapsRequest) GetRouteMapAction() *string {
	return s.RouteMapAction
}

func (s *ListErRouteMapsRequest) GetTransmissionInstanceId() *string {
	return s.TransmissionInstanceId
}

func (s *ListErRouteMapsRequest) GetTransmissionInstanceName() *string {
	return s.TransmissionInstanceName
}

func (s *ListErRouteMapsRequest) GetTransmissionInstanceType() *string {
	return s.TransmissionInstanceType
}

func (s *ListErRouteMapsRequest) SetDestinationCidrBlock(v string) *ListErRouteMapsRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListErRouteMapsRequest) SetEnablePage(v bool) *ListErRouteMapsRequest {
	s.EnablePage = &v
	return s
}

func (s *ListErRouteMapsRequest) SetErId(v string) *ListErRouteMapsRequest {
	s.ErId = &v
	return s
}

func (s *ListErRouteMapsRequest) SetErRouteMapId(v string) *ListErRouteMapsRequest {
	s.ErRouteMapId = &v
	return s
}

func (s *ListErRouteMapsRequest) SetErRouteMapNum(v int32) *ListErRouteMapsRequest {
	s.ErRouteMapNum = &v
	return s
}

func (s *ListErRouteMapsRequest) SetPageNumber(v int32) *ListErRouteMapsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListErRouteMapsRequest) SetPageSize(v int32) *ListErRouteMapsRequest {
	s.PageSize = &v
	return s
}

func (s *ListErRouteMapsRequest) SetReceptionInstanceId(v string) *ListErRouteMapsRequest {
	s.ReceptionInstanceId = &v
	return s
}

func (s *ListErRouteMapsRequest) SetReceptionInstanceName(v string) *ListErRouteMapsRequest {
	s.ReceptionInstanceName = &v
	return s
}

func (s *ListErRouteMapsRequest) SetReceptionInstanceType(v string) *ListErRouteMapsRequest {
	s.ReceptionInstanceType = &v
	return s
}

func (s *ListErRouteMapsRequest) SetRegionId(v string) *ListErRouteMapsRequest {
	s.RegionId = &v
	return s
}

func (s *ListErRouteMapsRequest) SetResourceGroupId(v string) *ListErRouteMapsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListErRouteMapsRequest) SetRouteMapAction(v string) *ListErRouteMapsRequest {
	s.RouteMapAction = &v
	return s
}

func (s *ListErRouteMapsRequest) SetTransmissionInstanceId(v string) *ListErRouteMapsRequest {
	s.TransmissionInstanceId = &v
	return s
}

func (s *ListErRouteMapsRequest) SetTransmissionInstanceName(v string) *ListErRouteMapsRequest {
	s.TransmissionInstanceName = &v
	return s
}

func (s *ListErRouteMapsRequest) SetTransmissionInstanceType(v string) *ListErRouteMapsRequest {
	s.TransmissionInstanceType = &v
	return s
}

func (s *ListErRouteMapsRequest) Validate() error {
	return dara.Validate(s)
}
