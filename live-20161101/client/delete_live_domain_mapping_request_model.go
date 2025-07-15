// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveDomainMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteLiveDomainMappingRequest
	GetOwnerId() *int64
	SetPullDomain(v string) *DeleteLiveDomainMappingRequest
	GetPullDomain() *string
	SetPushDomain(v string) *DeleteLiveDomainMappingRequest
	GetPushDomain() *string
	SetSecurityToken(v string) *DeleteLiveDomainMappingRequest
	GetSecurityToken() *string
}

type DeleteLiveDomainMappingRequest struct {
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

func (s DeleteLiveDomainMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveDomainMappingRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveDomainMappingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveDomainMappingRequest) GetPullDomain() *string {
	return s.PullDomain
}

func (s *DeleteLiveDomainMappingRequest) GetPushDomain() *string {
	return s.PushDomain
}

func (s *DeleteLiveDomainMappingRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteLiveDomainMappingRequest) SetOwnerId(v int64) *DeleteLiveDomainMappingRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveDomainMappingRequest) SetPullDomain(v string) *DeleteLiveDomainMappingRequest {
	s.PullDomain = &v
	return s
}

func (s *DeleteLiveDomainMappingRequest) SetPushDomain(v string) *DeleteLiveDomainMappingRequest {
	s.PushDomain = &v
	return s
}

func (s *DeleteLiveDomainMappingRequest) SetSecurityToken(v string) *DeleteLiveDomainMappingRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveDomainMappingRequest) Validate() error {
	return dara.Validate(s)
}
