// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnmountInstanceSDGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *UnmountInstanceSDGRequest
	GetInstanceIds() []*string
	SetSDGId(v string) *UnmountInstanceSDGRequest
	GetSDGId() *string
}

type UnmountInstanceSDGRequest struct {
	// The IDs of the instances.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the SDG.
	//
	// This parameter is required.
	//
	// example:
	//
	// sdg-xxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s UnmountInstanceSDGRequest) String() string {
	return dara.Prettify(s)
}

func (s UnmountInstanceSDGRequest) GoString() string {
	return s.String()
}

func (s *UnmountInstanceSDGRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *UnmountInstanceSDGRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *UnmountInstanceSDGRequest) SetInstanceIds(v []*string) *UnmountInstanceSDGRequest {
	s.InstanceIds = v
	return s
}

func (s *UnmountInstanceSDGRequest) SetSDGId(v string) *UnmountInstanceSDGRequest {
	s.SDGId = &v
	return s
}

func (s *UnmountInstanceSDGRequest) Validate() error {
	return dara.Validate(s)
}
