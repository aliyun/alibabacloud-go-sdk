// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyFileUploadLeaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryType(v string) *ApplyFileUploadLeaseRequest
	GetCategoryType() *string
	SetFileName(v string) *ApplyFileUploadLeaseRequest
	GetFileName() *string
	SetMd5(v string) *ApplyFileUploadLeaseRequest
	GetMd5() *string
	SetSizeInBytes(v string) *ApplyFileUploadLeaseRequest
	GetSizeInBytes() *string
	SetUseInternalEndpoint(v bool) *ApplyFileUploadLeaseRequest
	GetUseInternalEndpoint() *bool
}

type ApplyFileUploadLeaseRequest struct {
	// The category type. If this parameter is not specified, the default value is UNSTRUCTURED. Valid values:
	//
	// UNSTRUCTURED: category for building [knowledge base](https://help.aliyun.com/document_detail/2807740.html) scenarios.
	//
	// <props="china">
	//
	// SESSION_FILE: upload files for agent application [conversational interactions](https://help.aliyun.com/zh/model-studio/user-guide/file-interaction).
	//
	//
	// > To create a new data table and upload data, use the Alibaba Cloud Model Studio console. This is not supported through the API.
	//
	// >
	//
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// When uploading files for building a [knowledge base](https://help.aliyun.com/document_detail/2807740.html):
	//
	// - This field represents the name of the uploaded file. The file name must include the file format extension. Supported formats:
	//
	//      - Documents (less than 150 MB): doc, docx, wps, ppt, pptx, xls, xlsx, md, txt, pdf, epub, mobi.
	//
	//      - Spreadsheets (recommended within 10 MB and 100,000 rows): xls, xlsx.
	//
	//      - Plain text (recommended not to exceed 10 MB): md, txt.
	//
	//      - Images (less than 20 MB, shortest side > 15 px, longest side < 8192 px, aspect ratio < 50): png, jpg, jpeg, bmp, gif.
	//
	//      - Audio: aac, amr, flac, flv, m4a, mp3, mpeg, ogg, opus, wav, webm, wma.
	//
	//      - Video: mp4, mkv, avi, mov, wmv.
	//
	// - The file name must be 4 to 128 characters in length. For other limits, see [Knowledge base quotas and limits](https://help.aliyun.com/document_detail/2880605.html).
	//
	// > To create a new data table and upload data, use the Alibaba Cloud Model Studio console. This is not supported through the API.
	//
	// >
	//
	// <props="china">
	//
	// When uploading files for agent application [conversational interactions](https://help.aliyun.com/zh/model-studio/user-guide/file-interaction):
	//
	// - This field represents the name of the uploaded file. The file name must include the file format extension. Supported formats:
	//
	//      - Documents: doc, docx, wps, ppt, pptx, xls, xlsx, md, txt, pdf, epub, mobi.
	//
	//      - Images: png, jpg, jpeg, bmp, gif.
	//
	//      - Audio: aac, amr, flac, flv, m4a, mp3, mpeg, ogg, opus, wav, webm, wma.
	//
	//      - Video: mp4, mkv, avi, mov, wmv.
	//
	// - The file name must be 4 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// XXXX产品清单.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The MD5 value of the uploaded file. The server will verify this field (currently not enabled). Please fill in the correct value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 19657c391f6c70bcea63c154d8606bb3
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The size of the uploaded file in bytes. The server will verify this field (currently not enabled). Please fill in the correct value. Valid values: 1 B to 100 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	SizeInBytes *string `json:"SizeInBytes,omitempty" xml:"SizeInBytes,omitempty"`
	// <props="china">
	//
	// If you have enabled [Alibaba Cloud Model Studio secure storage](https://help.aliyun.com/zh/model-studio/configure-resources-in-private-network) and need to generate a lease URL that is only accessible from the Alibaba Cloud internal network in the same region, you can set this parameter to true to improve security. If this parameter is not specified, the default value is false, which generates a publicly accessible lease URL.
	//
	// > If you have not enabled Alibaba Cloud Model Studio secure storage, or are unsure whether you are using it, do not set this parameter to true (upload will fail).
	//
	//
	//
	// <props="intl">
	//
	// If you have enabled Alibaba Cloud Model Studio secure storage and need to generate a lease URL that is only accessible from the Alibaba Cloud internal network in the same region, you can set this parameter to true to improve security. If this parameter is not specified, the default value is false, which generates a publicly accessible lease URL.
	//
	// > If you have not enabled Alibaba Cloud Model Studio secure storage, or are unsure whether you are using it, do not set this parameter to true (upload will fail).
	//
	// example:
	//
	// false
	UseInternalEndpoint *bool `json:"UseInternalEndpoint,omitempty" xml:"UseInternalEndpoint,omitempty"`
}

func (s ApplyFileUploadLeaseRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyFileUploadLeaseRequest) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseRequest) GetCategoryType() *string {
	return s.CategoryType
}

func (s *ApplyFileUploadLeaseRequest) GetFileName() *string {
	return s.FileName
}

func (s *ApplyFileUploadLeaseRequest) GetMd5() *string {
	return s.Md5
}

func (s *ApplyFileUploadLeaseRequest) GetSizeInBytes() *string {
	return s.SizeInBytes
}

func (s *ApplyFileUploadLeaseRequest) GetUseInternalEndpoint() *bool {
	return s.UseInternalEndpoint
}

func (s *ApplyFileUploadLeaseRequest) SetCategoryType(v string) *ApplyFileUploadLeaseRequest {
	s.CategoryType = &v
	return s
}

func (s *ApplyFileUploadLeaseRequest) SetFileName(v string) *ApplyFileUploadLeaseRequest {
	s.FileName = &v
	return s
}

func (s *ApplyFileUploadLeaseRequest) SetMd5(v string) *ApplyFileUploadLeaseRequest {
	s.Md5 = &v
	return s
}

func (s *ApplyFileUploadLeaseRequest) SetSizeInBytes(v string) *ApplyFileUploadLeaseRequest {
	s.SizeInBytes = &v
	return s
}

func (s *ApplyFileUploadLeaseRequest) SetUseInternalEndpoint(v bool) *ApplyFileUploadLeaseRequest {
	s.UseInternalEndpoint = &v
	return s
}

func (s *ApplyFileUploadLeaseRequest) Validate() error {
	return dara.Validate(s)
}
