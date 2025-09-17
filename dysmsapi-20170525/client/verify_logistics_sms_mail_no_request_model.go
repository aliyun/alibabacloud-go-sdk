// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyLogisticsSmsMailNoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpressCompanyCode(v string) *VerifyLogisticsSmsMailNoRequest
	GetExpressCompanyCode() *string
	SetMailNo(v string) *VerifyLogisticsSmsMailNoRequest
	GetMailNo() *string
	SetOwnerId(v int64) *VerifyLogisticsSmsMailNoRequest
	GetOwnerId() *int64
	SetPlatformCompanyCode(v string) *VerifyLogisticsSmsMailNoRequest
	GetPlatformCompanyCode() *string
	SetResourceOwnerAccount(v string) *VerifyLogisticsSmsMailNoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *VerifyLogisticsSmsMailNoRequest
	GetResourceOwnerId() *int64
}

type VerifyLogisticsSmsMailNoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	ExpressCompanyCode *string `json:"ExpressCompanyCode,omitempty" xml:"ExpressCompanyCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	MailNo  *string `json:"MailNo,omitempty" xml:"MailNo,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 示例值示例值
	PlatformCompanyCode  *string `json:"PlatformCompanyCode,omitempty" xml:"PlatformCompanyCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s VerifyLogisticsSmsMailNoRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyLogisticsSmsMailNoRequest) GoString() string {
	return s.String()
}

func (s *VerifyLogisticsSmsMailNoRequest) GetExpressCompanyCode() *string {
	return s.ExpressCompanyCode
}

func (s *VerifyLogisticsSmsMailNoRequest) GetMailNo() *string {
	return s.MailNo
}

func (s *VerifyLogisticsSmsMailNoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *VerifyLogisticsSmsMailNoRequest) GetPlatformCompanyCode() *string {
	return s.PlatformCompanyCode
}

func (s *VerifyLogisticsSmsMailNoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *VerifyLogisticsSmsMailNoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *VerifyLogisticsSmsMailNoRequest) SetExpressCompanyCode(v string) *VerifyLogisticsSmsMailNoRequest {
	s.ExpressCompanyCode = &v
	return s
}

func (s *VerifyLogisticsSmsMailNoRequest) SetMailNo(v string) *VerifyLogisticsSmsMailNoRequest {
	s.MailNo = &v
	return s
}

func (s *VerifyLogisticsSmsMailNoRequest) SetOwnerId(v int64) *VerifyLogisticsSmsMailNoRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyLogisticsSmsMailNoRequest) SetPlatformCompanyCode(v string) *VerifyLogisticsSmsMailNoRequest {
	s.PlatformCompanyCode = &v
	return s
}

func (s *VerifyLogisticsSmsMailNoRequest) SetResourceOwnerAccount(v string) *VerifyLogisticsSmsMailNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VerifyLogisticsSmsMailNoRequest) SetResourceOwnerId(v int64) *VerifyLogisticsSmsMailNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *VerifyLogisticsSmsMailNoRequest) Validate() error {
	return dara.Validate(s)
}
