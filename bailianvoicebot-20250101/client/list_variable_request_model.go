// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *ListVariableRequest
	GetBusinessUnitId() *string
	SetPageNumber(v int32) *ListVariableRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVariableRequest
	GetPageSize() *int32
	SetSearchPattern(v string) *ListVariableRequest
	GetSearchPattern() *string
}

type ListVariableRequest struct {
	// example:
	//
	// llm-zop7ukgtksltamo4
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// age
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
}

func (s ListVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVariableRequest) GoString() string {
	return s.String()
}

func (s *ListVariableRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *ListVariableRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVariableRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVariableRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListVariableRequest) SetBusinessUnitId(v string) *ListVariableRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *ListVariableRequest) SetPageNumber(v int32) *ListVariableRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVariableRequest) SetPageSize(v int32) *ListVariableRequest {
	s.PageSize = &v
	return s
}

func (s *ListVariableRequest) SetSearchPattern(v string) *ListVariableRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListVariableRequest) Validate() error {
	return dara.Validate(s)
}
