// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReopenInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReopenInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *ReopenInstanceRequest
	GetRegionId() *string
}

type ReopenInstanceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp91inkw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReopenInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReopenInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReopenInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReopenInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReopenInstanceRequest) SetInstanceId(v string) *ReopenInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReopenInstanceRequest) SetRegionId(v string) *ReopenInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ReopenInstanceRequest) Validate() error {
	return dara.Validate(s)
}
