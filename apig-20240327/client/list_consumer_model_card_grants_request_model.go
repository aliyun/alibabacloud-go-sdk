// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerModelCardGrantsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelCardId(v string) *ListConsumerModelCardGrantsRequest
	GetModelCardId() *string
	SetPageNumber(v int32) *ListConsumerModelCardGrantsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListConsumerModelCardGrantsRequest
	GetPageSize() *int32
	SetPrincipalId(v string) *ListConsumerModelCardGrantsRequest
	GetPrincipalId() *string
	SetPrincipalType(v string) *ListConsumerModelCardGrantsRequest
	GetPrincipalType() *string
}

type ListConsumerModelCardGrantsRequest struct {
	// Filters results by model card ID.
	//
	// example:
	//
	// msc-d00000000000000000000
	ModelCardId *string `json:"modelCardId,omitempty" xml:"modelCardId,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Filters results by principal ID. This parameter must be specified together with principalType.
	//
	// example:
	//
	// cs-d00000000000000000000
	PrincipalId *string `json:"principalId,omitempty" xml:"principalId,omitempty"`
	// Filters results by principal type. Valid values: Consumer and ConsumerGroup. This parameter must be specified together with principalId.
	//
	// example:
	//
	// Consumer
	PrincipalType *string `json:"principalType,omitempty" xml:"principalType,omitempty"`
}

func (s ListConsumerModelCardGrantsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerModelCardGrantsRequest) GoString() string {
	return s.String()
}

func (s *ListConsumerModelCardGrantsRequest) GetModelCardId() *string {
	return s.ModelCardId
}

func (s *ListConsumerModelCardGrantsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConsumerModelCardGrantsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConsumerModelCardGrantsRequest) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListConsumerModelCardGrantsRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListConsumerModelCardGrantsRequest) SetModelCardId(v string) *ListConsumerModelCardGrantsRequest {
	s.ModelCardId = &v
	return s
}

func (s *ListConsumerModelCardGrantsRequest) SetPageNumber(v int32) *ListConsumerModelCardGrantsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerModelCardGrantsRequest) SetPageSize(v int32) *ListConsumerModelCardGrantsRequest {
	s.PageSize = &v
	return s
}

func (s *ListConsumerModelCardGrantsRequest) SetPrincipalId(v string) *ListConsumerModelCardGrantsRequest {
	s.PrincipalId = &v
	return s
}

func (s *ListConsumerModelCardGrantsRequest) SetPrincipalType(v string) *ListConsumerModelCardGrantsRequest {
	s.PrincipalType = &v
	return s
}

func (s *ListConsumerModelCardGrantsRequest) Validate() error {
	return dara.Validate(s)
}
