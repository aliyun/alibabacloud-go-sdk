// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseColdDataVolumeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ReleaseColdDataVolumeRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *ReleaseColdDataVolumeRequest
	GetRegionId() *string
}

type ReleaseColdDataVolumeRequest struct {
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

func (s ReleaseColdDataVolumeRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseColdDataVolumeRequest) GoString() string {
	return s.String()
}

func (s *ReleaseColdDataVolumeRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ReleaseColdDataVolumeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleaseColdDataVolumeRequest) SetDBInstanceName(v string) *ReleaseColdDataVolumeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ReleaseColdDataVolumeRequest) SetRegionId(v string) *ReleaseColdDataVolumeRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseColdDataVolumeRequest) Validate() error {
	return dara.Validate(s)
}
