// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnAttachmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVpnAttachmentsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnAttachmentsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVpnAttachmentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpnAttachmentsResponseBody
	GetTotalCount() *int32
	SetVpnAttachments(v []*DescribeVpnAttachmentsResponseBodyVpnAttachments) *DescribeVpnAttachmentsResponseBody
	GetVpnAttachments() []*DescribeVpnAttachmentsResponseBodyVpnAttachments
}

type DescribeVpnAttachmentsResponseBody struct {
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9F0725BB-186A-3564-91C3-AAE48042F853
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of IPsec-VPN connections associated with the transit router.
	VpnAttachments []*DescribeVpnAttachmentsResponseBodyVpnAttachments `json:"VpnAttachments,omitempty" xml:"VpnAttachments,omitempty" type:"Repeated"`
}

func (s DescribeVpnAttachmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpnAttachmentsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnAttachmentsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnAttachmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpnAttachmentsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpnAttachmentsResponseBody) GetVpnAttachments() []*DescribeVpnAttachmentsResponseBodyVpnAttachments {
	return s.VpnAttachments
}

func (s *DescribeVpnAttachmentsResponseBody) SetPageNumber(v int32) *DescribeVpnAttachmentsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnAttachmentsResponseBody) SetPageSize(v int32) *DescribeVpnAttachmentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnAttachmentsResponseBody) SetRequestId(v string) *DescribeVpnAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpnAttachmentsResponseBody) SetTotalCount(v int32) *DescribeVpnAttachmentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpnAttachmentsResponseBody) SetVpnAttachments(v []*DescribeVpnAttachmentsResponseBodyVpnAttachments) *DescribeVpnAttachmentsResponseBody {
	s.VpnAttachments = v
	return s
}

func (s *DescribeVpnAttachmentsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnAttachmentsResponseBodyVpnAttachments struct {
	// The type of resource that is associated with the IPsec-VPN connection. The value is set to **CEN**, which indicates that the IPsec-VPN connection is associated with a transit router.
	//
	// example:
	//
	// CEN
	AttachType *string `json:"AttachType,omitempty" xml:"AttachType,omitempty"`
	// Indicates whether the IPsec-VPN connection is associated with a transit router that belongs to another Alibaba Cloud account. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CrossAccountAuthorized *bool `json:"CrossAccountAuthorized,omitempty" xml:"CrossAccountAuthorized,omitempty"`
	// The description of the IPsec-VPN connection.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the IPsec-VPN connection.
	//
	// example:
	//
	// vco-p0w2jpkhi2eeop6q6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the IPsec-VPN connection.
	//
	// example:
	//
	// nametest1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The system tags of the IPsec-VPN connection.
	//
	// You can check whether an IPsec-VPN connection supports BGP based on the system tags.
	//
	// **BGPSupport**: indicates whether the IPsec-VPN connection supports BGP.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// {\\"description\\":\\"forwarding 1.7.22\\",\\"VisuallySsl\\":\\"true\\",\\"PbrPriority\\":\\"true\\",\\"BGPSupport\\":\\"true\\",\\"IDaaSNewVersion\\":\\"true\\"}
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The list of tags to be added to the IPsec-VPN connection.
	Tags []*DescribeVpnAttachmentsResponseBodyVpnAttachmentsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the transit router with which the IPsec-VPN connection is associated.
	//
	// example:
	//
	// tr-p0wkh4yryb1dnanqw****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The name of the transit router.
	//
	// example:
	//
	// nametest2
	TransitRouterName *string `json:"TransitRouterName,omitempty" xml:"TransitRouterName,omitempty"`
}

func (s DescribeVpnAttachmentsResponseBodyVpnAttachments) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnAttachmentsResponseBodyVpnAttachments) GoString() string {
	return s.String()
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) GetAttachType() *string {
	return s.AttachType
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) GetCrossAccountAuthorized() *bool {
	return s.CrossAccountAuthorized
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) GetDescription() *string {
	return s.Description
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) GetName() *string {
	return s.Name
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) GetTag() *string {
	return s.Tag
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) GetTags() []*DescribeVpnAttachmentsResponseBodyVpnAttachmentsTags {
	return s.Tags
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) GetTransitRouterName() *string {
	return s.TransitRouterName
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) SetAttachType(v string) *DescribeVpnAttachmentsResponseBodyVpnAttachments {
	s.AttachType = &v
	return s
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) SetCrossAccountAuthorized(v bool) *DescribeVpnAttachmentsResponseBodyVpnAttachments {
	s.CrossAccountAuthorized = &v
	return s
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) SetDescription(v string) *DescribeVpnAttachmentsResponseBodyVpnAttachments {
	s.Description = &v
	return s
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) SetInstanceId(v string) *DescribeVpnAttachmentsResponseBodyVpnAttachments {
	s.InstanceId = &v
	return s
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) SetName(v string) *DescribeVpnAttachmentsResponseBodyVpnAttachments {
	s.Name = &v
	return s
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) SetTag(v string) *DescribeVpnAttachmentsResponseBodyVpnAttachments {
	s.Tag = &v
	return s
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) SetTags(v []*DescribeVpnAttachmentsResponseBodyVpnAttachmentsTags) *DescribeVpnAttachmentsResponseBodyVpnAttachments {
	s.Tags = v
	return s
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) SetTransitRouterId(v string) *DescribeVpnAttachmentsResponseBodyVpnAttachments {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) SetTransitRouterName(v string) *DescribeVpnAttachmentsResponseBodyVpnAttachments {
	s.TransitRouterName = &v
	return s
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachments) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnAttachmentsResponseBodyVpnAttachmentsTags struct {
	// The tag key of the IPsec-VPN connection.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the IPsec-VPN connection.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVpnAttachmentsResponseBodyVpnAttachmentsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnAttachmentsResponseBodyVpnAttachmentsTags) GoString() string {
	return s.String()
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachmentsTags) GetKey() *string {
	return s.Key
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachmentsTags) GetValue() *string {
	return s.Value
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachmentsTags) SetKey(v string) *DescribeVpnAttachmentsResponseBodyVpnAttachmentsTags {
	s.Key = &v
	return s
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachmentsTags) SetValue(v string) *DescribeVpnAttachmentsResponseBodyVpnAttachmentsTags {
	s.Value = &v
	return s
}

func (s *DescribeVpnAttachmentsResponseBodyVpnAttachmentsTags) Validate() error {
	return dara.Validate(s)
}
