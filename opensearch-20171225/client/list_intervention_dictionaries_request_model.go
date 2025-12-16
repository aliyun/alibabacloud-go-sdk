// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterventionDictionariesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListInterventionDictionariesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInterventionDictionariesRequest
	GetPageSize() *int32
	SetTypes(v string) *ListInterventionDictionariesRequest
	GetTypes() *string
}

type ListInterventionDictionariesRequest struct {
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The type of the intervention dictionary. Valid values:
	//
	// 	- stopword: an intervention dictionary for stop word filtering
	//
	// 	- synonym: an intervention dictionary for synonym configuration
	//
	// 	- correction: an intervention dictionary for spelling correction
	//
	// 	- category_prediction: an intervention dictionary for category prediction
	//
	// 	- ner: an intervention dictionary for named entity recognition (NER)
	//
	// 	- term_weighting: an intervention dictionary for term weight analysis
	//
	// example:
	//
	// ["synonym"]
	Types *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s ListInterventionDictionariesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInterventionDictionariesRequest) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionariesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInterventionDictionariesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInterventionDictionariesRequest) GetTypes() *string {
	return s.Types
}

func (s *ListInterventionDictionariesRequest) SetPageNumber(v int32) *ListInterventionDictionariesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInterventionDictionariesRequest) SetPageSize(v int32) *ListInterventionDictionariesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInterventionDictionariesRequest) SetTypes(v string) *ListInterventionDictionariesRequest {
	s.Types = &v
	return s
}

func (s *ListInterventionDictionariesRequest) Validate() error {
	return dara.Validate(s)
}
