// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-oss/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddressDetail struct {
	// This parameter is required.
	//
	// example:
	//
	// test_access_id
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_secret_key
	AccessSecret *string `json:"AccessSecret,omitempty" xml:"AccessSecret,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ossinv
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	AgentList   *string `json:"AgentList,omitempty" xml:"AgentList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_domain
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// test_inv_access_id
	InvAccessId *string `json:"InvAccessId,omitempty" xml:"InvAccessId,omitempty"`
	// example:
	//
	// test_inv_secret_key
	InvAccessSecret *string `json:"InvAccessSecret,omitempty" xml:"InvAccessSecret,omitempty"`
	// example:
	//
	// test_inv_bucket
	InvBucket *string `json:"InvBucket,omitempty" xml:"InvBucket,omitempty"`
	// example:
	//
	// test_inv_domain
	InvDomain *string `json:"InvDomain,omitempty" xml:"InvDomain,omitempty"`
	// example:
	//
	// oss
	InvLocation *string `json:"InvLocation,omitempty" xml:"InvLocation,omitempty"`
	// example:
	//
	// manifest.json
	InvPath *string `json:"InvPath,omitempty" xml:"InvPath,omitempty"`
	// example:
	//
	// test_inv_region_id
	InvRegionId *string `json:"InvRegionId,omitempty" xml:"InvRegionId,omitempty"`
	// example:
	//
	// test_inv_role
	InvRole *string `json:"InvRole,omitempty" xml:"InvRole,omitempty"`
	// example:
	//
	// test_prefix
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// example:
	//
	// test_region_id
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// test_role
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s AddressDetail) String() string {
	return tea.Prettify(s)
}

func (s AddressDetail) GoString() string {
	return s.String()
}

func (s *AddressDetail) SetAccessId(v string) *AddressDetail {
	s.AccessId = &v
	return s
}

func (s *AddressDetail) SetAccessSecret(v string) *AddressDetail {
	s.AccessSecret = &v
	return s
}

func (s *AddressDetail) SetAddressType(v string) *AddressDetail {
	s.AddressType = &v
	return s
}

func (s *AddressDetail) SetAgentList(v string) *AddressDetail {
	s.AgentList = &v
	return s
}

func (s *AddressDetail) SetBucket(v string) *AddressDetail {
	s.Bucket = &v
	return s
}

func (s *AddressDetail) SetDomain(v string) *AddressDetail {
	s.Domain = &v
	return s
}

func (s *AddressDetail) SetInvAccessId(v string) *AddressDetail {
	s.InvAccessId = &v
	return s
}

func (s *AddressDetail) SetInvAccessSecret(v string) *AddressDetail {
	s.InvAccessSecret = &v
	return s
}

func (s *AddressDetail) SetInvBucket(v string) *AddressDetail {
	s.InvBucket = &v
	return s
}

func (s *AddressDetail) SetInvDomain(v string) *AddressDetail {
	s.InvDomain = &v
	return s
}

func (s *AddressDetail) SetInvLocation(v string) *AddressDetail {
	s.InvLocation = &v
	return s
}

func (s *AddressDetail) SetInvPath(v string) *AddressDetail {
	s.InvPath = &v
	return s
}

func (s *AddressDetail) SetInvRegionId(v string) *AddressDetail {
	s.InvRegionId = &v
	return s
}

func (s *AddressDetail) SetInvRole(v string) *AddressDetail {
	s.InvRole = &v
	return s
}

func (s *AddressDetail) SetPrefix(v string) *AddressDetail {
	s.Prefix = &v
	return s
}

func (s *AddressDetail) SetRegionId(v string) *AddressDetail {
	s.RegionId = &v
	return s
}

func (s *AddressDetail) SetRole(v string) *AddressDetail {
	s.Role = &v
	return s
}

type Audit struct {
	// example:
	//
	// off
	LogMode *string `json:"LogMode,omitempty" xml:"LogMode,omitempty"`
}

func (s Audit) String() string {
	return tea.Prettify(s)
}

func (s Audit) GoString() string {
	return s.String()
}

func (s *Audit) SetLogMode(v string) *Audit {
	s.LogMode = &v
	return s
}

type CreateAddressInfo struct {
	// This parameter is required.
	AddressDetail *AddressDetail `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// K1:V1,K2:V2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateAddressInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAddressInfo) GoString() string {
	return s.String()
}

func (s *CreateAddressInfo) SetAddressDetail(v *AddressDetail) *CreateAddressInfo {
	s.AddressDetail = v
	return s
}

func (s *CreateAddressInfo) SetName(v string) *CreateAddressInfo {
	s.Name = &v
	return s
}

func (s *CreateAddressInfo) SetTags(v string) *CreateAddressInfo {
	s.Tags = &v
	return s
}

type CreateAgentInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// vpc
	AgentEndpoint *string `json:"AgentEndpoint,omitempty" xml:"AgentEndpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	DeployMethod *string `json:"DeployMethod,omitempty" xml:"DeployMethod,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// K1:V1,K2:V2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_tunnel_id
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
}

func (s CreateAgentInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentInfo) GoString() string {
	return s.String()
}

func (s *CreateAgentInfo) SetAgentEndpoint(v string) *CreateAgentInfo {
	s.AgentEndpoint = &v
	return s
}

func (s *CreateAgentInfo) SetDeployMethod(v string) *CreateAgentInfo {
	s.DeployMethod = &v
	return s
}

func (s *CreateAgentInfo) SetName(v string) *CreateAgentInfo {
	s.Name = &v
	return s
}

func (s *CreateAgentInfo) SetTags(v string) *CreateAgentInfo {
	s.Tags = &v
	return s
}

func (s *CreateAgentInfo) SetTunnelId(v string) *CreateAgentInfo {
	s.TunnelId = &v
	return s
}

type CreateJobInfo struct {
	Audit *Audit `json:"Audit,omitempty" xml:"Audit,omitempty"`
	// example:
	//
	// false
	ConvertSymlinkTarget *bool `json:"ConvertSymlinkTarget,omitempty" xml:"ConvertSymlinkTarget,omitempty"`
	CreateReport         *bool `json:"CreateReport,omitempty" xml:"CreateReport,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_dest_address
	DestAddress           *string     `json:"DestAddress,omitempty" xml:"DestAddress,omitempty"`
	EnableMultiVersioning *bool       `json:"EnableMultiVersioning,omitempty" xml:"EnableMultiVersioning,omitempty"`
	FilterRule            *FilterRule `json:"FilterRule,omitempty" xml:"FilterRule,omitempty"`
	ImportQos             *ImportQos  `json:"ImportQos,omitempty" xml:"ImportQos,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// always
	OverwriteMode *string       `json:"OverwriteMode,omitempty" xml:"OverwriteMode,omitempty"`
	ParentVersion *string       `json:"ParentVersion,omitempty" xml:"ParentVersion,omitempty"`
	ScheduleRule  *ScheduleRule `json:"ScheduleRule,omitempty" xml:"ScheduleRule,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_src_address
	SrcAddress *string `json:"SrcAddress,omitempty" xml:"SrcAddress,omitempty"`
	// example:
	//
	// K1:V1,K2:V2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// all
	TransferMode *string `json:"TransferMode,omitempty" xml:"TransferMode,omitempty"`
}

func (s CreateJobInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateJobInfo) GoString() string {
	return s.String()
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

func (s *CreateJobInfo) SetTransferMode(v string) *CreateJobInfo {
	s.TransferMode = &v
	return s
}

type CreateReportInfo struct {
	// example:
	//
	// test_job_name
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 1
	RuntimeId *int32 `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	// example:
	//
	// test_job_id
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateReportInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateReportInfo) GoString() string {
	return s.String()
}

func (s *CreateReportInfo) SetJobName(v string) *CreateReportInfo {
	s.JobName = &v
	return s
}

func (s *CreateReportInfo) SetRuntimeId(v int32) *CreateReportInfo {
	s.RuntimeId = &v
	return s
}

func (s *CreateReportInfo) SetVersion(v string) *CreateReportInfo {
	s.Version = &v
	return s
}

