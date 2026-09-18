// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScanSbomExportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadUrl(v string) *CreateScanSbomExportResponseBody
	GetDownloadUrl() *string
	SetFileName(v string) *CreateScanSbomExportResponseBody
	GetFileName() *string
	SetFormat(v string) *CreateScanSbomExportResponseBody
	GetFormat() *string
	SetMediaType(v string) *CreateScanSbomExportResponseBody
	GetMediaType() *string
	SetRequestId(v string) *CreateScanSbomExportResponseBody
	GetRequestId() *string
	SetSha256(v string) *CreateScanSbomExportResponseBody
	GetSha256() *string
	SetSizeBytes(v int64) *CreateScanSbomExportResponseBody
	GetSizeBytes() *int64
}

type CreateScanSbomExportResponseBody struct {
	// 短时签名的 GET URL（有效期见 **`url_expires_at`**），对象上已带 Content-Disposition
	//
	// example:
	//
	// https://krypton-codesafe.oss-cn-hangzhou.aliyuncs.com/1609837153086803%2F1000108%2F1000893%2F1001080%2Fartifacts%2Fcyclonedx-json%2F1f8dc54097780e9c32941289d2aff5161df694b5bfb2d0aa3fe87fe72a751363.cdx.json?Expires=1789713461&OSSAccessKeyId=STS.NYDdfgGGhFqA4XBNu2EQebMeV&Signature=pnZPAML9CPKOIDyC4b1DK6gQmhs%3D&response-content-disposition=attachment%3B%20filename%3D%22test-sbom-cyclonedx.cdx.json%22%3B%20filename%2A%3DUTF-8%27%27test-sbom-cyclonedx.cdx.json&security-token=CAIS*
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// 最终生效的文件名（客户端指定或后端默认），已签进 **`download_url`*	- 的 **`Content-Disposition`**， 浏览器直连下载即按此落盘。
	//
	// example:
	//
	// test-sbom-cyclonedx.cdx.json
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// cyclonedx-json
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// 交付文件 MIME 类型
	//
	// example:
	//
	// application/vnd.cyclonedx+json
	MediaType *string `json:"mediaType,omitempty" xml:"mediaType,omitempty"`
	// example:
	//
	// 9A1F403F-0A85-5578-8B7C-55E3E9408659
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 交付文件本身（注入后重算）的摘要，供下载完整性校验；与源制品的 sha256 不同
	//
	// example:
	//
	// 1f8dc54097780e9c32941289d2aff5161df694b5bfb2d0aa3fe87fe72a751363
	Sha256 *string `json:"sha256,omitempty" xml:"sha256,omitempty"`
	// 交付文件字节数
	//
	// example:
	//
	// 791355
	SizeBytes *int64 `json:"sizeBytes,omitempty" xml:"sizeBytes,omitempty"`
}

func (s CreateScanSbomExportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScanSbomExportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScanSbomExportResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *CreateScanSbomExportResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *CreateScanSbomExportResponseBody) GetFormat() *string {
	return s.Format
}

func (s *CreateScanSbomExportResponseBody) GetMediaType() *string {
	return s.MediaType
}

func (s *CreateScanSbomExportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScanSbomExportResponseBody) GetSha256() *string {
	return s.Sha256
}

func (s *CreateScanSbomExportResponseBody) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *CreateScanSbomExportResponseBody) SetDownloadUrl(v string) *CreateScanSbomExportResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *CreateScanSbomExportResponseBody) SetFileName(v string) *CreateScanSbomExportResponseBody {
	s.FileName = &v
	return s
}

func (s *CreateScanSbomExportResponseBody) SetFormat(v string) *CreateScanSbomExportResponseBody {
	s.Format = &v
	return s
}

func (s *CreateScanSbomExportResponseBody) SetMediaType(v string) *CreateScanSbomExportResponseBody {
	s.MediaType = &v
	return s
}

func (s *CreateScanSbomExportResponseBody) SetRequestId(v string) *CreateScanSbomExportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScanSbomExportResponseBody) SetSha256(v string) *CreateScanSbomExportResponseBody {
	s.Sha256 = &v
	return s
}

func (s *CreateScanSbomExportResponseBody) SetSizeBytes(v int64) *CreateScanSbomExportResponseBody {
	s.SizeBytes = &v
	return s
}

func (s *CreateScanSbomExportResponseBody) Validate() error {
	return dara.Validate(s)
}
