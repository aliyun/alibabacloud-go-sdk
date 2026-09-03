// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAppendableToNormal(v bool) *CreateJobInfo
	GetAppendableToNormal() *bool
	SetAudit(v *Audit) *CreateJobInfo
	GetAudit() *Audit
	SetConvertSymlinkTarget(v bool) *CreateJobInfo
	GetConvertSymlinkTarget() *bool
	SetCreateReport(v bool) *CreateJobInfo
	GetCreateReport() *bool
	SetDestAddress(v string) *CreateJobInfo
	GetDestAddress() *string
	SetEnableMultiVersioning(v bool) *CreateJobInfo
	GetEnableMultiVersioning() *bool
	SetFilterRule(v *FilterRule) *CreateJobInfo
	GetFilterRule() *FilterRule
	SetImportQos(v *ImportQos) *CreateJobInfo
	GetImportQos() *ImportQos
	SetName(v string) *CreateJobInfo
	GetName() *string
	SetOverwriteMode(v string) *CreateJobInfo
	GetOverwriteMode() *string
	SetParentVersion(v string) *CreateJobInfo
	GetParentVersion() *string
	SetScheduleRule(v *ScheduleRule) *CreateJobInfo
	GetScheduleRule() *ScheduleRule
	SetSrcAddress(v string) *CreateJobInfo
	GetSrcAddress() *string
	SetTags(v string) *CreateJobInfo
	GetTags() *string
	SetTargetStorageClass(v string) *CreateJobInfo
	GetTargetStorageClass() *string
	SetTransferMode(v string) *CreateJobInfo
	GetTransferMode() *string
	SetWithLastModifyTime(v bool) *CreateJobInfo
	GetWithLastModifyTime() *bool
	SetWithStorageClass(v bool) *CreateJobInfo
	GetWithStorageClass() *bool
}

type CreateJobInfo struct {
	// Specifies whether to migrate appendable files as normal or multipart files. Default value: false.
	//
	// example:
	//
	// false
	AppendableToNormal *bool `json:"AppendableToNormal,omitempty" xml:"AppendableToNormal,omitempty"`
	// The audit method.
	Audit *Audit `json:"Audit,omitempty" xml:"Audit,omitempty"`
	// Specifies whether to transform the target of a symbolic link. When migrating data from OSS to a local server, from a local server to OSS, or between two local servers, set this parameter to \\`true\\` to ensure that symbolic links can be accessed after migration.
	//
	// example:
	//
	// false
	ConvertSymlinkTarget *bool `json:"ConvertSymlinkTarget,omitempty" xml:"ConvertSymlinkTarget,omitempty"`
	// Specifies whether to create a migration report.
	//
	// example:
	//
	// true
	CreateReport *bool `json:"CreateReport,omitempty" xml:"CreateReport,omitempty"`
	// The name of the destination data address.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_dest_address
	DestAddress *string `json:"DestAddress,omitempty" xml:"DestAddress,omitempty"`
	// Specifies whether to migrate multiple object versions. Multi-version migration is not supported.
	//
	// example:
	//
	// false
	EnableMultiVersioning *bool `json:"EnableMultiVersioning,omitempty" xml:"EnableMultiVersioning,omitempty"`
	// The filter rule.
	FilterRule *FilterRule `json:"FilterRule,omitempty" xml:"FilterRule,omitempty"`
	// The task throttling settings.
	ImportQos *ImportQos `json:"ImportQos,omitempty" xml:"ImportQos,omitempty"`
	// The task name.<br>
	//
	// The name must be 3 to 63 characters in length and can contain lowercase letters, digits, hyphens (-), and underscores (_). The name is case-sensitive and must be UTF-8 encoded. It cannot start with a hyphen (-) or an underscore (_). This parameter cannot be empty.<br>
	//
	// This parameter is required.
	//
	// example:
	//
	// test_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The file overwrite mode.<br>
	//
	// Valid values: \\`never\\` and \\`always\\`. \\`never\\`: Does not overwrite existing files. \\`always\\`: Overwrites existing files.<br>
	//
	// This parameter is required.
	//
	// example:
	//
	// always
	OverwriteMode *string `json:"OverwriteMode,omitempty" xml:"OverwriteMode,omitempty"`
	// The parent task ID. Specify this ID when you create a subtask to retry failed file transfers.
	//
	// example:
	//
	// 6af62558-970d-4f44-8663-4e297170fd6a
	ParentVersion *string `json:"ParentVersion,omitempty" xml:"ParentVersion,omitempty"`
	// The scheduling rule.
	ScheduleRule *ScheduleRule `json:"ScheduleRule,omitempty" xml:"ScheduleRule,omitempty"`
	// The name of the source data address.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_src_address
	SrcAddress *string `json:"SrcAddress,omitempty" xml:"SrcAddress,omitempty"`
	// The tags, in key-value format.<br>
	//
	// Allowed characters include uppercase and lowercase letters, digits, hyphens (-), and underscores (_). The maximum length is 1024 characters.<br>
	//
	// example:
	//
	// K1:V1,K2:V2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Specifies the StorageClass for destination files. The destination address can only be OSS. Valid values: Standard, IA, Archive, ColdArchive, DeepColdArchive.
	//
	// example:
	//
	// Standard
	TargetStorageClass *string `json:"TargetStorageClass,omitempty" xml:"TargetStorageClass,omitempty"`
	// The file transfer mode.<br>
	//
	// Valid values: \\`all\\` (full transfer) and \\`lastmodified\\` (incremental transfer).<br>
	//
	// \\`OverwriteMode\\` and \\`TransferMode\\` are used together:<br><br>
	//
	// - \\`always\\` and \\`all\\`: Forces a full overwrite.
	//
	// - \\`always\\` and \\`lastmodified\\`: Overwrites files based on their last modified time.
	//
	// - \\`never\\` and an empty value for \\`TransferMode\\`: Does not overwrite files with the same name.
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	TransferMode *string `json:"TransferMode,omitempty" xml:"TransferMode,omitempty"`
	// Specifies whether to preserve the lastModifyTime. Default value: true.
	//
	// example:
	//
	// true
	WithLastModifyTime *bool `json:"WithLastModifyTime,omitempty" xml:"WithLastModifyTime,omitempty"`
	// Specifies whether to migrate the StorageClass property. This is allowed only for OSS-to-OSS migration.
	//
	// example:
	//
	// false
	WithStorageClass *bool `json:"WithStorageClass,omitempty" xml:"WithStorageClass,omitempty"`
}

