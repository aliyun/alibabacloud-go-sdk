// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterSetMemberAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedModelGroupConfig(v string) *ModelRouterSetMemberAuthorizationRequest
	GetAllowedModelGroupConfig() *string
	SetAllowedModels(v string) *ModelRouterSetMemberAuthorizationRequest
	GetAllowedModels() *string
}

type ModelRouterSetMemberAuthorizationRequest struct {
	// The authorization configuration (JSON string, overwrite mode): {"model_ids":[...],"group_ids":["mg_xxx"]}. The internal key names use a fixed underscore style and are not converted to the camelCase convention used by the API. If this field is specified together with allowedModels, this field takes precedence.
	//
	// example:
	//
	// {"model_ids":[],"group_ids":["mg_qwen_49"]}
	AllowedModelGroupConfig *string `json:"allowedModelGroupConfig,omitempty" xml:"allowedModelGroupConfig,omitempty"`
	// The legacy authorization field (comma-separated numeric model IDs). This field is retained during the canary release of group-based authorization: tenants that have not enabled the grouping feature continue to use this field. If this field is specified together with allowedModelGroupConfig, the latter takes precedence.
	//
	// example:
	//
	// 101,102,103
	AllowedModels *string `json:"allowedModels,omitempty" xml:"allowedModels,omitempty"`
}

func (s ModelRouterSetMemberAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterSetMemberAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterSetMemberAuthorizationRequest) GetAllowedModelGroupConfig() *string {
	return s.AllowedModelGroupConfig
}

func (s *ModelRouterSetMemberAuthorizationRequest) GetAllowedModels() *string {
	return s.AllowedModels
}

func (s *ModelRouterSetMemberAuthorizationRequest) SetAllowedModelGroupConfig(v string) *ModelRouterSetMemberAuthorizationRequest {
	s.AllowedModelGroupConfig = &v
	return s
}

func (s *ModelRouterSetMemberAuthorizationRequest) SetAllowedModels(v string) *ModelRouterSetMemberAuthorizationRequest {
	s.AllowedModels = &v
	return s
}

func (s *ModelRouterSetMemberAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
