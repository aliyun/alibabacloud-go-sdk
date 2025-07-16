// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetReqHeaderConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *SetReqHeaderConfigRequest
	GetConfigId() *int64
	SetDomainName(v string) *SetReqHeaderConfigRequest
	GetDomainName() *string
	SetKey(v string) *SetReqHeaderConfigRequest
	GetKey() *string
	SetOwnerId(v int64) *SetReqHeaderConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *SetReqHeaderConfigRequest
	GetSecurityToken() *string
	SetValue(v string) *SetReqHeaderConfigRequest
	GetValue() *string
}

type SetReqHeaderConfigRequest struct {
	// The ID of the configuration.
	//
	// example:
	//
	// 123
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The accelerated domain name. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the custom header.
	//
	// This parameter is required.
	//
	// example:
	//
	// testkey
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The value of the custom header.
	//
	// This parameter is required.
	//
	// example:
	//
	// testvalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SetReqHeaderConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetReqHeaderConfigRequest) GoString() string {
	return s.String()
}

func (s *SetReqHeaderConfigRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *SetReqHeaderConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetReqHeaderConfigRequest) GetKey() *string {
	return s.Key
}

func (s *SetReqHeaderConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetReqHeaderConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetReqHeaderConfigRequest) GetValue() *string {
	return s.Value
}

func (s *SetReqHeaderConfigRequest) SetConfigId(v int64) *SetReqHeaderConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *SetReqHeaderConfigRequest) SetDomainName(v string) *SetReqHeaderConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetReqHeaderConfigRequest) SetKey(v string) *SetReqHeaderConfigRequest {
	s.Key = &v
	return s
}

func (s *SetReqHeaderConfigRequest) SetOwnerId(v int64) *SetReqHeaderConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetReqHeaderConfigRequest) SetSecurityToken(v string) *SetReqHeaderConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetReqHeaderConfigRequest) SetValue(v string) *SetReqHeaderConfigRequest {
	s.Value = &v
	return s
}

func (s *SetReqHeaderConfigRequest) Validate() error {
	return dara.Validate(s)
}
