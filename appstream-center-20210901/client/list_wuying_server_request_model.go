// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWuyingServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddVirtualNodePoolStatusList(v []*string) *ListWuyingServerRequest
	GetAddVirtualNodePoolStatusList() []*string
	SetBizRegionId(v string) *ListWuyingServerRequest
	GetBizRegionId() *string
	SetBizType(v int32) *ListWuyingServerRequest
	GetBizType() *int32
	SetChargeType(v string) *ListWuyingServerRequest
	GetChargeType() *string
	SetCreateTimeEnd(v string) *ListWuyingServerRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *ListWuyingServerRequest
	GetCreateTimeStart() *string
	SetExpiredTimeEnd(v string) *ListWuyingServerRequest
	GetExpiredTimeEnd() *string
	SetExpiredTimeStart(v string) *ListWuyingServerRequest
	GetExpiredTimeStart() *string
	SetImageId(v string) *ListWuyingServerRequest
	GetImageId() *string
	SetNetworkInterfaceIp(v string) *ListWuyingServerRequest
	GetNetworkInterfaceIp() *string
	SetOfficeSiteId(v string) *ListWuyingServerRequest
	GetOfficeSiteId() *string
	SetPageNumber(v int32) *ListWuyingServerRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWuyingServerRequest
	GetPageSize() *int32
	SetProductType(v string) *ListWuyingServerRequest
	GetProductType() *string
	SetServerInstanceType(v string) *ListWuyingServerRequest
	GetServerInstanceType() *string
	SetStatus(v string) *ListWuyingServerRequest
	GetStatus() *string
	SetUsers(v []*string) *ListWuyingServerRequest
	GetUsers() []*string
	SetVirtualNodePoolId(v string) *ListWuyingServerRequest
	GetVirtualNodePoolId() *string
	SetWuyingServerIdList(v []*string) *ListWuyingServerRequest
	GetWuyingServerIdList() []*string
	SetWuyingServerNameOrId(v string) *ListWuyingServerRequest
	GetWuyingServerNameOrId() *string
	SetZoneId(v string) *ListWuyingServerRequest
	GetZoneId() *string
}

