// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListErRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListErRouteEntriesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListErRouteEntriesResponseBody
	GetCode() *int32
	SetContent(v *ListErRouteEntriesResponseBodyContent) *ListErRouteEntriesResponseBody
	GetContent() *ListErRouteEntriesResponseBodyContent
	SetMessage(v string) *ListErRouteEntriesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListErRouteEntriesResponseBody
	GetRequestId() *string
}

type ListErRouteEntriesResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *ListErRouteEntriesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListErRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListErRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListErRouteEntriesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListErRouteEntriesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListErRouteEntriesResponseBody) GetContent() *ListErRouteEntriesResponseBodyContent {
	return s.Content
}

func (s *ListErRouteEntriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListErRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListErRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *ListErRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListErRouteEntriesResponseBody) SetCode(v int32) *ListErRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListErRouteEntriesResponseBody) SetContent(v *ListErRouteEntriesResponseBodyContent) *ListErRouteEntriesResponseBody {
	s.Content = v
	return s
}

func (s *ListErRouteEntriesResponseBody) SetMessage(v string) *ListErRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListErRouteEntriesResponseBody) SetRequestId(v string) *ListErRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListErRouteEntriesResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListErRouteEntriesResponseBodyContent struct {
	// Lingjun HUB Route Entry Information List
	Data []*ListErRouteEntriesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListErRouteEntriesResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListErRouteEntriesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListErRouteEntriesResponseBodyContent) GetData() []*ListErRouteEntriesResponseBodyContentData {
	return s.Data
}

func (s *ListErRouteEntriesResponseBodyContent) GetTotal() *int64 {
	return s.Total
}

func (s *ListErRouteEntriesResponseBodyContent) SetData(v []*ListErRouteEntriesResponseBodyContentData) *ListErRouteEntriesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListErRouteEntriesResponseBodyContent) SetTotal(v int64) *ListErRouteEntriesResponseBodyContent {
	s.Total = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContent) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListErRouteEntriesResponseBodyContentData struct {
	// Destination CIDR Block
	//
	// example:
	//
	// 100.64.1.100/32
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Lingjun HUB Instance ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The ID of the route entry.
	//
	// example:
	//
	// er-rte-maysfadg
	ErRouteEntryId *string `json:"ErRouteEntryId,omitempty" xml:"ErRouteEntryId,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 1640930901000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// vcc-cn-209300qha01
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// Next Hop Instance Type
	//
	// example:
	//
	// VCC
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aekzb3n5lgk2ieq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant to which the resource belongs.
	//
	// example:
	//
	// 1111156667137893
	ResourceTenantId *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	// Route type
	//
	// example:
	//
	// VCC
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The task status. Valid values:
	//
	// 	- Synchronizing
	//
	// 	- Available
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1111156667137893
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListErRouteEntriesResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s ListErRouteEntriesResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListErRouteEntriesResponseBodyContentData) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ListErRouteEntriesResponseBodyContentData) GetErId() *string {
	return s.ErId
}

func (s *ListErRouteEntriesResponseBodyContentData) GetErRouteEntryId() *string {
	return s.ErRouteEntryId
}

func (s *ListErRouteEntriesResponseBodyContentData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListErRouteEntriesResponseBodyContentData) GetNextHopId() *string {
	return s.NextHopId
}

func (s *ListErRouteEntriesResponseBodyContentData) GetNextHopType() *string {
	return s.NextHopType
}

func (s *ListErRouteEntriesResponseBodyContentData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListErRouteEntriesResponseBodyContentData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListErRouteEntriesResponseBodyContentData) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *ListErRouteEntriesResponseBodyContentData) GetRouteType() *string {
	return s.RouteType
}

func (s *ListErRouteEntriesResponseBodyContentData) GetStatus() *string {
	return s.Status
}

func (s *ListErRouteEntriesResponseBodyContentData) GetTenantId() *string {
	return s.TenantId
}

func (s *ListErRouteEntriesResponseBodyContentData) SetDestinationCidrBlock(v string) *ListErRouteEntriesResponseBodyContentData {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetErId(v string) *ListErRouteEntriesResponseBodyContentData {
	s.ErId = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetErRouteEntryId(v string) *ListErRouteEntriesResponseBodyContentData {
	s.ErRouteEntryId = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetGmtModified(v string) *ListErRouteEntriesResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetNextHopId(v string) *ListErRouteEntriesResponseBodyContentData {
	s.NextHopId = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetNextHopType(v string) *ListErRouteEntriesResponseBodyContentData {
	s.NextHopType = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetRegionId(v string) *ListErRouteEntriesResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetResourceGroupId(v string) *ListErRouteEntriesResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetResourceTenantId(v string) *ListErRouteEntriesResponseBodyContentData {
	s.ResourceTenantId = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetRouteType(v string) *ListErRouteEntriesResponseBodyContentData {
	s.RouteType = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetStatus(v string) *ListErRouteEntriesResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) SetTenantId(v string) *ListErRouteEntriesResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListErRouteEntriesResponseBodyContentData) Validate() error {
	return dara.Validate(s)
}
