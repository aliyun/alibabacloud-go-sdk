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
	SetCloudAccountSite(v string) *CreateCloudAccountRequest
	GetCloudAccountSite() *string
	SetCloudAccountVendorType(v string) *CreateCloudAccountRequest
	GetCloudAccountVendorType() *string
	SetDescription(v string) *CreateCloudAccountRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateCloudAccountRequest
	GetInstanceId() *string
}

type CreateCloudAccountRequest struct {
	// A client token that ensures the idempotence of the request. The client must generate a unique token for each request. The token must consist of only ASCII characters and be no more than 64 characters in length. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The external unique identifier of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	CloudAccountExternalId *string `json:"CloudAccountExternalId,omitempty" xml:"CloudAccountExternalId,omitempty"`
	// The name of the cloud account.
	//
	// example:
	//
	// cloud_accout_xxxx
	CloudAccountName *string `json:"CloudAccountName,omitempty" xml:"CloudAccountName,omitempty"`
	// The name of the identity provider.
	//
	// example:
	//
	// idaas-eiam-oidc-provider
	CloudAccountProviderName *string `json:"CloudAccountProviderName,omitempty" xml:"CloudAccountProviderName,omitempty"`
	// The site of the cloud account. Valid values:
	//
	// - global: international site
	//
	// - china_mainland: China mainland site
	//
	// example:
	//
	// china_mainland
	CloudAccountSite *string `json:"CloudAccountSite,omitempty" xml:"CloudAccountSite,omitempty"`
	// The type of the cloud account. Valid value:
	//
	// - alibaba_cloud: Alibaba Cloud
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba_cloud(Alibaba Cloud)
	CloudAccountVendorType *string `json:"CloudAccountVendorType,omitempty" xml:"CloudAccountVendorType,omitempty"`
	// The description of the cloud account.
	//
	// example:
	//
	// cloud_accout_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
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

func (s *CreateCloudAccountRequest) GetCloudAccountSite() *string {
	return s.CloudAccountSite
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

func (s *CreateCloudAccountRequest) SetCloudAccountSite(v string) *CreateCloudAccountRequest {
	s.CloudAccountSite = &v
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
