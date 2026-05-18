// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachVscFromMountPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetachInfos(v []*DetachVscFromMountPointsRequestDetachInfos) *DetachVscFromMountPointsRequest
	GetDetachInfos() []*DetachVscFromMountPointsRequestDetachInfos
	SetInputRegionId(v string) *DetachVscFromMountPointsRequest
	GetInputRegionId() *string
	SetUseAssumeRoleChkServerPerm(v bool) *DetachVscFromMountPointsRequest
	GetUseAssumeRoleChkServerPerm() *bool
}

type DetachVscFromMountPointsRequest struct {
	DetachInfos []*DetachVscFromMountPointsRequestDetachInfos `json:"DetachInfos,omitempty" xml:"DetachInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	InputRegionId              *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	UseAssumeRoleChkServerPerm *bool   `json:"UseAssumeRoleChkServerPerm,omitempty" xml:"UseAssumeRoleChkServerPerm,omitempty"`
}

func (s DetachVscFromMountPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachVscFromMountPointsRequest) GoString() string {
	return s.String()
}

func (s *DetachVscFromMountPointsRequest) GetDetachInfos() []*DetachVscFromMountPointsRequestDetachInfos {
	return s.DetachInfos
}

func (s *DetachVscFromMountPointsRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *DetachVscFromMountPointsRequest) GetUseAssumeRoleChkServerPerm() *bool {
	return s.UseAssumeRoleChkServerPerm
}

func (s *DetachVscFromMountPointsRequest) SetDetachInfos(v []*DetachVscFromMountPointsRequestDetachInfos) *DetachVscFromMountPointsRequest {
	s.DetachInfos = v
	return s
}

func (s *DetachVscFromMountPointsRequest) SetInputRegionId(v string) *DetachVscFromMountPointsRequest {
	s.InputRegionId = &v
	return s
}

func (s *DetachVscFromMountPointsRequest) SetUseAssumeRoleChkServerPerm(v bool) *DetachVscFromMountPointsRequest {
	s.UseAssumeRoleChkServerPerm = &v
	return s
}

func (s *DetachVscFromMountPointsRequest) Validate() error {
	if s.DetachInfos != nil {
		for _, item := range s.DetachInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetachVscFromMountPointsRequestDetachInfos struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MountPointId *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	VscId        *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	VscName      *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	VscType      *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s DetachVscFromMountPointsRequestDetachInfos) String() string {
	return dara.Prettify(s)
}

func (s DetachVscFromMountPointsRequestDetachInfos) GoString() string {
	return s.String()
}

func (s *DetachVscFromMountPointsRequestDetachInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachVscFromMountPointsRequestDetachInfos) GetMountPointId() *string {
	return s.MountPointId
}

func (s *DetachVscFromMountPointsRequestDetachInfos) GetVscId() *string {
	return s.VscId
}

func (s *DetachVscFromMountPointsRequestDetachInfos) GetVscName() *string {
	return s.VscName
}

func (s *DetachVscFromMountPointsRequestDetachInfos) GetVscType() *string {
	return s.VscType
}

func (s *DetachVscFromMountPointsRequestDetachInfos) SetInstanceId(v string) *DetachVscFromMountPointsRequestDetachInfos {
	s.InstanceId = &v
	return s
}

func (s *DetachVscFromMountPointsRequestDetachInfos) SetMountPointId(v string) *DetachVscFromMountPointsRequestDetachInfos {
	s.MountPointId = &v
	return s
}

func (s *DetachVscFromMountPointsRequestDetachInfos) SetVscId(v string) *DetachVscFromMountPointsRequestDetachInfos {
	s.VscId = &v
	return s
}

func (s *DetachVscFromMountPointsRequestDetachInfos) SetVscName(v string) *DetachVscFromMountPointsRequestDetachInfos {
	s.VscName = &v
	return s
}

func (s *DetachVscFromMountPointsRequestDetachInfos) SetVscType(v string) *DetachVscFromMountPointsRequestDetachInfos {
	s.VscType = &v
	return s
}

func (s *DetachVscFromMountPointsRequestDetachInfos) Validate() error {
	return dara.Validate(s)
}
