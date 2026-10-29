// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterUserKubeconfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivateIpAddress(v bool) *DescribeClusterUserKubeconfigRequest
	GetPrivateIpAddress() *bool
	SetTemporaryDurationMinutes(v int64) *DescribeClusterUserKubeconfigRequest
	GetTemporaryDurationMinutes() *int64
}

type DescribeClusterUserKubeconfigRequest struct {
	// Specifies whether to obtain the internal network connection configuration. Valid values:
	//
	// - `true`: obtains only the internal network connection credentials.
	//
	// - `false`: obtains only the public network connection credentials.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	PrivateIpAddress *bool `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The validity period of the temporary kubeconfig file, in minutes. Valid values: 15 to 4320.
	//
	// > If you do not specify this parameter, the system automatically determines a longer validity period. You can determine the specific expiration time based on the value of the `expiration` field in the response.
	//
	// example:
	//
	// 15
	TemporaryDurationMinutes *int64 `json:"TemporaryDurationMinutes,omitempty" xml:"TemporaryDurationMinutes,omitempty"`
}

func (s DescribeClusterUserKubeconfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterUserKubeconfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterUserKubeconfigRequest) GetPrivateIpAddress() *bool {
	return s.PrivateIpAddress
}

func (s *DescribeClusterUserKubeconfigRequest) GetTemporaryDurationMinutes() *int64 {
	return s.TemporaryDurationMinutes
}

func (s *DescribeClusterUserKubeconfigRequest) SetPrivateIpAddress(v bool) *DescribeClusterUserKubeconfigRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeClusterUserKubeconfigRequest) SetTemporaryDurationMinutes(v int64) *DescribeClusterUserKubeconfigRequest {
	s.TemporaryDurationMinutes = &v
	return s
}

func (s *DescribeClusterUserKubeconfigRequest) Validate() error {
	return dara.Validate(s)
}
