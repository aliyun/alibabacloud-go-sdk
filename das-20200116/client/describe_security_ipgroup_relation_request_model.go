// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityIPGroupRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeSecurityIPGroupRelationRequest
	GetInstanceId() *string
	SetRegionName(v string) *DescribeSecurityIPGroupRelationRequest
	GetRegionName() *string
}

type DescribeSecurityIPGroupRelationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rm-2zegsc57ofexxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeSecurityIPGroupRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPGroupRelationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSecurityIPGroupRelationRequest) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeSecurityIPGroupRelationRequest) SetInstanceId(v string) *DescribeSecurityIPGroupRelationRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSecurityIPGroupRelationRequest) SetRegionName(v string) *DescribeSecurityIPGroupRelationRequest {
	s.RegionName = &v
	return s
}

func (s *DescribeSecurityIPGroupRelationRequest) Validate() error {
	return dara.Validate(s)
}