type ListWuyingServerRequest struct {
	// The list of statuses for joining a virtual node pool.
	//
	// example:
	//
	// RUNNING
	AddVirtualNodePoolStatusList []*string `json:"AddVirtualNodePoolStatusList,omitempty" xml:"AddVirtualNodePoolStatusList,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The business type.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The billing type.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The end time of the creation time range, in ISO 8601 format. This time point is exclusive.
	//
	// example:
	//
	// 2026-08-01T00:00:00Z
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// The start time of the creation time range, in ISO 8601 format. This time point is inclusive.
	//
	// example:
	//
	// 2026-07-01T00:00:00Z
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// The end time of the expiration time range, in ISO 8601 format. This time point is exclusive.
	//
	// example:
	//
	// 2026-08-01T00:00:00Z
	ExpiredTimeEnd *string `json:"ExpiredTimeEnd,omitempty" xml:"ExpiredTimeEnd,omitempty"`
	// The start time of the expiration time range, in ISO 8601 format. This time point is inclusive.
	//
	// example:
	//
	// 2026-07-01T00:00:00Z
	ExpiredTimeStart *string `json:"ExpiredTimeStart,omitempty" xml:"ExpiredTimeStart,omitempty"`
	// The image ID.
	//
	// example:
	//
	// img-bp13mu****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The internal IP address.
	//
	// example:
	//
	// 10.31.1.1
	NetworkInterfaceIp *string `json:"NetworkInterfaceIp,omitempty" xml:"NetworkInterfaceIp,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-172301****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type.
	//
	// example:
	//
	// wuying_server
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The workstation instance type.
	//
	// example:
	//
	// eds.proworkstation_flagship_elite.32c64g.48g1x
	ServerInstanceType *string `json:"ServerInstanceType,omitempty" xml:"ServerInstanceType,omitempty"`
	// The workstation status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of authorized users.
	//
	// example:
	//
	// user1
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
	// The virtual node pool ID.
	//
	// example:
	//
	// vnp-bp1234567890abcde
	VirtualNodePoolId *string `json:"VirtualNodePoolId,omitempty" xml:"VirtualNodePoolId,omitempty"`
	// The list of workstation IDs.
	//
	// example:
	//
	// 1
	WuyingServerIdList []*string `json:"WuyingServerIdList,omitempty" xml:"WuyingServerIdList,omitempty" type:"Repeated"`
	// The workstation name or workstation ID.
	//
	// example:
	//
	// exampleServerName
	WuyingServerNameOrId *string `json:"WuyingServerNameOrId,omitempty" xml:"WuyingServerNameOrId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListWuyingServerRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWuyingServerRequest) GoString() string {
	return s.String()
}

func (s *ListWuyingServerRequest) GetAddVirtualNodePoolStatusList() []*string {
	return s.AddVirtualNodePoolStatusList
}

func (s *ListWuyingServerRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ListWuyingServerRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ListWuyingServerRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListWuyingServerRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *ListWuyingServerRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *ListWuyingServerRequest) GetExpiredTimeEnd() *string {
	return s.ExpiredTimeEnd
}

func (s *ListWuyingServerRequest) GetExpiredTimeStart() *string {
	return s.ExpiredTimeStart
}

func (s *ListWuyingServerRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ListWuyingServerRequest) GetNetworkInterfaceIp() *string {
	return s.NetworkInterfaceIp
}

func (s *ListWuyingServerRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ListWuyingServerRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWuyingServerRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWuyingServerRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListWuyingServerRequest) GetServerInstanceType() *string {
	return s.ServerInstanceType
}

func (s *ListWuyingServerRequest) GetStatus() *string {
	return s.Status
}

func (s *ListWuyingServerRequest) GetUsers() []*string {
	return s.Users
}

func (s *ListWuyingServerRequest) GetVirtualNodePoolId() *string {
	return s.VirtualNodePoolId
}

func (s *ListWuyingServerRequest) GetWuyingServerIdList() []*string {
	return s.WuyingServerIdList
}

func (s *ListWuyingServerRequest) GetWuyingServerNameOrId() *string {
	return s.WuyingServerNameOrId
}

func (s *ListWuyingServerRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListWuyingServerRequest) SetAddVirtualNodePoolStatusList(v []*string) *ListWuyingServerRequest {
	s.AddVirtualNodePoolStatusList = v
	return s
}

func (s *ListWuyingServerRequest) SetBizRegionId(v string) *ListWuyingServerRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListWuyingServerRequest) SetBizType(v int32) *ListWuyingServerRequest {
	s.BizType = &v
	return s
}

func (s *ListWuyingServerRequest) SetChargeType(v string) *ListWuyingServerRequest {
	s.ChargeType = &v
	return s
}

func (s *ListWuyingServerRequest) SetCreateTimeEnd(v string) *ListWuyingServerRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListWuyingServerRequest) SetCreateTimeStart(v string) *ListWuyingServerRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListWuyingServerRequest) SetExpiredTimeEnd(v string) *ListWuyingServerRequest {
	s.ExpiredTimeEnd = &v
	return s
}

func (s *ListWuyingServerRequest) SetExpiredTimeStart(v string) *ListWuyingServerRequest {
	s.ExpiredTimeStart = &v
	return s
}

func (s *ListWuyingServerRequest) SetImageId(v string) *ListWuyingServerRequest {
	s.ImageId = &v
	return s
}

func (s *ListWuyingServerRequest) SetNetworkInterfaceIp(v string) *ListWuyingServerRequest {
	s.NetworkInterfaceIp = &v
	return s
}

func (s *ListWuyingServerRequest) SetOfficeSiteId(v string) *ListWuyingServerRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ListWuyingServerRequest) SetPageNumber(v int32) *ListWuyingServerRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWuyingServerRequest) SetPageSize(v int32) *ListWuyingServerRequest {
	s.PageSize = &v
	return s
}

func (s *ListWuyingServerRequest) SetProductType(v string) *ListWuyingServerRequest {
	s.ProductType = &v
	return s
}

func (s *ListWuyingServerRequest) SetServerInstanceType(v string) *ListWuyingServerRequest {
	s.ServerInstanceType = &v
	return s
}

func (s *ListWuyingServerRequest) SetStatus(v string) *ListWuyingServerRequest {
	s.Status = &v
	return s
}

func (s *ListWuyingServerRequest) SetUsers(v []*string) *ListWuyingServerRequest {
	s.Users = v
	return s
}

func (s *ListWuyingServerRequest) SetVirtualNodePoolId(v string) *ListWuyingServerRequest {
	s.VirtualNodePoolId = &v
	return s
}

func (s *ListWuyingServerRequest) SetWuyingServerIdList(v []*string) *ListWuyingServerRequest {
	s.WuyingServerIdList = v
	return s
}

func (s *ListWuyingServerRequest) SetWuyingServerNameOrId(v string) *ListWuyingServerRequest {
	s.WuyingServerNameOrId = &v
	return s
}

func (s *ListWuyingServerRequest) SetZoneId(v string) *ListWuyingServerRequest {
	s.ZoneId = &v
	return s
}

func (s *ListWuyingServerRequest) Validate() error {
	return dara.Validate(s)
}
