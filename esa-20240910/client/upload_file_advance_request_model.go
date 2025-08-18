// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iUploadFileAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *UploadFileAdvanceRequest
	GetSiteId() *int64
	SetType(v string) *UploadFileAdvanceRequest
	GetType() *string
	SetUploadTaskName(v string) *UploadFileAdvanceRequest
	GetUploadTaskName() *string
	SetUrlObject(v io.Reader) *UploadFileAdvanceRequest
	GetUrlObject() io.Reader
}

type UploadFileAdvanceRequest struct {
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
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UploadFileAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadFileAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UploadFileAdvanceRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UploadFileAdvanceRequest) GetType() *string {
	return s.Type
}

func (s *UploadFileAdvanceRequest) GetUploadTaskName() *string {
	return s.UploadTaskName
}

func (s *UploadFileAdvanceRequest) GetUrlObject() io.Reader {
	return s.UrlObject
}

func (s *UploadFileAdvanceRequest) SetSiteId(v int64) *UploadFileAdvanceRequest {
	s.SiteId = &v
	return s
}

func (s *UploadFileAdvanceRequest) SetType(v string) *UploadFileAdvanceRequest {
	s.Type = &v
	return s
}

func (s *UploadFileAdvanceRequest) SetUploadTaskName(v string) *UploadFileAdvanceRequest {
	s.UploadTaskName = &v
	return s
}

func (s *UploadFileAdvanceRequest) SetUrlObject(v io.Reader) *UploadFileAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *UploadFileAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
