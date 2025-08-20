// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLakeStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteLakeStorageRequest
	GetDBClusterId() *string
	SetLakeStorageId(v string) *DeleteLakeStorageRequest
	GetLakeStorageId() *string
	SetRegionId(v string) *DeleteLakeStorageRequest
	GetRegionId() *string
}

type DeleteLakeStorageRequest struct {
	// The ID of the AnalyticDB for MySQL cluster.
	//
	// example:
	//
	// amv-bp*********
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the lake storage.
	//
	// This parameter is required.
	//
	// example:
	//
	// -
	LakeStorageId *string `json:"LakeStorageId,omitempty" xml:"LakeStorageId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteLakeStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLakeStorageRequest) GoString() string {
	return s.String()
}

func (s *DeleteLakeStorageRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteLakeStorageRequest) GetLakeStorageId() *string {
	return s.LakeStorageId
}

func (s *DeleteLakeStorageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLakeStorageRequest) SetDBClusterId(v string) *DeleteLakeStorageRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteLakeStorageRequest) SetLakeStorageId(v string) *DeleteLakeStorageRequest {
	s.LakeStorageId = &v
	return s
}

func (s *DeleteLakeStorageRequest) SetRegionId(v string) *DeleteLakeStorageRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLakeStorageRequest) Validate() error {
	return dara.Validate(s)
}
