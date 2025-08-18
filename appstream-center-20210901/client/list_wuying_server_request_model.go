// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWuyingServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *ListWuyingServerRequest
	GetBizRegionId() *string
	SetChargeType(v string) *ListWuyingServerRequest
	GetChargeType() *string
	SetImageId(v string) *ListWuyingServerRequest
	GetImageId() *string
	SetOfficeSiteId(v string) *ListWuyingServerRequest
	GetOfficeSiteId() *string
	SetPageNumber(v int32) *ListWuyingServerRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWuyingServerRequest
	GetPageSize() *int32
	SetServerInstanceType(v string) *ListWuyingServerRequest
	GetServerInstanceType() *string
	SetStatus(v string) *ListWuyingServerRequest
	GetStatus() *string
	SetWuyingServerIdList(v []*string) *ListWuyingServerRequest
	GetWuyingServerIdList() []*string
	SetWuyingServerNameOrId(v string) *ListWuyingServerRequest
	GetWuyingServerNameOrId() *string
}

type ListWuyingServerRequest struct {
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// img-bp13mu****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-172301****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// eds.proworkstation_flagship_elite.32c64g.48g1x
	ServerInstanceType *string `json:"ServerInstanceType,omitempty" xml:"ServerInstanceType,omitempty"`
	// example:
	//
	// RUNNING
	Status             *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	WuyingServerIdList []*string `json:"WuyingServerIdList,omitempty" xml:"WuyingServerIdList,omitempty" type:"Repeated"`
	// example:
	//
	// exampleServerName
	WuyingServerNameOrId *string `json:"WuyingServerNameOrId,omitempty" xml:"WuyingServerNameOrId,omitempty"`
}

func (s ListWuyingServerRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWuyingServerRequest) GoString() string {
	return s.String()
}

func (s *ListWuyingServerRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ListWuyingServerRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListWuyingServerRequest) GetImageId() *string {
	return s.ImageId
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

func (s *ListWuyingServerRequest) GetServerInstanceType() *string {
	return s.ServerInstanceType
}

func (s *ListWuyingServerRequest) GetStatus() *string {
	return s.Status
}

func (s *ListWuyingServerRequest) GetWuyingServerIdList() []*string {
	return s.WuyingServerIdList
}

func (s *ListWuyingServerRequest) GetWuyingServerNameOrId() *string {
	return s.WuyingServerNameOrId
}

func (s *ListWuyingServerRequest) SetBizRegionId(v string) *ListWuyingServerRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListWuyingServerRequest) SetChargeType(v string) *ListWuyingServerRequest {
	s.ChargeType = &v
	return s
}

func (s *ListWuyingServerRequest) SetImageId(v string) *ListWuyingServerRequest {
	s.ImageId = &v
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

func (s *ListWuyingServerRequest) SetServerInstanceType(v string) *ListWuyingServerRequest {
	s.ServerInstanceType = &v
	return s
}

func (s *ListWuyingServerRequest) SetStatus(v string) *ListWuyingServerRequest {
	s.Status = &v
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

func (s *ListWuyingServerRequest) Validate() error {
	return dara.Validate(s)
}
