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
