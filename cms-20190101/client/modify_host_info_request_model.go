// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostName(v string) *ModifyHostInfoRequest
	GetHostName() *string
	SetInstanceId(v string) *ModifyHostInfoRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyHostInfoRequest
	GetRegionId() *string
}

type ModifyHostInfoRequest struct {
	// The host name.
	//
	// example:
	//
	// portalHost
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The instance ID. Only non-Alibaba Cloud hosts are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// host-R_NSWNV****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyHostInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostInfoRequest) GetHostName() *string {
	return s.HostName
}

func (s *ModifyHostInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHostInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHostInfoRequest) SetHostName(v string) *ModifyHostInfoRequest {
	s.HostName = &v
	return s
}

func (s *ModifyHostInfoRequest) SetInstanceId(v string) *ModifyHostInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostInfoRequest) SetRegionId(v string) *ModifyHostInfoRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHostInfoRequest) Validate() error {
	return dara.Validate(s)
}
