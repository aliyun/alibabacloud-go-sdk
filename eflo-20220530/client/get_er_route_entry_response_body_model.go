// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetErRouteEntryResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetErRouteEntryResponseBody
	GetCode() *int32
	SetContent(v *GetErRouteEntryResponseBodyContent) *GetErRouteEntryResponseBody
	GetContent() *GetErRouteEntryResponseBodyContent
	SetMessage(v string) *GetErRouteEntryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetErRouteEntryResponseBody
	GetRequestId() *string
}

type GetErRouteEntryResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Content *GetErRouteEntryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// code: 400, Request was denied due to request throttling. request id: 7D177459-C1CF-5690-BB23-321D208B37D5
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 1F38A2E6-CB47-5369-95D2-96D0C287B4A5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetErRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetErRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *GetErRouteEntryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetErRouteEntryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetErRouteEntryResponseBody) GetContent() *GetErRouteEntryResponseBodyContent {
	return s.Content
}

func (s *GetErRouteEntryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetErRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetErRouteEntryResponseBody) SetAccessDeniedDetail(v string) *GetErRouteEntryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetErRouteEntryResponseBody) SetCode(v int32) *GetErRouteEntryResponseBody {
	s.Code = &v
	return s
}

func (s *GetErRouteEntryResponseBody) SetContent(v *GetErRouteEntryResponseBodyContent) *GetErRouteEntryResponseBody {
	s.Content = v
	return s
}

func (s *GetErRouteEntryResponseBody) SetMessage(v string) *GetErRouteEntryResponseBody {
	s.Message = &v
	return s
}

func (s *GetErRouteEntryResponseBody) SetRequestId(v string) *GetErRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetErRouteEntryResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetErRouteEntryResponseBodyContent struct {
	// Destination CIDR Block
	//
	// example:
	//
	// 11.0.0.0/16
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Lingjun HUB Instance ID
	//
	// example:
	//
	// er-aueyxxsy
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The ID of the route entry.
	//
	// example:
	//
	// er-rte-4q0jbylz
	ErRouteEntryId *string `json:"ErRouteEntryId,omitempty" xml:"ErRouteEntryId,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 1666677783000
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
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aekzb3n5lgk2ieq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Route type
	//
	// example:
	//
	// System
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
}

func (s GetErRouteEntryResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetErRouteEntryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetErRouteEntryResponseBodyContent) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *GetErRouteEntryResponseBodyContent) GetErId() *string {
	return s.ErId
}

func (s *GetErRouteEntryResponseBodyContent) GetErRouteEntryId() *string {
	return s.ErRouteEntryId
}

func (s *GetErRouteEntryResponseBodyContent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetErRouteEntryResponseBodyContent) GetNextHopId() *string {
	return s.NextHopId
}

func (s *GetErRouteEntryResponseBodyContent) GetNextHopType() *string {
	return s.NextHopType
}

func (s *GetErRouteEntryResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetErRouteEntryResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetErRouteEntryResponseBodyContent) GetRouteType() *string {
	return s.RouteType
}

func (s *GetErRouteEntryResponseBodyContent) GetStatus() *string {
	return s.Status
}

func (s *GetErRouteEntryResponseBodyContent) GetTenantId() *string {
	return s.TenantId
}

func (s *GetErRouteEntryResponseBodyContent) SetDestinationCidrBlock(v string) *GetErRouteEntryResponseBodyContent {
	s.DestinationCidrBlock = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetErId(v string) *GetErRouteEntryResponseBodyContent {
	s.ErId = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetErRouteEntryId(v string) *GetErRouteEntryResponseBodyContent {
	s.ErRouteEntryId = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetGmtModified(v string) *GetErRouteEntryResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetNextHopId(v string) *GetErRouteEntryResponseBodyContent {
	s.NextHopId = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetNextHopType(v string) *GetErRouteEntryResponseBodyContent {
	s.NextHopType = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetRegionId(v string) *GetErRouteEntryResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetResourceGroupId(v string) *GetErRouteEntryResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetRouteType(v string) *GetErRouteEntryResponseBodyContent {
	s.RouteType = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetStatus(v string) *GetErRouteEntryResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) SetTenantId(v string) *GetErRouteEntryResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetErRouteEntryResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
