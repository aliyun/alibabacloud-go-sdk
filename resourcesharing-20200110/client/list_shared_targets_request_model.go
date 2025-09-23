// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSharedTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSharedTargetsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSharedTargetsRequest
	GetNextToken() *string
	SetResourceArn(v string) *ListSharedTargetsRequest
	GetResourceArn() *string
	SetResourceId(v string) *ListSharedTargetsRequest
	GetResourceId() *string
	SetResourceOwner(v string) *ListSharedTargetsRequest
	GetResourceOwner() *string
	SetResourceShareIds(v []*string) *ListSharedTargetsRequest
	GetResourceShareIds() []*string
	SetResourceType(v string) *ListSharedTargetsRequest
	GetResourceType() *string
	SetTargets(v []*string) *ListSharedTargetsRequest
	GetTargets() []*string
}

type ListSharedTargetsRequest struct {
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// The ID of the shared resource.
	//
	// example:
	//
	// vsw-bp1upw03qyz8n7us9****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The owner of the resource share.
	//
	// 	- Self: your account. If you set the value to Self, the principals that are associated with your resource shares are queried.
	//
	// 	- OtherAccounts: another account. If you set the value to OtherAccounts, the resource shares with which your account is associated and the owners of the resource shares are queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// Self
	ResourceOwner *string `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The ID of a resource share.
	//
	// Valid values of N: 1 to 5. This indicates that a maximum of five resource shares can be specified at a time.
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
	// The information about the principals.
	//
	// example:
	//
	// 114240524784****
	Targets []*string `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s ListSharedTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSharedTargetsRequest) GoString() string {
	return s.String()
}

func (s *ListSharedTargetsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSharedTargetsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSharedTargetsRequest) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *ListSharedTargetsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListSharedTargetsRequest) GetResourceOwner() *string {
	return s.ResourceOwner
}

func (s *ListSharedTargetsRequest) GetResourceShareIds() []*string {
	return s.ResourceShareIds
}

func (s *ListSharedTargetsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListSharedTargetsRequest) GetTargets() []*string {
	return s.Targets
}

func (s *ListSharedTargetsRequest) SetMaxResults(v int32) *ListSharedTargetsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSharedTargetsRequest) SetNextToken(v string) *ListSharedTargetsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSharedTargetsRequest) SetResourceArn(v string) *ListSharedTargetsRequest {
	s.ResourceArn = &v
	return s
}

func (s *ListSharedTargetsRequest) SetResourceId(v string) *ListSharedTargetsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListSharedTargetsRequest) SetResourceOwner(v string) *ListSharedTargetsRequest {
	s.ResourceOwner = &v
	return s
}

func (s *ListSharedTargetsRequest) SetResourceShareIds(v []*string) *ListSharedTargetsRequest {
	s.ResourceShareIds = v
	return s
}

func (s *ListSharedTargetsRequest) SetResourceType(v string) *ListSharedTargetsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListSharedTargetsRequest) SetTargets(v []*string) *ListSharedTargetsRequest {
	s.Targets = v
	return s
}

func (s *ListSharedTargetsRequest) Validate() error {
	return dara.Validate(s)
}
