// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSyncEcsHostTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateSyncEcsHostTaskRequest
	GetLang() *string
	SetRegion(v []*UpdateSyncEcsHostTaskRequestRegion) *UpdateSyncEcsHostTaskRequest
	GetRegion() []*UpdateSyncEcsHostTaskRequestRegion
	SetStatus(v string) *UpdateSyncEcsHostTaskRequest
	GetStatus() *string
	SetZoneId(v string) *UpdateSyncEcsHostTaskRequest
	GetZoneId() *string
}

type UpdateSyncEcsHostTaskRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The regions to be synchronized.
	//
	// This parameter is required.
	Region []*UpdateSyncEcsHostTaskRequestRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
	// The state of the hostname synchronization task. Valid values:
	//
	// 	- ON: The task is started.
	//
	// 	- OFF: The task is ended.
	//
	// This parameter is required.
	//
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The zone ID. This ID uniquely identifies the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// df2d03865266bd9842306db586d3****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateSyncEcsHostTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSyncEcsHostTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateSyncEcsHostTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateSyncEcsHostTaskRequest) GetRegion() []*UpdateSyncEcsHostTaskRequestRegion {
	return s.Region
}

func (s *UpdateSyncEcsHostTaskRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateSyncEcsHostTaskRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateSyncEcsHostTaskRequest) SetLang(v string) *UpdateSyncEcsHostTaskRequest {
	s.Lang = &v
	return s
}

func (s *UpdateSyncEcsHostTaskRequest) SetRegion(v []*UpdateSyncEcsHostTaskRequestRegion) *UpdateSyncEcsHostTaskRequest {
	s.Region = v
	return s
}

func (s *UpdateSyncEcsHostTaskRequest) SetStatus(v string) *UpdateSyncEcsHostTaskRequest {
	s.Status = &v
	return s
}

func (s *UpdateSyncEcsHostTaskRequest) SetZoneId(v string) *UpdateSyncEcsHostTaskRequest {
	s.ZoneId = &v
	return s
}

func (s *UpdateSyncEcsHostTaskRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateSyncEcsHostTaskRequestRegion struct {
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID to which the region belongs. This parameter is used in cross-account synchronization scenarios.
	//
	// example:
	//
	// 141339776561****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateSyncEcsHostTaskRequestRegion) String() string {
	return dara.Prettify(s)
}

func (s UpdateSyncEcsHostTaskRequestRegion) GoString() string {
	return s.String()
}

func (s *UpdateSyncEcsHostTaskRequestRegion) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSyncEcsHostTaskRequestRegion) GetUserId() *int64 {
	return s.UserId
}

func (s *UpdateSyncEcsHostTaskRequestRegion) SetRegionId(v string) *UpdateSyncEcsHostTaskRequestRegion {
	s.RegionId = &v
	return s
}

func (s *UpdateSyncEcsHostTaskRequestRegion) SetUserId(v int64) *UpdateSyncEcsHostTaskRequestRegion {
	s.UserId = &v
	return s
}

func (s *UpdateSyncEcsHostTaskRequestRegion) Validate() error {
	return dara.Validate(s)
}
