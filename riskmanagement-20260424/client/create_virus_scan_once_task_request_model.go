// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirusScanOnceTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateVirusScanOnceTaskRequest
	GetInstanceId() *string
	SetIp(v string) *CreateVirusScanOnceTaskRequest
	GetIp() *string
	SetRegionId(v string) *CreateVirusScanOnceTaskRequest
	GetRegionId() *string
}

type CreateVirusScanOnceTaskRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateVirusScanOnceTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVirusScanOnceTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVirusScanOnceTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateVirusScanOnceTaskRequest) GetIp() *string {
	return s.Ip
}

func (s *CreateVirusScanOnceTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVirusScanOnceTaskRequest) SetInstanceId(v string) *CreateVirusScanOnceTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateVirusScanOnceTaskRequest) SetIp(v string) *CreateVirusScanOnceTaskRequest {
	s.Ip = &v
	return s
}

func (s *CreateVirusScanOnceTaskRequest) SetRegionId(v string) *CreateVirusScanOnceTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVirusScanOnceTaskRequest) Validate() error {
	return dara.Validate(s)
}
