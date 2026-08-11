// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapabilitiesShrink(v string) *ListModelsShrinkRequest
	GetCapabilitiesShrink() *string
	SetContextWindow(v int64) *ListModelsShrinkRequest
	GetContextWindow() *int64
	SetFeaturesShrink(v string) *ListModelsShrinkRequest
	GetFeaturesShrink() *string
	SetLanguage(v string) *ListModelsShrinkRequest
	GetLanguage() *string
	SetMaxResults(v int64) *ListModelsShrinkRequest
	GetMaxResults() *int64
	SetModel(v string) *ListModelsShrinkRequest
	GetModel() *string
	SetName(v string) *ListModelsShrinkRequest
	GetName() *string
	SetNextToken(v string) *ListModelsShrinkRequest
	GetNextToken() *string
	SetProvidersShrink(v string) *ListModelsShrinkRequest
	GetProvidersShrink() *string
}

type ListModelsShrinkRequest struct {
	CapabilitiesShrink *string `json:"capabilities,omitempty" xml:"capabilities,omitempty"`
	// example:
	//
	// 10
	ContextWindow  *int64  `json:"contextWindow,omitempty" xml:"contextWindow,omitempty"`
	FeaturesShrink *string `json:"features,omitempty" xml:"features,omitempty"`
	// example:
	//
	// zh-CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// qwen-plus
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// example:
	//
	// qwen-plus
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// lwytFRtLdNk=
	NextToken       *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ProvidersShrink *string `json:"providers,omitempty" xml:"providers,omitempty"`
}

func (s ListModelsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListModelsShrinkRequest) GetCapabilitiesShrink() *string {
	return s.CapabilitiesShrink
}

func (s *ListModelsShrinkRequest) GetContextWindow() *int64 {
	return s.ContextWindow
}

func (s *ListModelsShrinkRequest) GetFeaturesShrink() *string {
	return s.FeaturesShrink
}

func (s *ListModelsShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListModelsShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListModelsShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *ListModelsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListModelsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelsShrinkRequest) GetProvidersShrink() *string {
	return s.ProvidersShrink
}

func (s *ListModelsShrinkRequest) SetCapabilitiesShrink(v string) *ListModelsShrinkRequest {
	s.CapabilitiesShrink = &v
	return s
}

func (s *ListModelsShrinkRequest) SetContextWindow(v int64) *ListModelsShrinkRequest {
	s.ContextWindow = &v
	return s
}

func (s *ListModelsShrinkRequest) SetFeaturesShrink(v string) *ListModelsShrinkRequest {
	s.FeaturesShrink = &v
	return s
}

func (s *ListModelsShrinkRequest) SetLanguage(v string) *ListModelsShrinkRequest {
	s.Language = &v
	return s
}

func (s *ListModelsShrinkRequest) SetMaxResults(v int64) *ListModelsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListModelsShrinkRequest) SetModel(v string) *ListModelsShrinkRequest {
	s.Model = &v
	return s
}

func (s *ListModelsShrinkRequest) SetName(v string) *ListModelsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListModelsShrinkRequest) SetNextToken(v string) *ListModelsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListModelsShrinkRequest) SetProvidersShrink(v string) *ListModelsShrinkRequest {
	s.ProvidersShrink = &v
	return s
}

func (s *ListModelsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
