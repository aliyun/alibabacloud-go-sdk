// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContactBlockListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCountTotalRow(v bool) *GetContactBlockListRequest
	GetCountTotalRow() *bool
	SetInstanceId(v string) *GetContactBlockListRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *GetContactBlockListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetContactBlockListRequest
	GetPageSize() *int32
}

type GetContactBlockListRequest struct {
	// example:
	//
	// 100
	CountTotalRow *bool `json:"CountTotalRow,omitempty" xml:"CountTotalRow,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 846e20ae-e113-4231-a792-cb354187c9f6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetContactBlockListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContactBlockListRequest) GoString() string {
	return s.String()
}

func (s *GetContactBlockListRequest) GetCountTotalRow() *bool {
	return s.CountTotalRow
}

func (s *GetContactBlockListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetContactBlockListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetContactBlockListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetContactBlockListRequest) SetCountTotalRow(v bool) *GetContactBlockListRequest {
	s.CountTotalRow = &v
	return s
}

func (s *GetContactBlockListRequest) SetInstanceId(v string) *GetContactBlockListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetContactBlockListRequest) SetPageNumber(v int32) *GetContactBlockListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetContactBlockListRequest) SetPageSize(v int32) *GetContactBlockListRequest {
	s.PageSize = &v
	return s
}

func (s *GetContactBlockListRequest) Validate() error {
	return dara.Validate(s)
}
