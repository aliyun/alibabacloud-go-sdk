// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserDefinedSgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyUserDefinedSgRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyUserDefinedSgRequest
	GetRegionId() *string
	SetSgIdList(v []*string) *ModifyUserDefinedSgRequest
	GetSgIdList() []*string
}

type ModifyUserDefinedSgRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SgIdList []*string `json:"SgIdList,omitempty" xml:"SgIdList,omitempty" type:"Repeated"`
}

func (s ModifyUserDefinedSgRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserDefinedSgRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserDefinedSgRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyUserDefinedSgRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyUserDefinedSgRequest) GetSgIdList() []*string {
	return s.SgIdList
}

func (s *ModifyUserDefinedSgRequest) SetInstanceId(v string) *ModifyUserDefinedSgRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUserDefinedSgRequest) SetRegionId(v string) *ModifyUserDefinedSgRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserDefinedSgRequest) SetSgIdList(v []*string) *ModifyUserDefinedSgRequest {
	s.SgIdList = v
	return s
}

func (s *ModifyUserDefinedSgRequest) Validate() error {
	return dara.Validate(s)
}
