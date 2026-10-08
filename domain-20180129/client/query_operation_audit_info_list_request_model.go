// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOperationAuditInfoListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditStatus(v int32) *QueryOperationAuditInfoListRequest
	GetAuditStatus() *int32
	SetAuditType(v int32) *QueryOperationAuditInfoListRequest
	GetAuditType() *int32
	SetDomainName(v string) *QueryOperationAuditInfoListRequest
	GetDomainName() *string
	SetLang(v string) *QueryOperationAuditInfoListRequest
	GetLang() *string
	SetPageNum(v int32) *QueryOperationAuditInfoListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryOperationAuditInfoListRequest
	GetPageSize() *int32
}

type QueryOperationAuditInfoListRequest struct {
	// Review status. Valid values:
	//
	// - **0**: Information pending completion.
	//
	// - **1**, **2**, **3**, **4**: Under review.
	//
	// - **5**: Review failed.
	//
	// - **6**: Review succeeded.
	//
	// - **7**: Review canceled.
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// Review type. Valid value:
	//
	// **1**: Offline domain name transfer.
	//
	// example:
	//
	// 1
	AuditType *int32 `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	// Domain name to query.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Language of error messages returned by the API. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of records per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryOperationAuditInfoListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryOperationAuditInfoListRequest) GoString() string {
	return s.String()
}

func (s *QueryOperationAuditInfoListRequest) GetAuditStatus() *int32 {
	return s.AuditStatus
}

func (s *QueryOperationAuditInfoListRequest) GetAuditType() *int32 {
	return s.AuditType
}

func (s *QueryOperationAuditInfoListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryOperationAuditInfoListRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryOperationAuditInfoListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryOperationAuditInfoListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryOperationAuditInfoListRequest) SetAuditStatus(v int32) *QueryOperationAuditInfoListRequest {
	s.AuditStatus = &v
	return s
}

func (s *QueryOperationAuditInfoListRequest) SetAuditType(v int32) *QueryOperationAuditInfoListRequest {
	s.AuditType = &v
	return s
}

func (s *QueryOperationAuditInfoListRequest) SetDomainName(v string) *QueryOperationAuditInfoListRequest {
	s.DomainName = &v
	return s
}

func (s *QueryOperationAuditInfoListRequest) SetLang(v string) *QueryOperationAuditInfoListRequest {
	s.Lang = &v
	return s
}

func (s *QueryOperationAuditInfoListRequest) SetPageNum(v int32) *QueryOperationAuditInfoListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryOperationAuditInfoListRequest) SetPageSize(v int32) *QueryOperationAuditInfoListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryOperationAuditInfoListRequest) Validate() error {
	return dara.Validate(s)
}
