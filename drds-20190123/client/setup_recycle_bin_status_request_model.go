// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetupRecycleBinStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *SetupRecycleBinStatusRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *SetupRecycleBinStatusRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *SetupRecycleBinStatusRequest
	GetRegionId() *string
	SetStatusAction(v string) *SetupRecycleBinStatusRequest
	GetStatusAction() *string
}

type SetupRecycleBinStatusRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies the status of the table recycle bin. Valid values:
	//
	// 	- enable: The table recycle bin is enabled.
	//
	// 	- disable: The table recycle bin is disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// enable
	StatusAction *string `json:"StatusAction,omitempty" xml:"StatusAction,omitempty"`
}

func (s SetupRecycleBinStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetupRecycleBinStatusRequest) GoString() string {
	return s.String()
}

func (s *SetupRecycleBinStatusRequest) GetDbName() *string {
	return s.DbName
}

func (s *SetupRecycleBinStatusRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *SetupRecycleBinStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetupRecycleBinStatusRequest) GetStatusAction() *string {
	return s.StatusAction
}

func (s *SetupRecycleBinStatusRequest) SetDbName(v string) *SetupRecycleBinStatusRequest {
	s.DbName = &v
	return s
}

func (s *SetupRecycleBinStatusRequest) SetDrdsInstanceId(v string) *SetupRecycleBinStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetupRecycleBinStatusRequest) SetRegionId(v string) *SetupRecycleBinStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetupRecycleBinStatusRequest) SetStatusAction(v string) *SetupRecycleBinStatusRequest {
	s.StatusAction = &v
	return s
}

func (s *SetupRecycleBinStatusRequest) Validate() error {
	return dara.Validate(s)
}
