// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemediationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleIds(v string) *ListRemediationsRequest
	GetConfigRuleIds() *string
	SetPageNumber(v int64) *ListRemediationsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRemediationsRequest
	GetPageSize() *int64
}

type ListRemediationsRequest struct {
	// The rule IDs. Separate multiple rule IDs with commas (,).
	//
	// For more information about how to obtain the ID of a rule, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
	//
	// example:
	//
	// cr-6b7c626622af00b4****
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRemediationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRemediationsRequest) GoString() string {
	return s.String()
}

func (s *ListRemediationsRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *ListRemediationsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRemediationsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRemediationsRequest) SetConfigRuleIds(v string) *ListRemediationsRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *ListRemediationsRequest) SetPageNumber(v int64) *ListRemediationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRemediationsRequest) SetPageSize(v int64) *ListRemediationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRemediationsRequest) Validate() error {
	return dara.Validate(s)
}
