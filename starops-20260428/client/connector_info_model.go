// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnectorInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAuthentication(v *ConnectorAuthentication) *ConnectorInfo
	GetAuthentication() *ConnectorAuthentication
	SetCapabilityGrants(v []map[string]interface{}) *ConnectorInfo
	GetCapabilityGrants() []map[string]interface{}
	SetConfiguration(v map[string]interface{}) *ConnectorInfo
	GetConfiguration() map[string]interface{}
	SetConnectorName(v string) *ConnectorInfo
	GetConnectorName() *string
	SetCreateTime(v string) *ConnectorInfo
	GetCreateTime() *string
	SetDescription(v string) *ConnectorInfo
	GetDescription() *string
	SetDisplayName(v string) *ConnectorInfo
	GetDisplayName() *string
	SetEnabled(v bool) *ConnectorInfo
	GetEnabled() *bool
	SetEtag(v string) *ConnectorInfo
	GetEtag() *string
	SetName(v string) *ConnectorInfo
	GetName() *string
	SetPolicy(v map[string]interface{}) *ConnectorInfo
	GetPolicy() map[string]interface{}
	SetProvider(v string) *ConnectorInfo
	GetProvider() *string
	SetRevision(v int64) *ConnectorInfo
	GetRevision() *int64
	SetRuntime(v *ConnectorRuntime) *ConnectorInfo
	GetRuntime() *ConnectorRuntime
	SetStatus(v map[string]interface{}) *ConnectorInfo
	GetStatus() map[string]interface{}
	SetTarget(v map[string]interface{}) *ConnectorInfo
	GetTarget() map[string]interface{}
	SetUpdateTime(v string) *ConnectorInfo
	GetUpdateTime() *string
}

type ConnectorInfo struct {
	// Safe authentication identity
	//
	// This parameter is required.
	Authentication *ConnectorAuthentication `json:"authentication,omitempty" xml:"authentication,omitempty"`
	// The list of capabilities granted to the Connector.
	//
	// This parameter is required.
	CapabilityGrants []map[string]interface{} `json:"capabilityGrants,omitempty" xml:"capabilityGrants,omitempty" type:"Repeated"`
	// Provider configuration
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	Configuration map[string]interface{} `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// Connector name
	//
	// This parameter is required.
	//
	// example:
	//
	// cms2-prod
	ConnectorName *string `json:"connectorName,omitempty" xml:"connectorName,omitempty"`
	// Creation time
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-08-25T12:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Description
	//
	// example:
	//
	// Production observability data
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Display name
	//
	// This parameter is required.
	//
	// example:
	//
	// CMS 2.0 production workspace
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Specifies whether the Connector is enabled.
	//
	// This parameter is required.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// ETag
	//
	// This parameter is required.
	//
	// example:
	//
	// "connector-rev-1"
	Etag *string `json:"etag,omitempty" xml:"etag,omitempty"`
	// Digital employee name
	//
	// This parameter is required.
	//
	// example:
	//
	// production-ops
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The execution policy of the Connector.
	//
	// This parameter is required.
	Policy map[string]interface{} `json:"policy,omitempty" xml:"policy,omitempty"`
	// Provider
	//
	// This parameter is required.
	//
	// example:
	//
	// AlibabaCloudCms
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// Revision
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
	// The runtime configuration of the Connector.
	//
	// This parameter is required.
	Runtime *ConnectorRuntime `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// Resource status
	//
	// This parameter is required.
	Status map[string]interface{} `json:"status,omitempty" xml:"status,omitempty"`
	// Provider target
	//
	// This parameter is required.
	Target map[string]interface{} `json:"target,omitempty" xml:"target,omitempty"`
	// Update time
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-08-25T12:00:00Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ConnectorInfo) String() string {
	return dara.Prettify(s)
}

func (s ConnectorInfo) GoString() string {
	return s.String()
}

