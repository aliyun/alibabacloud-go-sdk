// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StartInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *StartInstanceRequest
	GetRegionId() *string
	SetVswitchId(v string) *StartInstanceRequest
	GetVswitchId() *string
}

type StartInstanceRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VswitchId  *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartInstanceRequest) GetVswitchId() *string {
	return s.VswitchId
}

func (s *StartInstanceRequest) SetInstanceId(v string) *StartInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceRequest) SetRegionId(v string) *StartInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartInstanceRequest) SetVswitchId(v string) *StartInstanceRequest {
	s.VswitchId = &v
	return s
}

func (s *StartInstanceRequest) Validate() error {
	return dara.Validate(s)
}
