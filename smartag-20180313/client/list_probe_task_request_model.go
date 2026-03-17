// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProbeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListProbeTaskRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProbeTaskRequest
	GetPageSize() *int32
	SetProbeTaskId(v string) *ListProbeTaskRequest
	GetProbeTaskId() *string
	SetProtocol(v string) *ListProbeTaskRequest
	GetProtocol() *string
	SetRegionId(v string) *ListProbeTaskRequest
	GetRegionId() *string
	SetSagId(v string) *ListProbeTaskRequest
	GetSagId() *string
	SetSagName(v string) *ListProbeTaskRequest
	GetSagName() *string
	SetSn(v string) *ListProbeTaskRequest
	GetSn() *string
	SetTaskName(v string) *ListProbeTaskRequest
	GetTaskName() *string
}

type ListProbeTaskRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the probe task.
	//
	// example:
	//
	// probe-****
	ProbeTaskId *string `json:"ProbeTaskId,omitempty" xml:"ProbeTaskId,omitempty"`
	// The protocol of the probe task. Valid values:
	//
	// 	- **ICMP**
	//
	// 	- **TCP**
	//
	// 	- **HTTP**
	//
	// >  Tasks that probe private networks support only ICMP and TCP.
	//
	// example:
	//
	// ICMP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region ID of the SAG instance.
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
	// example:
	//
	// sag-****
	SagId *string `json:"SagId,omitempty" xml:"SagId,omitempty"`
	// The name of the SAG instance.
	//
	// example:
	//
	// shanghai-office
	SagName *string `json:"SagName,omitempty" xml:"SagName,omitempty"`
	// The serial number of the SAG device.
	//
	// example:
	//
	// sag****
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// The name of the probe task.
	//
	// example:
	//
	// test-ping
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListProbeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProbeTaskRequest) GoString() string {
	return s.String()
}

func (s *ListProbeTaskRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProbeTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProbeTaskRequest) GetProbeTaskId() *string {
	return s.ProbeTaskId
}

func (s *ListProbeTaskRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *ListProbeTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListProbeTaskRequest) GetSagId() *string {
	return s.SagId
}

func (s *ListProbeTaskRequest) GetSagName() *string {
	return s.SagName
}

func (s *ListProbeTaskRequest) GetSn() *string {
	return s.Sn
}

func (s *ListProbeTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ListProbeTaskRequest) SetPageNumber(v int32) *ListProbeTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProbeTaskRequest) SetPageSize(v int32) *ListProbeTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListProbeTaskRequest) SetProbeTaskId(v string) *ListProbeTaskRequest {
	s.ProbeTaskId = &v
	return s
}

func (s *ListProbeTaskRequest) SetProtocol(v string) *ListProbeTaskRequest {
	s.Protocol = &v
	return s
}

func (s *ListProbeTaskRequest) SetRegionId(v string) *ListProbeTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ListProbeTaskRequest) SetSagId(v string) *ListProbeTaskRequest {
	s.SagId = &v
	return s
}

func (s *ListProbeTaskRequest) SetSagName(v string) *ListProbeTaskRequest {
	s.SagName = &v
	return s
}

func (s *ListProbeTaskRequest) SetSn(v string) *ListProbeTaskRequest {
	s.Sn = &v
	return s
}

func (s *ListProbeTaskRequest) SetTaskName(v string) *ListProbeTaskRequest {
	s.TaskName = &v
	return s
}

func (s *ListProbeTaskRequest) Validate() error {
	return dara.Validate(s)
}
