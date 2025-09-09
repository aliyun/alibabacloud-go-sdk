// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceStoreInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceStoreInfoRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetInstanceStoreInfoRequest
	GetRegionId() *string
}

type GetInstanceStoreInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetInstanceStoreInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceStoreInfoRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceStoreInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceStoreInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceStoreInfoRequest) SetInstanceId(v string) *GetInstanceStoreInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceStoreInfoRequest) SetRegionId(v string) *GetInstanceStoreInfoRequest {
	s.RegionId = &v
	return s
}

func (s *GetInstanceStoreInfoRequest) Validate() error {
	return dara.Validate(s)
}
