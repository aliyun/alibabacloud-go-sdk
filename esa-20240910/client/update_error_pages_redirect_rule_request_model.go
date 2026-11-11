// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateErrorPagesRedirectRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateErrorPagesRedirectRuleRequest
	GetConfigId() *int64
	SetErrorPagesRedirect(v []*UpdateErrorPagesRedirectRuleRequestErrorPagesRedirect) *UpdateErrorPagesRedirectRuleRequest
	GetErrorPagesRedirect() []*UpdateErrorPagesRedirectRuleRequestErrorPagesRedirect
	SetRule(v string) *UpdateErrorPagesRedirectRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateErrorPagesRedirectRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateErrorPagesRedirectRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateErrorPagesRedirectRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateErrorPagesRedirectRuleRequest
	GetSiteId() *int64
}

type UpdateErrorPagesRedirectRuleRequest struct {
	// The configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The error code redirect configuration.
	ErrorPagesRedirect []*UpdateErrorPagesRedirectRuleRequestErrorPagesRedirect `json:"ErrorPagesRedirect,omitempty" xml:"ErrorPagesRedirect,omitempty" type:"Repeated"`
	// The content of the rule, which uses a conditional expression to match user requests. You do not need to specify this parameter when you add a global configuration. This parameter applies to the following scenarios:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specific requests: Set the value to a custom expression, such as (http.host eq \\"video.example.com\\").
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. You do not need to specify this parameter when you add a global configuration. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the rule. You do not need to specify this parameter when you add a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution sequence of the rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The ID of the site. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5407498413****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateErrorPagesRedirectRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateErrorPagesRedirectRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateErrorPagesRedirectRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateErrorPagesRedirectRuleRequest) GetErrorPagesRedirect() []*UpdateErrorPagesRedirectRuleRequestErrorPagesRedirect {
	return s.ErrorPagesRedirect
}

func (s *UpdateErrorPagesRedirectRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateErrorPagesRedirectRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateErrorPagesRedirectRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateErrorPagesRedirectRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateErrorPagesRedirectRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateErrorPagesRedirectRuleRequest) SetConfigId(v int64) *UpdateErrorPagesRedirectRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleRequest) SetErrorPagesRedirect(v []*UpdateErrorPagesRedirectRuleRequestErrorPagesRedirect) *UpdateErrorPagesRedirectRuleRequest {
	s.ErrorPagesRedirect = v
	return s
}

func (s *UpdateErrorPagesRedirectRuleRequest) SetRule(v string) *UpdateErrorPagesRedirectRuleRequest {
	s.Rule = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleRequest) SetRuleEnable(v string) *UpdateErrorPagesRedirectRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleRequest) SetRuleName(v string) *UpdateErrorPagesRedirectRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleRequest) SetSequence(v int32) *UpdateErrorPagesRedirectRuleRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleRequest) SetSiteId(v int64) *UpdateErrorPagesRedirectRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleRequest) Validate() error {
	if s.ErrorPagesRedirect != nil {
		for _, item := range s.ErrorPagesRedirect {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateErrorPagesRedirectRuleRequestErrorPagesRedirect struct {
	// The HTTP status code used by the node when responding to the client with a redirect address. Valid values:
	//
	// - 400
	//
	// - 403
	//
	// - 404
	//
	// - 405
	//
	// - 414
	//
	// - 416
	//
	// - 500
	//
	// - 501
	//
	// - 502
	//
	// - 503
	//
	// - 504
	//
	// This parameter is required.
	//
	// example:
	//
	// 400
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The destination URL after the redirect.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example.com/test
	TargetURL *string `json:"TargetURL,omitempty" xml:"TargetURL,omitempty"`
}

func (s UpdateErrorPagesRedirectRuleRequestErrorPagesRedirect) String() string {
	return dara.Prettify(s)
}

func (s UpdateErrorPagesRedirectRuleRequestErrorPagesRedirect) GoString() string {
	return s.String()
}

func (s *UpdateErrorPagesRedirectRuleRequestErrorPagesRedirect) GetStatusCode() *string {
	return s.StatusCode
}

func (s *UpdateErrorPagesRedirectRuleRequestErrorPagesRedirect) GetTargetURL() *string {
	return s.TargetURL
}

func (s *UpdateErrorPagesRedirectRuleRequestErrorPagesRedirect) SetStatusCode(v string) *UpdateErrorPagesRedirectRuleRequestErrorPagesRedirect {
	s.StatusCode = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleRequestErrorPagesRedirect) SetTargetURL(v string) *UpdateErrorPagesRedirectRuleRequestErrorPagesRedirect {
	s.TargetURL = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleRequestErrorPagesRedirect) Validate() error {
	return dara.Validate(s)
}
