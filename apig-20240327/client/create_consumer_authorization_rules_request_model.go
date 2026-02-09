// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerAuthorizationRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRules(v []*CreateConsumerAuthorizationRulesRequestAuthorizationRules) *CreateConsumerAuthorizationRulesRequest
	GetAuthorizationRules() []*CreateConsumerAuthorizationRulesRequestAuthorizationRules
}

type CreateConsumerAuthorizationRulesRequest struct {
	// The consumer authentication rules to be created.
	AuthorizationRules []*CreateConsumerAuthorizationRulesRequestAuthorizationRules `json:"authorizationRules,omitempty" xml:"authorizationRules,omitempty" type:"Repeated"`
}

func (s CreateConsumerAuthorizationRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerAuthorizationRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerAuthorizationRulesRequest) GetAuthorizationRules() []*CreateConsumerAuthorizationRulesRequestAuthorizationRules {
	return s.AuthorizationRules
}

func (s *CreateConsumerAuthorizationRulesRequest) SetAuthorizationRules(v []*CreateConsumerAuthorizationRulesRequestAuthorizationRules) *CreateConsumerAuthorizationRulesRequest {
	s.AuthorizationRules = v
	return s
}

func (s *CreateConsumerAuthorizationRulesRequest) Validate() error {
	if s.AuthorizationRules != nil {
		for _, item := range s.AuthorizationRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateConsumerAuthorizationRulesRequestAuthorizationRules struct {
	// The consumer ID.
	//
	// example:
	//
	// cs-cu08olem1hkokaut34i0
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// The expiration mode. Valid values: LongTerm and ShortTerm.
	//
	// example:
	//
	// LongTerm
	ExpireMode *string `json:"expireMode,omitempty" xml:"expireMode,omitempty"`
	// The expiration timestamp.
	//
	// example:
	//
	// 174116222x
	ExpireTimestamp *int64 `json:"expireTimestamp,omitempty" xml:"expireTimestamp,omitempty"`
	// The resource identifier, which is provided to non-standard code sources for space reuse.
	ResourceIdentifier *CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier `json:"resourceIdentifier,omitempty" xml:"resourceIdentifier,omitempty" type:"Struct"`
	// The resource type.
	//
	// example:
	//
	// HttpApiRoute
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s CreateConsumerAuthorizationRulesRequestAuthorizationRules) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerAuthorizationRulesRequestAuthorizationRules) GoString() string {
	return s.String()
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRules) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRules) GetExpireMode() *string {
	return s.ExpireMode
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRules) GetExpireTimestamp() *int64 {
	return s.ExpireTimestamp
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRules) GetResourceIdentifier() *CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier {
	return s.ResourceIdentifier
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRules) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRules) SetConsumerId(v string) *CreateConsumerAuthorizationRulesRequestAuthorizationRules {
	s.ConsumerId = &v
	return s
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRules) SetExpireMode(v string) *CreateConsumerAuthorizationRulesRequestAuthorizationRules {
	s.ExpireMode = &v
	return s
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRules) SetExpireTimestamp(v int64) *CreateConsumerAuthorizationRulesRequestAuthorizationRules {
	s.ExpireTimestamp = &v
	return s
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRules) SetResourceIdentifier(v *CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier) *CreateConsumerAuthorizationRulesRequestAuthorizationRules {
	s.ResourceIdentifier = v
	return s
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRules) SetResourceType(v string) *CreateConsumerAuthorizationRulesRequestAuthorizationRules {
	s.ResourceType = &v
	return s
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRules) Validate() error {
	if s.ResourceIdentifier != nil {
		if err := s.ResourceIdentifier.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier struct {
	// The environment ID.
	//
	// example:
	//
	// env-cti17hem1hktoruj98ug
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Parent resource ID
	//
	// example:
	//
	// api-******
	ParentResourceId *string `json:"parentResourceId,omitempty" xml:"parentResourceId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// ha-cn-li942gy8p03
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// List of resources
	Resources []*string `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
}

func (s CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier) GoString() string {
	return s.String()
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier) GetParentResourceId() *string {
	return s.ParentResourceId
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier) GetResources() []*string {
	return s.Resources
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier) SetEnvironmentId(v string) *CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier {
	s.EnvironmentId = &v
	return s
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier) SetParentResourceId(v string) *CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier {
	s.ParentResourceId = &v
	return s
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier) SetResourceId(v string) *CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier {
	s.ResourceId = &v
	return s
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier) SetResources(v []*string) *CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier {
	s.Resources = v
	return s
}

func (s *CreateConsumerAuthorizationRulesRequestAuthorizationRulesResourceIdentifier) Validate() error {
	return dara.Validate(s)
}
