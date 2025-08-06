// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCustomerConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIConfig(v string) *SetCustomerConfigRequest
	GetAIConfig() *string
	SetAppId(v string) *SetCustomerConfigRequest
	GetAppId() *string
	SetAuditConfig(v string) *SetCustomerConfigRequest
	GetAuditConfig() *string
	SetDownloadSwitch(v string) *SetCustomerConfigRequest
	GetDownloadSwitch() *string
	SetMetricConfig(v string) *SetCustomerConfigRequest
	GetMetricConfig() *string
	SetOwnerId(v int64) *SetCustomerConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SetCustomerConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetCustomerConfigRequest
	GetResourceOwnerId() *int64
}

type SetCustomerConfigRequest struct {
	AIConfig             *string `json:"AIConfig,omitempty" xml:"AIConfig,omitempty"`
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AuditConfig          *string `json:"AuditConfig,omitempty" xml:"AuditConfig,omitempty"`
	DownloadSwitch       *string `json:"DownloadSwitch,omitempty" xml:"DownloadSwitch,omitempty"`
	MetricConfig         *string `json:"MetricConfig,omitempty" xml:"MetricConfig,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetCustomerConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCustomerConfigRequest) GoString() string {
	return s.String()
}

func (s *SetCustomerConfigRequest) GetAIConfig() *string {
	return s.AIConfig
}

func (s *SetCustomerConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *SetCustomerConfigRequest) GetAuditConfig() *string {
	return s.AuditConfig
}

func (s *SetCustomerConfigRequest) GetDownloadSwitch() *string {
	return s.DownloadSwitch
}

func (s *SetCustomerConfigRequest) GetMetricConfig() *string {
	return s.MetricConfig
}

func (s *SetCustomerConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetCustomerConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetCustomerConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetCustomerConfigRequest) SetAIConfig(v string) *SetCustomerConfigRequest {
	s.AIConfig = &v
	return s
}

func (s *SetCustomerConfigRequest) SetAppId(v string) *SetCustomerConfigRequest {
	s.AppId = &v
	return s
}

func (s *SetCustomerConfigRequest) SetAuditConfig(v string) *SetCustomerConfigRequest {
	s.AuditConfig = &v
	return s
}

func (s *SetCustomerConfigRequest) SetDownloadSwitch(v string) *SetCustomerConfigRequest {
	s.DownloadSwitch = &v
	return s
}

func (s *SetCustomerConfigRequest) SetMetricConfig(v string) *SetCustomerConfigRequest {
	s.MetricConfig = &v
	return s
}

func (s *SetCustomerConfigRequest) SetOwnerId(v int64) *SetCustomerConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCustomerConfigRequest) SetResourceOwnerAccount(v string) *SetCustomerConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetCustomerConfigRequest) SetResourceOwnerId(v int64) *SetCustomerConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetCustomerConfigRequest) Validate() error {
	return dara.Validate(s)
}
