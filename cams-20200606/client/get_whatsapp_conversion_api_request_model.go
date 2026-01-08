// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWhatsappConversionApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetWhatsappConversionApiRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *GetWhatsappConversionApiRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetWhatsappConversionApiRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetWhatsappConversionApiRequest
	GetResourceOwnerId() *int64
}

type GetWhatsappConversionApiRequest struct {
	// example:
	//
	// chatbot-cn-VBe6QXXX
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetWhatsappConversionApiRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappConversionApiRequest) GoString() string {
	return s.String()
}

func (s *GetWhatsappConversionApiRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWhatsappConversionApiRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetWhatsappConversionApiRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetWhatsappConversionApiRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetWhatsappConversionApiRequest) SetInstanceId(v string) *GetWhatsappConversionApiRequest {
	s.InstanceId = &v
	return s
}

func (s *GetWhatsappConversionApiRequest) SetOwnerId(v int64) *GetWhatsappConversionApiRequest {
	s.OwnerId = &v
	return s
}

func (s *GetWhatsappConversionApiRequest) SetResourceOwnerAccount(v string) *GetWhatsappConversionApiRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetWhatsappConversionApiRequest) SetResourceOwnerId(v int64) *GetWhatsappConversionApiRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetWhatsappConversionApiRequest) Validate() error {
	return dara.Validate(s)
}
