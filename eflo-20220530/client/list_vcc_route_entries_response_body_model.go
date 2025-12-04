// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVccRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListVccRouteEntriesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListVccRouteEntriesResponseBody
	GetCode() *int32
	SetContent(v *ListVccRouteEntriesResponseBodyContent) *ListVccRouteEntriesResponseBody
	GetContent() *ListVccRouteEntriesResponseBodyContent
	SetMessage(v string) *ListVccRouteEntriesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVccRouteEntriesResponseBody
	GetRequestId() *string
}

type ListVccRouteEntriesResponseBody struct {
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
	Content *ListVccRouteEntriesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// response message, if the success request is
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

func (s ListVccRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVccRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVccRouteEntriesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListVccRouteEntriesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListVccRouteEntriesResponseBody) GetContent() *ListVccRouteEntriesResponseBodyContent {
	return s.Content
}

func (s *ListVccRouteEntriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVccRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVccRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *ListVccRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListVccRouteEntriesResponseBody) SetCode(v int32) *ListVccRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListVccRouteEntriesResponseBody) SetContent(v *ListVccRouteEntriesResponseBodyContent) *ListVccRouteEntriesResponseBody {
	s.Content = v
	return s
}

func (s *ListVccRouteEntriesResponseBody) SetMessage(v string) *ListVccRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListVccRouteEntriesResponseBody) SetRequestId(v string) *ListVccRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVccRouteEntriesResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVccRouteEntriesResponseBodyContent struct {
	// List of Lingjun Connection Route Entries
	Data []*ListVccRouteEntriesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVccRouteEntriesResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListVccRouteEntriesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVccRouteEntriesResponseBodyContent) GetData() []*ListVccRouteEntriesResponseBodyContentData {
	return s.Data
}

func (s *ListVccRouteEntriesResponseBodyContent) GetTotal() *int64 {
	return s.Total
}

func (s *ListVccRouteEntriesResponseBodyContent) SetData(v []*ListVccRouteEntriesResponseBodyContentData) *ListVccRouteEntriesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContent) SetTotal(v int64) *ListVccRouteEntriesResponseBodyContent {
	s.Total = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContent) Validate() error {
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

type ListVccRouteEntriesResponseBodyContentData struct {
	// Destination CIDR block
	//
	// example:
	//
	// 10.192.32.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 1642745758000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// Next Hop Type
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
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant to which the resource belongs.
	//
	// example:
	//
	// 1655449505171
	ResourceTenantId *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	// Route type
	//
	// example:
	//
	// BGP
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// The ID of the route entry.
	//
	// example:
	//
	// vcc-rte-maysfadg
	VccRouteEntryId *string `json:"VccRouteEntryId,omitempty" xml:"VccRouteEntryId,omitempty"`
}

func (s ListVccRouteEntriesResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s ListVccRouteEntriesResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListVccRouteEntriesResponseBodyContentData) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ListVccRouteEntriesResponseBodyContentData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListVccRouteEntriesResponseBodyContentData) GetMessage() *string {
	return s.Message
}

func (s *ListVccRouteEntriesResponseBodyContentData) GetNextHopId() *string {
	return s.NextHopId
}

func (s *ListVccRouteEntriesResponseBodyContentData) GetNextHopType() *string {
	return s.NextHopType
}

func (s *ListVccRouteEntriesResponseBodyContentData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVccRouteEntriesResponseBodyContentData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVccRouteEntriesResponseBodyContentData) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *ListVccRouteEntriesResponseBodyContentData) GetRouteType() *string {
	return s.RouteType
}

func (s *ListVccRouteEntriesResponseBodyContentData) GetStatus() *string {
	return s.Status
}

func (s *ListVccRouteEntriesResponseBodyContentData) GetTenantId() *string {
	return s.TenantId
}

func (s *ListVccRouteEntriesResponseBodyContentData) GetVccId() *string {
	return s.VccId
}

func (s *ListVccRouteEntriesResponseBodyContentData) GetVccRouteEntryId() *string {
	return s.VccRouteEntryId
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetDestinationCidrBlock(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetGmtModified(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetMessage(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetNextHopId(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.NextHopId = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetNextHopType(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.NextHopType = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetRegionId(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetResourceGroupId(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetResourceTenantId(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.ResourceTenantId = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetRouteType(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.RouteType = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetStatus(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetTenantId(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetVccId(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.VccId = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) SetVccRouteEntryId(v string) *ListVccRouteEntriesResponseBodyContentData {
	s.VccRouteEntryId = &v
	return s
}

func (s *ListVccRouteEntriesResponseBodyContentData) Validate() error {
	return dara.Validate(s)
}
