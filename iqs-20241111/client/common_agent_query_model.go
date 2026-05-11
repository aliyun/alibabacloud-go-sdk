// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonAgentQuery interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *CommonAgentQuery
	GetLimit() *int32
	SetQuery(v string) *CommonAgentQuery
	GetQuery() *string
	SetQuerySceneEnumCode(v string) *CommonAgentQuery
	GetQuerySceneEnumCode() *string
	SetSearchModel(v string) *CommonAgentQuery
	GetSearchModel() *string
}

type CommonAgentQuery struct {
	Limit              *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	Query              *string `json:"query,omitempty" xml:"query,omitempty"`
	QuerySceneEnumCode *string `json:"querySceneEnumCode,omitempty" xml:"querySceneEnumCode,omitempty"`
	SearchModel        *string `json:"searchModel,omitempty" xml:"searchModel,omitempty"`
}

func (s CommonAgentQuery) String() string {
	return dara.Prettify(s)
}

func (s CommonAgentQuery) GoString() string {
	return s.String()
}

func (s *CommonAgentQuery) GetLimit() *int32 {
	return s.Limit
}

func (s *CommonAgentQuery) GetQuery() *string {
	return s.Query
}

func (s *CommonAgentQuery) GetQuerySceneEnumCode() *string {
	return s.QuerySceneEnumCode
}

func (s *CommonAgentQuery) GetSearchModel() *string {
	return s.SearchModel
}

func (s *CommonAgentQuery) SetLimit(v int32) *CommonAgentQuery {
	s.Limit = &v
	return s
}

func (s *CommonAgentQuery) SetQuery(v string) *CommonAgentQuery {
	s.Query = &v
	return s
}

func (s *CommonAgentQuery) SetQuerySceneEnumCode(v string) *CommonAgentQuery {
	s.QuerySceneEnumCode = &v
	return s
}

func (s *CommonAgentQuery) SetSearchModel(v string) *CommonAgentQuery {
	s.SearchModel = &v
	return s
}

func (s *CommonAgentQuery) Validate() error {
	return dara.Validate(s)
}
