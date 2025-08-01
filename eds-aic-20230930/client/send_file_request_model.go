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
	// The IDs of the cloud phone instances.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	AutoInstall           *bool     `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	// The path to which you want to upload the pushed file in the cloud phone instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// /data
	SourceFilePath *string `json:"SourceFilePath,omitempty" xml:"SourceFilePath,omitempty"`
	// The name of the file uploaded from the Object Storage Service (OSS) to the cloud phone instance.
	//
	// >  If UploadType is set to OSS, you must specify TargetFileName. If TargetFileName is empty, the file uploaded from the OSS bucket to the cloud phone instance retains its original name. If TargetFileName is provided with a value, the uploaded file in the SourceFilePath directory uses the specified name (TargetFileName). If UploadType is set to DOWNLOAD_URL, TargetFileName does not take effect.
	//
	// example:
	//
	// test.txt
	TargetFileName *string `json:"TargetFileName,omitempty" xml:"TargetFileName,omitempty"`
	// The endpoint of the OSS bucket in which the file is stored.
	//
	// >  Set the value to an internal endpoint when the cloud phone instance and the OSS bucket are in the same region to improve transfer speed without incurring public traffic fees. Sample endpoint: `oss-cn-hangzhou-internal.aliyuncs.com`. For more information, see [OSS regions and endpoints](https://help.aliyun.com/document_detail/31837.html).
	//
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// The storage type of the file that you want to upload.
	//
	// 	- Set the value to OSS.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
	// The OSS URL of the file.
	//
	// >  The OSS bucket name must start with "cloudphone-saved-bucket-", for example, "cloudphone-saved-bucket-example". You must also create an OSS directory to store the backup data. Set the value for UploadUrl in this format: oss://\\<BucketName>/\\<OSSDirectoryName>\\<FileName>.
	//
	// This parameter is required.
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
