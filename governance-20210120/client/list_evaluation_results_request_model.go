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
	// The Alibaba Cloud account ID of the member. This parameter takes effect only when a multi-account governance maturity check is performed.
	//
	// example:
	//
	// 176618589410****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The filter conditions.
	Filters []*ListEvaluationResultsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// example:
	//
	// basic
	LensCode *string `json:"LensCode,omitempty" xml:"LensCode,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ResourceDirectory
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// es-bp1r**************
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
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
	// The key of the filter condition. Valid values:
	//
	// 	- ResourceId: the resource ID.
	//
	// 	- ResourceName: the name of the resource.
	//
	// 	- ResourceType: the resource type.
	//
	// example:
	//
	// ResourceId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The list of filter condition values.
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
