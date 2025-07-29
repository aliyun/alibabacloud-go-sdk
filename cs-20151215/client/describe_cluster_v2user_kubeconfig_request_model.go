// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterV2UserKubeconfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivateIpAddress(v bool) *DescribeClusterV2UserKubeconfigRequest
	GetPrivateIpAddress() *bool
	SetTemporaryDurationMinutes(v int64) *DescribeClusterV2UserKubeconfigRequest
	GetTemporaryDurationMinutes() *int64
}

type DescribeClusterV2UserKubeconfigRequest struct {
	PrivateIpAddress         *bool  `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	TemporaryDurationMinutes *int64 `json:"TemporaryDurationMinutes,omitempty" xml:"TemporaryDurationMinutes,omitempty"`
}

func (s DescribeClusterV2UserKubeconfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterV2UserKubeconfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2UserKubeconfigRequest) GetPrivateIpAddress() *bool {
	return s.PrivateIpAddress
}

func (s *DescribeClusterV2UserKubeconfigRequest) GetTemporaryDurationMinutes() *int64 {
	return s.TemporaryDurationMinutes
}

func (s *DescribeClusterV2UserKubeconfigRequest) SetPrivateIpAddress(v bool) *DescribeClusterV2UserKubeconfigRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeClusterV2UserKubeconfigRequest) SetTemporaryDurationMinutes(v int64) *DescribeClusterV2UserKubeconfigRequest {
	s.TemporaryDurationMinutes = &v
	return s
}

func (s *DescribeClusterV2UserKubeconfigRequest) Validate() error {
	return dara.Validate(s)
}
