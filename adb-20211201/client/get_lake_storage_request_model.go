// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLakeStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetLakeStorageRequest
	GetDBClusterId() *string
	SetLakeStorageId(v string) *GetLakeStorageRequest
	GetLakeStorageId() *string
	SetRegionId(v string) *GetLakeStorageRequest
	GetRegionId() *string
}

type GetLakeStorageRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// amv-*******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The lake storage ID.
	//
	// example:
	//
	// -
	LakeStorageId *string `json:"LakeStorageId,omitempty" xml:"LakeStorageId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetLakeStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLakeStorageRequest) GoString() string {
	return s.String()
}

func (s *GetLakeStorageRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetLakeStorageRequest) GetLakeStorageId() *string {
	return s.LakeStorageId
}

func (s *GetLakeStorageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLakeStorageRequest) SetDBClusterId(v string) *GetLakeStorageRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetLakeStorageRequest) SetLakeStorageId(v string) *GetLakeStorageRequest {
	s.LakeStorageId = &v
	return s
}

func (s *GetLakeStorageRequest) SetRegionId(v string) *GetLakeStorageRequest {
	s.RegionId = &v
	return s
}

func (s *GetLakeStorageRequest) Validate() error {
	return dara.Validate(s)
}
