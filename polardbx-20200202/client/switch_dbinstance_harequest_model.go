// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDBInstanceHARequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *SwitchDBInstanceHARequest
	GetDBInstanceName() *string
	SetRegionId(v string) *SwitchDBInstanceHARequest
	GetRegionId() *string
	SetSwitchTime(v string) *SwitchDBInstanceHARequest
	GetSwitchTime() *string
	SetSwitchTimeMode(v string) *SwitchDBInstanceHARequest
	GetSwitchTimeMode() *string
	SetTargetPrimaryAzoneId(v string) *SwitchDBInstanceHARequest
	GetTargetPrimaryAzoneId() *string
	SetTargetPrimaryRegionId(v string) *SwitchDBInstanceHARequest
	GetTargetPrimaryRegionId() *string
}

type SwitchDBInstanceHARequest struct {
	// This parameter is required.
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SwitchTime            *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	SwitchTimeMode        *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
	TargetPrimaryAzoneId  *string `json:"TargetPrimaryAzoneId,omitempty" xml:"TargetPrimaryAzoneId,omitempty"`
	TargetPrimaryRegionId *string `json:"TargetPrimaryRegionId,omitempty" xml:"TargetPrimaryRegionId,omitempty"`
}

func (s SwitchDBInstanceHARequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchDBInstanceHARequest) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceHARequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *SwitchDBInstanceHARequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SwitchDBInstanceHARequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *SwitchDBInstanceHARequest) GetSwitchTimeMode() *string {
	return s.SwitchTimeMode
}

func (s *SwitchDBInstanceHARequest) GetTargetPrimaryAzoneId() *string {
	return s.TargetPrimaryAzoneId
}

func (s *SwitchDBInstanceHARequest) GetTargetPrimaryRegionId() *string {
	return s.TargetPrimaryRegionId
}

func (s *SwitchDBInstanceHARequest) SetDBInstanceName(v string) *SwitchDBInstanceHARequest {
	s.DBInstanceName = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetRegionId(v string) *SwitchDBInstanceHARequest {
	s.RegionId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetSwitchTime(v string) *SwitchDBInstanceHARequest {
	s.SwitchTime = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetSwitchTimeMode(v string) *SwitchDBInstanceHARequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetTargetPrimaryAzoneId(v string) *SwitchDBInstanceHARequest {
	s.TargetPrimaryAzoneId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetTargetPrimaryRegionId(v string) *SwitchDBInstanceHARequest {
	s.TargetPrimaryRegionId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) Validate() error {
	return dara.Validate(s)
}
