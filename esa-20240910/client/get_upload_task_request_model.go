// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetUploadTaskRequest
	GetSiteId() *int64
	SetUploadId(v int64) *GetUploadTaskRequest
	GetUploadId() *int64
}

type GetUploadTaskRequest struct {
	SiteId   *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	UploadId *int64 `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s GetUploadTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUploadTaskRequest) GoString() string {
	return s.String()
}

func (s *GetUploadTaskRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetUploadTaskRequest) GetUploadId() *int64 {
	return s.UploadId
}

func (s *GetUploadTaskRequest) SetSiteId(v int64) *GetUploadTaskRequest {
	s.SiteId = &v
	return s
}

func (s *GetUploadTaskRequest) SetUploadId(v int64) *GetUploadTaskRequest {
	s.UploadId = &v
	return s
}

func (s *GetUploadTaskRequest) Validate() error {
	return dara.Validate(s)
}