func (s *ConnectorInfo) GetAuthentication() *ConnectorAuthentication {
	return s.Authentication
}

func (s *ConnectorInfo) GetCapabilityGrants() []map[string]interface{} {
	return s.CapabilityGrants
}

func (s *ConnectorInfo) GetConfiguration() map[string]interface{} {
	return s.Configuration
}

func (s *ConnectorInfo) GetConnectorName() *string {
	return s.ConnectorName
}

func (s *ConnectorInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ConnectorInfo) GetDescription() *string {
	return s.Description
}

func (s *ConnectorInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ConnectorInfo) GetEnabled() *bool {
	return s.Enabled
}

func (s *ConnectorInfo) GetEtag() *string {
	return s.Etag
}

func (s *ConnectorInfo) GetName() *string {
	return s.Name
}

func (s *ConnectorInfo) GetPolicy() map[string]interface{} {
	return s.Policy
}

func (s *ConnectorInfo) GetProvider() *string {
	return s.Provider
}

func (s *ConnectorInfo) GetRevision() *int64 {
	return s.Revision
}

func (s *ConnectorInfo) GetRuntime() *ConnectorRuntime {
	return s.Runtime
}

func (s *ConnectorInfo) GetStatus() map[string]interface{} {
	return s.Status
}

func (s *ConnectorInfo) GetTarget() map[string]interface{} {
	return s.Target
}

func (s *ConnectorInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ConnectorInfo) SetAuthentication(v *ConnectorAuthentication) *ConnectorInfo {
	s.Authentication = v
	return s
}

func (s *ConnectorInfo) SetCapabilityGrants(v []map[string]interface{}) *ConnectorInfo {
	s.CapabilityGrants = v
	return s
}

func (s *ConnectorInfo) SetConfiguration(v map[string]interface{}) *ConnectorInfo {
	s.Configuration = v
	return s
}

func (s *ConnectorInfo) SetConnectorName(v string) *ConnectorInfo {
	s.ConnectorName = &v
	return s
}

func (s *ConnectorInfo) SetCreateTime(v string) *ConnectorInfo {
	s.CreateTime = &v
	return s
}

func (s *ConnectorInfo) SetDescription(v string) *ConnectorInfo {
	s.Description = &v
	return s
}

func (s *ConnectorInfo) SetDisplayName(v string) *ConnectorInfo {
	s.DisplayName = &v
	return s
}

func (s *ConnectorInfo) SetEnabled(v bool) *ConnectorInfo {
	s.Enabled = &v
	return s
}

func (s *ConnectorInfo) SetEtag(v string) *ConnectorInfo {
	s.Etag = &v
	return s
}

func (s *ConnectorInfo) SetName(v string) *ConnectorInfo {
	s.Name = &v
	return s
}

func (s *ConnectorInfo) SetPolicy(v map[string]interface{}) *ConnectorInfo {
	s.Policy = v
	return s
}

func (s *ConnectorInfo) SetProvider(v string) *ConnectorInfo {
	s.Provider = &v
	return s
}

func (s *ConnectorInfo) SetRevision(v int64) *ConnectorInfo {
	s.Revision = &v
	return s
}

func (s *ConnectorInfo) SetRuntime(v *ConnectorRuntime) *ConnectorInfo {
	s.Runtime = v
	return s
}

func (s *ConnectorInfo) SetStatus(v map[string]interface{}) *ConnectorInfo {
	s.Status = v
	return s
}

func (s *ConnectorInfo) SetTarget(v map[string]interface{}) *ConnectorInfo {
	s.Target = v
	return s
}

func (s *ConnectorInfo) SetUpdateTime(v string) *ConnectorInfo {
	s.UpdateTime = &v
	return s
}

func (s *ConnectorInfo) Validate() error {
	if s.Authentication != nil {
		if err := s.Authentication.Validate(); err != nil {
			return err
		}
	}
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	return nil
}
