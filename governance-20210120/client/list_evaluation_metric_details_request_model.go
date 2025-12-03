// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationMetricDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *ListEvaluationMetricDetailsRequest
	GetAccountId() *int64
	SetDate(v string) *ListEvaluationMetricDetailsRequest
	GetDate() *string
	SetId(v string) *ListEvaluationMetricDetailsRequest
	GetId() *string
	SetMaxResults(v int32) *ListEvaluationMetricDetailsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEvaluationMetricDetailsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListEvaluationMetricDetailsRequest
	GetRegionId() *string
	SetScope(v string) *ListEvaluationMetricDetailsRequest
	GetScope() *string
	SetSnapshotId(v string) *ListEvaluationMetricDetailsRequest
	GetSnapshotId() *string
}

type ListEvaluationMetricDetailsRequest struct {
	// The Alibaba Cloud account ID of the member. This parameter takes effect only when a multi-account governance maturity check is performed.
	//
	// example:
	//
	// 103144549568****
	AccountId *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Date      *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The ID of the check item.
	//
	// You can call the [ListEvaluationMetadata](https://help.aliyun.com/document_detail/2841889.html) operation to query the ID of the check item.
	//
	// example:
	//
	// xfyve5****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The maximum number of entries to return for a single request. Default value: 5.
	//
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAGEaXR18y1rqykZHIqRuBejOqED4S3Xne33c7zbn****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scope      *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ListEvaluationMetricDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationMetricDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetricDetailsRequest) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListEvaluationMetricDetailsRequest) GetDate() *string {
	return s.Date
}

func (s *ListEvaluationMetricDetailsRequest) GetId() *string {
	return s.Id
}

func (s *ListEvaluationMetricDetailsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEvaluationMetricDetailsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEvaluationMetricDetailsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEvaluationMetricDetailsRequest) GetScope() *string {
	return s.Scope
}

func (s *ListEvaluationMetricDetailsRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ListEvaluationMetricDetailsRequest) SetAccountId(v int64) *ListEvaluationMetricDetailsRequest {
	s.AccountId = &v
	return s
}

func (s *ListEvaluationMetricDetailsRequest) SetDate(v string) *ListEvaluationMetricDetailsRequest {
	s.Date = &v
	return s
}

func (s *ListEvaluationMetricDetailsRequest) SetId(v string) *ListEvaluationMetricDetailsRequest {
	s.Id = &v
	return s
}

func (s *ListEvaluationMetricDetailsRequest) SetMaxResults(v int32) *ListEvaluationMetricDetailsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEvaluationMetricDetailsRequest) SetNextToken(v string) *ListEvaluationMetricDetailsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEvaluationMetricDetailsRequest) SetRegionId(v string) *ListEvaluationMetricDetailsRequest {
	s.RegionId = &v
	return s
}

func (s *ListEvaluationMetricDetailsRequest) SetScope(v string) *ListEvaluationMetricDetailsRequest {
	s.Scope = &v
	return s
}

func (s *ListEvaluationMetricDetailsRequest) SetSnapshotId(v string) *ListEvaluationMetricDetailsRequest {
	s.SnapshotId = &v
	return s
}

func (s *ListEvaluationMetricDetailsRequest) Validate() error {
	return dara.Validate(s)
}
