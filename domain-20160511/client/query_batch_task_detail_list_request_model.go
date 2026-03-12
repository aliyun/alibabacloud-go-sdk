// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBatchTaskDetailListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *QueryBatchTaskDetailListRequest
	GetDomainName() *string
	SetLang(v string) *QueryBatchTaskDetailListRequest
	GetLang() *string
	SetPageNum(v int32) *QueryBatchTaskDetailListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryBatchTaskDetailListRequest
	GetPageSize() *int32
	SetSaleId(v string) *QueryBatchTaskDetailListRequest
	GetSaleId() *string
	SetTaskNo(v string) *QueryBatchTaskDetailListRequest
	GetTaskNo() *string
	SetTaskStatus(v int32) *QueryBatchTaskDetailListRequest
	GetTaskStatus() *int32
	SetUserClientIp(v string) *QueryBatchTaskDetailListRequest
	GetUserClientIp() *string
}

type QueryBatchTaskDetailListRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SaleId   *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	// This parameter is required.
	TaskNo       *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskStatus   *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryBatchTaskDetailListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryBatchTaskDetailListRequest) GoString() string {
	return s.String()
}

func (s *QueryBatchTaskDetailListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryBatchTaskDetailListRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryBatchTaskDetailListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryBatchTaskDetailListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryBatchTaskDetailListRequest) GetSaleId() *string {
	return s.SaleId
}

func (s *QueryBatchTaskDetailListRequest) GetTaskNo() *string {
	return s.TaskNo
}

func (s *QueryBatchTaskDetailListRequest) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *QueryBatchTaskDetailListRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryBatchTaskDetailListRequest) SetDomainName(v string) *QueryBatchTaskDetailListRequest {
	s.DomainName = &v
	return s
}

func (s *QueryBatchTaskDetailListRequest) SetLang(v string) *QueryBatchTaskDetailListRequest {
	s.Lang = &v
	return s
}

func (s *QueryBatchTaskDetailListRequest) SetPageNum(v int32) *QueryBatchTaskDetailListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryBatchTaskDetailListRequest) SetPageSize(v int32) *QueryBatchTaskDetailListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryBatchTaskDetailListRequest) SetSaleId(v string) *QueryBatchTaskDetailListRequest {
	s.SaleId = &v
	return s
}

func (s *QueryBatchTaskDetailListRequest) SetTaskNo(v string) *QueryBatchTaskDetailListRequest {
	s.TaskNo = &v
	return s
}

func (s *QueryBatchTaskDetailListRequest) SetTaskStatus(v int32) *QueryBatchTaskDetailListRequest {
	s.TaskStatus = &v
	return s
}

func (s *QueryBatchTaskDetailListRequest) SetUserClientIp(v string) *QueryBatchTaskDetailListRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryBatchTaskDetailListRequest) Validate() error {
	return dara.Validate(s)
}
