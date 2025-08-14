// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditPageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAuditPageListRequest
	GetLang() *string
	SetAuditStatus(v string) *DescribeAuditPageListRequest
	GetAuditStatus() *string
	SetCurrentPage(v string) *DescribeAuditPageListRequest
	GetCurrentPage() *string
	SetEventCode(v string) *DescribeAuditPageListRequest
	GetEventCode() *string
	SetPageSize(v string) *DescribeAuditPageListRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeAuditPageListRequest
	GetRegId() *string
	SetRuleName(v string) *DescribeAuditPageListRequest
	GetRuleName() *string
}

type DescribeAuditPageListRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Audit status
	//
	// example:
	//
	// SUCCESS
	AuditStatus *string `json:"auditStatus,omitempty" xml:"auditStatus,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Event code
	//
	// example:
	//
	// de_awukck7117
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Page size, default value is 10
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Policy name
	//
	// example:
	//
	// 手机号MD5命中人脸测试名单
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s DescribeAuditPageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditPageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditPageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAuditPageListRequest) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *DescribeAuditPageListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeAuditPageListRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeAuditPageListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAuditPageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAuditPageListRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeAuditPageListRequest) SetLang(v string) *DescribeAuditPageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAuditPageListRequest) SetAuditStatus(v string) *DescribeAuditPageListRequest {
	s.AuditStatus = &v
	return s
}

func (s *DescribeAuditPageListRequest) SetCurrentPage(v string) *DescribeAuditPageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAuditPageListRequest) SetEventCode(v string) *DescribeAuditPageListRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeAuditPageListRequest) SetPageSize(v string) *DescribeAuditPageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditPageListRequest) SetRegId(v string) *DescribeAuditPageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAuditPageListRequest) SetRuleName(v string) *DescribeAuditPageListRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeAuditPageListRequest) Validate() error {
	return dara.Validate(s)
}