type CreateTunnelInfo struct {
	// example:
	//
	// K1:V1,K2:V2
	Tags      *string    `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TunnelQos *TunnelQos `json:"TunnelQos,omitempty" xml:"TunnelQos,omitempty"`
}

func (s CreateTunnelInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateTunnelInfo) GoString() string {
	return s.String()
}

func (s *CreateTunnelInfo) SetTags(v string) *CreateTunnelInfo {
	s.Tags = &v
	return s
}

func (s *CreateTunnelInfo) SetTunnelQos(v *TunnelQos) *CreateTunnelInfo {
	s.TunnelQos = v
	return s
}

type FileTypeFilters struct {
	// example:
	//
	// fasle
	ExcludeDir *bool `json:"ExcludeDir,omitempty" xml:"ExcludeDir,omitempty"`
	// example:
	//
	// fasle
	ExcludeSymlink *bool `json:"ExcludeSymlink,omitempty" xml:"ExcludeSymlink,omitempty"`
}

func (s FileTypeFilters) String() string {
	return tea.Prettify(s)
}

func (s FileTypeFilters) GoString() string {
	return s.String()
}

func (s *FileTypeFilters) SetExcludeDir(v bool) *FileTypeFilters {
	s.ExcludeDir = &v
	return s
}

func (s *FileTypeFilters) SetExcludeSymlink(v bool) *FileTypeFilters {
	s.ExcludeSymlink = &v
	return s
}

type FilterRule struct {
	FileTypeFilters     *FileTypeFilters     `json:"FileTypeFilters,omitempty" xml:"FileTypeFilters,omitempty"`
	KeyFilters          *KeyFilters          `json:"KeyFilters,omitempty" xml:"KeyFilters,omitempty"`
	LastModifiedFilters *LastModifiedFilters `json:"LastModifiedFilters,omitempty" xml:"LastModifiedFilters,omitempty"`
}

func (s FilterRule) String() string {
	return tea.Prettify(s)
}

func (s FilterRule) GoString() string {
	return s.String()
}

func (s *FilterRule) SetFileTypeFilters(v *FileTypeFilters) *FilterRule {
	s.FileTypeFilters = v
	return s
}

func (s *FilterRule) SetKeyFilters(v *KeyFilters) *FilterRule {
	s.KeyFilters = v
	return s
}

func (s *FilterRule) SetLastModifiedFilters(v *LastModifiedFilters) *FilterRule {
	s.LastModifiedFilters = v
	return s
}

type GetAddressResp struct {
	AddressDetail *AddressDetail `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// test_owner
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// avaliable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// K1:V1,K2:V2
	Tags         *string     `json:"Tags,omitempty" xml:"Tags,omitempty"`
	VerifyResult *VerifyResp `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	VerifyTime *string `json:"VerifyTime,omitempty" xml:"VerifyTime,omitempty"`
	// example:
	//
	// test_id
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAddressResp) String() string {
	return tea.Prettify(s)
}

func (s GetAddressResp) GoString() string {
	return s.String()
}

func (s *GetAddressResp) SetAddressDetail(v *AddressDetail) *GetAddressResp {
	s.AddressDetail = v
	return s
}

func (s *GetAddressResp) SetCreateTime(v string) *GetAddressResp {
	s.CreateTime = &v
	return s
}

func (s *GetAddressResp) SetModifyTime(v string) *GetAddressResp {
	s.ModifyTime = &v
	return s
}

func (s *GetAddressResp) SetName(v string) *GetAddressResp {
	s.Name = &v
	return s
}

func (s *GetAddressResp) SetOwner(v string) *GetAddressResp {
	s.Owner = &v
	return s
}

func (s *GetAddressResp) SetStatus(v string) *GetAddressResp {
	s.Status = &v
	return s
}

func (s *GetAddressResp) SetTags(v string) *GetAddressResp {
	s.Tags = &v
	return s
}

func (s *GetAddressResp) SetVerifyResult(v *VerifyResp) *GetAddressResp {
	s.VerifyResult = v
	return s
}

func (s *GetAddressResp) SetVerifyTime(v string) *GetAddressResp {
	s.VerifyTime = &v
	return s
}

func (s *GetAddressResp) SetVersion(v string) *GetAddressResp {
	s.Version = &v
	return s
}

type GetAgentResp struct {
	ActivationKey *string `json:"ActivationKey,omitempty" xml:"ActivationKey,omitempty"`
	// example:
	//
	// vpc
	AgentEndpoint *string `json:"AgentEndpoint,omitempty" xml:"AgentEndpoint,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// default
	DeployMethod *string `json:"DeployMethod,omitempty" xml:"DeployMethod,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// test_owner
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// K1:V1,K2:V2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// test_tunnel_id
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// example:
	//
	// test_agent_id
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAgentResp) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResp) GoString() string {
	return s.String()
}

func (s *GetAgentResp) SetActivationKey(v string) *GetAgentResp {
	s.ActivationKey = &v
	return s
}

func (s *GetAgentResp) SetAgentEndpoint(v string) *GetAgentResp {
	s.AgentEndpoint = &v
	return s
}

func (s *GetAgentResp) SetCreateTime(v string) *GetAgentResp {
	s.CreateTime = &v
	return s
}

func (s *GetAgentResp) SetDeployMethod(v string) *GetAgentResp {
	s.DeployMethod = &v
	return s
}

func (s *GetAgentResp) SetModifyTime(v string) *GetAgentResp {
	s.ModifyTime = &v
	return s
}

func (s *GetAgentResp) SetName(v string) *GetAgentResp {
	s.Name = &v
	return s
}

func (s *GetAgentResp) SetOwner(v string) *GetAgentResp {
	s.Owner = &v
	return s
}

func (s *GetAgentResp) SetTags(v string) *GetAgentResp {
	s.Tags = &v
	return s
}

func (s *GetAgentResp) SetTunnelId(v string) *GetAgentResp {
	s.TunnelId = &v
	return s
}

func (s *GetAgentResp) SetVersion(v string) *GetAgentResp {
	s.Version = &v
	return s
}

type GetAgentStatusResp struct {
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAgentStatusResp) String() string {
	return tea.Prettify(s)
}

func (s GetAgentStatusResp) GoString() string {
	return s.String()
}

func (s *GetAgentStatusResp) SetStatus(v string) *GetAgentStatusResp {
	s.Status = &v
	return s
}

type GetJobResp struct {
	Audit *Audit `json:"Audit,omitempty" xml:"Audit,omitempty"`
	// example:
	//
	// false
	ConvertSymlinkTarget *bool `json:"ConvertSymlinkTarget,omitempty" xml:"ConvertSymlinkTarget,omitempty"`
	// example:
	//
	// false
	CreateReport *bool `json:"CreateReport,omitempty" xml:"CreateReport,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// test_dest_address
	DestAddress *string `json:"DestAddress,omitempty" xml:"DestAddress,omitempty"`
	// example:
	//
	// false
	EnableMultiVersioning *bool       `json:"EnableMultiVersioning,omitempty" xml:"EnableMultiVersioning,omitempty"`
	FilterRule            *FilterRule `json:"FilterRule,omitempty" xml:"FilterRule,omitempty"`
	ImportQos             *ImportQos  `json:"ImportQos,omitempty" xml:"ImportQos,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// always
	OverwriteMode *string       `json:"OverwriteMode,omitempty" xml:"OverwriteMode,omitempty"`
	Owner         *string       `json:"Owner,omitempty" xml:"Owner,omitempty"`
	ParentName    *string       `json:"ParentName,omitempty" xml:"ParentName,omitempty"`
	ParentVersion *string       `json:"ParentVersion,omitempty" xml:"ParentVersion,omitempty"`
	ScheduleRule  *ScheduleRule `json:"ScheduleRule,omitempty" xml:"ScheduleRule,omitempty"`
	// example:
	//
	// test_src_address
	SrcAddress *string `json:"SrcAddress,omitempty" xml:"SrcAddress,omitempty"`
	// example:
	//
	// IMPORT_JOB_DOING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// K1:V1,K2:V2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// all
	TransferMode *string `json:"TransferMode,omitempty" xml:"TransferMode,omitempty"`
	// example:
	//
	// test_id
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetJobResp) String() string {
	return tea.Prettify(s)
}

func (s GetJobResp) GoString() string {
	return s.String()
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

func (s *GetJobResp) SetTransferMode(v string) *GetJobResp {
	s.TransferMode = &v
	return s
}

func (s *GetJobResp) SetVersion(v string) *GetJobResp {
	s.Version = &v
	return s
}

type GetJobResultResp struct {
	// example:
	//
	// ossinv
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// example:
	//
	// 800
	CopiedObjectCount *int64 `json:"CopiedObjectCount,omitempty" xml:"CopiedObjectCount,omitempty"`
	// example:
	//
	// 800
	CopiedObjectSize *int64 `json:"CopiedObjectSize,omitempty" xml:"CopiedObjectSize,omitempty"`
	// example:
	//
	// 200
	FailedObjectCount *int64 `json:"FailedObjectCount,omitempty" xml:"FailedObjectCount,omitempty"`
	// example:
	//
	// test_access_id
	InvAccessId *string `json:"InvAccessId,omitempty" xml:"InvAccessId,omitempty"`
	// example:
	//
	// test_secret_key
	InvAccessSecret *string `json:"InvAccessSecret,omitempty" xml:"InvAccessSecret,omitempty"`
	// example:
	//
	// test_sys_bucket
	InvBucket *string `json:"InvBucket,omitempty" xml:"InvBucket,omitempty"`
	// example:
	//
	// test_domain
	InvDomain *string `json:"InvDomain,omitempty" xml:"InvDomain,omitempty"`
	// example:
	//
	// oss
	InvLocation *string `json:"InvLocation,omitempty" xml:"InvLocation,omitempty"`
	// example:
	//
	// mainfest.json
	InvPath *string `json:"InvPath,omitempty" xml:"InvPath,omitempty"`
	// example:
	//
	// test_region_id
	InvRegionId *string `json:"InvRegionId,omitempty" xml:"InvRegionId,omitempty"`
	// example:
	//
	// Ready
	ReadyRetry *string `json:"ReadyRetry,omitempty" xml:"ReadyRetry,omitempty"`
	// example:
	//
	// 1000
	TotalObjectCount *int64 `json:"TotalObjectCount,omitempty" xml:"TotalObjectCount,omitempty"`
	// example:
	//
	// 1000
	TotalObjectSize *int64 `json:"TotalObjectSize,omitempty" xml:"TotalObjectSize,omitempty"`
	// example:
	//
	// test_job_id
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetJobResultResp) String() string {
	return tea.Prettify(s)
}

func (s GetJobResultResp) GoString() string {
	return s.String()
}

func (s *GetJobResultResp) SetAddressType(v string) *GetJobResultResp {
	s.AddressType = &v
	return s
}

func (s *GetJobResultResp) SetCopiedObjectCount(v int64) *GetJobResultResp {
	s.CopiedObjectCount = &v
	return s
}

func (s *GetJobResultResp) SetCopiedObjectSize(v int64) *GetJobResultResp {
	s.CopiedObjectSize = &v
	return s
}

func (s *GetJobResultResp) SetFailedObjectCount(v int64) *GetJobResultResp {
	s.FailedObjectCount = &v
	return s
}

func (s *GetJobResultResp) SetInvAccessId(v string) *GetJobResultResp {
	s.InvAccessId = &v
	return s
}

func (s *GetJobResultResp) SetInvAccessSecret(v string) *GetJobResultResp {
	s.InvAccessSecret = &v
	return s
}

func (s *GetJobResultResp) SetInvBucket(v string) *GetJobResultResp {
	s.InvBucket = &v
	return s
}

func (s *GetJobResultResp) SetInvDomain(v string) *GetJobResultResp {
	s.InvDomain = &v
	return s
}

func (s *GetJobResultResp) SetInvLocation(v string) *GetJobResultResp {
	s.InvLocation = &v
	return s
}

func (s *GetJobResultResp) SetInvPath(v string) *GetJobResultResp {
	s.InvPath = &v
	return s
}

func (s *GetJobResultResp) SetInvRegionId(v string) *GetJobResultResp {
	s.InvRegionId = &v
	return s
}

func (s *GetJobResultResp) SetReadyRetry(v string) *GetJobResultResp {
	s.ReadyRetry = &v
	return s
}

func (s *GetJobResultResp) SetTotalObjectCount(v int64) *GetJobResultResp {
	s.TotalObjectCount = &v
	return s
}

func (s *GetJobResultResp) SetTotalObjectSize(v int64) *GetJobResultResp {
	s.TotalObjectSize = &v
	return s
}

func (s *GetJobResultResp) SetVersion(v string) *GetJobResultResp {
	s.Version = &v
	return s
}

type GetReportResp struct {
	// example:
	//
	// 800
	CopiedCount  *int64  `json:"CopiedCount,omitempty" xml:"CopiedCount,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 100
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// example:
	//
	// test_failed_prefix/
	FailedListPrefix *string `json:"FailedListPrefix,omitempty" xml:"FailedListPrefix,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	JobCreateTime *string `json:"JobCreateTime,omitempty" xml:"JobCreateTime,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	JobEndTime *string `json:"JobEndTime,omitempty" xml:"JobEndTime,omitempty"`
	// example:
	//
	// 1000
	JobExecuteTime *string `json:"JobExecuteTime,omitempty" xml:"JobExecuteTime,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	ReportCreateTime *string `json:"ReportCreateTime,omitempty" xml:"ReportCreateTime,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	ReportEndTime *string `json:"ReportEndTime,omitempty" xml:"ReportEndTime,omitempty"`
	// example:
	//
	// 100
	SkippedCount *int64 `json:"SkippedCount,omitempty" xml:"SkippedCount,omitempty"`
	// example:
	//
	// test_skipped_prefix/
	SkippedListPrefix *string `json:"SkippedListPrefix,omitempty" xml:"SkippedListPrefix,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// test_total_prefix/
	TotalListPrefix *string `json:"TotalListPrefix,omitempty" xml:"TotalListPrefix,omitempty"`
}

func (s GetReportResp) String() string {
	return tea.Prettify(s)
}

func (s GetReportResp) GoString() string {
	return s.String()
}

func (s *GetReportResp) SetCopiedCount(v int64) *GetReportResp {
	s.CopiedCount = &v
	return s
}

func (s *GetReportResp) SetErrorMessage(v string) *GetReportResp {
	s.ErrorMessage = &v
	return s
}

func (s *GetReportResp) SetFailedCount(v int64) *GetReportResp {
	s.FailedCount = &v
	return s
}

func (s *GetReportResp) SetFailedListPrefix(v string) *GetReportResp {
	s.FailedListPrefix = &v
	return s
}

func (s *GetReportResp) SetJobCreateTime(v string) *GetReportResp {
	s.JobCreateTime = &v
	return s
}

func (s *GetReportResp) SetJobEndTime(v string) *GetReportResp {
	s.JobEndTime = &v
	return s
}

func (s *GetReportResp) SetJobExecuteTime(v string) *GetReportResp {
	s.JobExecuteTime = &v
	return s
}

func (s *GetReportResp) SetReportCreateTime(v string) *GetReportResp {
	s.ReportCreateTime = &v
	return s
}

func (s *GetReportResp) SetReportEndTime(v string) *GetReportResp {
	s.ReportEndTime = &v
	return s
}

func (s *GetReportResp) SetSkippedCount(v int64) *GetReportResp {
	s.SkippedCount = &v
	return s
}

func (s *GetReportResp) SetSkippedListPrefix(v string) *GetReportResp {
	s.SkippedListPrefix = &v
	return s
}

func (s *GetReportResp) SetStatus(v string) *GetReportResp {
	s.Status = &v
	return s
}

func (s *GetReportResp) SetTotalCount(v int64) *GetReportResp {
	s.TotalCount = &v
	return s
}

func (s *GetReportResp) SetTotalListPrefix(v string) *GetReportResp {
	s.TotalListPrefix = &v
	return s
}

type GetTunnelResp struct {
	// example:
	//
	// 2024-05-01 12:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test_owner
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// K1:V1,K2:V2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// test_tunnel_id
	TunnelId  *string    `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	TunnelQos *TunnelQos `json:"TunnelQos,omitempty" xml:"TunnelQos,omitempty"`
}

