// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateModelGroupRequest
	GetDescription() *string
	SetExpectedVersion(v int64) *UpdateModelGroupRequest
	GetExpectedVersion() *int64
	SetModelCardIds(v []*string) *UpdateModelGroupRequest
	GetModelCardIds() []*string
	SetName(v string) *UpdateModelGroupRequest
	GetName() *string
}

type UpdateModelGroupRequest struct {
	// The nullable description. The maximum length is 1024 UTF-16 code units.
	//
	// example:
	//
	// Model group for service integration
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The group version obtained from the most recent Get operation. A stale version returns Conflict.ModelGroupChanged. Submitting the same expectedVersion, name, description, and modelCardIds retries the Failed or PartiallyFailed configuration for the same version, reuses the operationId, and does not increment the group version. Pending, Running, or Succeeded states are not forcibly reset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ExpectedVersion *int64 `json:"expectedVersion,omitempty" xml:"expectedVersion,omitempty"`
	// The actual or logical model identities within the same instance. The list can be null and must not contain duplicates. Model groups are not accepted as members.
	//
	// This parameter is required.
	ModelCardIds []*string `json:"modelCardIds,omitempty" xml:"modelCardIds,omitempty" type:"Repeated"`
	// The unique name within the instance. The length is 1 to 128 UTF-8 bytes after trimming.
	//
	// This parameter is required.
	//
	// example:
	//
	// beta-model-group
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateModelGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelGroupRequest) GetExpectedVersion() *int64 {
	return s.ExpectedVersion
}

func (s *UpdateModelGroupRequest) GetModelCardIds() []*string {
	return s.ModelCardIds
}

func (s *UpdateModelGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateModelGroupRequest) SetDescription(v string) *UpdateModelGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateModelGroupRequest) SetExpectedVersion(v int64) *UpdateModelGroupRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *UpdateModelGroupRequest) SetModelCardIds(v []*string) *UpdateModelGroupRequest {
	s.ModelCardIds = v
	return s
}

func (s *UpdateModelGroupRequest) SetName(v string) *UpdateModelGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateModelGroupRequest) Validate() error {
	return dara.Validate(s)
}
