// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveDomainMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *AddLiveDomainMappingRequest
	GetOwnerId() *int64
	SetPullDomain(v string) *AddLiveDomainMappingRequest
	GetPullDomain() *string
	SetPushDomain(v string) *AddLiveDomainMappingRequest
	GetPushDomain() *string
	SetSecurityToken(v string) *AddLiveDomainMappingRequest
	GetSecurityToken() *string
}

type AddLiveDomainMappingRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The streaming domain. The type of the domain name is **liveVideo**.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	PullDomain *string `json:"PullDomain,omitempty" xml:"PullDomain,omitempty"`
	// The ingest domain. The type of the domain name is **liveEdge**.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	PushDomain    *string `json:"PushDomain,omitempty" xml:"PushDomain,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AddLiveDomainMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveDomainMappingRequest) GoString() string {
	return s.String()
}

func (s *AddLiveDomainMappingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveDomainMappingRequest) GetPullDomain() *string {
	return s.PullDomain
}

func (s *AddLiveDomainMappingRequest) GetPushDomain() *string {
	return s.PushDomain
}

func (s *AddLiveDomainMappingRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AddLiveDomainMappingRequest) SetOwnerId(v int64) *AddLiveDomainMappingRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveDomainMappingRequest) SetPullDomain(v string) *AddLiveDomainMappingRequest {
	s.PullDomain = &v
	return s
}

func (s *AddLiveDomainMappingRequest) SetPushDomain(v string) *AddLiveDomainMappingRequest {
	s.PushDomain = &v
	return s
}

func (s *AddLiveDomainMappingRequest) SetSecurityToken(v string) *AddLiveDomainMappingRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddLiveDomainMappingRequest) Validate() error {
	return dara.Validate(s)
}
