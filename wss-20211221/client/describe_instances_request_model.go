// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribeInstancesRequest
	GetInstanceIds() []*string
	SetMaxResults(v int32) *DescribeInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInstancesRequest
	GetNextToken() *string
	SetNoResourceId(v bool) *DescribeInstancesRequest
	GetNoResourceId() *bool
	SetPropertyKeyList(v []*string) *DescribeInstancesRequest
	GetPropertyKeyList() []*string
	SetResourceIds(v []*string) *DescribeInstancesRequest
	GetResourceIds() []*string
	SetResourceType(v string) *DescribeInstancesRequest
	GetResourceType() *string
	SetStateList(v []*string) *DescribeInstancesRequest
	GetStateList() []*string
}

type DescribeInstancesRequest struct {
	// The list of sold instance IDs, which is used to specify the instances for this operation.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The maximum number of records to return per query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Leave it empty for the first query. For subsequent queries, pass the NextToken value returned in the previous response as is. Do not construct or decode it yourself.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6k=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Specifies whether to query only instances that are not associated with a business resource ID. Valid values:
	//
	// - true: Queries only unassociated instances.
	//
	// - false or empty: Does not filter by this condition.
	NoResourceId *bool `json:"NoResourceId,omitempty" xml:"NoResourceId,omitempty"`
	// The list of instance property names to be filtered.
	PropertyKeyList []*string `json:"PropertyKeyList,omitempty" xml:"PropertyKeyList,omitempty" type:"Repeated"`
	// The list of associated business resource IDs.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The resource type. For example, CorePackage indicates a core-hour package. It must be consistent with the resource type of the target instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// CreditPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of instance statuses used for filtering. You can use NORMAL, DELETED, PENDING, and EXPIRED. ARCHIVED indicates the historical archived status.
	StateList []*string `json:"StateList,omitempty" xml:"StateList,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstancesRequest) GetNoResourceId() *bool {
	return s.NoResourceId
}

func (s *DescribeInstancesRequest) GetPropertyKeyList() []*string {
	return s.PropertyKeyList
}

func (s *DescribeInstancesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DescribeInstancesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeInstancesRequest) GetStateList() []*string {
	return s.StateList
}

func (s *DescribeInstancesRequest) SetInstanceIds(v []*string) *DescribeInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeInstancesRequest) SetMaxResults(v int32) *DescribeInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstancesRequest) SetNextToken(v string) *DescribeInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInstancesRequest) SetNoResourceId(v bool) *DescribeInstancesRequest {
	s.NoResourceId = &v
	return s
}

func (s *DescribeInstancesRequest) SetPropertyKeyList(v []*string) *DescribeInstancesRequest {
	s.PropertyKeyList = v
	return s
}

func (s *DescribeInstancesRequest) SetResourceIds(v []*string) *DescribeInstancesRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeInstancesRequest) SetResourceType(v string) *DescribeInstancesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeInstancesRequest) SetStateList(v []*string) *DescribeInstancesRequest {
	s.StateList = v
	return s
}

func (s *DescribeInstancesRequest) Validate() error {
	return dara.Validate(s)
}
