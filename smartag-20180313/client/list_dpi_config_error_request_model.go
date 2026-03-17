// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDpiConfigErrorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDpiConfigType(v string) *ListDpiConfigErrorRequest
	GetDpiConfigType() *string
	SetMaxResults(v int32) *ListDpiConfigErrorRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDpiConfigErrorRequest
	GetNextToken() *string
	SetRegionId(v string) *ListDpiConfigErrorRequest
	GetRegionId() *string
	SetRuleInstanceId(v string) *ListDpiConfigErrorRequest
	GetRuleInstanceId() *string
	SetSmartAGId(v string) *ListDpiConfigErrorRequest
	GetSmartAGId() *string
}

type ListDpiConfigErrorRequest struct {
	// The type of the instance for which the DPI feature is configured. Valid values:
	//
	// 	- **acl**
	//
	// 	- **qos**
	//
	// This parameter is required.
	//
	// example:
	//
	// qos
	DpiConfigType *string `json:"DpiConfigType,omitempty" xml:"DpiConfigType,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: **1*	- to **100**.
	//
	// Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to query the next page.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the SAG instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the instance for which the DPI feature is configured.
	//
	// example:
	//
	// qos-1strcafl4wghpb****
	RuleInstanceId *string `json:"RuleInstanceId,omitempty" xml:"RuleInstanceId,omitempty"`
	// The ID of the SAG instance.
	//
	// example:
	//
	// sag-1e8sgws6b133b8****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s ListDpiConfigErrorRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDpiConfigErrorRequest) GoString() string {
	return s.String()
}

func (s *ListDpiConfigErrorRequest) GetDpiConfigType() *string {
	return s.DpiConfigType
}

func (s *ListDpiConfigErrorRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDpiConfigErrorRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDpiConfigErrorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDpiConfigErrorRequest) GetRuleInstanceId() *string {
	return s.RuleInstanceId
}

func (s *ListDpiConfigErrorRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ListDpiConfigErrorRequest) SetDpiConfigType(v string) *ListDpiConfigErrorRequest {
	s.DpiConfigType = &v
	return s
}

func (s *ListDpiConfigErrorRequest) SetMaxResults(v int32) *ListDpiConfigErrorRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDpiConfigErrorRequest) SetNextToken(v string) *ListDpiConfigErrorRequest {
	s.NextToken = &v
	return s
}

func (s *ListDpiConfigErrorRequest) SetRegionId(v string) *ListDpiConfigErrorRequest {
	s.RegionId = &v
	return s
}

func (s *ListDpiConfigErrorRequest) SetRuleInstanceId(v string) *ListDpiConfigErrorRequest {
	s.RuleInstanceId = &v
	return s
}

func (s *ListDpiConfigErrorRequest) SetSmartAGId(v string) *ListDpiConfigErrorRequest {
	s.SmartAGId = &v
	return s
}

func (s *ListDpiConfigErrorRequest) Validate() error {
	return dara.Validate(s)
}
