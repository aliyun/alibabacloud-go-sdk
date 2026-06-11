// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketHttpRoute interface {
	dara.Model
	String() string
	GoString() string
	SetBuiltin(v bool) *HiMarketHttpRoute
	GetBuiltin() *bool
	SetDescription(v string) *HiMarketHttpRoute
	GetDescription() *string
	SetDomains(v []*HiMarketDomain) *HiMarketHttpRoute
	GetDomains() []*HiMarketDomain
	SetMatch(v *HiMarketHttpRouteMatch) *HiMarketHttpRoute
	GetMatch() *HiMarketHttpRouteMatch
}

type HiMarketHttpRoute struct {
	// Indicates whether this is a system-defined route. Users cannot modify or delete built-in routes. Defaults to `false`.
	Builtin *bool `json:"builtin,omitempty" xml:"builtin,omitempty"`
	// An optional description for the HTTP route. This field is for informational purposes only.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// A list of hostnames to which this route applies. The request\\"s `Host` header must match one of the hostnames in this list.
	Domains []*HiMarketDomain `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
	// Defines the matching criteria for an incoming HTTP request. The request must meet all specified conditions for this route to apply.
	Match *HiMarketHttpRouteMatch `json:"match,omitempty" xml:"match,omitempty" type:"Struct"`
}

func (s HiMarketHttpRoute) String() string {
	return dara.Prettify(s)
}

func (s HiMarketHttpRoute) GoString() string {
	return s.String()
}

func (s *HiMarketHttpRoute) GetBuiltin() *bool {
	return s.Builtin
}

func (s *HiMarketHttpRoute) GetDescription() *string {
	return s.Description
}

func (s *HiMarketHttpRoute) GetDomains() []*HiMarketDomain {
	return s.Domains
}

func (s *HiMarketHttpRoute) GetMatch() *HiMarketHttpRouteMatch {
	return s.Match
}

func (s *HiMarketHttpRoute) SetBuiltin(v bool) *HiMarketHttpRoute {
	s.Builtin = &v
	return s
}

func (s *HiMarketHttpRoute) SetDescription(v string) *HiMarketHttpRoute {
	s.Description = &v
	return s
}

func (s *HiMarketHttpRoute) SetDomains(v []*HiMarketDomain) *HiMarketHttpRoute {
	s.Domains = v
	return s
}

func (s *HiMarketHttpRoute) SetMatch(v *HiMarketHttpRouteMatch) *HiMarketHttpRoute {
	s.Match = v
	return s
}

func (s *HiMarketHttpRoute) Validate() error {
	if s.Domains != nil {
		for _, item := range s.Domains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HiMarketHttpRouteMatch struct {
	// A list of HTTP header match conditions. The request must match all of these conditions.
	Headers []*HiMarketHttpRouteMatchHeaders `json:"headers,omitempty" xml:"headers,omitempty" type:"Repeated"`
	// A list of HTTP methods to match, such as `GET` or `POST`. If this field is not specified, the route matches requests with any HTTP method.
	Methods []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
	// A list of conditions for matching against a data model. Use this to validate the request body or other structured data.
	ModelMatches []*HiMarketHttpRouteMatchModelMatches `json:"modelMatches,omitempty" xml:"modelMatches,omitempty" type:"Repeated"`
	// Specifies the conditions for matching the request path.
	Path *HiMarketHttpRouteMatchPath `json:"path,omitempty" xml:"path,omitempty" type:"Struct"`
	// A list of URL query parameter match conditions. The request must match all of these conditions.
	QueryParams []*HiMarketHttpRouteMatchQueryParams `json:"queryParams,omitempty" xml:"queryParams,omitempty" type:"Repeated"`
}

func (s HiMarketHttpRouteMatch) String() string {
	return dara.Prettify(s)
}

func (s HiMarketHttpRouteMatch) GoString() string {
	return s.String()
}

func (s *HiMarketHttpRouteMatch) GetHeaders() []*HiMarketHttpRouteMatchHeaders {
	return s.Headers
}

func (s *HiMarketHttpRouteMatch) GetMethods() []*string {
	return s.Methods
}

func (s *HiMarketHttpRouteMatch) GetModelMatches() []*HiMarketHttpRouteMatchModelMatches {
	return s.ModelMatches
}

func (s *HiMarketHttpRouteMatch) GetPath() *HiMarketHttpRouteMatchPath {
	return s.Path
}

func (s *HiMarketHttpRouteMatch) GetQueryParams() []*HiMarketHttpRouteMatchQueryParams {
	return s.QueryParams
}

func (s *HiMarketHttpRouteMatch) SetHeaders(v []*HiMarketHttpRouteMatchHeaders) *HiMarketHttpRouteMatch {
	s.Headers = v
	return s
}

func (s *HiMarketHttpRouteMatch) SetMethods(v []*string) *HiMarketHttpRouteMatch {
	s.Methods = v
	return s
}

func (s *HiMarketHttpRouteMatch) SetModelMatches(v []*HiMarketHttpRouteMatchModelMatches) *HiMarketHttpRouteMatch {
	s.ModelMatches = v
	return s
}

func (s *HiMarketHttpRouteMatch) SetPath(v *HiMarketHttpRouteMatchPath) *HiMarketHttpRouteMatch {
	s.Path = v
	return s
}

func (s *HiMarketHttpRouteMatch) SetQueryParams(v []*HiMarketHttpRouteMatchQueryParams) *HiMarketHttpRouteMatch {
	s.QueryParams = v
	return s
}

func (s *HiMarketHttpRouteMatch) Validate() error {
	if s.Headers != nil {
		for _, item := range s.Headers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ModelMatches != nil {
		for _, item := range s.ModelMatches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Path != nil {
		if err := s.Path.Validate(); err != nil {
			return err
		}
	}
	if s.QueryParams != nil {
		for _, item := range s.QueryParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HiMarketHttpRouteMatchHeaders struct {
	// Specifies whether the header match is case-sensitive. Defaults to `true`.
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// The name of the HTTP header to match, such as `Content-Type`.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of header match. Valid values are `Exact` and `RegularExpression`. Defaults to `Exact`.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value to match against the header. The match `type` determines how this value is interpreted.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HiMarketHttpRouteMatchHeaders) String() string {
	return dara.Prettify(s)
}

func (s HiMarketHttpRouteMatchHeaders) GoString() string {
	return s.String()
}

func (s *HiMarketHttpRouteMatchHeaders) GetCaseSensitive() *bool {
	return s.CaseSensitive
}

func (s *HiMarketHttpRouteMatchHeaders) GetName() *string {
	return s.Name
}

func (s *HiMarketHttpRouteMatchHeaders) GetType() *string {
	return s.Type
}

func (s *HiMarketHttpRouteMatchHeaders) GetValue() *string {
	return s.Value
}

func (s *HiMarketHttpRouteMatchHeaders) SetCaseSensitive(v bool) *HiMarketHttpRouteMatchHeaders {
	s.CaseSensitive = &v
	return s
}

func (s *HiMarketHttpRouteMatchHeaders) SetName(v string) *HiMarketHttpRouteMatchHeaders {
	s.Name = &v
	return s
}

func (s *HiMarketHttpRouteMatchHeaders) SetType(v string) *HiMarketHttpRouteMatchHeaders {
	s.Type = &v
	return s
}

func (s *HiMarketHttpRouteMatchHeaders) SetValue(v string) *HiMarketHttpRouteMatchHeaders {
	s.Value = &v
	return s
}

func (s *HiMarketHttpRouteMatchHeaders) Validate() error {
	return dara.Validate(s)
}

type HiMarketHttpRouteMatchModelMatches struct {
	// Specifies whether the model field match is case-sensitive. Defaults to `true`.
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// The name of the model field to match.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of match, such as `Exact`, `Pattern`, or `Range`.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value or pattern to match against the model field.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HiMarketHttpRouteMatchModelMatches) String() string {
	return dara.Prettify(s)
}

func (s HiMarketHttpRouteMatchModelMatches) GoString() string {
	return s.String()
}

func (s *HiMarketHttpRouteMatchModelMatches) GetCaseSensitive() *bool {
	return s.CaseSensitive
}

func (s *HiMarketHttpRouteMatchModelMatches) GetName() *string {
	return s.Name
}

func (s *HiMarketHttpRouteMatchModelMatches) GetType() *string {
	return s.Type
}

func (s *HiMarketHttpRouteMatchModelMatches) GetValue() *string {
	return s.Value
}

func (s *HiMarketHttpRouteMatchModelMatches) SetCaseSensitive(v bool) *HiMarketHttpRouteMatchModelMatches {
	s.CaseSensitive = &v
	return s
}

func (s *HiMarketHttpRouteMatchModelMatches) SetName(v string) *HiMarketHttpRouteMatchModelMatches {
	s.Name = &v
	return s
}

func (s *HiMarketHttpRouteMatchModelMatches) SetType(v string) *HiMarketHttpRouteMatchModelMatches {
	s.Type = &v
	return s
}

func (s *HiMarketHttpRouteMatchModelMatches) SetValue(v string) *HiMarketHttpRouteMatchModelMatches {
	s.Value = &v
	return s
}

func (s *HiMarketHttpRouteMatchModelMatches) Validate() error {
	return dara.Validate(s)
}

type HiMarketHttpRouteMatchPath struct {
	// Specifies whether the path match is case-sensitive. Defaults to `true`.
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// The type of path match. Valid values are `Exact` and `Prefix`. Defaults to `Exact` if not specified.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The path value to match. The specified `type` determines how this value is interpreted.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HiMarketHttpRouteMatchPath) String() string {
	return dara.Prettify(s)
}

func (s HiMarketHttpRouteMatchPath) GoString() string {
	return s.String()
}

func (s *HiMarketHttpRouteMatchPath) GetCaseSensitive() *bool {
	return s.CaseSensitive
}

func (s *HiMarketHttpRouteMatchPath) GetType() *string {
	return s.Type
}

func (s *HiMarketHttpRouteMatchPath) GetValue() *string {
	return s.Value
}

func (s *HiMarketHttpRouteMatchPath) SetCaseSensitive(v bool) *HiMarketHttpRouteMatchPath {
	s.CaseSensitive = &v
	return s
}

func (s *HiMarketHttpRouteMatchPath) SetType(v string) *HiMarketHttpRouteMatchPath {
	s.Type = &v
	return s
}

func (s *HiMarketHttpRouteMatchPath) SetValue(v string) *HiMarketHttpRouteMatchPath {
	s.Value = &v
	return s
}

func (s *HiMarketHttpRouteMatchPath) Validate() error {
	return dara.Validate(s)
}

type HiMarketHttpRouteMatchQueryParams struct {
	// Specifies whether the query parameter match is case-sensitive. Defaults to `true`.
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// The name of the query parameter to match.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of query parameter match. Valid values are `Exact` and `RegularExpression`. Defaults to `Exact`.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value to match against the query parameter. The match `type` determines how this value is interpreted.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HiMarketHttpRouteMatchQueryParams) String() string {
	return dara.Prettify(s)
}

func (s HiMarketHttpRouteMatchQueryParams) GoString() string {
	return s.String()
}

func (s *HiMarketHttpRouteMatchQueryParams) GetCaseSensitive() *bool {
	return s.CaseSensitive
}

func (s *HiMarketHttpRouteMatchQueryParams) GetName() *string {
	return s.Name
}

func (s *HiMarketHttpRouteMatchQueryParams) GetType() *string {
	return s.Type
}

func (s *HiMarketHttpRouteMatchQueryParams) GetValue() *string {
	return s.Value
}

func (s *HiMarketHttpRouteMatchQueryParams) SetCaseSensitive(v bool) *HiMarketHttpRouteMatchQueryParams {
	s.CaseSensitive = &v
	return s
}

func (s *HiMarketHttpRouteMatchQueryParams) SetName(v string) *HiMarketHttpRouteMatchQueryParams {
	s.Name = &v
	return s
}

func (s *HiMarketHttpRouteMatchQueryParams) SetType(v string) *HiMarketHttpRouteMatchQueryParams {
	s.Type = &v
	return s
}

func (s *HiMarketHttpRouteMatchQueryParams) SetValue(v string) *HiMarketHttpRouteMatchQueryParams {
	s.Value = &v
	return s
}

func (s *HiMarketHttpRouteMatchQueryParams) Validate() error {
	return dara.Validate(s)
}
