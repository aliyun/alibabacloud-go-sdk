// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledPreloadJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInsertWay(v string) *CreateScheduledPreloadJobRequest
	GetInsertWay() *string
	SetName(v string) *CreateScheduledPreloadJobRequest
	GetName() *string
	SetOssUrl(v string) *CreateScheduledPreloadJobRequest
	GetOssUrl() *string
	SetSiteId(v int64) *CreateScheduledPreloadJobRequest
	GetSiteId() *int64
	SetUrlList(v string) *CreateScheduledPreloadJobRequest
	GetUrlList() *string
}

type CreateScheduledPreloadJobRequest struct {
	// The method for uploading the preload file. Valid values are `Textbox` and `OSS`.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss
	InsertWay *string `json:"InsertWay,omitempty" xml:"InsertWay,omitempty"`
	// The name of the scheduled preload job.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The URL of the OSS file that contains the URLs to preload.
	//
	// example:
	//
	// https://xxxobject.oss-cn-reginon.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// The ID of the site. You can get this ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 190007158391808
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The list of URLs to preload. This parameter is used when `InsertWay` is set to `Textbox`.
	//
	// example:
	//
	// http://testurl.com/a.txt
	//
	// http://testurl.com/b.txt
	UrlList *string `json:"UrlList,omitempty" xml:"UrlList,omitempty"`
}

func (s CreateScheduledPreloadJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledPreloadJobRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadJobRequest) GetInsertWay() *string {
	return s.InsertWay
}

func (s *CreateScheduledPreloadJobRequest) GetName() *string {
	return s.Name
}

func (s *CreateScheduledPreloadJobRequest) GetOssUrl() *string {
	return s.OssUrl
}

func (s *CreateScheduledPreloadJobRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateScheduledPreloadJobRequest) GetUrlList() *string {
	return s.UrlList
}

func (s *CreateScheduledPreloadJobRequest) SetInsertWay(v string) *CreateScheduledPreloadJobRequest {
	s.InsertWay = &v
	return s
}

func (s *CreateScheduledPreloadJobRequest) SetName(v string) *CreateScheduledPreloadJobRequest {
	s.Name = &v
	return s
}

func (s *CreateScheduledPreloadJobRequest) SetOssUrl(v string) *CreateScheduledPreloadJobRequest {
	s.OssUrl = &v
	return s
}

func (s *CreateScheduledPreloadJobRequest) SetSiteId(v int64) *CreateScheduledPreloadJobRequest {
	s.SiteId = &v
	return s
}

func (s *CreateScheduledPreloadJobRequest) SetUrlList(v string) *CreateScheduledPreloadJobRequest {
	s.UrlList = &v
	return s
}

func (s *CreateScheduledPreloadJobRequest) Validate() error {
	return dara.Validate(s)
}
