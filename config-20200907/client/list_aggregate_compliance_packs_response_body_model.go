// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateCompliancePacksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePacksResult(v *ListAggregateCompliancePacksResponseBodyCompliancePacksResult) *ListAggregateCompliancePacksResponseBody
	GetCompliancePacksResult() *ListAggregateCompliancePacksResponseBodyCompliancePacksResult
	SetRequestId(v string) *ListAggregateCompliancePacksResponseBody
	GetRequestId() *string
}

type ListAggregateCompliancePacksResponseBody struct {
	// The compliance packages returned.
	CompliancePacksResult *ListAggregateCompliancePacksResponseBodyCompliancePacksResult `json:"CompliancePacksResult,omitempty" xml:"CompliancePacksResult,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B5806142-3090-4F86-A84E-12B3FE52C1C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateCompliancePacksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateCompliancePacksResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateCompliancePacksResponseBody) GetCompliancePacksResult() *ListAggregateCompliancePacksResponseBodyCompliancePacksResult {
	return s.CompliancePacksResult
}

func (s *ListAggregateCompliancePacksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAggregateCompliancePacksResponseBody) SetCompliancePacksResult(v *ListAggregateCompliancePacksResponseBodyCompliancePacksResult) *ListAggregateCompliancePacksResponseBody {
	s.CompliancePacksResult = v
	return s
}

func (s *ListAggregateCompliancePacksResponseBody) SetRequestId(v string) *ListAggregateCompliancePacksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBody) Validate() error {
	if s.CompliancePacksResult != nil {
		if err := s.CompliancePacksResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAggregateCompliancePacksResponseBodyCompliancePacksResult struct {
	// The compliance packages.
	CompliancePacks []*ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks `json:"CompliancePacks,omitempty" xml:"CompliancePacks,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of compliance packages returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAggregateCompliancePacksResponseBodyCompliancePacksResult) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateCompliancePacksResponseBodyCompliancePacksResult) GoString() string {
	return s.String()
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResult) GetCompliancePacks() []*ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	return s.CompliancePacks
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResult) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResult) SetCompliancePacks(v []*ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) *ListAggregateCompliancePacksResponseBodyCompliancePacksResult {
	s.CompliancePacks = v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResult) SetPageNumber(v int32) *ListAggregateCompliancePacksResponseBodyCompliancePacksResult {
	s.PageNumber = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResult) SetPageSize(v int32) *ListAggregateCompliancePacksResponseBodyCompliancePacksResult {
	s.PageSize = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResult) SetTotalCount(v int64) *ListAggregateCompliancePacksResponseBodyCompliancePacksResult {
	s.TotalCount = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResult) Validate() error {
	if s.CompliancePacks != nil {
		for _, item := range s.CompliancePacks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks struct {
	// The ID of the management account to which the compliance package belongs.
	//
	// example:
	//
	// 100931896542****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the account group.
	//
	// example:
	//
	// ca-f632626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the compliance package.
	//
	// example:
	//
	// cp-fdc8626622af00f9****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The name of the compliance package.
	//
	// example:
	//
	// example-name
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	// The ID of the compliance package template.
	//
	// example:
	//
	// ct-5f26ff4e06a300c4****
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	// The timestamp when the compliance package was created. Unit: milliseconds.
	//
	// example:
	//
	// 1624243657000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The description of the compliance package.
	//
	// example:
	//
	// example-description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The risk level of the resources that are not compliant with the managed rules in the compliance package. Valid values:
	//
	// 	- 1: high risk level.
	//
	// 	- 2: medium risk level.
	//
	// 	- 3: low risk level.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The status of the compliance package. Valid values:
	//
	// 	- ACTIVE: The compliance package is available for use.
	//
	// 	- CREATING: The compliance package is being created.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GoString() string {
	return s.String()
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetCompliancePackTemplateId() *string {
	return s.CompliancePackTemplateId
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetDescription() *string {
	return s.Description
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetStatus() *string {
	return s.Status
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetTags() []*ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags {
	return s.Tags
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetAccountId(v int64) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.AccountId = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetAggregatorId(v string) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetCompliancePackId(v string) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.CompliancePackId = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetCompliancePackName(v string) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.CompliancePackName = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetCompliancePackTemplateId(v string) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetCreateTimestamp(v int64) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.CreateTimestamp = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetDescription(v string) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.Description = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetRiskLevel(v int32) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.RiskLevel = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetStatus(v string) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.Status = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetTags(v []*ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.Tags = v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags struct {
	// The tag key.
	//
	// example:
	//
	// key-1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value-1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags) GoString() string {
	return s.String()
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags) SetTagKey(v string) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags {
	s.TagKey = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags) SetTagValue(v string) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags {
	s.TagValue = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags) Validate() error {
	return dara.Validate(s)
}
