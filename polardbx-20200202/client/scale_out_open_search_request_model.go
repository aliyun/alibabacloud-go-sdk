// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleOutOpenSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ScaleOutOpenSearchRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *ScaleOutOpenSearchRequest
	GetRegionId() *string
	SetSearchNodeCount(v string) *ScaleOutOpenSearchRequest
	GetSearchNodeCount() *string
}

type ScaleOutOpenSearchRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-spsil01pww4hfz
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The total number of data nodes after the scale-out. The value must be a positive integer and greater than the current number of data nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	SearchNodeCount *string `json:"SearchNodeCount,omitempty" xml:"SearchNodeCount,omitempty"`
}

func (s ScaleOutOpenSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleOutOpenSearchRequest) GoString() string {
	return s.String()
}

func (s *ScaleOutOpenSearchRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ScaleOutOpenSearchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ScaleOutOpenSearchRequest) GetSearchNodeCount() *string {
	return s.SearchNodeCount
}

func (s *ScaleOutOpenSearchRequest) SetDBInstanceName(v string) *ScaleOutOpenSearchRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ScaleOutOpenSearchRequest) SetRegionId(v string) *ScaleOutOpenSearchRequest {
	s.RegionId = &v
	return s
}

func (s *ScaleOutOpenSearchRequest) SetSearchNodeCount(v string) *ScaleOutOpenSearchRequest {
	s.SearchNodeCount = &v
	return s
}

func (s *ScaleOutOpenSearchRequest) Validate() error {
	return dara.Validate(s)
}
