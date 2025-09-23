// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSharedResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSharedResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSharedResourcesRequest
	GetNextToken() *string
	SetResourceArns(v []*string) *ListSharedResourcesRequest
	GetResourceArns() []*string
	SetResourceIds(v []*string) *ListSharedResourcesRequest
	GetResourceIds() []*string
	SetResourceOwner(v string) *ListSharedResourcesRequest
	GetResourceOwner() *string
	SetResourceShareIds(v []*string) *ListSharedResourcesRequest
	GetResourceShareIds() []*string
	SetResourceType(v string) *ListSharedResourcesRequest
	GetResourceType() *string
	SetTarget(v string) *ListSharedResourcesRequest
	GetTarget() *string
}

type ListSharedResourcesRequest struct {
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The `token` that is used to initiate the next request. If the response of the current request is truncated, you can use the token to initiate another request and obtain the remaining records.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceArns []*string `json:"ResourceArns,omitempty" xml:"ResourceArns,omitempty" type:"Repeated"`
	// The ID of a shared resource.
	//
	// example:
	//
	// vsw-bp1upw03qyz8n7us9****
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The owner of the resource shares. Valid values:
	//
	// 	- Self: your account. If you set the value to Self, the resources you share with other accounts are queried.
	//
	// 	- OtherAccounts: another account. If you set the value to OtherAccounts, the resources other accounts share with you are queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// Self
	ResourceOwner *string `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The ID of a resource share.
	//
	// example:
	//
	// rs-6GRmdD3X****
	ResourceShareIds []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	// The type of the shared resources.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](https://help.aliyun.com/document_detail/450526.html).
	//
	// example:
	//
	// VSwitch
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the principal or resource owner.
	//
	// 	- If the value of `ResourceOwner` is `Self`, set this parameter to the ID of a principal.
	//
	// 	- If the value of `ResourceOwner` is `OtherAccounts`, set this parameter to the ID of a resource owner.
	//
	// example:
	//
	// 172050525300****
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ListSharedResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSharedResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListSharedResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSharedResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSharedResourcesRequest) GetResourceArns() []*string {
	return s.ResourceArns
}

func (s *ListSharedResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *ListSharedResourcesRequest) GetResourceOwner() *string {
	return s.ResourceOwner
}

func (s *ListSharedResourcesRequest) GetResourceShareIds() []*string {
	return s.ResourceShareIds
}

func (s *ListSharedResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListSharedResourcesRequest) GetTarget() *string {
	return s.Target
}

func (s *ListSharedResourcesRequest) SetMaxResults(v int32) *ListSharedResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSharedResourcesRequest) SetNextToken(v string) *ListSharedResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSharedResourcesRequest) SetResourceArns(v []*string) *ListSharedResourcesRequest {
	s.ResourceArns = v
	return s
}

func (s *ListSharedResourcesRequest) SetResourceIds(v []*string) *ListSharedResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *ListSharedResourcesRequest) SetResourceOwner(v string) *ListSharedResourcesRequest {
	s.ResourceOwner = &v
	return s
}

func (s *ListSharedResourcesRequest) SetResourceShareIds(v []*string) *ListSharedResourcesRequest {
	s.ResourceShareIds = v
	return s
}

func (s *ListSharedResourcesRequest) SetResourceType(v string) *ListSharedResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListSharedResourcesRequest) SetTarget(v string) *ListSharedResourcesRequest {
	s.Target = &v
	return s
}

func (s *ListSharedResourcesRequest) Validate() error {
	return dara.Validate(s)
}
