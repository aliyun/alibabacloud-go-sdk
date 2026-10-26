// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectAuthorizationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListProjectAuthorizationsRequest
	GetInstanceId() *string
	SetKeyword(v string) *ListProjectAuthorizationsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListProjectAuthorizationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProjectAuthorizationsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListProjectAuthorizationsRequest
	GetRegionId() *string
	SetSkip(v int32) *ListProjectAuthorizationsRequest
	GetSkip() *int32
}

type ListProjectAuthorizationsRequest struct {
	// The ID of the bastion host instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The keyword for fuzzy search by project authorization name.
	//
	// example:
	//
	// name
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The maximum number of project authorizations to return per request. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The request token. Leave this parameter empty or set it to an empty string for the first request.
	//
	// example:
	//
	// 4ieSWJCwxvW3dk3wF.BqkrZmP72nWu5zJ5NWydMqyEs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of entries to skip.
	//
	// > This parameter takes effect only when NextToken is empty.
	//
	// example:
	//
	// 0
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
}

func (s ListProjectAuthorizationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectAuthorizationsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectAuthorizationsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListProjectAuthorizationsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListProjectAuthorizationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProjectAuthorizationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectAuthorizationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListProjectAuthorizationsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListProjectAuthorizationsRequest) SetInstanceId(v string) *ListProjectAuthorizationsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListProjectAuthorizationsRequest) SetKeyword(v string) *ListProjectAuthorizationsRequest {
	s.Keyword = &v
	return s
}

func (s *ListProjectAuthorizationsRequest) SetMaxResults(v int32) *ListProjectAuthorizationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectAuthorizationsRequest) SetNextToken(v string) *ListProjectAuthorizationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProjectAuthorizationsRequest) SetRegionId(v string) *ListProjectAuthorizationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListProjectAuthorizationsRequest) SetSkip(v int32) *ListProjectAuthorizationsRequest {
	s.Skip = &v
	return s
}

func (s *ListProjectAuthorizationsRequest) Validate() error {
	return dara.Validate(s)
}
