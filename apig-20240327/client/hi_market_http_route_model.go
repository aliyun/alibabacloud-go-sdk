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
	Builtin     *bool                   `json:"builtin,omitempty" xml:"builtin,omitempty"`
	Description *string                 `json:"description,omitempty" xml:"description,omitempty"`
	Domains     []*HiMarketDomain       `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
	Match       *HiMarketHttpRouteMatch `json:"match,omitempty" xml:"match,omitempty" type:"Struct"`
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
	Headers      []*HiMarketHttpRouteMatchHeaders      `json:"headers,omitempty" xml:"headers,omitempty" type:"Repeated"`
	Methods      []*string                             `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
	ModelMatches []*HiMarketHttpRouteMatchModelMatches `json:"modelMatches,omitempty" xml:"modelMatches,omitempty" type:"Repeated"`
	Path         *HiMarketHttpRouteMatchPath           `json:"path,omitempty" xml:"path,omitempty" type:"Struct"`
	QueryParams  []*HiMarketHttpRouteMatchQueryParams  `json:"queryParams,omitempty" xml:"queryParams,omitempty" type:"Repeated"`
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
	CaseSensitive *bool   `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
	Value         *string `json:"value,omitempty" xml:"value,omitempty"`
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
	CaseSensitive *bool   `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
	Value         *string `json:"value,omitempty" xml:"value,omitempty"`
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
	CaseSensitive *bool   `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
	Value         *string `json:"value,omitempty" xml:"value,omitempty"`
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
	CaseSensitive *bool   `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
	Value         *string `json:"value,omitempty" xml:"value,omitempty"`
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
