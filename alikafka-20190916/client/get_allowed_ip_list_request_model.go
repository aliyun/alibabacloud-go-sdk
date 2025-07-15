// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllowedIpListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetAllowedIpListRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetAllowedIpListRequest
	GetRegionId() *string
}

type GetAllowedIpListRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp91inkw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAllowedIpListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAllowedIpListRequest) GoString() string {
	return s.String()
}

func (s *GetAllowedIpListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAllowedIpListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAllowedIpListRequest) SetInstanceId(v string) *GetAllowedIpListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAllowedIpListRequest) SetRegionId(v string) *GetAllowedIpListRequest {
	s.RegionId = &v
	return s
}

func (s *GetAllowedIpListRequest) Validate() error {
	return dara.Validate(s)
}
