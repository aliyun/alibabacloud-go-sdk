// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLogisticsSmsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpressCompanyCode(v string) *SendLogisticsSmsRequest
	GetExpressCompanyCode() *string
	SetMailNo(v string) *SendLogisticsSmsRequest
	GetMailNo() *string
	SetOutId(v string) *SendLogisticsSmsRequest
	GetOutId() *string
	SetOwnerId(v int64) *SendLogisticsSmsRequest
	GetOwnerId() *int64
	SetPlatformCompanyCode(v string) *SendLogisticsSmsRequest
	GetPlatformCompanyCode() *string
	SetResourceOwnerAccount(v string) *SendLogisticsSmsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SendLogisticsSmsRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *SendLogisticsSmsRequest
	GetSignName() *string
	SetTemplateCode(v string) *SendLogisticsSmsRequest
	GetTemplateCode() *string
	SetTemplateParam(v string) *SendLogisticsSmsRequest
	GetTemplateParam() *string
}

type SendLogisticsSmsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	ExpressCompanyCode *string `json:"ExpressCompanyCode,omitempty" xml:"ExpressCompanyCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	MailNo *string `json:"MailNo,omitempty" xml:"MailNo,omitempty"`
	// example:
	//
	// 示例值
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 示例值
	PlatformCompanyCode  *string `json:"PlatformCompanyCode,omitempty" xml:"PlatformCompanyCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	TemplateParam *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
}

func (s SendLogisticsSmsRequest) String() string {
	return dara.Prettify(s)
}

func (s SendLogisticsSmsRequest) GoString() string {
	return s.String()
}

func (s *SendLogisticsSmsRequest) GetExpressCompanyCode() *string {
	return s.ExpressCompanyCode
}

func (s *SendLogisticsSmsRequest) GetMailNo() *string {
	return s.MailNo
}

func (s *SendLogisticsSmsRequest) GetOutId() *string {
	return s.OutId
}

func (s *SendLogisticsSmsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SendLogisticsSmsRequest) GetPlatformCompanyCode() *string {
	return s.PlatformCompanyCode
}

func (s *SendLogisticsSmsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SendLogisticsSmsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SendLogisticsSmsRequest) GetSignName() *string {
	return s.SignName
}

func (s *SendLogisticsSmsRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *SendLogisticsSmsRequest) GetTemplateParam() *string {
	return s.TemplateParam
}

func (s *SendLogisticsSmsRequest) SetExpressCompanyCode(v string) *SendLogisticsSmsRequest {
	s.ExpressCompanyCode = &v
	return s
}

func (s *SendLogisticsSmsRequest) SetMailNo(v string) *SendLogisticsSmsRequest {
	s.MailNo = &v
	return s
}

func (s *SendLogisticsSmsRequest) SetOutId(v string) *SendLogisticsSmsRequest {
	s.OutId = &v
	return s
}

func (s *SendLogisticsSmsRequest) SetOwnerId(v int64) *SendLogisticsSmsRequest {
	s.OwnerId = &v
	return s
}

func (s *SendLogisticsSmsRequest) SetPlatformCompanyCode(v string) *SendLogisticsSmsRequest {
	s.PlatformCompanyCode = &v
	return s
}

func (s *SendLogisticsSmsRequest) SetResourceOwnerAccount(v string) *SendLogisticsSmsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendLogisticsSmsRequest) SetResourceOwnerId(v int64) *SendLogisticsSmsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendLogisticsSmsRequest) SetSignName(v string) *SendLogisticsSmsRequest {
	s.SignName = &v
	return s
}

func (s *SendLogisticsSmsRequest) SetTemplateCode(v string) *SendLogisticsSmsRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendLogisticsSmsRequest) SetTemplateParam(v string) *SendLogisticsSmsRequest {
	s.TemplateParam = &v
	return s
}

func (s *SendLogisticsSmsRequest) Validate() error {
	return dara.Validate(s)
}
