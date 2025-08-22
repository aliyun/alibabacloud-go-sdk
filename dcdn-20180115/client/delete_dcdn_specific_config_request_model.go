// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnSpecificConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *DeleteDcdnSpecificConfigRequest
	GetConfigId() *string
	SetDomainName(v string) *DeleteDcdnSpecificConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteDcdnSpecificConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteDcdnSpecificConfigRequest
	GetSecurityToken() *string
}

type DeleteDcdnSpecificConfigRequest struct {
	// The ID of the configuration. Separate multiple configuration IDs with commas (,). For more information about ConfigId, see [Usage notes on ConfigId](https://help.aliyun.com/document_detail/410558.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2117
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

func (s DeleteDcdnSpecificConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnSpecificConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSpecificConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DeleteDcdnSpecificConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteDcdnSpecificConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDcdnSpecificConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteDcdnSpecificConfigRequest) SetConfigId(v string) *DeleteDcdnSpecificConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteDcdnSpecificConfigRequest) SetDomainName(v string) *DeleteDcdnSpecificConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteDcdnSpecificConfigRequest) SetOwnerId(v int64) *DeleteDcdnSpecificConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDcdnSpecificConfigRequest) SetSecurityToken(v string) *DeleteDcdnSpecificConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteDcdnSpecificConfigRequest) Validate() error {
	return dara.Validate(s)
}
