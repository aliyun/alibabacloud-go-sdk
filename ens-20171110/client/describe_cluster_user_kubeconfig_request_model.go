// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterUserKubeconfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterUserKubeconfigRequest
	GetClusterId() *string
	SetPrivateIpAddress(v bool) *DescribeClusterUserKubeconfigRequest
	GetPrivateIpAddress() *bool
}

type DescribeClusterUserKubeconfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxxx
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PrivateIpAddress *bool   `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s DescribeClusterUserKubeconfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterUserKubeconfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterUserKubeconfigRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterUserKubeconfigRequest) GetPrivateIpAddress() *bool {
	return s.PrivateIpAddress
}

func (s *DescribeClusterUserKubeconfigRequest) SetClusterId(v string) *DescribeClusterUserKubeconfigRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterUserKubeconfigRequest) SetPrivateIpAddress(v bool) *DescribeClusterUserKubeconfigRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeClusterUserKubeconfigRequest) Validate() error {
	return dara.Validate(s)
}
