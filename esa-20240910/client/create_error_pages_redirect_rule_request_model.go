// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateErrorPagesRedirectRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetErrorPagesRedirect(v []*CreateErrorPagesRedirectRuleRequestErrorPagesRedirect) *CreateErrorPagesRedirectRuleRequest
	GetErrorPagesRedirect() []*CreateErrorPagesRedirectRuleRequestErrorPagesRedirect
	SetRule(v string) *CreateErrorPagesRedirectRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateErrorPagesRedirectRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateErrorPagesRedirectRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateErrorPagesRedirectRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateErrorPagesRedirectRuleRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateErrorPagesRedirectRuleRequest
	GetSiteVersion() *int32
}

type CreateErrorPagesRedirectRuleRequest struct {
	// The error code redirect configuration.
	//
	// This parameter is required.
	ErrorPagesRedirect []*CreateErrorPagesRedirectRuleRequestErrorPagesRedirect `json:"ErrorPagesRedirect,omitempty" xml:"ErrorPagesRedirect,omitempty" type:"Repeated"`
	// The rule content, which uses conditional expressions to match user requests. You do not need to specify this parameter when adding a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specific requests: Set the value to a custom expression, such as (http.host eq "video.example.com").
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. You do not need to specify this parameter when adding a global configuration. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the rule. You do not need to specify this parameter when adding a global configuration.
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
	// The ID of the site. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 478016908379824
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with configuration version management enabled, you can use this parameter to specify the site version to which the configuration applies. Default value: 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateErrorPagesRedirectRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateErrorPagesRedirectRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateErrorPagesRedirectRuleRequest) GetErrorPagesRedirect() []*CreateErrorPagesRedirectRuleRequestErrorPagesRedirect {
	return s.ErrorPagesRedirect
}

func (s *CreateErrorPagesRedirectRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateErrorPagesRedirectRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateErrorPagesRedirectRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateErrorPagesRedirectRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateErrorPagesRedirectRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateErrorPagesRedirectRuleRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateErrorPagesRedirectRuleRequest) SetErrorPagesRedirect(v []*CreateErrorPagesRedirectRuleRequestErrorPagesRedirect) *CreateErrorPagesRedirectRuleRequest {
	s.ErrorPagesRedirect = v
	return s
}

func (s *CreateErrorPagesRedirectRuleRequest) SetRule(v string) *CreateErrorPagesRedirectRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleRequest) SetRuleEnable(v string) *CreateErrorPagesRedirectRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleRequest) SetRuleName(v string) *CreateErrorPagesRedirectRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleRequest) SetSequence(v int32) *CreateErrorPagesRedirectRuleRequest {
	s.Sequence = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleRequest) SetSiteId(v int64) *CreateErrorPagesRedirectRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleRequest) SetSiteVersion(v int32) *CreateErrorPagesRedirectRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleRequest) Validate() error {
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

type CreateErrorPagesRedirectRuleRequestErrorPagesRedirect struct {
	// The response status code used by the node when responding to the client with a redirect address. Valid values:
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
	// The target URL after redirection.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example.com/test
	TargetURL *string `json:"TargetURL,omitempty" xml:"TargetURL,omitempty"`
}

func (s CreateErrorPagesRedirectRuleRequestErrorPagesRedirect) String() string {
	return dara.Prettify(s)
}

func (s CreateErrorPagesRedirectRuleRequestErrorPagesRedirect) GoString() string {
	return s.String()
}

func (s *CreateErrorPagesRedirectRuleRequestErrorPagesRedirect) GetStatusCode() *string {
	return s.StatusCode
}

func (s *CreateErrorPagesRedirectRuleRequestErrorPagesRedirect) GetTargetURL() *string {
	return s.TargetURL
}

func (s *CreateErrorPagesRedirectRuleRequestErrorPagesRedirect) SetStatusCode(v string) *CreateErrorPagesRedirectRuleRequestErrorPagesRedirect {
	s.StatusCode = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleRequestErrorPagesRedirect) SetTargetURL(v string) *CreateErrorPagesRedirectRuleRequestErrorPagesRedirect {
	s.TargetURL = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleRequestErrorPagesRedirect) Validate() error {
	return dara.Validate(s)
}
