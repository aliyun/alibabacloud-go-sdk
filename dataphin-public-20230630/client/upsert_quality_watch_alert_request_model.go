// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityWatchAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpsertQualityWatchAlertRequest
	GetOpTenantId() *int64
	SetUpsertCommand(v *UpsertQualityWatchAlertRequestUpsertCommand) *UpsertQualityWatchAlertRequest
	GetUpsertCommand() *UpsertQualityWatchAlertRequestUpsertCommand
}

type UpsertQualityWatchAlertRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpsertCommand *UpsertQualityWatchAlertRequestUpsertCommand `json:"UpsertCommand,omitempty" xml:"UpsertCommand,omitempty" type:"Struct"`
}

func (s UpsertQualityWatchAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityWatchAlertRequest) GoString() string {
	return s.String()
}

func (s *UpsertQualityWatchAlertRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpsertQualityWatchAlertRequest) GetUpsertCommand() *UpsertQualityWatchAlertRequestUpsertCommand {
	return s.UpsertCommand
}

func (s *UpsertQualityWatchAlertRequest) SetOpTenantId(v int64) *UpsertQualityWatchAlertRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpsertQualityWatchAlertRequest) SetUpsertCommand(v *UpsertQualityWatchAlertRequestUpsertCommand) *UpsertQualityWatchAlertRequest {
	s.UpsertCommand = v
	return s
}

func (s *UpsertQualityWatchAlertRequest) Validate() error {
	if s.UpsertCommand != nil {
		if err := s.UpsertCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpsertQualityWatchAlertRequestUpsertCommand struct {
	// This parameter is required.
	QualityAlertInfo *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo `json:"QualityAlertInfo,omitempty" xml:"QualityAlertInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 11
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s UpsertQualityWatchAlertRequestUpsertCommand) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityWatchAlertRequestUpsertCommand) GoString() string {
	return s.String()
}

func (s *UpsertQualityWatchAlertRequestUpsertCommand) GetQualityAlertInfo() *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo {
	return s.QualityAlertInfo
}

func (s *UpsertQualityWatchAlertRequestUpsertCommand) GetWatchId() *int64 {
	return s.WatchId
}

func (s *UpsertQualityWatchAlertRequestUpsertCommand) SetQualityAlertInfo(v *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo) *UpsertQualityWatchAlertRequestUpsertCommand {
	s.QualityAlertInfo = v
	return s
}

func (s *UpsertQualityWatchAlertRequestUpsertCommand) SetWatchId(v int64) *UpsertQualityWatchAlertRequestUpsertCommand {
	s.WatchId = &v
	return s
}

func (s *UpsertQualityWatchAlertRequestUpsertCommand) Validate() error {
	if s.QualityAlertInfo != nil {
		if err := s.QualityAlertInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo struct {
	AlertDutyChannelList         []*string                                                                   `json:"AlertDutyChannelList,omitempty" xml:"AlertDutyChannelList,omitempty" type:"Repeated"`
	AlertDutyList                []*UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertDutyList `json:"AlertDutyList,omitempty" xml:"AlertDutyList,omitempty" type:"Repeated"`
	AlertQualityOwnerChannelList []*string                                                                   `json:"AlertQualityOwnerChannelList,omitempty" xml:"AlertQualityOwnerChannelList,omitempty" type:"Repeated"`
	AlertUserChannelList         []*string                                                                   `json:"AlertUserChannelList,omitempty" xml:"AlertUserChannelList,omitempty" type:"Repeated"`
	AlertUserList                []*UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertUserList `json:"AlertUserList,omitempty" xml:"AlertUserList,omitempty" type:"Repeated"`
	EnableAlertQualityOwner      *bool                                                                       `json:"EnableAlertQualityOwner,omitempty" xml:"EnableAlertQualityOwner,omitempty"`
}

func (s UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo) GoString() string {
	return s.String()
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo) GetAlertDutyChannelList() []*string {
	return s.AlertDutyChannelList
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo) GetAlertDutyList() []*UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertDutyList {
	return s.AlertDutyList
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo) GetAlertQualityOwnerChannelList() []*string {
	return s.AlertQualityOwnerChannelList
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo) GetAlertUserChannelList() []*string {
	return s.AlertUserChannelList
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo) GetAlertUserList() []*UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertUserList {
	return s.AlertUserList
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo) GetEnableAlertQualityOwner() *bool {
	return s.EnableAlertQualityOwner
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo) SetAlertDutyChannelList(v []*string) *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo {
	s.AlertDutyChannelList = v
	return s
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo) SetAlertDutyList(v []*UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertDutyList) *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo {
	s.AlertDutyList = v
	return s
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo) SetAlertQualityOwnerChannelList(v []*string) *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo {
	s.AlertQualityOwnerChannelList = v
	return s
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo) SetAlertUserChannelList(v []*string) *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo {
	s.AlertUserChannelList = v
	return s
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo) SetAlertUserList(v []*UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertUserList) *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo {
	s.AlertUserList = v
	return s
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo) SetEnableAlertQualityOwner(v bool) *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo {
	s.EnableAlertQualityOwner = &v
	return s
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfo) Validate() error {
	if s.AlertDutyList != nil {
		for _, item := range s.AlertDutyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AlertUserList != nil {
		for _, item := range s.AlertUserList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertDutyList struct {
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertDutyList) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertDutyList) GoString() string {
	return s.String()
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertDutyList) GetId() *string {
	return s.Id
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertDutyList) GetName() *string {
	return s.Name
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertDutyList) SetId(v string) *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertDutyList {
	s.Id = &v
	return s
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertDutyList) SetName(v string) *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertDutyList {
	s.Name = &v
	return s
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertDutyList) Validate() error {
	return dara.Validate(s)
}

type UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertUserList struct {
	// example:
	//
	// 30012011
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertUserList) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertUserList) GoString() string {
	return s.String()
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertUserList) GetId() *string {
	return s.Id
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertUserList) GetName() *string {
	return s.Name
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertUserList) SetId(v string) *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertUserList {
	s.Id = &v
	return s
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertUserList) SetName(v string) *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertUserList {
	s.Name = &v
	return s
}

func (s *UpsertQualityWatchAlertRequestUpsertCommandQualityAlertInfoAlertUserList) Validate() error {
	return dara.Validate(s)
}
