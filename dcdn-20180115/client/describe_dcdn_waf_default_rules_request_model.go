// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafDefaultRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQueryArgs(v string) *DescribeDcdnWafDefaultRulesRequest
	GetQueryArgs() *string
}

type DescribeDcdnWafDefaultRulesRequest struct {
	// The query conditions. The value is a string in the JSON format. Format: `QueryArgs={"DefenseScene":"anti_scan"}`
	//
	// example:
	//
	// {"DefenseScene":"anti_scan"}
	QueryArgs *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
}

func (s DescribeDcdnWafDefaultRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDefaultRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDefaultRulesRequest) GetQueryArgs() *string {
	return s.QueryArgs
}

func (s *DescribeDcdnWafDefaultRulesRequest) SetQueryArgs(v string) *DescribeDcdnWafDefaultRulesRequest {
	s.QueryArgs = &v
	return s
}

func (s *DescribeDcdnWafDefaultRulesRequest) Validate() error {
	return dara.Validate(s)
}
