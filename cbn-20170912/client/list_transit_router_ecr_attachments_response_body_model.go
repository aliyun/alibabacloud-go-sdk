// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterEcrAttachmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterEcrAttachmentsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterEcrAttachmentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTransitRouterEcrAttachmentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransitRouterEcrAttachmentsResponseBody
	GetTotalCount() *int32
	SetTransitRouterAttachments(v []*ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) *ListTransitRouterEcrAttachmentsResponseBody
	GetTransitRouterAttachments() []*ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments
}

type ListTransitRouterEcrAttachmentsResponseBody struct {
	// The number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 461EC1B5-04A8-4706-8764-8F5BCEF48A6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the ECR connections.
	TransitRouterAttachments []*ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments `json:"TransitRouterAttachments,omitempty" xml:"TransitRouterAttachments,omitempty" type:"Repeated"`
}

func (s ListTransitRouterEcrAttachmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterEcrAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterEcrAttachmentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterEcrAttachmentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterEcrAttachmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRouterEcrAttachmentsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransitRouterEcrAttachmentsResponseBody) GetTransitRouterAttachments() []*ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments {
	return s.TransitRouterAttachments
}

func (s *ListTransitRouterEcrAttachmentsResponseBody) SetMaxResults(v int32) *ListTransitRouterEcrAttachmentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBody) SetNextToken(v string) *ListTransitRouterEcrAttachmentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBody) SetRequestId(v string) *ListTransitRouterEcrAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBody) SetTotalCount(v int32) *ListTransitRouterEcrAttachmentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBody) SetTransitRouterAttachments(v []*ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) *ListTransitRouterEcrAttachmentsResponseBody {
	s.TransitRouterAttachments = v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments struct {
	// Indicates whether the Enterprise Edition transit router can automatically advertise routes to ECRs.
	//
	// The value is **true**, which indicates that the Enterprise Edition transit router can automatically advertise routes to ECRs.
	//
	// example:
	//
	// true
	AutoPublishRouteEnabled *bool `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The time when the ECR connection was created.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-15T02:14Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the ECR with which the ECR connection is associated.
	//
	// example:
	//
	// ecr-n78omt2qsko06y****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the Alibaba Cloud account to which the ECR connection belongs.
	//
	// example:
	//
	// 1688111111111111
	EcrOwnerId *int64 `json:"EcrOwnerId,omitempty" xml:"EcrOwnerId,omitempty"`
	// The entity that pays the fees of the network instance. Valid values: Valid values:
	//
	// 	- **PayByCenOwner**: The Alibaba Cloud account to which the transit router belongs pays the connection and data forwarding fees of the ECR.
	//
	// 	- **PayByResourceOwner**: The Alibaba Cloud account to which the ECR belongs pays the connection and data forwarding fees of the ECR.
	//
	// example:
	//
	// PayByCenOwner
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The type of resource to which the transit router is connected. Valid values:
	//
	// The value is **ECR**, which indicates ECR connections.
	//
	// example:
	//
	// ECR
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the ECR connection. Valid values:
	//
	// 	- **Attached**
	//
	// 	- **Attaching**
	//
	// 	- **Detaching**
	//
	// example:
	//
	// Attached
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachmentsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The description of the ECR connection.
	//
	// example:
	//
	// desctest
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The ID of the ECR connection.
	//
	// example:
	//
	// tr-attach-nls9fzkfat8934****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The name of the ECR connection.
	//
	// example:
	//
	// testname
	TransitRouterAttachmentName *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	// The ID of the Enterprise Edition transit router.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The region ID of the transit router.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-shanghai
	TransitRouterRegionId *string `json:"TransitRouterRegionId,omitempty" xml:"TransitRouterRegionId,omitempty"`
}

func (s ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) GoString() string {
	return s.String()
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) GetAutoPublishRouteEnabled() *bool {
	return s.AutoPublishRouteEnabled
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) GetCenId() *string {
	return s.CenId
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) GetEcrId() *string {
	return s.EcrId
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) GetEcrOwnerId() *int64 {
	return s.EcrOwnerId
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) GetOrderType() *string {
	return s.OrderType
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) GetTags() []*ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachmentsTags {
	return s.Tags
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterAttachmentDescription() *string {
	return s.TransitRouterAttachmentDescription
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterAttachmentName() *string {
	return s.TransitRouterAttachmentName
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterRegionId() *string {
	return s.TransitRouterRegionId
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) SetAutoPublishRouteEnabled(v bool) *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) SetCenId(v string) *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) SetCreationTime(v string) *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments {
	s.CreationTime = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) SetEcrId(v string) *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments {
	s.EcrId = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) SetEcrOwnerId(v int64) *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments {
	s.EcrOwnerId = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) SetOrderType(v string) *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments {
	s.OrderType = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) SetResourceType(v string) *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments {
	s.ResourceType = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) SetStatus(v string) *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments {
	s.Status = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) SetTags(v []*ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachmentsTags) *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments {
	s.Tags = v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentDescription(v string) *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentId(v string) *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentName(v string) *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterId(v string) *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterRegionId(v string) *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterRegionId = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachmentsTags struct {
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
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachmentsTags) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachmentsTags) GoString() string {
	return s.String()
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachmentsTags) GetKey() *string {
	return s.Key
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachmentsTags) GetValue() *string {
	return s.Value
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachmentsTags) SetKey(v string) *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachmentsTags {
	s.Key = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachmentsTags) SetValue(v string) *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachmentsTags {
	s.Value = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachmentsTags) Validate() error {
	return dara.Validate(s)
}
