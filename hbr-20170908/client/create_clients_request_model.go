// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertSetting(v string) *CreateClientsRequest
	GetAlertSetting() *string
	SetClientInfo(v string) *CreateClientsRequest
	GetClientInfo() *string
	SetCrossAccountRoleName(v string) *CreateClientsRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *CreateClientsRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *CreateClientsRequest
	GetCrossAccountUserId() *int64
	SetResourceGroupId(v string) *CreateClientsRequest
	GetResourceGroupId() *string
	SetUseHttps(v bool) *CreateClientsRequest
	GetUseHttps() *bool
	SetVaultId(v string) *CreateClientsRequest
	GetVaultId() *string
}

type CreateClientsRequest struct {
	// The alert settings. Valid value: INHERITED, which indicates that the HBR client sends alert notifications by using the same method configured for the backup vault.
	//
	// example:
	//
	// INHERITED
	AlertSetting *string `json:"AlertSetting,omitempty" xml:"AlertSetting,omitempty"`
	// The installation information of the HBR clients.
	//
	// example:
	//
	// [  {    "instanceId": "i-bp116lr******te9q2",    "accessKeyId": "",    "accessKeySecret": "",    "clusterId": "cl-000csy09q******9rfz9",    "sourceTypes": [      "HANA"    ]  },  {    "instanceId": "i-bp116lrux******hte9q4",    "accessKeyId": "",    "accessKeySecret": "",    "clusterId": "cl-000csy09q5094vw9rfz9",    "sourceTypes": [      "HANA"    ]  }]
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	// The name of the Resource Access Management (RAM) role that is created within the source Alibaba Cloud account and assigned to the current Alibaba Cloud account to authorize the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// hbrcrossrole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// The backup type. Valid values:
	//
	// - **SELF_ACCOUNT**: Data is backed up within the same Alibaba Cloud account.
	//
	// - **CROSS_ACCOUNT**: Data is backed up across Alibaba Cloud accounts.
	//
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the source Alibaba Cloud account that authorizes the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// 158975xxxxx4625
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzvx7d3c4kpny
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to transmit data over HTTPS. Valid values:
	//
	// 	- true: transmits data over HTTPS.
	//
	// 	- false: transmits data over HTTP.
	//
	// example:
	//
	// false
	UseHttps *bool `json:"UseHttps,omitempty" xml:"UseHttps,omitempty"`
	// The ID of the backup vault.
	//
	// This parameter is required.
	//
	// example:
	//
	// v-0001ufe******kgm
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClientsRequest) GoString() string {
	return s.String()
}

func (s *CreateClientsRequest) GetAlertSetting() *string {
	return s.AlertSetting
}

func (s *CreateClientsRequest) GetClientInfo() *string {
	return s.ClientInfo
}

func (s *CreateClientsRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *CreateClientsRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *CreateClientsRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *CreateClientsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateClientsRequest) GetUseHttps() *bool {
	return s.UseHttps
}

func (s *CreateClientsRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *CreateClientsRequest) SetAlertSetting(v string) *CreateClientsRequest {
	s.AlertSetting = &v
	return s
}

func (s *CreateClientsRequest) SetClientInfo(v string) *CreateClientsRequest {
	s.ClientInfo = &v
	return s
}

func (s *CreateClientsRequest) SetCrossAccountRoleName(v string) *CreateClientsRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *CreateClientsRequest) SetCrossAccountType(v string) *CreateClientsRequest {
	s.CrossAccountType = &v
	return s
}

func (s *CreateClientsRequest) SetCrossAccountUserId(v int64) *CreateClientsRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *CreateClientsRequest) SetResourceGroupId(v string) *CreateClientsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClientsRequest) SetUseHttps(v bool) *CreateClientsRequest {
	s.UseHttps = &v
	return s
}

func (s *CreateClientsRequest) SetVaultId(v string) *CreateClientsRequest {
	s.VaultId = &v
	return s
}

func (s *CreateClientsRequest) Validate() error {
	return dara.Validate(s)
}
