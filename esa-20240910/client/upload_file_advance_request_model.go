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
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	UploadTaskName *string `json:"UploadTaskName,omitempty" xml:"UploadTaskName,omitempty"`
	// This parameter is required.
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
