// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AssociateGroupRequest
	GetClientToken() *string
	SetProjectId(v string) *AssociateGroupRequest
	GetProjectId() *string
	SetResourceIds(v []*string) *AssociateGroupRequest
	GetResourceIds() []*string
	SetResourceType(v string) *AssociateGroupRequest
	GetResourceType() *string
}

type AssociateGroupRequest struct {
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// p-433aead7560571a87349d054b4
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// This parameter is required.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// Task
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s AssociateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateGroupRequest) GoString() string {
	return s.String()
}

func (s *AssociateGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateGroupRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *AssociateGroupRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *AssociateGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *AssociateGroupRequest) SetClientToken(v string) *AssociateGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateGroupRequest) SetProjectId(v string) *AssociateGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *AssociateGroupRequest) SetResourceIds(v []*string) *AssociateGroupRequest {
	s.ResourceIds = v
	return s
}

func (s *AssociateGroupRequest) SetResourceType(v string) *AssociateGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *AssociateGroupRequest) Validate() error {
	return dara.Validate(s)
}
