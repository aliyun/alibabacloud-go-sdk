// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespaceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainCustom(v bool) *DescribeNamespaceListRequest
	GetContainCustom() *bool
	SetHybridCloudExclude(v bool) *DescribeNamespaceListRequest
	GetHybridCloudExclude() *bool
}

type DescribeNamespaceListRequest struct {
	// Specifies whether to return custom namespaces. Valid values:
	//
	// 	- **true**: The system returns custom namespaces.
	//
	// 	- **false**: The system does not return custom namespaces.
	//
	// example:
	//
	// true
	ContainCustom *bool `json:"ContainCustom,omitempty" xml:"ContainCustom,omitempty"`
	// Indicates whether hybrid cloud namespaces are excluded. Valid values:
	//
	// 	- **true**: Hybrid cloud namespaces are excluded.
	//
	// 	- **false**: Hybrid cloud namespaces are included.
	//
	// example:
	//
	// true
	HybridCloudExclude *bool `json:"HybridCloudExclude,omitempty" xml:"HybridCloudExclude,omitempty"`
}

func (s DescribeNamespaceListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceListRequest) GetContainCustom() *bool {
	return s.ContainCustom
}

func (s *DescribeNamespaceListRequest) GetHybridCloudExclude() *bool {
	return s.HybridCloudExclude
}

func (s *DescribeNamespaceListRequest) SetContainCustom(v bool) *DescribeNamespaceListRequest {
	s.ContainCustom = &v
	return s
}

func (s *DescribeNamespaceListRequest) SetHybridCloudExclude(v bool) *DescribeNamespaceListRequest {
	s.HybridCloudExclude = &v
	return s
}

func (s *DescribeNamespaceListRequest) Validate() error {
	return dara.Validate(s)
}
