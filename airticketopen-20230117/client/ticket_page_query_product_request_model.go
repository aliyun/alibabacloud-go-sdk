// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketPageQueryProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *TicketPageQueryProductRequest
	GetAccountNo() *int64
	SetPageNo(v int32) *TicketPageQueryProductRequest
	GetPageNo() *int32
	SetPageSize(v int32) *TicketPageQueryProductRequest
	GetPageSize() *int32
	SetScenicId(v int64) *TicketPageQueryProductRequest
	GetScenicId() *int64
}

type TicketPageQueryProductRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 123456
	ScenicId *int64 `json:"ScenicId,omitempty" xml:"ScenicId,omitempty"`
}

func (s TicketPageQueryProductRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductRequest) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *TicketPageQueryProductRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *TicketPageQueryProductRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *TicketPageQueryProductRequest) GetScenicId() *int64 {
	return s.ScenicId
}

func (s *TicketPageQueryProductRequest) SetAccountNo(v int64) *TicketPageQueryProductRequest {
	s.AccountNo = &v
	return s
}

func (s *TicketPageQueryProductRequest) SetPageNo(v int32) *TicketPageQueryProductRequest {
	s.PageNo = &v
	return s
}

func (s *TicketPageQueryProductRequest) SetPageSize(v int32) *TicketPageQueryProductRequest {
	s.PageSize = &v
	return s
}

func (s *TicketPageQueryProductRequest) SetScenicId(v int64) *TicketPageQueryProductRequest {
	s.ScenicId = &v
	return s
}

func (s *TicketPageQueryProductRequest) Validate() error {
	return dara.Validate(s)
}
