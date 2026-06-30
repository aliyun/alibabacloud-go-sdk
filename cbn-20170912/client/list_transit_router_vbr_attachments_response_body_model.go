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
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query.
	//
	// - If this parameter is empty, no more data is returned.
	//
	// - If a value is returned for this parameter, it is the token that you can use to retrieve the next page of results.
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
	if s.TransitRouterAttachments != nil {
		for _, item := range s.TransitRouterAttachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments struct {
	// Indicates whether the Enterprise Edition transit router automatically advertises routes to the VBR.
	//
	// - **false**: no.
	//
	// - **true**: yes.
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
	// The time is displayed in the YYYY-MM-DDThh:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-15T15:20Z
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ManagedService *string `json:"ManagedService,omitempty" xml:"ManagedService,omitempty"`
	// The payer for the network instance. Valid values:
	//
	// - **PayByCenOwner**: The connection fee and data transfer fee for the VBR are paid by the account that owns the transit router.
	//
	// - **PayByResourceOwner**: The connection fee and data transfer fee for the VBR are paid by the account that owns the VBR.
	//
	// example:
	//
	// PayByCenOwner
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The resource type of the connection.
	//
	// The value is set to **VBR**, which indicates a VBR instance.
	//
	// example:
	//
	// VBR
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the VBR connection.
	//
	// - **Attached**: The connection is established.
	//
	// - **Attaching**: The connection is being established.
	//
	// - **Detaching**: The connection is being removed.
	//
	// example:
	//
	// Attached
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
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
	// The ID of the Enterprise Edition transit router.
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
	// The ID of the region where the VBR is deployed.
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

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) GetManagedService() *string {
	return s.ManagedService
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

func (s *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments) SetManagedService(v string) *ListTransitRouterVbrAttachmentsResponseBodyTransitRouterAttachments {
	s.ManagedService = &v
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
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
