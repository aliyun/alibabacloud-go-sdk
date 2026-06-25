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
	SetClientToken(v string) *FetchFileRequest
	GetClientToken() *string
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
	// A list of cloud phone instance IDs.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// A client-generated token, up to 100 characters long, that ensures the idempotency of the request.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The path of the file or folder to fetch from the cloud phone.
	//
	// This parameter is required.
	//
	// example:
	//
	// /data/a.txt
	SourceFilePath *string `json:"SourceFilePath,omitempty" xml:"SourceFilePath,omitempty"`
	// The endpoint for uploading files to OSS.
	//
	// > If the cloud phone and the destination OSS bucket are in the same region, you can use an internal endpoint to accelerate the transfer and avoid public network charges. For example, in the China (Hangzhou) region, use `oss-cn-hangzhou-internal.aliyuncs.com`. For a complete list of endpoints, see [OSS regions and endpoints](https://help.aliyun.com/document_detail/31837.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// oss-cn-hangzhou-internal.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// The type of storage service for the fetched file.
	//
	// > Currently, only Object Storage Service (OSS) is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
	// The destination URL in OSS.
	//
	// > The destination bucket name must be prefixed with `cloudphone-saved-bucket-`. For example, `cloudphone-saved-bucket-example`. You must also create a folder in the bucket to serve as the destination directory. The `UploadUrl` must follow the format: `oss://<bucket_name>/<folder_name>`.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://cloudphone-saved-bucket-example/received
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

func (s *FetchFileRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *FetchFileRequest) SetClientToken(v string) *FetchFileRequest {
	s.ClientToken = &v
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
