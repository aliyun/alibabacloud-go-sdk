// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContactWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCountTotalRow(v bool) *GetContactWhiteListRequest
	GetCountTotalRow() *bool
	SetInstanceId(v string) *GetContactWhiteListRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *GetContactWhiteListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetContactWhiteListRequest
	GetPageSize() *int32
}

type GetContactWhiteListRequest struct {
	// Indicates whether the total number of entries is required
	//
	// example:
	//
	// true
	CountTotalRow *bool `json:"CountTotalRow,omitempty" xml:"CountTotalRow,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 2a830781-324e-4568-ae96-309f93090fe1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of entries per page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetContactWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContactWhiteListRequest) GoString() string {
	return s.String()
}

func (s *GetContactWhiteListRequest) GetCountTotalRow() *bool {
	return s.CountTotalRow
}

func (s *GetContactWhiteListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetContactWhiteListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetContactWhiteListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetContactWhiteListRequest) SetCountTotalRow(v bool) *GetContactWhiteListRequest {
	s.CountTotalRow = &v
	return s
}

func (s *GetContactWhiteListRequest) SetInstanceId(v string) *GetContactWhiteListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetContactWhiteListRequest) SetPageNumber(v int32) *GetContactWhiteListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetContactWhiteListRequest) SetPageSize(v int32) *GetContactWhiteListRequest {
	s.PageSize = &v
	return s
}

func (s *GetContactWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
