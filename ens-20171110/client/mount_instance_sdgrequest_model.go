// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMountInstanceSDGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *MountInstanceSDGRequest
	GetInstanceIds() []*string
	SetSDGId(v string) *MountInstanceSDGRequest
	GetSDGId() *string
}

type MountInstanceSDGRequest struct {
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

func (s MountInstanceSDGRequest) String() string {
	return dara.Prettify(s)
}

func (s MountInstanceSDGRequest) GoString() string {
	return s.String()
}

func (s *MountInstanceSDGRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *MountInstanceSDGRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *MountInstanceSDGRequest) SetInstanceIds(v []*string) *MountInstanceSDGRequest {
	s.InstanceIds = v
	return s
}

func (s *MountInstanceSDGRequest) SetSDGId(v string) *MountInstanceSDGRequest {
	s.SDGId = &v
	return s
}

func (s *MountInstanceSDGRequest) Validate() error {
	return dara.Validate(s)
}
