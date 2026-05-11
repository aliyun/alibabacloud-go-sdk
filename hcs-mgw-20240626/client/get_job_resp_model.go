// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResp interface {
	dara.Model
	String() string
	GoString() string
	SetAppendableToNormal(v bool) *GetJobResp
	GetAppendableToNormal() *bool
	SetAudit(v *Audit) *GetJobResp
	GetAudit() *Audit
	SetConvertSymlinkTarget(v bool) *GetJobResp
	GetConvertSymlinkTarget() *bool
	SetCreateReport(v bool) *GetJobResp
	GetCreateReport() *bool
	SetCreateTime(v string) *GetJobResp
	GetCreateTime() *string
	SetDestAddress(v string) *GetJobResp
	GetDestAddress() *string
	SetEnableMultiVersioning(v bool) *GetJobResp
	GetEnableMultiVersioning() *bool
	SetFilterRule(v *FilterRule) *GetJobResp
	GetFilterRule() *FilterRule
	SetImportQos(v *ImportQos) *GetJobResp
	GetImportQos() *ImportQos
	SetModifyTime(v string) *GetJobResp
	GetModifyTime() *string
	SetName(v string) *GetJobResp
	GetName() *string
	SetOverwriteMode(v string) *GetJobResp
	GetOverwriteMode() *string
	SetOwner(v string) *GetJobResp
	GetOwner() *string
	SetParentName(v string) *GetJobResp
	GetParentName() *string
	SetParentVersion(v string) *GetJobResp
	GetParentVersion() *string
	SetScheduleRule(v *ScheduleRule) *GetJobResp
	GetScheduleRule() *ScheduleRule
	SetSrcAddress(v string) *GetJobResp
	GetSrcAddress() *string
	SetStatus(v string) *GetJobResp
	GetStatus() *string
	SetTags(v string) *GetJobResp
	GetTags() *string
	SetTargetStorageClass(v string) *GetJobResp
	GetTargetStorageClass() *string
	SetTransferMode(v string) *GetJobResp
	GetTransferMode() *string
	SetVersion(v string) *GetJobResp
	GetVersion() *string
	SetWithLastModifyTime(v bool) *GetJobResp
	GetWithLastModifyTime() *bool
	SetWithStorageClass(v bool) *GetJobResp
	GetWithStorageClass() *bool
}

