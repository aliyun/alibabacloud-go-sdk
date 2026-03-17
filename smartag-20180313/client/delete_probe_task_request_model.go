// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProbeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProbeTaskId(v string) *DeleteProbeTaskRequest
	GetProbeTaskId() *string
	SetRegionId(v string) *DeleteProbeTaskRequest
	GetRegionId() *string
	SetSagId(v string) *DeleteProbeTaskRequest
	GetSagId() *string
	SetSn(v string) *DeleteProbeTaskRequest
	GetSn() *string
}

type DeleteProbeTaskRequest struct {
	// The ID of the probe task.
	//
	// This parameter is required.
	//
	// example:
	//
	// probe-****
	ProbeTaskId *string `json:"ProbeTaskId,omitempty" xml:"ProbeTaskId,omitempty"`
	// The region ID of the Smart Access Gateway (SAG) instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-****
	SagId *string `json:"SagId,omitempty" xml:"SagId,omitempty"`
	// The serial number of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag****
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s DeleteProbeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProbeTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteProbeTaskRequest) GetProbeTaskId() *string {
	return s.ProbeTaskId
}

func (s *DeleteProbeTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteProbeTaskRequest) GetSagId() *string {
	return s.SagId
}

func (s *DeleteProbeTaskRequest) GetSn() *string {
	return s.Sn
}

func (s *DeleteProbeTaskRequest) SetProbeTaskId(v string) *DeleteProbeTaskRequest {
	s.ProbeTaskId = &v
	return s
}

func (s *DeleteProbeTaskRequest) SetRegionId(v string) *DeleteProbeTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteProbeTaskRequest) SetSagId(v string) *DeleteProbeTaskRequest {
	s.SagId = &v
	return s
}

func (s *DeleteProbeTaskRequest) SetSn(v string) *DeleteProbeTaskRequest {
	s.Sn = &v
	return s
}

func (s *DeleteProbeTaskRequest) Validate() error {
	return dara.Validate(s)
}
