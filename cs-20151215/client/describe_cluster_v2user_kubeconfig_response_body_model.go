// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterV2UserKubeconfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *DescribeClusterV2UserKubeconfigResponseBody
	GetConfig() *string
}

type DescribeClusterV2UserKubeconfigResponseBody struct {
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
}

func (s DescribeClusterV2UserKubeconfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterV2UserKubeconfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2UserKubeconfigResponseBody) GetConfig() *string {
	return s.Config
}

func (s *DescribeClusterV2UserKubeconfigResponseBody) SetConfig(v string) *DescribeClusterV2UserKubeconfigResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeClusterV2UserKubeconfigResponseBody) Validate() error {
	return dara.Validate(s)
}
