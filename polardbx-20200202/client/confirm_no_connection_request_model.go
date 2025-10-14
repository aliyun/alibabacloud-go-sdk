// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmNoConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ConfirmNoConnectionRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *ConfirmNoConnectionRequest
	GetRegionId() *string
	SetSlinkTaskId(v string) *ConfirmNoConnectionRequest
	GetSlinkTaskId() *string
}

type ConfirmNoConnectionRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
}

func (s ConfirmNoConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmNoConnectionRequest) GoString() string {
	return s.String()
}

func (s *ConfirmNoConnectionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ConfirmNoConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfirmNoConnectionRequest) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *ConfirmNoConnectionRequest) SetDBInstanceName(v string) *ConfirmNoConnectionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ConfirmNoConnectionRequest) SetRegionId(v string) *ConfirmNoConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *ConfirmNoConnectionRequest) SetSlinkTaskId(v string) *ConfirmNoConnectionRequest {
	s.SlinkTaskId = &v
	return s
}

func (s *ConfirmNoConnectionRequest) Validate() error {
	return dara.Validate(s)
}
