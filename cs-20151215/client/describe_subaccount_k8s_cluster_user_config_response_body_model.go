// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubaccountK8sClusterUserConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *DescribeSubaccountK8sClusterUserConfigResponseBody
	GetConfig() *string
	SetExpiration(v string) *DescribeSubaccountK8sClusterUserConfigResponseBody
	GetExpiration() *string
}

type DescribeSubaccountK8sClusterUserConfigResponseBody struct {
	// The KubeConfig of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// apiVersion: v1\\nclusters:\\n- cluster:\\n    server: https://114.55.xx.xx:6443\\n    certificate-authority-data: LS0tLS****
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The expiration time of the KubeConfig. Format: UTC time in RFC 3339 format.
	//
	// example:
	//
	// 2028-04-09T06:20:47Z
	Expiration *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
}

func (s DescribeSubaccountK8sClusterUserConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubaccountK8sClusterUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubaccountK8sClusterUserConfigResponseBody) GetConfig() *string {
	return s.Config
}

func (s *DescribeSubaccountK8sClusterUserConfigResponseBody) GetExpiration() *string {
	return s.Expiration
}

func (s *DescribeSubaccountK8sClusterUserConfigResponseBody) SetConfig(v string) *DescribeSubaccountK8sClusterUserConfigResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeSubaccountK8sClusterUserConfigResponseBody) SetExpiration(v string) *DescribeSubaccountK8sClusterUserConfigResponseBody {
	s.Expiration = &v
	return s
}

func (s *DescribeSubaccountK8sClusterUserConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
