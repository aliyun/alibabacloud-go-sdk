// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateCompliancePacksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *ListAggregateCompliancePacksRequest
	GetAggregatorId() *string
	SetPageNumber(v int32) *ListAggregateCompliancePacksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAggregateCompliancePacksRequest
	GetPageSize() *int32
	SetRiskLevel(v int32) *ListAggregateCompliancePacksRequest
	GetRiskLevel() *int32
	SetStatus(v string) *ListAggregateCompliancePacksRequest
	GetStatus() *string
	SetTag(v []*ListAggregateCompliancePacksRequestTag) *ListAggregateCompliancePacksRequest
	GetTag() []*ListAggregateCompliancePacksRequestTag
}

type ListAggregateCompliancePacksRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-f632626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The status of the one or more compliance packages to be queried. Valid values:
	//
	// 	- ACTIVE: compliance packages that are available for use.
	//
	// 	- CREATING: compliance packages that are being created.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tag []*ListAggregateCompliancePacksRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAggregateCompliancePacksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateCompliancePacksRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateCompliancePacksRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateCompliancePacksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAggregateCompliancePacksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAggregateCompliancePacksRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListAggregateCompliancePacksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAggregateCompliancePacksRequest) GetTag() []*ListAggregateCompliancePacksRequestTag {
	return s.Tag
}

func (s *ListAggregateCompliancePacksRequest) SetAggregatorId(v string) *ListAggregateCompliancePacksRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateCompliancePacksRequest) SetPageNumber(v int32) *ListAggregateCompliancePacksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAggregateCompliancePacksRequest) SetPageSize(v int32) *ListAggregateCompliancePacksRequest {
	s.PageSize = &v
	return s
}

func (s *ListAggregateCompliancePacksRequest) SetRiskLevel(v int32) *ListAggregateCompliancePacksRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListAggregateCompliancePacksRequest) SetStatus(v string) *ListAggregateCompliancePacksRequest {
	s.Status = &v
	return s
}

func (s *ListAggregateCompliancePacksRequest) SetTag(v []*ListAggregateCompliancePacksRequestTag) *ListAggregateCompliancePacksRequest {
	s.Tag = v
	return s
}

func (s *ListAggregateCompliancePacksRequest) Validate() error {
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

type ListAggregateCompliancePacksRequestTag struct {
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

func (s ListAggregateCompliancePacksRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateCompliancePacksRequestTag) GoString() string {
	return s.String()
}

func (s *ListAggregateCompliancePacksRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListAggregateCompliancePacksRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListAggregateCompliancePacksRequestTag) SetKey(v string) *ListAggregateCompliancePacksRequestTag {
	s.Key = &v
	return s
}

func (s *ListAggregateCompliancePacksRequestTag) SetValue(v string) *ListAggregateCompliancePacksRequestTag {
	s.Value = &v
	return s
}

func (s *ListAggregateCompliancePacksRequestTag) Validate() error {
	return dara.Validate(s)
}
