// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeApplicationsResponseBodyItems) *DescribeApplicationsResponseBody
	GetItems() *DescribeApplicationsResponseBodyItems
	SetPageNumber(v int32) *DescribeApplicationsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeApplicationsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeApplicationsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeApplicationsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeApplicationsResponseBody struct {
	Items *DescribeApplicationsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationsResponseBody) GetItems() *DescribeApplicationsResponseBodyItems {
	return s.Items
}

func (s *DescribeApplicationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApplicationsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeApplicationsResponseBody) SetItems(v *DescribeApplicationsResponseBodyItems) *DescribeApplicationsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeApplicationsResponseBody) SetPageNumber(v int32) *DescribeApplicationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApplicationsResponseBody) SetPageRecordCount(v int32) *DescribeApplicationsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeApplicationsResponseBody) SetRequestId(v string) *DescribeApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationsResponseBody) SetTotalRecordCount(v int32) *DescribeApplicationsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeApplicationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationsResponseBodyItems struct {
	Applications []*DescribeApplicationsResponseBodyItemsApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
}

func (s DescribeApplicationsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeApplicationsResponseBodyItems) GetApplications() []*DescribeApplicationsResponseBodyItemsApplications {
	return s.Applications
}

func (s *DescribeApplicationsResponseBodyItems) SetApplications(v []*DescribeApplicationsResponseBodyItemsApplications) *DescribeApplicationsResponseBodyItems {
	s.Applications = v
	return s
}

func (s *DescribeApplicationsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationsResponseBodyItemsApplications struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// supabase
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// example:
	//
	// 2025-03-25T09:37:10Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// myapp
	Description *string                                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	Endpoints   *DescribeApplicationsResponseBodyItemsApplicationsEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	// example:
	//
	// 1.0.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// 2025-06-25T09:37:10Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// pcs-**************
	PolarFSInstanceId *string `json:"PolarFSInstanceId,omitempty" xml:"PolarFSInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Activated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeApplicationsResponseBodyItemsApplications) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationsResponseBodyItemsApplications) GoString() string {
	return s.String()
}

func (s *DescribeApplicationsResponseBodyItemsApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationsResponseBodyItemsApplications) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *DescribeApplicationsResponseBodyItemsApplications) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeApplicationsResponseBodyItemsApplications) GetDescription() *string {
	return s.Description
}

func (s *DescribeApplicationsResponseBodyItemsApplications) GetEndpoints() *DescribeApplicationsResponseBodyItemsApplicationsEndpoints {
	return s.Endpoints
}

func (s *DescribeApplicationsResponseBodyItemsApplications) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeApplicationsResponseBodyItemsApplications) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeApplicationsResponseBodyItemsApplications) GetExpired() *string {
	return s.Expired
}

func (s *DescribeApplicationsResponseBodyItemsApplications) GetPayType() *string {
	return s.PayType
}

func (s *DescribeApplicationsResponseBodyItemsApplications) GetPolarFSInstanceId() *string {
	return s.PolarFSInstanceId
}

func (s *DescribeApplicationsResponseBodyItemsApplications) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationsResponseBodyItemsApplications) GetStatus() *string {
	return s.Status
}

func (s *DescribeApplicationsResponseBodyItemsApplications) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeApplicationsResponseBodyItemsApplications) SetApplicationId(v string) *DescribeApplicationsResponseBodyItemsApplications {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplications) SetApplicationType(v string) *DescribeApplicationsResponseBodyItemsApplications {
	s.ApplicationType = &v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplications) SetCreationTime(v string) *DescribeApplicationsResponseBodyItemsApplications {
	s.CreationTime = &v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplications) SetDescription(v string) *DescribeApplicationsResponseBodyItemsApplications {
	s.Description = &v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplications) SetEndpoints(v *DescribeApplicationsResponseBodyItemsApplicationsEndpoints) *DescribeApplicationsResponseBodyItemsApplications {
	s.Endpoints = v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplications) SetEngineVersion(v string) *DescribeApplicationsResponseBodyItemsApplications {
	s.EngineVersion = &v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplications) SetExpireTime(v string) *DescribeApplicationsResponseBodyItemsApplications {
	s.ExpireTime = &v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplications) SetExpired(v string) *DescribeApplicationsResponseBodyItemsApplications {
	s.Expired = &v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplications) SetPayType(v string) *DescribeApplicationsResponseBodyItemsApplications {
	s.PayType = &v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplications) SetPolarFSInstanceId(v string) *DescribeApplicationsResponseBodyItemsApplications {
	s.PolarFSInstanceId = &v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplications) SetRegionId(v string) *DescribeApplicationsResponseBodyItemsApplications {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplications) SetStatus(v string) *DescribeApplicationsResponseBodyItemsApplications {
	s.Status = &v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplications) SetZoneId(v string) *DescribeApplicationsResponseBodyItemsApplications {
	s.ZoneId = &v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplications) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationsResponseBodyItemsApplicationsEndpoints struct {
	Endpoint []*DescribeApplicationsResponseBodyItemsApplicationsEndpointsEndpoint `json:"endpoint,omitempty" xml:"endpoint,omitempty" type:"Repeated"`
}

func (s DescribeApplicationsResponseBodyItemsApplicationsEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationsResponseBodyItemsApplicationsEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeApplicationsResponseBodyItemsApplicationsEndpoints) GetEndpoint() []*DescribeApplicationsResponseBodyItemsApplicationsEndpointsEndpoint {
	return s.Endpoint
}

func (s *DescribeApplicationsResponseBodyItemsApplicationsEndpoints) SetEndpoint(v []*DescribeApplicationsResponseBodyItemsApplicationsEndpointsEndpoint) *DescribeApplicationsResponseBodyItemsApplicationsEndpoints {
	s.Endpoint = v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplicationsEndpoints) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationsResponseBodyItemsApplicationsEndpointsEndpoint struct {
	// example:
	//
	// 192.168.0.3
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// example:
	//
	// Public
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeApplicationsResponseBodyItemsApplicationsEndpointsEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationsResponseBodyItemsApplicationsEndpointsEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeApplicationsResponseBodyItemsApplicationsEndpointsEndpoint) GetIP() *string {
	return s.IP
}

func (s *DescribeApplicationsResponseBodyItemsApplicationsEndpointsEndpoint) GetNetType() *string {
	return s.NetType
}

func (s *DescribeApplicationsResponseBodyItemsApplicationsEndpointsEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeApplicationsResponseBodyItemsApplicationsEndpointsEndpoint) SetIP(v string) *DescribeApplicationsResponseBodyItemsApplicationsEndpointsEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplicationsEndpointsEndpoint) SetNetType(v string) *DescribeApplicationsResponseBodyItemsApplicationsEndpointsEndpoint {
	s.NetType = &v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplicationsEndpointsEndpoint) SetPort(v string) *DescribeApplicationsResponseBodyItemsApplicationsEndpointsEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeApplicationsResponseBodyItemsApplicationsEndpointsEndpoint) Validate() error {
	return dara.Validate(s)
}
