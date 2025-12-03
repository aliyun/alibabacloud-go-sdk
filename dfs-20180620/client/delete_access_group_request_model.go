// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupId(v string) *DeleteAccessGroupRequest
	GetAccessGroupId() *string
	SetInputRegionId(v string) *DeleteAccessGroupRequest
	GetInputRegionId() *string
}

type DeleteAccessGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// acg-e3755fb0-358d-4286-9942-8d461048****
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
}

func (s DeleteAccessGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessGroupRequest) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *DeleteAccessGroupRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *DeleteAccessGroupRequest) SetAccessGroupId(v string) *DeleteAccessGroupRequest {
	s.AccessGroupId = &v
	return s
}

func (s *DeleteAccessGroupRequest) SetInputRegionId(v string) *DeleteAccessGroupRequest {
	s.InputRegionId = &v
	return s
}

func (s *DeleteAccessGroupRequest) Validate() error {
	return dara.Validate(s)
}
