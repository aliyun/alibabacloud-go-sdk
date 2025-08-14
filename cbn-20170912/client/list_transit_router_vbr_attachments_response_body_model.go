// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterVbrAttachmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterVbrAttachmentsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterVbrAttachmentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTransitRouterVbrAttachmentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransitRouterVbrAttachmentsResponseBody
	GetTotalCount() *int32
	SetTransitRouterAttachments(v []*ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) *ListTransitRouterVbrAttachmentsResponseBody
	GetTransitRouterAttachments() []*ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments
}

type ListTransitRouterVbrAttachmentsResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// dd20****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F42D9616-29EB-4E75-8CA8-9654D4E07501
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// A list of VBR connections.
	TransitRouterAttachments []*ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments `json:"TransitRouterAttachments,omitempty" xml:"TransitRouterAttachments,omitempty" type:"Repeated"`
}

func (s ListTransitRouterVbrAttachmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVbrAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVbrAttachmentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterVbrAttachmentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterVbrAttachmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRouterVbrAttachmentsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransitRouterVbrAttachmentsResponseBody) GetTransitRouterAttachments() []*ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	return s.TransitRouterAttachments
}

func (s *ListTransitRouterVbrAttachmentsResponseBody) SetMaxResults(v int32) *ListTransitRouterVbrAttachmentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBody) SetNextToken(v string) *ListTransitRouterVbrAttachmentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBody) SetRequestId(v string) *ListTransitRouterVbrAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBody) SetTotalCount(v int32) *ListTransitRouterVbrAttachmentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBody) SetTransitRouterAttachments(v []*ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) *ListTransitRouterVbrAttachmentsResponseBody {
	s.TransitRouterAttachments = v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments struct {
	// Indicates whether the Enterprise Edition transit router is allowed to automatically advertise routes to the VBR. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// false
	AutoPublishRouteEnabled *bool `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The time when the VBR connection was created.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-15T15:20Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The entity that pays the fees of the network instance. Valid values:
	//
	// 	- **PayByCenOwner**: the Alibaba Cloud account that owns the CEN instance.
	//
	// 	- **PayByResourceOwner**: the Alibaba Cloud account that owns the network instance.
	//
	// example:
	//
	// PayByCenOwner
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The type of resource to which the transit router is connected. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **CCN**
	//
	// 	- **VBR**
	//
	// 	- **TR**
	//
	// example:
	//
	// VBR
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the VBR connection. Valid values:
	//
	// 	- **Attached**
	//
	// 	- **Attaching**
	//
	// 	- **Detaching**
	//
	// 	- **Detached**
	//
	// example:
	//
	// Attached
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// A list of tags.
	Tags []*ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachmentsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The description of the VBR connection.
	//
	// example:
	//
	// testdesc
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The ID of the VBR connection.
	//
	// example:
	//
	// tr-attach-oyf70wfuorwx87****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The name of the VBR connection.
	//
	// example:
	//
	// testa
	TransitRouterAttachmentName *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	// The description of the Enterprise Edition transit router.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The VBR ID.
	//
	// example:
	//
	// vbr-bp1svadp4lq38janc****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	// The ID of the Alibaba Cloud account to which the VBR belongs.
	//
	// example:
	//
	// 1688111111111111
	VbrOwnerId *int64 `json:"VbrOwnerId,omitempty" xml:"VbrOwnerId,omitempty"`
	// The region ID of the VBR.
	//
	// example:
	//
	// cn-hangzhou
	VbrRegionId *string `json:"VbrRegionId,omitempty" xml:"VbrRegionId,omitempty"`
}

func (s ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GetAutoPublishRouteEnabled() *bool {
	return s.AutoPublishRouteEnabled
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GetCenId() *string {
	return s.CenId
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GetOrderType() *string {
	return s.OrderType
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GetTags() []*ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachmentsTags {
	return s.Tags
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterAttachmentDescription() *string {
	return s.TransitRouterAttachmentDescription
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterAttachmentName() *string {
	return s.TransitRouterAttachmentName
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GetVbrId() *string {
	return s.VbrId
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GetVbrOwnerId() *int64 {
	return s.VbrOwnerId
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GetVbrRegionId() *string {
	return s.VbrRegionId
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetAutoPublishRouteEnabled(v bool) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetCenId(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetCreationTime(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.CreationTime = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetOrderType(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.OrderType = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetResourceType(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.ResourceType = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetStatus(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.Status = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetTags(v []*ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachmentsTags) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.Tags = v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentDescription(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentId(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentName(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterId(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetVbrId(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.VbrId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetVbrOwnerId(v int64) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.VbrOwnerId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetVbrRegionId(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.VbrRegionId = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachmentsTags struct {
	// The tag key.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value_A1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachmentsTags) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachmentsTags) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachmentsTags) GetKey() *string {
	return s.Key
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachmentsTags) GetValue() *string {
	return s.Value
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachmentsTags) SetKey(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachmentsTags {
	s.Key = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachmentsTags) SetValue(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachmentsTags {
	s.Value = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachmentsTags) Validate() error {
	return dara.Validate(s)
}
