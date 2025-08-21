// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachInstanceSDGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DetachInstanceSDGRequest
	GetInstanceIds() []*string
	SetSDGId(v string) *DetachInstanceSDGRequest
	GetSDGId() *string
}

type DetachInstanceSDGRequest struct {
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
	// sdg-xxxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s DetachInstanceSDGRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachInstanceSDGRequest) GoString() string {
	return s.String()
}

func (s *DetachInstanceSDGRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DetachInstanceSDGRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *DetachInstanceSDGRequest) SetInstanceIds(v []*string) *DetachInstanceSDGRequest {
	s.InstanceIds = v
	return s
}

func (s *DetachInstanceSDGRequest) SetSDGId(v string) *DetachInstanceSDGRequest {
	s.SDGId = &v
	return s
}

func (s *DetachInstanceSDGRequest) Validate() error {
	return dara.Validate(s)
}
