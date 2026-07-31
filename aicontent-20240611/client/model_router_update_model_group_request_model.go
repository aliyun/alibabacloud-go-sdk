// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateModelGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelList(v []*int64) *ModelRouterUpdateModelGroupRequest
	GetModelList() []*int64
	SetName(v string) *ModelRouterUpdateModelGroupRequest
	GetName() *string
}

type ModelRouterUpdateModelGroupRequest struct {
	// The full member array. An empty array clears all members.
	//
	// This parameter is required.
	//
	// example:
	//
	// [101, 102, 103]
	ModelList []*int64 `json:"modelList,omitempty" xml:"modelList,omitempty" type:"Repeated"`
	// The group name. This parameter performs a full overwrite. Pass the current name even if you do not want to rename the group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Professional Plan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ModelRouterUpdateModelGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateModelGroupRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateModelGroupRequest) GetModelList() []*int64 {
	return s.ModelList
}

func (s *ModelRouterUpdateModelGroupRequest) GetName() *string {
	return s.Name
}

func (s *ModelRouterUpdateModelGroupRequest) SetModelList(v []*int64) *ModelRouterUpdateModelGroupRequest {
	s.ModelList = v
	return s
}

func (s *ModelRouterUpdateModelGroupRequest) SetName(v string) *ModelRouterUpdateModelGroupRequest {
	s.Name = &v
	return s
}

func (s *ModelRouterUpdateModelGroupRequest) Validate() error {
	return dara.Validate(s)
}