type GetJobResp struct {
	// example:
	//
	// true
	AppendableToNormal *bool `json:"AppendableToNormal,omitempty" xml:"AppendableToNormal,omitempty"`
	// The audit method.
	Audit *Audit `json:"Audit,omitempty" xml:"Audit,omitempty"`
	// Indicates whether the Target attribute value of the symbolic links at the source data address is converted.
	//
	// example:
	//
	// false
	ConvertSymlinkTarget *bool `json:"ConvertSymlinkTarget,omitempty" xml:"ConvertSymlinkTarget,omitempty"`
	// Indicates whether a migration report is created.
	//
	// example:
	//
	// false
	CreateReport *bool `json:"CreateReport,omitempty" xml:"CreateReport,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2024-05-01 12:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the destination data address.
	//
	// example:
	//
	// test_dest_address
	DestAddress *string `json:"DestAddress,omitempty" xml:"DestAddress,omitempty"`
	// Indicates whether multi-version migration is enabled.
	//
	// example:
	//
	// false
	EnableMultiVersioning *bool `json:"EnableMultiVersioning,omitempty" xml:"EnableMultiVersioning,omitempty"`
	// The filtering rule.
	FilterRule *FilterRule `json:"FilterRule,omitempty" xml:"FilterRule,omitempty"`
	// The throttling settings of the task.
	ImportQos *ImportQos `json:"ImportQos,omitempty" xml:"ImportQos,omitempty"`
	// The time when the task was modified.
	//
	// example:
	//
	// 2024-05-01 12:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The task name.
	//
	// example:
	//
	// test_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The file overwriting mode.
	//
	// example:
	//
	// always
	OverwriteMode *string `json:"OverwriteMode,omitempty" xml:"OverwriteMode,omitempty"`
	// The task owner.
	//
	// example:
	//
	// test_owner
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the parent task.
	//
	// example:
	//
	// test_parent_name
	ParentName *string `json:"ParentName,omitempty" xml:"ParentName,omitempty"`
	// The ID of the parent task.
	//
	// example:
	//
	// 7db93837-a5ee-4e3a-b3c8-800e7947dabc
	ParentVersion *string `json:"ParentVersion,omitempty" xml:"ParentVersion,omitempty"`
	// The scheduling rule.
	ScheduleRule *ScheduleRule `json:"ScheduleRule,omitempty" xml:"ScheduleRule,omitempty"`
	// The name of the source data address.
	//
	// example:
	//
	// test_src_address
	SrcAddress *string `json:"SrcAddress,omitempty" xml:"SrcAddress,omitempty"`
	// The task state. Valid values:
	//
	// IMPORT_JOB_BEGIN: The task is created.
	//
	// IMPORT_JOB_LAUNCHING: The task is being started.
	//
	// IMPORT_JOB_PREPARING: The system is preparing for the task.
	//
	// IMPORT_JOB_DOING: The task is running.
	//
	// IMPORT_JOB_SUSPEND: The task is paused.
	//
	// IMPORT_JOB_CLOSING: The task is being closed.
	//
	// IMPORT_JOB_FINISHED: The task is complete.
	//
	// IMPORT_JOB_INTERRUPTED: The task is abnormally interrupted.
	//
	// IMPORT_JOB_CONFIRMED: The task is complete and the user has confirmed the data integrity and consistency.
	//
	// example:
	//
	// IMPORT_JOB_DOING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	//
	// example:
	//
	// K1:V1,K2:V2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// Standard
	TargetStorageClass *string `json:"TargetStorageClass,omitempty" xml:"TargetStorageClass,omitempty"`
	// The file transfer mode.
	//
	// example:
	//
	// all
	TransferMode *string `json:"TransferMode,omitempty" xml:"TransferMode,omitempty"`
	// The task ID.
	//
	// example:
	//
	// test_id
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// true
	WithLastModifyTime *bool `json:"WithLastModifyTime,omitempty" xml:"WithLastModifyTime,omitempty"`
	// example:
	//
	// false
	WithStorageClass *bool `json:"WithStorageClass,omitempty" xml:"WithStorageClass,omitempty"`
}

func (s GetJobResp) String() string {
	return dara.Prettify(s)
}

func (s GetJobResp) GoString() string {
	return s.String()
}

func (s *GetJobResp) GetAppendableToNormal() *bool {
	return s.AppendableToNormal
}

func (s *GetJobResp) GetAudit() *Audit {
	return s.Audit
}

func (s *GetJobResp) GetConvertSymlinkTarget() *bool {
	return s.ConvertSymlinkTarget
}

func (s *GetJobResp) GetCreateReport() *bool {
	return s.CreateReport
}

func (s *GetJobResp) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetJobResp) GetDestAddress() *string {
	return s.DestAddress
}

func (s *GetJobResp) GetEnableMultiVersioning() *bool {
	return s.EnableMultiVersioning
}

func (s *GetJobResp) GetFilterRule() *FilterRule {
	return s.FilterRule
}

func (s *GetJobResp) GetImportQos() *ImportQos {
	return s.ImportQos
}

func (s *GetJobResp) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetJobResp) GetName() *string {
	return s.Name
}

func (s *GetJobResp) GetOverwriteMode() *string {
	return s.OverwriteMode
}

func (s *GetJobResp) GetOwner() *string {
	return s.Owner
}

func (s *GetJobResp) GetParentName() *string {
	return s.ParentName
}

func (s *GetJobResp) GetParentVersion() *string {
	return s.ParentVersion
}

func (s *GetJobResp) GetScheduleRule() *ScheduleRule {
	return s.ScheduleRule
}

func (s *GetJobResp) GetSrcAddress() *string {
	return s.SrcAddress
}

func (s *GetJobResp) GetStatus() *string {
	return s.Status
}

func (s *GetJobResp) GetTags() *string {
	return s.Tags
}

func (s *GetJobResp) GetTargetStorageClass() *string {
	return s.TargetStorageClass
}

func (s *GetJobResp) GetTransferMode() *string {
	return s.TransferMode
}

func (s *GetJobResp) GetVersion() *string {
	return s.Version
}

func (s *GetJobResp) GetWithLastModifyTime() *bool {
	return s.WithLastModifyTime
}

func (s *GetJobResp) GetWithStorageClass() *bool {
	return s.WithStorageClass
}

func (s *GetJobResp) SetAppendableToNormal(v bool) *GetJobResp {
	s.AppendableToNormal = &v
	return s
}

func (s *GetJobResp) SetAudit(v *Audit) *GetJobResp {
	s.Audit = v
	return s
}

func (s *GetJobResp) SetConvertSymlinkTarget(v bool) *GetJobResp {
	s.ConvertSymlinkTarget = &v
	return s
}

func (s *GetJobResp) SetCreateReport(v bool) *GetJobResp {
	s.CreateReport = &v
	return s
}

func (s *GetJobResp) SetCreateTime(v string) *GetJobResp {
	s.CreateTime = &v
	return s
}

func (s *GetJobResp) SetDestAddress(v string) *GetJobResp {
	s.DestAddress = &v
	return s
}

func (s *GetJobResp) SetEnableMultiVersioning(v bool) *GetJobResp {
	s.EnableMultiVersioning = &v
	return s
}

func (s *GetJobResp) SetFilterRule(v *FilterRule) *GetJobResp {
	s.FilterRule = v
	return s
}

func (s *GetJobResp) SetImportQos(v *ImportQos) *GetJobResp {
	s.ImportQos = v
	return s
}

func (s *GetJobResp) SetModifyTime(v string) *GetJobResp {
	s.ModifyTime = &v
	return s
}

func (s *GetJobResp) SetName(v string) *GetJobResp {
	s.Name = &v
	return s
}

func (s *GetJobResp) SetOverwriteMode(v string) *GetJobResp {
	s.OverwriteMode = &v
	return s
}

func (s *GetJobResp) SetOwner(v string) *GetJobResp {
	s.Owner = &v
	return s
}

func (s *GetJobResp) SetParentName(v string) *GetJobResp {
	s.ParentName = &v
	return s
}

func (s *GetJobResp) SetParentVersion(v string) *GetJobResp {
	s.ParentVersion = &v
	return s
}

func (s *GetJobResp) SetScheduleRule(v *ScheduleRule) *GetJobResp {
	s.ScheduleRule = v
	return s
}

func (s *GetJobResp) SetSrcAddress(v string) *GetJobResp {
	s.SrcAddress = &v
	return s
}

func (s *GetJobResp) SetStatus(v string) *GetJobResp {
	s.Status = &v
	return s
}

func (s *GetJobResp) SetTags(v string) *GetJobResp {
	s.Tags = &v
	return s
}

func (s *GetJobResp) SetTargetStorageClass(v string) *GetJobResp {
	s.TargetStorageClass = &v
	return s
}

func (s *GetJobResp) SetTransferMode(v string) *GetJobResp {
	s.TransferMode = &v
	return s
}

func (s *GetJobResp) SetVersion(v string) *GetJobResp {
	s.Version = &v
	return s
}

func (s *GetJobResp) SetWithLastModifyTime(v bool) *GetJobResp {
	s.WithLastModifyTime = &v
	return s
}

func (s *GetJobResp) SetWithStorageClass(v bool) *GetJobResp {
	s.WithStorageClass = &v
	return s
}

func (s *GetJobResp) Validate() error {
	if s.Audit != nil {
		if err := s.Audit.Validate(); err != nil {
			return err
		}
	}
	if s.FilterRule != nil {
		if err := s.FilterRule.Validate(); err != nil {
			return err
		}
	}
	if s.ImportQos != nil {
		if err := s.ImportQos.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduleRule != nil {
		if err := s.ScheduleRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}
