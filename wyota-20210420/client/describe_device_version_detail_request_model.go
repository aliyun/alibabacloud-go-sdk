// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceVersionDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModel(v string) *DescribeDeviceVersionDetailRequest
	GetModel() *string
	SetNetworkType(v string) *DescribeDeviceVersionDetailRequest
	GetNetworkType() *string
	SetRegion(v string) *DescribeDeviceVersionDetailRequest
	GetRegion() *string
	SetVersionName(v string) *DescribeDeviceVersionDetailRequest
	GetVersionName() *string
}

type DescribeDeviceVersionDetailRequest struct {
	Model       *string `json:"Model,omitempty" xml:"Model,omitempty"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s DescribeDeviceVersionDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceVersionDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceVersionDetailRequest) GetModel() *string {
	return s.Model
}

func (s *DescribeDeviceVersionDetailRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDeviceVersionDetailRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeDeviceVersionDetailRequest) GetVersionName() *string {
	return s.VersionName
}

func (s *DescribeDeviceVersionDetailRequest) SetModel(v string) *DescribeDeviceVersionDetailRequest {
	s.Model = &v
	return s
}

func (s *DescribeDeviceVersionDetailRequest) SetNetworkType(v string) *DescribeDeviceVersionDetailRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeDeviceVersionDetailRequest) SetRegion(v string) *DescribeDeviceVersionDetailRequest {
	s.Region = &v
	return s
}

func (s *DescribeDeviceVersionDetailRequest) SetVersionName(v string) *DescribeDeviceVersionDetailRequest {
	s.VersionName = &v
	return s
}

func (s *DescribeDeviceVersionDetailRequest) Validate() error {
	return dara.Validate(s)
}
