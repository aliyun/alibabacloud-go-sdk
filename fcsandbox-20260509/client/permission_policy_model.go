// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPermissionPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetAllow(v *PermissionPolicyAllow) *PermissionPolicy
	GetAllow() *PermissionPolicyAllow
	SetCatalogVersion(v int64) *PermissionPolicy
	GetCatalogVersion() *int64
	SetDeny(v *PermissionPolicyDeny) *PermissionPolicy
	GetDeny() *PermissionPolicyDeny
	SetSchemaVersion(v int64) *PermissionPolicy
	GetSchemaVersion() *int64
}

type PermissionPolicy struct {
	// The allow policy.
	Allow *PermissionPolicyAllow `json:"allow,omitempty" xml:"allow,omitempty" type:"Struct"`
	// catalog version
	CatalogVersion *int64 `json:"catalogVersion,omitempty" xml:"catalogVersion,omitempty"`
	// The deny policy.
	Deny *PermissionPolicyDeny `json:"deny,omitempty" xml:"deny,omitempty" type:"Struct"`
	// schema version
	SchemaVersion *int64 `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
}

func (s PermissionPolicy) String() string {
	return dara.Prettify(s)
}

func (s PermissionPolicy) GoString() string {
	return s.String()
}

func (s *PermissionPolicy) GetAllow() *PermissionPolicyAllow {
	return s.Allow
}

func (s *PermissionPolicy) GetCatalogVersion() *int64 {
	return s.CatalogVersion
}

func (s *PermissionPolicy) GetDeny() *PermissionPolicyDeny {
	return s.Deny
}

func (s *PermissionPolicy) GetSchemaVersion() *int64 {
	return s.SchemaVersion
}

func (s *PermissionPolicy) SetAllow(v *PermissionPolicyAllow) *PermissionPolicy {
	s.Allow = v
	return s
}

func (s *PermissionPolicy) SetCatalogVersion(v int64) *PermissionPolicy {
	s.CatalogVersion = &v
	return s
}

func (s *PermissionPolicy) SetDeny(v *PermissionPolicyDeny) *PermissionPolicy {
	s.Deny = v
	return s
}

func (s *PermissionPolicy) SetSchemaVersion(v int64) *PermissionPolicy {
	s.SchemaVersion = &v
	return s
}

func (s *PermissionPolicy) Validate() error {
	if s.Allow != nil {
		if err := s.Allow.Validate(); err != nil {
			return err
		}
	}
	if s.Deny != nil {
		if err := s.Deny.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PermissionPolicyAllow struct {
	// The actions.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The capabilities.
	Capabilities []*string `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Repeated"`
}

func (s PermissionPolicyAllow) String() string {
	return dara.Prettify(s)
}

func (s PermissionPolicyAllow) GoString() string {
	return s.String()
}

func (s *PermissionPolicyAllow) GetActions() []*string {
	return s.Actions
}

func (s *PermissionPolicyAllow) GetCapabilities() []*string {
	return s.Capabilities
}

func (s *PermissionPolicyAllow) SetActions(v []*string) *PermissionPolicyAllow {
	s.Actions = v
	return s
}

func (s *PermissionPolicyAllow) SetCapabilities(v []*string) *PermissionPolicyAllow {
	s.Capabilities = v
	return s
}

func (s *PermissionPolicyAllow) Validate() error {
	return dara.Validate(s)
}

type PermissionPolicyDeny struct {
	// The actions.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The capabilities.
	Capabilities []*string `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Repeated"`
}

func (s PermissionPolicyDeny) String() string {
	return dara.Prettify(s)
}

func (s PermissionPolicyDeny) GoString() string {
	return s.String()
}

func (s *PermissionPolicyDeny) GetActions() []*string {
	return s.Actions
}

func (s *PermissionPolicyDeny) GetCapabilities() []*string {
	return s.Capabilities
}

func (s *PermissionPolicyDeny) SetActions(v []*string) *PermissionPolicyDeny {
	s.Actions = v
	return s
}

func (s *PermissionPolicyDeny) SetCapabilities(v []*string) *PermissionPolicyDeny {
	s.Capabilities = v
	return s
}

func (s *PermissionPolicyDeny) Validate() error {
	return dara.Validate(s)
}
