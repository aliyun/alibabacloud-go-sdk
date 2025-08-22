// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnIpaSpecificConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *DeleteDcdnIpaSpecificConfigRequest
	GetConfigId() *string
	SetDomainName(v string) *DeleteDcdnIpaSpecificConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteDcdnIpaSpecificConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteDcdnIpaSpecificConfigRequest
	GetSecurityToken() *string
}

type DeleteDcdnIpaSpecificConfigRequest struct {
	// The ID of the configuration. You can call the [DescribeDcdnDomainConfigs](https://help.aliyun.com/document_detail/130625.html) operation to query configuration IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50035**
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The accelerated domain name. You can specify only one domain name in each request.
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

func (s DeleteDcdnIpaSpecificConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnIpaSpecificConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnIpaSpecificConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DeleteDcdnIpaSpecificConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteDcdnIpaSpecificConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDcdnIpaSpecificConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteDcdnIpaSpecificConfigRequest) SetConfigId(v string) *DeleteDcdnIpaSpecificConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteDcdnIpaSpecificConfigRequest) SetDomainName(v string) *DeleteDcdnIpaSpecificConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteDcdnIpaSpecificConfigRequest) SetOwnerId(v int64) *DeleteDcdnIpaSpecificConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDcdnIpaSpecificConfigRequest) SetSecurityToken(v string) *DeleteDcdnIpaSpecificConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteDcdnIpaSpecificConfigRequest) Validate() error {
	return dara.Validate(s)
}
