// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVccRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetVccRouteEntryResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetVccRouteEntryResponseBody
	GetCode() *int32
	SetContent(v *GetVccRouteEntryResponseBodyContent) *GetVccRouteEntryResponseBody
	GetContent() *GetVccRouteEntryResponseBodyContent
	SetMessage(v string) *GetVccRouteEntryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVccRouteEntryResponseBody
	GetRequestId() *string
}

type GetVccRouteEntryResponseBody struct {
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
	Content *GetVccRouteEntryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// BDBCC783-84CA-5733-8EEA-645C88B9009C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVccRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVccRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *GetVccRouteEntryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetVccRouteEntryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetVccRouteEntryResponseBody) GetContent() *GetVccRouteEntryResponseBodyContent {
	return s.Content
}

func (s *GetVccRouteEntryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVccRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVccRouteEntryResponseBody) SetAccessDeniedDetail(v string) *GetVccRouteEntryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetVccRouteEntryResponseBody) SetCode(v int32) *GetVccRouteEntryResponseBody {
	s.Code = &v
	return s
}

func (s *GetVccRouteEntryResponseBody) SetContent(v *GetVccRouteEntryResponseBodyContent) *GetVccRouteEntryResponseBody {
	s.Content = v
	return s
}

func (s *GetVccRouteEntryResponseBody) SetMessage(v string) *GetVccRouteEntryResponseBody {
	s.Message = &v
	return s
}

func (s *GetVccRouteEntryResponseBody) SetRequestId(v string) *GetVccRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVccRouteEntryResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVccRouteEntryResponseBodyContent struct {
	// Destination CIDR Block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1648085472000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// local
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
	// rg-aek2l4sq6l7u***
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
	// vcc-rte-31ocvdhq
	VccRouteEntryId *string `json:"VccRouteEntryId,omitempty" xml:"VccRouteEntryId,omitempty"`
}

func (s GetVccRouteEntryResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetVccRouteEntryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetVccRouteEntryResponseBodyContent) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *GetVccRouteEntryResponseBodyContent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetVccRouteEntryResponseBodyContent) GetMessage() *string {
	return s.Message
}

func (s *GetVccRouteEntryResponseBodyContent) GetNextHopId() *string {
	return s.NextHopId
}

func (s *GetVccRouteEntryResponseBodyContent) GetNextHopType() *string {
	return s.NextHopType
}

func (s *GetVccRouteEntryResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVccRouteEntryResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetVccRouteEntryResponseBodyContent) GetRouteType() *string {
	return s.RouteType
}

func (s *GetVccRouteEntryResponseBodyContent) GetStatus() *string {
	return s.Status
}

func (s *GetVccRouteEntryResponseBodyContent) GetTenantId() *string {
	return s.TenantId
}

func (s *GetVccRouteEntryResponseBodyContent) GetVccId() *string {
	return s.VccId
}

func (s *GetVccRouteEntryResponseBodyContent) GetVccRouteEntryId() *string {
	return s.VccRouteEntryId
}

func (s *GetVccRouteEntryResponseBodyContent) SetDestinationCidrBlock(v string) *GetVccRouteEntryResponseBodyContent {
	s.DestinationCidrBlock = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetGmtModified(v string) *GetVccRouteEntryResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetMessage(v string) *GetVccRouteEntryResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetNextHopId(v string) *GetVccRouteEntryResponseBodyContent {
	s.NextHopId = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetNextHopType(v string) *GetVccRouteEntryResponseBodyContent {
	s.NextHopType = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetRegionId(v string) *GetVccRouteEntryResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetResourceGroupId(v string) *GetVccRouteEntryResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetRouteType(v string) *GetVccRouteEntryResponseBodyContent {
	s.RouteType = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetStatus(v string) *GetVccRouteEntryResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetTenantId(v string) *GetVccRouteEntryResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetVccId(v string) *GetVccRouteEntryResponseBodyContent {
	s.VccId = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) SetVccRouteEntryId(v string) *GetVccRouteEntryResponseBodyContent {
	s.VccRouteEntryId = &v
	return s
}

func (s *GetVccRouteEntryResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
