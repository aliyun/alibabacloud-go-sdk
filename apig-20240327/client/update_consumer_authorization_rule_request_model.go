// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConsumerAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationResourceInfos(v []*AuthorizationResourceInfo) *UpdateConsumerAuthorizationRuleRequest
	GetAuthorizationResourceInfos() []*AuthorizationResourceInfo
	SetExpireMode(v string) *UpdateConsumerAuthorizationRuleRequest
	GetExpireMode() *string
	SetExpireTimestamp(v int64) *UpdateConsumerAuthorizationRuleRequest
	GetExpireTimestamp() *int64
}

type UpdateConsumerAuthorizationRuleRequest struct {
	AuthorizationResourceInfos []*AuthorizationResourceInfo `json:"authorizationResourceInfos,omitempty" xml:"authorizationResourceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// LongTerm
	ExpireMode *string `json:"expireMode,omitempty" xml:"expireMode,omitempty"`
	// example:
	//
	// 1750852089975
	ExpireTimestamp *int64 `json:"expireTimestamp,omitempty" xml:"expireTimestamp,omitempty"`
}

func (s UpdateConsumerAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateConsumerAuthorizationRuleRequest) GetAuthorizationResourceInfos() []*AuthorizationResourceInfo {
	return s.AuthorizationResourceInfos
}

func (s *UpdateConsumerAuthorizationRuleRequest) GetExpireMode() *string {
	return s.ExpireMode
}

func (s *UpdateConsumerAuthorizationRuleRequest) GetExpireTimestamp() *int64 {
	return s.ExpireTimestamp
}

func (s *UpdateConsumerAuthorizationRuleRequest) SetAuthorizationResourceInfos(v []*AuthorizationResourceInfo) *UpdateConsumerAuthorizationRuleRequest {
	s.AuthorizationResourceInfos = v
	return s
}

func (s *UpdateConsumerAuthorizationRuleRequest) SetExpireMode(v string) *UpdateConsumerAuthorizationRuleRequest {
	s.ExpireMode = &v
	return s
}

func (s *UpdateConsumerAuthorizationRuleRequest) SetExpireTimestamp(v int64) *UpdateConsumerAuthorizationRuleRequest {
	s.ExpireTimestamp = &v
	return s
}

func (s *UpdateConsumerAuthorizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
