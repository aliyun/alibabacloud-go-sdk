// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErrorPagesRedirectRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetErrorPagesRedirectRuleResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetErrorPagesRedirectRuleResponseBody
	GetConfigType() *string
	SetErrorPagesRedirect(v []*GetErrorPagesRedirectRuleResponseBodyErrorPagesRedirect) *GetErrorPagesRedirectRuleResponseBody
	GetErrorPagesRedirect() []*GetErrorPagesRedirectRuleResponseBodyErrorPagesRedirect
	SetRequestId(v string) *GetErrorPagesRedirectRuleResponseBody
	GetRequestId() *string
	SetRule(v string) *GetErrorPagesRedirectRuleResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetErrorPagesRedirectRuleResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetErrorPagesRedirectRuleResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetErrorPagesRedirectRuleResponseBody
	GetSequence() *int32
	SetSiteVersion(v int32) *GetErrorPagesRedirectRuleResponseBody
	GetSiteVersion() *int32
}

type GetErrorPagesRedirectRuleResponseBody struct {
	// The configuration ID. You can obtain the configuration ID by calling the [ListCompressionRules](https://help.aliyun.com/document_detail/2867498.html) operation.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. You can use this parameter to query global configurations or rule configurations. Valid values:
	//
	// - global: Queries global configurations.
	//
	// - rule: Queries rule configurations.
	//
	// This parameter is optional. If you do not specify this parameter, global and rule configurations are not distinguished. This parameter takes effect only when the functionName parameter is specified.
	//
	// example:
	//
	// rule
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The error code redirect.
	ErrorPagesRedirect []*GetErrorPagesRedirectRuleResponseBodyErrorPagesRedirect `json:"ErrorPagesRedirect,omitempty" xml:"ErrorPagesRedirect,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-280B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The rule content, which uses a conditional expression to match user requests. You do not need to set this parameter when adding a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specified requests: Set the value to a custom expression, such as (http.host eq "video.example.com").
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. You do not need to set this parameter when adding a global configuration. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. You do not need to set this parameter when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The version number of the site configuration. For sites with configuration version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetErrorPagesRedirectRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetErrorPagesRedirectRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetErrorPagesRedirectRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetErrorPagesRedirectRuleResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetErrorPagesRedirectRuleResponseBody) GetErrorPagesRedirect() []*GetErrorPagesRedirectRuleResponseBodyErrorPagesRedirect {
	return s.ErrorPagesRedirect
}

func (s *GetErrorPagesRedirectRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetErrorPagesRedirectRuleResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetErrorPagesRedirectRuleResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetErrorPagesRedirectRuleResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetErrorPagesRedirectRuleResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetErrorPagesRedirectRuleResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetErrorPagesRedirectRuleResponseBody) SetConfigId(v int64) *GetErrorPagesRedirectRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetErrorPagesRedirectRuleResponseBody) SetConfigType(v string) *GetErrorPagesRedirectRuleResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetErrorPagesRedirectRuleResponseBody) SetErrorPagesRedirect(v []*GetErrorPagesRedirectRuleResponseBodyErrorPagesRedirect) *GetErrorPagesRedirectRuleResponseBody {
	s.ErrorPagesRedirect = v
	return s
}

func (s *GetErrorPagesRedirectRuleResponseBody) SetRequestId(v string) *GetErrorPagesRedirectRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetErrorPagesRedirectRuleResponseBody) SetRule(v string) *GetErrorPagesRedirectRuleResponseBody {
	s.Rule = &v
	return s
}

func (s *GetErrorPagesRedirectRuleResponseBody) SetRuleEnable(v string) *GetErrorPagesRedirectRuleResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetErrorPagesRedirectRuleResponseBody) SetRuleName(v string) *GetErrorPagesRedirectRuleResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetErrorPagesRedirectRuleResponseBody) SetSequence(v int32) *GetErrorPagesRedirectRuleResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetErrorPagesRedirectRuleResponseBody) SetSiteVersion(v int32) *GetErrorPagesRedirectRuleResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetErrorPagesRedirectRuleResponseBody) Validate() error {
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

type GetErrorPagesRedirectRuleResponseBodyErrorPagesRedirect struct {
	// The response status code used when the node responds to the client with a redirect address. Valid values:
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
	// example:
	//
	// 400
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The target URL after redirection.
	//
	// example:
	//
	// http://example.com/test
	TargetURL *string `json:"TargetURL,omitempty" xml:"TargetURL,omitempty"`
}

func (s GetErrorPagesRedirectRuleResponseBodyErrorPagesRedirect) String() string {
	return dara.Prettify(s)
}

func (s GetErrorPagesRedirectRuleResponseBodyErrorPagesRedirect) GoString() string {
	return s.String()
}

func (s *GetErrorPagesRedirectRuleResponseBodyErrorPagesRedirect) GetStatusCode() *string {
	return s.StatusCode
}

func (s *GetErrorPagesRedirectRuleResponseBodyErrorPagesRedirect) GetTargetURL() *string {
	return s.TargetURL
}

func (s *GetErrorPagesRedirectRuleResponseBodyErrorPagesRedirect) SetStatusCode(v string) *GetErrorPagesRedirectRuleResponseBodyErrorPagesRedirect {
	s.StatusCode = &v
	return s
}

func (s *GetErrorPagesRedirectRuleResponseBodyErrorPagesRedirect) SetTargetURL(v string) *GetErrorPagesRedirectRuleResponseBodyErrorPagesRedirect {
	s.TargetURL = &v
	return s
}

func (s *GetErrorPagesRedirectRuleResponseBodyErrorPagesRedirect) Validate() error {
	return dara.Validate(s)
}
