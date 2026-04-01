// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRCInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StartRCInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *StartRCInstanceRequest
	GetRegionId() *string
}

type StartRCInstanceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rc-l02u59b2kjfd2us0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartRCInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRCInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartRCInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartRCInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartRCInstanceRequest) SetInstanceId(v string) *StartRCInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartRCInstanceRequest) SetRegionId(v string) *StartRCInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartRCInstanceRequest) Validate() error {
	return dara.Validate(s)
}
