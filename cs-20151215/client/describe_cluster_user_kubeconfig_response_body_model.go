// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterUserKubeconfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *DescribeClusterUserKubeconfigResponseBody
	GetConfig() *string
	SetExpiration(v string) *DescribeClusterUserKubeconfigResponseBody
	GetExpiration() *string
}

type DescribeClusterUserKubeconfigResponseBody struct {
	// The kubeconfig file of the cluster.
	//
	// example:
	//
	// apiVersion: v1****
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The expiration time of the kubeconfig file. Format: the UTC time in the RFC3339 format.
	//
	// example:
	//
	// 2024-03-10T09:56:17Z
	Expiration *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
}

func (s DescribeClusterUserKubeconfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterUserKubeconfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterUserKubeconfigResponseBody) GetConfig() *string {
	return s.Config
}

func (s *DescribeClusterUserKubeconfigResponseBody) GetExpiration() *string {
	return s.Expiration
}

func (s *DescribeClusterUserKubeconfigResponseBody) SetConfig(v string) *DescribeClusterUserKubeconfigResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeClusterUserKubeconfigResponseBody) SetExpiration(v string) *DescribeClusterUserKubeconfigResponseBody {
	s.Expiration = &v
	return s
}

func (s *DescribeClusterUserKubeconfigResponseBody) Validate() error {
	return dara.Validate(s)
}
