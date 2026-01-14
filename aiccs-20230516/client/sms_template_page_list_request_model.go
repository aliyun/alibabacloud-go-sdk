// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmsTemplatePageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *SmsTemplatePageListRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *SmsTemplatePageListRequest
	GetPageNo() *int64
	SetPageSize(v int64) *SmsTemplatePageListRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *SmsTemplatePageListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SmsTemplatePageListRequest
	GetResourceOwnerId() *int64
	SetSign(v string) *SmsTemplatePageListRequest
	GetSign() *string
	SetSmsType(v int64) *SmsTemplatePageListRequest
	GetSmsType() *int64
	SetStatus(v int64) *SmsTemplatePageListRequest
	GetStatus() *int64
	SetTemplateId(v int64) *SmsTemplatePageListRequest
	GetTemplateId() *int64
	SetTemplateType(v int64) *SmsTemplatePageListRequest
	GetTemplateType() *int64
}

type SmsTemplatePageListRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 页码
	//
	// example:
	//
	// 24
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// 每页条数
	//
	// example:
	//
	// 97
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 短信签名
	//
	// example:
	//
	// 114ah23m
	Sign *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	// 短信类型
	//
	// example:
	//
	// 42
	SmsType *int64 `json:"SmsType,omitempty" xml:"SmsType,omitempty"`
	// 模板状态
	//
	// example:
	//
	// 92
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板ID
	//
	// example:
	//
	// 83
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 模板类型
	//
	// example:
	//
	// 19
	TemplateType *int64 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s SmsTemplatePageListRequest) String() string {
	return dara.Prettify(s)
}

func (s SmsTemplatePageListRequest) GoString() string {
	return s.String()
}

func (s *SmsTemplatePageListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SmsTemplatePageListRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *SmsTemplatePageListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *SmsTemplatePageListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SmsTemplatePageListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SmsTemplatePageListRequest) GetSign() *string {
	return s.Sign
}

func (s *SmsTemplatePageListRequest) GetSmsType() *int64 {
	return s.SmsType
}

func (s *SmsTemplatePageListRequest) GetStatus() *int64 {
	return s.Status
}

func (s *SmsTemplatePageListRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *SmsTemplatePageListRequest) GetTemplateType() *int64 {
	return s.TemplateType
}

func (s *SmsTemplatePageListRequest) SetOwnerId(v int64) *SmsTemplatePageListRequest {
	s.OwnerId = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetPageNo(v int64) *SmsTemplatePageListRequest {
	s.PageNo = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetPageSize(v int64) *SmsTemplatePageListRequest {
	s.PageSize = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetResourceOwnerAccount(v string) *SmsTemplatePageListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetResourceOwnerId(v int64) *SmsTemplatePageListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetSign(v string) *SmsTemplatePageListRequest {
	s.Sign = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetSmsType(v int64) *SmsTemplatePageListRequest {
	s.SmsType = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetStatus(v int64) *SmsTemplatePageListRequest {
	s.Status = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetTemplateId(v int64) *SmsTemplatePageListRequest {
	s.TemplateId = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetTemplateType(v int64) *SmsTemplatePageListRequest {
	s.TemplateType = &v
	return s
}

func (s *SmsTemplatePageListRequest) Validate() error {
	return dara.Validate(s)
}
