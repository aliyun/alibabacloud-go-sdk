// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultimodalQueryContext interface {
	dara.Model
	String() string
	GoString() string
	SetOriginalQuery(v *MultimodalOriginalQuery) *MultimodalQueryContext
	GetOriginalQuery() *MultimodalOriginalQuery
}

type MultimodalQueryContext struct {
	OriginalQuery *MultimodalOriginalQuery `json:"originalQuery,omitempty" xml:"originalQuery,omitempty"`
}

func (s MultimodalQueryContext) String() string {
	return dara.Prettify(s)
}

func (s MultimodalQueryContext) GoString() string {
	return s.String()
}

func (s *MultimodalQueryContext) GetOriginalQuery() *MultimodalOriginalQuery {
	return s.OriginalQuery
}

func (s *MultimodalQueryContext) SetOriginalQuery(v *MultimodalOriginalQuery) *MultimodalQueryContext {
	s.OriginalQuery = v
	return s
}

func (s *MultimodalQueryContext) Validate() error {
	if s.OriginalQuery != nil {
		if err := s.OriginalQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}
