// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSharedResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListSharedResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSharedResourcesResponseBody
	GetRequestId() *string
	SetSharedResources(v []*ListSharedResourcesResponseBodySharedResources) *ListSharedResourcesResponseBody
	GetSharedResources() []*ListSharedResourcesResponseBodySharedResources
}

type ListSharedResourcesResponseBody struct {
	// The token that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04677DCA-7C33-464B-8811-1B1DA3C3D197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the shared resources.
	SharedResources []*ListSharedResourcesResponseBodySharedResources `json:"SharedResources,omitempty" xml:"SharedResources,omitempty" type:"Repeated"`
}

func (s ListSharedResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSharedResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSharedResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSharedResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSharedResourcesResponseBody) GetSharedResources() []*ListSharedResourcesResponseBodySharedResources {
	return s.SharedResources
}

func (s *ListSharedResourcesResponseBody) SetNextToken(v string) *ListSharedResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSharedResourcesResponseBody) SetRequestId(v string) *ListSharedResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSharedResourcesResponseBody) SetSharedResources(v []*ListSharedResourcesResponseBodySharedResources) *ListSharedResourcesResponseBody {
	s.SharedResources = v
	return s
}

func (s *ListSharedResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSharedResourcesResponseBodySharedResources struct {
	// The time when the shared resource was associated with the resource share.
	//
	// example:
	//
	// 2020-12-07T07:39:02.921Z
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// The ID of the shared resource.
	//
	// example:
	//
	// vsw-bp1upw03qyz8n7us9****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the resource share.
	//
	// example:
	//
	// rs-6GRmdD3X****
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The status of the shared resource. This parameter is returned only when you query the resources that other accounts share with you.
	//
	// Valid values:
	//
	// 	- Available: The resource is available.
	//
	// 	- ZonalResourceInaccessible: The resource is unavailable in the current zone.
	//
	// 	- LimitExceeded: The resource is unavailable because the maximum number of resources that other accounts can share with you exceeds the upper limit.
	//
	// 	- Unavailable: The resource is unavailable.
	//
	// example:
	//
	// Available
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The cause of the association failure.
	//
	// example:
	//
	// The reason for the association failure.
	ResourceStatusMessage *string `json:"ResourceStatusMessage,omitempty" xml:"ResourceStatusMessage,omitempty"`
	// The type of the shared resource.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](https://help.aliyun.com/document_detail/450526.html).
	//
	// example:
	//
	// VSwitch
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The time when the association of the shared resource was updated.
	//
	// example:
	//
	// 2020-12-07T07:39:02.921Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListSharedResourcesResponseBodySharedResources) String() string {
	return dara.Prettify(s)
}

func (s ListSharedResourcesResponseBodySharedResources) GoString() string {
	return s.String()
}

func (s *ListSharedResourcesResponseBodySharedResources) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSharedResourcesResponseBodySharedResources) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *ListSharedResourcesResponseBodySharedResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListSharedResourcesResponseBodySharedResources) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *ListSharedResourcesResponseBodySharedResources) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *ListSharedResourcesResponseBodySharedResources) GetResourceStatusMessage() *string {
	return s.ResourceStatusMessage
}

func (s *ListSharedResourcesResponseBodySharedResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListSharedResourcesResponseBodySharedResources) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListSharedResourcesResponseBodySharedResources) SetCreateTime(v string) *ListSharedResourcesResponseBodySharedResources {
	s.CreateTime = &v
	return s
}

func (s *ListSharedResourcesResponseBodySharedResources) SetResourceArn(v string) *ListSharedResourcesResponseBodySharedResources {
	s.ResourceArn = &v
	return s
}

func (s *ListSharedResourcesResponseBodySharedResources) SetResourceId(v string) *ListSharedResourcesResponseBodySharedResources {
	s.ResourceId = &v
	return s
}

func (s *ListSharedResourcesResponseBodySharedResources) SetResourceShareId(v string) *ListSharedResourcesResponseBodySharedResources {
	s.ResourceShareId = &v
	return s
}

func (s *ListSharedResourcesResponseBodySharedResources) SetResourceStatus(v string) *ListSharedResourcesResponseBodySharedResources {
	s.ResourceStatus = &v
	return s
}

func (s *ListSharedResourcesResponseBodySharedResources) SetResourceStatusMessage(v string) *ListSharedResourcesResponseBodySharedResources {
	s.ResourceStatusMessage = &v
	return s
}

func (s *ListSharedResourcesResponseBodySharedResources) SetResourceType(v string) *ListSharedResourcesResponseBodySharedResources {
	s.ResourceType = &v
	return s
}

func (s *ListSharedResourcesResponseBodySharedResources) SetUpdateTime(v string) *ListSharedResourcesResponseBodySharedResources {
	s.UpdateTime = &v
	return s
}

func (s *ListSharedResourcesResponseBodySharedResources) Validate() error {
	return dara.Validate(s)
}