func (s GetTunnelResp) String() string {
	return tea.Prettify(s)
}

func (s GetTunnelResp) GoString() string {
	return s.String()
}

func (s *GetTunnelResp) SetCreateTime(v string) *GetTunnelResp {
	s.CreateTime = &v
	return s
}

func (s *GetTunnelResp) SetModifyTime(v string) *GetTunnelResp {
	s.ModifyTime = &v
	return s
}

func (s *GetTunnelResp) SetOwner(v string) *GetTunnelResp {
	s.Owner = &v
	return s
}

func (s *GetTunnelResp) SetTags(v string) *GetTunnelResp {
	s.Tags = &v
	return s
}

func (s *GetTunnelResp) SetTunnelId(v string) *GetTunnelResp {
	s.TunnelId = &v
	return s
}

func (s *GetTunnelResp) SetTunnelQos(v *TunnelQos) *GetTunnelResp {
	s.TunnelQos = v
	return s
}

type ImportQos struct {
	// example:
	//
	// 1073741824
	MaxBandWidth *int64 `json:"MaxBandWidth,omitempty" xml:"MaxBandWidth,omitempty"`
	// example:
	//
	// 1000
	MaxImportTaskQps *int64 `json:"MaxImportTaskQps,omitempty" xml:"MaxImportTaskQps,omitempty"`
}

func (s ImportQos) String() string {
	return tea.Prettify(s)
}

func (s ImportQos) GoString() string {
	return s.String()
}

func (s *ImportQos) SetMaxBandWidth(v int64) *ImportQos {
	s.MaxBandWidth = &v
	return s
}

func (s *ImportQos) SetMaxImportTaskQps(v int64) *ImportQos {
	s.MaxImportTaskQps = &v
	return s
}

type JobHistory struct {
	// example:
	//
	// 2
	CommitId *string `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	// example:
	//
	// 900
	CopiedCount *int64 `json:"CopiedCount,omitempty" xml:"CopiedCount,omitempty"`
	// example:
	//
	// 1000
	CopiedSize *int64 `json:"CopiedSize,omitempty" xml:"CopiedSize,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 100
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// example:
	//
	// test_id
	JobVersion *string `json:"JobVersion,omitempty" xml:"JobVersion,omitempty"`
	// example:
	//
	// Listing
	ListStatus *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// test_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// user
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 1
	RuntimeId *string `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	// example:
	//
	// Normal
	RuntimeState *string `json:"RuntimeState,omitempty" xml:"RuntimeState,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// IMPORT_JOB_DOING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1000
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s JobHistory) String() string {
	return tea.Prettify(s)
}

func (s JobHistory) GoString() string {
	return s.String()
}

func (s *JobHistory) SetCommitId(v string) *JobHistory {
	s.CommitId = &v
	return s
}

func (s *JobHistory) SetCopiedCount(v int64) *JobHistory {
	s.CopiedCount = &v
	return s
}

func (s *JobHistory) SetCopiedSize(v int64) *JobHistory {
	s.CopiedSize = &v
	return s
}

func (s *JobHistory) SetEndTime(v string) *JobHistory {
	s.EndTime = &v
	return s
}

func (s *JobHistory) SetFailedCount(v int64) *JobHistory {
	s.FailedCount = &v
	return s
}

func (s *JobHistory) SetJobVersion(v string) *JobHistory {
	s.JobVersion = &v
	return s
}

func (s *JobHistory) SetListStatus(v string) *JobHistory {
	s.ListStatus = &v
	return s
}

func (s *JobHistory) SetMessage(v string) *JobHistory {
	s.Message = &v
	return s
}

func (s *JobHistory) SetName(v string) *JobHistory {
	s.Name = &v
	return s
}

func (s *JobHistory) SetOperator(v string) *JobHistory {
	s.Operator = &v
	return s
}

func (s *JobHistory) SetRuntimeId(v string) *JobHistory {
	s.RuntimeId = &v
	return s
}

func (s *JobHistory) SetRuntimeState(v string) *JobHistory {
	s.RuntimeState = &v
	return s
}

func (s *JobHistory) SetStartTime(v string) *JobHistory {
	s.StartTime = &v
	return s
}

func (s *JobHistory) SetStatus(v string) *JobHistory {
	s.Status = &v
	return s
}

func (s *JobHistory) SetTotalCount(v int64) *JobHistory {
	s.TotalCount = &v
	return s
}

func (s *JobHistory) SetTotalSize(v int64) *JobHistory {
	s.TotalSize = &v
	return s
}

type KeyFilterItem struct {
	Regex []*string `json:"Regex,omitempty" xml:"Regex,omitempty" type:"Repeated"`
}

func (s KeyFilterItem) String() string {
	return tea.Prettify(s)
}

func (s KeyFilterItem) GoString() string {
	return s.String()
}

func (s *KeyFilterItem) SetRegex(v []*string) *KeyFilterItem {
	s.Regex = v
	return s
}

type KeyFilters struct {
	Excludes *KeyFilterItem `json:"Excludes,omitempty" xml:"Excludes,omitempty"`
	Includes *KeyFilterItem `json:"Includes,omitempty" xml:"Includes,omitempty"`
}

func (s KeyFilters) String() string {
	return tea.Prettify(s)
}

func (s KeyFilters) GoString() string {
	return s.String()
}

func (s *KeyFilters) SetExcludes(v *KeyFilterItem) *KeyFilters {
	s.Excludes = v
	return s
}

func (s *KeyFilters) SetIncludes(v *KeyFilterItem) *KeyFilters {
	s.Includes = v
	return s
}

type LastModifiedFilters struct {
	Excludes *LastModifyFilterItem `json:"Excludes,omitempty" xml:"Excludes,omitempty"`
	Includes *LastModifyFilterItem `json:"Includes,omitempty" xml:"Includes,omitempty"`
}

func (s LastModifiedFilters) String() string {
	return tea.Prettify(s)
}

func (s LastModifiedFilters) GoString() string {
	return s.String()
}

func (s *LastModifiedFilters) SetExcludes(v *LastModifyFilterItem) *LastModifiedFilters {
	s.Excludes = v
	return s
}

func (s *LastModifiedFilters) SetIncludes(v *LastModifyFilterItem) *LastModifiedFilters {
	s.Includes = v
	return s
}

type LastModifyFilterItem struct {
	TimeFilter []*TimeFilter `json:"TimeFilter,omitempty" xml:"TimeFilter,omitempty" type:"Repeated"`
}

func (s LastModifyFilterItem) String() string {
	return tea.Prettify(s)
}

func (s LastModifyFilterItem) GoString() string {
	return s.String()
}

func (s *LastModifyFilterItem) SetTimeFilter(v []*TimeFilter) *LastModifyFilterItem {
	s.TimeFilter = v
	return s
}

type ListAddressResp struct {
	ImportAddress []*GetAddressResp `json:"ImportAddress,omitempty" xml:"ImportAddress,omitempty" type:"Repeated"`
	// example:
	//
	// test_marker
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	// example:
	//
	// true
	Truncated *bool `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s ListAddressResp) String() string {
	return tea.Prettify(s)
}

func (s ListAddressResp) GoString() string {
	return s.String()
}

func (s *ListAddressResp) SetImportAddress(v []*GetAddressResp) *ListAddressResp {
	s.ImportAddress = v
	return s
}

func (s *ListAddressResp) SetNextMarker(v string) *ListAddressResp {
	s.NextMarker = &v
	return s
}

func (s *ListAddressResp) SetTruncated(v bool) *ListAddressResp {
	s.Truncated = &v
	return s
}

type ListAgentResp struct {
	ImportAgent []*GetAgentResp `json:"ImportAgent,omitempty" xml:"ImportAgent,omitempty" type:"Repeated"`
	// example:
	//
	// test_next_marker
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	// example:
	//
	// true
	Truncated *bool `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s ListAgentResp) String() string {
	return tea.Prettify(s)
}

