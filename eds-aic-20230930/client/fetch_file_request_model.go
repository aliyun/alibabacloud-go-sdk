// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIdList(v []*string) *FetchFileRequest
	GetAndroidInstanceIdList() []*string
	SetSourceFilePath(v string) *FetchFileRequest
	GetSourceFilePath() *string
	SetUploadEndpoint(v string) *FetchFileRequest
	GetUploadEndpoint() *string
	SetUploadType(v string) *FetchFileRequest
	GetUploadType() *string
	SetUploadUrl(v string) *FetchFileRequest
	GetUploadUrl() *string
}

type FetchFileRequest struct {
	// The IDs of the cloud phone instances.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// The path to the file that you want to pull from the cloud phone instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// /data/a.txt
	SourceFilePath *string `json:"SourceFilePath,omitempty" xml:"SourceFilePath,omitempty"`
	// The endpoint of the OSS bucket in which you want to store the pulled file.
	//
	// >  Set the value to an internal endpoint when the cloud phone instance and the OSS bucket are in the same region to improve upload speed without incurring public traffic fees. Sample endpoint: `oss-cn-hangzhou-internal.aliyuncs.com`. For more information, see [OSS regions and endpoints](https://help.aliyun.com/document_detail/31837.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// The type of the storage service.
	//
	// >  Currently, only OSS is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
	// The OSS URL of the pulled file.
	//
	// >  The OSS bucket name must start with "cloudphone-saved-bucket-", for example, "cloudphone-saved-bucket-example". You must also create an OSS directory to store the backup data. Set the value for UploadUrl in this format: oss://\\<BucketName>/\\<OSSDirectoryName>.
	//
	// This parameter is required.
	UploadUrl *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
}

func (s FetchFileRequest) String() string {
	return dara.Prettify(s)
}

func (s FetchFileRequest) GoString() string {
	return s.String()
}

func (s *FetchFileRequest) GetAndroidInstanceIdList() []*string {
	return s.AndroidInstanceIdList
}

func (s *FetchFileRequest) GetSourceFilePath() *string {
	return s.SourceFilePath
}

func (s *FetchFileRequest) GetUploadEndpoint() *string {
	return s.UploadEndpoint
}

func (s *FetchFileRequest) GetUploadType() *string {
	return s.UploadType
}

func (s *FetchFileRequest) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *FetchFileRequest) SetAndroidInstanceIdList(v []*string) *FetchFileRequest {
	s.AndroidInstanceIdList = v
	return s
}

func (s *FetchFileRequest) SetSourceFilePath(v string) *FetchFileRequest {
	s.SourceFilePath = &v
	return s
}

func (s *FetchFileRequest) SetUploadEndpoint(v string) *FetchFileRequest {
	s.UploadEndpoint = &v
	return s
}

func (s *FetchFileRequest) SetUploadType(v string) *FetchFileRequest {
	s.UploadType = &v
	return s
}

func (s *FetchFileRequest) SetUploadUrl(v string) *FetchFileRequest {
	s.UploadUrl = &v
	return s
}

func (s *FetchFileRequest) Validate() error {
	return dara.Validate(s)
}
