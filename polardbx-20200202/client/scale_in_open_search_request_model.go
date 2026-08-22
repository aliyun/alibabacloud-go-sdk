// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleInOpenSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ScaleInOpenSearchRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *ScaleInOpenSearchRequest
	GetRegionId() *string
	SetSearchNodeCount(v string) *ScaleInOpenSearchRequest
	GetSearchNodeCount() *string
}

type ScaleInOpenSearchRequest struct {
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
	// The total number of data nodes after the scale-in. The value must be a positive integer and less than the current number of data nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	SearchNodeCount *string `json:"SearchNodeCount,omitempty" xml:"SearchNodeCount,omitempty"`
}

func (s ScaleInOpenSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleInOpenSearchRequest) GoString() string {
	return s.String()
}

func (s *ScaleInOpenSearchRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ScaleInOpenSearchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ScaleInOpenSearchRequest) GetSearchNodeCount() *string {
	return s.SearchNodeCount
}

func (s *ScaleInOpenSearchRequest) SetDBInstanceName(v string) *ScaleInOpenSearchRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ScaleInOpenSearchRequest) SetRegionId(v string) *ScaleInOpenSearchRequest {
	s.RegionId = &v
	return s
}

func (s *ScaleInOpenSearchRequest) SetSearchNodeCount(v string) *ScaleInOpenSearchRequest {
	s.SearchNodeCount = &v
	return s
}

func (s *ScaleInOpenSearchRequest) Validate() error {
	return dara.Validate(s)
}