func (s ListAgentResp) GoString() string {
	return s.String()
}

func (s *ListAgentResp) SetImportAgent(v []*GetAgentResp) *ListAgentResp {
	s.ImportAgent = v
	return s
}

func (s *ListAgentResp) SetNextMarker(v string) *ListAgentResp {
	s.NextMarker = &v
	return s
}

func (s *ListAgentResp) SetTruncated(v bool) *ListAgentResp {
	s.Truncated = &v
	return s
}

type ListJobHistoryResp struct {
	JobHistory []*JobHistory `json:"JobHistory,omitempty" xml:"JobHistory,omitempty" type:"Repeated"`
	// example:
	//
	// test_next_marker
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	// example:
	//
	// true
	Truncated *string `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s ListJobHistoryResp) String() string {
	return tea.Prettify(s)
}

func (s ListJobHistoryResp) GoString() string {
	return s.String()
}

func (s *ListJobHistoryResp) SetJobHistory(v []*JobHistory) *ListJobHistoryResp {
	s.JobHistory = v
	return s
}

func (s *ListJobHistoryResp) SetNextMarker(v string) *ListJobHistoryResp {
	s.NextMarker = &v
	return s
}

func (s *ListJobHistoryResp) SetTruncated(v string) *ListJobHistoryResp {
	s.Truncated = &v
	return s
}

type ListJobInfo struct {
	ImportJob []*CreateJobInfo `json:"ImportJob,omitempty" xml:"ImportJob,omitempty" type:"Repeated"`
	// example:
	//
	// test_next_marker
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	// example:
	//
	// true
	Truncated *bool `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s ListJobInfo) String() string {
	return tea.Prettify(s)
}

func (s ListJobInfo) GoString() string {
	return s.String()
}

func (s *ListJobInfo) SetImportJob(v []*CreateJobInfo) *ListJobInfo {
	s.ImportJob = v
	return s
}

func (s *ListJobInfo) SetNextMarker(v string) *ListJobInfo {
	s.NextMarker = &v
	return s
}

func (s *ListJobInfo) SetTruncated(v bool) *ListJobInfo {
	s.Truncated = &v
	return s
}

type ListJobResp struct {
	ImportJob  []*GetJobResp `json:"ImportJob,omitempty" xml:"ImportJob,omitempty" type:"Repeated"`
	NextMarker *string       `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	Truncated  *bool         `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s ListJobResp) String() string {
	return tea.Prettify(s)
}

func (s ListJobResp) GoString() string {
	return s.String()
}

func (s *ListJobResp) SetImportJob(v []*GetJobResp) *ListJobResp {
	s.ImportJob = v
	return s
}

func (s *ListJobResp) SetNextMarker(v string) *ListJobResp {
	s.NextMarker = &v
	return s
}

func (s *ListJobResp) SetTruncated(v bool) *ListJobResp {
	s.Truncated = &v
	return s
}

type ListTunnelResp struct {
	ImportTunnel []*GetTunnelResp `json:"ImportTunnel,omitempty" xml:"ImportTunnel,omitempty" type:"Repeated"`
	NextMarker   *string          `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	Truncated    *bool            `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s ListTunnelResp) String() string {
	return tea.Prettify(s)
}

func (s ListTunnelResp) GoString() string {
	return s.String()
}

func (s *ListTunnelResp) SetImportTunnel(v []*GetTunnelResp) *ListTunnelResp {
	s.ImportTunnel = v
	return s
}

func (s *ListTunnelResp) SetNextMarker(v string) *ListTunnelResp {
	s.NextMarker = &v
	return s
}

func (s *ListTunnelResp) SetTruncated(v bool) *ListTunnelResp {
	s.Truncated = &v
	return s
}

type ScheduleRule struct {
	MaxScheduleCount      *int64  `json:"MaxScheduleCount,omitempty" xml:"MaxScheduleCount,omitempty"`
	StartCronExpression   *string `json:"StartCronExpression,omitempty" xml:"StartCronExpression,omitempty"`
	SuspendCronExpression *string `json:"SuspendCronExpression,omitempty" xml:"SuspendCronExpression,omitempty"`
}

func (s ScheduleRule) String() string {
	return tea.Prettify(s)
}

func (s ScheduleRule) GoString() string {
	return s.String()
}

func (s *ScheduleRule) SetMaxScheduleCount(v int64) *ScheduleRule {
	s.MaxScheduleCount = &v
	return s
}

func (s *ScheduleRule) SetStartCronExpression(v string) *ScheduleRule {
	s.StartCronExpression = &v
	return s
}

func (s *ScheduleRule) SetSuspendCronExpression(v string) *ScheduleRule {
	s.SuspendCronExpression = &v
	return s
}

type TimeFilter struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s TimeFilter) String() string {
	return tea.Prettify(s)
}

func (s TimeFilter) GoString() string {
	return s.String()
}

func (s *TimeFilter) SetEndTime(v string) *TimeFilter {
	s.EndTime = &v
	return s
}

func (s *TimeFilter) SetStartTime(v string) *TimeFilter {
	s.StartTime = &v
	return s
}

type TunnelQos struct {
	// example:
	//
	// 1073741824
	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" xml:"MaxBandwidth,omitempty"`
	// example:
	//
	// 100
	MaxQps *int32 `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
}

func (s TunnelQos) String() string {
	return tea.Prettify(s)
}

func (s TunnelQos) GoString() string {
	return s.String()
}

func (s *TunnelQos) SetMaxBandwidth(v int64) *TunnelQos {
	s.MaxBandwidth = &v
	return s
}

func (s *TunnelQos) SetMaxQps(v int32) *TunnelQos {
	s.MaxQps = &v
	return s
}

type UpdateAddressInfo struct {
	AgentList *string `json:"AgentList,omitempty" xml:"AgentList,omitempty"`
}

func (s UpdateAddressInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateAddressInfo) GoString() string {
	return s.String()
}

func (s *UpdateAddressInfo) SetAgentList(v string) *UpdateAddressInfo {
	s.AgentList = &v
	return s
}

type UpdateJobInfo struct {
	ImportQos *ImportQos `json:"ImportQos,omitempty" xml:"ImportQos,omitempty"`
	// example:
	//
	// IMPORT_JOB_LAUNCHING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateJobInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobInfo) GoString() string {
	return s.String()
}

func (s *UpdateJobInfo) SetImportQos(v *ImportQos) *UpdateJobInfo {
	s.ImportQos = v
	return s
}

func (s *UpdateJobInfo) SetStatus(v string) *UpdateJobInfo {
	s.Status = &v
	return s
}

type UpdateTunnelInfo struct {
	// example:
	//
	// k1=v1;k2=v2
	Tags      *string    `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TunnelQos *TunnelQos `json:"TunnelQos,omitempty" xml:"TunnelQos,omitempty"`
}

func (s UpdateTunnelInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateTunnelInfo) GoString() string {
	return s.String()
}

func (s *UpdateTunnelInfo) SetTags(v string) *UpdateTunnelInfo {
	s.Tags = &v
	return s
}

func (s *UpdateTunnelInfo) SetTunnelQos(v *TunnelQos) *UpdateTunnelInfo {
	s.TunnelQos = v
	return s
}

type VerifyAddressResp struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// avaliable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	VerifyTime *string `json:"VerifyTime,omitempty" xml:"VerifyTime,omitempty"`
}

func (s VerifyAddressResp) String() string {
	return tea.Prettify(s)
}

func (s VerifyAddressResp) GoString() string {
	return s.String()
}

func (s *VerifyAddressResp) SetErrorCode(v string) *VerifyAddressResp {
	s.ErrorCode = &v
	return s
}

func (s *VerifyAddressResp) SetErrorMessage(v string) *VerifyAddressResp {
	s.ErrorMessage = &v
	return s
}

func (s *VerifyAddressResp) SetStatus(v string) *VerifyAddressResp {
	s.Status = &v
	return s
}

func (s *VerifyAddressResp) SetVerifyTime(v string) *VerifyAddressResp {
	s.VerifyTime = &v
	return s
}

type VerifyResp struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 200
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
}

func (s VerifyResp) String() string {
	return tea.Prettify(s)
}

func (s VerifyResp) GoString() string {
	return s.String()
}

func (s *VerifyResp) SetErrorCode(v string) *VerifyResp {
	s.ErrorCode = &v
	return s
}

func (s *VerifyResp) SetErrorMsg(v string) *VerifyResp {
	s.ErrorMsg = &v
	return s
}

func (s *VerifyResp) SetHttpCode(v string) *VerifyResp {
	s.HttpCode = &v
	return s
}

type CreateAddressRequest struct {
	// The details for creating the data address.
	ImportAddress *CreateAddressInfo `json:"ImportAddress,omitempty" xml:"ImportAddress,omitempty"`
}

func (s CreateAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateAddressRequest) SetImportAddress(v *CreateAddressInfo) *CreateAddressRequest {
	s.ImportAddress = v
	return s
}

type CreateAddressResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateAddressResponse) SetHeaders(v map[string]*string) *CreateAddressResponse {
	s.Headers = v
	return s
}

func (s *CreateAddressResponse) SetStatusCode(v int32) *CreateAddressResponse {
	s.StatusCode = &v
	return s
}

type CreateAgentRequest struct {
	// The details for creating the agent.
	ImportAgent *CreateAgentInfo `json:"ImportAgent,omitempty" xml:"ImportAgent,omitempty"`
}

func (s CreateAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentRequest) SetImportAgent(v *CreateAgentInfo) *CreateAgentRequest {
	s.ImportAgent = v
	return s
}

type CreateAgentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentResponse) SetHeaders(v map[string]*string) *CreateAgentResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentResponse) SetStatusCode(v int32) *CreateAgentResponse {
	s.StatusCode = &v
	return s
}

type CreateJobRequest struct {
	// The details for creating the migration task.
	//
	// This parameter is required.
	ImportJob *CreateJobInfo `json:"ImportJob,omitempty" xml:"ImportJob,omitempty"`
}

func (s CreateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) SetImportJob(v *CreateJobInfo) *CreateJobRequest {
	s.ImportJob = v
	return s
}

type CreateJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponse) GoString() string {
	return s.String()
}

func (s *CreateJobResponse) SetHeaders(v map[string]*string) *CreateJobResponse {
	s.Headers = v
	return s
}

