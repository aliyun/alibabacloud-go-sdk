// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketPageQueryScenicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *TicketPageQueryScenicRequest
	GetAccountNo() *int64
	SetPageNo(v int32) *TicketPageQueryScenicRequest
	GetPageNo() *int32
	SetPageSize(v int32) *TicketPageQueryScenicRequest
	GetPageSize() *int32
}

type TicketPageQueryScenicRequest struct {
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
}

func (s TicketPageQueryScenicRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryScenicRequest) GoString() string {
	return s.String()
}

func (s *TicketPageQueryScenicRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *TicketPageQueryScenicRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *TicketPageQueryScenicRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *TicketPageQueryScenicRequest) SetAccountNo(v int64) *TicketPageQueryScenicRequest {
	s.AccountNo = &v
	return s
}

func (s *TicketPageQueryScenicRequest) SetPageNo(v int32) *TicketPageQueryScenicRequest {
	s.PageNo = &v
	return s
}

func (s *TicketPageQueryScenicRequest) SetPageSize(v int32) *TicketPageQueryScenicRequest {
	s.PageSize = &v
	return s
}

func (s *TicketPageQueryScenicRequest) Validate() error {
	return dara.Validate(s)
}
