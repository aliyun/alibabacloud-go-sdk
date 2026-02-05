// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultimodalSearchBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedParams(v map[string]interface{}) *MultimodalSearchBody
	GetAdvancedParams() map[string]interface{}
	SetQuery(v string) *MultimodalSearchBody
	GetQuery() *string
}

type MultimodalSearchBody struct {
	AdvancedParams map[string]interface{} `json:"advancedParams,omitempty" xml:"advancedParams,omitempty"`
	Query          *string                `json:"query,omitempty" xml:"query,omitempty"`
}

func (s MultimodalSearchBody) String() string {
	return dara.Prettify(s)
}

func (s MultimodalSearchBody) GoString() string {
	return s.String()
}

func (s *MultimodalSearchBody) GetAdvancedParams() map[string]interface{} {
	return s.AdvancedParams
}

func (s *MultimodalSearchBody) GetQuery() *string {
	return s.Query
}

func (s *MultimodalSearchBody) SetAdvancedParams(v map[string]interface{}) *MultimodalSearchBody {
	s.AdvancedParams = v
	return s
}

func (s *MultimodalSearchBody) SetQuery(v string) *MultimodalSearchBody {
	s.Query = &v
	return s
}

func (s *MultimodalSearchBody) Validate() error {
	return dara.Validate(s)
}
