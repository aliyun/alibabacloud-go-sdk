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
	// The website ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The type of the purge or prefetch task. Valid values:
	//
	// 	- **file*	- (default): purges the cache by file.
	//
	// 	- **preload**: prefetches the file.
	//
	// 	- **directory**: purges the cache by directory.
	//
	// 	- **ignoreParams**: purges the cache by URL with specified parameters ignored.
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
	// The OSS URL of the file that contains resources to be purged or prefetched.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xxxxx.oss-cn-shenzhen.aliyuncs.com/test_oss_file?Expires=1708659191&OSSAccessKeyId=**********&Signature=**********
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
