// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerIdeInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListServerIdeInstancesRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListServerIdeInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServerIdeInstancesRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListServerIdeInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListServerIdeInstancesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListServerIdeInstancesRequest
	GetProjectId() *int64
	SetRelatedUserId(v string) *ListServerIdeInstancesRequest
	GetRelatedUserId() *string
	SetResourceGroupId(v string) *ListServerIdeInstancesRequest
	GetResourceGroupId() *string
	SetSubType(v string) *ListServerIdeInstancesRequest
	GetSubType() *string
}

type ListServerIdeInstancesRequest struct {
	// The keyword for fuzzy match by instance ID or instance name.
	//
	// example:
	//
	// notebook_dev
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The maximum number of records to return in a single request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// CAESG****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. Minimum value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The account ID of the user who owns the instance. Used to filter instances by owner.
	//
	// example:
	//
	// 20933221576142****
	RelatedUserId *string `json:"RelatedUserId,omitempty" xml:"RelatedUserId,omitempty"`
	// The DataWorks resource group identifier. You can specify a numeric resource group ID or a full identifier in the format of Serverless_res_group_{tenantId}_{resgId}.
	//
	// example:
	//
	// Serverless_res_group_123456789012345_9876543210****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance subtype. Valid values:
	//
	// - PERSONAL_DEV: personal development environment.
	//
	// - DATA_AGENT: Data Agent.
	//
	// example:
	//
	// PERSONAL_DEV
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s ListServerIdeInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListServerIdeInstancesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListServerIdeInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServerIdeInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServerIdeInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListServerIdeInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServerIdeInstancesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListServerIdeInstancesRequest) GetRelatedUserId() *string {
	return s.RelatedUserId
}

func (s *ListServerIdeInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListServerIdeInstancesRequest) GetSubType() *string {
	return s.SubType
}

func (s *ListServerIdeInstancesRequest) SetKeyword(v string) *ListServerIdeInstancesRequest {
	s.Keyword = &v
	return s
}

func (s *ListServerIdeInstancesRequest) SetMaxResults(v int32) *ListServerIdeInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServerIdeInstancesRequest) SetNextToken(v string) *ListServerIdeInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServerIdeInstancesRequest) SetPageNumber(v int32) *ListServerIdeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServerIdeInstancesRequest) SetPageSize(v int32) *ListServerIdeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListServerIdeInstancesRequest) SetProjectId(v int64) *ListServerIdeInstancesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListServerIdeInstancesRequest) SetRelatedUserId(v string) *ListServerIdeInstancesRequest {
	s.RelatedUserId = &v
	return s
}

func (s *ListServerIdeInstancesRequest) SetResourceGroupId(v string) *ListServerIdeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServerIdeInstancesRequest) SetSubType(v string) *ListServerIdeInstancesRequest {
	s.SubType = &v
	return s
}

func (s *ListServerIdeInstancesRequest) Validate() error {
	return dara.Validate(s)
}
