// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSpecificStagingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *DeleteSpecificStagingConfigRequest
	GetConfigId() *string
	SetDomainName(v string) *DeleteSpecificStagingConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteSpecificStagingConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteSpecificStagingConfigRequest
	GetSecurityToken() *string
}

type DeleteSpecificStagingConfigRequest struct {
	// The configuration IDs. Separate configuration IDs with commas (,). For more information about ConfigId, see [Usage notes on ConfigId](https://help.aliyun.com/document_detail/388994.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2317
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The accelerated domain name. You can specify up to 50 domain names in each request. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteSpecificStagingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSpecificStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteSpecificStagingConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DeleteSpecificStagingConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteSpecificStagingConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSpecificStagingConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteSpecificStagingConfigRequest) SetConfigId(v string) *DeleteSpecificStagingConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteSpecificStagingConfigRequest) SetDomainName(v string) *DeleteSpecificStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteSpecificStagingConfigRequest) SetOwnerId(v int64) *DeleteSpecificStagingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSpecificStagingConfigRequest) SetSecurityToken(v string) *DeleteSpecificStagingConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteSpecificStagingConfigRequest) Validate() error {
	return dara.Validate(s)
}
