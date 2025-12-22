// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserDefinedSgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AddUserDefinedSgRequest
	GetInstanceId() *string
	SetRegionId(v string) *AddUserDefinedSgRequest
	GetRegionId() *string
	SetSgIdList(v []*string) *AddUserDefinedSgRequest
	GetSgIdList() []*string
}

type AddUserDefinedSgRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SgIdList []*string `json:"SgIdList,omitempty" xml:"SgIdList,omitempty" type:"Repeated"`
}

func (s AddUserDefinedSgRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserDefinedSgRequest) GoString() string {
	return s.String()
}

func (s *AddUserDefinedSgRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddUserDefinedSgRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddUserDefinedSgRequest) GetSgIdList() []*string {
	return s.SgIdList
}

func (s *AddUserDefinedSgRequest) SetInstanceId(v string) *AddUserDefinedSgRequest {
	s.InstanceId = &v
	return s
}

func (s *AddUserDefinedSgRequest) SetRegionId(v string) *AddUserDefinedSgRequest {
	s.RegionId = &v
	return s
}

func (s *AddUserDefinedSgRequest) SetSgIdList(v []*string) *AddUserDefinedSgRequest {
	s.SgIdList = v
	return s
}

func (s *AddUserDefinedSgRequest) Validate() error {
	return dara.Validate(s)
}
