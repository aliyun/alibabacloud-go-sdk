// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationResourceInfos(v []*AuthorizationResourceInfo) *CreateConsumerAuthorizationRuleRequest
	GetAuthorizationResourceInfos() []*AuthorizationResourceInfo
	SetExpireMode(v string) *CreateConsumerAuthorizationRuleRequest
	GetExpireMode() *string
	SetExpireTimestamp(v int64) *CreateConsumerAuthorizationRuleRequest
	GetExpireTimestamp() *int64
	SetParentResourceType(v string) *CreateConsumerAuthorizationRuleRequest
	GetParentResourceType() *string
	SetResourceType(v string) *CreateConsumerAuthorizationRuleRequest
	GetResourceType() *string
}

type CreateConsumerAuthorizationRuleRequest struct {
	AuthorizationResourceInfos []*AuthorizationResourceInfo `json:"authorizationResourceInfos,omitempty" xml:"authorizationResourceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// LongTerm
	ExpireMode *string `json:"expireMode,omitempty" xml:"expireMode,omitempty"`
	// example:
	//
	// 1750852089975
	ExpireTimestamp *int64 `json:"expireTimestamp,omitempty" xml:"expireTimestamp,omitempty"`
	// example:
	//
	// API
	ParentResourceType *string `json:"parentResourceType,omitempty" xml:"parentResourceType,omitempty"`
	// example:
	//
	// API
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s CreateConsumerAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerAuthorizationRuleRequest) GetAuthorizationResourceInfos() []*AuthorizationResourceInfo {
	return s.AuthorizationResourceInfos
}

func (s *CreateConsumerAuthorizationRuleRequest) GetExpireMode() *string {
	return s.ExpireMode
}

func (s *CreateConsumerAuthorizationRuleRequest) GetExpireTimestamp() *int64 {
	return s.ExpireTimestamp
}

func (s *CreateConsumerAuthorizationRuleRequest) GetParentResourceType() *string {
	return s.ParentResourceType
}

func (s *CreateConsumerAuthorizationRuleRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateConsumerAuthorizationRuleRequest) SetAuthorizationResourceInfos(v []*AuthorizationResourceInfo) *CreateConsumerAuthorizationRuleRequest {
	s.AuthorizationResourceInfos = v
	return s
}

func (s *CreateConsumerAuthorizationRuleRequest) SetExpireMode(v string) *CreateConsumerAuthorizationRuleRequest {
	s.ExpireMode = &v
	return s
}

func (s *CreateConsumerAuthorizationRuleRequest) SetExpireTimestamp(v int64) *CreateConsumerAuthorizationRuleRequest {
	s.ExpireTimestamp = &v
	return s
}

func (s *CreateConsumerAuthorizationRuleRequest) SetParentResourceType(v string) *CreateConsumerAuthorizationRuleRequest {
	s.ParentResourceType = &v
	return s
}

func (s *CreateConsumerAuthorizationRuleRequest) SetResourceType(v string) *CreateConsumerAuthorizationRuleRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateConsumerAuthorizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
