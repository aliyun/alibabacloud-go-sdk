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
	// example:
	//
	// 0D298375-F92F-5B65-82E4-EA68F02521F1
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Version   *GetModuleVersionResponseBodyVersion `json:"version,omitempty" xml:"version,omitempty" type:"Struct"`
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
	// example:
	//
	// 2022-09-08T18:07:40Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// mod-4267dcfbf1b6dfffbc27e218d1b66
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/terraform.tfstate
	StatePath        *string                `json:"statePath,omitempty" xml:"statePath,omitempty"`
	TerraformContext map[string]interface{} `json:"terraformContext,omitempty" xml:"terraformContext,omitempty"`
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
