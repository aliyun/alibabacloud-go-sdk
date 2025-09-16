// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRestQueryResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndexName(v string) *ListRestQueryResultRequest
	GetIndexName() *string
	SetQuery(v map[string]interface{}) *ListRestQueryResultRequest
	GetQuery() map[string]interface{}
}

type ListRestQueryResultRequest struct {
	// The name of the index table.
	//
	// example:
	//
	// main_index
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	// The rest query statement.
	//
	// example:
	//
	// query%3Drelation_id%3A%221151274675_2%22%26%26cluster%3Dgeneral%26%26config%3Dstart%3A0%2Chit%3A10%2Cformat%3Ajson
	Query map[string]interface{} `json:"query,omitempty" xml:"query,omitempty"`
}

func (s ListRestQueryResultRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRestQueryResultRequest) GoString() string {
	return s.String()
}

func (s *ListRestQueryResultRequest) GetIndexName() *string {
	return s.IndexName
}

func (s *ListRestQueryResultRequest) GetQuery() map[string]interface{} {
	return s.Query
}

func (s *ListRestQueryResultRequest) SetIndexName(v string) *ListRestQueryResultRequest {
	s.IndexName = &v
	return s
}

func (s *ListRestQueryResultRequest) SetQuery(v map[string]interface{}) *ListRestQueryResultRequest {
	s.Query = v
	return s
}

func (s *ListRestQueryResultRequest) Validate() error {
	return dara.Validate(s)
}
