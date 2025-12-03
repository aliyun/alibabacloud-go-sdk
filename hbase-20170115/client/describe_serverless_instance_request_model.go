// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerlessInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeServerlessInstanceRequest
	GetClientToken() *string
	SetInstanceId(v string) *DescribeServerlessInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeServerlessInstanceRequest
	GetRegionId() *string
	SetZoneId(v string) *DescribeServerlessInstanceRequest
	GetZoneId() *string
}

type DescribeServerlessInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeServerlessInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerlessInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeServerlessInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeServerlessInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeServerlessInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeServerlessInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeServerlessInstanceRequest) SetClientToken(v string) *DescribeServerlessInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeServerlessInstanceRequest) SetInstanceId(v string) *DescribeServerlessInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeServerlessInstanceRequest) SetRegionId(v string) *DescribeServerlessInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeServerlessInstanceRequest) SetZoneId(v string) *DescribeServerlessInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeServerlessInstanceRequest) Validate() error {
	return dara.Validate(s)
}
