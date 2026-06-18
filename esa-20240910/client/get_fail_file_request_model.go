// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFailFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetFailFileRequest
	GetSiteId() *int64
	SetUploadId(v int64) *GetFailFileRequest
	GetUploadId() *int64
}

type GetFailFileRequest struct {
	// The site ID. You can obtain this value by calling the [ListSites](~~ListSites~~) operation.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The ID of the file upload task.
	//
	// example:
	//
	// 1593805857882113
	UploadId *int64 `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s GetFailFileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFailFileRequest) GoString() string {
	return s.String()
}

func (s *GetFailFileRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetFailFileRequest) GetUploadId() *int64 {
	return s.UploadId
}

func (s *GetFailFileRequest) SetSiteId(v int64) *GetFailFileRequest {
	s.SiteId = &v
	return s
}

func (s *GetFailFileRequest) SetUploadId(v int64) *GetFailFileRequest {
	s.UploadId = &v
	return s
}

func (s *GetFailFileRequest) Validate() error {
	return dara.Validate(s)
}
