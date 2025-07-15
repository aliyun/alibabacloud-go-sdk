// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomCallTaggingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallTagNameList(v string) *ListCustomCallTaggingRequest
	GetCallTagNameList() *string
	SetInstanceId(v string) *ListCustomCallTaggingRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListCustomCallTaggingRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomCallTaggingRequest
	GetPageSize() *int32
	SetSearchPattern(v string) *ListCustomCallTaggingRequest
	GetSearchPattern() *string
}

type ListCustomCallTaggingRequest struct {
	// example:
	//
	// ["TagA"]
	CallTagNameList *string `json:"CallTagNameList,omitempty" xml:"CallTagNameList,omitempty"`
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
	// 1312121****
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
}

func (s ListCustomCallTaggingRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomCallTaggingRequest) GoString() string {
	return s.String()
}

func (s *ListCustomCallTaggingRequest) GetCallTagNameList() *string {
	return s.CallTagNameList
}

func (s *ListCustomCallTaggingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCustomCallTaggingRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomCallTaggingRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomCallTaggingRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListCustomCallTaggingRequest) SetCallTagNameList(v string) *ListCustomCallTaggingRequest {
	s.CallTagNameList = &v
	return s
}

func (s *ListCustomCallTaggingRequest) SetInstanceId(v string) *ListCustomCallTaggingRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCustomCallTaggingRequest) SetPageNumber(v int32) *ListCustomCallTaggingRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomCallTaggingRequest) SetPageSize(v int32) *ListCustomCallTaggingRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomCallTaggingRequest) SetSearchPattern(v string) *ListCustomCallTaggingRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListCustomCallTaggingRequest) Validate() error {
	return dara.Validate(s)
}
