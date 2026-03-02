// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLaneMqGroupInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAllowClean(v bool) *LaneMqGroupInfo
	GetAllowClean() *bool
	SetLaneId(v int64) *LaneMqGroupInfo
	GetLaneId() *int64
	SetLaneName(v string) *LaneMqGroupInfo
	GetLaneName() *string
	SetMqGroups(v *MqGroup) *LaneMqGroupInfo
	GetMqGroups() *MqGroup
}

type LaneMqGroupInfo struct {
	AllowClean *bool    `json:"allowClean,omitempty" xml:"allowClean,omitempty"`
	LaneId     *int64   `json:"laneId,omitempty" xml:"laneId,omitempty"`
	LaneName   *string  `json:"laneName,omitempty" xml:"laneName,omitempty"`
	MqGroups   *MqGroup `json:"mqGroups,omitempty" xml:"mqGroups,omitempty"`
}

func (s LaneMqGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s LaneMqGroupInfo) GoString() string {
	return s.String()
}

func (s *LaneMqGroupInfo) GetAllowClean() *bool {
	return s.AllowClean
}

func (s *LaneMqGroupInfo) GetLaneId() *int64 {
	return s.LaneId
}

func (s *LaneMqGroupInfo) GetLaneName() *string {
	return s.LaneName
}

func (s *LaneMqGroupInfo) GetMqGroups() *MqGroup {
	return s.MqGroups
}

func (s *LaneMqGroupInfo) SetAllowClean(v bool) *LaneMqGroupInfo {
	s.AllowClean = &v
	return s
}

func (s *LaneMqGroupInfo) SetLaneId(v int64) *LaneMqGroupInfo {
	s.LaneId = &v
	return s
}

func (s *LaneMqGroupInfo) SetLaneName(v string) *LaneMqGroupInfo {
	s.LaneName = &v
	return s
}

func (s *LaneMqGroupInfo) SetMqGroups(v *MqGroup) *LaneMqGroupInfo {
	s.MqGroups = v
	return s
}

func (s *LaneMqGroupInfo) Validate() error {
	if s.MqGroups != nil {
		if err := s.MqGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}
