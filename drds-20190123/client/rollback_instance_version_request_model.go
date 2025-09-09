// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackInstanceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *RollbackInstanceVersionRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *RollbackInstanceVersionRequest
	GetRegionId() *string
}

type RollbackInstanceVersionRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RollbackInstanceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackInstanceVersionRequest) GoString() string {
	return s.String()
}

func (s *RollbackInstanceVersionRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *RollbackInstanceVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RollbackInstanceVersionRequest) SetDrdsInstanceId(v string) *RollbackInstanceVersionRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RollbackInstanceVersionRequest) SetRegionId(v string) *RollbackInstanceVersionRequest {
	s.RegionId = &v
	return s
}

func (s *RollbackInstanceVersionRequest) Validate() error {
	return dara.Validate(s)
}
