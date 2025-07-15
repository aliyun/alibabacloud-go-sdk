// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAndroidInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceGroupIds(v []*string) *DeleteAndroidInstanceGroupRequest
	GetInstanceGroupIds() []*string
}

type DeleteAndroidInstanceGroupRequest struct {
	// The IDs of the instance groups.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
}

func (s DeleteAndroidInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAndroidInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAndroidInstanceGroupRequest) GetInstanceGroupIds() []*string {
	return s.InstanceGroupIds
}

func (s *DeleteAndroidInstanceGroupRequest) SetInstanceGroupIds(v []*string) *DeleteAndroidInstanceGroupRequest {
	s.InstanceGroupIds = v
	return s
}

func (s *DeleteAndroidInstanceGroupRequest) Validate() error {
	return dara.Validate(s)
}
