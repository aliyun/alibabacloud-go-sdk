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
	// The number of entries per page for a paged query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next paged query.
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
	// The total number of entries.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of ECR connection information.
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

type ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachments struct {
	// Indicates whether the Enterprise Edition transit router automatically publishes route entries to the ECR instance.
	//
	// The value is **true*	- only, which indicates that route entries are automatically published.
	//
	// example:
	//
	// true
	AutoPublishRouteEnabled *bool `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	// The CEN instance ID.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The time when the ECR connection was created.
	//
	// The time is displayed in the ISO 8601 standard in UTC. Format: YYYY-MM-DDThh:mmZ.
	//
	// example:
	//
	// 2021-06-15T02:14Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The instance ID of the associated Express Connect Router (ECR).
	//
	// example:
	//
	// ecr-n78omt2qsko06y****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the Alibaba Cloud account to which the ECR instance belongs.
	//
	// example:
	//
	// 1210123456123456
	EcrOwnerId *int64 `json:"EcrOwnerId,omitempty" xml:"EcrOwnerId,omitempty"`
	// The payer of the network instance. Valid values:
	//
	// - **PayByCenOwner**: The connection fee and data processing fee of the ECR instance are paid by the account that owns the transit router instance.
	//
	// - **PayByResourceOwner**: The connection fee and data processing fee of the ECR instance are paid by the account that owns the ECR instance.
	//
	// example:
	//
	// PayByCenOwner
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The resource type of the connection.
	//
	// The value is **ECR*	- only, which indicates an Express Connect Router (ECR) instance.
	//
	// example:
	//
	// ECR
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the ECR connection.
	//
	// - **Attached**: attached.
	//
	// - **Attaching**: being attached.
	//
	// - **Detaching**: being detached.
	//
	// example:
	//
	// Attached
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*ListTransitRouterEcrAttachmentsResponseBodyTransitRouterAttachmentsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The description of the ECR connection.
	//
	// example:
	//
	// desctest
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The ECR connection ID.
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
	// The Enterprise Edition transit router instance ID.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The region ID of the transit router.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the region information corresponding to the region ID.
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
