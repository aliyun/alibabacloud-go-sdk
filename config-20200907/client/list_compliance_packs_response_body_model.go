// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCompliancePacksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePacksResult(v *ListCompliancePacksResponseBodyCompliancePacksResult) *ListCompliancePacksResponseBody
	GetCompliancePacksResult() *ListCompliancePacksResponseBodyCompliancePacksResult
	SetRequestId(v string) *ListCompliancePacksResponseBody
	GetRequestId() *string
}

type ListCompliancePacksResponseBody struct {
	// The query results for the compliance packs.
	CompliancePacksResult *ListCompliancePacksResponseBodyCompliancePacksResult `json:"CompliancePacksResult,omitempty" xml:"CompliancePacksResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCompliancePacksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCompliancePacksResponseBody) GoString() string {
	return s.String()
}

func (s *ListCompliancePacksResponseBody) GetCompliancePacksResult() *ListCompliancePacksResponseBodyCompliancePacksResult {
	return s.CompliancePacksResult
}

func (s *ListCompliancePacksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCompliancePacksResponseBody) SetCompliancePacksResult(v *ListCompliancePacksResponseBodyCompliancePacksResult) *ListCompliancePacksResponseBody {
	s.CompliancePacksResult = v
	return s
}

func (s *ListCompliancePacksResponseBody) SetRequestId(v string) *ListCompliancePacksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCompliancePacksResponseBody) Validate() error {
	if s.CompliancePacksResult != nil {
		if err := s.CompliancePacksResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCompliancePacksResponseBodyCompliancePacksResult struct {
	// A list of compliance packs.
	CompliancePacks []*ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks `json:"CompliancePacks,omitempty" xml:"CompliancePacks,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of compliance packs.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCompliancePacksResponseBodyCompliancePacksResult) String() string {
	return dara.Prettify(s)
}

func (s ListCompliancePacksResponseBodyCompliancePacksResult) GoString() string {
	return s.String()
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResult) GetCompliancePacks() []*ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	return s.CompliancePacks
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResult) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResult) SetCompliancePacks(v []*ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) *ListCompliancePacksResponseBodyCompliancePacksResult {
	s.CompliancePacks = v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResult) SetPageNumber(v int32) *ListCompliancePacksResponseBodyCompliancePacksResult {
	s.PageNumber = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResult) SetPageSize(v int32) *ListCompliancePacksResponseBodyCompliancePacksResult {
	s.PageSize = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResult) SetTotalCount(v int64) *ListCompliancePacksResponseBodyCompliancePacksResult {
	s.TotalCount = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResult) Validate() error {
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

type ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks struct {
	// The ID of the Alibaba Cloud account to which the compliance pack belongs.
	//
	// example:
	//
	// 120886317861****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the compliance pack.
	//
	// example:
	//
	// cp-fdc8626622af00f9****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The name of the compliance pack.
	//
	// example:
	//
	// BestPracticesForResourceStability
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	// The ID of the compliance pack template.
	//
	// example:
	//
	// ct-484cff4e06a30062****
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	// The timestamp when the compliance pack was created. Unit: milliseconds.
	//
	// example:
	//
	// 1621325046000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The description of the compliance pack.
	//
	// example:
	//
	// Check the stability of Alibaba Cloud resources based on the high-availability infrastructure, capacity protection, change management, monitoring management, backup management, and fault isolation. This helps you identify issues at the earliest opportunity and improve stability and the efficiency of O&M.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The risk level of the compliance pack. Valid values:
	//
	// - 1: high risk.
	//
	// - 2: medium risk.
	//
	// - 3: low risk.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The status of the compliance pack. Valid values:
	//
	// - ACTIVE: The compliance pack is active.
	//
	// - CREATING: The compliance pack is being created.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the resource.
	Tags []*ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) String() string {
	return dara.Prettify(s)
}

func (s ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GoString() string {
	return s.String()
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetCompliancePackTemplateId() *string {
	return s.CompliancePackTemplateId
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetDescription() *string {
	return s.Description
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetStatus() *string {
	return s.Status
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GetTags() []*ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags {
	return s.Tags
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetAccountId(v int64) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.AccountId = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetCompliancePackId(v string) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.CompliancePackId = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetCompliancePackName(v string) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.CompliancePackName = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetCompliancePackTemplateId(v string) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetCreateTimestamp(v int64) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.CreateTimestamp = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetDescription(v string) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.Description = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetRiskLevel(v int32) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.RiskLevel = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetStatus(v string) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.Status = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetTags(v []*ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.Tags = v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) Validate() error {
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

type ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags struct {
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

func (s ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags) String() string {
	return dara.Prettify(s)
}

func (s ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags) GoString() string {
	return s.String()
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags) SetTagKey(v string) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags {
	s.TagKey = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags) SetTagValue(v string) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags {
	s.TagValue = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacksTags) Validate() error {
	return dara.Validate(s)
}
