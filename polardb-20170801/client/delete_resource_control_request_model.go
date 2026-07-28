// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteResourceControlRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DeleteResourceControlRequest
	GetRegionId() *string
	SetResourceControlName(v string) *DeleteResourceControlRequest
	GetResourceControlName() *string
}

type DeleteResourceControlRequest struct {
	// The database cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource control name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-rc
	ResourceControlName *string `json:"ResourceControlName,omitempty" xml:"ResourceControlName,omitempty"`
}

func (s DeleteResourceControlRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceControlRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceControlRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteResourceControlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteResourceControlRequest) GetResourceControlName() *string {
	return s.ResourceControlName
}

func (s *DeleteResourceControlRequest) SetDBClusterId(v string) *DeleteResourceControlRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteResourceControlRequest) SetRegionId(v string) *DeleteResourceControlRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteResourceControlRequest) SetResourceControlName(v string) *DeleteResourceControlRequest {
	s.ResourceControlName = &v
	return s
}

func (s *DeleteResourceControlRequest) Validate() error {
	return dara.Validate(s)
}
