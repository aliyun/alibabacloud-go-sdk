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
	// This parameter is required.
	InsertWay *string `json:"InsertWay,omitempty" xml:"InsertWay,omitempty"`
	// This parameter is required.
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// This parameter is required.
	SiteId  *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
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
