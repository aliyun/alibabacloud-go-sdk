// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerNoteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerUid(v string) *CustomerNoteListRequest
	GetCustomerUid() *string
	SetPageNum(v int32) *CustomerNoteListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *CustomerNoteListRequest
	GetPageSize() *int32
}

type CustomerNoteListRequest struct {
	// example:
	//
	// 5625862916391497
	CustomerUid *string `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s CustomerNoteListRequest) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteListRequest) GoString() string {
	return s.String()
}

func (s *CustomerNoteListRequest) GetCustomerUid() *string {
	return s.CustomerUid
}

func (s *CustomerNoteListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *CustomerNoteListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CustomerNoteListRequest) SetCustomerUid(v string) *CustomerNoteListRequest {
	s.CustomerUid = &v
	return s
}

func (s *CustomerNoteListRequest) SetPageNum(v int32) *CustomerNoteListRequest {
	s.PageNum = &v
	return s
}

func (s *CustomerNoteListRequest) SetPageSize(v int32) *CustomerNoteListRequest {
	s.PageSize = &v
	return s
}

func (s *CustomerNoteListRequest) Validate() error {
	return dara.Validate(s)
}
