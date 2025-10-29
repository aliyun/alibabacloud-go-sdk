// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListRoutesResponseBodyPagingInfo) *ListRoutesResponseBody
	GetPagingInfo() *ListRoutesResponseBodyPagingInfo
	SetRequestId(v string) *ListRoutesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRoutesResponseBody
	GetSuccess() *bool
}

type ListRoutesResponseBody struct {
	// The pagination information.
	PagingInfo *ListRoutesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The ID of the request. It is used to locate logs and troubleshoot problems.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRoutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoutesResponseBody) GetPagingInfo() *ListRoutesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListRoutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRoutesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRoutesResponseBody) SetPagingInfo(v *ListRoutesResponseBodyPagingInfo) *ListRoutesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListRoutesResponseBody) SetRequestId(v string) *ListRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoutesResponseBody) SetSuccess(v bool) *ListRoutesResponseBody {
	s.Success = &v
	return s
}

func (s *ListRoutesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRoutesResponseBodyPagingInfo struct {
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
	// The list of network resource routing information obtained.
	RouteList []*ListRoutesResponseBodyPagingInfoRouteList `json:"RouteList,omitempty" xml:"RouteList,omitempty" type:"Repeated"`
	// All data entries
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRoutesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListRoutesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListRoutesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRoutesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRoutesResponseBodyPagingInfo) GetRouteList() []*ListRoutesResponseBodyPagingInfoRouteList {
	return s.RouteList
}

func (s *ListRoutesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRoutesResponseBodyPagingInfo) SetPageNumber(v int32) *ListRoutesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListRoutesResponseBodyPagingInfo) SetPageSize(v int32) *ListRoutesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListRoutesResponseBodyPagingInfo) SetRouteList(v []*ListRoutesResponseBodyPagingInfoRouteList) *ListRoutesResponseBodyPagingInfo {
	s.RouteList = v
	return s
}

func (s *ListRoutesResponseBodyPagingInfo) SetTotalCount(v int32) *ListRoutesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListRoutesResponseBodyPagingInfo) Validate() error {
	if s.RouteList != nil {
		for _, item := range s.RouteList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRoutesResponseBodyPagingInfoRouteList struct {
	// The creation time, which is a 64-bit timestamp.
	//
	// example:
	//
	// 1727055811000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Route destination CIDR
	//
	// example:
	//
	// 192.168.0.0/16
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// Route ID
	//
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Network Resource ID
	//
	// example:
	//
	// 1000
	NetworkId *int64 `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// Unique identifier of the resource group to which it belongs
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Unique identifier of network resource
	//
	// example:
	//
	// ns-679XXXXXX
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s ListRoutesResponseBodyPagingInfoRouteList) String() string {
	return dara.Prettify(s)
}

func (s ListRoutesResponseBodyPagingInfoRouteList) GoString() string {
	return s.String()
}

func (s *ListRoutesResponseBodyPagingInfoRouteList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListRoutesResponseBodyPagingInfoRouteList) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *ListRoutesResponseBodyPagingInfoRouteList) GetId() *int64 {
	return s.Id
}

func (s *ListRoutesResponseBodyPagingInfoRouteList) GetNetworkId() *int64 {
	return s.NetworkId
}

func (s *ListRoutesResponseBodyPagingInfoRouteList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListRoutesResponseBodyPagingInfoRouteList) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListRoutesResponseBodyPagingInfoRouteList) SetCreateTime(v int64) *ListRoutesResponseBodyPagingInfoRouteList {
	s.CreateTime = &v
	return s
}

func (s *ListRoutesResponseBodyPagingInfoRouteList) SetDestinationCidr(v string) *ListRoutesResponseBodyPagingInfoRouteList {
	s.DestinationCidr = &v
	return s
}

func (s *ListRoutesResponseBodyPagingInfoRouteList) SetId(v int64) *ListRoutesResponseBodyPagingInfoRouteList {
	s.Id = &v
	return s
}

func (s *ListRoutesResponseBodyPagingInfoRouteList) SetNetworkId(v int64) *ListRoutesResponseBodyPagingInfoRouteList {
	s.NetworkId = &v
	return s
}

func (s *ListRoutesResponseBodyPagingInfoRouteList) SetResourceGroupId(v string) *ListRoutesResponseBodyPagingInfoRouteList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListRoutesResponseBodyPagingInfoRouteList) SetResourceId(v string) *ListRoutesResponseBodyPagingInfoRouteList {
	s.ResourceId = &v
	return s
}

func (s *ListRoutesResponseBodyPagingInfoRouteList) Validate() error {
	return dara.Validate(s)
}
