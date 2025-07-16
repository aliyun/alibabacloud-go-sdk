// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSpecificConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *DeleteSpecificConfigRequest
	GetConfigId() *string
	SetDomainName(v string) *DeleteSpecificConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteSpecificConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteSpecificConfigRequest
	GetSecurityToken() *string
}

type DeleteSpecificConfigRequest struct {
	// The ID of the configuration. Separate multiple configuration IDs with commas (,). For more information about ConfigId, see [Usage notes on ConfigId](https://help.aliyun.com/document_detail/388994.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2317
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The accelerated domain name. You can specify only one domain name.
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

func (s DeleteSpecificConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSpecificConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteSpecificConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DeleteSpecificConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteSpecificConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSpecificConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteSpecificConfigRequest) SetConfigId(v string) *DeleteSpecificConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteSpecificConfigRequest) SetDomainName(v string) *DeleteSpecificConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteSpecificConfigRequest) SetOwnerId(v int64) *DeleteSpecificConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSpecificConfigRequest) SetSecurityToken(v string) *DeleteSpecificConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteSpecificConfigRequest) Validate() error {
	return dara.Validate(s)
}
