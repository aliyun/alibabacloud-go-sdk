// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnifiedQueryContext interface {
	dara.Model
	String() string
	GoString() string
	SetEngineType(v string) *UnifiedQueryContext
	GetEngineType() *string
	SetOriginalQuery(v *UnifiedOriginalQuery) *UnifiedQueryContext
	GetOriginalQuery() *UnifiedOriginalQuery
	SetRewrite(v *UnifiedRewrite) *UnifiedQueryContext
	GetRewrite() *UnifiedRewrite
}

type UnifiedQueryContext struct {
	// example:
	//
	// Generic
	EngineType    *string               `json:"engineType,omitempty" xml:"engineType,omitempty"`
	OriginalQuery *UnifiedOriginalQuery `json:"originalQuery,omitempty" xml:"originalQuery,omitempty"`
	Rewrite       *UnifiedRewrite       `json:"rewrite,omitempty" xml:"rewrite,omitempty"`
}

func (s UnifiedQueryContext) String() string {
	return dara.Prettify(s)
}

func (s UnifiedQueryContext) GoString() string {
	return s.String()
}

func (s *UnifiedQueryContext) GetEngineType() *string {
	return s.EngineType
}

func (s *UnifiedQueryContext) GetOriginalQuery() *UnifiedOriginalQuery {
	return s.OriginalQuery
}

func (s *UnifiedQueryContext) GetRewrite() *UnifiedRewrite {
	return s.Rewrite
}

func (s *UnifiedQueryContext) SetEngineType(v string) *UnifiedQueryContext {
	s.EngineType = &v
	return s
}

func (s *UnifiedQueryContext) SetOriginalQuery(v *UnifiedOriginalQuery) *UnifiedQueryContext {
	s.OriginalQuery = v
	return s
}

func (s *UnifiedQueryContext) SetRewrite(v *UnifiedRewrite) *UnifiedQueryContext {
	s.Rewrite = v
	return s
}

func (s *UnifiedQueryContext) Validate() error {
	if s.OriginalQuery != nil {
		if err := s.OriginalQuery.Validate(); err != nil {
			return err
		}
	}
	if s.Rewrite != nil {
		if err := s.Rewrite.Validate(); err != nil {
			return err
		}
	}
	return nil
}
