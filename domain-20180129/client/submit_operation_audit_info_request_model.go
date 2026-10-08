// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitOperationAuditInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditInfo(v string) *SubmitOperationAuditInfoRequest
	GetAuditInfo() *string
	SetAuditType(v int32) *SubmitOperationAuditInfoRequest
	GetAuditType() *int32
	SetDomainName(v string) *SubmitOperationAuditInfoRequest
	GetDomainName() *string
	SetId(v int64) *SubmitOperationAuditInfoRequest
	GetId() *int64
	SetLang(v string) *SubmitOperationAuditInfoRequest
	GetLang() *string
}

type SubmitOperationAuditInfoRequest struct {
	// The information to be reviewed. The displayed information varies by business type.
	//
	// example:
	//
	// 个人 {"regType":1,"registrantName":"张三","registrantNo":"2201919190**","telephone":"1390123****","account":"zhangsan@alimail.com","reason":1,"remark":"账号丢失"} 企业 {"regType":2,"registrantName":"华大信通","operatorName":"王武","operatorNo":"2201811987101901**",      "operatorPhone":"1390123****","account":"wangwu@alimail.com","companyNo":"91361100MA35N6****","reason":2,"remark":"账号丢失"}
	AuditInfo *string `json:"AuditInfo,omitempty" xml:"AuditInfo,omitempty"`
	// The business type. Valid values:
	//
	// **1**: Transfer a domain name offline, that is, transfer the domain name from the current Alibaba Cloud account to another Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuditType *int32 `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	// The domain name. You can specify one or more domain names, separated by commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx.com,yyyy.cn
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The review ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the error message returned by the API. Valid values:
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
}

func (s SubmitOperationAuditInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitOperationAuditInfoRequest) GoString() string {
	return s.String()
}

func (s *SubmitOperationAuditInfoRequest) GetAuditInfo() *string {
	return s.AuditInfo
}

func (s *SubmitOperationAuditInfoRequest) GetAuditType() *int32 {
	return s.AuditType
}

func (s *SubmitOperationAuditInfoRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SubmitOperationAuditInfoRequest) GetId() *int64 {
	return s.Id
}

func (s *SubmitOperationAuditInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *SubmitOperationAuditInfoRequest) SetAuditInfo(v string) *SubmitOperationAuditInfoRequest {
	s.AuditInfo = &v
	return s
}

func (s *SubmitOperationAuditInfoRequest) SetAuditType(v int32) *SubmitOperationAuditInfoRequest {
	s.AuditType = &v
	return s
}

func (s *SubmitOperationAuditInfoRequest) SetDomainName(v string) *SubmitOperationAuditInfoRequest {
	s.DomainName = &v
	return s
}

func (s *SubmitOperationAuditInfoRequest) SetId(v int64) *SubmitOperationAuditInfoRequest {
	s.Id = &v
	return s
}

func (s *SubmitOperationAuditInfoRequest) SetLang(v string) *SubmitOperationAuditInfoRequest {
	s.Lang = &v
	return s
}

func (s *SubmitOperationAuditInfoRequest) Validate() error {
	return dara.Validate(s)
}
