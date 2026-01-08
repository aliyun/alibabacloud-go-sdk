// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWhatsappConversionApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWhatsappConversionApiRequest
	GetCode() *string
	SetInstanceId(v string) *CreateWhatsappConversionApiRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *CreateWhatsappConversionApiRequest
	GetOwnerId() *int64
	SetPermissions(v []*string) *CreateWhatsappConversionApiRequest
	GetPermissions() []*string
	SetResourceOwnerAccount(v string) *CreateWhatsappConversionApiRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateWhatsappConversionApiRequest
	GetResourceOwnerId() *int64
}

type CreateWhatsappConversionApiRequest struct {
	// example:
	//
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 131
	InstanceId           *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Permissions          []*string `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateWhatsappConversionApiRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWhatsappConversionApiRequest) GoString() string {
	return s.String()
}

func (s *CreateWhatsappConversionApiRequest) GetCode() *string {
	return s.Code
}

func (s *CreateWhatsappConversionApiRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateWhatsappConversionApiRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateWhatsappConversionApiRequest) GetPermissions() []*string {
	return s.Permissions
}

func (s *CreateWhatsappConversionApiRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateWhatsappConversionApiRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateWhatsappConversionApiRequest) SetCode(v string) *CreateWhatsappConversionApiRequest {
	s.Code = &v
	return s
}

func (s *CreateWhatsappConversionApiRequest) SetInstanceId(v string) *CreateWhatsappConversionApiRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateWhatsappConversionApiRequest) SetOwnerId(v int64) *CreateWhatsappConversionApiRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateWhatsappConversionApiRequest) SetPermissions(v []*string) *CreateWhatsappConversionApiRequest {
	s.Permissions = v
	return s
}

func (s *CreateWhatsappConversionApiRequest) SetResourceOwnerAccount(v string) *CreateWhatsappConversionApiRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateWhatsappConversionApiRequest) SetResourceOwnerId(v int64) *CreateWhatsappConversionApiRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateWhatsappConversionApiRequest) Validate() error {
	return dara.Validate(s)
}
