// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAddonMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersion(v string) *DescribeClusterAddonMetadataRequest
	GetVersion() *string
}

type DescribeClusterAddonMetadataRequest struct {
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeClusterAddonMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAddonMetadataRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonMetadataRequest) GetVersion() *string {
	return s.Version
}

func (s *DescribeClusterAddonMetadataRequest) SetVersion(v string) *DescribeClusterAddonMetadataRequest {
	s.Version = &v
	return s
}

func (s *DescribeClusterAddonMetadataRequest) Validate() error {
	return dara.Validate(s)
}
