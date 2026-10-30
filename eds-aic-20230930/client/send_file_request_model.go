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
	// The list of instances.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// Specifies whether to automatically install the uploaded application.
	//
	// example:
	//
	// true
	AutoInstall *bool `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. The token can be up to 100 characters in length.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The MD5 value of the input file.
	FileMd5 *string `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	// The upload path inside the cloud phone to which the file is uploaded.
	//
	// > If UploadType is OSS or OSS_BRIDGED, the upload path SourceFilePath inside the cloud phone must be a directory, such as /sdcard/Download. If UploadType is DOWNLOAD_URL, the upload path SourceFilePath inside the cloud phone must be a file, such as /sdcard/Download/MyFile.txt.
	//
	// >
	//
	// This parameter is required.
	//
	// example:
	//
	// /sdcard/Download
	SourceFilePath *string `json:"SourceFilePath,omitempty" xml:"SourceFilePath,omitempty"`
	// The name of the file uploaded from OSS to the cloud phone.
	//
	// > If UploadType is OSS or OSS_BRIDGED, the TargetFileName parameter is optional. If TargetFileName is empty, the original file name is retained when the OSS file is uploaded to the cloud phone. If TargetFileName is not empty, the file name in the SourceFilePath directory on the cloud phone is set to the value of TargetFileName. If UploadType is DOWNLOAD_URL, the TargetFileName parameter does not take effect.
	//
	// >
	//
	// example:
	//
	// test.txt
	TargetFileName *string `json:"TargetFileName,omitempty" xml:"TargetFileName,omitempty"`
	// The service domain name for uploading files to OSS. This parameter is required when UploadType is set to OSS or OSS_BRIDGED.
	//
	// > If the cloud phone and the OSS bucket are in the same region, you can use an internal endpoint to accelerate the transmission speed and avoid Internet traffic fees. For example, the internal endpoint for the China (Hangzhou) region is `oss-cn-hangzhou-internal.aliyuncs.com`. For more information about the mapping between regions and endpoints, see [Regions and endpoints](https://help.aliyun.com/document_detail/31837.html).
	//
	// example:
	//
	// oss-cn-hangzhou-internal.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// The storage type of the file to be uploaded.
	//
	// - OSS: The file to be uploaded is an OSS resource.
	//
	// - DOWNLOAD_URL: The file to be uploaded is a public download URL.
	//
	// - OSS_BRIDGED: The file is synchronized to OSS through a public download URL and then distributed to cloud phone instances over the OSS internal network.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
	// - If UploadType is OSS, UploadUrl is the cloud address of the file uploaded to OSS.
	//
	// > The OSS bucket name must start with the prefix "cloudphone-saved-bucket-", such as "cloudphone-saved-bucket-example". You must create an OSS folder as the backup directory. The format of UploadUrl is "oss://" + "Bucket name" + "/" + "Folder name" + "Uploaded file name".
	//
	// - If UploadType is DOWNLOAD or OSS_BRIDGED, UploadUrl is the public download URL of the file.
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
