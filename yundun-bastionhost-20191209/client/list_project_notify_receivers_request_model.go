// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectNotifyReceiversRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListProjectNotifyReceiversRequest
	GetInstanceId() *string
	SetKeyword(v string) *ListProjectNotifyReceiversRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListProjectNotifyReceiversRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProjectNotifyReceiversRequest
	GetNextToken() *string
	SetProjectId(v int64) *ListProjectNotifyReceiversRequest
	GetProjectId() *int64
	SetRegionId(v string) *ListProjectNotifyReceiversRequest
	GetRegionId() *string
	SetSkip(v int32) *ListProjectNotifyReceiversRequest
	GetSkip() *int32
}

type ListProjectNotifyReceiversRequest struct {
	// The instance ID of the bastion host.
	//
	// > You can invoke [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) to obtain this parameter.
	//
	// example:
	//
	// bastionhost-cn-zvp282aly06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name, phone number, or email address of the notification recipient to query. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of entries to query. Default value: 20. Maximum value: 100.
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
	// The project ID.
	//
	// > You can call ListProjects to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
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

func (s ListProjectNotifyReceiversRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectNotifyReceiversRequest) GoString() string {
	return s.String()
}

func (s *ListProjectNotifyReceiversRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListProjectNotifyReceiversRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListProjectNotifyReceiversRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProjectNotifyReceiversRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectNotifyReceiversRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListProjectNotifyReceiversRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListProjectNotifyReceiversRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListProjectNotifyReceiversRequest) SetInstanceId(v string) *ListProjectNotifyReceiversRequest {
	s.InstanceId = &v
	return s
}

func (s *ListProjectNotifyReceiversRequest) SetKeyword(v string) *ListProjectNotifyReceiversRequest {
	s.Keyword = &v
	return s
}

func (s *ListProjectNotifyReceiversRequest) SetMaxResults(v int32) *ListProjectNotifyReceiversRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectNotifyReceiversRequest) SetNextToken(v string) *ListProjectNotifyReceiversRequest {
	s.NextToken = &v
	return s
}

func (s *ListProjectNotifyReceiversRequest) SetProjectId(v int64) *ListProjectNotifyReceiversRequest {
	s.ProjectId = &v
	return s
}

func (s *ListProjectNotifyReceiversRequest) SetRegionId(v string) *ListProjectNotifyReceiversRequest {
	s.RegionId = &v
	return s
}

func (s *ListProjectNotifyReceiversRequest) SetSkip(v int32) *ListProjectNotifyReceiversRequest {
	s.Skip = &v
	return s
}

func (s *ListProjectNotifyReceiversRequest) Validate() error {
	return dara.Validate(s)
}
