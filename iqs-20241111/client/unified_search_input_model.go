// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnifiedSearchInput interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedParams(v map[string]interface{}) *UnifiedSearchInput
	GetAdvancedParams() map[string]interface{}
	SetCategory(v string) *UnifiedSearchInput
	GetCategory() *string
	SetContents(v *RequestContents) *UnifiedSearchInput
	GetContents() *RequestContents
	SetEngineType(v string) *UnifiedSearchInput
	GetEngineType() *string
	SetLocation(v string) *UnifiedSearchInput
	GetLocation() *string
	SetQuery(v string) *UnifiedSearchInput
	GetQuery() *string
	SetTimeRange(v string) *UnifiedSearchInput
	GetTimeRange() *string
}

type UnifiedSearchInput struct {
	AdvancedParams map[string]interface{} `json:"advancedParams,omitempty" xml:"advancedParams,omitempty"`
	Category       *string                `json:"category,omitempty" xml:"category,omitempty"`
	Contents       *RequestContents       `json:"contents,omitempty" xml:"contents,omitempty"`
	EngineType     *string                `json:"engineType,omitempty" xml:"engineType,omitempty"`
	Location       *string                `json:"location,omitempty" xml:"location,omitempty"`
	Query          *string                `json:"query,omitempty" xml:"query,omitempty"`
	TimeRange      *string                `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s UnifiedSearchInput) String() string {
	return dara.Prettify(s)
}

func (s UnifiedSearchInput) GoString() string {
	return s.String()
}

func (s *UnifiedSearchInput) GetAdvancedParams() map[string]interface{} {
	return s.AdvancedParams
}

func (s *UnifiedSearchInput) GetCategory() *string {
	return s.Category
}

func (s *UnifiedSearchInput) GetContents() *RequestContents {
	return s.Contents
}

func (s *UnifiedSearchInput) GetEngineType() *string {
	return s.EngineType
}

func (s *UnifiedSearchInput) GetLocation() *string {
	return s.Location
}

func (s *UnifiedSearchInput) GetQuery() *string {
	return s.Query
}

func (s *UnifiedSearchInput) GetTimeRange() *string {
	return s.TimeRange
}

func (s *UnifiedSearchInput) SetAdvancedParams(v map[string]interface{}) *UnifiedSearchInput {
	s.AdvancedParams = v
	return s
}

func (s *UnifiedSearchInput) SetCategory(v string) *UnifiedSearchInput {
	s.Category = &v
	return s
}

func (s *UnifiedSearchInput) SetContents(v *RequestContents) *UnifiedSearchInput {
	s.Contents = v
	return s
}

func (s *UnifiedSearchInput) SetEngineType(v string) *UnifiedSearchInput {
	s.EngineType = &v
	return s
}

func (s *UnifiedSearchInput) SetLocation(v string) *UnifiedSearchInput {
	s.Location = &v
	return s
}

func (s *UnifiedSearchInput) SetQuery(v string) *UnifiedSearchInput {
	s.Query = &v
	return s
}

func (s *UnifiedSearchInput) SetTimeRange(v string) *UnifiedSearchInput {
	s.TimeRange = &v
	return s
}

func (s *UnifiedSearchInput) Validate() error {
	if s.Contents != nil {
		if err := s.Contents.Validate(); err != nil {
			return err
		}
	}
	return nil
}
