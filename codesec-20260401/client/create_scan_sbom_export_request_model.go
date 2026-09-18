// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScanSbomExportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *CreateScanSbomExportRequest
	GetFileName() *string
	SetFormat(v string) *CreateScanSbomExportRequest
	GetFormat() *string
}

type CreateScanSbomExportRequest struct {
	// 选填。指定下载落盘的文件名（含扩展名），会签进下载地址的 Content-Disposition。
	//
	// 留空时后端按 `项目名-sbom-<format>.<扩展名>` 生成默认值。
	//
	// 不得含控制字符或路径分隔符（`/`、`\`）、长度不超过 255，否则 → 400 **`InvalidFileName`**。
	//
	// example:
	//
	// test-sbom-cyclonedx.cdx.json
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 闭合枚举，须是该次扫描 **`artifacts`*	- 里 **`status=success`*	- 的 **`artifact_kind`**。
	//
	// 未知值 → 400 **`InvalidArtifactFormat`**，且绝不参与 OSS key 构造。
	//
	// This parameter is required.
	//
	// example:
	//
	// cyclonedx-json
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
}

func (s CreateScanSbomExportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScanSbomExportRequest) GoString() string {
	return s.String()
}

func (s *CreateScanSbomExportRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateScanSbomExportRequest) GetFormat() *string {
	return s.Format
}

func (s *CreateScanSbomExportRequest) SetFileName(v string) *CreateScanSbomExportRequest {
	s.FileName = &v
	return s
}

func (s *CreateScanSbomExportRequest) SetFormat(v string) *CreateScanSbomExportRequest {
	s.Format = &v
	return s
}

func (s *CreateScanSbomExportRequest) Validate() error {
	return dara.Validate(s)
}
