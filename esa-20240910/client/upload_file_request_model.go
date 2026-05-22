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
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	UploadTaskName *string `json:"UploadTaskName,omitempty" xml:"UploadTaskName,omitempty"`
	// This parameter is required.
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
