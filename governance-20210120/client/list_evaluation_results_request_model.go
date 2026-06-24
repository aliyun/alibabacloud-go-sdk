// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *ListEvaluationResultsRequest
	GetAccountId() *int64
	SetEvaluationDomain(v string) *ListEvaluationResultsRequest
	GetEvaluationDomain() *string
	SetFilters(v []*ListEvaluationResultsRequestFilters) *ListEvaluationResultsRequest
	GetFilters() []*ListEvaluationResultsRequestFilters
	SetLensCode(v string) *ListEvaluationResultsRequest
	GetLensCode() *string
	SetRegionId(v string) *ListEvaluationResultsRequest
	GetRegionId() *string
	SetScope(v string) *ListEvaluationResultsRequest
	GetScope() *string
	SetSnapshotId(v string) *ListEvaluationResultsRequest
	GetSnapshotId() *string
	SetTopicCode(v string) *ListEvaluationResultsRequest
	GetTopicCode() *string
}

type ListEvaluationResultsRequest struct {
	// Member account ID. This parameter is only applicable to multi-account evaluation mode.
	//
	// example:
	//
	// 176618589410****
	AccountId        *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	EvaluationDomain *string `json:"EvaluationDomain,omitempty" xml:"EvaluationDomain,omitempty"`
	// Filter conditions.
	Filters []*ListEvaluationResultsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// Special evaluation code. Valid values:
	//
	// - basic (default): Basic model (governance maturity) evaluation.
	//
	// - ack: Container construction special evaluation.
	//
	// - ai: Machine learning special evaluation.
	//
	// - nis: Network service special evaluation.
	//
	// example:
	//
	// basic
	LensCode *string `json:"LensCode,omitempty" xml:"LensCode,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Governance maturity evaluation scope. Valid values:
	//
	// - Account (default): Performs single-account governance maturity evaluation, evaluating only the current account.
	//
	// - ResourceDirectory: Performs multi-account governance maturity evaluation, evaluating all member accounts in the resource directory. Before performing this operation, you must first upgrade to multi-account governance maturity evaluation.
	//
	// example:
	//
	// ResourceDirectory
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// Evaluation snapshot ID.
	//
	// example:
	//
	// es-bp1r**************
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// Governance topic code.
	//
	// example:
	//
	// IdentityAndAccessManagement
	TopicCode *string `json:"TopicCode,omitempty" xml:"TopicCode,omitempty"`
}

func (s ListEvaluationResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsRequest) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListEvaluationResultsRequest) GetEvaluationDomain() *string {
	return s.EvaluationDomain
}

func (s *ListEvaluationResultsRequest) GetFilters() []*ListEvaluationResultsRequestFilters {
	return s.Filters
}

func (s *ListEvaluationResultsRequest) GetLensCode() *string {
	return s.LensCode
}

func (s *ListEvaluationResultsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEvaluationResultsRequest) GetScope() *string {
	return s.Scope
}

func (s *ListEvaluationResultsRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ListEvaluationResultsRequest) GetTopicCode() *string {
	return s.TopicCode
}

func (s *ListEvaluationResultsRequest) SetAccountId(v int64) *ListEvaluationResultsRequest {
	s.AccountId = &v
	return s
}

func (s *ListEvaluationResultsRequest) SetEvaluationDomain(v string) *ListEvaluationResultsRequest {
	s.EvaluationDomain = &v
	return s
}

func (s *ListEvaluationResultsRequest) SetFilters(v []*ListEvaluationResultsRequestFilters) *ListEvaluationResultsRequest {
	s.Filters = v
	return s
}

func (s *ListEvaluationResultsRequest) SetLensCode(v string) *ListEvaluationResultsRequest {
	s.LensCode = &v
	return s
}

func (s *ListEvaluationResultsRequest) SetRegionId(v string) *ListEvaluationResultsRequest {
	s.RegionId = &v
	return s
}

func (s *ListEvaluationResultsRequest) SetScope(v string) *ListEvaluationResultsRequest {
	s.Scope = &v
	return s
}

func (s *ListEvaluationResultsRequest) SetSnapshotId(v string) *ListEvaluationResultsRequest {
	s.SnapshotId = &v
	return s
}

func (s *ListEvaluationResultsRequest) SetTopicCode(v string) *ListEvaluationResultsRequest {
	s.TopicCode = &v
	return s
}

func (s *ListEvaluationResultsRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEvaluationResultsRequestFilters struct {
	// Filter condition key. Valid values:
	//
	// - ResourceId: Resource ID.
	//
	// - ResourceName: Resource name.
	//
	// - ResourceType: Resource type.
	//
	// example:
	//
	// ResourceId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// List of filter condition values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListEvaluationResultsRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationResultsRequestFilters) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsRequestFilters) GetKey() *string {
	return s.Key
}

func (s *ListEvaluationResultsRequestFilters) GetValues() []*string {
	return s.Values
}

func (s *ListEvaluationResultsRequestFilters) SetKey(v string) *ListEvaluationResultsRequestFilters {
	s.Key = &v
	return s
}

func (s *ListEvaluationResultsRequestFilters) SetValues(v []*string) *ListEvaluationResultsRequestFilters {
	s.Values = v
	return s
}

func (s *ListEvaluationResultsRequestFilters) Validate() error {
	return dara.Validate(s)
}
