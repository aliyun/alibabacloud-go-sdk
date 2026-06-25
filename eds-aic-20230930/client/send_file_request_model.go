// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIdList(v []*string) *SendFileRequest
	GetAndroidInstanceIdList() []*string
	SetAutoInstall(v bool) *SendFileRequest
	GetAutoInstall() *bool
	SetClientToken(v string) *SendFileRequest
	GetClientToken() *string
	SetFileMd5(v string) *SendFileRequest
	GetFileMd5() *string
	SetSourceFilePath(v string) *SendFileRequest
	GetSourceFilePath() *string
	SetTargetFileName(v string) *SendFileRequest
	GetTargetFileName() *string
	SetUploadEndpoint(v string) *SendFileRequest
	GetUploadEndpoint() *string
	SetUploadType(v string) *SendFileRequest
	GetUploadType() *string
	SetUploadUrl(v string) *SendFileRequest
	GetUploadUrl() *string
}

type SendFileRequest struct {
	// The IDs of one or more cloud phone instances.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// Specifies whether to automatically install the application after the file is uploaded.
	//
	// example:
	//
	// true
	AutoInstall *bool `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	// A client-generated token that ensures request idempotence and prevents duplicate submissions. The token can contain up to 100 characters.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FileMd5     *string `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	// The destination path on the cloud phone.
	//
	// > If `UploadType` is `OSS` or `OSS_BRIDGED`, `SourceFilePath` must specify a directory, for example, `/sdcard/Download/`. If `UploadType` is `DOWNLOAD_URL`, `SourceFilePath` must specify a full file path, for example, `/sdcard/Download/MyFile.txt`.
	//
	// This parameter is required.
	//
	// example:
	//
	// /sdcard/Download
	SourceFilePath *string `json:"SourceFilePath,omitempty" xml:"SourceFilePath,omitempty"`
	// The name for the destination file on the cloud phone.
	//
	// > This parameter is optional and takes effect only when `UploadType` is set to `OSS` or `OSS_BRIDGED`. If you specify this parameter, the file is saved with this name in the path specified by `SourceFilePath`. If you leave this parameter empty, the source file name is used. This parameter is ignored when `UploadType` is set to `DOWNLOAD_URL`.
	//
	// example:
	//
	// test.txt
	TargetFileName *string `json:"TargetFileName,omitempty" xml:"TargetFileName,omitempty"`
	// The service endpoint of Object Storage Service (OSS). This parameter is required if `UploadType` is `OSS` or `OSS_BRIDGED`.
	//
	// > If the cloud phone instance and the OSS bucket are in the same region, you can specify an internal endpoint to accelerate data transfer and avoid public data transfer costs. For example, the internal endpoint for the China (Hangzhou) region is `oss-cn-hangzhou-internal.aliyuncs.com`. For a complete list of endpoints, see [OSS regions and endpoints](https://help.aliyun.com/document_detail/31837.html).
	//
	// example:
	//
	// oss-cn-hangzhou-internal.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// The storage type of the source file. Valid values:
	//
	// - **OSS**: The file is stored in Object Storage Service (OSS).
	//
	// - **DOWNLOAD_URL**: The file is accessible via a public download link.
	//
	// - **OSS_BRIDGED**: The service first downloads the file from a public download link to an internal OSS bucket, and then distributes it to the cloud phone instances over the internal network.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
	// - If `UploadType` is `OSS`, this parameter specifies the URI of the source object in Object Storage Service (OSS).
	//
	// > The URI must be in the `oss://<bucket-name>/<object-key>` format. The specified OSS bucket name must have the `cloudphone-saved-bucket-` prefix, for example, `cloudphone-saved-bucket-example`.
	//
	// - If `UploadType` is `DOWNLOAD_URL` or `OSS_BRIDGED`, this parameter specifies the public download link of the source file.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://cloudphone-saved-bucket-example/send/a.txt
	UploadUrl *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
}

func (s SendFileRequest) String() string {
	return dara.Prettify(s)
}

func (s SendFileRequest) GoString() string {
	return s.String()
}

func (s *SendFileRequest) GetAndroidInstanceIdList() []*string {
	return s.AndroidInstanceIdList
}

func (s *SendFileRequest) GetAutoInstall() *bool {
	return s.AutoInstall
}

func (s *SendFileRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SendFileRequest) GetFileMd5() *string {
	return s.FileMd5
}

func (s *SendFileRequest) GetSourceFilePath() *string {
	return s.SourceFilePath
}

func (s *SendFileRequest) GetTargetFileName() *string {
	return s.TargetFileName
}

func (s *SendFileRequest) GetUploadEndpoint() *string {
	return s.UploadEndpoint
}

func (s *SendFileRequest) GetUploadType() *string {
	return s.UploadType
}

func (s *SendFileRequest) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *SendFileRequest) SetAndroidInstanceIdList(v []*string) *SendFileRequest {
	s.AndroidInstanceIdList = v
	return s
}

func (s *SendFileRequest) SetAutoInstall(v bool) *SendFileRequest {
	s.AutoInstall = &v
	return s
}

func (s *SendFileRequest) SetClientToken(v string) *SendFileRequest {
	s.ClientToken = &v
	return s
}

func (s *SendFileRequest) SetFileMd5(v string) *SendFileRequest {
	s.FileMd5 = &v
	return s
}

func (s *SendFileRequest) SetSourceFilePath(v string) *SendFileRequest {
	s.SourceFilePath = &v
	return s
}

func (s *SendFileRequest) SetTargetFileName(v string) *SendFileRequest {
	s.TargetFileName = &v
	return s
}

func (s *SendFileRequest) SetUploadEndpoint(v string) *SendFileRequest {
	s.UploadEndpoint = &v
	return s
}

func (s *SendFileRequest) SetUploadType(v string) *SendFileRequest {
	s.UploadType = &v
	return s
}

func (s *SendFileRequest) SetUploadUrl(v string) *SendFileRequest {
	s.UploadUrl = &v
	return s
}

func (s *SendFileRequest) Validate() error {
	return dara.Validate(s)
}
