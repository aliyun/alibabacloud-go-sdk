// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModuleVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetModuleVersionResponseBody
	GetRequestId() *string
	SetVersion(v *GetModuleVersionResponseBodyVersion) *GetModuleVersionResponseBody
	GetVersion() *GetModuleVersionResponseBodyVersion
}

type GetModuleVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0D298375-F92F-5B65-82E4-EA68F02521F1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The version details.
	Version *GetModuleVersionResponseBodyVersion `json:"version,omitempty" xml:"version,omitempty" type:"Struct"`
}

func (s GetModuleVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModuleVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetModuleVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModuleVersionResponseBody) GetVersion() *GetModuleVersionResponseBodyVersion {
	return s.Version
}

func (s *GetModuleVersionResponseBody) SetRequestId(v string) *GetModuleVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModuleVersionResponseBody) SetVersion(v *GetModuleVersionResponseBodyVersion) *GetModuleVersionResponseBody {
	s.Version = v
	return s
}

func (s *GetModuleVersionResponseBody) Validate() error {
	if s.Version != nil {
		if err := s.Version.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModuleVersionResponseBodyVersion struct {
	// The time when the version was created.
	//
	// example:
	//
	// 2022-09-08T18:07:40Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The version description.
	//
	// example:
	//
	// this is description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The template ID.
	//
	// example:
	//
	// mod-4267dcfbf1b6dfffbc27e218d1b66
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// The template version number.
	//
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// The version name.
	//
	// example:
	//
	// versionName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version source. Valid values:
	//
	// - OSS: imported from OSS.
	//
	// - Registry: created by using a template from the template center.
	//
	// - ExportTask: exported from a resource export task.
	//
	// - Upload: uploaded as a file.
	//
	// - Shared: cloned from a shared source.
	//
	// - Editor: edited online.
	//
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The path of the version source.
	//
	// - If the source is Registry, the value is in the format of <workspace name>/<module name>:<module version>. Example: terraform-alicloud-modules/rds:1.0.0.
	//
	// - If the source is OSS, the value is in the format of oss::<file link>. Example: oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip.
	//
	// - If the source is ExportTask, the value is in the format of <export task ID>:<exported version>. Example: ex-3b6cb9fa4751afff298da723c24ac:v1.
	//
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// The path of the State file that corresponds to the template. Currently, only OSS paths are supported. The value is in the format of oss::<OSS file path>/terraform.tfstate.
	//
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/terraform.tfstate
	StatePath *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
	// The Terraform content.
	TerraformContext map[string]interface{} `json:"terraformContext,omitempty" xml:"terraformContext,omitempty"`
	// The version generation strategy. Valid values:
	//
	// - Manual: manually generate a version. This is the default value.
	//
	// - SourcePathUpdated: a new version is generated when the sourcePath is modified.
	//
	// example:
	//
	// Manual
	VersionStrategy *string `json:"versionStrategy,omitempty" xml:"versionStrategy,omitempty"`
}

func (s GetModuleVersionResponseBodyVersion) String() string {
	return dara.Prettify(s)
}

func (s GetModuleVersionResponseBodyVersion) GoString() string {
	return s.String()
}

func (s *GetModuleVersionResponseBodyVersion) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetModuleVersionResponseBodyVersion) GetDescription() *string {
	return s.Description
}

func (s *GetModuleVersionResponseBodyVersion) GetModuleId() *string {
	return s.ModuleId
}

func (s *GetModuleVersionResponseBodyVersion) GetModuleVersion() *string {
	return s.ModuleVersion
}

func (s *GetModuleVersionResponseBodyVersion) GetName() *string {
	return s.Name
}

func (s *GetModuleVersionResponseBodyVersion) GetSource() *string {
	return s.Source
}

func (s *GetModuleVersionResponseBodyVersion) GetSourcePath() *string {
	return s.SourcePath
}

func (s *GetModuleVersionResponseBodyVersion) GetStatePath() *string {
	return s.StatePath
}

func (s *GetModuleVersionResponseBodyVersion) GetTerraformContext() map[string]interface{} {
	return s.TerraformContext
}

func (s *GetModuleVersionResponseBodyVersion) GetVersionStrategy() *string {
	return s.VersionStrategy
}

func (s *GetModuleVersionResponseBodyVersion) SetCreateTime(v string) *GetModuleVersionResponseBodyVersion {
	s.CreateTime = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetDescription(v string) *GetModuleVersionResponseBodyVersion {
	s.Description = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetModuleId(v string) *GetModuleVersionResponseBodyVersion {
	s.ModuleId = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetModuleVersion(v string) *GetModuleVersionResponseBodyVersion {
	s.ModuleVersion = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetName(v string) *GetModuleVersionResponseBodyVersion {
	s.Name = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetSource(v string) *GetModuleVersionResponseBodyVersion {
	s.Source = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetSourcePath(v string) *GetModuleVersionResponseBodyVersion {
	s.SourcePath = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetStatePath(v string) *GetModuleVersionResponseBodyVersion {
	s.StatePath = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetTerraformContext(v map[string]interface{}) *GetModuleVersionResponseBodyVersion {
	s.TerraformContext = v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetVersionStrategy(v string) *GetModuleVersionResponseBodyVersion {
	s.VersionStrategy = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) Validate() error {
	return dara.Validate(s)
}
