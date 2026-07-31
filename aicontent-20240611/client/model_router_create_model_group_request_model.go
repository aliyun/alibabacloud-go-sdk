// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateModelGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelList(v []*int64) *ModelRouterCreateModelGroupRequest
	GetModelList() []*int64
	SetName(v string) *ModelRouterCreateModelGroupRequest
	GetName() *string
}

type ModelRouterCreateModelGroupRequest struct {
	// The array of model IDs. At least one element is required. Each element must be the numeric model ID, not the model identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// [101, 102, 103]
	ModelList []*int64 `json:"modelList,omitempty" xml:"modelList,omitempty" type:"Repeated"`
	// The group name. The name must be 1 to 50 characters in length and must be unique within the tenant (case-insensitive).
	//
	// This parameter is required.
	//
	// example:
	//
	// Professional Plan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ModelRouterCreateModelGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateModelGroupRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateModelGroupRequest) GetModelList() []*int64 {
	return s.ModelList
}

func (s *ModelRouterCreateModelGroupRequest) GetName() *string {
	return s.Name
}

func (s *ModelRouterCreateModelGroupRequest) SetModelList(v []*int64) *ModelRouterCreateModelGroupRequest {
	s.ModelList = v
	return s
}

func (s *ModelRouterCreateModelGroupRequest) SetName(v string) *ModelRouterCreateModelGroupRequest {
	s.Name = &v
	return s
}

func (s *ModelRouterCreateModelGroupRequest) Validate() error {
	return dara.Validate(s)
}
