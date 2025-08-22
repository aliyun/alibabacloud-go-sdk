// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnSpecificStagingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *DeleteDcdnSpecificStagingConfigRequest
	GetConfigId() *string
	SetDomainName(v string) *DeleteDcdnSpecificStagingConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteDcdnSpecificStagingConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteDcdnSpecificStagingConfigRequest
	GetSecurityToken() *string
}

type DeleteDcdnSpecificStagingConfigRequest struct {
	// The ID of the configuration to be deleted. You can specify multiple configuration IDs and separate them with commas (,).
	//
	// You can call the DescribeDcdnDomainStagingConfig operation to query the environment configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2317
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The accelerated domain name. You can specify only one domain name in each call.
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

func (s DeleteDcdnSpecificStagingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnSpecificStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSpecificStagingConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DeleteDcdnSpecificStagingConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteDcdnSpecificStagingConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDcdnSpecificStagingConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteDcdnSpecificStagingConfigRequest) SetConfigId(v string) *DeleteDcdnSpecificStagingConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteDcdnSpecificStagingConfigRequest) SetDomainName(v string) *DeleteDcdnSpecificStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteDcdnSpecificStagingConfigRequest) SetOwnerId(v int64) *DeleteDcdnSpecificStagingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDcdnSpecificStagingConfigRequest) SetSecurityToken(v string) *DeleteDcdnSpecificStagingConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteDcdnSpecificStagingConfigRequest) Validate() error {
	return dara.Validate(s)
}
