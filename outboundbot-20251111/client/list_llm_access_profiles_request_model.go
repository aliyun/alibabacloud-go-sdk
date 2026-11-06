// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLlmAccessProfilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListLlmAccessProfilesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListLlmAccessProfilesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLlmAccessProfilesRequest
	GetPageSize() *int32
}

type ListLlmAccessProfilesRequest struct {
	// The instance ID.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListLlmAccessProfilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLlmAccessProfilesRequest) GoString() string {
	return s.String()
}

func (s *ListLlmAccessProfilesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListLlmAccessProfilesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLlmAccessProfilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLlmAccessProfilesRequest) SetInstanceId(v string) *ListLlmAccessProfilesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListLlmAccessProfilesRequest) SetPageNumber(v int32) *ListLlmAccessProfilesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLlmAccessProfilesRequest) SetPageSize(v int32) *ListLlmAccessProfilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLlmAccessProfilesRequest) Validate() error {
	return dara.Validate(s)
}
