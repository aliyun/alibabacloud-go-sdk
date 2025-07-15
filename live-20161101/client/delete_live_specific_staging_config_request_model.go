// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveSpecificStagingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *DeleteLiveSpecificStagingConfigRequest
	GetConfigId() *string
	SetDomainName(v string) *DeleteLiveSpecificStagingConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLiveSpecificStagingConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteLiveSpecificStagingConfigRequest
	GetSecurityToken() *string
}

type DeleteLiveSpecificStagingConfigRequest struct {
	// The ID of the configuration that you want to delete. If you want to specify multiple IDs, separate them with commas (,). You can call the [DescribeLiveDomainStagingConfig](https://help.aliyun.com/document_detail/297374.html) operation to obtain the configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6295
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The accelerated domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// developer.aliyundoc.com
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteLiveSpecificStagingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSpecificStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveSpecificStagingConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DeleteLiveSpecificStagingConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveSpecificStagingConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveSpecificStagingConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteLiveSpecificStagingConfigRequest) SetConfigId(v string) *DeleteLiveSpecificStagingConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteLiveSpecificStagingConfigRequest) SetDomainName(v string) *DeleteLiveSpecificStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveSpecificStagingConfigRequest) SetOwnerId(v int64) *DeleteLiveSpecificStagingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveSpecificStagingConfigRequest) SetSecurityToken(v string) *DeleteLiveSpecificStagingConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveSpecificStagingConfigRequest) Validate() error {
	return dara.Validate(s)
}
