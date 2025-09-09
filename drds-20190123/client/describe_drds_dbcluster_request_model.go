// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstanceId(v string) *DescribeDrdsDBClusterRequest
	GetDbInstanceId() *string
	SetDbName(v string) *DescribeDrdsDBClusterRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeDrdsDBClusterRequest
	GetDrdsInstanceId() *string
}

type DescribeDrdsDBClusterRequest struct {
	// The ID of the PolarDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*****************
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The name of the DRDS database.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of a DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds*********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DescribeDrdsDBClusterRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDrdsDBClusterRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsDBClusterRequest) SetDbInstanceId(v string) *DescribeDrdsDBClusterRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDrdsDBClusterRequest) SetDbName(v string) *DescribeDrdsDBClusterRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBClusterRequest) SetDrdsInstanceId(v string) *DescribeDrdsDBClusterRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
