// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLangfuseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteLangfuseInstanceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DeleteLangfuseInstanceRequest
	GetRegionId() *string
}

type DeleteLangfuseInstanceRequest struct {
	// The Langfuse instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lfs-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteLangfuseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteLangfuseInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLangfuseInstanceRequest) SetDBInstanceId(v string) *DeleteLangfuseInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteLangfuseInstanceRequest) SetRegionId(v string) *DeleteLangfuseInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLangfuseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
