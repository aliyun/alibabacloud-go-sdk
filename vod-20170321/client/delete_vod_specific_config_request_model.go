// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodSpecificConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *DeleteVodSpecificConfigRequest
	GetConfigId() *string
	SetDomainName(v string) *DeleteVodSpecificConfigRequest
	GetDomainName() *string
	SetEnv(v string) *DeleteVodSpecificConfigRequest
	GetEnv() *string
	SetOwnerId(v int64) *DeleteVodSpecificConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteVodSpecificConfigRequest
	GetSecurityToken() *string
}

type DeleteVodSpecificConfigRequest struct {
	// The ID of the configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2317****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The accelerated domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The environment from which the domain name configurations are deleted. Valid values:
	//
	// 	- online: production environment
	//
	// 	- gray: simulation environment
	//
	// example:
	//
	// online
	Env           *string `json:"Env,omitempty" xml:"Env,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteVodSpecificConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodSpecificConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteVodSpecificConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DeleteVodSpecificConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteVodSpecificConfigRequest) GetEnv() *string {
	return s.Env
}

func (s *DeleteVodSpecificConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVodSpecificConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteVodSpecificConfigRequest) SetConfigId(v string) *DeleteVodSpecificConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteVodSpecificConfigRequest) SetDomainName(v string) *DeleteVodSpecificConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteVodSpecificConfigRequest) SetEnv(v string) *DeleteVodSpecificConfigRequest {
	s.Env = &v
	return s
}

func (s *DeleteVodSpecificConfigRequest) SetOwnerId(v int64) *DeleteVodSpecificConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVodSpecificConfigRequest) SetSecurityToken(v string) *DeleteVodSpecificConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteVodSpecificConfigRequest) Validate() error {
	return dara.Validate(s)
}
