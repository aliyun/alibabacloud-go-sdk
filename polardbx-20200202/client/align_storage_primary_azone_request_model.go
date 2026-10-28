// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlignStoragePrimaryAzoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *AlignStoragePrimaryAzoneRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *AlignStoragePrimaryAzoneRequest
	GetRegionId() *string
	SetStorageInstanceName(v string) *AlignStoragePrimaryAzoneRequest
	GetStorageInstanceName() *string
	SetSwitchTime(v string) *AlignStoragePrimaryAzoneRequest
	GetSwitchTime() *string
	SetSwitchTimeMode(v string) *AlignStoragePrimaryAzoneRequest
	GetSwitchTimeMode() *string
}

type AlignStoragePrimaryAzoneRequest struct {
	// The name of the PolarDB-X instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0ori2r****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the storage instance whose primary node zone needs to be aligned.
	//
	// example:
	//
	// pxc-xdb-m-pxcbjrylg49skcxb17394
	StorageInstanceName *string `json:"StorageInstanceName,omitempty" xml:"StorageInstanceName,omitempty"`
	// The start time of the switch for the storage primary node, in the UTC time format YYYY-MM-DDThh:mm:ssZ.
	//
	// example:
	//
	// 2024-12-11T17:10:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The switchover mode. Valid values:
	//
	// - 0: immediate switchover
	//
	// - 1: switchover during the O&M window
	//
	// - 2: switchover at a specified time
	//
	// example:
	//
	// 0
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s AlignStoragePrimaryAzoneRequest) String() string {
	return dara.Prettify(s)
}

func (s AlignStoragePrimaryAzoneRequest) GoString() string {
	return s.String()
}

func (s *AlignStoragePrimaryAzoneRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *AlignStoragePrimaryAzoneRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AlignStoragePrimaryAzoneRequest) GetStorageInstanceName() *string {
	return s.StorageInstanceName
}

func (s *AlignStoragePrimaryAzoneRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *AlignStoragePrimaryAzoneRequest) GetSwitchTimeMode() *string {
	return s.SwitchTimeMode
}

func (s *AlignStoragePrimaryAzoneRequest) SetDBInstanceName(v string) *AlignStoragePrimaryAzoneRequest {
	s.DBInstanceName = &v
	return s
}

func (s *AlignStoragePrimaryAzoneRequest) SetRegionId(v string) *AlignStoragePrimaryAzoneRequest {
	s.RegionId = &v
	return s
}

func (s *AlignStoragePrimaryAzoneRequest) SetStorageInstanceName(v string) *AlignStoragePrimaryAzoneRequest {
	s.StorageInstanceName = &v
	return s
}

func (s *AlignStoragePrimaryAzoneRequest) SetSwitchTime(v string) *AlignStoragePrimaryAzoneRequest {
	s.SwitchTime = &v
	return s
}

func (s *AlignStoragePrimaryAzoneRequest) SetSwitchTimeMode(v string) *AlignStoragePrimaryAzoneRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *AlignStoragePrimaryAzoneRequest) Validate() error {
	return dara.Validate(s)
}
