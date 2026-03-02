// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMicroServiceDigest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceId(v int64) *MicroServiceDigest
	GetServiceId() *int64
	SetServiceName(v string) *MicroServiceDigest
	GetServiceName() *string
	SetLaneMqGroupInfos(v *LaneMqGroupInfo) *MicroServiceDigest
	GetLaneMqGroupInfos() *LaneMqGroupInfo
}

type MicroServiceDigest struct {
	ServiceId        *int64           `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName      *string          `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	LaneMqGroupInfos *LaneMqGroupInfo `json:"laneMqGroupInfos,omitempty" xml:"laneMqGroupInfos,omitempty"`
}

func (s MicroServiceDigest) String() string {
	return dara.Prettify(s)
}

func (s MicroServiceDigest) GoString() string {
	return s.String()
}

func (s *MicroServiceDigest) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *MicroServiceDigest) GetServiceName() *string {
	return s.ServiceName
}

func (s *MicroServiceDigest) GetLaneMqGroupInfos() *LaneMqGroupInfo {
	return s.LaneMqGroupInfos
}

func (s *MicroServiceDigest) SetServiceId(v int64) *MicroServiceDigest {
	s.ServiceId = &v
	return s
}

func (s *MicroServiceDigest) SetServiceName(v string) *MicroServiceDigest {
	s.ServiceName = &v
	return s
}

func (s *MicroServiceDigest) SetLaneMqGroupInfos(v *LaneMqGroupInfo) *MicroServiceDigest {
	s.LaneMqGroupInfos = v
	return s
}

func (s *MicroServiceDigest) Validate() error {
	if s.LaneMqGroupInfos != nil {
		if err := s.LaneMqGroupInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}
