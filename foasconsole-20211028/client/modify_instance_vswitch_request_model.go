// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVswitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHaVSwitchIds(v []*string) *ModifyInstanceVswitchRequest
	GetHaVSwitchIds() []*string
	SetInstanceId(v string) *ModifyInstanceVswitchRequest
	GetInstanceId() *string
	SetVSwitchIds(v []*string) *ModifyInstanceVswitchRequest
	GetVSwitchIds() []*string
}

type ModifyInstanceVswitchRequest struct {
	HaVSwitchIds []*string `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// sc_flinkserverless_public_cn-7e22ae****
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s ModifyInstanceVswitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVswitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVswitchRequest) GetHaVSwitchIds() []*string {
	return s.HaVSwitchIds
}

func (s *ModifyInstanceVswitchRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceVswitchRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ModifyInstanceVswitchRequest) SetHaVSwitchIds(v []*string) *ModifyInstanceVswitchRequest {
	s.HaVSwitchIds = v
	return s
}

func (s *ModifyInstanceVswitchRequest) SetInstanceId(v string) *ModifyInstanceVswitchRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceVswitchRequest) SetVSwitchIds(v []*string) *ModifyInstanceVswitchRequest {
	s.VSwitchIds = v
	return s
}

func (s *ModifyInstanceVswitchRequest) Validate() error {
	return dara.Validate(s)
}
