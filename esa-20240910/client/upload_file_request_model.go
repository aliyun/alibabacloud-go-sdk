// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *UploadFileRequest
	GetSiteId() *int64
	SetType(v string) *UploadFileRequest
	GetType() *string
	SetUploadTaskName(v string) *UploadFileRequest
	GetUploadTaskName() *string
	SetUrl(v string) *UploadFileRequest
	GetUrl() *string
}

type UploadFileRequest struct {
	// The site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The type of the refresh or prefetch task. Valid values:
	//
	// - **file*	- (default): file refresh.
	//
	// - **preload**: file prefetch.
	//
	// - **directory**: directory refresh.
	//
	// - **ignoreParams**: parameter-ignored refresh.
	//
	// This parameter is required.
	//
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The name of the upload task.
	//
	// This parameter is required.
	//
	// example:
	//
	// purge_task_2024_11_11
	UploadTaskName *string `json:"UploadTaskName,omitempty" xml:"UploadTaskName,omitempty"`
	// The URL of the refresh or prefetch file stored in OSS. The Url parameter accepts URLs in two formats:
	//
	// - Transit URL (recommended): automatically generated through the file transfer feature of the ESA console or SDK.
	//
	// - OSS pre-signed HTTPS URL: generated after you upload the file to an authorized OSS bucket. The isFileTransferUrl field specifies whether to use the transit URL mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://XXXXXX.oss-cn-hangzhou.aliyuncs.com/{prefix}_{account_uid}_{hash}
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UploadFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadFileRequest) GoString() string {
	return s.String()
}

func (s *UploadFileRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UploadFileRequest) GetType() *string {
	return s.Type
}

func (s *UploadFileRequest) GetUploadTaskName() *string {
	return s.UploadTaskName
}

func (s *UploadFileRequest) GetUrl() *string {
	return s.Url
}

func (s *UploadFileRequest) SetSiteId(v int64) *UploadFileRequest {
	s.SiteId = &v
	return s
}

func (s *UploadFileRequest) SetType(v string) *UploadFileRequest {
	s.Type = &v
	return s
}

func (s *UploadFileRequest) SetUploadTaskName(v string) *UploadFileRequest {
	s.UploadTaskName = &v
	return s
}

func (s *UploadFileRequest) SetUrl(v string) *UploadFileRequest {
	s.Url = &v
	return s
}

func (s *UploadFileRequest) Validate() error {
	return dara.Validate(s)
}
