// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAINodePoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListAINodePoolsRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *ListAINodePoolsRequest
	GetRegionId() *string
}

type ListAINodePoolsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAINodePoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAINodePoolsRequest) GoString() string {
	return s.String()
}

func (s *ListAINodePoolsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListAINodePoolsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAINodePoolsRequest) SetDBInstanceId(v string) *ListAINodePoolsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListAINodePoolsRequest) SetRegionId(v string) *ListAINodePoolsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAINodePoolsRequest) Validate() error {
	return dara.Validate(s)
}
