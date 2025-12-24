// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateFileImportTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileDownloadUrl(v string) *GenerateFileImportTemplateResponseBody
	GetFileDownloadUrl() *string
	SetRequestId(v string) *GenerateFileImportTemplateResponseBody
	GetRequestId() *string
}

type GenerateFileImportTemplateResponseBody struct {
	// example:
	//
	// https://test.oss-cn-hangzhou.aliyuncs.com/idaas_ly77wa2oexrciw5v672vxxxx/tmp/eiam_v2_user_import_1766469463365.csv
	FileDownloadUrl *string `json:"FileDownloadUrl,omitempty" xml:"FileDownloadUrl,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateFileImportTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateFileImportTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateFileImportTemplateResponseBody) GetFileDownloadUrl() *string {
	return s.FileDownloadUrl
}

func (s *GenerateFileImportTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateFileImportTemplateResponseBody) SetFileDownloadUrl(v string) *GenerateFileImportTemplateResponseBody {
	s.FileDownloadUrl = &v
	return s
}

func (s *GenerateFileImportTemplateResponseBody) SetRequestId(v string) *GenerateFileImportTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateFileImportTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
