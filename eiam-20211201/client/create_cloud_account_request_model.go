// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateCloudAccountRequest
	GetClientToken() *string
	SetCloudAccountExternalId(v string) *CreateCloudAccountRequest
	GetCloudAccountExternalId() *string
	SetCloudAccountName(v string) *CreateCloudAccountRequest
	GetCloudAccountName() *string
	SetCloudAccountProviderName(v string) *CreateCloudAccountRequest
	GetCloudAccountProviderName() *string
	SetCloudAccountVendorType(v string) *CreateCloudAccountRequest
	GetCloudAccountVendorType() *string
	SetDescription(v string) *CreateCloudAccountRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateCloudAccountRequest
	GetInstanceId() *string
}

type CreateCloudAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 云账号唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	CloudAccountExternalId *string `json:"CloudAccountExternalId,omitempty" xml:"CloudAccountExternalId,omitempty"`
	// 云账号名称
	//
	// example:
	//
	// cloud_accout_xxxx
	CloudAccountName *string `json:"CloudAccountName,omitempty" xml:"CloudAccountName,omitempty"`
	// 云账号提供商名称
	//
	// example:
	//
	// idaas-eiam-oidc-provider
	CloudAccountProviderName *string `json:"CloudAccountProviderName,omitempty" xml:"CloudAccountProviderName,omitempty"`
	// 云账号类型
	//
	// This parameter is required.
	CloudAccountVendorType *string `json:"CloudAccountVendorType,omitempty" xml:"CloudAccountVendorType,omitempty"`
	// 云账号描述
	//
	// example:
	//
	// cloud_accout_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateCloudAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudAccountRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCloudAccountRequest) GetCloudAccountExternalId() *string {
	return s.CloudAccountExternalId
}

func (s *CreateCloudAccountRequest) GetCloudAccountName() *string {
	return s.CloudAccountName
}

func (s *CreateCloudAccountRequest) GetCloudAccountProviderName() *string {
	return s.CloudAccountProviderName
}

func (s *CreateCloudAccountRequest) GetCloudAccountVendorType() *string {
	return s.CloudAccountVendorType
}

func (s *CreateCloudAccountRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCloudAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCloudAccountRequest) SetClientToken(v string) *CreateCloudAccountRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCloudAccountRequest) SetCloudAccountExternalId(v string) *CreateCloudAccountRequest {
	s.CloudAccountExternalId = &v
	return s
}

func (s *CreateCloudAccountRequest) SetCloudAccountName(v string) *CreateCloudAccountRequest {
	s.CloudAccountName = &v
	return s
}

func (s *CreateCloudAccountRequest) SetCloudAccountProviderName(v string) *CreateCloudAccountRequest {
	s.CloudAccountProviderName = &v
	return s
}

func (s *CreateCloudAccountRequest) SetCloudAccountVendorType(v string) *CreateCloudAccountRequest {
	s.CloudAccountVendorType = &v
	return s
}

func (s *CreateCloudAccountRequest) SetDescription(v string) *CreateCloudAccountRequest {
	s.Description = &v
	return s
}

func (s *CreateCloudAccountRequest) SetInstanceId(v string) *CreateCloudAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCloudAccountRequest) Validate() error {
	return dara.Validate(s)
}
