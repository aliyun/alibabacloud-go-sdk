// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserDefinedSgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteUserDefinedSgRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteUserDefinedSgRequest
	GetRegionId() *string
	SetSgIdList(v []*string) *DeleteUserDefinedSgRequest
	GetSgIdList() []*string
}

type DeleteUserDefinedSgRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SgIdList []*string `json:"SgIdList,omitempty" xml:"SgIdList,omitempty" type:"Repeated"`
}

func (s DeleteUserDefinedSgRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserDefinedSgRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserDefinedSgRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteUserDefinedSgRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteUserDefinedSgRequest) GetSgIdList() []*string {
	return s.SgIdList
}

func (s *DeleteUserDefinedSgRequest) SetInstanceId(v string) *DeleteUserDefinedSgRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteUserDefinedSgRequest) SetRegionId(v string) *DeleteUserDefinedSgRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUserDefinedSgRequest) SetSgIdList(v []*string) *DeleteUserDefinedSgRequest {
	s.SgIdList = v
	return s
}

func (s *DeleteUserDefinedSgRequest) Validate() error {
	return dara.Validate(s)
}
