// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCClusterConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *DescribeRCClusterConfigResponseBody
	GetConfig() *string
	SetExpiration(v string) *DescribeRCClusterConfigResponseBody
	GetExpiration() *string
	SetRequestId(v string) *DescribeRCClusterConfigResponseBody
	GetRequestId() *string
}

type DescribeRCClusterConfigResponseBody struct {
	// The kubeconfig file of the cluster.
	//
	// example:
	//
	// apiVersion: v1****
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The expiration time of the kubeconfig file. Format: the UTC time in the RFC3339 format.
	//
	// example:
	//
	// 2024-03-10T09:56:17Z
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E9DD55F4-1A5F-48CA-BA57-DFB3CA8C4C34
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRCClusterConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCClusterConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCClusterConfigResponseBody) GetConfig() *string {
	return s.Config
}

func (s *DescribeRCClusterConfigResponseBody) GetExpiration() *string {
	return s.Expiration
}

func (s *DescribeRCClusterConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCClusterConfigResponseBody) SetConfig(v string) *DescribeRCClusterConfigResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeRCClusterConfigResponseBody) SetExpiration(v string) *DescribeRCClusterConfigResponseBody {
	s.Expiration = &v
	return s
}

func (s *DescribeRCClusterConfigResponseBody) SetRequestId(v string) *DescribeRCClusterConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCClusterConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