func (s CreateJobInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateJobInfo) GoString() string {
	return s.String()
}

func (s *CreateJobInfo) GetAppendableToNormal() *bool {
	return s.AppendableToNormal
}

func (s *CreateJobInfo) GetAudit() *Audit {
	return s.Audit
}

func (s *CreateJobInfo) GetConvertSymlinkTarget() *bool {
	return s.ConvertSymlinkTarget
}

func (s *CreateJobInfo) GetCreateReport() *bool {
	return s.CreateReport
}

func (s *CreateJobInfo) GetDestAddress() *string {
	return s.DestAddress
}

func (s *CreateJobInfo) GetEnableMultiVersioning() *bool {
	return s.EnableMultiVersioning
}

func (s *CreateJobInfo) GetFilterRule() *FilterRule {
	return s.FilterRule
}

func (s *CreateJobInfo) GetImportQos() *ImportQos {
	return s.ImportQos
}

func (s *CreateJobInfo) GetName() *string {
	return s.Name
}

func (s *CreateJobInfo) GetOverwriteMode() *string {
	return s.OverwriteMode
}

func (s *CreateJobInfo) GetParentVersion() *string {
	return s.ParentVersion
}

func (s *CreateJobInfo) GetScheduleRule() *ScheduleRule {
	return s.ScheduleRule
}

func (s *CreateJobInfo) GetSrcAddress() *string {
	return s.SrcAddress
}

func (s *CreateJobInfo) GetTags() *string {
	return s.Tags
}

func (s *CreateJobInfo) GetTargetStorageClass() *string {
	return s.TargetStorageClass
}

func (s *CreateJobInfo) GetTransferMode() *string {
	return s.TransferMode
}

func (s *CreateJobInfo) GetWithLastModifyTime() *bool {
	return s.WithLastModifyTime
}

func (s *CreateJobInfo) GetWithStorageClass() *bool {
	return s.WithStorageClass
}

func (s *CreateJobInfo) SetAppendableToNormal(v bool) *CreateJobInfo {
	s.AppendableToNormal = &v
	return s
}

func (s *CreateJobInfo) SetAudit(v *Audit) *CreateJobInfo {
	s.Audit = v
	return s
}

func (s *CreateJobInfo) SetConvertSymlinkTarget(v bool) *CreateJobInfo {
	s.ConvertSymlinkTarget = &v
	return s
}

func (s *CreateJobInfo) SetCreateReport(v bool) *CreateJobInfo {
	s.CreateReport = &v
	return s
}

func (s *CreateJobInfo) SetDestAddress(v string) *CreateJobInfo {
	s.DestAddress = &v
	return s
}

func (s *CreateJobInfo) SetEnableMultiVersioning(v bool) *CreateJobInfo {
	s.EnableMultiVersioning = &v
	return s
}

func (s *CreateJobInfo) SetFilterRule(v *FilterRule) *CreateJobInfo {
	s.FilterRule = v
	return s
}

func (s *CreateJobInfo) SetImportQos(v *ImportQos) *CreateJobInfo {
	s.ImportQos = v
	return s
}

func (s *CreateJobInfo) SetName(v string) *CreateJobInfo {
	s.Name = &v
	return s
}

func (s *CreateJobInfo) SetOverwriteMode(v string) *CreateJobInfo {
	s.OverwriteMode = &v
	return s
}

func (s *CreateJobInfo) SetParentVersion(v string) *CreateJobInfo {
	s.ParentVersion = &v
	return s
}

func (s *CreateJobInfo) SetScheduleRule(v *ScheduleRule) *CreateJobInfo {
	s.ScheduleRule = v
	return s
}

func (s *CreateJobInfo) SetSrcAddress(v string) *CreateJobInfo {
	s.SrcAddress = &v
	return s
}

func (s *CreateJobInfo) SetTags(v string) *CreateJobInfo {
	s.Tags = &v
	return s
}

func (s *CreateJobInfo) SetTargetStorageClass(v string) *CreateJobInfo {
	s.TargetStorageClass = &v
	return s
}

func (s *CreateJobInfo) SetTransferMode(v string) *CreateJobInfo {
	s.TransferMode = &v
	return s
}

func (s *CreateJobInfo) SetWithLastModifyTime(v bool) *CreateJobInfo {
	s.WithLastModifyTime = &v
	return s
}

func (s *CreateJobInfo) SetWithStorageClass(v bool) *CreateJobInfo {
	s.WithStorageClass = &v
	return s
}

func (s *CreateJobInfo) Validate() error {
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
