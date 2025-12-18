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
	// Pages start from page 1. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100. Minimum value: 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The status of the compliance package to be queried. Valid values:
	//
	// 	- ACTIVE: The compliance package is active.
	//
	// 	- CREATING: The compliance package is being created.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
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
	// The tag key of the resource. You can specify up to 20 tag keys.
	//
	// The tag key cannot be an empty string. The tag key must be 1 to 64 characters in length and cannot start with `aliyun` or `acs`:. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values.
	//
	// The tag values can be an empty string or up to 128 characters in length. The tag values cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// Each key-value must be unique. You can specify at most 20 tag values in each call.
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
