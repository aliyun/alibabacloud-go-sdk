// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupId(v string) *GetAccessGroupRequest
	GetAccessGroupId() *string
	SetInputRegionId(v string) *GetAccessGroupRequest
	GetInputRegionId() *string
}

type GetAccessGroupRequest struct {
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

func (s GetAccessGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *GetAccessGroupRequest) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *GetAccessGroupRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *GetAccessGroupRequest) SetAccessGroupId(v string) *GetAccessGroupRequest {
	s.AccessGroupId = &v
	return s
}

func (s *GetAccessGroupRequest) SetInputRegionId(v string) *GetAccessGroupRequest {
	s.InputRegionId = &v
	return s
}

func (s *GetAccessGroupRequest) Validate() error {
	return dara.Validate(s)
}
