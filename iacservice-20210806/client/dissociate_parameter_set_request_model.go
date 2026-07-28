// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateParameterSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParameterSetIds(v []*string) *DissociateParameterSetRequest
	GetParameterSetIds() []*string
	SetResourceId(v string) *DissociateParameterSetRequest
	GetResourceId() *string
	SetResourceType(v string) *DissociateParameterSetRequest
	GetResourceType() *string
}

type DissociateParameterSetRequest struct {
	// The list of parameter set IDs to be associated with the resource. Maximum length: 5.
	//
	// This parameter is required.
	ParameterSetIds []*string `json:"parameterSetIds,omitempty" xml:"parameterSetIds,omitempty" type:"Repeated"`
	// The resource ID. If the resource type is ModuleVersion, the value is a combination of <moduleId>-<moduleversion>, such as mod-34535345df123fr-v3.
	//
	// This parameter is required.
	//
	// example:
	//
	// mod-39cd1e5e58c50e79dd8cd901cd
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type. Valid values:
	//
	// - Module: template
	//
	// - ModuleVersion: template version
	//
	// - Task: node
	//
	// - Stack: resource stack.
	//
	// This parameter is required.
	//
	// example:
	//
	// Module
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s DissociateParameterSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateParameterSetRequest) GoString() string {
	return s.String()
}

func (s *DissociateParameterSetRequest) GetParameterSetIds() []*string {
	return s.ParameterSetIds
}

func (s *DissociateParameterSetRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DissociateParameterSetRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DissociateParameterSetRequest) SetParameterSetIds(v []*string) *DissociateParameterSetRequest {
	s.ParameterSetIds = v
	return s
}

func (s *DissociateParameterSetRequest) SetResourceId(v string) *DissociateParameterSetRequest {
	s.ResourceId = &v
	return s
}

func (s *DissociateParameterSetRequest) SetResourceType(v string) *DissociateParameterSetRequest {
	s.ResourceType = &v
	return s
}

func (s *DissociateParameterSetRequest) Validate() error {
	return dara.Validate(s)
}
