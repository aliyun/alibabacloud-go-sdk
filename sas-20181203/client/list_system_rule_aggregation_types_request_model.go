// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemRuleAggregationTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListSystemRuleAggregationTypesRequest
	GetLang() *string
}

type ListSystemRuleAggregationTypesRequest struct {
	// The language of the content within the request and response. Default value: zh. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ListSystemRuleAggregationTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSystemRuleAggregationTypesRequest) GoString() string {
	return s.String()
}

func (s *ListSystemRuleAggregationTypesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListSystemRuleAggregationTypesRequest) SetLang(v string) *ListSystemRuleAggregationTypesRequest {
	s.Lang = &v
	return s
}

func (s *ListSystemRuleAggregationTypesRequest) Validate() error {
	return dara.Validate(s)
}
