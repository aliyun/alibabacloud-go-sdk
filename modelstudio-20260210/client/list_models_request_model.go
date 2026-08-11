// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapabilities(v []*string) *ListModelsRequest
	GetCapabilities() []*string
	SetContextWindow(v int64) *ListModelsRequest
	GetContextWindow() *int64
	SetFeatures(v []*string) *ListModelsRequest
	GetFeatures() []*string
	SetLanguage(v string) *ListModelsRequest
	GetLanguage() *string
	SetMaxResults(v int64) *ListModelsRequest
	GetMaxResults() *int64
	SetModel(v string) *ListModelsRequest
	GetModel() *string
	SetName(v string) *ListModelsRequest
	GetName() *string
	SetNextToken(v string) *ListModelsRequest
	GetNextToken() *string
	SetProviders(v []*string) *ListModelsRequest
	GetProviders() []*string
}

type ListModelsRequest struct {
	Capabilities []*string `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	ContextWindow *int64    `json:"contextWindow,omitempty" xml:"contextWindow,omitempty"`
	Features      []*string `json:"features,omitempty" xml:"features,omitempty" type:"Repeated"`
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
	NextToken *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Providers []*string `json:"providers,omitempty" xml:"providers,omitempty" type:"Repeated"`
}

func (s ListModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelsRequest) GoString() string {
	return s.String()
}

func (s *ListModelsRequest) GetCapabilities() []*string {
	return s.Capabilities
}

func (s *ListModelsRequest) GetContextWindow() *int64 {
	return s.ContextWindow
}

func (s *ListModelsRequest) GetFeatures() []*string {
	return s.Features
}

func (s *ListModelsRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListModelsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListModelsRequest) GetModel() *string {
	return s.Model
}

func (s *ListModelsRequest) GetName() *string {
	return s.Name
}

func (s *ListModelsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelsRequest) GetProviders() []*string {
	return s.Providers
}

func (s *ListModelsRequest) SetCapabilities(v []*string) *ListModelsRequest {
	s.Capabilities = v
	return s
}

func (s *ListModelsRequest) SetContextWindow(v int64) *ListModelsRequest {
	s.ContextWindow = &v
	return s
}

func (s *ListModelsRequest) SetFeatures(v []*string) *ListModelsRequest {
	s.Features = v
	return s
}

func (s *ListModelsRequest) SetLanguage(v string) *ListModelsRequest {
	s.Language = &v
	return s
}

func (s *ListModelsRequest) SetMaxResults(v int64) *ListModelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListModelsRequest) SetModel(v string) *ListModelsRequest {
	s.Model = &v
	return s
}

func (s *ListModelsRequest) SetName(v string) *ListModelsRequest {
	s.Name = &v
	return s
}

func (s *ListModelsRequest) SetNextToken(v string) *ListModelsRequest {
	s.NextToken = &v
	return s
}

func (s *ListModelsRequest) SetProviders(v []*string) *ListModelsRequest {
	s.Providers = v
	return s
}

func (s *ListModelsRequest) Validate() error {
	return dara.Validate(s)
}
