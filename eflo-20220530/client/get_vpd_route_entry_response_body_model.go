// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpdRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetVpdRouteEntryResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetVpdRouteEntryResponseBody
	GetCode() *int32
	SetContent(v *GetVpdRouteEntryResponseBodyContent) *GetVpdRouteEntryResponseBody
	GetContent() *GetVpdRouteEntryResponseBodyContent
	SetMessage(v string) *GetVpdRouteEntryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVpdRouteEntryResponseBody
	GetRequestId() *string
}

type GetVpdRouteEntryResponseBody struct {
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
	Content *GetVpdRouteEntryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF80857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVpdRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVpdRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpdRouteEntryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetVpdRouteEntryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetVpdRouteEntryResponseBody) GetContent() *GetVpdRouteEntryResponseBodyContent {
	return s.Content
}

func (s *GetVpdRouteEntryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVpdRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVpdRouteEntryResponseBody) SetAccessDeniedDetail(v string) *GetVpdRouteEntryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetVpdRouteEntryResponseBody) SetCode(v int32) *GetVpdRouteEntryResponseBody {
	s.Code = &v
	return s
}

func (s *GetVpdRouteEntryResponseBody) SetContent(v *GetVpdRouteEntryResponseBodyContent) *GetVpdRouteEntryResponseBody {
	s.Content = v
	return s
}

func (s *GetVpdRouteEntryResponseBody) SetMessage(v string) *GetVpdRouteEntryResponseBody {
	s.Message = &v
	return s
}

func (s *GetVpdRouteEntryResponseBody) SetRequestId(v string) *GetVpdRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVpdRouteEntryResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVpdRouteEntryResponseBodyContent struct {
	// Destination CIDR block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The time when the agent was last modified.
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
	// rg-acfmv7mcq63uyhq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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
	// vpd-ze3na0wf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// Lingjun CIDR block route entry ID
	//
	// example:
	//
	// vpd-rte-toekyqel
	VpdRouteEntryId *string `json:"VpdRouteEntryId,omitempty" xml:"VpdRouteEntryId,omitempty"`
}

func (s GetVpdRouteEntryResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetVpdRouteEntryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetVpdRouteEntryResponseBodyContent) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *GetVpdRouteEntryResponseBodyContent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetVpdRouteEntryResponseBodyContent) GetNextHopId() *string {
	return s.NextHopId
}

func (s *GetVpdRouteEntryResponseBodyContent) GetNextHopType() *string {
	return s.NextHopType
}

func (s *GetVpdRouteEntryResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpdRouteEntryResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetVpdRouteEntryResponseBodyContent) GetRouteType() *string {
	return s.RouteType
}

func (s *GetVpdRouteEntryResponseBodyContent) GetStatus() *string {
	return s.Status
}

func (s *GetVpdRouteEntryResponseBodyContent) GetTenantId() *string {
	return s.TenantId
}

func (s *GetVpdRouteEntryResponseBodyContent) GetVpdId() *string {
	return s.VpdId
}

func (s *GetVpdRouteEntryResponseBodyContent) GetVpdRouteEntryId() *string {
	return s.VpdRouteEntryId
}

func (s *GetVpdRouteEntryResponseBodyContent) SetDestinationCidrBlock(v string) *GetVpdRouteEntryResponseBodyContent {
	s.DestinationCidrBlock = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetGmtModified(v string) *GetVpdRouteEntryResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetNextHopId(v string) *GetVpdRouteEntryResponseBodyContent {
	s.NextHopId = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetNextHopType(v string) *GetVpdRouteEntryResponseBodyContent {
	s.NextHopType = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetRegionId(v string) *GetVpdRouteEntryResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetResourceGroupId(v string) *GetVpdRouteEntryResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetRouteType(v string) *GetVpdRouteEntryResponseBodyContent {
	s.RouteType = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetStatus(v string) *GetVpdRouteEntryResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetTenantId(v string) *GetVpdRouteEntryResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetVpdId(v string) *GetVpdRouteEntryResponseBodyContent {
	s.VpdId = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) SetVpdRouteEntryId(v string) *GetVpdRouteEntryResponseBodyContent {
	s.VpdRouteEntryId = &v
	return s
}

func (s *GetVpdRouteEntryResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
