// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterventionDictionaryNerResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuery(v string) *ListInterventionDictionaryNerResultsRequest
	GetQuery() *string
}

type ListInterventionDictionaryNerResultsRequest struct {
	// Query keywords.
	//
	// This parameter is required.
	//
	// example:
	//
	// "hello world"
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s ListInterventionDictionaryNerResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInterventionDictionaryNerResultsRequest) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryNerResultsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListInterventionDictionaryNerResultsRequest) SetQuery(v string) *ListInterventionDictionaryNerResultsRequest {
	s.Query = &v
	return s
}

func (s *ListInterventionDictionaryNerResultsRequest) Validate() error {
	return dara.Validate(s)
}
