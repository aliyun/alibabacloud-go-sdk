// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetInstanceRequest
	GetRegionId() *string
	SetInstanceId(v string) *GetInstanceRequest
	GetInstanceId() *string
}

type GetInstanceRequest struct {
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c-xxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s GetInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceRequest) SetRegionId(v string) *GetInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetInstanceRequest) SetInstanceId(v string) *GetInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceRequest) Validate() error {
	return dara.Validate(s)
}
