// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateColdDataVolumeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *AllocateColdDataVolumeRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *AllocateColdDataVolumeRequest
	GetRegionId() *string
}

type AllocateColdDataVolumeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzravgpt8q****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AllocateColdDataVolumeRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateColdDataVolumeRequest) GoString() string {
	return s.String()
}

func (s *AllocateColdDataVolumeRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *AllocateColdDataVolumeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocateColdDataVolumeRequest) SetDBInstanceName(v string) *AllocateColdDataVolumeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *AllocateColdDataVolumeRequest) SetRegionId(v string) *AllocateColdDataVolumeRequest {
	s.RegionId = &v
	return s
}

func (s *AllocateColdDataVolumeRequest) Validate() error {
	return dara.Validate(s)
}
