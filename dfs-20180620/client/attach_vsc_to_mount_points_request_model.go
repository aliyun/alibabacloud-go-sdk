// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVscToMountPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachInfos(v []*AttachVscToMountPointsRequestAttachInfos) *AttachVscToMountPointsRequest
	GetAttachInfos() []*AttachVscToMountPointsRequestAttachInfos
	SetInputRegionId(v string) *AttachVscToMountPointsRequest
	GetInputRegionId() *string
	SetUseAssumeRoleChkServerPerm(v bool) *AttachVscToMountPointsRequest
	GetUseAssumeRoleChkServerPerm() *bool
}

type AttachVscToMountPointsRequest struct {
	AttachInfos []*AttachVscToMountPointsRequestAttachInfos `json:"AttachInfos,omitempty" xml:"AttachInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// example:
	//
	// false
	UseAssumeRoleChkServerPerm *bool `json:"UseAssumeRoleChkServerPerm,omitempty" xml:"UseAssumeRoleChkServerPerm,omitempty"`
}

func (s AttachVscToMountPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachVscToMountPointsRequest) GoString() string {
	return s.String()
}

func (s *AttachVscToMountPointsRequest) GetAttachInfos() []*AttachVscToMountPointsRequestAttachInfos {
	return s.AttachInfos
}

func (s *AttachVscToMountPointsRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *AttachVscToMountPointsRequest) GetUseAssumeRoleChkServerPerm() *bool {
	return s.UseAssumeRoleChkServerPerm
}

func (s *AttachVscToMountPointsRequest) SetAttachInfos(v []*AttachVscToMountPointsRequestAttachInfos) *AttachVscToMountPointsRequest {
	s.AttachInfos = v
	return s
}

func (s *AttachVscToMountPointsRequest) SetInputRegionId(v string) *AttachVscToMountPointsRequest {
	s.InputRegionId = &v
	return s
}

func (s *AttachVscToMountPointsRequest) SetUseAssumeRoleChkServerPerm(v bool) *AttachVscToMountPointsRequest {
	s.UseAssumeRoleChkServerPerm = &v
	return s
}

func (s *AttachVscToMountPointsRequest) Validate() error {
	if s.AttachInfos != nil {
		for _, item := range s.AttachInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachVscToMountPointsRequestAttachInfos struct {
	// example:
	//
	// i-2zehyz70ednszl6rrfj6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// f-cfea9ae2ref87.cn-zhangjiakou.dfs.aliyuncs.com
	MountPointId *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	// example:
	//
	// vsc-bp19yqmujug2r762cnabal
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	// example:
	//
	// xc
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	// example:
	//
	// Primary
	VscType *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s AttachVscToMountPointsRequestAttachInfos) String() string {
	return dara.Prettify(s)
}

func (s AttachVscToMountPointsRequestAttachInfos) GoString() string {
	return s.String()
}

func (s *AttachVscToMountPointsRequestAttachInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachVscToMountPointsRequestAttachInfos) GetMountPointId() *string {
	return s.MountPointId
}

func (s *AttachVscToMountPointsRequestAttachInfos) GetVscId() *string {
	return s.VscId
}

func (s *AttachVscToMountPointsRequestAttachInfos) GetVscName() *string {
	return s.VscName
}

func (s *AttachVscToMountPointsRequestAttachInfos) GetVscType() *string {
	return s.VscType
}

func (s *AttachVscToMountPointsRequestAttachInfos) SetInstanceId(v string) *AttachVscToMountPointsRequestAttachInfos {
	s.InstanceId = &v
	return s
}

func (s *AttachVscToMountPointsRequestAttachInfos) SetMountPointId(v string) *AttachVscToMountPointsRequestAttachInfos {
	s.MountPointId = &v
	return s
}

func (s *AttachVscToMountPointsRequestAttachInfos) SetVscId(v string) *AttachVscToMountPointsRequestAttachInfos {
	s.VscId = &v
	return s
}

func (s *AttachVscToMountPointsRequestAttachInfos) SetVscName(v string) *AttachVscToMountPointsRequestAttachInfos {
	s.VscName = &v
	return s
}

func (s *AttachVscToMountPointsRequestAttachInfos) SetVscType(v string) *AttachVscToMountPointsRequestAttachInfos {
	s.VscType = &v
	return s
}

func (s *AttachVscToMountPointsRequestAttachInfos) Validate() error {
	return dara.Validate(s)
}
