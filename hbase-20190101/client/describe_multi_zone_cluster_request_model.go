// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiZoneClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeMultiZoneClusterRequest
	GetClusterId() *string
}

type DescribeMultiZoneClusterRequest struct {
	// The ID of the multi-zone instance. You can call [DescribeInstances](https://help.aliyun.com/document_detail/144595.html) to obtain the list. The multi-zone instance has **DbType*	- set to hbaseue and **ModuleStackVersion*	- set to 2.0.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-t4nn71xa0yn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeMultiZoneClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeMultiZoneClusterRequest) SetClusterId(v string) *DescribeMultiZoneClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeMultiZoneClusterRequest) Validate() error {
	return dara.Validate(s)
}
