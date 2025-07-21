// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListNetworkRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNetworkRulesRequest
	GetPageSize() *int32
}

type ListNetworkRulesRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListNetworkRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkRulesRequest) GoString() string {
	return s.String()
}

func (s *ListNetworkRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNetworkRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNetworkRulesRequest) SetPageNumber(v int32) *ListNetworkRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNetworkRulesRequest) SetPageSize(v int32) *ListNetworkRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNetworkRulesRequest) Validate() error {
	return dara.Validate(s)
}
