// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpenSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DeleteOpenSearchRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DeleteOpenSearchRequest
	GetRegionId() *string
}

type DeleteOpenSearchRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-spsil01pww4hfz
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteOpenSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSearchRequest) GoString() string {
	return s.String()
}

func (s *DeleteOpenSearchRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteOpenSearchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteOpenSearchRequest) SetDBInstanceName(v string) *DeleteOpenSearchRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteOpenSearchRequest) SetRegionId(v string) *DeleteOpenSearchRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteOpenSearchRequest) Validate() error {
	return dara.Validate(s)
}
