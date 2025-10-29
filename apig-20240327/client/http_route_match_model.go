// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpRouteMatch interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v []*HttpRouteMatchHeaders) *HttpRouteMatch
	GetHeaders() []*HttpRouteMatchHeaders
	SetIgnoreUriCase(v bool) *HttpRouteMatch
	GetIgnoreUriCase() *bool
	SetMethods(v []*string) *HttpRouteMatch
	GetMethods() []*string
	SetPath(v *HttpRouteMatchPath) *HttpRouteMatch
	GetPath() *HttpRouteMatchPath
	SetQueryParams(v []*HttpRouteMatchQueryParams) *HttpRouteMatch
	GetQueryParams() []*HttpRouteMatchQueryParams
}

type HttpRouteMatch struct {
	Headers []*HttpRouteMatchHeaders `json:"headers,omitempty" xml:"headers,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IgnoreUriCase *bool                        `json:"ignoreUriCase,omitempty" xml:"ignoreUriCase,omitempty"`
	Methods       []*string                    `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
	Path          *HttpRouteMatchPath          `json:"path,omitempty" xml:"path,omitempty" type:"Struct"`
	QueryParams   []*HttpRouteMatchQueryParams `json:"queryParams,omitempty" xml:"queryParams,omitempty" type:"Repeated"`
}

func (s HttpRouteMatch) String() string {
	return dara.Prettify(s)
}

func (s HttpRouteMatch) GoString() string {
	return s.String()
}

func (s *HttpRouteMatch) GetHeaders() []*HttpRouteMatchHeaders {
	return s.Headers
}

func (s *HttpRouteMatch) GetIgnoreUriCase() *bool {
	return s.IgnoreUriCase
}

func (s *HttpRouteMatch) GetMethods() []*string {
	return s.Methods
}

func (s *HttpRouteMatch) GetPath() *HttpRouteMatchPath {
	return s.Path
}

func (s *HttpRouteMatch) GetQueryParams() []*HttpRouteMatchQueryParams {
	return s.QueryParams
}

func (s *HttpRouteMatch) SetHeaders(v []*HttpRouteMatchHeaders) *HttpRouteMatch {
	s.Headers = v
	return s
}

func (s *HttpRouteMatch) SetIgnoreUriCase(v bool) *HttpRouteMatch {
	s.IgnoreUriCase = &v
	return s
}

func (s *HttpRouteMatch) SetMethods(v []*string) *HttpRouteMatch {
	s.Methods = v
	return s
}

func (s *HttpRouteMatch) SetPath(v *HttpRouteMatchPath) *HttpRouteMatch {
	s.Path = v
	return s
}

func (s *HttpRouteMatch) SetQueryParams(v []*HttpRouteMatchQueryParams) *HttpRouteMatch {
	s.QueryParams = v
	return s
}

func (s *HttpRouteMatch) Validate() error {
	if s.Headers != nil {
		for _, item := range s.Headers {
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

type HttpRouteMatchHeaders struct {
	// example:
	//
	// dev
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Exact
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// true
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HttpRouteMatchHeaders) String() string {
	return dara.Prettify(s)
}

func (s HttpRouteMatchHeaders) GoString() string {
	return s.String()
}

func (s *HttpRouteMatchHeaders) GetName() *string {
	return s.Name
}

func (s *HttpRouteMatchHeaders) GetType() *string {
	return s.Type
}

func (s *HttpRouteMatchHeaders) GetValue() *string {
	return s.Value
}

func (s *HttpRouteMatchHeaders) SetName(v string) *HttpRouteMatchHeaders {
	s.Name = &v
	return s
}

func (s *HttpRouteMatchHeaders) SetType(v string) *HttpRouteMatchHeaders {
	s.Type = &v
	return s
}

func (s *HttpRouteMatchHeaders) SetValue(v string) *HttpRouteMatchHeaders {
	s.Value = &v
	return s
}

func (s *HttpRouteMatchHeaders) Validate() error {
	return dara.Validate(s)
}

type HttpRouteMatchPath struct {
	// example:
	//
	// Prefix
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// /user
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HttpRouteMatchPath) String() string {
	return dara.Prettify(s)
}

func (s HttpRouteMatchPath) GoString() string {
	return s.String()
}

func (s *HttpRouteMatchPath) GetType() *string {
	return s.Type
}

func (s *HttpRouteMatchPath) GetValue() *string {
	return s.Value
}

func (s *HttpRouteMatchPath) SetType(v string) *HttpRouteMatchPath {
	s.Type = &v
	return s
}

func (s *HttpRouteMatchPath) SetValue(v string) *HttpRouteMatchPath {
	s.Value = &v
	return s
}

func (s *HttpRouteMatchPath) Validate() error {
	return dara.Validate(s)
}

type HttpRouteMatchQueryParams struct {
	// example:
	//
	// age
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Exact
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 17
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HttpRouteMatchQueryParams) String() string {
	return dara.Prettify(s)
}

func (s HttpRouteMatchQueryParams) GoString() string {
	return s.String()
}

func (s *HttpRouteMatchQueryParams) GetName() *string {
	return s.Name
}

func (s *HttpRouteMatchQueryParams) GetType() *string {
	return s.Type
}

func (s *HttpRouteMatchQueryParams) GetValue() *string {
	return s.Value
}

func (s *HttpRouteMatchQueryParams) SetName(v string) *HttpRouteMatchQueryParams {
	s.Name = &v
	return s
}

func (s *HttpRouteMatchQueryParams) SetType(v string) *HttpRouteMatchQueryParams {
	s.Type = &v
	return s
}

func (s *HttpRouteMatchQueryParams) SetValue(v string) *HttpRouteMatchQueryParams {
	s.Value = &v
	return s
}

func (s *HttpRouteMatchQueryParams) Validate() error {
	return dara.Validate(s)
}
