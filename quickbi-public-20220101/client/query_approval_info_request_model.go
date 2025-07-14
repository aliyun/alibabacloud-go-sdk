// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryApprovalInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int32) *QueryApprovalInfoRequest
	GetPage() *int32
	SetPageSize(v int32) *QueryApprovalInfoRequest
	GetPageSize() *int32
	SetStatus(v int32) *QueryApprovalInfoRequest
	GetStatus() *int32
	SetUserId(v string) *QueryApprovalInfoRequest
	GetUserId() *string
}

type QueryApprovalInfoRequest struct {
	// Page number, default is 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// Number of rows per page, default is 1000.
	//
	// example:
	//
	// 1000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Approval status:
	//
	// - 0: Pending
	//
	// - 1: Processed
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Current approver user ID, qbi user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12352fasdavsa
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryApprovalInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryApprovalInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryApprovalInfoRequest) GetPage() *int32 {
	return s.Page
}

func (s *QueryApprovalInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryApprovalInfoRequest) GetStatus() *int32 {
	return s.Status
}

func (s *QueryApprovalInfoRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryApprovalInfoRequest) SetPage(v int32) *QueryApprovalInfoRequest {
	s.Page = &v
	return s
}

func (s *QueryApprovalInfoRequest) SetPageSize(v int32) *QueryApprovalInfoRequest {
	s.PageSize = &v
	return s
}

func (s *QueryApprovalInfoRequest) SetStatus(v int32) *QueryApprovalInfoRequest {
	s.Status = &v
	return s
}

func (s *QueryApprovalInfoRequest) SetUserId(v string) *QueryApprovalInfoRequest {
	s.UserId = &v
	return s
}

func (s *QueryApprovalInfoRequest) Validate() error {
	return dara.Validate(s)
}
