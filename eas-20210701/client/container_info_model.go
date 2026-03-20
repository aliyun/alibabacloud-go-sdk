// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContainerInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentReaon(v string) *ContainerInfo
	GetCurrentReaon() *string
	SetCurrentStatus(v string) *ContainerInfo
	GetCurrentStatus() *string
	SetCurrentTimestamp(v string) *ContainerInfo
	GetCurrentTimestamp() *string
	SetImage(v string) *ContainerInfo
	GetImage() *string
	SetLastReason(v string) *ContainerInfo
	GetLastReason() *string
	SetLastStatus(v string) *ContainerInfo
	GetLastStatus() *string
	SetLastTimestamp(v string) *ContainerInfo
	GetLastTimestamp() *string
	SetName(v string) *ContainerInfo
	GetName() *string
	SetPort(v int32) *ContainerInfo
	GetPort() *int32
	SetReady(v bool) *ContainerInfo
	GetReady() *bool
	SetRestartCount(v int32) *ContainerInfo
	GetRestartCount() *int32
}

type ContainerInfo struct {
	// The reason why the container is in the current state.
	//
	// example:
	//
	// PodInitializing
	CurrentReaon *string `json:"CurrentReaon,omitempty" xml:"CurrentReaon,omitempty"`
	// The current state of the container. Valid values:
	//
	// 	- Waiting
	//
	// 	- Running
	//
	// 	- Terminated
	//
	// example:
	//
	// Waiting
	CurrentStatus *string `json:"CurrentStatus,omitempty" xml:"CurrentStatus,omitempty"`
	// The time when the container entered the current state.
	//
	// example:
	//
	// 2022-03-21T06:17:57Z
	CurrentTimestamp *string `json:"CurrentTimestamp,omitempty" xml:"CurrentTimestamp,omitempty"`
	// The image.
	//
	// example:
	//
	// registry.cn-shanghai.aliyuncs.com/eas/echo_cn-shanghai:v0.0.1-20210129111320
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The reason why the container is in the last state.
	//
	// example:
	//
	// Error
	LastReason *string `json:"LastReason,omitempty" xml:"LastReason,omitempty"`
	// The last state of the container. Valid values:
	//
	// 	- Waiting
	//
	// 	- Running
	//
	// 	- Terminated
	//
	// example:
	//
	// Terminated
	LastStatus *string `json:"LastStatus,omitempty" xml:"LastStatus,omitempty"`
	// The time when the container entered the last state.
	//
	// example:
	//
	// 2022-03-21T05:17:57Z
	LastTimestamp *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The container name.
	//
	// example:
	//
	// worker0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port number.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// Indicates whether the container passed the health check.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Ready *bool `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// The number of times the container restarted.
	//
	// example:
	//
	// 0
	RestartCount *int32 `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
}

func (s ContainerInfo) String() string {
	return dara.Prettify(s)
}

func (s ContainerInfo) GoString() string {
	return s.String()
}

func (s *ContainerInfo) GetCurrentReaon() *string {
	return s.CurrentReaon
}

func (s *ContainerInfo) GetCurrentStatus() *string {
	return s.CurrentStatus
}

func (s *ContainerInfo) GetCurrentTimestamp() *string {
	return s.CurrentTimestamp
}

func (s *ContainerInfo) GetImage() *string {
	return s.Image
}

func (s *ContainerInfo) GetLastReason() *string {
	return s.LastReason
}

func (s *ContainerInfo) GetLastStatus() *string {
	return s.LastStatus
}

func (s *ContainerInfo) GetLastTimestamp() *string {
	return s.LastTimestamp
}

func (s *ContainerInfo) GetName() *string {
	return s.Name
}

func (s *ContainerInfo) GetPort() *int32 {
	return s.Port
}

func (s *ContainerInfo) GetReady() *bool {
	return s.Ready
}

func (s *ContainerInfo) GetRestartCount() *int32 {
	return s.RestartCount
}

func (s *ContainerInfo) SetCurrentReaon(v string) *ContainerInfo {
	s.CurrentReaon = &v
	return s
}

func (s *ContainerInfo) SetCurrentStatus(v string) *ContainerInfo {
	s.CurrentStatus = &v
	return s
}

func (s *ContainerInfo) SetCurrentTimestamp(v string) *ContainerInfo {
	s.CurrentTimestamp = &v
	return s
}

func (s *ContainerInfo) SetImage(v string) *ContainerInfo {
	s.Image = &v
	return s
}

func (s *ContainerInfo) SetLastReason(v string) *ContainerInfo {
	s.LastReason = &v
	return s
}

func (s *ContainerInfo) SetLastStatus(v string) *ContainerInfo {
	s.LastStatus = &v
	return s
}

func (s *ContainerInfo) SetLastTimestamp(v string) *ContainerInfo {
	s.LastTimestamp = &v
	return s
}

func (s *ContainerInfo) SetName(v string) *ContainerInfo {
	s.Name = &v
	return s
}

func (s *ContainerInfo) SetPort(v int32) *ContainerInfo {
	s.Port = &v
	return s
}

func (s *ContainerInfo) SetReady(v bool) *ContainerInfo {
	s.Ready = &v
	return s
}

func (s *ContainerInfo) SetRestartCount(v int32) *ContainerInfo {
	s.RestartCount = &v
	return s
}

func (s *ContainerInfo) Validate() error {
	return dara.Validate(s)
}
