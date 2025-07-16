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
	CurrentReaon     *string `json:"CurrentReaon,omitempty" xml:"CurrentReaon,omitempty"`
	CurrentStatus    *string `json:"CurrentStatus,omitempty" xml:"CurrentStatus,omitempty"`
	CurrentTimestamp *string `json:"CurrentTimestamp,omitempty" xml:"CurrentTimestamp,omitempty"`
	Image            *string `json:"Image,omitempty" xml:"Image,omitempty"`
	LastReason       *string `json:"LastReason,omitempty" xml:"LastReason,omitempty"`
	LastStatus       *string `json:"LastStatus,omitempty" xml:"LastStatus,omitempty"`
	LastTimestamp    *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Ready            *bool   `json:"Ready,omitempty" xml:"Ready,omitempty"`
	RestartCount     *int32  `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
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
