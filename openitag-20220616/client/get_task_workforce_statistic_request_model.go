// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskWorkforceStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *GetTaskWorkforceStatisticRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetTaskWorkforceStatisticRequest
	GetPageSize() *int32
	SetStatType(v string) *GetTaskWorkforceStatisticRequest
	GetStatType() *string
}

type GetTaskWorkforceStatisticRequest struct {
	// The page number of the member list. The value starts from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of members per page in a paged query. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The statistics type. Valid values:
	//
	// - ITEM: Statistics are collected based on individual review records.
	//
	// - OPERATORCELL: Statistics are collected based on operation units. A single ITEM may contain multiple operation units.
	//
	// example:
	//
	// ITEM
	StatType *string `json:"StatType,omitempty" xml:"StatType,omitempty"`
}

func (s GetTaskWorkforceStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskWorkforceStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetTaskWorkforceStatisticRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetTaskWorkforceStatisticRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetTaskWorkforceStatisticRequest) GetStatType() *string {
	return s.StatType
}

func (s *GetTaskWorkforceStatisticRequest) SetPageNumber(v int32) *GetTaskWorkforceStatisticRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTaskWorkforceStatisticRequest) SetPageSize(v int32) *GetTaskWorkforceStatisticRequest {
	s.PageSize = &v
	return s
}

func (s *GetTaskWorkforceStatisticRequest) SetStatType(v string) *GetTaskWorkforceStatisticRequest {
	s.StatType = &v
	return s
}

func (s *GetTaskWorkforceStatisticRequest) Validate() error {
	return dara.Validate(s)
}
