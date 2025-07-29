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
	// Specifies whether to obtain the kubeconfig file that is used to connect to the cluster over the internal network. You can obtain the terminal ID by calling one of the following operations:
	//
	// 	- `true`: obtains the kubeconfig file that is used to connect to the master instance over the internal network.
	//
	// 	- `false`: obtains the kubeconfig file that is used to connect to the master instance over the Internet.
	//
	// Default value: `false`
	//
	// example:
	//
	// true
	PrivateIpAddress *bool `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The validity period of the temporary kubeconfig file. Unit: minutes. Valid values: 15 to 4320 (3 days).
	//
	// **
	//
	// **Usage notes*	- If you do not specify this parameter, the system specifies a longer validity period. The validity period is returned in the `expiration` parameter.
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
