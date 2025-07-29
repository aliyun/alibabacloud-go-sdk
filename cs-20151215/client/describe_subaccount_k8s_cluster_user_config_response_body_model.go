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
	// The cluster kubeconfig file. For more information about how to view the kubeconfig file content, see [Configure cluster credentials](https://help.aliyun.com/document_detail/86494.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// apiVersion: v1****
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The expiration date of the kubeconfig file. The value is the UTC time displayed in RFC3339 format.
	//
	// example:
	//
	// 2024-03-10T09:56:17Z
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
