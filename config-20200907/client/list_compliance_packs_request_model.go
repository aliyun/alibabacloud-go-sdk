// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCompliancePacksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListCompliancePacksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCompliancePacksRequest
	GetPageSize() *int32
	SetRiskLevel(v int32) *ListCompliancePacksRequest
	GetRiskLevel() *int32
	SetStatus(v string) *ListCompliancePacksRequest
	GetStatus() *string
	SetTag(v []*ListCompliancePacksRequestTag) *ListCompliancePacksRequest
	GetTag() []*ListCompliancePacksRequestTag
}

type ListCompliancePacksRequest struct {
	// The page number.
	//
	// Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	//
	// You can attach up to 20 tags to a resource.
	Tag []*ListCompliancePacksRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListCompliancePacksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCompliancePacksRequest) GoString() string {
	return s.String()
}

func (s *ListCompliancePacksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCompliancePacksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCompliancePacksRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListCompliancePacksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCompliancePacksRequest) GetTag() []*ListCompliancePacksRequestTag {
	return s.Tag
}

func (s *ListCompliancePacksRequest) SetPageNumber(v int32) *ListCompliancePacksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCompliancePacksRequest) SetPageSize(v int32) *ListCompliancePacksRequest {
	s.PageSize = &v
	return s
}

func (s *ListCompliancePacksRequest) SetRiskLevel(v int32) *ListCompliancePacksRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListCompliancePacksRequest) SetStatus(v string) *ListCompliancePacksRequest {
	s.Status = &v
	return s
}

func (s *ListCompliancePacksRequest) SetTag(v []*ListCompliancePacksRequestTag) *ListCompliancePacksRequest {
	s.Tag = v
	return s
}

func (s *ListCompliancePacksRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCompliancePacksRequestTag struct {
	// The tag key of the resource.
	//
	// You can attach up to 20 tag keys to a resource.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// You can attach up to 20 tag values to a resource.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListCompliancePacksRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListCompliancePacksRequestTag) GoString() string {
	return s.String()
}

func (s *ListCompliancePacksRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListCompliancePacksRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListCompliancePacksRequestTag) SetKey(v string) *ListCompliancePacksRequestTag {
	s.Key = &v
	return s
}

func (s *ListCompliancePacksRequestTag) SetValue(v string) *ListCompliancePacksRequestTag {
	s.Value = &v
	return s
}

func (s *ListCompliancePacksRequestTag) Validate() error {
	return dara.Validate(s)
}
