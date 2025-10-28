// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScaleOutEcuRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListScaleOutEcuRequest
	GetAppId() *string
	SetClusterId(v string) *ListScaleOutEcuRequest
	GetClusterId() *string
	SetCpu(v int32) *ListScaleOutEcuRequest
	GetCpu() *int32
	SetGroupId(v string) *ListScaleOutEcuRequest
	GetGroupId() *string
	SetInstanceNum(v int32) *ListScaleOutEcuRequest
	GetInstanceNum() *int32
	SetLogicalRegionId(v string) *ListScaleOutEcuRequest
	GetLogicalRegionId() *string
	SetMem(v int32) *ListScaleOutEcuRequest
	GetMem() *int32
}

type ListScaleOutEcuRequest struct {
	// The ID of the application. Specify this parameter if you want to query the available ECUs in the cluster where the application is deployed.
	//
	// >  Specify at least one of the ClusterId and AppId parameters as the query parameter.
	//
	// example:
	//
	// b93024fd-8a9d-4ef7-99f1-5f0d65cc****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the cluster. Specify this parameter if you want to query the available ECUs in the cluster.
	//
	// > Specify at least one of the ClusterId and AppId parameters as the query parameter.
	//
	// example:
	//
	// 52984524-6d48-4bbd-******************
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of CPU cores based on which you want to query the available ECUs in the cluster.
	//
	// example:
	//
	// 1
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The ID of the instance group. Specify this parameter if you want to query the available ECUs in the cluster where the instance group resides.
	//
	// example:
	//
	// 8123db90-880f-486f-****-****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of ECUs that you want to query. If this parameter is not specified, all the ECUs that meet the query conditions are returned.
	//
	// example:
	//
	// 1
	InstanceNum *int32 `json:"InstanceNum,omitempty" xml:"InstanceNum,omitempty"`
	// The ID of the namespace.
	//
	// 	- The ID of a custom namespace is in the `region ID:namespace identifier` format. Example: cn-beijing:test.
	//
	// 	- The ID of the default namespace is in the `region ID` format. Example: cn-beijing.
	//
	// example:
	//
	// cn-beijing:test
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	// The size of available memory based on which you want to query the available ECUs in the cluster. Unit: MB.
	//
	// example:
	//
	// 200
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
}

func (s ListScaleOutEcuRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScaleOutEcuRequest) GoString() string {
	return s.String()
}

func (s *ListScaleOutEcuRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListScaleOutEcuRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListScaleOutEcuRequest) GetCpu() *int32 {
	return s.Cpu
}

func (s *ListScaleOutEcuRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListScaleOutEcuRequest) GetInstanceNum() *int32 {
	return s.InstanceNum
}

func (s *ListScaleOutEcuRequest) GetLogicalRegionId() *string {
	return s.LogicalRegionId
}

func (s *ListScaleOutEcuRequest) GetMem() *int32 {
	return s.Mem
}

func (s *ListScaleOutEcuRequest) SetAppId(v string) *ListScaleOutEcuRequest {
	s.AppId = &v
	return s
}

func (s *ListScaleOutEcuRequest) SetClusterId(v string) *ListScaleOutEcuRequest {
	s.ClusterId = &v
	return s
}

func (s *ListScaleOutEcuRequest) SetCpu(v int32) *ListScaleOutEcuRequest {
	s.Cpu = &v
	return s
}

func (s *ListScaleOutEcuRequest) SetGroupId(v string) *ListScaleOutEcuRequest {
	s.GroupId = &v
	return s
}

func (s *ListScaleOutEcuRequest) SetInstanceNum(v int32) *ListScaleOutEcuRequest {
	s.InstanceNum = &v
	return s
}

func (s *ListScaleOutEcuRequest) SetLogicalRegionId(v string) *ListScaleOutEcuRequest {
	s.LogicalRegionId = &v
	return s
}

func (s *ListScaleOutEcuRequest) SetMem(v int32) *ListScaleOutEcuRequest {
	s.Mem = &v
	return s
}

func (s *ListScaleOutEcuRequest) Validate() error {
	return dara.Validate(s)
}
