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
	// Specifies whether to obtain the internal network connection configuration. Valid values:
	//
	// - `true`: obtains only the KubeConfig credentials for internal network connections
	//
	// - `false`: obtains only the KubeConfig credentials for public network connections
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	PrivateIpAddress *bool `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The validity period of the temporary KubeConfig, in minutes.
	//
	// Valid values: 15 to 4320, which indicates a maximum of 3 days.
	//
	// > If you do not specify this parameter, the system automatically determines a longer validity period. The specific expiration time is determined by the value of the expiration field in the response.
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
