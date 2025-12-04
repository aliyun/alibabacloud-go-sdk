// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpdRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListVpdRouteEntriesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListVpdRouteEntriesResponseBody
	GetCode() *int32
	SetContent(v *ListVpdRouteEntriesResponseBodyContent) *ListVpdRouteEntriesResponseBody
	GetContent() *ListVpdRouteEntriesResponseBodyContent
	SetMessage(v string) *ListVpdRouteEntriesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVpdRouteEntriesResponseBody
	GetRequestId() *string
}

type ListVpdRouteEntriesResponseBody struct {
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
	Content *ListVpdRouteEntriesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// 039C3C3A-3C37-5672-80D5-D8CD48C676D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVpdRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpdRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpdRouteEntriesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListVpdRouteEntriesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListVpdRouteEntriesResponseBody) GetContent() *ListVpdRouteEntriesResponseBodyContent {
	return s.Content
}

func (s *ListVpdRouteEntriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVpdRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpdRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *ListVpdRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBody) SetCode(v int32) *ListVpdRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBody) SetContent(v *ListVpdRouteEntriesResponseBodyContent) *ListVpdRouteEntriesResponseBody {
	s.Content = v
	return s
}

func (s *ListVpdRouteEntriesResponseBody) SetMessage(v string) *ListVpdRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBody) SetRequestId(v string) *ListVpdRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVpdRouteEntriesResponseBodyContent struct {
	// Lingjun CIDR block route entry list
	Data []*ListVpdRouteEntriesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVpdRouteEntriesResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListVpdRouteEntriesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVpdRouteEntriesResponseBodyContent) GetData() []*ListVpdRouteEntriesResponseBodyContentData {
	return s.Data
}

func (s *ListVpdRouteEntriesResponseBodyContent) GetTotal() *int64 {
	return s.Total
}

func (s *ListVpdRouteEntriesResponseBodyContent) SetData(v []*ListVpdRouteEntriesResponseBodyContentData) *ListVpdRouteEntriesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContent) SetTotal(v int64) *ListVpdRouteEntriesResponseBodyContent {
	s.Total = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContent) Validate() error {
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

type ListVpdRouteEntriesResponseBodyContentData struct {
	// Destination CIDR block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 1678273219000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// er-bmlqiym1
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// Next Hop Instance Type
	//
	// example:
	//
	// ER
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
	// rg-acfmxhucx5ewuwy
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
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-eoiy88ju
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The ID of the route entry.
	//
	// example:
	//
	// vpd-rte-toekyqel
	VpdRouteEntryId *string `json:"VpdRouteEntryId,omitempty" xml:"VpdRouteEntryId,omitempty"`
}

func (s ListVpdRouteEntriesResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s ListVpdRouteEntriesResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListVpdRouteEntriesResponseBodyContentData) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ListVpdRouteEntriesResponseBodyContentData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListVpdRouteEntriesResponseBodyContentData) GetNextHopId() *string {
	return s.NextHopId
}

func (s *ListVpdRouteEntriesResponseBodyContentData) GetNextHopType() *string {
	return s.NextHopType
}

func (s *ListVpdRouteEntriesResponseBodyContentData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpdRouteEntriesResponseBodyContentData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpdRouteEntriesResponseBodyContentData) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *ListVpdRouteEntriesResponseBodyContentData) GetRouteType() *string {
	return s.RouteType
}

func (s *ListVpdRouteEntriesResponseBodyContentData) GetStatus() *string {
	return s.Status
}

func (s *ListVpdRouteEntriesResponseBodyContentData) GetTenantId() *string {
	return s.TenantId
}

func (s *ListVpdRouteEntriesResponseBodyContentData) GetVpdId() *string {
	return s.VpdId
}

func (s *ListVpdRouteEntriesResponseBodyContentData) GetVpdRouteEntryId() *string {
	return s.VpdRouteEntryId
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetDestinationCidrBlock(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetGmtModified(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetNextHopId(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.NextHopId = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetNextHopType(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.NextHopType = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetRegionId(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetResourceGroupId(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetResourceTenantId(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.ResourceTenantId = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetRouteType(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.RouteType = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetStatus(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetTenantId(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetVpdId(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.VpdId = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) SetVpdRouteEntryId(v string) *ListVpdRouteEntriesResponseBodyContentData {
	s.VpdRouteEntryId = &v
	return s
}

func (s *ListVpdRouteEntriesResponseBodyContentData) Validate() error {
	return dara.Validate(s)
}
