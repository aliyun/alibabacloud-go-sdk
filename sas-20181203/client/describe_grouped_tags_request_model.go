// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupedTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMachineTypes(v string) *DescribeGroupedTagsRequest
	GetMachineTypes() *string
}

type DescribeGroupedTagsRequest struct {
	// The type of the asset to query. If you do not specify this parameter, the tags of all asset types are queried. Valid values:
	//
	// 	- **ecs**: server
	//
	// 	- **cloud_product**: Alibaba Cloud service
	//
	// example:
	//
	// ecs
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
}

func (s DescribeGroupedTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupedTagsRequest) GetMachineTypes() *string {
	return s.MachineTypes
}

func (s *DescribeGroupedTagsRequest) SetMachineTypes(v string) *DescribeGroupedTagsRequest {
	s.MachineTypes = &v
	return s
}

func (s *DescribeGroupedTagsRequest) Validate() error {
	return dara.Validate(s)
}
