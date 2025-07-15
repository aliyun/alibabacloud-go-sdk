// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoNotCallNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListDoNotCallNumbersRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListDoNotCallNumbersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDoNotCallNumbersRequest
	GetPageSize() *int32
	SetScope(v string) *ListDoNotCallNumbersRequest
	GetScope() *string
	SetSearchPattern(v string) *ListDoNotCallNumbersRequest
	GetSearchPattern() *string
}

type ListDoNotCallNumbersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// INSTANCE
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// RemarkA
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
}

func (s ListDoNotCallNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoNotCallNumbersRequest) GoString() string {
	return s.String()
}

func (s *ListDoNotCallNumbersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDoNotCallNumbersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDoNotCallNumbersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDoNotCallNumbersRequest) GetScope() *string {
	return s.Scope
}

func (s *ListDoNotCallNumbersRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListDoNotCallNumbersRequest) SetInstanceId(v string) *ListDoNotCallNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDoNotCallNumbersRequest) SetPageNumber(v int32) *ListDoNotCallNumbersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDoNotCallNumbersRequest) SetPageSize(v int32) *ListDoNotCallNumbersRequest {
	s.PageSize = &v
	return s
}

func (s *ListDoNotCallNumbersRequest) SetScope(v string) *ListDoNotCallNumbersRequest {
	s.Scope = &v
	return s
}

func (s *ListDoNotCallNumbersRequest) SetSearchPattern(v string) *ListDoNotCallNumbersRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListDoNotCallNumbersRequest) Validate() error {
	return dara.Validate(s)
}
