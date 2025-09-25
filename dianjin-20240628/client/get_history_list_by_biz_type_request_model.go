// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoryListByBizTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetHistoryListByBizTypeRequest
	GetBizId() *string
	SetBizType(v string) *GetHistoryListByBizTypeRequest
	GetBizType() *string
	SetPage(v int32) *GetHistoryListByBizTypeRequest
	GetPage() *int32
	SetPageSize(v int32) *GetHistoryListByBizTypeRequest
	GetPageSize() *int32
}

type GetHistoryListByBizTypeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// GysYBsxx
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// LibraryChat
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetHistoryListByBizTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHistoryListByBizTypeRequest) GoString() string {
	return s.String()
}

func (s *GetHistoryListByBizTypeRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetHistoryListByBizTypeRequest) GetBizType() *string {
	return s.BizType
}

func (s *GetHistoryListByBizTypeRequest) GetPage() *int32 {
	return s.Page
}

func (s *GetHistoryListByBizTypeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetHistoryListByBizTypeRequest) SetBizId(v string) *GetHistoryListByBizTypeRequest {
	s.BizId = &v
	return s
}

func (s *GetHistoryListByBizTypeRequest) SetBizType(v string) *GetHistoryListByBizTypeRequest {
	s.BizType = &v
	return s
}

func (s *GetHistoryListByBizTypeRequest) SetPage(v int32) *GetHistoryListByBizTypeRequest {
	s.Page = &v
	return s
}

func (s *GetHistoryListByBizTypeRequest) SetPageSize(v int32) *GetHistoryListByBizTypeRequest {
	s.PageSize = &v
	return s
}

func (s *GetHistoryListByBizTypeRequest) Validate() error {
	return dara.Validate(s)
}