func (s *CreateJobResponse) SetStatusCode(v int32) *CreateJobResponse {
	s.StatusCode = &v
	return s
}

type CreateReportRequest struct {
	// The details for creating the migration report.
	CreateReport *CreateReportInfo `json:"CreateReport,omitempty" xml:"CreateReport,omitempty"`
}

func (s CreateReportRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateReportRequest) GoString() string {
	return s.String()
}

func (s *CreateReportRequest) SetCreateReport(v *CreateReportInfo) *CreateReportRequest {
	s.CreateReport = v
	return s
}

type CreateReportResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateReportResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateReportResponse) GoString() string {
	return s.String()
}

func (s *CreateReportResponse) SetHeaders(v map[string]*string) *CreateReportResponse {
	s.Headers = v
	return s
}

func (s *CreateReportResponse) SetStatusCode(v int32) *CreateReportResponse {
	s.StatusCode = &v
	return s
}

type CreateTunnelRequest struct {
	// The details for creating the tunnel.
	ImportTunnel *CreateTunnelInfo `json:"ImportTunnel,omitempty" xml:"ImportTunnel,omitempty"`
}

func (s CreateTunnelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTunnelRequest) GoString() string {
	return s.String()
}

func (s *CreateTunnelRequest) SetImportTunnel(v *CreateTunnelInfo) *CreateTunnelRequest {
	s.ImportTunnel = v
	return s
}

type CreateTunnelResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateTunnelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTunnelResponse) GoString() string {
	return s.String()
}

func (s *CreateTunnelResponse) SetHeaders(v map[string]*string) *CreateTunnelResponse {
	s.Headers = v
	return s
}

func (s *CreateTunnelResponse) SetStatusCode(v int32) *CreateTunnelResponse {
	s.StatusCode = &v
	return s
}

type DeleteAddressResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAddressResponse) GoString() string {
	return s.String()
}

func (s *DeleteAddressResponse) SetHeaders(v map[string]*string) *DeleteAddressResponse {
	s.Headers = v
	return s
}

func (s *DeleteAddressResponse) SetStatusCode(v int32) *DeleteAddressResponse {
	s.StatusCode = &v
	return s
}

type DeleteAgentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAgentResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentResponse) SetHeaders(v map[string]*string) *DeleteAgentResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentResponse) SetStatusCode(v int32) *DeleteAgentResponse {
	s.StatusCode = &v
	return s
}

type DeleteJobRequest struct {
	// Specifies whether to force delete the subtask. If the task has subtasks and you set this parameter to true, the task and its subtasks are forcibly deleted. If this parameter is set to false, the task and its subtasks fail to be deleted.
	//
	// example:
	//
	// true
	ForceDelete *bool `json:"forceDelete,omitempty" xml:"forceDelete,omitempty"`
}

func (s DeleteJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobRequest) SetForceDelete(v bool) *DeleteJobRequest {
	s.ForceDelete = &v
	return s
}

type DeleteJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobResponse) SetHeaders(v map[string]*string) *DeleteJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobResponse) SetStatusCode(v int32) *DeleteJobResponse {
	s.StatusCode = &v
	return s
}

type DeleteTunnelResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteTunnelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTunnelResponse) GoString() string {
	return s.String()
}

func (s *DeleteTunnelResponse) SetHeaders(v map[string]*string) *DeleteTunnelResponse {
	s.Headers = v
	return s
}

func (s *DeleteTunnelResponse) SetStatusCode(v int32) *DeleteTunnelResponse {
	s.StatusCode = &v
	return s
}

type GetAddressResponseBody struct {
	// The details for obtaining the data address.
	//
	// Valid values:
	//
	// 	- 1:1
	ImportAddress *GetAddressResp `json:"ImportAddress,omitempty" xml:"ImportAddress,omitempty"`
}

func (s GetAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAddressResponseBody) GoString() string {
	return s.String()
}

func (s *GetAddressResponseBody) SetImportAddress(v *GetAddressResp) *GetAddressResponseBody {
	s.ImportAddress = v
	return s
}

type GetAddressResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAddressResponse) GoString() string {
	return s.String()
}

func (s *GetAddressResponse) SetHeaders(v map[string]*string) *GetAddressResponse {
	s.Headers = v
	return s
}

func (s *GetAddressResponse) SetStatusCode(v int32) *GetAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAddressResponse) SetBody(v *GetAddressResponseBody) *GetAddressResponse {
	s.Body = v
	return s
}

type GetAgentResponseBody struct {
	// The details for obtaining the details of the agent.
	ImportAgent *GetAgentResp `json:"ImportAgent,omitempty" xml:"ImportAgent,omitempty"`
}

func (s GetAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBody) SetImportAgent(v *GetAgentResp) *GetAgentResponseBody {
	s.ImportAgent = v
	return s
}

type GetAgentResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponse) GoString() string {
	return s.String()
}

func (s *GetAgentResponse) SetHeaders(v map[string]*string) *GetAgentResponse {
	s.Headers = v
	return s
}

func (s *GetAgentResponse) SetStatusCode(v int32) *GetAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentResponse) SetBody(v *GetAgentResponseBody) *GetAgentResponse {
	s.Body = v
	return s
}

type GetAgentStatusResponseBody struct {
	// The details for obtaining the status of the agent.
	ImportAgentStatus *GetAgentStatusResp `json:"ImportAgentStatus,omitempty" xml:"ImportAgentStatus,omitempty"`
}

func (s GetAgentStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentStatusResponseBody) SetImportAgentStatus(v *GetAgentStatusResp) *GetAgentStatusResponseBody {
	s.ImportAgentStatus = v
	return s
}

type GetAgentStatusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAgentStatusResponse) SetHeaders(v map[string]*string) *GetAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAgentStatusResponse) SetStatusCode(v int32) *GetAgentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentStatusResponse) SetBody(v *GetAgentStatusResponseBody) *GetAgentStatusResponse {
	s.Body = v
	return s
}

type GetJobRequest struct {
	// Specifies whether to obtain the details of the migration task by using the task ID.
	//
	// example:
	//
	// false
	ByVersion *string `json:"byVersion,omitempty" xml:"byVersion,omitempty"`
}

func (s GetJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) SetByVersion(v string) *GetJobRequest {
	s.ByVersion = &v
	return s
}

