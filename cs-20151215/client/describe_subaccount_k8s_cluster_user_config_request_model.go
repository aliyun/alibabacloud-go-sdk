// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubaccountK8sClusterUserConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivateIpAddress(v bool) *DescribeSubaccountK8sClusterUserConfigRequest
	GetPrivateIpAddress() *bool
	SetTemporaryDurationMinutes(v int64) *DescribeSubaccountK8sClusterUserConfigRequest
	GetTemporaryDurationMinutes() *int64
}

type DescribeSubaccountK8sClusterUserConfigRequest struct {
	// Specifies whether to obtain the kubeconfig file used to connect to the cluster over the internal network. Valid values:
	//
	// 	- `true`: Obtain the kubeconfig file used to connect to the cluster over the internal network.
	//
	// 	- `false`: Obtain the kubeconfig file used to connect to the cluster over the Internet.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	PrivateIpAddress *bool `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The validity period of the temporary kubeconfig file. Unit: minutes.
	//
	// Valid values: 15 to 4320 (three days).
	//
	// > If you leave this parameter empty, the system sets a longer validity period and returns the value in the expiration parameter of the response.
	//
	// example:
	//
	// 15
	TemporaryDurationMinutes *int64 `json:"TemporaryDurationMinutes,omitempty" xml:"TemporaryDurationMinutes,omitempty"`
}

func (s DescribeSubaccountK8sClusterUserConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubaccountK8sClusterUserConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubaccountK8sClusterUserConfigRequest) GetPrivateIpAddress() *bool {
	return s.PrivateIpAddress
}

func (s *DescribeSubaccountK8sClusterUserConfigRequest) GetTemporaryDurationMinutes() *int64 {
	return s.TemporaryDurationMinutes
}

func (s *DescribeSubaccountK8sClusterUserConfigRequest) SetPrivateIpAddress(v bool) *DescribeSubaccountK8sClusterUserConfigRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeSubaccountK8sClusterUserConfigRequest) SetTemporaryDurationMinutes(v int64) *DescribeSubaccountK8sClusterUserConfigRequest {
	s.TemporaryDurationMinutes = &v
	return s
}

func (s *DescribeSubaccountK8sClusterUserConfigRequest) Validate() error {
	return dara.Validate(s)
}
