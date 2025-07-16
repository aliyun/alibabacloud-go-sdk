// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableRightsSeparationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DisableRightsSeparationRequest
	GetDBInstanceName() *string
	SetDbaAccountName(v string) *DisableRightsSeparationRequest
	GetDbaAccountName() *string
	SetDbaAccountPassword(v string) *DisableRightsSeparationRequest
	GetDbaAccountPassword() *string
	SetRegionId(v string) *DisableRightsSeparationRequest
	GetRegionId() *string
}

type DisableRightsSeparationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-sprcym7g7w****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// account_1
	DbaAccountName *string `json:"DbaAccountName,omitempty" xml:"DbaAccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// *****
	DbaAccountPassword *string `json:"DbaAccountPassword,omitempty" xml:"DbaAccountPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hanghzou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableRightsSeparationRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableRightsSeparationRequest) GoString() string {
	return s.String()
}

func (s *DisableRightsSeparationRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DisableRightsSeparationRequest) GetDbaAccountName() *string {
	return s.DbaAccountName
}

func (s *DisableRightsSeparationRequest) GetDbaAccountPassword() *string {
	return s.DbaAccountPassword
}

func (s *DisableRightsSeparationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableRightsSeparationRequest) SetDBInstanceName(v string) *DisableRightsSeparationRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DisableRightsSeparationRequest) SetDbaAccountName(v string) *DisableRightsSeparationRequest {
	s.DbaAccountName = &v
	return s
}

func (s *DisableRightsSeparationRequest) SetDbaAccountPassword(v string) *DisableRightsSeparationRequest {
	s.DbaAccountPassword = &v
	return s
}

func (s *DisableRightsSeparationRequest) SetRegionId(v string) *DisableRightsSeparationRequest {
	s.RegionId = &v
	return s
}

func (s *DisableRightsSeparationRequest) Validate() error {
	return dara.Validate(s)
}
