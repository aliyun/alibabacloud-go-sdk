// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserWafRulesetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListUserWafRulesetsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListUserWafRulesetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserWafRulesetsRequest
	GetPageSize() *int32
	SetPhase(v string) *ListUserWafRulesetsRequest
	GetPhase() *string
	SetQueryArgs(v *ListUserWafRulesetsRequestQueryArgs) *ListUserWafRulesetsRequest
	GetQueryArgs() *ListUserWafRulesetsRequestQueryArgs
}

type ListUserWafRulesetsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// esa-xxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies the execution phase of the WAF rule.
	//
	// - `http_whitelist`: whitelist rule
	//
	// - `http_custom`: custom rule
	//
	// - `http_managed`: managed rule
	//
	// - `http_anti_scan`: anti-scan rule
	//
	// - `http_ratelimit`: rate limit rule
	//
	// - `ip_access_rule`: IP access rule
	//
	// - `http_bot`: advanced mode bot
	//
	// - `http_security_level_rule`: security rule
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// Parameters for filtering and sorting the results.
	QueryArgs *ListUserWafRulesetsRequestQueryArgs `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty" type:"Struct"`
}

func (s ListUserWafRulesetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserWafRulesetsRequest) GoString() string {
	return s.String()
}

func (s *ListUserWafRulesetsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserWafRulesetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserWafRulesetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserWafRulesetsRequest) GetPhase() *string {
	return s.Phase
}

func (s *ListUserWafRulesetsRequest) GetQueryArgs() *ListUserWafRulesetsRequestQueryArgs {
	return s.QueryArgs
}

func (s *ListUserWafRulesetsRequest) SetInstanceId(v string) *ListUserWafRulesetsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserWafRulesetsRequest) SetPageNumber(v int32) *ListUserWafRulesetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserWafRulesetsRequest) SetPageSize(v int32) *ListUserWafRulesetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserWafRulesetsRequest) SetPhase(v string) *ListUserWafRulesetsRequest {
	s.Phase = &v
	return s
}

func (s *ListUserWafRulesetsRequest) SetQueryArgs(v *ListUserWafRulesetsRequestQueryArgs) *ListUserWafRulesetsRequest {
	s.QueryArgs = v
	return s
}

func (s *ListUserWafRulesetsRequest) Validate() error {
	if s.QueryArgs != nil {
		if err := s.QueryArgs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserWafRulesetsRequestQueryArgs struct {
	// Specifies whether to sort results in descending order.
	Desc *bool `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// A string for a fuzzy search of ruleset names.
	//
	// example:
	//
	// example
	NameLike *string `json:"NameLike,omitempty" xml:"NameLike,omitempty"`
	// The field by which to sort results.
	//
	// example:
	//
	// id
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
}

func (s ListUserWafRulesetsRequestQueryArgs) String() string {
	return dara.Prettify(s)
}

func (s ListUserWafRulesetsRequestQueryArgs) GoString() string {
	return s.String()
}

func (s *ListUserWafRulesetsRequestQueryArgs) GetDesc() *bool {
	return s.Desc
}

func (s *ListUserWafRulesetsRequestQueryArgs) GetNameLike() *string {
	return s.NameLike
}

func (s *ListUserWafRulesetsRequestQueryArgs) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListUserWafRulesetsRequestQueryArgs) SetDesc(v bool) *ListUserWafRulesetsRequestQueryArgs {
	s.Desc = &v
	return s
}

func (s *ListUserWafRulesetsRequestQueryArgs) SetNameLike(v string) *ListUserWafRulesetsRequestQueryArgs {
	s.NameLike = &v
	return s
}

func (s *ListUserWafRulesetsRequestQueryArgs) SetOrderBy(v string) *ListUserWafRulesetsRequestQueryArgs {
	s.OrderBy = &v
	return s
}

func (s *ListUserWafRulesetsRequestQueryArgs) Validate() error {
	return dara.Validate(s)
}
