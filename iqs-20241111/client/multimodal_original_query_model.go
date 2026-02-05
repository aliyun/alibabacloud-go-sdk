// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultimodalOriginalQuery interface {
	dara.Model
	String() string
	GoString() string
	SetQuery(v string) *MultimodalOriginalQuery
	GetQuery() *string
}

type MultimodalOriginalQuery struct {
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s MultimodalOriginalQuery) String() string {
	return dara.Prettify(s)
}

func (s MultimodalOriginalQuery) GoString() string {
	return s.String()
}

func (s *MultimodalOriginalQuery) GetQuery() *string {
	return s.Query
}

func (s *MultimodalOriginalQuery) SetQuery(v string) *MultimodalOriginalQuery {
	s.Query = &v
	return s
}

func (s *MultimodalOriginalQuery) Validate() error {
	return dara.Validate(s)
}
