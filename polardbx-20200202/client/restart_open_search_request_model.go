// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartOpenSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *RestartOpenSearchRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *RestartOpenSearchRequest
	GetRegionId() *string
}

type RestartOpenSearchRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RestartOpenSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartOpenSearchRequest) GoString() string {
	return s.String()
}

func (s *RestartOpenSearchRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *RestartOpenSearchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartOpenSearchRequest) SetDBInstanceName(v string) *RestartOpenSearchRequest {
	s.DBInstanceName = &v
	return s
}

func (s *RestartOpenSearchRequest) SetRegionId(v string) *RestartOpenSearchRequest {
	s.RegionId = &v
	return s
}

func (s *RestartOpenSearchRequest) Validate() error {
	return dara.Validate(s)
}
