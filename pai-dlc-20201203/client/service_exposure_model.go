// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceExposure interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *ServiceExposure
	GetDisplayName() *string
	SetJobId(v string) *ServiceExposure
	GetJobId() *string
	SetMessage(v string) *ServiceExposure
	GetMessage() *string
	SetPodId(v string) *ServiceExposure
	GetPodId() *string
	SetPort(v int32) *ServiceExposure
	GetPort() *int32
	SetServiceId(v string) *ServiceExposure
	GetServiceId() *string
	SetStatus(v string) *ServiceExposure
	GetStatus() *string
	SetType(v string) *ServiceExposure
	GetType() *string
}

type ServiceExposure struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	JobId       *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PodId       *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	Port        *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ServiceId   *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ServiceExposure) String() string {
	return dara.Prettify(s)
}

func (s ServiceExposure) GoString() string {
	return s.String()
}

func (s *ServiceExposure) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ServiceExposure) GetJobId() *string {
	return s.JobId
}

func (s *ServiceExposure) GetMessage() *string {
	return s.Message
}

func (s *ServiceExposure) GetPodId() *string {
	return s.PodId
}

func (s *ServiceExposure) GetPort() *int32 {
	return s.Port
}

func (s *ServiceExposure) GetServiceId() *string {
	return s.ServiceId
}

func (s *ServiceExposure) GetStatus() *string {
	return s.Status
}

func (s *ServiceExposure) GetType() *string {
	return s.Type
}

func (s *ServiceExposure) SetDisplayName(v string) *ServiceExposure {
	s.DisplayName = &v
	return s
}

func (s *ServiceExposure) SetJobId(v string) *ServiceExposure {
	s.JobId = &v
	return s
}

func (s *ServiceExposure) SetMessage(v string) *ServiceExposure {
	s.Message = &v
	return s
}

func (s *ServiceExposure) SetPodId(v string) *ServiceExposure {
	s.PodId = &v
	return s
}

func (s *ServiceExposure) SetPort(v int32) *ServiceExposure {
	s.Port = &v
	return s
}

func (s *ServiceExposure) SetServiceId(v string) *ServiceExposure {
	s.ServiceId = &v
	return s
}

func (s *ServiceExposure) SetStatus(v string) *ServiceExposure {
	s.Status = &v
	return s
}

func (s *ServiceExposure) SetType(v string) *ServiceExposure {
	s.Type = &v
	return s
}

func (s *ServiceExposure) Validate() error {
	return dara.Validate(s)
}