type GetJobResponseBody struct {
	// The details for obtaining the details of the migration task.
	ImportJob *GetJobResp `json:"ImportJob,omitempty" xml:"ImportJob,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) SetImportJob(v *GetJobResp) *GetJobResponseBody {
	s.ImportJob = v
	return s
}

type GetJobResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponse) GoString() string {
	return s.String()
}

func (s *GetJobResponse) SetHeaders(v map[string]*string) *GetJobResponse {
	s.Headers = v
	return s
}

func (s *GetJobResponse) SetStatusCode(v int32) *GetJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobResponse) SetBody(v *GetJobResponseBody) *GetJobResponse {
	s.Body = v
	return s
}

type GetJobResultRequest struct {
	// The execution ID of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	RuntimeId *int32 `json:"runtimeId,omitempty" xml:"runtimeId,omitempty"`
}

func (s GetJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetJobResultRequest) SetRuntimeId(v int32) *GetJobResultRequest {
	s.RuntimeId = &v
	return s
}

type GetJobResultResponseBody struct {
	// The details for obtaining the retries of the migration task.
	ImportJobResult *GetJobResultResp `json:"ImportJobResult,omitempty" xml:"ImportJobResult,omitempty"`
}

func (s GetJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResultResponseBody) SetImportJobResult(v *GetJobResultResp) *GetJobResultResponseBody {
	s.ImportJobResult = v
	return s
}

type GetJobResultResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetJobResultResponse) SetHeaders(v map[string]*string) *GetJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetJobResultResponse) SetStatusCode(v int32) *GetJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobResultResponse) SetBody(v *GetJobResultResponseBody) *GetJobResultResponse {
	s.Body = v
	return s
}

type GetReportRequest struct {
	// The execution ID of the migration task.
	//
	// example:
	//
	// 1
	RuntimeId *int32 `json:"runtimeId,omitempty" xml:"runtimeId,omitempty"`
	// The ID of the migration task.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_job_id
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetReportRequest) GoString() string {
	return s.String()
}

func (s *GetReportRequest) SetRuntimeId(v int32) *GetReportRequest {
	s.RuntimeId = &v
	return s
}

func (s *GetReportRequest) SetVersion(v string) *GetReportRequest {
	s.Version = &v
	return s
}

type GetReportResponseBody struct {
	// The details for obtaining the migration report.
	GetReportResponse *GetReportResp `json:"GetReportResponse,omitempty" xml:"GetReportResponse,omitempty"`
}

func (s GetReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetReportResponseBody) SetGetReportResponse(v *GetReportResp) *GetReportResponseBody {
	s.GetReportResponse = v
	return s
}

type GetReportResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetReportResponse) GoString() string {
	return s.String()
}

func (s *GetReportResponse) SetHeaders(v map[string]*string) *GetReportResponse {
	s.Headers = v
	return s
}

func (s *GetReportResponse) SetStatusCode(v int32) *GetReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReportResponse) SetBody(v *GetReportResponseBody) *GetReportResponse {
	s.Body = v
	return s
}

type GetTunnelResponseBody struct {
	// The details for obtaining the details of the tunnel.
	ImportTunnel *GetTunnelResp `json:"ImportTunnel,omitempty" xml:"ImportTunnel,omitempty"`
}

func (s GetTunnelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTunnelResponseBody) GoString() string {
	return s.String()
}

func (s *GetTunnelResponseBody) SetImportTunnel(v *GetTunnelResp) *GetTunnelResponseBody {
	s.ImportTunnel = v
	return s
}

type GetTunnelResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTunnelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTunnelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTunnelResponse) GoString() string {
	return s.String()
}

func (s *GetTunnelResponse) SetHeaders(v map[string]*string) *GetTunnelResponse {
	s.Headers = v
	return s
}

func (s *GetTunnelResponse) SetStatusCode(v int32) *GetTunnelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTunnelResponse) SetBody(v *GetTunnelResponseBody) *GetTunnelResponse {
	s.Body = v
	return s
}

type ListAddressRequest struct {
	// Specifies the number of migration addresses to be returned.\\
	//
	// Valid values: 0 - 1000 (excluding 0).\\
	//
	// Default value: 1000.
	//
	// example:
	//
	// 100
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The marker after which the migration addresses are listed.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// test_marker
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAddressRequest) GoString() string {
	return s.String()
}

func (s *ListAddressRequest) SetCount(v int32) *ListAddressRequest {
	s.Count = &v
	return s
}

func (s *ListAddressRequest) SetMarker(v string) *ListAddressRequest {
	s.Marker = &v
	return s
}

type ListAddressResponseBody struct {
	// The details of migration addresses.
	ImportAddressList *ListAddressResp `json:"ImportAddressList,omitempty" xml:"ImportAddressList,omitempty"`
}

func (s ListAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddressResponseBody) SetImportAddressList(v *ListAddressResp) *ListAddressResponseBody {
	s.ImportAddressList = v
	return s
}

type ListAddressResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAddressResponse) GoString() string {
	return s.String()
}

func (s *ListAddressResponse) SetHeaders(v map[string]*string) *ListAddressResponse {
	s.Headers = v
	return s
}

func (s *ListAddressResponse) SetStatusCode(v int32) *ListAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAddressResponse) SetBody(v *ListAddressResponseBody) *ListAddressResponse {
	s.Body = v
	return s
}

type ListAgentRequest struct {
	// Specifies the number of agents to be returned.\\
	//
	// Valid values: 0 - 1000.\\
	//
	// Default value: 1000.
	//
	// example:
	//
	// 100
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The marker after which the agents are listed.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// test_agent
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAgentRequest) GoString() string {
	return s.String()
}

func (s *ListAgentRequest) SetCount(v int32) *ListAgentRequest {
	s.Count = &v
	return s
}

func (s *ListAgentRequest) SetMarker(v string) *ListAgentRequest {
	s.Marker = &v
	return s
}

type ListAgentResponseBody struct {
	// The details of the agents.
	ImportAgentList *ListAgentResp `json:"ImportAgentList,omitempty" xml:"ImportAgentList,omitempty"`
}

func (s ListAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAgentResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentResponseBody) SetImportAgentList(v *ListAgentResp) *ListAgentResponseBody {
	s.ImportAgentList = v
	return s
}

type ListAgentResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAgentResponse) GoString() string {
	return s.String()
}

func (s *ListAgentResponse) SetHeaders(v map[string]*string) *ListAgentResponse {
	s.Headers = v
	return s
}

func (s *ListAgentResponse) SetStatusCode(v int32) *ListAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentResponse) SetBody(v *ListAgentResponseBody) *ListAgentResponse {
	s.Body = v
	return s
}

type ListJobRequest struct {
	// Specifies whether to return subtasks.\\
	//
	// Valid values: true and false.
	//
	// example:
	//
	// true
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// Specifies the number of migration tasks to be returned.\\
	//
	// Valid values: 0 - 1000 (excluding 0).\\
	//
	// Default value: 1000.
	//
	// example:
	//
	// 1000
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The marker after which the migration tasks are listed.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// test_marker
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The name of the parent task. If this parameter is specified, all subtasks of the parent task are returned.
	//
	// example:
	//
	// test_parent_job_name
	ParentName *string `json:"parentName,omitempty" xml:"parentName,omitempty"`
}

func (s ListJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobRequest) GoString() string {
	return s.String()
}

func (s *ListJobRequest) SetAll(v bool) *ListJobRequest {
	s.All = &v
	return s
}

func (s *ListJobRequest) SetCount(v int32) *ListJobRequest {
	s.Count = &v
	return s
}

func (s *ListJobRequest) SetMarker(v string) *ListJobRequest {
	s.Marker = &v
	return s
}

func (s *ListJobRequest) SetParentName(v string) *ListJobRequest {
	s.ParentName = &v
	return s
}

type ListJobResponseBody struct {
	// The queried migration tasks.
	ImportJobList *ListJobResp `json:"ImportJobList,omitempty" xml:"ImportJobList,omitempty"`
}

func (s ListJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobResponseBody) SetImportJobList(v *ListJobResp) *ListJobResponseBody {
	s.ImportJobList = v
	return s
}

type ListJobResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobResponse) GoString() string {
	return s.String()
}

func (s *ListJobResponse) SetHeaders(v map[string]*string) *ListJobResponse {
	s.Headers = v
	return s
}

func (s *ListJobResponse) SetStatusCode(v int32) *ListJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobResponse) SetBody(v *ListJobResponseBody) *ListJobResponse {
	s.Body = v
	return s
}

type ListJobHistoryRequest struct {
	// Specifies the number of running records of the migration task to be returned.\\
	//
	// Valid values: 0 - 1000.\\
	//
	// Default value: 1000.
	//
	// example:
	//
	// 100
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The marker after which the running history of the task is listed.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// test_marker
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The execution ID of the task. If you specify an execution ID, only the running history related to the execution ID is listed.
	//
	// example:
	//
	// 1
	RuntimeId *int32 `json:"runtimeId,omitempty" xml:"runtimeId,omitempty"`
}

func (s ListJobHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListJobHistoryRequest) SetCount(v int32) *ListJobHistoryRequest {
	s.Count = &v
	return s
}

func (s *ListJobHistoryRequest) SetMarker(v string) *ListJobHistoryRequest {
	s.Marker = &v
	return s
}

func (s *ListJobHistoryRequest) SetRuntimeId(v int32) *ListJobHistoryRequest {
	s.RuntimeId = &v
	return s
}

type ListJobHistoryResponseBody struct {
	// The running history of the migration task.
	JobHistoryList *ListJobHistoryResp `json:"JobHistoryList,omitempty" xml:"JobHistoryList,omitempty"`
}

func (s ListJobHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobHistoryResponseBody) SetJobHistoryList(v *ListJobHistoryResp) *ListJobHistoryResponseBody {
	s.JobHistoryList = v
	return s
}

type ListJobHistoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListJobHistoryResponse) SetHeaders(v map[string]*string) *ListJobHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListJobHistoryResponse) SetStatusCode(v int32) *ListJobHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobHistoryResponse) SetBody(v *ListJobHistoryResponseBody) *ListJobHistoryResponse {
	s.Body = v
	return s
}

type ListTunnelRequest struct {
	// Specifies the number of tunnels to be returned.\\
	//
	// Valid values: 0 - 1000.\\
	//
	// Default value: 1000.
	//
	// example:
	//
	// 2
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The marker after which tunnels are listed.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// 1
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListTunnelRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTunnelRequest) GoString() string {
	return s.String()
}

func (s *ListTunnelRequest) SetCount(v int32) *ListTunnelRequest {
	s.Count = &v
	return s
}

func (s *ListTunnelRequest) SetMarker(v string) *ListTunnelRequest {
	s.Marker = &v
	return s
}

type ListTunnelResponseBody struct {
	// The details of the tunnels.
	ImportTunnelList *ListTunnelResp `json:"ImportTunnelList,omitempty" xml:"ImportTunnelList,omitempty"`
}

func (s ListTunnelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTunnelResponseBody) GoString() string {
	return s.String()
}

func (s *ListTunnelResponseBody) SetImportTunnelList(v *ListTunnelResp) *ListTunnelResponseBody {
	s.ImportTunnelList = v
	return s
}

type ListTunnelResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTunnelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTunnelResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTunnelResponse) GoString() string {
	return s.String()
}

func (s *ListTunnelResponse) SetHeaders(v map[string]*string) *ListTunnelResponse {
	s.Headers = v
	return s
}

func (s *ListTunnelResponse) SetStatusCode(v int32) *ListTunnelResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTunnelResponse) SetBody(v *ListTunnelResponseBody) *ListTunnelResponse {
	s.Body = v
	return s
}

type UpdateAddressRequest struct {
	// The details for updating the data address.
	ImportAddress *UpdateAddressInfo `json:"ImportAddress,omitempty" xml:"ImportAddress,omitempty"`
}

func (s UpdateAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAddressRequest) GoString() string {
	return s.String()
}

func (s *UpdateAddressRequest) SetImportAddress(v *UpdateAddressInfo) *UpdateAddressRequest {
	s.ImportAddress = v
	return s
}

type UpdateAddressResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAddressResponse) GoString() string {
	return s.String()
}

func (s *UpdateAddressResponse) SetHeaders(v map[string]*string) *UpdateAddressResponse {
	s.Headers = v
	return s
}

func (s *UpdateAddressResponse) SetStatusCode(v int32) *UpdateAddressResponse {
	s.StatusCode = &v
	return s
}

type UpdateJobRequest struct {
	// The details for updating the task.
	ImportJob *UpdateJobInfo `json:"ImportJob,omitempty" xml:"ImportJob,omitempty"`
}

func (s UpdateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobRequest) SetImportJob(v *UpdateJobInfo) *UpdateJobRequest {
	s.ImportJob = v
	return s
}

type UpdateJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateJobResponse) SetHeaders(v map[string]*string) *UpdateJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateJobResponse) SetStatusCode(v int32) *UpdateJobResponse {
	s.StatusCode = &v
	return s
}

type UpdateTunnelRequest struct {
	// The details for updating the tunnel.
	ImportTunnel *UpdateTunnelInfo `json:"ImportTunnel,omitempty" xml:"ImportTunnel,omitempty"`
}

func (s UpdateTunnelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTunnelRequest) GoString() string {
	return s.String()
}

func (s *UpdateTunnelRequest) SetImportTunnel(v *UpdateTunnelInfo) *UpdateTunnelRequest {
	s.ImportTunnel = v
	return s
}

type UpdateTunnelResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateTunnelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTunnelResponse) GoString() string {
	return s.String()
}

func (s *UpdateTunnelResponse) SetHeaders(v map[string]*string) *UpdateTunnelResponse {
	s.Headers = v
	return s
}

func (s *UpdateTunnelResponse) SetStatusCode(v int32) *UpdateTunnelResponse {
	s.StatusCode = &v
	return s
}

type VerifyAddressResponseBody struct {
	// 校验数据地址详情。
	VerifyAddressResponse *VerifyAddressResp `json:"VerifyAddressResponse,omitempty" xml:"VerifyAddressResponse,omitempty"`
}

func (s VerifyAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyAddressResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyAddressResponseBody) SetVerifyAddressResponse(v *VerifyAddressResp) *VerifyAddressResponseBody {
	s.VerifyAddressResponse = v
	return s
}

type VerifyAddressResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyAddressResponse) GoString() string {
	return s.String()
}

func (s *VerifyAddressResponse) SetHeaders(v map[string]*string) *VerifyAddressResponse {
	s.Headers = v
	return s
}

func (s *VerifyAddressResponse) SetStatusCode(v int32) *VerifyAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyAddressResponse) SetBody(v *VerifyAddressResponseBody) *VerifyAddressResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.ProductId = tea.String("hcs-mgw")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = tea.String("")
	return nil
}

// Summary:
//
// Creates a data address.
//
// Description:
//
//   To create a data address, you must have the permission on mgw:CreateImportAddress.
//
// 	- If you want to use an agent to migrate data, you must create an agent first and then associate the agent with a data address when you create the data address.
//
// @param request - CreateAddressRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAddressResponse
func (client *Client) CreateAddressWithOptions(userid *string, request *CreateAddressRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImportAddress)) {
		body["ImportAddress"] = request.ImportAddress
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAddress"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/address"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &CreateAddressResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data address.
//
// Description:
//
//   To create a data address, you must have the permission on mgw:CreateImportAddress.
//
// 	- If you want to use an agent to migrate data, you must create an agent first and then associate the agent with a data address when you create the data address.
//
// @param request - CreateAddressRequest
//
// @return CreateAddressResponse
func (client *Client) CreateAddress(userid *string, request *CreateAddressRequest) (_result *CreateAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAddressResponse{}
	_body, _err := client.CreateAddressWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The request boy for creating the agent.
//
// Description:
//
//   To create an agent, you must have the permission on mgw:CreateImportAgent.
//
// 	- If you want to migrate data to Alibaba Cloud over an Express Connect circuit or a VPN gateway, or migrate data from a self-managed storage space to Alibaba Cloud, you can deploy an agent.
//
// 	- Before you create an agent, you must create a tunnel. An agent must be associated with a tunnel.
//
// @param request - CreateAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentResponse
func (client *Client) CreateAgentWithOptions(userid *string, request *CreateAgentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImportAgent)) {
		body["ImportAgent"] = request.ImportAgent
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAgent"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/agent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &CreateAgentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The request boy for creating the agent.
//
// Description:
//
//   To create an agent, you must have the permission on mgw:CreateImportAgent.
//
// 	- If you want to migrate data to Alibaba Cloud over an Express Connect circuit or a VPN gateway, or migrate data from a self-managed storage space to Alibaba Cloud, you can deploy an agent.
//
// 	- Before you create an agent, you must create a tunnel. An agent must be associated with a tunnel.
//
// @param request - CreateAgentRequest
//
// @return CreateAgentResponse
func (client *Client) CreateAgent(userid *string, request *CreateAgentRequest) (_result *CreateAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAgentResponse{}
	_body, _err := client.CreateAgentWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a migration task.
//
// Description:
//
//   To create a migration task, you must have the permission on mgw:CreateImportJob.
//
// 	- Before you create a migration task, you must create data addresses.
//
// 	- A migration task can run multiple rounds. Each round has an execution ID.
//
// @param request - CreateJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobResponse
func (client *Client) CreateJobWithOptions(userid *string, request *CreateJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImportJob)) {
		body["ImportJob"] = request.ImportJob
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateJob"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/job"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a migration task.
//
// Description:
//
//   To create a migration task, you must have the permission on mgw:CreateImportJob.
//
// 	- Before you create a migration task, you must create data addresses.
//
// 	- A migration task can run multiple rounds. Each round has an execution ID.
//
// @param request - CreateJobRequest
//
// @return CreateJobResponse
func (client *Client) CreateJob(userid *string, request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a migration report.
//
// Description:
//
//   To create a migration report, you must have the permission on mgw:CreateImportReport.
//
// 	- If you specify that a migration report is to be generated when you create a migration task, you do not need to call this operation. If you do not specify that a migration report is to be generated when you create a migration task, you can call this operation to create a migration report for an execution with the specified ID.
//
// @param request - CreateReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReportResponse
func (client *Client) CreateReportWithOptions(userid *string, request *CreateReportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateReport)) {
		body["CreateReport"] = request.CreateReport
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateReport"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/report"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &CreateReportResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a migration report.
//
// Description:
//
//   To create a migration report, you must have the permission on mgw:CreateImportReport.
//
// 	- If you specify that a migration report is to be generated when you create a migration task, you do not need to call this operation. If you do not specify that a migration report is to be generated when you create a migration task, you can call this operation to create a migration report for an execution with the specified ID.
//
// @param request - CreateReportRequest
//
// @return CreateReportResponse
func (client *Client) CreateReport(userid *string, request *CreateReportRequest) (_result *CreateReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateReportResponse{}
	_body, _err := client.CreateReportWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a tunnel.
//
// Description:
//
//   To create a tunnel, you must have the permission on mgw:CreateImportTunnel.
//
// 	- When you use an agent to migrate data, the agent must be associated with a tunnel.
//
// 	- A tunnel can be associated with multiple agents. You can throttle the traffic of the agents that are associated with the same tunnel by setting the bandwidth and the number of requests per second for the tunnel.
//
// @param request - CreateTunnelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTunnelResponse
func (client *Client) CreateTunnelWithOptions(userid *string, request *CreateTunnelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTunnelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImportTunnel)) {
		body["ImportTunnel"] = request.ImportTunnel
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTunnel"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tunnel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &CreateTunnelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a tunnel.
//
// Description:
//
//   To create a tunnel, you must have the permission on mgw:CreateImportTunnel.
//
// 	- When you use an agent to migrate data, the agent must be associated with a tunnel.
//
// 	- A tunnel can be associated with multiple agents. You can throttle the traffic of the agents that are associated with the same tunnel by setting the bandwidth and the number of requests per second for the tunnel.
//
// @param request - CreateTunnelRequest
//
// @return CreateTunnelResponse
func (client *Client) CreateTunnel(userid *string, request *CreateTunnelRequest) (_result *CreateTunnelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTunnelResponse{}
	_body, _err := client.CreateTunnelWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data address.
//
// Description:
//
//   To delete a data address, you must have the permission on mgw:DeleteImportAddress.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAddressResponse
func (client *Client) DeleteAddressWithOptions(userid *string, addressName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAddressResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAddress"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/address/" + tea.StringValue(addressName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteAddressResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data address.
//
// Description:
//
//   To delete a data address, you must have the permission on mgw:DeleteImportAddress.
//
// @return DeleteAddressResponse
func (client *Client) DeleteAddress(userid *string, addressName *string) (_result *DeleteAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAddressResponse{}
	_body, _err := client.DeleteAddressWithOptions(userid, addressName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an agent.
//
// Description:
//
//   To delete an agent, you must have the permission on mgw:DeleteImportAgent.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentResponse
func (client *Client) DeleteAgentWithOptions(userid *string, agentName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAgentResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAgent"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/agent/" + tea.StringValue(agentName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteAgentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an agent.
//
// Description:
//
//   To delete an agent, you must have the permission on mgw:DeleteImportAgent.
//
// @return DeleteAgentResponse
func (client *Client) DeleteAgent(userid *string, agentName *string) (_result *DeleteAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAgentResponse{}
	_body, _err := client.DeleteAgentWithOptions(userid, agentName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a migration task.
//
// Description:
//
//   To delete a migration task, you must have the permission on mgw:DeleteImportJob.
//
// 	- The operation to delete a migration task is asynchronous. The migration task remains in the Deleting state until it is deleted.
//
// @param request - DeleteJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobResponse
func (client *Client) DeleteJobWithOptions(userid *string, jobName *string, request *DeleteJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForceDelete)) {
		query["forceDelete"] = request.ForceDelete
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteJob"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/job/" + tea.StringValue(jobName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a migration task.
//
// Description:
//
//   To delete a migration task, you must have the permission on mgw:DeleteImportJob.
//
// 	- The operation to delete a migration task is asynchronous. The migration task remains in the Deleting state until it is deleted.
//
// @param request - DeleteJobRequest
//
// @return DeleteJobResponse
func (client *Client) DeleteJob(userid *string, jobName *string, request *DeleteJobRequest) (_result *DeleteJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteJobResponse{}
	_body, _err := client.DeleteJobWithOptions(userid, jobName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a tunnel.
//
// Description:
//
//   To delete a tunnel, you must have the permission on mgw:DeleteImportTunnel.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTunnelResponse
func (client *Client) DeleteTunnelWithOptions(userid *string, tunnelId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTunnelResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTunnel"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tunnel/" + tea.StringValue(tunnelId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteTunnelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a tunnel.
//
// Description:
//
//   To delete a tunnel, you must have the permission on mgw:DeleteImportTunnel.
//
// @return DeleteTunnelResponse
func (client *Client) DeleteTunnel(userid *string, tunnelId *string) (_result *DeleteTunnelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTunnelResponse{}
	_body, _err := client.DeleteTunnelWithOptions(userid, tunnelId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of a data address.
//
// Description:
//
//   To query the information about a data address, you must have the permission on mgw:GetImportAddress.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAddressResponse
func (client *Client) GetAddressWithOptions(userid *string, addressName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAddressResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAddress"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/address/" + tea.StringValue(addressName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetAddressResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a data address.
//
// Description:
//
//   To query the information about a data address, you must have the permission on mgw:GetImportAddress.
//
// @return GetAddressResponse
func (client *Client) GetAddress(userid *string, addressName *string) (_result *GetAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAddressResponse{}
	_body, _err := client.GetAddressWithOptions(userid, addressName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of an agent.
//
// Description:
//
//   To query the information about an agent, you must have the permission on mgw:GetImportAgent.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentResponse
func (client *Client) GetAgentWithOptions(userid *string, agentName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAgentResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAgent"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/agent/" + tea.StringValue(agentName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetAgentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of an agent.
//
// Description:
//
//   To query the information about an agent, you must have the permission on mgw:GetImportAgent.
//
// @return GetAgentResponse
func (client *Client) GetAgent(userid *string, agentName *string) (_result *GetAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentResponse{}
	_body, _err := client.GetAgentWithOptions(userid, agentName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the running status of an agent.
//
// Description:
//
//   To query the status of an agent, you must have the permission on mgw:GetImportAgent.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentStatusResponse
func (client *Client) GetAgentStatusWithOptions(userid *string, agentName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAgentStatusResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAgentStatus"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/agent/" + tea.StringValue(agentName) + "?status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetAgentStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the running status of an agent.
//
// Description:
//
//   To query the status of an agent, you must have the permission on mgw:GetImportAgent.
//
// @return GetAgentStatusResponse
func (client *Client) GetAgentStatus(userid *string, agentName *string) (_result *GetAgentStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentStatusResponse{}
	_body, _err := client.GetAgentStatusWithOptions(userid, agentName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of a migration task.
//
// Description:
//
//   To query the information about a migration task, you must have the permission on mgw:GetImportJob.
//
// @param request - GetJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
func (client *Client) GetJobWithOptions(userid *string, jobName *string, request *GetJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ByVersion)) {
		query["byVersion"] = request.ByVersion
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJob"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/job/" + tea.StringValue(jobName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a migration task.
//
// Description:
//
//   To query the information about a migration task, you must have the permission on mgw:GetImportJob.
//
// @param request - GetJobRequest
//
// @return GetJobResponse
func (client *Client) GetJob(userid *string, jobName *string, request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(userid, jobName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the list of files that fail to be migrated when files fail to be migrated during a migration task.
//
// Description:
//
//   To query the retry information about a migration task, you must have the permission on mgw:GetImportJobResult.
//
// 	- If files fail to be migrated during a migration task, a list of files that fail to be migrated is generated. You can call this operation to query this list. You can create a data address based on this list and create a subtask. This way, you can migrate these files again.
//
// @param request - GetJobResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResultResponse
func (client *Client) GetJobResultWithOptions(userid *string, jobName *string, request *GetJobResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RuntimeId)) {
		query["runtimeId"] = request.RuntimeId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobResult"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/job/" + tea.StringValue(jobName) + "?jobResult"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetJobResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the list of files that fail to be migrated when files fail to be migrated during a migration task.
//
// Description:
//
//   To query the retry information about a migration task, you must have the permission on mgw:GetImportJobResult.
//
// 	- If files fail to be migrated during a migration task, a list of files that fail to be migrated is generated. You can call this operation to query this list. You can create a data address based on this list and create a subtask. This way, you can migrate these files again.
//
// @param request - GetJobResultRequest
//
// @return GetJobResultResponse
func (client *Client) GetJobResult(userid *string, jobName *string, request *GetJobResultRequest) (_result *GetJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobResultResponse{}
	_body, _err := client.GetJobResultWithOptions(userid, jobName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of a migration report.
//
// Description:
//
//   To query the information about a migration report, you must have the permission on mgw:GetImportReport.
//
// 	- The migration report is pushed to the destination data address. For more information, see the "View a migration report" section of the "Subsequent operations" topic in migration tutorials.
//
// @param request - GetReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReportResponse
func (client *Client) GetReportWithOptions(userid *string, request *GetReportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RuntimeId)) {
		query["runtimeId"] = request.RuntimeId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetReport"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/report"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetReportResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a migration report.
//
// Description:
//
//   To query the information about a migration report, you must have the permission on mgw:GetImportReport.
//
// 	- The migration report is pushed to the destination data address. For more information, see the "View a migration report" section of the "Subsequent operations" topic in migration tutorials.
//
// @param request - GetReportRequest
//
// @return GetReportResponse
func (client *Client) GetReport(userid *string, request *GetReportRequest) (_result *GetReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetReportResponse{}
	_body, _err := client.GetReportWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of a tunnel.
//
// Description:
//
//   To query the information about a tunnel, you must have the permission on mgw:GetImportTunnel.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTunnelResponse
func (client *Client) GetTunnelWithOptions(userid *string, tunnelId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTunnelResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTunnel"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tunnel/" + tea.StringValue(tunnelId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetTunnelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a tunnel.
//
// Description:
//
//   To query the information about a tunnel, you must have the permission on mgw:GetImportTunnel.
//
// @return GetTunnelResponse
func (client *Client) GetTunnel(userid *string, tunnelId *string) (_result *GetTunnelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTunnelResponse{}
	_body, _err := client.GetTunnelWithOptions(userid, tunnelId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the data addresses created by a user in the specific region.
//
// Description:
//
//   To query a list of data addresses, you must have the permission on mgw:ListImportAddress.
//
// @param request - ListAddressRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddressResponse
func (client *Client) ListAddressWithOptions(userid *string, request *ListAddressRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAddress"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/addresslist"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &ListAddressResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the data addresses created by a user in the specific region.
//
// Description:
//
//   To query a list of data addresses, you must have the permission on mgw:ListImportAddress.
//
// @param request - ListAddressRequest
//
// @return ListAddressResponse
func (client *Client) ListAddress(userid *string, request *ListAddressRequest) (_result *ListAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAddressResponse{}
	_body, _err := client.ListAddressWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the agents created by a user in the specific region.
//
// Description:
//
//   To query a list of agents, you must have the permission on mgw:ListImportAgent.
//
// @param request - ListAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentResponse
func (client *Client) ListAgentWithOptions(userid *string, request *ListAgentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAgent"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/agentlist"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &ListAgentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the agents created by a user in the specific region.
//
// Description:
//
//   To query a list of agents, you must have the permission on mgw:ListImportAgent.
//
// @param request - ListAgentRequest
//
// @return ListAgentResponse
func (client *Client) ListAgent(userid *string, request *ListAgentRequest) (_result *ListAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentResponse{}
	_body, _err := client.ListAgentWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the migration tasks created by a user in the specific region.
//
// Description:
//
//   To query a list of migration tasks, you must have the permission on mgw:ListImportJob.
//
// @param request - ListJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobResponse
func (client *Client) ListJobWithOptions(userid *string, request *ListJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["all"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.ParentName)) {
		query["parentName"] = request.ParentName
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJob"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/joblist"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &ListJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the migration tasks created by a user in the specific region.
//
// Description:
//
//   To query a list of migration tasks, you must have the permission on mgw:ListImportJob.
//
// @param request - ListJobRequest
//
// @return ListJobResponse
func (client *Client) ListJob(userid *string, request *ListJobRequest) (_result *ListJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobResponse{}
	_body, _err := client.ListJobWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the running history of a migration task.
//
// Description:
//
//   To query the execution history of a migration task, you must have the permission on mgw:ListImportJobHistory.
//
// 	- A migration task can run multiple rounds. A unique execution ID is generated for each round.
//
// 	- The execution history of a migration task records the change history of the task status.
//
// @param request - ListJobHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobHistoryResponse
func (client *Client) ListJobHistoryWithOptions(userid *string, jobName *string, request *ListJobHistoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListJobHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.RuntimeId)) {
		query["runtimeId"] = request.RuntimeId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobHistory"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/jobhistory/" + tea.StringValue(jobName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &ListJobHistoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the running history of a migration task.
//
// Description:
//
//   To query the execution history of a migration task, you must have the permission on mgw:ListImportJobHistory.
//
// 	- A migration task can run multiple rounds. A unique execution ID is generated for each round.
//
// 	- The execution history of a migration task records the change history of the task status.
//
// @param request - ListJobHistoryRequest
//
// @return ListJobHistoryResponse
func (client *Client) ListJobHistory(userid *string, jobName *string, request *ListJobHistoryRequest) (_result *ListJobHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobHistoryResponse{}
	_body, _err := client.ListJobHistoryWithOptions(userid, jobName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists tunnels.
//
// Description:
//
//   To query a list of tunnels, you must have the permission on mgw:ListImportTunnel.
//
// @param request - ListTunnelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTunnelResponse
func (client *Client) ListTunnelWithOptions(userid *string, request *ListTunnelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTunnelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTunnel"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tunnellist"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &ListTunnelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists tunnels.
//
// Description:
//
//   To query a list of tunnels, you must have the permission on mgw:ListImportTunnel.
//
// @param request - ListTunnelRequest
//
// @return ListTunnelResponse
func (client *Client) ListTunnel(userid *string, request *ListTunnelRequest) (_result *ListTunnelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTunnelResponse{}
	_body, _err := client.ListTunnelWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a data address.
//
// Description:
//
//   To update a data address, you must have the permission on mgw:UpdateImportAddress.
//
// 	- If the data address is associated with an agent, you can scale up or down the agent.
//
// @param request - UpdateAddressRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAddressResponse
func (client *Client) UpdateAddressWithOptions(userid *string, addressName *string, request *UpdateAddressRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImportAddress)) {
		body["ImportAddress"] = request.ImportAddress
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAddress"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/address/" + tea.StringValue(addressName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &UpdateAddressResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data address.
//
// Description:
//
//   To update a data address, you must have the permission on mgw:UpdateImportAddress.
//
// 	- If the data address is associated with an agent, you can scale up or down the agent.
//
// @param request - UpdateAddressRequest
//
// @return UpdateAddressResponse
func (client *Client) UpdateAddress(userid *string, addressName *string, request *UpdateAddressRequest) (_result *UpdateAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAddressResponse{}
	_body, _err := client.UpdateAddressWithOptions(userid, addressName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the status or throttling of a task.
//
// Description:
//
//   To update a migration task, you must have the permission on mgw:UpdateImportJob.
//
// 	- You can update only the status or throttling settings of a task in a single request.
//
// @param request - UpdateJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateJobResponse
func (client *Client) UpdateJobWithOptions(userid *string, jobName *string, request *UpdateJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImportJob)) {
		body["ImportJob"] = request.ImportJob
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateJob"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/job/" + tea.StringValue(jobName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &UpdateJobResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status or throttling of a task.
//
// Description:
//
//   To update a migration task, you must have the permission on mgw:UpdateImportJob.
//
// 	- You can update only the status or throttling settings of a task in a single request.
//
// @param request - UpdateJobRequest
//
// @return UpdateJobResponse
func (client *Client) UpdateJob(userid *string, jobName *string, request *UpdateJobRequest) (_result *UpdateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateJobResponse{}
	_body, _err := client.UpdateJobWithOptions(userid, jobName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a tunnel.
//
// Description:
//
//   To update a tunnel, you must have the permission on mgw:UpdateImportTunnel.
//
// @param request - UpdateTunnelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTunnelResponse
func (client *Client) UpdateTunnelWithOptions(userid *string, tunnelId *string, request *UpdateTunnelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTunnelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImportTunnel)) {
		body["ImportTunnel"] = request.ImportTunnel
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTunnel"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tunnel/" + tea.StringValue(tunnelId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &UpdateTunnelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a tunnel.
//
// Description:
//
//   To update a tunnel, you must have the permission on mgw:UpdateImportTunnel.
//
// @param request - UpdateTunnelRequest
//
// @return UpdateTunnelResponse
func (client *Client) UpdateTunnel(userid *string, tunnelId *string, request *UpdateTunnelRequest) (_result *UpdateTunnelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTunnelResponse{}
	_body, _err := client.UpdateTunnelWithOptions(userid, tunnelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies whether a data address is available.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyAddressResponse
func (client *Client) VerifyAddressWithOptions(userid *string, addressName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VerifyAddressResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyAddress"),
		Version:     tea.String("2024-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/address/" + tea.StringValue(addressName) + "?verify"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &VerifyAddressResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies whether a data address is available.
//
// @return VerifyAddressResponse
func (client *Client) VerifyAddress(userid *string, addressName *string) (_result *VerifyAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VerifyAddressResponse{}
	_body, _err := client.VerifyAddressWithOptions(userid, addressName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
